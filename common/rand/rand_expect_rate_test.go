package rand

import (
	"sync"
	"testing"
)

func TestPickItemByExpectRate(t *testing.T) {
	poolItems := map[string]int{
		"apple":     1,
		"banana":    1,
		"raspberry": 1,
	}

	expectedRate := map[string]float64{
		"apple":     0.5,
		"banana":    0.3,
		"raspberry": 0.2,
	}

	item, err := PickItemByExpectRate(poolItems, expectedRate)
	if err != nil {
		t.Error(err)
	}

	t.Log("Pick item is", item)

	if item != "apple" {
		t.FailNow()
	}
}

func TestPickItemByExpectRateStats(t *testing.T) {
	poolItems := map[string]int{
		"apple":     0,
		"banana":    0,
		"raspberry": 0,
	}

	expectedRate := map[string]float64{
		"apple":     0.5, // 50%
		"banana":    0.3, // 30%
		"raspberry": 0.2, // 20%
	}

	numTest := 1_000_000
	mu := sync.Mutex{}
	count := 0

	for i := 0; i < numTest; i++ {
		item, err := PickItemByExpectRate(poolItems, expectedRate)
		if err != nil {
			t.Error(err)
		}

		mu.Lock()
		poolItems[item]++
		count++
		mu.Unlock()
	}

	t.Log("Total rand:", count)
	for item, appeared := range poolItems {
		t.Logf("Item %s appeared %d times, expect rate: %.2f%%", item, appeared, expectedRate[item]*100)
	}
}

func TestPickMultipleItemsByExpectRate(t *testing.T) {
	poolItems := map[string]int{
		"apple":     1,
		"banana":    1,
		"raspberry": 1,
	}

	expectedRate := map[string]float64{
		"apple":     0.5,
		"banana":    0.3,
		"raspberry": 0.2,
	}

	items, err := PickMultipleItemsByExpectRate(poolItems, expectedRate, 2)
	if err != nil {
		t.Error(err)
	}

	t.Log("Pick item is:", items)

	if items[0] != "apple" || items[1] != "banana" {
		t.FailNow()
	}
}

func TestPickMultipleItemsByExpectRateStats(t *testing.T) {
	poolItems := map[string]int{
		"apple":     0,
		"banana":    0,
		"raspberry": 0,
	}

	expectedRate := map[string]float64{
		"apple":     0.5, // 50%
		"banana":    0.3, // 30%
		"raspberry": 0.2, // 20%
	}

	numTest := 1_000_000
	mu := sync.Mutex{}
	count := 0

	for i := 0; i < numTest; i++ {
		items, err := PickMultipleItemsByExpectRate(poolItems, expectedRate, 1)
		if err != nil {
			t.Error(err)
		}

		for _, item := range items {
			mu.Lock()
			poolItems[item]++
			count++
			mu.Unlock()
		}
	}

	t.Log("Total rand:", count)
	for item, appeared := range poolItems {
		t.Logf("Item %s appeared %d times, expect rate: %.2f%%", item, appeared, expectedRate[item]*100)
	}
}
