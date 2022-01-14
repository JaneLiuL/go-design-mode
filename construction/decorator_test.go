package construction

import (
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	now := time.Now()
	Call()
	duration := time.Now().Sub(now)
	if duration < 5*time.Second {
		t.Fatal("at least should large than 5 seconds")
	}
}
