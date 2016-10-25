package main

import (
	"errors"
	"fmt"
	"strings"

	"zombiezen.com/go/capnproto2"
	"zombiezen.com/go/capnproto2/std/capnp/schema"
)

type node struct {
	schema.Node
	pkg   string
	imp   string
	nodes []*node // only for file nodes
	Name  string
}

func (n *node) codeOrderFields() []field {
	fields := n.StructNode().Fields(nil)
	numFields := fields.Len()
	mbrs := make([]field, numFields)
	for i := 0; i < numFields; i++ {
		f := fields.At(i)
		fann := f.Annotations(nil)
		fname := f.Name(nil)
		fname = parseAnnotations(fann).Rename(fname)
		mbrs[f.CodeOrder()] = field{Field: f, Name: fname}
	}
	return mbrs
}

// DiscriminantOffset returns the byte offset of the struct union discriminant.
func (n *node) DiscriminantOffset() (uint32, error) {
	if n == nil {
		return 0, errors.New("discriminant offset called on nil node")
	}
	if n.Which() != schema.Node_Which_structNode {
		return 0, fmt.Errorf("discriminant offset called on %v node", n.Which())
	}
	return n.StructNode().DiscriminantOffset() * 2, nil
}

func (n *node) shortDisplayName() string {
	dn := n.DisplayName(nil)
	return dn[n.DisplayNamePrefixLength():]
}

// String returns the node's display name.
func (n *node) String() string {
	return displayName(n)
}

func displayName(n interface {
	DisplayName(*capnp.ErrorSet) string
}) string {
	dn := n.DisplayName(nil)
	return dn
}

type field struct {
	schema.Field
	Name string
}

// HasDiscriminant reports whether the field is in a union.
func (f field) HasDiscriminant() bool {
	return f.DiscriminantValue() != schema.FieldnoDiscriminant
}

type enumval struct {
	schema.Enumerant
	Name   string
	Val    int
	Tag    string
	parent *node
}

func makeEnumval(enum *node, i int, e schema.Enumerant) enumval {
	eann := e.Annotations(nil)
	ann := parseAnnotations(eann)
	name := e.Name(nil)
	name = ann.Rename(name)
	t := ann.Tag(name)
	return enumval{e, name, i, t, enum}
}

func (e *enumval) FullName() string {
	return strings.ToUpper(e.parent.Name + "_" + e.Name)
}

type interfaceMethod struct {
	schema.Method
	Interface    *node
	ID           int
	Name         string
	OriginalName string
	Params       *node
	Results      *node
}

func methodSet(methods []interfaceMethod, n *node, nodes nodeMap) ([]interfaceMethod, error) {
	ms := n.Interface().Methods(nil)
	for i := 0; i < ms.Len(); i++ {
		m := ms.At(i)
		mname := m.Name(nil)
		mann := m.Annotations(nil)
		pn, err := nodes.mustFind(m.ParamStructType())
		if err != nil {
			return methods, fmt.Errorf("could not find param type for %s.%s", n.shortDisplayName(), mname)
		}
		rn, err := nodes.mustFind(m.ResultStructType())
		if err != nil {
			return methods, fmt.Errorf("could not find result type for %s.%s", n.shortDisplayName(), mname)
		}
		methods = append(methods, interfaceMethod{
			Method:       m,
			Interface:    n,
			ID:           i,
			OriginalName: mname,
			Name:         parseAnnotations(mann).Rename(mname),
			Params:       pn,
			Results:      rn,
		})
	}
	// TODO(light): sort added methods by code order

	supers := n.Interface().Superclasses(nil)
	for i := 0; i < supers.Len(); i++ {
		s := supers.At(i)
		sn, err := nodes.mustFind(s.Id())
		if err != nil {
			return methods, fmt.Errorf("could not find superclass %#x of %s", s.Id(), n)
		}
		methods, err = methodSet(methods, sn, nodes)
		if err != nil {
			return methods, err
		}
	}
	return methods, nil
}

// Tag types
const (
	defaultTag = iota
	noTag
	customTag
)

type annotations struct {
	Doc       string
	Package   string
	Import    string
	TagType   int
	CustomTag string
	Name      string
}

func parseAnnotations(list schema.Annotation_List) *annotations {
	ann := new(annotations)
	for i, n := 0, list.Len(); i < n; i++ {
		a := list.At(i)
		val := a.Value(nil)
		text := val.Text(nil)
		switch a.Id() {
		case capnp.Doc:
			ann.Doc = text
		case capnp.Package:
			ann.Package = text
		case capnp.Import:
			ann.Import = text
		case capnp.Tag:
			ann.TagType = customTag
			ann.CustomTag = text
		case capnp.Notag:
			ann.TagType = noTag
		case capnp.Name:
			ann.Name = text
		}
	}
	return ann
}

// Tag returns the string value that an enumerant value called name should have.
// An empty string indicates that this enumerant value has no tag.
func (ann *annotations) Tag(name string) string {
	switch ann.TagType {
	case noTag:
		return ""
	case customTag:
		return ann.CustomTag
	case defaultTag:
		fallthrough
	default:
		return name
	}
}

// Rename returns the overridden name from the annotations or the given name
// if no annotation was found.
func (ann *annotations) Rename(given string) string {
	if ann.Name == "" {
		return given
	}
	return ann.Name
}

type nodeMap map[uint64]*node

func buildNodeMap(req schema.CodeGeneratorRequest) (nodeMap, error) {
	var es capnp.ErrorSet
	rnodes := req.Nodes(&es)
	if es != nil {
		return nil, es
	}
	nodes := make(nodeMap, rnodes.Len())
	var allfiles []*node
	for i := 0; i < rnodes.Len(); i++ {
		ni := rnodes.At(i)
		n := &node{Node: ni}
		nodes[n.Id()] = n
		if n.Which() == schema.Node_Which_file {
			allfiles = append(allfiles, n)
		}
	}
	for _, f := range allfiles {
		fann := f.Annotations(&es)
		if es != nil {
			return nil, es
		}
		ann := parseAnnotations(fann)
		f.pkg = ann.Package
		f.imp = ann.Import
		nnodes := f.NestedNodes(nil)
		for i := 0; i < nnodes.Len(); i++ {
			nn := nnodes.At(i)
			if ni := nodes[nn.Id()]; ni != nil {
				nname := nn.Name(nil)
				if err := resolveName(nodes, ni, "", nname, f); err != nil {
					return nil, err
				}
			}
		}
	}
	return nodes, nil
}

// resolveName is called as part of building up a node map to populate the name field of n.
func resolveName(nodes nodeMap, n *node, base, name string, file *node) error {
	var es capnp.ErrorSet
	na := n.Annotations(&es)
	if es != nil {
		return fmt.Errorf("reading annotations for %s: %v", n, es)
	}
	name = parseAnnotations(na).Rename(name)
	if base == "" {
		n.Name = strings.Title(name)
	} else {
		n.Name = base + name
	}
	n.pkg = file.pkg
	n.imp = file.imp
	file.nodes = append(file.nodes, n)

	nnodes := n.NestedNodes(&es)
	if es != nil {
		return fmt.Errorf("listing nested nodes for %s: %v", n, es)
	}
	for i := 0; i < nnodes.Len(); i++ {
		nn := nnodes.At(i)
		ni := nodes[nn.Id()]
		if ni == nil {
			continue
		}
		nname := nn.Name(&es)
		if es != nil {
			return fmt.Errorf("reading name of nested node %d in %s: %v", i+1, n, es)
		}
		if err := resolveName(nodes, ni, n.Name, nname, file); err != nil {
			return err
		}
	}

	switch n.Which() {
	case schema.Node_Which_structNode:
		fields := n.StructNode().Fields(nil)
		for i := 0; i < fields.Len(); i++ {
			f := fields.At(i)
			if f.Which() != schema.Field_Which_group {
				continue
			}
			fa := f.Annotations(nil)
			fname := f.Name(nil)
			grp := nodes[f.Group().TypeId()]
			if grp == nil {
				return fmt.Errorf("could not find type information for group %s in %s", fname, n)
			}
			fname = parseAnnotations(fa).Rename(fname)
			if err := resolveName(nodes, grp, n.Name, fname, file); err != nil {
				return err
			}
		}
	case schema.Node_Which_interface:
		m := n.Interface().Methods(nil)
		methodResolve := func(id uint64, mname string, base string, name string) error {
			x := nodes[id]
			if x == nil {
				return fmt.Errorf("could not find type %#x for %s.%s", id, n, mname)
			}
			if x.ScopeId() != 0 {
				return nil
			}
			return resolveName(nodes, x, base, name, file)
		}
		for i := 0; i < m.Len(); i++ {
			mm := m.At(i)
			mname := mm.Name(nil)
			mann := mm.Annotations(nil)
			base := n.Name + parseAnnotations(mann).Rename(mname)
			if err := methodResolve(mm.ParamStructType(), mname, base, "Params"); err != nil {
				return err
			}
			if err := methodResolve(mm.ResultStructType(), mname, base, "Results"); err != nil {
				return err
			}
		}
	}
	return nil
}

func (nm nodeMap) mustFind(id uint64) (*node, error) {
	n := nm[id]
	if n == nil {
		return nil, fmt.Errorf("could not find node %#x in schema", id)
	}
	return n, nil
}
