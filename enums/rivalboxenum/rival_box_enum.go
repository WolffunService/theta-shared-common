package rivalboxenum

type RivalBoxType int // enum:"RivalBoxType"

const (
	RivalBox RivalBoxType = iota + 1
	BigBox
	MegaBox
	EventBox

	_endRivalBoxType
)

func (r RivalBoxType) IsValid() bool {
	return r > 0 && r < _endRivalBoxType
}

var _boxStr = map[RivalBoxType]string{
	RivalBox: "rival_box",
	BigBox:   "big_box",
	MegaBox:  "mega_box",
	EventBox: "event_box",
}

func (r RivalBoxType) String() string {
	return _boxStr[r]
}
