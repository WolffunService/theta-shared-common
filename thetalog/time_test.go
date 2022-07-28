package thetalog

import (
	"fmt"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	logger := NewBizLogger("bizname")
	for i := 0; i < b.N; i++ {
		logger.Err(fmt.Errorf("error ")).Op("main").Var("z", 1).Msg("hahahaaaaa")
	}
}
