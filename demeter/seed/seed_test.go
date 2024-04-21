package seed

import (
	"testing"
)

func TestNewCountSeed(t *testing.T) {

    test := struct {
        index   int
        super   string
    }{}

    seed := NewCountSeed(test, 0)
    t.Logf("Created seed: %v", seed)
}
