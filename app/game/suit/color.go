package suit

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

var (
	ColorRed       = Color{Key: "red", Name: "Red", Code: "r"}
	ColorBlack     = Color{Key: "black", Name: "Black", Code: "b"}
	ColorGreen     = Color{Key: "green", Name: "Green", Code: "g"}
	ColorBlue      = Color{Key: "blue", Name: "Blue", Code: "u"}
	ColorColorless = Color{Key: "colorless", Name: "Colorless", Code: "x"}
	ColorUnknown   = Color{Key: "unknown", Name: "Unknown", Code: "?"}

	AllColors = Colors{ColorRed, ColorBlack, ColorGreen, ColorBlue, ColorColorless, ColorUnknown}
)

type Color struct {
	Key         string
	Name        string
	Description string
	Icon        string

	Code string
}

func (c Color) String() string {
	return c.Code
}

func (c Color) NameSafe() string {
	if c.Name != "" {
		return c.Name
	}
	return c.String()
}

func (c Color) Matches(xx Color) bool {
	return c.Key == xx.Key
}

func (c Color) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(c.Code, false), nil
}

func (c *Color) UnmarshalJSON(data []byte) error {
	key, err := util.FromJSONString(data)
	if err != nil {
		return err
	}
	*c = AllColors.GetByCode(key, nil)
	return nil
}

func (c Color) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(c.Code, start)
}

func (c *Color) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*c = AllColors.GetByCode(key, nil)
	return nil
}

func (c Color) Value() (driver.Value, error) {
	return c.Code, nil
}

func (c *Color) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, err := util.Cast[string](converted); err == nil {
			*c = AllColors.GetByCode(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan Color enum from value [%v]", value)
}

func ColorParse(logger util.Logger, keys ...string) Colors {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) Color {
		return AllColors.GetByCode(x, logger)
	})
}

type Colors []Color

func (c Colors) Keys() []string {
	return lo.Map(c, func(x Color, _ int) string {
		return x.Key
	})
}

func (c Colors) Strings() []string {
	return lo.Map(c, func(x Color, _ int) string {
		return x.String()
	})
}

func (c Colors) NamesSafe() []string {
	return lo.Map(c, func(x Color, _ int) string {
		return x.NameSafe()
	})
}

func (c Colors) Help() string {
	return "Available color options: [" + util.StringJoin(c.Strings(), ", ") + "]"
}

func (c Colors) Get(key string, logger util.Logger) Color {
	for _, value := range c {
		if strings.EqualFold(value.Code, key) {
			return value
		}
	}
	if key == "" {
		return ColorUnknown
	}
	msg := fmt.Sprintf("unable to find [Color] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return ColorUnknown
}

func (c Colors) GetByName(name string, logger util.Logger) Color {
	for _, value := range c {
		if strings.EqualFold(value.Name, name) {
			return value
		}
	}
	if name == "" {
		return ColorUnknown
	}
	msg := fmt.Sprintf("unable to find [Color] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return ColorUnknown
}

func (c Colors) GetByCode(input string, logger util.Logger) Color {
	for _, value := range c {
		if value.Code == input {
			return value
		}
	}
	if input == "" {
		return ColorUnknown
	}
	msg := fmt.Sprintf("unable to find [Color] with Code [%s]", input)
	if logger != nil {
		logger.Warn(msg)
	}
	return ColorUnknown
}

func (c Colors) Random() Color {
	return util.RandomElement(c)
}
