package card

import (
	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Color struct {
	Key  string
	Name string
}

func (c *Color) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(c.Key, false), nil
}

func (c *Color) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid color key [%s]", string(data))
	}
	curr, ok := colorsByKey[key]
	if !ok {
		return errors.Wrapf(err, "Incorrect color key [%s]", string(data))
	}
	c.ReplaceFrom(curr)
	return nil
}

func (c *Color) ReplaceFrom(x *Color) {
	c.Key = x.Key
	c.Name = x.Name
}

var (
	ColorRed       = &Color{Key: "R", Name: "Red"}
	ColorBlack     = &Color{Key: "B", Name: "Black"}
	ColorGreen     = &Color{Key: "G", Name: "Green"}
	ColorBlue      = &Color{Key: "U", Name: "Blue"}
	ColorColorless = &Color{Key: "X", Name: "Colorless"}
	ColorUnknown   = &Color{Key: "?", Name: "Unknown"}

	colorsByKey = map[string]*Color{
		ColorRed.Key:       ColorRed,
		ColorBlack.Key:     ColorBlack,
		ColorGreen.Key:     ColorGreen,
		ColorBlue.Key:      ColorBlue,
		ColorColorless.Key: ColorColorless,
		ColorUnknown.Key:   ColorUnknown,
	}
)
