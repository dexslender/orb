package commands

import "testing"

func TestBar(t *testing.T) {
	bar := GenUsageBar(50, 25, 15)
	t.Log(bar)
}
