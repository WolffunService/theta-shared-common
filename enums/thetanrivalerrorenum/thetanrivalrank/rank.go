package thetanrivalrank

type Rank int

const (
	_Start Rank = iota - 1
	Tutorial
	Recruit
	Private
	Bronze
	Silver
	Golden
	Platinum
	Diamond
	Master
	Champion
	SuperChampion
	_End
)
const (
	MinRank = Tutorial
	MaxRank = SuperChampion
)

var _rankName = map[Rank]string{
	Tutorial:      "TUTORIAL",
	Recruit:       "RECRUIT",
	Private:       "PRIVATE",
	Bronze:        "BRONZE",
	Silver:        "SILVER",
	Golden:        "GOLDEN",
	Platinum:      "PLATINUM",
	Diamond:       "DIAMOND",
	Master:        "MASTER",
	Champion:      "CHAMPION",
	SuperChampion: "SUPER_CHAMPION",
}

func (r Rank) String() string {
	return _rankName[r]
}
func (r Rank) IsValid() bool {
	return r > _Start && r < _End
}

func (Rank) Min() Rank {
	return _Start + 1
}

func (Rank) Max() Rank {
	return _End - 1
}
