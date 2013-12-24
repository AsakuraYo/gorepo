package sched

import "testing"

func TestHomeDir(t *testing.T) {
    homeDir, err := HomeDir()
    if len(homeDir) == 0 && err == nil {
        t.Error("Logical error")
    } else if len(homeDir) != 0 && err != nil {
        t.Error("Logical error")
    }
}
