package capnp_test

import (
	"testing"

	capnp "zombiezen.com/go/capnproto2"
	air "zombiezen.com/go/capnproto2/internal/aircraftlib"
)

func TestCopy(t *testing.T) {
	// Create object
	// Fill in fields
	// Set fields to null on original object
	// Do copy
	// Compare serialized bytes && size
	// Compare individual fields
	_, seg, err := capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		panic(err)
	}

	z, err := air.NewRootAircraft(seg)
	if err != nil {
		panic(err)
	}

	b7, err := z.NewB737()
	if err != nil {
		panic(err)
	}

	base, err := b7.NewBase()
	if err != nil {
		panic(err)
	}

	base.SetName("asdf")
	base.SetRating(int64(12381))
	base.SetCanFly(false)
	base.SetMaxSpeed(float64(3410.111))
	homes, err := base.NewHomes(50)

	// Something big for more memory usage
	for i := 0; i < 50; i++ {
		homes.Set(i, air.Airport_jfk)
	}

	b, err := z.Segment().Message().Marshal()
	if err != nil {
		panic(err)
	}

	_, seg, err = capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		panic(err)
	}

	// Nil pointer, of a different type
	z.SetVoid()

	y, err := z.Copy(seg)
	if err != nil {
		panic(err)
	}

	b2, err := y.Segment().Message().Marshal()
	if err != nil {
		panic(err)
	}
	if len(b) <= len(b2) {
		t.Errorf("Copied buffer is not smaller")
	}
	if y.HasB737() {
		t.Errorf("New struct should not have old union")
	}

}

func TestCopyUnions(t *testing.T) {
	_, seg, err := capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create object
	z, err := air.NewRootZ(seg)
	if err != nil {
		panic(err)
	}

	ac, _ := z.NewAircraft()
	b7, _ := ac.NewB737()
	base, _ := b7.NewBase()
	homes, _ := base.NewHomes(50)

	// Something big for more memory usage
	for i := 0; i < 50; i++ {
		homes.Set(i, air.Airport_jfk)
	}

	// Marshal original bytes
	b, err := z.Segment().Message().Marshal()
	if err != nil {
		panic(err)
	}

	_, seg, err = capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		panic(err)
	}

	// set fields / union type
	z.SetF64(11.11)

	y, err := z.Copy(seg)
	if err != nil {
		panic(err)
	}

	b2, err := y.Segment().Message().Marshal()
	if err != nil {
		panic(err)
	}
	if len(b) <= len(b2) {
		t.Errorf("Copied buffer is not smaller")
	}
	if y.Which() != air.Z_Which_f64 {
		t.Errorf("New struct should have new union type. old: %v new: %v", z.Which(), y.Which())
	}
}