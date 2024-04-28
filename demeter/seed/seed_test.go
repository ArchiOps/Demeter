package seed

import (
	"testing"
)

func TestNewCountSeed(t *testing.T) {

    type Model struct {
        Index int
        Super string
    }

    var test Model
	seed := NewCountSeed(&test, 0)
	t.Logf("Created seed: %v", seed)

    props := seed.ReadOne()
    t.Logf("With struct: %v %v", seed.meta.Struct, props)
}
