package suit

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Color uint8

const (
	ColorRed Color = iota + 1
	ColorBlack
	ColorGreen
	ColorBlue
	ColorColorless
	ColorUnknown
)

var colorMap = map[string]Color{
	"r": ColorRed,
	"b": ColorBlack,
	"g": ColorGreen,
	"u": ColorBlue,
	"x": ColorColorless,
	"?": ColorUnknown,
}

func (c Color) Key() string {
	switch c {
	case ColorRed:
		return "R"
	case ColorBlack:
		return "B"
	case ColorGreen:
		return "G"
	case ColorBlue:
		return "U"
	case ColorColorless:
		return "X"
	case ColorUnknown:
		return "?"
	default:
		return "?"
	}
}

func (c Color) Name() string {
	switch c {
	case ColorRed:
		return "Red"
	case ColorBlack:
		return "Black"
	case ColorGreen:
		return "Green"
	case ColorBlue:
		return "Blue"
	case ColorColorless:
		return "Colorless"
	case ColorUnknown:
		return util.KeyUnknown
	default:
		return fmt.Sprintf("Unknown: %d", c)
	}
}

func (c *Color) MarshalText() ([]byte, error) {
	return []byte(c.Key()), nil
}

func (c *Color) UnmarshalText(data []byte) error {
	var err error
	*c, err = ParseColor(string(data))
	return err
}

func ParseColor(key string) (Color, error) {
	key = strings.TrimSpace(strings.ToLower(key))
	ret, ok := colorMap[key]
	if !ok {
		return Color(0), errors.Errorf("invalid color [%s]", key)
	}
	return ret, nil
}
