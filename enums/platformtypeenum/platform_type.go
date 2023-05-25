package platformtypeenum

type PlatformType int32

const (
	ANDROID PlatformType = iota
	IOS
	PC
	MACOS
	end_enum
)

func (p PlatformType) IsValid() bool {
	return p >= 0 && p < end_enum
}
