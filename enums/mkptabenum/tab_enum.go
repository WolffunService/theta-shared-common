package mkptabenum

type MKPTabEnum int

const (
	_Start MKPTabEnum = iota
	Dashboard
	ThetanBox
	Buy
	Rent
	SpecialEvent
	Staking
	MKPQuest
	ConvertionProgram
	VestingSafe
	CreatorCorner
	Profile
	_End
)

func (t MKPTabEnum) IsValid() bool {
	return t > _Start && t < _End
}

func ForEach(fn func(tab MKPTabEnum)) {
	for i := _Start + 1; i < _End; i++ {
		fn(i)
	}
}
