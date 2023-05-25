package smcenum

type ScanOffset uint64

const (
	SO_Small  ScanOffset = 3 // default is 3
	SO_Medium ScanOffset = 15
	// when adding please add to GetAllOffsets function too
)

func GetAllOffsets() []ScanOffset {
	return []ScanOffset{SO_Small, SO_Medium}
}
