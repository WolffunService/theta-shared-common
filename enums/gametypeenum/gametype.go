package gametypeenum

type TypeGame int

const (
	ThetanArena TypeGame = iota
	ThetanRivals
	ThetanUGC
	end

	// HeroesStrikeTheta is alias of Thetan Arena
	HeroesStrikeTheta = ThetanArena
)

var keys = map[TypeGame]string{
	ThetanArena:  "thetan-arena",
	ThetanRivals: "thetan-rivals",
	ThetanUGC:    "thetan-ugc",
}

var identifies = map[TypeGame]string{
	ThetanArena:  "TA",
	ThetanRivals: "TR",
	ThetanUGC:    "UG",
}

var names = map[TypeGame]string{
	ThetanArena:  "Thetan Arena",
	ThetanRivals: "Thetan Rivals",
	ThetanUGC:    "Thetan UGC",
}

func (t TypeGame) IsValid() bool {
	return t >= 0 && t < end
}

func (t TypeGame) Key() string {
	if key, found := keys[t]; found {
		return key
	}

	return ""
}

func (t TypeGame) Identify() string {
	if identify, found := identifies[t]; found {
		return identify
	}

	return "UK" // Unknown
}

func (t TypeGame) String() string {
	if name, found := names[t]; found {
		return name
	}

	return "Unknown"
}
