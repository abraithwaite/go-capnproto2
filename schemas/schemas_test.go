package schemas_test

import (
	"testing"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/schemas"
	"zombiezen.com/go/capnproto2/std/capnp/schema"
)

func TestDefaultFind(t *testing.T) {
	if s := schemas.Find(0xdeadbeef); s != nil {
		t.Errorf("schemas.Find(0xdeadbeef) = %d-byte slice; want nil", len(s))
	}
	s := schemas.Find(capnp.Package)
	if s == nil {
		t.Fatalf("schemas.Find(%#x) = nil", capnp.Package)
	}
	msg, err := capnp.Unmarshal(s)
	if err != nil {
		t.Fatalf("capnp.Unmarshal(schemas.Find(%#x)) error: %v", capnp.Package, err)
	}
	req, err := schema.ReadRootCodeGeneratorRequest(msg)
	if err != nil {
		t.Fatalf("ReadRootCodeGeneratorRequest error: %v", err)
	}
	var es capnp.ErrorSet
	nodes := req.Nodes(&es)
	if es != nil {
		t.Fatalf("req.Nodes() error: %v", es)
	}
	for i := 0; i < nodes.Len(); i++ {
		n := nodes.At(i)
		if n.Id() == capnp.Package {
			// Found
			if n.Which() != schema.Node_Which_annotation {
				t.Errorf("found node %#x which = %v; want annotation", capnp.Package, n.Which())
			}
			return
		}
	}
	t.Fatalf("could not find node %#x in registry", capnp.Package)
}

func TestNotFound(t *testing.T) {
	reg := new(schemas.Registry)
	_, err := reg.Find(0)
	if err == nil {
		t.Error("new(schemas.Registry).Find(0) = nil; want not found error")
	}
	if !schemas.IsNotFound(err) {
		t.Errorf("new(schemas.Registry).Find(0) = %v; want not found error", err)
	}
}
