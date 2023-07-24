package card

const (
	TwoSpades   = "🂢"
	TwoHearts   = "🂲"
	TwoDiamonds = "🃂"
	TwoClubs    = "🃒"

	ThreeSpades   = "🂣"
	ThreeHearts   = "🂳"
	ThreeDiamonds = "🃃"
	ThreeClubs    = "🃓"

	FourSpades   = "🂤"
	FourHearts   = "🂴"
	FourDiamonds = "🃄"
	FourClubs    = "🃔"

	FiveSpades   = "🂥"
	FiveHearts   = "🂵"
	FiveDiamonds = "🃅"
	FiveClubs    = "🃕"

	SixSpades   = "🂦"
	SixHearts   = "🂶"
	SixDiamonds = "🃆"
	SixClubs    = "🃖"

	SevenSpades   = "🂧"
	SevenHearts   = "🂷"
	SevenDiamonds = "🃇"
	SevenClubs    = "🃗"

	EightSpades   = "🂨"
	EightHearts   = "🂸"
	EightDiamonds = "🃈"
	EightClubs    = "🃘"

	NineSpades   = "🂩"
	NineHearts   = "🂹"
	NineDiamonds = "🃉"
	NineClubs    = "🃙"

	TenSpades   = "🂪"
	TenHearts   = "🂺"
	TenDiamonds = "🃊"
	TenClubs    = "🃚"

	JackSpades   = "🂫"
	JackHearts   = "🂻"
	JackDiamonds = "🃋"
	JackClubs    = "🃛"

	QueenSpades   = "🂭"
	QueenHearts   = "🂽"
	QueenDiamonds = "🃍"
	QueenClubs    = "🃝"

	KingSpades   = "🂮"
	KingHearts   = "🂾"
	KingDiamonds = "🃎"
	KingClubs    = "🃞"

	AceSpades   = "🂡"
	AceHearts   = "🂱"
	AceDiamonds = "🃁"
	AceClubs    = "🃑"

	CardBack = "🂠"
	Joker    = "🃟"
)

var UnicodeCards = map[string]string{
	"2S": TwoSpades, "2H": TwoHearts, "2D": TwoDiamonds, "2C": TwoClubs,
	"3S": ThreeSpades, "3H": ThreeHearts, "3D": ThreeDiamonds, "3C": ThreeClubs,
	"4S": FourSpades, "4H": FourHearts, "4D": FourDiamonds, "4C": FourClubs,
	"5S": FiveSpades, "5H": FiveHearts, "5D": FiveDiamonds, "5C": FiveClubs,
	"6S": SixSpades, "6H": SixHearts, "6D": SixDiamonds, "6C": SixClubs,
	"7S": SevenSpades, "7H": SevenHearts, "7D": SevenDiamonds, "7C": SevenClubs,
	"8S": EightSpades, "8H": EightHearts, "8D": EightDiamonds, "8C": EightClubs,
	"9S": NineSpades, "9H": NineHearts, "9D": NineDiamonds, "9C": NineClubs,
	"XS": TenSpades, "XH": TenHearts, "XD": TenDiamonds, "XC": TenClubs,
	"JS": JackSpades, "JH": JackHearts, "JD": JackDiamonds, "JC": JackClubs,
	"QS": QueenSpades, "QH": QueenHearts, "QD": QueenDiamonds, "QC": QueenClubs,
	"KS": KingSpades, "KH": KingHearts, "KD": KingDiamonds, "KC": KingClubs,
	"AS": AceSpades, "AH": AceHearts, "AD": AceDiamonds, "AC": AceClubs,
}
