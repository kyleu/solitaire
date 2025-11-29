package stock

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/solitaire/app/util"
)

//nolint:lll
var (
	DealToWaste               = DealTo{Key: "waste", Name: "Waste"}
	DealToWasteOrPairManually = DealTo{Key: "waste-or-pair-manually", Name: "Waste Or Pair Manually"}
	DealToTableau             = DealTo{Key: "tableau", Name: "Tableau"}
	DealToTableauFirstSet     = DealTo{Key: "tableau-first-set", Name: "Tableau First Set"}
	DealToTableauIfNoneEmpty  = DealTo{Key: "tableau-if-none-empty", Name: "Tableau If None Empty"}
	DealToTableauEmpty        = DealTo{Key: "tableau-empty", Name: "Tableau Empty"}
	DealToTableauNonEmpty     = DealTo{Key: "tableau-non-empty", Name: "Tableau Non Empty"}
	DealToFoundation          = DealTo{Key: "foundation", Name: "Foundation"}
	DealToReserve             = DealTo{Key: "reserve", Name: "Reserve"}
	DealToManually            = DealTo{Key: "manually", Name: "Manually"}
	DealToNever               = DealTo{Key: "never", Name: "Never"}

	AllDealTos = DealTos{DealToWaste, DealToWasteOrPairManually, DealToTableau, DealToTableauFirstSet, DealToTableauIfNoneEmpty, DealToTableauEmpty, DealToTableauNonEmpty, DealToFoundation, DealToReserve, DealToManually, DealToNever}
)

type DealTo struct {
	Key         string
	Name        string
	Description string
	Icon        string
}

func (d DealTo) String() string {
	return d.Key
}

func (d DealTo) NameSafe() string {
	if d.Name != "" {
		return d.Name
	}
	return d.String()
}

func (d DealTo) Matches(xx DealTo) bool {
	return d.Key == xx.Key
}

func (d DealTo) MarshalJSON() ([]byte, error) {
	return util.ToJSONBytes(d.Key, false), nil
}

func (d *DealTo) UnmarshalJSON(data []byte) error {
	key, err := util.FromJSONString(data)
	if err != nil {
		return err
	}
	*d = AllDealTos.Get(key, nil)
	return nil
}

func (d DealTo) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	return enc.EncodeElement(d.Key, start)
}

func (d *DealTo) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var key string
	if err := dec.DecodeElement(&key, &start); err != nil {
		return err
	}
	*d = AllDealTos.Get(key, nil)
	return nil
}

func (d DealTo) Value() (driver.Value, error) {
	return d.Key, nil
}

func (d *DealTo) Scan(value any) error {
	if value == nil {
		return nil
	}
	if converted, err := driver.String.ConvertValue(value); err == nil {
		if str, err := util.Cast[string](converted); err == nil {
			*d = AllDealTos.Get(str, nil)
			return nil
		}
	}
	return errors.Errorf("failed to scan DealTo enum from value [%v]", value)
}

func DealToParse(logger util.Logger, keys ...string) DealTos {
	if len(keys) == 0 {
		return nil
	}
	return lo.Map(keys, func(x string, _ int) DealTo {
		return AllDealTos.Get(x, logger)
	})
}

type DealTos []DealTo

func (d DealTos) Keys() []string {
	return lo.Map(d, func(x DealTo, _ int) string {
		return x.Key
	})
}

func (d DealTos) Strings() []string {
	return lo.Map(d, func(x DealTo, _ int) string {
		return x.String()
	})
}

func (d DealTos) NamesSafe() []string {
	return lo.Map(d, func(x DealTo, _ int) string {
		return x.NameSafe()
	})
}

func (d DealTos) Help() string {
	return "Available deal to options: [" + util.StringJoin(d.Strings(), ", ") + "]"
}

func (d DealTos) Get(key string, logger util.Logger) DealTo {
	for _, value := range d {
		if strings.EqualFold(value.Key, key) {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [DealTo] with key [%s]", key)
	if logger != nil {
		logger.Warn(msg)
	}
	return DealTo{Key: "_error", Name: "error: " + msg}
}

func (d DealTos) GetByName(name string, logger util.Logger) DealTo {
	for _, value := range d {
		if strings.EqualFold(value.Name, name) {
			return value
		}
	}
	msg := fmt.Sprintf("unable to find [DealTo] with name [%s]", name)
	if logger != nil {
		logger.Warn(msg)
	}
	return DealTo{Key: "_error", Name: "error: " + msg}
}

func (d DealTos) Random() DealTo {
	return util.RandomElement(d)
}
