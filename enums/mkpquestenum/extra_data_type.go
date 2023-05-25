package mkpquestenum

import (
	"errors"
)

type ExtraDataType int

const (
	none ExtraDataType = iota
	ExtDataHeroTypeId
	ExtDataItemRarity
)

var (
	ErrUndefinedType = errors.New("Cannot parse this type")
)
