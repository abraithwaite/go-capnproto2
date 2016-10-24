package schema

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Constants defined in schema.capnp.
const (
	FieldnoDiscriminant = uint16(65535)
)

type Node struct{ capnp.Struct }
type NodestructNode Node
type Nodeenum Node
type Nodeinterface Node
type Nodeconst Node
type Nodeannotation Node
type Node_Which uint16

const (
	Node_Which_file       Node_Which = 0
	Node_Which_structNode Node_Which = 1
	Node_Which_enum       Node_Which = 2
	Node_Which_interface  Node_Which = 3
	Node_Which_const      Node_Which = 4
	Node_Which_annotation Node_Which = 5
)

func (w Node_Which) String() string {
	const s = "filestructNodeenuminterfaceconstannotation"
	switch w {
	case Node_Which_file:
		return s[0:4]
	case Node_Which_structNode:
		return s[4:14]
	case Node_Which_enum:
		return s[14:18]
	case Node_Which_interface:
		return s[18:27]
	case Node_Which_const:
		return s[27:32]
	case Node_Which_annotation:
		return s[32:42]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xe682ab4cf923a417

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(12))
}
func (s Node) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) DisplayName(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s Node) HasDisplayName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) DisplayNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetDisplayName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Node) DisplayNamePrefixLength() uint32 {
	return s.Struct.Uint32(8)
}

func (s Node) SetDisplayNamePrefixLength(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Node) ScopeId() uint64 {
	return s.Struct.Uint64(16)
}

func (s Node) SetScopeId(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Node) Parameters(e *capnp.ErrorSet) NodeParameter_List {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		*e = append(*e, err)
		return NodeParameter_List{}
	}
	return NodeParameter_List{List: p.List()}
}

func (s Node) HasParameters() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Node) SetParameters(v NodeParameter_List) error {
	return s.Struct.SetPtr(5, v.List.ToPtr())
}

// NewParameters sets the parameters field to a newly
// allocated NodeParameter_List, preferring placement in s's segment.
func (s Node) NewParameters(n int32) (NodeParameter_List, error) {
	l, err := NewNodeParameter_List(s.Struct.Segment(), n)
	if err != nil {
		return NodeParameter_List{}, err
	}
	err = s.Struct.SetPtr(5, l.List.ToPtr())
	return l, err
}

func (s Node) IsGeneric() bool {
	return s.Struct.Bit(288)
}

func (s Node) SetIsGeneric(v bool) {
	s.Struct.SetBit(288, v)
}

func (s Node) NestedNodes(e *capnp.ErrorSet) NodeNestedNode_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return NodeNestedNode_List{}
	}
	return NodeNestedNode_List{List: p.List()}
}

func (s Node) HasNestedNodes() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetNestedNodes(v NodeNestedNode_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewNestedNodes sets the nestedNodes field to a newly
// allocated NodeNestedNode_List, preferring placement in s's segment.
func (s Node) NewNestedNodes(n int32) (NodeNestedNode_List, error) {
	l, err := NewNodeNestedNode_List(s.Struct.Segment(), n)
	if err != nil {
		return NodeNestedNode_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Node) Annotations(e *capnp.ErrorSet) Annotation_List {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		*e = append(*e, err)
		return Annotation_List{}
	}
	return Annotation_List{List: p.List()}
}

func (s Node) HasAnnotations() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Node) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Node) SetFile() {
	s.Struct.SetUint16(12, 0)

}

func (s Node) StructNode() NodestructNode { return NodestructNode(s) }

func (s Node) SetStructNode() {
	s.Struct.SetUint16(12, 1)
}

func (s NodestructNode) DataWordCount() uint16 {
	return s.Struct.Uint16(14)
}

func (s NodestructNode) SetDataWordCount(v uint16) {
	s.Struct.SetUint16(14, v)
}

func (s NodestructNode) PointerCount() uint16 {
	return s.Struct.Uint16(24)
}

func (s NodestructNode) SetPointerCount(v uint16) {
	s.Struct.SetUint16(24, v)
}

func (s NodestructNode) PreferredListEncoding() ElementSize {
	return ElementSize(s.Struct.Uint16(26))
}

func (s NodestructNode) SetPreferredListEncoding(v ElementSize) {
	s.Struct.SetUint16(26, uint16(v))
}

func (s NodestructNode) IsGroup() bool {
	return s.Struct.Bit(224)
}

func (s NodestructNode) SetIsGroup(v bool) {
	s.Struct.SetBit(224, v)
}

func (s NodestructNode) DiscriminantCount() uint16 {
	return s.Struct.Uint16(30)
}

func (s NodestructNode) SetDiscriminantCount(v uint16) {
	s.Struct.SetUint16(30, v)
}

func (s NodestructNode) DiscriminantOffset() uint32 {
	return s.Struct.Uint32(32)
}

func (s NodestructNode) SetDiscriminantOffset(v uint32) {
	s.Struct.SetUint32(32, v)
}

func (s NodestructNode) Fields(e *capnp.ErrorSet) Field_List {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Field_List{}
	}
	return Field_List{List: p.List()}
}

func (s NodestructNode) HasFields() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s NodestructNode) SetFields(v Field_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewFields sets the fields field to a newly
// allocated Field_List, preferring placement in s's segment.
func (s NodestructNode) NewFields(n int32) (Field_List, error) {
	l, err := NewField_List(s.Struct.Segment(), n)
	if err != nil {
		return Field_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Node) Enum() Nodeenum { return Nodeenum(s) }

func (s Node) SetEnum() {
	s.Struct.SetUint16(12, 2)
}

func (s Nodeenum) Enumerants(e *capnp.ErrorSet) Enumerant_List {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Enumerant_List{}
	}
	return Enumerant_List{List: p.List()}
}

func (s Nodeenum) HasEnumerants() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Nodeenum) SetEnumerants(v Enumerant_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewEnumerants sets the enumerants field to a newly
// allocated Enumerant_List, preferring placement in s's segment.
func (s Nodeenum) NewEnumerants(n int32) (Enumerant_List, error) {
	l, err := NewEnumerant_List(s.Struct.Segment(), n)
	if err != nil {
		return Enumerant_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Node) Interface() Nodeinterface { return Nodeinterface(s) }

func (s Node) SetInterface() {
	s.Struct.SetUint16(12, 3)
}

func (s Nodeinterface) Methods(e *capnp.ErrorSet) Method_List {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Method_List{}
	}
	return Method_List{List: p.List()}
}

func (s Nodeinterface) HasMethods() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Nodeinterface) SetMethods(v Method_List) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewMethods sets the methods field to a newly
// allocated Method_List, preferring placement in s's segment.
func (s Nodeinterface) NewMethods(n int32) (Method_List, error) {
	l, err := NewMethod_List(s.Struct.Segment(), n)
	if err != nil {
		return Method_List{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Nodeinterface) Superclasses(e *capnp.ErrorSet) Superclass_List {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		*e = append(*e, err)
		return Superclass_List{}
	}
	return Superclass_List{List: p.List()}
}

func (s Nodeinterface) HasSuperclasses() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Nodeinterface) SetSuperclasses(v Superclass_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewSuperclasses sets the superclasses field to a newly
// allocated Superclass_List, preferring placement in s's segment.
func (s Nodeinterface) NewSuperclasses(n int32) (Superclass_List, error) {
	l, err := NewSuperclass_List(s.Struct.Segment(), n)
	if err != nil {
		return Superclass_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s Node) Const() Nodeconst { return Nodeconst(s) }

func (s Node) SetConst() {
	s.Struct.SetUint16(12, 4)
}

func (s Nodeconst) Type(e *capnp.ErrorSet) Type {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Type{}
	}
	return Type{Struct: p.Struct()}
}

func (s Nodeconst) HasType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Nodeconst) SetType(v Type) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Nodeconst) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Nodeconst) Value(e *capnp.ErrorSet) Value {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		*e = append(*e, err)
		return Value{}
	}
	return Value{Struct: p.Struct()}
}

func (s Nodeconst) HasValue() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Nodeconst) SetValue(v Value) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Nodeconst) NewValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Annotation() Nodeannotation { return Nodeannotation(s) }

func (s Node) SetAnnotation() {
	s.Struct.SetUint16(12, 5)
}

func (s Nodeannotation) Type(e *capnp.ErrorSet) Type {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Type{}
	}
	return Type{Struct: p.Struct()}
}

func (s Nodeannotation) HasType() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Nodeannotation) SetType(v Type) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Nodeannotation) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Nodeannotation) TargetsFile() bool {
	return s.Struct.Bit(112)
}

func (s Nodeannotation) SetTargetsFile(v bool) {
	s.Struct.SetBit(112, v)
}

func (s Nodeannotation) TargetsConst() bool {
	return s.Struct.Bit(113)
}

func (s Nodeannotation) SetTargetsConst(v bool) {
	s.Struct.SetBit(113, v)
}

func (s Nodeannotation) TargetsEnum() bool {
	return s.Struct.Bit(114)
}

func (s Nodeannotation) SetTargetsEnum(v bool) {
	s.Struct.SetBit(114, v)
}

func (s Nodeannotation) TargetsEnumerant() bool {
	return s.Struct.Bit(115)
}

func (s Nodeannotation) SetTargetsEnumerant(v bool) {
	s.Struct.SetBit(115, v)
}

func (s Nodeannotation) TargetsStruct() bool {
	return s.Struct.Bit(116)
}

func (s Nodeannotation) SetTargetsStruct(v bool) {
	s.Struct.SetBit(116, v)
}

func (s Nodeannotation) TargetsField() bool {
	return s.Struct.Bit(117)
}

func (s Nodeannotation) SetTargetsField(v bool) {
	s.Struct.SetBit(117, v)
}

func (s Nodeannotation) TargetsUnion() bool {
	return s.Struct.Bit(118)
}

func (s Nodeannotation) SetTargetsUnion(v bool) {
	s.Struct.SetBit(118, v)
}

func (s Nodeannotation) TargetsGroup() bool {
	return s.Struct.Bit(119)
}

func (s Nodeannotation) SetTargetsGroup(v bool) {
	s.Struct.SetBit(119, v)
}

func (s Nodeannotation) TargetsInterface() bool {
	return s.Struct.Bit(120)
}

func (s Nodeannotation) SetTargetsInterface(v bool) {
	s.Struct.SetBit(120, v)
}

func (s Nodeannotation) TargetsMethod() bool {
	return s.Struct.Bit(121)
}

func (s Nodeannotation) SetTargetsMethod(v bool) {
	s.Struct.SetBit(121, v)
}

func (s Nodeannotation) TargetsParam() bool {
	return s.Struct.Bit(122)
}

func (s Nodeannotation) SetTargetsParam(v bool) {
	s.Struct.SetBit(122, v)
}

func (s Nodeannotation) TargetsAnnotation() bool {
	return s.Struct.Bit(123)
}

func (s Nodeannotation) SetTargetsAnnotation(v bool) {
	s.Struct.SetBit(123, v)
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) StructNode() NodestructNode_Promise { return NodestructNode_Promise{p.Pipeline} }

// NodestructNode_Promise is a wrapper for a NodestructNode promised by a client call.
type NodestructNode_Promise struct{ *capnp.Pipeline }

func (p NodestructNode_Promise) Struct() (NodestructNode, error) {
	s, err := p.Pipeline.Struct()
	return NodestructNode{s}, err
}

func (p Node_Promise) Enum() Nodeenum_Promise { return Nodeenum_Promise{p.Pipeline} }

// Nodeenum_Promise is a wrapper for a Nodeenum promised by a client call.
type Nodeenum_Promise struct{ *capnp.Pipeline }

func (p Nodeenum_Promise) Struct() (Nodeenum, error) {
	s, err := p.Pipeline.Struct()
	return Nodeenum{s}, err
}

func (p Node_Promise) Interface() Nodeinterface_Promise { return Nodeinterface_Promise{p.Pipeline} }

// Nodeinterface_Promise is a wrapper for a Nodeinterface promised by a client call.
type Nodeinterface_Promise struct{ *capnp.Pipeline }

func (p Nodeinterface_Promise) Struct() (Nodeinterface, error) {
	s, err := p.Pipeline.Struct()
	return Nodeinterface{s}, err
}

func (p Node_Promise) Const() Nodeconst_Promise { return Nodeconst_Promise{p.Pipeline} }

// Nodeconst_Promise is a wrapper for a Nodeconst promised by a client call.
type Nodeconst_Promise struct{ *capnp.Pipeline }

func (p Nodeconst_Promise) Struct() (Nodeconst, error) {
	s, err := p.Pipeline.Struct()
	return Nodeconst{s}, err
}

func (p Nodeconst_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Nodeconst_Promise) Value() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Node_Promise) Annotation() Nodeannotation_Promise { return Nodeannotation_Promise{p.Pipeline} }

// Nodeannotation_Promise is a wrapper for a Nodeannotation promised by a client call.
type Nodeannotation_Promise struct{ *capnp.Pipeline }

func (p Nodeannotation_Promise) Struct() (Nodeannotation, error) {
	s, err := p.Pipeline.Struct()
	return Nodeannotation{s}, err
}

func (p Nodeannotation_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type NodeParameter struct{ capnp.Struct }

// NodeParameter_TypeID is the unique identifier for the type NodeParameter.
const NodeParameter_TypeID = 0xb9521bccf10fa3b1

func NewNodeParameter(s *capnp.Segment) (NodeParameter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return NodeParameter{st}, err
}

func NewRootNodeParameter(s *capnp.Segment) (NodeParameter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return NodeParameter{st}, err
}

func ReadRootNodeParameter(msg *capnp.Message) (NodeParameter, error) {
	root, err := msg.RootPtr()
	return NodeParameter{root.Struct()}, err
}

func (s NodeParameter) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s NodeParameter) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NodeParameter) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s NodeParameter) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// NodeParameter_List is a list of NodeParameter.
type NodeParameter_List struct{ capnp.List }

// NewNodeParameter creates a new list of NodeParameter.
func NewNodeParameter_List(s *capnp.Segment, sz int32) (NodeParameter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return NodeParameter_List{l}, err
}

func (s NodeParameter_List) At(i int) NodeParameter { return NodeParameter{s.List.Struct(i)} }

func (s NodeParameter_List) Set(i int, v NodeParameter) error { return s.List.SetStruct(i, v.Struct) }

// NodeParameter_Promise is a wrapper for a NodeParameter promised by a client call.
type NodeParameter_Promise struct{ *capnp.Pipeline }

func (p NodeParameter_Promise) Struct() (NodeParameter, error) {
	s, err := p.Pipeline.Struct()
	return NodeParameter{s}, err
}

type NodeNestedNode struct{ capnp.Struct }

// NodeNestedNode_TypeID is the unique identifier for the type NodeNestedNode.
const NodeNestedNode_TypeID = 0xdebf55bbfa0fc242

func NewNodeNestedNode(s *capnp.Segment) (NodeNestedNode, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return NodeNestedNode{st}, err
}

func NewRootNodeNestedNode(s *capnp.Segment) (NodeNestedNode, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return NodeNestedNode{st}, err
}

func ReadRootNodeNestedNode(msg *capnp.Message) (NodeNestedNode, error) {
	root, err := msg.RootPtr()
	return NodeNestedNode{root.Struct()}, err
}

func (s NodeNestedNode) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s NodeNestedNode) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s NodeNestedNode) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s NodeNestedNode) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s NodeNestedNode) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s NodeNestedNode) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

// NodeNestedNode_List is a list of NodeNestedNode.
type NodeNestedNode_List struct{ capnp.List }

// NewNodeNestedNode creates a new list of NodeNestedNode.
func NewNodeNestedNode_List(s *capnp.Segment, sz int32) (NodeNestedNode_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return NodeNestedNode_List{l}, err
}

func (s NodeNestedNode_List) At(i int) NodeNestedNode { return NodeNestedNode{s.List.Struct(i)} }

func (s NodeNestedNode_List) Set(i int, v NodeNestedNode) error { return s.List.SetStruct(i, v.Struct) }

// NodeNestedNode_Promise is a wrapper for a NodeNestedNode promised by a client call.
type NodeNestedNode_Promise struct{ *capnp.Pipeline }

func (p NodeNestedNode_Promise) Struct() (NodeNestedNode, error) {
	s, err := p.Pipeline.Struct()
	return NodeNestedNode{s}, err
}

type Field struct{ capnp.Struct }
type Fieldslot Field
type Fieldgroup Field
type Fieldordinal Field
type Field_Which uint16

const (
	Field_Which_slot  Field_Which = 0
	Field_Which_group Field_Which = 1
)

func (w Field_Which) String() string {
	const s = "slotgroup"
	switch w {
	case Field_Which_slot:
		return s[0:4]
	case Field_Which_group:
		return s[4:9]

	}
	return "Field_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Fieldordinal_Which uint16

const (
	Fieldordinal_Which_implicit Fieldordinal_Which = 0
	Fieldordinal_Which_explicit Fieldordinal_Which = 1
)

func (w Fieldordinal_Which) String() string {
	const s = "implicitexplicit"
	switch w {
	case Fieldordinal_Which_implicit:
		return s[0:8]
	case Fieldordinal_Which_explicit:
		return s[8:16]

	}
	return "Fieldordinal_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Field_TypeID is the unique identifier for the type Field.
const Field_TypeID = 0x9aad50a41f4af45f

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field{st}, err
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field{st}, err
}

func ReadRootField(msg *capnp.Message) (Field, error) {
	root, err := msg.RootPtr()
	return Field{root.Struct()}, err
}

func (s Field) Which() Field_Which {
	return Field_Which(s.Struct.Uint16(8))
}
func (s Field) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s Field) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Field) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Field) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Field) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Field) Annotations(e *capnp.ErrorSet) Annotation_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return Annotation_List{}
	}
	return Annotation_List{List: p.List()}
}

func (s Field) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Field) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Field) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Field) DiscriminantValue() uint16 {
	return s.Struct.Uint16(2) ^ 65535
}

func (s Field) SetDiscriminantValue(v uint16) {
	s.Struct.SetUint16(2, v^65535)
}

func (s Field) Slot() Fieldslot { return Fieldslot(s) }

func (s Field) SetSlot() {
	s.Struct.SetUint16(8, 0)
}

func (s Fieldslot) Offset() uint32 {
	return s.Struct.Uint32(4)
}

func (s Fieldslot) SetOffset(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Fieldslot) Type(e *capnp.ErrorSet) Type {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		*e = append(*e, err)
		return Type{}
	}
	return Type{Struct: p.Struct()}
}

func (s Fieldslot) HasType() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Fieldslot) SetType(v Type) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Fieldslot) NewType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Fieldslot) DefaultValue(e *capnp.ErrorSet) Value {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Value{}
	}
	return Value{Struct: p.Struct()}
}

func (s Fieldslot) HasDefaultValue() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Fieldslot) SetDefaultValue(v Value) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDefaultValue sets the defaultValue field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Fieldslot) NewDefaultValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Fieldslot) HadExplicitDefault() bool {
	return s.Struct.Bit(128)
}

func (s Fieldslot) SetHadExplicitDefault(v bool) {
	s.Struct.SetBit(128, v)
}

func (s Field) Group() Fieldgroup { return Fieldgroup(s) }

func (s Field) SetGroup() {
	s.Struct.SetUint16(8, 1)
}

func (s Fieldgroup) TypeId() uint64 {
	return s.Struct.Uint64(16)
}

func (s Fieldgroup) SetTypeId(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Field) Ordinal() Fieldordinal { return Fieldordinal(s) }

func (s Fieldordinal) Which() Fieldordinal_Which {
	return Fieldordinal_Which(s.Struct.Uint16(10))
}
func (s Fieldordinal) SetImplicit() {
	s.Struct.SetUint16(10, 0)

}

func (s Fieldordinal) Explicit() uint16 {
	return s.Struct.Uint16(12)
}

func (s Fieldordinal) SetExplicit(v uint16) {
	s.Struct.SetUint16(10, 1)
	s.Struct.SetUint16(12, v)
}

// Field_List is a list of Field.
type Field_List struct{ capnp.List }

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4}, sz)
	return Field_List{l}, err
}

func (s Field_List) At(i int) Field { return Field{s.List.Struct(i)} }

func (s Field_List) Set(i int, v Field) error { return s.List.SetStruct(i, v.Struct) }

// Field_Promise is a wrapper for a Field promised by a client call.
type Field_Promise struct{ *capnp.Pipeline }

func (p Field_Promise) Struct() (Field, error) {
	s, err := p.Pipeline.Struct()
	return Field{s}, err
}

func (p Field_Promise) Slot() Fieldslot_Promise { return Fieldslot_Promise{p.Pipeline} }

// Fieldslot_Promise is a wrapper for a Fieldslot promised by a client call.
type Fieldslot_Promise struct{ *capnp.Pipeline }

func (p Fieldslot_Promise) Struct() (Fieldslot, error) {
	s, err := p.Pipeline.Struct()
	return Fieldslot{s}, err
}

func (p Fieldslot_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Fieldslot_Promise) DefaultValue() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Field_Promise) Group() Fieldgroup_Promise { return Fieldgroup_Promise{p.Pipeline} }

// Fieldgroup_Promise is a wrapper for a Fieldgroup promised by a client call.
type Fieldgroup_Promise struct{ *capnp.Pipeline }

func (p Fieldgroup_Promise) Struct() (Fieldgroup, error) {
	s, err := p.Pipeline.Struct()
	return Fieldgroup{s}, err
}

func (p Field_Promise) Ordinal() Fieldordinal_Promise { return Fieldordinal_Promise{p.Pipeline} }

// Fieldordinal_Promise is a wrapper for a Fieldordinal promised by a client call.
type Fieldordinal_Promise struct{ *capnp.Pipeline }

func (p Fieldordinal_Promise) Struct() (Fieldordinal, error) {
	s, err := p.Pipeline.Struct()
	return Fieldordinal{s}, err
}

type Enumerant struct{ capnp.Struct }

// Enumerant_TypeID is the unique identifier for the type Enumerant.
const Enumerant_TypeID = 0x978a7cebdc549a4d

func NewEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant{st}, err
}

func NewRootEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant{st}, err
}

func ReadRootEnumerant(msg *capnp.Message) (Enumerant, error) {
	root, err := msg.RootPtr()
	return Enumerant{root.Struct()}, err
}

func (s Enumerant) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s Enumerant) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Enumerant) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Enumerant) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Enumerant) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Enumerant) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Enumerant) Annotations(e *capnp.ErrorSet) Annotation_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return Annotation_List{}
	}
	return Annotation_List{List: p.List()}
}

func (s Enumerant) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Enumerant) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Enumerant) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Enumerant_List is a list of Enumerant.
type Enumerant_List struct{ capnp.List }

// NewEnumerant creates a new list of Enumerant.
func NewEnumerant_List(s *capnp.Segment, sz int32) (Enumerant_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Enumerant_List{l}, err
}

func (s Enumerant_List) At(i int) Enumerant { return Enumerant{s.List.Struct(i)} }

func (s Enumerant_List) Set(i int, v Enumerant) error { return s.List.SetStruct(i, v.Struct) }

// Enumerant_Promise is a wrapper for a Enumerant promised by a client call.
type Enumerant_Promise struct{ *capnp.Pipeline }

func (p Enumerant_Promise) Struct() (Enumerant, error) {
	s, err := p.Pipeline.Struct()
	return Enumerant{s}, err
}

type Superclass struct{ capnp.Struct }

// Superclass_TypeID is the unique identifier for the type Superclass.
const Superclass_TypeID = 0xa9962a9ed0a4d7f8

func NewSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass{st}, err
}

func NewRootSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass{st}, err
}

func ReadRootSuperclass(msg *capnp.Message) (Superclass, error) {
	root, err := msg.RootPtr()
	return Superclass{root.Struct()}, err
}

func (s Superclass) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Superclass) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Superclass) Brand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Superclass) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Superclass) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Superclass) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Superclass_List is a list of Superclass.
type Superclass_List struct{ capnp.List }

// NewSuperclass creates a new list of Superclass.
func NewSuperclass_List(s *capnp.Segment, sz int32) (Superclass_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Superclass_List{l}, err
}

func (s Superclass_List) At(i int) Superclass { return Superclass{s.List.Struct(i)} }

func (s Superclass_List) Set(i int, v Superclass) error { return s.List.SetStruct(i, v.Struct) }

// Superclass_Promise is a wrapper for a Superclass promised by a client call.
type Superclass_Promise struct{ *capnp.Pipeline }

func (p Superclass_Promise) Struct() (Superclass, error) {
	s, err := p.Pipeline.Struct()
	return Superclass{s}, err
}

func (p Superclass_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Method struct{ capnp.Struct }

// Method_TypeID is the unique identifier for the type Method.
const Method_TypeID = 0x9500cce23b334d80

func NewMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method{st}, err
}

func NewRootMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method{st}, err
}

func ReadRootMethod(msg *capnp.Message) (Method, error) {
	root, err := msg.RootPtr()
	return Method{root.Struct()}, err
}

func (s Method) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s Method) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Method) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Method) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Method) CodeOrder() uint16 {
	return s.Struct.Uint16(0)
}

func (s Method) SetCodeOrder(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Method) ImplicitParameters(e *capnp.ErrorSet) NodeParameter_List {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		*e = append(*e, err)
		return NodeParameter_List{}
	}
	return NodeParameter_List{List: p.List()}
}

func (s Method) HasImplicitParameters() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Method) SetImplicitParameters(v NodeParameter_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewImplicitParameters sets the implicitParameters field to a newly
// allocated NodeParameter_List, preferring placement in s's segment.
func (s Method) NewImplicitParameters(n int32) (NodeParameter_List, error) {
	l, err := NewNodeParameter_List(s.Struct.Segment(), n)
	if err != nil {
		return NodeParameter_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

func (s Method) ParamStructType() uint64 {
	return s.Struct.Uint64(8)
}

func (s Method) SetParamStructType(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Method) ParamBrand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Method) HasParamBrand() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Method) SetParamBrand(v Brand) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewParamBrand sets the paramBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewParamBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Method) ResultStructType() uint64 {
	return s.Struct.Uint64(16)
}

func (s Method) SetResultStructType(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Method) ResultBrand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Method) HasResultBrand() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Method) SetResultBrand(v Brand) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewResultBrand sets the resultBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewResultBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Method) Annotations(e *capnp.ErrorSet) Annotation_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return Annotation_List{}
	}
	return Annotation_List{List: p.List()}
}

func (s Method) HasAnnotations() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Method) SetAnnotations(v Annotation_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Method) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(s.Struct.Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Method_List is a list of Method.
type Method_List struct{ capnp.List }

// NewMethod creates a new list of Method.
func NewMethod_List(s *capnp.Segment, sz int32) (Method_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5}, sz)
	return Method_List{l}, err
}

func (s Method_List) At(i int) Method { return Method{s.List.Struct(i)} }

func (s Method_List) Set(i int, v Method) error { return s.List.SetStruct(i, v.Struct) }

// Method_Promise is a wrapper for a Method promised by a client call.
type Method_Promise struct{ *capnp.Pipeline }

func (p Method_Promise) Struct() (Method, error) {
	s, err := p.Pipeline.Struct()
	return Method{s}, err
}

func (p Method_Promise) ParamBrand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Method_Promise) ResultBrand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type Type struct{ capnp.Struct }
type Typelist Type
type Typeenum Type
type TypestructType Type
type Typeinterface Type
type TypeanyPointer Type
type TypeanyPointerunconstrained Type
type TypeanyPointerparameter Type
type TypeanyPointerimplicitMethodParameter Type
type Type_Which uint16

const (
	Type_Which_void       Type_Which = 0
	Type_Which_bool       Type_Which = 1
	Type_Which_int8       Type_Which = 2
	Type_Which_int16      Type_Which = 3
	Type_Which_int32      Type_Which = 4
	Type_Which_int64      Type_Which = 5
	Type_Which_uint8      Type_Which = 6
	Type_Which_uint16     Type_Which = 7
	Type_Which_uint32     Type_Which = 8
	Type_Which_uint64     Type_Which = 9
	Type_Which_float32    Type_Which = 10
	Type_Which_float64    Type_Which = 11
	Type_Which_text       Type_Which = 12
	Type_Which_data       Type_Which = 13
	Type_Which_list       Type_Which = 14
	Type_Which_enum       Type_Which = 15
	Type_Which_structType Type_Which = 16
	Type_Which_interface  Type_Which = 17
	Type_Which_anyPointer Type_Which = 18
)

func (w Type_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructTypeinterfaceanyPointer"
	switch w {
	case Type_Which_void:
		return s[0:4]
	case Type_Which_bool:
		return s[4:8]
	case Type_Which_int8:
		return s[8:12]
	case Type_Which_int16:
		return s[12:17]
	case Type_Which_int32:
		return s[17:22]
	case Type_Which_int64:
		return s[22:27]
	case Type_Which_uint8:
		return s[27:32]
	case Type_Which_uint16:
		return s[32:38]
	case Type_Which_uint32:
		return s[38:44]
	case Type_Which_uint64:
		return s[44:50]
	case Type_Which_float32:
		return s[50:57]
	case Type_Which_float64:
		return s[57:64]
	case Type_Which_text:
		return s[64:68]
	case Type_Which_data:
		return s[68:72]
	case Type_Which_list:
		return s[72:76]
	case Type_Which_enum:
		return s[76:80]
	case Type_Which_structType:
		return s[80:90]
	case Type_Which_interface:
		return s[90:99]
	case Type_Which_anyPointer:
		return s[99:109]

	}
	return "Type_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type TypeanyPointer_Which uint16

const (
	TypeanyPointer_Which_unconstrained           TypeanyPointer_Which = 0
	TypeanyPointer_Which_parameter               TypeanyPointer_Which = 1
	TypeanyPointer_Which_implicitMethodParameter TypeanyPointer_Which = 2
)

func (w TypeanyPointer_Which) String() string {
	const s = "unconstrainedparameterimplicitMethodParameter"
	switch w {
	case TypeanyPointer_Which_unconstrained:
		return s[0:13]
	case TypeanyPointer_Which_parameter:
		return s[13:22]
	case TypeanyPointer_Which_implicitMethodParameter:
		return s[22:45]

	}
	return "TypeanyPointer_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type TypeanyPointerunconstrained_Which uint16

const (
	TypeanyPointerunconstrained_Which_anyKind    TypeanyPointerunconstrained_Which = 0
	TypeanyPointerunconstrained_Which_struct     TypeanyPointerunconstrained_Which = 1
	TypeanyPointerunconstrained_Which_list       TypeanyPointerunconstrained_Which = 2
	TypeanyPointerunconstrained_Which_capability TypeanyPointerunconstrained_Which = 3
)

func (w TypeanyPointerunconstrained_Which) String() string {
	const s = "anyKindstructlistcapability"
	switch w {
	case TypeanyPointerunconstrained_Which_anyKind:
		return s[0:7]
	case TypeanyPointerunconstrained_Which_struct:
		return s[7:13]
	case TypeanyPointerunconstrained_Which_list:
		return s[13:17]
	case TypeanyPointerunconstrained_Which_capability:
		return s[17:27]

	}
	return "TypeanyPointerunconstrained_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Type_TypeID is the unique identifier for the type Type.
const Type_TypeID = 0xd07378ede1f9cc60

func NewType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Type{st}, err
}

func NewRootType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Type{st}, err
}

func ReadRootType(msg *capnp.Message) (Type, error) {
	root, err := msg.RootPtr()
	return Type{root.Struct()}, err
}

func (s Type) Which() Type_Which {
	return Type_Which(s.Struct.Uint16(0))
}
func (s Type) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Type) SetBool() {
	s.Struct.SetUint16(0, 1)

}

func (s Type) SetInt8() {
	s.Struct.SetUint16(0, 2)

}

func (s Type) SetInt16() {
	s.Struct.SetUint16(0, 3)

}

func (s Type) SetInt32() {
	s.Struct.SetUint16(0, 4)

}

func (s Type) SetInt64() {
	s.Struct.SetUint16(0, 5)

}

func (s Type) SetUint8() {
	s.Struct.SetUint16(0, 6)

}

func (s Type) SetUint16() {
	s.Struct.SetUint16(0, 7)

}

func (s Type) SetUint32() {
	s.Struct.SetUint16(0, 8)

}

func (s Type) SetUint64() {
	s.Struct.SetUint16(0, 9)

}

func (s Type) SetFloat32() {
	s.Struct.SetUint16(0, 10)

}

func (s Type) SetFloat64() {
	s.Struct.SetUint16(0, 11)

}

func (s Type) SetText() {
	s.Struct.SetUint16(0, 12)

}

func (s Type) SetData() {
	s.Struct.SetUint16(0, 13)

}

func (s Type) List() Typelist { return Typelist(s) }

func (s Type) SetList() {
	s.Struct.SetUint16(0, 14)
}

func (s Typelist) ElementType(e *capnp.ErrorSet) Type {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Type{}
	}
	return Type{Struct: p.Struct()}
}

func (s Typelist) HasElementType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Typelist) SetElementType(v Type) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewElementType sets the elementType field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Typelist) NewElementType() (Type, error) {
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) Enum() Typeenum { return Typeenum(s) }

func (s Type) SetEnum() {
	s.Struct.SetUint16(0, 15)
}

func (s Typeenum) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Typeenum) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Typeenum) Brand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Typeenum) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Typeenum) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Typeenum) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) StructType() TypestructType { return TypestructType(s) }

func (s Type) SetStructType() {
	s.Struct.SetUint16(0, 16)
}

func (s TypestructType) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s TypestructType) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TypestructType) Brand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s TypestructType) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TypestructType) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s TypestructType) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) Interface() Typeinterface { return Typeinterface(s) }

func (s Type) SetInterface() {
	s.Struct.SetUint16(0, 17)
}

func (s Typeinterface) TypeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Typeinterface) SetTypeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Typeinterface) Brand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Typeinterface) HasBrand() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Typeinterface) SetBrand(v Brand) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Typeinterface) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Type) AnyPointer() TypeanyPointer { return TypeanyPointer(s) }

func (s Type) SetAnyPointer() {
	s.Struct.SetUint16(0, 18)
}

func (s TypeanyPointer) Which() TypeanyPointer_Which {
	return TypeanyPointer_Which(s.Struct.Uint16(4))
}
func (s TypeanyPointer) Unconstrained() TypeanyPointerunconstrained {
	return TypeanyPointerunconstrained(s)
}

func (s TypeanyPointer) SetUnconstrained() {
	s.Struct.SetUint16(4, 0)
}

func (s TypeanyPointerunconstrained) Which() TypeanyPointerunconstrained_Which {
	return TypeanyPointerunconstrained_Which(s.Struct.Uint16(2))
}
func (s TypeanyPointerunconstrained) SetAnyKind() {
	s.Struct.SetUint16(2, 0)

}

func (s TypeanyPointerunconstrained) SetStruct() {
	s.Struct.SetUint16(2, 1)

}

func (s TypeanyPointerunconstrained) SetList() {
	s.Struct.SetUint16(2, 2)

}

func (s TypeanyPointerunconstrained) SetCapability() {
	s.Struct.SetUint16(2, 3)

}

func (s TypeanyPointer) Parameter() TypeanyPointerparameter { return TypeanyPointerparameter(s) }

func (s TypeanyPointer) SetParameter() {
	s.Struct.SetUint16(4, 1)
}

func (s TypeanyPointerparameter) ScopeId() uint64 {
	return s.Struct.Uint64(8)
}

func (s TypeanyPointerparameter) SetScopeId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TypeanyPointerparameter) ParameterIndex() uint16 {
	return s.Struct.Uint16(2)
}

func (s TypeanyPointerparameter) SetParameterIndex(v uint16) {
	s.Struct.SetUint16(2, v)
}

func (s TypeanyPointer) ImplicitMethodParameter() TypeanyPointerimplicitMethodParameter {
	return TypeanyPointerimplicitMethodParameter(s)
}

func (s TypeanyPointer) SetImplicitMethodParameter() {
	s.Struct.SetUint16(4, 2)
}

func (s TypeanyPointerimplicitMethodParameter) ParameterIndex() uint16 {
	return s.Struct.Uint16(2)
}

func (s TypeanyPointerimplicitMethodParameter) SetParameterIndex(v uint16) {
	s.Struct.SetUint16(2, v)
}

// Type_List is a list of Type.
type Type_List struct{ capnp.List }

// NewType creates a new list of Type.
func NewType_List(s *capnp.Segment, sz int32) (Type_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Type_List{l}, err
}

func (s Type_List) At(i int) Type { return Type{s.List.Struct(i)} }

func (s Type_List) Set(i int, v Type) error { return s.List.SetStruct(i, v.Struct) }

// Type_Promise is a wrapper for a Type promised by a client call.
type Type_Promise struct{ *capnp.Pipeline }

func (p Type_Promise) Struct() (Type, error) {
	s, err := p.Pipeline.Struct()
	return Type{s}, err
}

func (p Type_Promise) List() Typelist_Promise { return Typelist_Promise{p.Pipeline} }

// Typelist_Promise is a wrapper for a Typelist promised by a client call.
type Typelist_Promise struct{ *capnp.Pipeline }

func (p Typelist_Promise) Struct() (Typelist, error) {
	s, err := p.Pipeline.Struct()
	return Typelist{s}, err
}

func (p Typelist_Promise) ElementType() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) Enum() Typeenum_Promise { return Typeenum_Promise{p.Pipeline} }

// Typeenum_Promise is a wrapper for a Typeenum promised by a client call.
type Typeenum_Promise struct{ *capnp.Pipeline }

func (p Typeenum_Promise) Struct() (Typeenum, error) {
	s, err := p.Pipeline.Struct()
	return Typeenum{s}, err
}

func (p Typeenum_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) StructType() TypestructType_Promise { return TypestructType_Promise{p.Pipeline} }

// TypestructType_Promise is a wrapper for a TypestructType promised by a client call.
type TypestructType_Promise struct{ *capnp.Pipeline }

func (p TypestructType_Promise) Struct() (TypestructType, error) {
	s, err := p.Pipeline.Struct()
	return TypestructType{s}, err
}

func (p TypestructType_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) Interface() Typeinterface_Promise { return Typeinterface_Promise{p.Pipeline} }

// Typeinterface_Promise is a wrapper for a Typeinterface promised by a client call.
type Typeinterface_Promise struct{ *capnp.Pipeline }

func (p Typeinterface_Promise) Struct() (Typeinterface, error) {
	s, err := p.Pipeline.Struct()
	return Typeinterface{s}, err
}

func (p Typeinterface_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Type_Promise) AnyPointer() TypeanyPointer_Promise { return TypeanyPointer_Promise{p.Pipeline} }

// TypeanyPointer_Promise is a wrapper for a TypeanyPointer promised by a client call.
type TypeanyPointer_Promise struct{ *capnp.Pipeline }

func (p TypeanyPointer_Promise) Struct() (TypeanyPointer, error) {
	s, err := p.Pipeline.Struct()
	return TypeanyPointer{s}, err
}

func (p TypeanyPointer_Promise) Unconstrained() TypeanyPointerunconstrained_Promise {
	return TypeanyPointerunconstrained_Promise{p.Pipeline}
}

// TypeanyPointerunconstrained_Promise is a wrapper for a TypeanyPointerunconstrained promised by a client call.
type TypeanyPointerunconstrained_Promise struct{ *capnp.Pipeline }

func (p TypeanyPointerunconstrained_Promise) Struct() (TypeanyPointerunconstrained, error) {
	s, err := p.Pipeline.Struct()
	return TypeanyPointerunconstrained{s}, err
}

func (p TypeanyPointer_Promise) Parameter() TypeanyPointerparameter_Promise {
	return TypeanyPointerparameter_Promise{p.Pipeline}
}

// TypeanyPointerparameter_Promise is a wrapper for a TypeanyPointerparameter promised by a client call.
type TypeanyPointerparameter_Promise struct{ *capnp.Pipeline }

func (p TypeanyPointerparameter_Promise) Struct() (TypeanyPointerparameter, error) {
	s, err := p.Pipeline.Struct()
	return TypeanyPointerparameter{s}, err
}

func (p TypeanyPointer_Promise) ImplicitMethodParameter() TypeanyPointerimplicitMethodParameter_Promise {
	return TypeanyPointerimplicitMethodParameter_Promise{p.Pipeline}
}

// TypeanyPointerimplicitMethodParameter_Promise is a wrapper for a TypeanyPointerimplicitMethodParameter promised by a client call.
type TypeanyPointerimplicitMethodParameter_Promise struct{ *capnp.Pipeline }

func (p TypeanyPointerimplicitMethodParameter_Promise) Struct() (TypeanyPointerimplicitMethodParameter, error) {
	s, err := p.Pipeline.Struct()
	return TypeanyPointerimplicitMethodParameter{s}, err
}

type Brand struct{ capnp.Struct }

// Brand_TypeID is the unique identifier for the type Brand.
const Brand_TypeID = 0x903455f06065422b

func NewBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand{st}, err
}

func NewRootBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand{st}, err
}

func ReadRootBrand(msg *capnp.Message) (Brand, error) {
	root, err := msg.RootPtr()
	return Brand{root.Struct()}, err
}

func (s Brand) Scopes(e *capnp.ErrorSet) BrandScope_List {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return BrandScope_List{}
	}
	return BrandScope_List{List: p.List()}
}

func (s Brand) HasScopes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Brand) SetScopes(v BrandScope_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewScopes sets the scopes field to a newly
// allocated BrandScope_List, preferring placement in s's segment.
func (s Brand) NewScopes(n int32) (BrandScope_List, error) {
	l, err := NewBrandScope_List(s.Struct.Segment(), n)
	if err != nil {
		return BrandScope_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Brand_List is a list of Brand.
type Brand_List struct{ capnp.List }

// NewBrand creates a new list of Brand.
func NewBrand_List(s *capnp.Segment, sz int32) (Brand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Brand_List{l}, err
}

func (s Brand_List) At(i int) Brand { return Brand{s.List.Struct(i)} }

func (s Brand_List) Set(i int, v Brand) error { return s.List.SetStruct(i, v.Struct) }

// Brand_Promise is a wrapper for a Brand promised by a client call.
type Brand_Promise struct{ *capnp.Pipeline }

func (p Brand_Promise) Struct() (Brand, error) {
	s, err := p.Pipeline.Struct()
	return Brand{s}, err
}

type BrandScope struct{ capnp.Struct }
type BrandScope_Which uint16

const (
	BrandScope_Which_bind    BrandScope_Which = 0
	BrandScope_Which_inherit BrandScope_Which = 1
)

func (w BrandScope_Which) String() string {
	const s = "bindinherit"
	switch w {
	case BrandScope_Which_bind:
		return s[0:4]
	case BrandScope_Which_inherit:
		return s[4:11]

	}
	return "BrandScope_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// BrandScope_TypeID is the unique identifier for the type BrandScope.
const BrandScope_TypeID = 0xabd73485a9636bc9

func NewBrandScope(s *capnp.Segment) (BrandScope, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return BrandScope{st}, err
}

func NewRootBrandScope(s *capnp.Segment) (BrandScope, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return BrandScope{st}, err
}

func ReadRootBrandScope(msg *capnp.Message) (BrandScope, error) {
	root, err := msg.RootPtr()
	return BrandScope{root.Struct()}, err
}

func (s BrandScope) Which() BrandScope_Which {
	return BrandScope_Which(s.Struct.Uint16(8))
}
func (s BrandScope) ScopeId() uint64 {
	return s.Struct.Uint64(0)
}

func (s BrandScope) SetScopeId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s BrandScope) Bind(e *capnp.ErrorSet) BrandBinding_List {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return BrandBinding_List{}
	}
	return BrandBinding_List{List: p.List()}
}

func (s BrandScope) HasBind() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BrandScope) SetBind(v BrandBinding_List) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBind sets the bind field to a newly
// allocated BrandBinding_List, preferring placement in s's segment.
func (s BrandScope) NewBind(n int32) (BrandBinding_List, error) {
	s.Struct.SetUint16(8, 0)
	l, err := NewBrandBinding_List(s.Struct.Segment(), n)
	if err != nil {
		return BrandBinding_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s BrandScope) SetInherit() {
	s.Struct.SetUint16(8, 1)

}

// BrandScope_List is a list of BrandScope.
type BrandScope_List struct{ capnp.List }

// NewBrandScope creates a new list of BrandScope.
func NewBrandScope_List(s *capnp.Segment, sz int32) (BrandScope_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return BrandScope_List{l}, err
}

func (s BrandScope_List) At(i int) BrandScope { return BrandScope{s.List.Struct(i)} }

func (s BrandScope_List) Set(i int, v BrandScope) error { return s.List.SetStruct(i, v.Struct) }

// BrandScope_Promise is a wrapper for a BrandScope promised by a client call.
type BrandScope_Promise struct{ *capnp.Pipeline }

func (p BrandScope_Promise) Struct() (BrandScope, error) {
	s, err := p.Pipeline.Struct()
	return BrandScope{s}, err
}

type BrandBinding struct{ capnp.Struct }
type BrandBinding_Which uint16

const (
	BrandBinding_Which_unbound BrandBinding_Which = 0
	BrandBinding_Which_type    BrandBinding_Which = 1
)

func (w BrandBinding_Which) String() string {
	const s = "unboundtype"
	switch w {
	case BrandBinding_Which_unbound:
		return s[0:7]
	case BrandBinding_Which_type:
		return s[7:11]

	}
	return "BrandBinding_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// BrandBinding_TypeID is the unique identifier for the type BrandBinding.
const BrandBinding_TypeID = 0xc863cd16969ee7fc

func NewBrandBinding(s *capnp.Segment) (BrandBinding, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return BrandBinding{st}, err
}

func NewRootBrandBinding(s *capnp.Segment) (BrandBinding, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return BrandBinding{st}, err
}

func ReadRootBrandBinding(msg *capnp.Message) (BrandBinding, error) {
	root, err := msg.RootPtr()
	return BrandBinding{root.Struct()}, err
}

func (s BrandBinding) Which() BrandBinding_Which {
	return BrandBinding_Which(s.Struct.Uint16(0))
}
func (s BrandBinding) SetUnbound() {
	s.Struct.SetUint16(0, 0)

}

func (s BrandBinding) Type(e *capnp.ErrorSet) Type {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Type{}
	}
	return Type{Struct: p.Struct()}
}

func (s BrandBinding) HasType() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BrandBinding) SetType(v Type) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s BrandBinding) NewType() (Type, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewType(s.Struct.Segment())
	if err != nil {
		return Type{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// BrandBinding_List is a list of BrandBinding.
type BrandBinding_List struct{ capnp.List }

// NewBrandBinding creates a new list of BrandBinding.
func NewBrandBinding_List(s *capnp.Segment, sz int32) (BrandBinding_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return BrandBinding_List{l}, err
}

func (s BrandBinding_List) At(i int) BrandBinding { return BrandBinding{s.List.Struct(i)} }

func (s BrandBinding_List) Set(i int, v BrandBinding) error { return s.List.SetStruct(i, v.Struct) }

// BrandBinding_Promise is a wrapper for a BrandBinding promised by a client call.
type BrandBinding_Promise struct{ *capnp.Pipeline }

func (p BrandBinding_Promise) Struct() (BrandBinding, error) {
	s, err := p.Pipeline.Struct()
	return BrandBinding{s}, err
}

func (p BrandBinding_Promise) Type() Type_Promise {
	return Type_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Value struct{ capnp.Struct }
type Value_Which uint16

const (
	Value_Which_void        Value_Which = 0
	Value_Which_bool        Value_Which = 1
	Value_Which_int8        Value_Which = 2
	Value_Which_int16       Value_Which = 3
	Value_Which_int32       Value_Which = 4
	Value_Which_int64       Value_Which = 5
	Value_Which_uint8       Value_Which = 6
	Value_Which_uint16      Value_Which = 7
	Value_Which_uint32      Value_Which = 8
	Value_Which_uint64      Value_Which = 9
	Value_Which_float32     Value_Which = 10
	Value_Which_float64     Value_Which = 11
	Value_Which_text        Value_Which = 12
	Value_Which_data        Value_Which = 13
	Value_Which_list        Value_Which = 14
	Value_Which_enum        Value_Which = 15
	Value_Which_structValue Value_Which = 16
	Value_Which_interface   Value_Which = 17
	Value_Which_anyPointer  Value_Which = 18
)

func (w Value_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructValueinterfaceanyPointer"
	switch w {
	case Value_Which_void:
		return s[0:4]
	case Value_Which_bool:
		return s[4:8]
	case Value_Which_int8:
		return s[8:12]
	case Value_Which_int16:
		return s[12:17]
	case Value_Which_int32:
		return s[17:22]
	case Value_Which_int64:
		return s[22:27]
	case Value_Which_uint8:
		return s[27:32]
	case Value_Which_uint16:
		return s[32:38]
	case Value_Which_uint32:
		return s[38:44]
	case Value_Which_uint64:
		return s[44:50]
	case Value_Which_float32:
		return s[50:57]
	case Value_Which_float64:
		return s[57:64]
	case Value_Which_text:
		return s[64:68]
	case Value_Which_data:
		return s[68:72]
	case Value_Which_list:
		return s[72:76]
	case Value_Which_enum:
		return s[76:80]
	case Value_Which_structValue:
		return s[80:91]
	case Value_Which_interface:
		return s[91:100]
	case Value_Which_anyPointer:
		return s[100:110]

	}
	return "Value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Value_TypeID is the unique identifier for the type Value.
const Value_TypeID = 0xce23dcd2d7b00c9b

func NewValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func NewRootValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value{st}, err
}

func ReadRootValue(msg *capnp.Message) (Value, error) {
	root, err := msg.RootPtr()
	return Value{root.Struct()}, err
}

func (s Value) Which() Value_Which {
	return Value_Which(s.Struct.Uint16(0))
}
func (s Value) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Value) Bool() bool {
	return s.Struct.Bit(16)
}

func (s Value) SetBool(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s Value) Int8() int8 {
	return int8(s.Struct.Uint8(2))
}

func (s Value) SetInt8(v int8) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint8(2, uint8(v))
}

func (s Value) Int16() int16 {
	return int16(s.Struct.Uint16(2))
}

func (s Value) SetInt16(v int16) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint16(2, uint16(v))
}

func (s Value) Int32() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Value) SetInt32(v int32) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint32(4, uint32(v))
}

func (s Value) Int64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Value) SetInt64(v int64) {
	s.Struct.SetUint16(0, 5)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Value) Uint8() uint8 {
	return s.Struct.Uint8(2)
}

func (s Value) SetUint8(v uint8) {
	s.Struct.SetUint16(0, 6)
	s.Struct.SetUint8(2, v)
}

func (s Value) Uint16() uint16 {
	return s.Struct.Uint16(2)
}

func (s Value) SetUint16(v uint16) {
	s.Struct.SetUint16(0, 7)
	s.Struct.SetUint16(2, v)
}

func (s Value) Uint32() uint32 {
	return s.Struct.Uint32(4)
}

func (s Value) SetUint32(v uint32) {
	s.Struct.SetUint16(0, 8)
	s.Struct.SetUint32(4, v)
}

func (s Value) Uint64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Value) SetUint64(v uint64) {
	s.Struct.SetUint16(0, 9)
	s.Struct.SetUint64(8, v)
}

func (s Value) Float32() float32 {
	return math.Float32frombits(s.Struct.Uint32(4))
}

func (s Value) SetFloat32(v float32) {
	s.Struct.SetUint16(0, 10)
	s.Struct.SetUint32(4, math.Float32bits(v))
}

func (s Value) Float64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Value) SetFloat64(v float64) {
	s.Struct.SetUint16(0, 11)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Value) Text(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s Value) HasText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Value) SetText(v string) error {
	s.Struct.SetUint16(0, 12)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Value) Data(e *capnp.ErrorSet) []byte {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return []byte(nil)
	}
	return []byte(p.Data())
}

func (s Value) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) SetData(v []byte) error {
	s.Struct.SetUint16(0, 13)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Value) List(e *capnp.ErrorSet) capnp.Pointer {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		*e = append(*e, err)
		return nil
	}
	return p
}

func (s Value) HasList() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) ListPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetList(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 14)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetListPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 14)
	return s.Struct.SetPtr(0, v)
}

func (s Value) Enum() uint16 {
	return s.Struct.Uint16(2)
}

func (s Value) SetEnum(v uint16) {
	s.Struct.SetUint16(0, 15)
	s.Struct.SetUint16(2, v)
}

func (s Value) StructValue(e *capnp.ErrorSet) capnp.Pointer {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		*e = append(*e, err)
		return nil
	}
	return p
}

func (s Value) HasStructValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) StructValuePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetStructValue(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetStructValuePtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPtr(0, v)
}

func (s Value) SetInterface() {
	s.Struct.SetUint16(0, 17)

}

func (s Value) AnyPointer(e *capnp.ErrorSet) capnp.Pointer {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		*e = append(*e, err)
		return nil
	}
	return p
}

func (s Value) HasAnyPointer() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Value) AnyPointerPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Value) SetAnyPointer(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPointer(0, v)
}

func (s Value) SetAnyPointerPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPtr(0, v)
}

// Value_List is a list of Value.
type Value_List struct{ capnp.List }

// NewValue creates a new list of Value.
func NewValue_List(s *capnp.Segment, sz int32) (Value_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Value_List{l}, err
}

func (s Value_List) At(i int) Value { return Value{s.List.Struct(i)} }

func (s Value_List) Set(i int, v Value) error { return s.List.SetStruct(i, v.Struct) }

// Value_Promise is a wrapper for a Value promised by a client call.
type Value_Promise struct{ *capnp.Pipeline }

func (p Value_Promise) Struct() (Value, error) {
	s, err := p.Pipeline.Struct()
	return Value{s}, err
}

func (p Value_Promise) List() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Value_Promise) StructValue() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Value_Promise) AnyPointer() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Annotation struct{ capnp.Struct }

// Annotation_TypeID is the unique identifier for the type Annotation.
const Annotation_TypeID = 0xf1c8950dab257542

func NewAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation{st}, err
}

func NewRootAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation{st}, err
}

func ReadRootAnnotation(msg *capnp.Message) (Annotation, error) {
	root, err := msg.RootPtr()
	return Annotation{root.Struct()}, err
}

func (s Annotation) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s Annotation) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Annotation) Brand(e *capnp.ErrorSet) Brand {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return Brand{}
	}
	return Brand{Struct: p.Struct()}
}

func (s Annotation) HasBrand() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Annotation) SetBrand(v Brand) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Annotation) NewBrand() (Brand, error) {
	ss, err := NewBrand(s.Struct.Segment())
	if err != nil {
		return Brand{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Annotation) Value(e *capnp.ErrorSet) Value {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Value{}
	}
	return Value{Struct: p.Struct()}
}

func (s Annotation) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Annotation) SetValue(v Value) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Annotation) NewValue() (Value, error) {
	ss, err := NewValue(s.Struct.Segment())
	if err != nil {
		return Value{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Annotation_List is a list of Annotation.
type Annotation_List struct{ capnp.List }

// NewAnnotation creates a new list of Annotation.
func NewAnnotation_List(s *capnp.Segment, sz int32) (Annotation_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Annotation_List{l}, err
}

func (s Annotation_List) At(i int) Annotation { return Annotation{s.List.Struct(i)} }

func (s Annotation_List) Set(i int, v Annotation) error { return s.List.SetStruct(i, v.Struct) }

// Annotation_Promise is a wrapper for a Annotation promised by a client call.
type Annotation_Promise struct{ *capnp.Pipeline }

func (p Annotation_Promise) Struct() (Annotation, error) {
	s, err := p.Pipeline.Struct()
	return Annotation{s}, err
}

func (p Annotation_Promise) Brand() Brand_Promise {
	return Brand_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Annotation_Promise) Value() Value_Promise {
	return Value_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type ElementSize uint16

// Values of ElementSize.
const (
	ELEMENTSIZE_EMPTY           ElementSize = 0
	ELEMENTSIZE_BIT             ElementSize = 1
	ELEMENTSIZE_BYTE            ElementSize = 2
	ELEMENTSIZE_TWOBYTES        ElementSize = 3
	ELEMENTSIZE_FOURBYTES       ElementSize = 4
	ELEMENTSIZE_EIGHTBYTES      ElementSize = 5
	ELEMENTSIZE_POINTER         ElementSize = 6
	ELEMENTSIZE_INLINECOMPOSITE ElementSize = 7
)

// String returns the enum's constant name.
func (c ElementSize) String() string {
	switch c {
	case ELEMENTSIZE_EMPTY:
		return "empty"
	case ELEMENTSIZE_BIT:
		return "bit"
	case ELEMENTSIZE_BYTE:
		return "byte"
	case ELEMENTSIZE_TWOBYTES:
		return "twoBytes"
	case ELEMENTSIZE_FOURBYTES:
		return "fourBytes"
	case ELEMENTSIZE_EIGHTBYTES:
		return "eightBytes"
	case ELEMENTSIZE_POINTER:
		return "pointer"
	case ELEMENTSIZE_INLINECOMPOSITE:
		return "inlineComposite"

	default:
		return ""
	}
}

// ElementSizeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ElementSizeFromString(c string) ElementSize {
	switch c {
	case "empty":
		return ELEMENTSIZE_EMPTY
	case "bit":
		return ELEMENTSIZE_BIT
	case "byte":
		return ELEMENTSIZE_BYTE
	case "twoBytes":
		return ELEMENTSIZE_TWOBYTES
	case "fourBytes":
		return ELEMENTSIZE_FOURBYTES
	case "eightBytes":
		return ELEMENTSIZE_EIGHTBYTES
	case "pointer":
		return ELEMENTSIZE_POINTER
	case "inlineComposite":
		return ELEMENTSIZE_INLINECOMPOSITE

	default:
		return 0
	}
}

type ElementSize_List struct{ capnp.List }

func NewElementSize_List(s *capnp.Segment, sz int32) (ElementSize_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ElementSize_List{l.List}, err
}

func (l ElementSize_List) At(i int) ElementSize {
	ul := capnp.UInt16List{List: l.List}
	return ElementSize(ul.At(i))
}

func (l ElementSize_List) Set(i int, v ElementSize) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type CodeGeneratorRequest struct{ capnp.Struct }

// CodeGeneratorRequest_TypeID is the unique identifier for the type CodeGeneratorRequest.
const CodeGeneratorRequest_TypeID = 0xbfc546f6210ad7ce

func NewCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CodeGeneratorRequest{st}, err
}

func NewRootCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CodeGeneratorRequest{st}, err
}

func ReadRootCodeGeneratorRequest(msg *capnp.Message) (CodeGeneratorRequest, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequest{root.Struct()}, err
}

func (s CodeGeneratorRequest) Nodes(e *capnp.ErrorSet) Node_List {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return Node_List{}
	}
	return Node_List{List: p.List()}
}

func (s CodeGeneratorRequest) HasNodes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest) SetNodes(v Node_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNodes sets the nodes field to a newly
// allocated Node_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewNodes(n int32) (Node_List, error) {
	l, err := NewNode_List(s.Struct.Segment(), n)
	if err != nil {
		return Node_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s CodeGeneratorRequest) RequestedFiles(e *capnp.ErrorSet) CodeGeneratorRequestRequestedFile_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return CodeGeneratorRequestRequestedFile_List{}
	}
	return CodeGeneratorRequestRequestedFile_List{List: p.List()}
}

func (s CodeGeneratorRequest) HasRequestedFiles() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequest) SetRequestedFiles(v CodeGeneratorRequestRequestedFile_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewRequestedFiles sets the requestedFiles field to a newly
// allocated CodeGeneratorRequestRequestedFile_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewRequestedFiles(n int32) (CodeGeneratorRequestRequestedFile_List, error) {
	l, err := NewCodeGeneratorRequestRequestedFile_List(s.Struct.Segment(), n)
	if err != nil {
		return CodeGeneratorRequestRequestedFile_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// CodeGeneratorRequest_List is a list of CodeGeneratorRequest.
type CodeGeneratorRequest_List struct{ capnp.List }

// NewCodeGeneratorRequest creates a new list of CodeGeneratorRequest.
func NewCodeGeneratorRequest_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return CodeGeneratorRequest_List{l}, err
}

func (s CodeGeneratorRequest_List) At(i int) CodeGeneratorRequest {
	return CodeGeneratorRequest{s.List.Struct(i)}
}

func (s CodeGeneratorRequest_List) Set(i int, v CodeGeneratorRequest) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequest_Promise is a wrapper for a CodeGeneratorRequest promised by a client call.
type CodeGeneratorRequest_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequest_Promise) Struct() (CodeGeneratorRequest, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequest{s}, err
}

type CodeGeneratorRequestRequestedFile struct{ capnp.Struct }

// CodeGeneratorRequestRequestedFile_TypeID is the unique identifier for the type CodeGeneratorRequestRequestedFile.
const CodeGeneratorRequestRequestedFile_TypeID = 0xcfea0eb02e810062

func NewCodeGeneratorRequestRequestedFile(s *capnp.Segment) (CodeGeneratorRequestRequestedFile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequestRequestedFile{st}, err
}

func NewRootCodeGeneratorRequestRequestedFile(s *capnp.Segment) (CodeGeneratorRequestRequestedFile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequestRequestedFile{st}, err
}

func ReadRootCodeGeneratorRequestRequestedFile(msg *capnp.Message) (CodeGeneratorRequestRequestedFile, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequestRequestedFile{root.Struct()}, err
}

func (s CodeGeneratorRequestRequestedFile) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s CodeGeneratorRequestRequestedFile) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s CodeGeneratorRequestRequestedFile) Filename(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s CodeGeneratorRequestRequestedFile) HasFilename() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequestRequestedFile) FilenameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequestRequestedFile) SetFilename(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s CodeGeneratorRequestRequestedFile) Imports(e *capnp.ErrorSet) CodeGeneratorRequestRequestedFileImport_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return CodeGeneratorRequestRequestedFileImport_List{}
	}
	return CodeGeneratorRequestRequestedFileImport_List{List: p.List()}
}

func (s CodeGeneratorRequestRequestedFile) HasImports() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequestRequestedFile) SetImports(v CodeGeneratorRequestRequestedFileImport_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewImports sets the imports field to a newly
// allocated CodeGeneratorRequestRequestedFileImport_List, preferring placement in s's segment.
func (s CodeGeneratorRequestRequestedFile) NewImports(n int32) (CodeGeneratorRequestRequestedFileImport_List, error) {
	l, err := NewCodeGeneratorRequestRequestedFileImport_List(s.Struct.Segment(), n)
	if err != nil {
		return CodeGeneratorRequestRequestedFileImport_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// CodeGeneratorRequestRequestedFile_List is a list of CodeGeneratorRequestRequestedFile.
type CodeGeneratorRequestRequestedFile_List struct{ capnp.List }

// NewCodeGeneratorRequestRequestedFile creates a new list of CodeGeneratorRequestRequestedFile.
func NewCodeGeneratorRequestRequestedFile_List(s *capnp.Segment, sz int32) (CodeGeneratorRequestRequestedFile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return CodeGeneratorRequestRequestedFile_List{l}, err
}

func (s CodeGeneratorRequestRequestedFile_List) At(i int) CodeGeneratorRequestRequestedFile {
	return CodeGeneratorRequestRequestedFile{s.List.Struct(i)}
}

func (s CodeGeneratorRequestRequestedFile_List) Set(i int, v CodeGeneratorRequestRequestedFile) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequestRequestedFile_Promise is a wrapper for a CodeGeneratorRequestRequestedFile promised by a client call.
type CodeGeneratorRequestRequestedFile_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequestRequestedFile_Promise) Struct() (CodeGeneratorRequestRequestedFile, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequestRequestedFile{s}, err
}

type CodeGeneratorRequestRequestedFileImport struct{ capnp.Struct }

// CodeGeneratorRequestRequestedFileImport_TypeID is the unique identifier for the type CodeGeneratorRequestRequestedFileImport.
const CodeGeneratorRequestRequestedFileImport_TypeID = 0xae504193122357e5

func NewCodeGeneratorRequestRequestedFileImport(s *capnp.Segment) (CodeGeneratorRequestRequestedFileImport, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequestRequestedFileImport{st}, err
}

func NewRootCodeGeneratorRequestRequestedFileImport(s *capnp.Segment) (CodeGeneratorRequestRequestedFileImport, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequestRequestedFileImport{st}, err
}

func ReadRootCodeGeneratorRequestRequestedFileImport(msg *capnp.Message) (CodeGeneratorRequestRequestedFileImport, error) {
	root, err := msg.RootPtr()
	return CodeGeneratorRequestRequestedFileImport{root.Struct()}, err
}

func (s CodeGeneratorRequestRequestedFileImport) Id() uint64 {
	return s.Struct.Uint64(0)
}

func (s CodeGeneratorRequestRequestedFileImport) SetId(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s CodeGeneratorRequestRequestedFileImport) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s CodeGeneratorRequestRequestedFileImport) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CodeGeneratorRequestRequestedFileImport) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequestRequestedFileImport) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// CodeGeneratorRequestRequestedFileImport_List is a list of CodeGeneratorRequestRequestedFileImport.
type CodeGeneratorRequestRequestedFileImport_List struct{ capnp.List }

// NewCodeGeneratorRequestRequestedFileImport creates a new list of CodeGeneratorRequestRequestedFileImport.
func NewCodeGeneratorRequestRequestedFileImport_List(s *capnp.Segment, sz int32) (CodeGeneratorRequestRequestedFileImport_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CodeGeneratorRequestRequestedFileImport_List{l}, err
}

func (s CodeGeneratorRequestRequestedFileImport_List) At(i int) CodeGeneratorRequestRequestedFileImport {
	return CodeGeneratorRequestRequestedFileImport{s.List.Struct(i)}
}

func (s CodeGeneratorRequestRequestedFileImport_List) Set(i int, v CodeGeneratorRequestRequestedFileImport) error {
	return s.List.SetStruct(i, v.Struct)
}

// CodeGeneratorRequestRequestedFileImport_Promise is a wrapper for a CodeGeneratorRequestRequestedFileImport promised by a client call.
type CodeGeneratorRequestRequestedFileImport_Promise struct{ *capnp.Pipeline }

func (p CodeGeneratorRequestRequestedFileImport_Promise) Struct() (CodeGeneratorRequestRequestedFileImport, error) {
	s, err := p.Pipeline.Struct()
	return CodeGeneratorRequestRequestedFileImport{s}, err
}

const schema_a93fc509624c72d9 = "x\xda\xacY}\x94\x14\xd5\x95\xbf\xf7U\x7f\xccG\x17" +
	"\xdd5\xaf\x10\x8784\xa8\xa8\x10\x9d\xc0\x0c\xb08\x91" +
	"\x1d\x18\x18\x0c,\x90)\x1aP8\xeb\x095\xd35L" +
	"\x99\x9e\xea\xa6\xbaZi\x16\xcf\x80'\x1e\xb3\xee\xba~" +
	"\xacD\x97$\x9el\x96\x9c\x13W\xdd\xc0\xae\x9c\x13\x89" +
	"Y\xcd\x1cY\x81\x83\x89\xee\x9acHt7xBt" +
	"\xd9e\x85]P\x91\x8f\xdas_UW\xd74=\xd1" +
	"\x9c\xf8WW\xdf\xfb\xea\xbe\xf7\xee\xbd\xbf\xfbU\xb3>" +
	"\xdf\xbc\x90\xcd\x8e\xa6'\x00h\xbb\xa31\xf7\xf1\x13\x1b" +
	"\x9b\xa6\xdf\xfc\xee\xfd\xa0\xb5 s7\x1e9w\xec\xe4" +
	"\x96\xe2k0\x11\xe3\x08\xc0\xcf\xc6\xf6\x03\xf2\xb3\xb1n" +
	"@w\xdd\xbc\xf5\x97\xca_\xf9\xe2_\x81\xd6\x86\xcc=" +
	"\xbd\xec\x8f\xbf\xfb^\xf7m\xa3\xb0\x16\xe3\x18A\xecl" +
	"\x8do@@>=\xfe.\xa0\xfb\xf9\x1ec\xe3\xa9\xb5" +
	"s\x1e\x06EF\xf7\xa8\xbd\xa2\xbf\xf1@\xf7S\x10\xc5" +
	"8\x00\xbf\x18\xdf\xc5\xa3\x0d\xd7\x03\xf0\xd6\x86n\xc0\x17" +
	"\xb7\xaf\xec\xfc\xe2;Gvj2J\xa1\xa5QZ\xba" +
	"\xac\xe1\xbb\\k\x88\x03t\xaelx\x05\x01\xdd\x95\xbb" +
	"\xd6\xbc\xf5_\xdb\x1ex\x1c4\x19\xc3\x82\x19\xad^\xdb" +
	"\xb4\x9f\xdf\xd1DO\xeb\x9b~\x00\xe8\xb6\xbcri\xdb" +
	"?\xaf\xd8\xfb8(<\xe2~\xe5\xcc\xf2\xf4\xee\xbeg" +
	"w\x01`gcs\x0b\xf2\xd6\xe68@Fm\x960" +
	"3\xb5\x99!@u\xc9\xd8\xa3\xf4F\xe2\x0c#\\i" +
	"\xde\xc5[\x9b'\x01tNo~\x88\xcer\xdfr\xfd" +
	"\xcc\xb5\x1f\xbe\xfe\xed\x1a}x\x9a\xeb,'\xbaH\x1d" +
	";\x12w\x03\xbaO\xdd?2y\xf3\x96\x09O\xd6W" +
	"\xf2\xb1\x04)\xf9\x98X9w\xc9\xf9?\xfa\xe6\xde\xef" +
	"\x88\x95Qw\xd2\xeek\xce\xadx\xfa\xde\xdf\xc2\xc4\x98" +
	"X\xb9@>\x0c\xd8\xb9H\x16\xfb\x7f\xf4\xe6\xee\xd7\x9e" +
	"\x9c\xf9\x8d\xa7ju!\x94|n\xc2(\xc7\xa4P\xf7" +
	"\x04\x92{\xe8\xab\x03O\xdd7\xe7\xcd\xa7A\xe3\xc8\xaa" +
	"\xe6\xe9Eq\xb7\xf5\xc9\xc3\xdc\x10\xab\xf5$i\xee\xdf" +
	"\xfee\xc5\xfb\x1b\xf3]\xcf\xd4?oc\xea0 \x97" +
	"S$\xf7\xf8m\xd7\xb4\xfc\xf5\xa2\xbe\x7f\x00m\x16\xe2" +
	"\xc5\xfe\x1d\xed{&\x9c\xf8\x998B\xa7\x99\xda\x8f|" +
	"G\x8a\xa4\xde#\xd6N}L\xde\xfa\xcc\xf7\x1e\xd8[" +
	"\xffnGS\xa3\x80\xfc\xa8X\xf9\xc4\xd9\xce\xc5\xf3\xfe" +
	"i\xf9\xbe\xfa+\xe7*\xa4\xaf\xb9\x0a9\xe5\xde\xbfK" +
	"\x9e>\xf2\xb9\xd5\xcf\x83\xd2\x82\xd5\x85\x9e\x0eL\xe5\x1d" +
	"^R\xe8i\xb3X\xeb\xbc\xb3.\xd1r\xe8\xfd\xfd\xf5" +
	"\x0d\xf6\x88\xf2\xf7d\xb0\xbf\x15K\x7f\xdb|\xff\xfd\xa3" +
	"?\x7f\xf8G\xa4.\xa9\xea\x18k#qd\x18\xe5\x07" +
	"\x94_\x02\xf2C\x0a\x9d\xf5\xa7o6M\xfb`\xe9\x81" +
	"\x17k\\\x9d<\xb2\xf3\xa6\x96\x16\xe4\x0bZ&\x91+" +
	"\xb7\xd0\xe2`\xd3\xb1\x8a%\x0cI\xc8\xf8\xa1\x96\x13\x80" +
	"\xfc\xd5\x16\xb2A\xdey\xee\xab\xb7F\xafy\xb9\xe6\x08" +
	"\x13#B\x07\xd38ik\x1a'\xb0]x\xf7\xc9o" +
	"\\\xf1\xea\xc0AZ\x89cm\x0b\xc0Q\xfd%\x97U" +
	"\xd2B\xa3JGP\xda~5\xf4\xabW/\x1c\xae/" +
	"\xd7P\xc9\xb6\xa6JJ\xf8fb\xcf\x9b\xff\xfa\xd65" +
	"?%\x07c!<`\x9c\x03\xf0?Ww\xf1GH" +
	"n\xe7\x83\xea\x17\"\x10\x18_\xbb\x1aCJ\xf1\xf40" +
	"\xb1\xf5^\xe43ZI\x0f\xb3[\xe9r\xc1\xcd\xeb\x8a" +
	"~\xbb\xf5Q~\xbc\x95^<\xd6J\xa2\xdd\xebZW" +
	"\xef\xbf\xe7\xa1\x9d\xaf\x83\"\x87V\x03\xf2\xb7\xdb\x0e\xf3" +
	"\xffl\xa3\xcb\x1do{\x05\xd0\xed\x19M~\xfc\xa3\xb5" +
	"/\xfe;\xe9\xf72\x7fX?\xe5\x047\xa6\x08/\x9f" +
	"B\x9a\x08\xd8\x9a\x8c\xd1\xd0!b\xf1\x18\xc6\xf8\x8f\xa7" +
	"<\xca\x0fL\xb9\x9e\x0e1e\x92\x04\xe8>4m\xf4" +
	"\xd4\xcf2\xd7\xbfW\xdf)/N{\x07\x90\xe3\xd5$" +
	"\xf8a\xd6\xb4\xf0\x8d\xd6+\xfe\xbb\xfe\xca\x1dW\x9f\x00" +
	"\xec\xbc\xef\xea\xff`\x80\xee\x8b\x89\xf3\x1f\x9a\x87\xff\xe2" +
	"d}\xa4\xed\x9dNB\xf7M'\xa1=\xa5\xe9O\xcb" +
	";\x0f\x9e\xae\x1b\xfa\x94\xebFy\xebu\xf44\xf1\xba" +
	"\x1f\xc0*\xb780d\x0c\xeb\xed\x03\xa8\x17\xacB\xd7" +
	"\x9ar\xa1\xdbh\xcf\x99EG\x8bH\x11\x00\x15'\x00" +
	"(r?\x80\x96\x90P\xbb\x92\xa1k\xe4\x8ca\xc3r" +
	"\xd6@\xbc\\00U=\x09 \xa6\x00\x03\x81\x91\x8a" +
	"@\xa3]\xb7\xca}y\xd3r\x0c\xbb\xbdd\x0d\xe4\xad" +
	"\xa2c\xeb\xa6%\x19Y-%E\x12\xae\xabb\x0b\x80" +
	"\xa2\xf7\x00h\x7f*\xa16\xc4P\xc6K\xae\x8a\xad\x00" +
	"\x8a\xd1\x05\xa0m\x94P\xcb1\x94\xd9EW\xc5\xc9\x00" +
	"\x8a9\x13@\xcbJ\xa8\x15\x18\xca\xd2\x05W\xc5\xcf\x01" +
	"(\xc3\x1b\x00\xb4\x9c\x84\xda\x16\x86#\xbaU\xfe\x13\xd3" +
	"\xcaB\xac\xbb\xe8\xd8\xa5\x01\x07bI\xba\x17\xc4\xdc\x01" +
	"\xbd\xa0\xf7\x9b9\x13$\xa7\x0c\xb1\x1a\x0d\xf4\xd8\xbad" +
	"e\xb5\x06\x0cEB\xa5\xb1\xa3\x0a\x1d%\xda\x93\xce\x0c" +
	"\xe4\x0b\xc6H\x8fieMk\x93\xa7\xa9\x08\x92\xa2\xe8" +
	"\xb0\x0d\x12j\xd72\xec.\xd2\xa2\"N\x00\xec\x93\x10" +
	"SUq\x80D\xac\xd9w\xa5\xe1\xc4\x87\xf2\xd9>D" +
	"mj \xefu\xba\xe6\x11\x09\xb5_0DT\x91h" +
	"?_\x0d\xa0\xbd!\xa1\xf6k\x86\x8a\x84*2\x00\xe5" +
	"\xed{\x01\xb4\xb7$\xd4\xdec\xa8D\x99\x8a\x12\x80r" +
	"\xfc\x01\x00\xed=\x09\xb53\x0c\x958\xaa\x18\x01PN" +
	"\x93-OI\x98I C%\xc2T\x8c\x12\xe0q\x03" +
	"@\xa6\x01%\xcc\xa8D\x8fI*\xc6\xc8[\xb0\x1f " +
	"\x93\"\xfaUDg\x11U\x00\xa4\x15w\x01d\xae\"" +
	"\xfa|d\x98\xb4\xf4a\x03\x13\xc00\x01\xe8\x0e\xe4\xb3" +
	"\xc6\x97\xed\xac\x01hc\x1c\x18\xc6\x01\xdd\x82n\xeb\xc3" +
	"\x19\xc7\xc6\xd2\x80C>\x01\xd8\x08\x0c\x1b\x01]\xdb(" +
	"\x96rN\xc6A\xbb\xc2\xaa\xf2t\xcb\xca;\xbacB" +
	"<o\x854\x198\xb8\xafI!\xbc\xc7\xd6A\xb2\xb2" +
	"\x98\xaa\xc65\xdf#\xbd\x1dzl\x88\xebu\xf9\xe6p" +
	"!g\x0e\x98\x0e\xf6\x91\x1c\xc31$;\xb4Y\x907" +
	"\xea\x9a\xad\xd7*u\x0f\x1b\xb6n9d\xb9D`\xb9" +
	"^\xb2\xdcB\x09\xb5\x15U\xcb-#\xcb}IBm" +
	"\x0di\xd2\xb7\x9cF\xf6\xe8\xf3\xfc\xfb\x93\xd5\xf8)\x15" +
	"R9\xa3$\xce\xb8\xd44r\xd9v+\xbf\xc4,\x0e" +
	"\xd8\xe6\xb0i\xe9\x16\xd2qI\xaa\x1cw\xdd\xcb.\xb5" +
	"\xd44\xa4\\V\x8b`\xb82\xc2\xadnE\x04t\x0b" +
	"!\x8evUp\xdf}t\xdf=\x12j/T\xef\xfb" +
	"<\xdd\xf7\x87\x12j/\x87\xee\xfb\x13\xba\xefK\x12j" +
	"o\xf9\xee+!*G\x1f\xad\xba\xaf\x1cq]\x0c%" +
	"5\xe5\xf8L`r\xf4\x12\x11\x83\x8c\xa4\xbc\xda\x01\x0c" +
	"c\x18\xca\xbf\xca\xbe\x1e`\x9f\x99\x06\xb3\xbe\xae\x90\xee" +
	"\xb9N\xcf\x95\xd0\xa8\xaa+Y\xcc\xe5\x9d\xf4&;_" +
	"*\x8c\xe4\xed\xaci\xe9\xb9\x1a\x95\xd7\x86\xbd\x82nw" +
	"\x0b\xd7\xb2\xb5\x06)\x92B\x159\x802\x83\"\xde\xb5" +
	"\x12j\xb3\x18*\xa45\x15@\xb9i+\x80v\xa3\x84" +
	"\xda|\x86#\"\x86,\xcb\x06\x80(\xf8\x1e\x0a\xdd\xf6" +
	"2+kl\x09\xaeU/\x8c\x1bVi\xb8\xb2]\x92" +
	"\xb6\xeb\xaanG\xbbM\xa4\xdd:\x00\xb4\x1b$\xd4\xe6" +
	"0\xecv\xca\xe1\xcd\xd2\xfdv}\xbcT\xf6bb\xaf" +
	"U\xf9\xac\xd1\xee\x87X \x8fH\xc5E\x8cP\xf6\xd9" +
	"\x00\xdas\x12j/\xd1\xed\x12*6\x00(?\xbe\x13" +
	"@{AB\xed \xf9\x84\xacb#\x80r\xe0\x1f\x01" +
	"\xb4\x83\x12jo\x90O\xfcZ\xc5&\x8a}=\xd5\xd8" +
	"\xa7D\x92*6S\xf0#G\xf9\x85\x84\xdao(\xce" +
	"5\xa8\x98\x00P\x8e\xed\x02\xd0~#\xa1v\xca\x8f\\" +
	"2\x80r\xb2\xcb\x0b~\x99\x082t\xb3\xba\xa3\xdf\x96" +
	"\xb7\xb3\x90^\x9c/YN5,y\xf6Y\x0c\xc9\xb1" +
	"d\xdb\x184l\xdb\xc0\xec\x0a\xb3\xe8\xf4Z\x03\xe9<" +
	"EzLV\xeb\x0b@L\x02\x8e\x98\xc5[\xc9\x0d\x10" +
	"\x81!\xd6\xfa\x0d\xed\x85U\xa9cx_\x1e\x1c,J" +
	"\x86\x83\x0d\xc0\xb0\x01\xb0{\x90@\x1ar\xc8P\x172" +
	"\x06\xd2\x9e\xd23\xa5\x82a\x0f\xe4\xf4b\x11(\xee4" +
	"\x048\x9c1y\xac\x8d\xb1\xc6\xc6\x92\xf9\xfb\xda\xb7\x87" +
	"V\xb5\x8b\x8c\x07P\x13\xe4z\xaaANF\xca\xe3\"" +
	"\xccQ(X\"\xa1\xb6\x9126\xe5q\xc2\xfd\x1d\xb4" +
	"\xf6v\x09\xb5\xec\xe5n\x9d\xec7\xadl\xf5\xeaA\xaa" +
	"\xf5\xae>bZC\x86m:\xa1T\xcd\xaa \x0b<" +
	"\xcfw\xf4T\x1dG\xbf\xe2\x0ft\xf4\xa8\xd8nq>" +
	"k\xdcjX\x86\xad;y{\xb5\xb1\xb9d\x14\x9dv" +
	"\xff\xd7\xc8.5sF{\xf7\xb2\xe1B\xdev>\x85" +
	"If\xd65\xc9\xd8\xf0U\x07g\xa2x\xf2\xee*\xf9" +
	"w\x9d\x19\x8e!\x11\x15\x95\x9a\xcb&\x9d\xbauZ\xfa" +
	".=W\"zP\xbb\xd7\\\x1a+\xbbV\"I\xc4" +
	"\xdbT\x14\x84\x1b\xfc\x82\xf0\x06*\x08\xad\x92\xc8\x80 " +
	"9!\xff\x0d\xfa\xef\xba\xfe+.\xe3\xa5\xdb\xb8c\xd8" +
	"\xa4\xb0P\x155\xd3\xaf\xa2T6\x8eF\xa2u\x83l" +
	"%\x8f\xaf4\x1c\xaa\xa4HzR\x84\xdc\x88\xe7\x1aW" +
	"\x92\xec\xad\xa1R\xf6\x13\xa3)\x0b\xe5O\x11\xe8%=" +
	"G\xca\x175+\x85\xb3\x19\xcb\xab\x8an\xc3Kn\xcc" +
	"\x0bh\xb3\x89<KB\xed\x16V-/\x00\xe2\xae\xb1" +
	"\xa5\xf2\x0c\x97m&\x8d\xe7e\xe8P*\xaetM\x0a" +
	"\xdan\xc5\xeb -\xfc.\xecl\x1d\xbe;,\xf4S" +
	"\x0a\x11\x17\xd0\xa5o\x91P\xbb\x9da\xda\xcag\xc3U" +
	"i\xd0s\xf8v\xb2+\xa2\xbb\x85\xe8\xea\xca\xca\xfeu" +
	"\xed\x192E\x92lAA\"!\xf2x0\x0bRz" +
	"m`T\xd7ch \xa2\xcc^\x0d\x8c\xcaz\x0c5" +
	"\xddJ\xdb(0\xb7\xd2)@Z7-#[5\x16" +
	"\xda\x81J\xa3!S\x93\xa5\xa1r\xaa\x90\xdd(W\x03" +
	"5\x1a)\xbf\x02\xd6\xbb\xaa}\x86\x82^\x01\xac\x183" +
	"\xabm\x86\xc2\xbc\xeaW1)Y\x0dI\xa89\x94\x97" +
	"\xb6{yi3e\x1bGBm;\xc3\xee\xfc\xe0`" +
	"\xb1\x1a\xc4\xc7\x81\x9a\x9b5\x06\xf5R\xceY\x07\xc9q" +
	"07\xa4g{\xc93p\xc0t\x96\xd0b)\xe7\x04" +
	")\xa5^8\x16\xad\x87dm\xf2#M\xc2u=\xeb" +
	"\x87\x0a\x0a\xaf\x85\xaa\x0d6#%\xab?_\xa2\xb6h" +
	"\xbc\xc3\xd6\xf3}Q\xea\x00\x08 1/\xbb\x07\xad\x8e" +
	"zY<\xad\x09!\xeb\xf4\x9cT2\xe8\xa4+\x82\x93" +
	"\xf2(\x9b\x09\xb0\x9aQ\x1f\xc2<\xe8\xa4\xc4ay#" +
	"12\x11\xe2\xa4\x88\xc3.\xbaLd\x10.\x0bN\x03" +
	"qT\xe2H\x17\\\x14\x8d\x0eWX\x07@&A\x9c" +
	"+\x89\x139\xefz\xb6\xe6\x13\x05'E\x9c\xab\x88\x13" +
	"\xfd\x988\xd4\xf1\xb4\x0a\x8eJ\x9c\xa9\xc4\x89\x9d\xa3}" +
	"\xa8\xe7i\x13\x9c+\x89s-q\xe2\x1f\xd1;\xd4\xf5" +
	"Lc]\xd4\xf5\x10\xe7\x06\xe24|H\x9c\x06\x00>" +
	"]p\xa6\x12\xe7F\xe24~@\x9cF\x00>Cp" +
	"\xae%\xce,\xe24\x9d%N\x13\x00\xbf\x89\xf5\x00d" +
	"n \xce\x1c\xe24\x9f!N3\x00\x9f-87\x12" +
	"g>c('\xfe\xcf\x15\x95\x0e\x9f+T0\x8b\x18" +
	"\xb7\x10C\xfe_WT;\xfcf\xc1\x98C\x8c\x85\xc4" +
	"\x98p\xda\x15\x91\x9a/\x10\x8c\xf9\xc4XB\x9b$O" +
	"\xb9^5\xc8\x17\x09\xce-\xc4\xf9\x12\xbd\x92z\xdf\x15" +
	"\x19\x85\xf7\xb2.\xde\xcb\xd2\x99!b9\xc4R\xfe\xc7" +
	"\x15y\x85of\xab\x012\x05bl#F\xcbI\xaf" +
	"w\xe7eF\xfd\xe3\x16b|\x8d1L\xde\x957\xc9" +
	"\xc5\xfa\xf3\xf9\\\xc5\x8f\x93\xa6\xe5\xccG\x06\x0c\x19`" +
	"\xda\xb4\x9c\xd9\xf3P\x02\x86\x92\xf7\xaf\xb3\x03#\xc00" +
	"\xe2\xfd\x9b7\x07\xa3\xc00\x0a\x98.\x89\xf7bT\xee" +
	"\x03v\x97\xbc\x17\xfd\xd8)\xfevv\x04ET\xc9{" +
	"\xd5\xf7\xc3\x91\xc1\\^'v\x130l\xaa\xfc\x9f7" +
	"\x07\x9b\x81a3\xc1\xd5\xd8\xe2TrK\x92JD\x94" +
	"\x81\xa1\x0c(&\x05\xd8\x02\x0c[\x00\x93\x94\xe0\x82\x1d" +
	"\xbdzC\x8b sO?\xf8\x85I-\x1b\x9f\x1f\x05" +
	"-\xc2p\x91\x8a\xa2\x18\xc5~\xd7[\xb2N\x87x\xae" +
	"D\xdd\xac/\xc7\x15YjP\x1f\x004 \xe6V\x12" +
	"\x17H\x86\x1d,\xf9}+\x0e\x10MZ0ZU\xb0" +
	"\xcb/A\xc2U\xda\xe4:\xad\xe8\xf2z\xad(E\x8f" +
	"\x15^`\x0c\x15%\xee\xa0\x993(\x0dS\xd2\xf2\xb5" +
	"5b\x8a]By$8C\xdd6yM\xb9\x80\xbf" +
	";\x04\x04\xe1\xaa6\x02xC\x9f:\x01\xc0\x9b\xfb\xd4" +
	"\xc1\xbf\x1c9\xef\xd6\x85\xbf\x1c\xfd\xd8\xad\x8b~9v" +
	"\xce\xad\x0b~9\xfe\x91[\x17\xfbr\xc3\x87n]\xe8" +
	"\xcb\x8d\x1f\xb8u\x91/7\x9du\xeb\x02_n>\xe3" +
	"~\x86\xb8G\xac~\xa9\x11\xf0g2\x81>\xf4i\x81" +
	"O\x17\xd4\xd4\xfbD\x0d\x06\xf8\\f]\\fi\x10" +
	"x\xc7\xd0\xc0\x91\xb71\xca\xd2-'\x89\x1a\xcc\xa5y" +
	"#\xdb\x00,\x8ct\x88\x09\x8cC\xccC\xb7\xf7\xdb\xd9" +
	"\xe1\xfd\xce\x9b\x031\x0f\xcb\x10\xf3Q\xec?\xd0\x0a\x1f" +
	"\xb9\x10\xab`\xb6\xf2D4\x81S\x88\x09\x84\xfaS<" +
	"\x81\xc9O\xc6\xe2\x06\x1f\x8bk\xca Q\xfb\x12\x86\xe0" +
	"\x18\x00\xd6d\xbd^o\xc6\x991\xb7\xfa=\xcfT\x01" +
	"\x90\xf5\x1d\xe4\xdd\x8av5\x002\xd1\xe7\xa0\xa4,Z" +
	"\x0e\x80\x11e\xc1j\x00\x8c*7o\x00\xc0\x982\xb7" +
	"\x87\x0a<\xe5\xa6{\x01\xd2\xc6p\xc1)\xc7\xfbM'" +
	"\xd9_v\x0c\xd7\xb9;\xdfSv\x8c\"\x00\xb8\x83\xf9" +
	"\x92M\x7f\x00\x8b\xaean\x1arz\xcat\x9c\xe2\x88" +
	"\xdf\x97\xba\xa6\x953-cq>2\\\xc8\x17M\xc7" +
	"\x18[\xe0\x88*z\x95\x08\x06Iz\xae\xe9;f\xd6" +
	"\xe9;&\x87\x1a\x83pi\x1d\xc6\xfbe\x1d\x00\x1ab" +
	"\x1eZ\xfd\x82\xd2\xb8\xba:?W\x1a7\xb8}\xa1\xda" +
	"\xcc;\xcf\xaa<HYC\x9b_9\x0d/\xe3d\x80" +
	"\x8c\x83\x12f\xb6cp ~\x8f\x180n#\xf2\xd7" +
	"\xc5\x80\xd1O\xf6\xf7\xe1(@\xe6\xebD\x7f\x8c\xe8\x92" +
	"7\xd3\xe4\x8f A\xe4/\x89\xfe\x84\x18T\xfa\x89~" +
	"\xa7\x90\xf3\x18\xd1\xf7\x10=\xee\x0f6\x9f\x15\xf4g\x88" +
	"~\x10\x09\xb6\xae\x0f\xf4\x03H\xc8y\x89\x18G\x88\xd1" +
	"(\x0a\xd3\xe0\xa3\x1a?\x84]\xfc\x10\x12\x1a\x9aDu" +
	"\x1a|h\xe2G\x91\xd0\xd3|\x81\xa8\xc1\xa4\x9f\xff\x04" +
	"\x09#\x89\xf3D\x0d>_\xf1g\xb1\x03\x98,\x7fL" +
	"\xd4`\xd6\xcfw\xe2\x06`J$\xaa\xe2T\x00\xbeC" +
	"\xcc^\xb7\xd3I\xbe\x85\x0cgG\xa7\xa2\x8a\xd3\x00\xf8" +
	"\xdf e\xdb'\x88\xb1\x1b\xc7\x86\xe4\xacY,\xe4\xf4" +
	"\xf2*\x88\x87\xdb\xa3\x0a\x95\xe9\xc3F\x9fm\x0c\x9a[" +
	"V\x18\xd6&g\x08*\x09\xf2\xb2y\x92U1V|" +
	"LS\x10\x18\xd7\x0f\xe6\x9fnh\x96\xa4D\x11\xcc\xdb" +
	"?\x05,='\x01\x10h\x0ec3-J\xff\xea\xb6" +
	"R\xde\xaa\xd6\xff\xbfs:k\x16E\xc24\x01\x07\xc6" +
	")\xa0\x05h\xbc\xbd\xe2\xfa\x80Q\xe9\xa5\x93\xe1\xf2y" +
	"\xa1\xdfK\xa7\xa9y\xba\xb3\xda<\x8d\x0c\x8b~\xa3\xba" +
	"\x7f\xe5\x93t%\xe9\xf9c\x19H\x16\x8ba}\x06\x1f" +
	"`\xc7o\x86\xbd\xbb&\x1d3ois\xbc#Qa" +
	"u\x87p\xd3\xdb\xc9\x07\xb2\xe4\xd7X\x10\xa3C\xae\x0b" +
	"\xbf\xdeH\xf4\x9c\xc0\xcdf1=\xe4&\xde\x09\x90\x19" +
	"\"\xba#pc\x8b9\x1f\xdf,\xd6\x17\x88\xbeM\x0c" +
	"\xfe\x8bb,\xc2\xcb\xf8\xc0\x18\xfcE\x1d\x15'\x09\xfc" +
	"\xd9\x00\x99\xaf\x11\xfda\x81\xb3\x92\xe8\x9f\xf9\x83B~" +
	"\x15\x7f\xf1\xbb\xc4g\x1a\xbeS\xd0\x05\xfe\xbeC\xf4\x86" +
	"\xbb\xc5\x87\x1a\xfemA\xff\x16\xd1\xbfO\xf4\xc6-\xe2" +
	"S\x0d\xff\x9e\xd8\xf7\xfbD\x7f\x8e\xe8Me\x15\xaf\x02" +
	"\xe0{\xc5\xbe{\x88\xfe\x02\xd1\x9b\xb7\xaa\xd8\x06\xc0\x9f" +
	"\x17r~H\xf4\x97\x89\x9e\xf83\x15\xa7\x00\xc1\xeeQ" +
	"\x80\xcc\xcbD\x7f\x0d\xc7\x1dv\xb8\x8eno2\x9c\xe2" +
	"R\x88\x9b9#p\x0c\x9f\xba\x18\x92\xe4n\xb5\xe4^" +
	"\x88S\xe1WKE\x7f\xdc\xe1\x90\xec\xb1\xbc\x0c\xa4\x85" +
	"_\xd7\xd2\x97B\x92\xda\xa8Z\xf2ZHZf\xde\xaa" +
	"%\xdf\x0a\xc91#E\x9f\xbc\x0c}|\x18\x97o\xbc" +
	"\x12\xd2\xc21k\xe9}\x90$\xd4\xd4\x92\x17\xa1\x0f\xac" +
	"<Z\xe3\xa0D4\xf45(Ao\xb8T;]\x9b" +
	"\xf4\x99\x8c\x91\x17U\x8ed\xd5\x0e\x19\xc3\xe5+\xf3\xcb" +
	"\xd7\x0e\x7f\xc4\xd8\xe7O9(;\xaf\xec\xa8\xd6\xb4\xe1" +
	"1\xe78\x83\xae\xf1\xce\xf5\xff\x01\x00\x00\xff\xff\xb0p" +
	"\xbc\xbd"

func init() {
	schemas.Register(schema_a93fc509624c72d9,
		0x87e739250a60ea97,
		0x8e3b5f79fe593656,
		0x903455f06065422b,
		0x9500cce23b334d80,
		0x978a7cebdc549a4d,
		0x97b14cbe7cfec712,
		0x9aad50a41f4af45f,
		0x9dd1f724f4614a85,
		0x9e0e78711a7f87a9,
		0x9ea0b19b37fb4435,
		0xa9962a9ed0a4d7f8,
		0xabd73485a9636bc9,
		0xac3a6f60ef4cc6d3,
		0xae504193122357e5,
		0xb18aa5ac7a0d9420,
		0xb54ab3364333f598,
		0xb9521bccf10fa3b1,
		0xbaefc9120c56e274,
		0xbb90d5c287870be6,
		0xbfc546f6210ad7ce,
		0xc2573fe8a23e49f1,
		0xc42305476bb4746f,
		0xc863cd16969ee7fc,
		0xcafccddb68db1d11,
		0xce23dcd2d7b00c9b,
		0xcfea0eb02e810062,
		0xd07378ede1f9cc60,
		0xd1958f7dba521926,
		0xdebf55bbfa0fc242,
		0xe682ab4cf923a417,
		0xe82753cff0c2218f,
		0xec1619d4400a0290,
		0xed8bca69f7fb0cbf,
		0xf1c8950dab257542)
}
