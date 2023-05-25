package ivtmodel

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// InventoryChangeType enum use when create transaction with inventory
type InventoryChangeType uint32

const (
	ChangeIncrease InventoryChangeType = 1 + iota
	ChangeDecrease
)

type Number interface {
	constraints.Integer | constraints.Float
}

type InventoryChangeSystem[T Number] struct {
	kind         InventoryKind
	typ          InventoryType
	amount       T
	systemAmount int64
}

type IInventoryChangeSystem interface {
	Validate() error
	GetInventoryKind() InventoryKind
	GetInventoryType() InventoryType
	GetRealAmount() any
	GetSystemAmount() int64
}

func NewICSFromSystemCurrency[T Number](kind InventoryKind, typ InventoryType, sysAmount int64) InventoryChangeSystem[T] {
	amount := T(sysAmount / SystemDecimal)
	return InventoryChangeSystem[T]{kind: kind, typ: typ, amount: amount, systemAmount: sysAmount}
}

func NewInventoryChangeSystem[T Number](kind InventoryKind, typ InventoryType, amount T) InventoryChangeSystem[T] {
	systemAmount := int64(float64(amount) * SystemDecimal)
	return InventoryChangeSystem[T]{kind: kind, typ: typ, amount: amount, systemAmount: systemAmount}
}

// Validate must call this method before execute transaction
func (i InventoryChangeSystem[T]) Validate() error {
	switch i.kind {
	case IKBox, IKCosmeticProfile, IKCosmeticAddIn:
		if !IsInteger(i.amount) {
			return fmt.Errorf("Amount Type Invalid, %T , IK %v ", i.amount, i.kind)
		}
	}
	return nil
}

func (i InventoryChangeSystem[T]) GetInventoryKind() InventoryKind {
	return i.kind
}

func (i InventoryChangeSystem[T]) GetInventoryType() InventoryType {
	return i.typ
}

func (i InventoryChangeSystem[T]) GetRealAmount() any {
	return i.amount
}

func (i InventoryChangeSystem[Number]) GetSystemAmount() int64 {
	return int64(i.systemAmount)
}

func IsInteger(t any) bool {
	switch t.(type) {
	case int, int32, int64, uint, uint32, uint64:
		return true
	}
	return false
}

type InventoryChangeOptions struct {
	//Source *inventoryenum.ChangeSource
}

func ChangeOptions() *InventoryChangeOptions {
	return &InventoryChangeOptions{}
}

//func (opts *InventoryChangeOptions) SetSource(source inventoryenum.ChangeSource) {
//	opts.Source = &source
//}

func MergeInventoryChangeOptions(opts ...*InventoryChangeOptions) *InventoryChangeOptions {
	co := ChangeOptions()

	//for _, opt := range opts {
	//	if opt.Source != nil {
	//		co.Source = opt.Source
	//	}
	//}

	return co
}
