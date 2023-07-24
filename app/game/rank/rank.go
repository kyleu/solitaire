package rank

import (
	"github.com/pkg/errors"

	"github.com/kyleu/solitaire/app/util"
)

type Rank uint8

const (
	Two Rank = iota + 1
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Unknown
)

func (r Rank) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(r.Key(), false), nil
}

func (r *Rank) UnmarshalJSON(data []byte) error {
	var key string
	err := util.FromJSON(data, &key)
	if err != nil {
		return errors.Wrapf(err, "Invalid suit key [%s]", string(data))
	}
	*r, err = Parse(key)
	return err
}
