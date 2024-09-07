package test

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func TestStats(t *testing.T) {
	t.Log(cpu.Counts(true))
	pc, err := cpu.Percent(time.Second, false)
	if err != nil {
		return
	}
	t.Log(pc[0])
}
