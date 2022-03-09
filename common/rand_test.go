package common

import "testing"

func TestRandString(t *testing.T) {
	const numTests = 1000
	const length = 10

	for i := 0; i < numTests; i++ {
		randStr := RandString(length)
		if len(randStr) == length {
			t.Log(randStr)
		} else {
			t.Fail()
		}
	}
}
