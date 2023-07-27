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

func (r Rank) Previous() Rank {
	switch r {
	case Two:
		return Ace
	case Three:
		return Two
	case Four:
		return Three
	case Five:
		return Four
	case Six:
		return Five
	case Seven:
		return Six
	case Eight:
		return Seven
	case Nine:
		return Eight
	case Ten:
		return Nine
	case Jack:
		return Ten
	case Queen:
		return Jack
	case King:
		return Queen
	case Ace:
		return King
	default:
		return Unknown
	}
}

func (r Rank) Next() Rank {
	switch r {
	case Two:
		return Three
	case Three:
		return Four
	case Four:
		return Five
	case Five:
		return Six
	case Six:
		return Seven
	case Seven:
		return Eight
	case Eight:
		return Nine
	case Nine:
		return Ten
	case Ten:
		return Jack
	case Jack:
		return Queen
	case Queen:
		return King
	case King:
		return Ace
	case Ace:
		return Two
	default:
		return Unknown
	}
}
