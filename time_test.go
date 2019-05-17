package msgpack

import (
	"github.com/Limard/msgpack/codes"
	"testing"
)

func Test111(t *testing.T) {
	t.Logf("%X", codes.FixedArrayLow|2)
}
