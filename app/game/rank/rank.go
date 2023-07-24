package rank

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

func (r Rank) MarshalText() ([]byte, error) {
	return []byte(r.Key()), nil
}

func (r *Rank) UnmarshalText(data []byte) error {
	var err error
	*r, err = Parse(string(data))
	return err
}
