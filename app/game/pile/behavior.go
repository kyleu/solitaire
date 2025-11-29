package pile

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
	BehaviorStock      = Behavior{Key: "stock", Name: "Stock", Code: "s", Weight: 0}
	BehaviorWaste      = Behavior{Key: "waste", Name: "Waste", Code: "w", Weight: 1}
	BehaviorPocket     = Behavior{Key: "pocket", Name: "Pocket", Code: "k", Weight: 0}
	BehaviorFoundation = Behavior{Key: "foundation", Name: "Foundation", Code: "f", Weight: 4}
	BehaviorTableau    = Behavior{Key: "tableau", Name: "Tableau", Code: "t", Weight: 4}
	BehaviorCell       = Behavior{Key: "cell", Name: "Cell", Code: "c", Weight: 2}
	BehaviorReserve    = Behavior{Key: "reserve", Name: "Reserve", Code: "r", Weight: 2}
	BehaviorPyramid    = Behavior{Key: "pyramid", Name: "Pyramid", Code: "p", Weight: 5}

	AllBehaviors = Behaviors{BehaviorStock, BehaviorWaste, BehaviorPocket, BehaviorFoundation, BehaviorTableau, BehaviorCell, BehaviorReserve, BehaviorPyramid}
)

type Behavior struct {
	Key         string
	Name        string
	Description string
	Icon        string

	Code   string
	Weight int
}

func (b Behavior) String() string {
	return b.Key
}

func (b Behavior) NameSafe() string {
	if b.Name != "" {
		return b.Name
	}
	return b.String()
}

func (b Behavior) Matches(xx Behavior) bool {
	return b.Key == xx.Key
}

func (b Behavior) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(b.Key, false), nil
}

func (b *Behavior) UnmarshalJSON(data []byte) error {
	key, err := util.FromJSONString(data)
	if err != nil {
		return err
	}
	*b = AllBehaviors.Get(key, nil)
	return nil
}

func (b Behavior) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(b.Key, start)
}

func (b *Behavior) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*b = AllBehaviors.Get(key, nil)
	return nil
}

func (b Behavior) Value() (driver.Value, error) {
	return b.Key, nil
}

func (b *Behavior) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, err := util.Cast[string](converted); err == nil {
			*b = AllBehaviors.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan Behavior enum from value [%v]", value)
}

func BehaviorParse(logger util.Logger, keys ...string) Behaviors {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) Behavior {
		return AllBehaviors.Get(x, logger)
	})
}

type Behaviors []Behavior

func (b Behaviors) Keys() []string {
	return lo.Map(b, func(x Behavior, _ int) string {
		return x.Key
	})
}

func (b Behaviors) Strings() []string {
	return lo.Map(b, func(x Behavior, _ int) string {
		return x.String()
	})
}

func (b Behaviors) NamesSafe() []string {
	return lo.Map(b, func(x Behavior, _ int) string {
		return x.NameSafe()
	})
}

func (b Behaviors) Help() string {
	return "Available behavior options: [" + util.StringJoin(b.Strings(), ", ") + "]"
}

func (b Behaviors) Get(key string, logger util.Logger) Behavior {
	for _, value := range b {
		if strings.EqualFold(value.Key, key) {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Behavior] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return Behavior{Key: "_error", Name: "error: " + msg}
}

func (b Behaviors) GetByName(name string, logger util.Logger) Behavior {
	for _, value := range b {
		if strings.EqualFold(value.Name, name) {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Behavior] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return Behavior{Key: "_error", Name: "error: " + msg}
}

func (b Behaviors) GetByCode(input string, logger util.Logger) Behavior {
	for _, value := range b {
		if value.Code == input {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [Behavior] with Code [%s]", input)
	if logger != nil {
		logger.Warn(msg)
	}
	return Behavior{Key: "_error", Name: "error: " + msg}
}

func (b Behaviors) GetByWeight(input int) Behaviors {
	ret := b
	for _, value := range b {
		if value.Weight == input {
			ret = append(ret, value)
		}
	}
	return ret
}

func (b Behaviors) Random() Behavior {
	return util.RandomElement(b)
}
