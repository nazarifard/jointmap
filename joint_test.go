package jointmap

import (
	"testing"

	"github.com/nazarifard/gomap"
)

func TestJointMap(t *testing.T) {
	a := gomap.New[string, float32]()
	a.Set("pi", 3.14159)
	a.Set("gamma", 0.57721)

	b := gomap.New[float32, rune]()
	b.Set(3.14159, 'π')
	b.Set(0.57721, 'γ')

	ab := New(a, b)

	for it := ab.Iterator(); it.Next(); {
		print(it.Key(), ":", string(it.Value()), "  ")
	}
	print("\n\n")
}
