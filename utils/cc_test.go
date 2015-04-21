package utils

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCC(t *testing.T) {
	deny := time.Second * 1
	cc := NewCC(time.Millisecond)
	cc.Start()
	time.Sleep(deny)
	cc.Stop()
	if !strings.EqualFold(cc.GetUnitStr(), "ms") {
		t.Error("time unit not ms")
	}

	fmt.Println(cc.CostStr())
	fmt.Println(cc.CostFloatStr())
}
