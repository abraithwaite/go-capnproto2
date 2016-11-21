package rpctwoparty

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Side uint16

// Values of Side.
const (
	Side_server Side = 0
	Side_client Side = 1
)

// String returns the enum's constant name.
func (c Side) String() string {
	switch c {
	case Side_server:
		return "server"
	case Side_client:
		return "client"

	default:
		return ""
	}
}

// SideFromString returns the enum value with a name,
// or the zero value if there's no such value.
func SideFromString(c string) Side {
	switch c {
	case "server":
		return Side_server
	case "client":
		return Side_client

	default:
		return 0
	}
}

type Side_List struct{ capnp.List }

func NewSide_List(s *capnp.Segment, sz int32) (Side_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Side_List{l.List}, err
}

func (l Side_List) At(i int) Side {
	ul := capnp.UInt16List{List: l.List}
	return Side(ul.At(i))
}

func (l Side_List) Set(i int, v Side) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

func (l Side_List) Copy(seg *capnp.Segment) (Side_List, error) {
	ul := capnp.UInt16List{List: l.List}
	t, err := ul.Copy(seg)
	return Side_List{t.List}, err
}

type VatId struct{ capnp.Struct }

// VatId_TypeID is the unique identifier for the type VatId.
const VatId_TypeID = 0xd20b909fee733a8e

func NewVatId(s *capnp.Segment) (VatId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VatId{st}, err
}

func NewRootVatId(s *capnp.Segment) (VatId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VatId{st}, err
}

func ReadRootVatId(msg *capnp.Message) (VatId, error) {
	root, err := msg.RootPtr()
	return VatId{root.Struct()}, err
}

func (s VatId) String() string {
	str, _ := text.Marshal(0xd20b909fee733a8e, s.Struct)
	return str
}

func (s VatId) Copy(seg *capnp.Segment) (VatId, error) {
	var err error
	t, err := NewVatId(seg)
	if err != nil {
		return t, err
	}
	t.SetSide(s.Side())
	return t, nil
}

func (s VatId) Side() Side {
	return Side(s.Struct.Uint16(0))
}

func (s VatId) SetSide(v Side) {
	s.Struct.SetUint16(0, uint16(v))
}

// VatId_List is a list of VatId.
type VatId_List struct{ capnp.List }

// NewVatId creates a new list of VatId.
func NewVatId_List(s *capnp.Segment, sz int32) (VatId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return VatId_List{l}, err
}

func (s VatId_List) At(i int) VatId { return VatId{s.List.Struct(i)} }

func (s VatId_List) Set(i int, v VatId) error { return s.List.SetStruct(i, v.Struct) }

func (s VatId_List) Copy(seg *capnp.Segment) (VatId_List, error) {
	var err error
	t, err := NewVatId_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VatId_Promise is a wrapper for a VatId promised by a client call.
type VatId_Promise struct{ *capnp.Pipeline }

func (p VatId_Promise) Struct() (VatId, error) {
	s, err := p.Pipeline.Struct()
	return VatId{s}, err
}

type ProvisionId struct{ capnp.Struct }

// ProvisionId_TypeID is the unique identifier for the type ProvisionId.
const ProvisionId_TypeID = 0xb88d09a9c5f39817

func NewProvisionId(s *capnp.Segment) (ProvisionId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ProvisionId{st}, err
}

func NewRootProvisionId(s *capnp.Segment) (ProvisionId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return ProvisionId{st}, err
}

func ReadRootProvisionId(msg *capnp.Message) (ProvisionId, error) {
	root, err := msg.RootPtr()
	return ProvisionId{root.Struct()}, err
}

func (s ProvisionId) String() string {
	str, _ := text.Marshal(0xb88d09a9c5f39817, s.Struct)
	return str
}

func (s ProvisionId) Copy(seg *capnp.Segment) (ProvisionId, error) {
	var err error
	t, err := NewProvisionId(seg)
	if err != nil {
		return t, err
	}
	t.SetJoinId(s.JoinId())
	return t, nil
}

func (s ProvisionId) JoinId() uint32 {
	return s.Struct.Uint32(0)
}

func (s ProvisionId) SetJoinId(v uint32) {
	s.Struct.SetUint32(0, v)
}

// ProvisionId_List is a list of ProvisionId.
type ProvisionId_List struct{ capnp.List }

// NewProvisionId creates a new list of ProvisionId.
func NewProvisionId_List(s *capnp.Segment, sz int32) (ProvisionId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return ProvisionId_List{l}, err
}

func (s ProvisionId_List) At(i int) ProvisionId { return ProvisionId{s.List.Struct(i)} }

func (s ProvisionId_List) Set(i int, v ProvisionId) error { return s.List.SetStruct(i, v.Struct) }

func (s ProvisionId_List) Copy(seg *capnp.Segment) (ProvisionId_List, error) {
	var err error
	t, err := NewProvisionId_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// ProvisionId_Promise is a wrapper for a ProvisionId promised by a client call.
type ProvisionId_Promise struct{ *capnp.Pipeline }

func (p ProvisionId_Promise) Struct() (ProvisionId, error) {
	s, err := p.Pipeline.Struct()
	return ProvisionId{s}, err
}

type RecipientId struct{ capnp.Struct }

// RecipientId_TypeID is the unique identifier for the type RecipientId.
const RecipientId_TypeID = 0x89f389b6fd4082c1

func NewRecipientId(s *capnp.Segment) (RecipientId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return RecipientId{st}, err
}

func NewRootRecipientId(s *capnp.Segment) (RecipientId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return RecipientId{st}, err
}

func ReadRootRecipientId(msg *capnp.Message) (RecipientId, error) {
	root, err := msg.RootPtr()
	return RecipientId{root.Struct()}, err
}

func (s RecipientId) String() string {
	str, _ := text.Marshal(0x89f389b6fd4082c1, s.Struct)
	return str
}

func (s RecipientId) Copy(seg *capnp.Segment) (RecipientId, error) {
	var err error
	t, err := NewRecipientId(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

// RecipientId_List is a list of RecipientId.
type RecipientId_List struct{ capnp.List }

// NewRecipientId creates a new list of RecipientId.
func NewRecipientId_List(s *capnp.Segment, sz int32) (RecipientId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return RecipientId_List{l}, err
}

func (s RecipientId_List) At(i int) RecipientId { return RecipientId{s.List.Struct(i)} }

func (s RecipientId_List) Set(i int, v RecipientId) error { return s.List.SetStruct(i, v.Struct) }

func (s RecipientId_List) Copy(seg *capnp.Segment) (RecipientId_List, error) {
	var err error
	t, err := NewRecipientId_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// RecipientId_Promise is a wrapper for a RecipientId promised by a client call.
type RecipientId_Promise struct{ *capnp.Pipeline }

func (p RecipientId_Promise) Struct() (RecipientId, error) {
	s, err := p.Pipeline.Struct()
	return RecipientId{s}, err
}

type ThirdPartyCapId struct{ capnp.Struct }

// ThirdPartyCapId_TypeID is the unique identifier for the type ThirdPartyCapId.
const ThirdPartyCapId_TypeID = 0xb47f4979672cb59d

func NewThirdPartyCapId(s *capnp.Segment) (ThirdPartyCapId, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ThirdPartyCapId{st}, err
}

func NewRootThirdPartyCapId(s *capnp.Segment) (ThirdPartyCapId, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return ThirdPartyCapId{st}, err
}

func ReadRootThirdPartyCapId(msg *capnp.Message) (ThirdPartyCapId, error) {
	root, err := msg.RootPtr()
	return ThirdPartyCapId{root.Struct()}, err
}

func (s ThirdPartyCapId) String() string {
	str, _ := text.Marshal(0xb47f4979672cb59d, s.Struct)
	return str
}

func (s ThirdPartyCapId) Copy(seg *capnp.Segment) (ThirdPartyCapId, error) {
	var err error
	t, err := NewThirdPartyCapId(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ThirdPartyCapId_List is a list of ThirdPartyCapId.
type ThirdPartyCapId_List struct{ capnp.List }

// NewThirdPartyCapId creates a new list of ThirdPartyCapId.
func NewThirdPartyCapId_List(s *capnp.Segment, sz int32) (ThirdPartyCapId_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return ThirdPartyCapId_List{l}, err
}

func (s ThirdPartyCapId_List) At(i int) ThirdPartyCapId { return ThirdPartyCapId{s.List.Struct(i)} }

func (s ThirdPartyCapId_List) Set(i int, v ThirdPartyCapId) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s ThirdPartyCapId_List) Copy(seg *capnp.Segment) (ThirdPartyCapId_List, error) {
	var err error
	t, err := NewThirdPartyCapId_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// ThirdPartyCapId_Promise is a wrapper for a ThirdPartyCapId promised by a client call.
type ThirdPartyCapId_Promise struct{ *capnp.Pipeline }

func (p ThirdPartyCapId_Promise) Struct() (ThirdPartyCapId, error) {
	s, err := p.Pipeline.Struct()
	return ThirdPartyCapId{s}, err
}

type JoinKeyPart struct{ capnp.Struct }

// JoinKeyPart_TypeID is the unique identifier for the type JoinKeyPart.
const JoinKeyPart_TypeID = 0x95b29059097fca83

func NewJoinKeyPart(s *capnp.Segment) (JoinKeyPart, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return JoinKeyPart{st}, err
}

func NewRootJoinKeyPart(s *capnp.Segment) (JoinKeyPart, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return JoinKeyPart{st}, err
}

func ReadRootJoinKeyPart(msg *capnp.Message) (JoinKeyPart, error) {
	root, err := msg.RootPtr()
	return JoinKeyPart{root.Struct()}, err
}

func (s JoinKeyPart) String() string {
	str, _ := text.Marshal(0x95b29059097fca83, s.Struct)
	return str
}

func (s JoinKeyPart) Copy(seg *capnp.Segment) (JoinKeyPart, error) {
	var err error
	t, err := NewJoinKeyPart(seg)
	if err != nil {
		return t, err
	}
	t.SetJoinId(s.JoinId())
	t.SetPartCount(s.PartCount())
	t.SetPartNum(s.PartNum())
	return t, nil
}

func (s JoinKeyPart) JoinId() uint32 {
	return s.Struct.Uint32(0)
}

func (s JoinKeyPart) SetJoinId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s JoinKeyPart) PartCount() uint16 {
	return s.Struct.Uint16(4)
}

func (s JoinKeyPart) SetPartCount(v uint16) {
	s.Struct.SetUint16(4, v)
}

func (s JoinKeyPart) PartNum() uint16 {
	return s.Struct.Uint16(6)
}

func (s JoinKeyPart) SetPartNum(v uint16) {
	s.Struct.SetUint16(6, v)
}

// JoinKeyPart_List is a list of JoinKeyPart.
type JoinKeyPart_List struct{ capnp.List }

// NewJoinKeyPart creates a new list of JoinKeyPart.
func NewJoinKeyPart_List(s *capnp.Segment, sz int32) (JoinKeyPart_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return JoinKeyPart_List{l}, err
}

func (s JoinKeyPart_List) At(i int) JoinKeyPart { return JoinKeyPart{s.List.Struct(i)} }

func (s JoinKeyPart_List) Set(i int, v JoinKeyPart) error { return s.List.SetStruct(i, v.Struct) }

func (s JoinKeyPart_List) Copy(seg *capnp.Segment) (JoinKeyPart_List, error) {
	var err error
	t, err := NewJoinKeyPart_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// JoinKeyPart_Promise is a wrapper for a JoinKeyPart promised by a client call.
type JoinKeyPart_Promise struct{ *capnp.Pipeline }

func (p JoinKeyPart_Promise) Struct() (JoinKeyPart, error) {
	s, err := p.Pipeline.Struct()
	return JoinKeyPart{s}, err
}

type JoinResult struct{ capnp.Struct }

// JoinResult_TypeID is the unique identifier for the type JoinResult.
const JoinResult_TypeID = 0x9d263a3630b7ebee

func NewJoinResult(s *capnp.Segment) (JoinResult, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return JoinResult{st}, err
}

func NewRootJoinResult(s *capnp.Segment) (JoinResult, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return JoinResult{st}, err
}

func ReadRootJoinResult(msg *capnp.Message) (JoinResult, error) {
	root, err := msg.RootPtr()
	return JoinResult{root.Struct()}, err
}

func (s JoinResult) String() string {
	str, _ := text.Marshal(0x9d263a3630b7ebee, s.Struct)
	return str
}

func (s JoinResult) Copy(seg *capnp.Segment) (JoinResult, error) {
	var err error
	t, err := NewJoinResult(seg)
	if err != nil {
		return t, err
	}
	t.SetJoinId(s.JoinId())
	t.SetSucceeded(s.Succeeded())
	return t, nil
}

func (s JoinResult) JoinId() uint32 {
	return s.Struct.Uint32(0)
}

func (s JoinResult) SetJoinId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s JoinResult) Succeeded() bool {
	return s.Struct.Bit(32)
}

func (s JoinResult) SetSucceeded(v bool) {
	s.Struct.SetBit(32, v)
}

func (s JoinResult) Cap() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s JoinResult) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JoinResult) CapPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s JoinResult) SetCap(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s JoinResult) SetCapPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// JoinResult_List is a list of JoinResult.
type JoinResult_List struct{ capnp.List }

// NewJoinResult creates a new list of JoinResult.
func NewJoinResult_List(s *capnp.Segment, sz int32) (JoinResult_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return JoinResult_List{l}, err
}

func (s JoinResult_List) At(i int) JoinResult { return JoinResult{s.List.Struct(i)} }

func (s JoinResult_List) Set(i int, v JoinResult) error { return s.List.SetStruct(i, v.Struct) }

func (s JoinResult_List) Copy(seg *capnp.Segment) (JoinResult_List, error) {
	var err error
	t, err := NewJoinResult_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// JoinResult_Promise is a wrapper for a JoinResult promised by a client call.
type JoinResult_Promise struct{ *capnp.Pipeline }

func (p JoinResult_Promise) Struct() (JoinResult, error) {
	s, err := p.Pipeline.Struct()
	return JoinResult{s}, err
}

func (p JoinResult_Promise) Cap() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

const schema_a184c7885cdaf2a1 = "x\xda|\x92\xcfk\x13A\x14\xc7\xdfw&1\x09Z" +
	"\xc2v\"\xa2\x17E\xd0\x83\xa8\xb4(\x0a\x0b\xb2\xc1 " +
	"\xb8Z$\x13[\xb1`\x0f\xcb\xee\xa2+\xed\xeevw" +
	"S\xc9A\x02*h\x0fJ/\x82\x07k\xe9\xd1\x93(" +
	"\xfe*\xe8AA\x84\x1e=x\xf0_\x10zho\x01" +
	"Y\x99\x05\x1b\xc9\x8f\xdef\x1e\x9f\xf7\xe5\xf3\xde\xcc\xd8" +
	"9T\xd9x\xbe\x9d#\x92\xa7\xf3\xbb\xd2/\xf7\xaa\x7f" +
	"\xde/n-\x92&\x90\xaen\xfe\xba\xf1\xe8\xfb\x83U" +
	"\xca\x15\x88\xc4\x11\xd6\x11\xe3\xac@<\xbd\xbf\xde.M" +
	"/\xbdyJR\xa0\x97\x1aa\x1d\xb1\x9f\xa9\xd3^\xf6" +
	"\x8a\x90n\xfc\xfe0vF?\xba\xdc\xc3\xe6\xa1\x90\xd7" +
	"lS|\xce\xe0\xb5\x0c~>\xbf\xf6\xf0\xd3\x8b\x9f+" +
	"\xa4\x09\xd6e\x09b\x8a\x7f\x153\\\x81\xd3\xfc,!" +
	"]~w\xfcf\xcbl\xbf\xed\xd7<5\xc3\x0f@\xcc" +
	"q\xe5\xb9\xef\xd9\xd6\xb7\x97\xa5\xc7\x1f\x07y\x9a\xbc#" +
	"\xa6\xb2D\xc9\x0dB\xfaD\x8f7V\x96v\xff\x18\xc4" +
	"\xce\xf3uq7c[\xdc\xa0Z\x1a\x85\xf6\x89\xe4N" +
	"\x102+JZ'm+\xf4C\xbd\xe1\xda^hx" +
	"\xae\x9f\x98N\x1d\x18\xc8\\\x0a<\xff\xb2\xe1\xb6\xeaV" +
	"\x94\xd4\x01\xb9\x87\xe7\x88r \xd2.\xe8D\xb2\xca!" +
	"'\x184\xb0\x0aT\xd1l\x10\xc9\x8b\x1cr\x92Ac" +
	"\xbc\x02F\xa4\xc9\xf3Dr\x82C^g0n\x07\x9e" +
	"o:(\x12C\x91\x90\x86V\x94\xd4\x82\xa6OHP" +
	" \x86\x02\xa1\xadjW\x9as\xff\xeeC\xbd\x1a\x07\xdd" +
	"\xb89\xbb\xa3\xd6\xa1~-\xa5\x9aY\x1d\x1en\x157" +
	"m\xdbu\x1d\x97\xe0\x00\xc4\x00B\xc1\xb6B\x8c\x12\xc3" +
	"\xe8\x10\xa3\xab\x9e\xe3\x92r)f\xf1\x9aN\x04h%" +
	"\x9d\xc8\x88\xddh\xc1\x8d\x0c{V\xedz\xbb\x99\xff\xd7" +
	"<y\xcb\x8b\x1c\xb5\xe3V\xcd\x0a\xf9\xf0\xe7\xa8G\xc1" +
	"\x82g\xc4^\xe0g\x8c\xccm\xcf=\xa2\xe6.r\xc8" +
	"J\xff8\x83\x92\xaeY\x89\xe9\x10\xf5\x84\x1c\xeb\x86\x94" +
	"c\xcfqQ\xee~p\x02\xca\x84\xbf\x01\x00\x00\xff\xff" +
	"+!\xf3\xa9"

func init() {
	schemas.Register(schema_a184c7885cdaf2a1,
		0x89f389b6fd4082c1,
		0x95b29059097fca83,
		0x9d263a3630b7ebee,
		0x9fd69ebc87b9719c,
		0xb47f4979672cb59d,
		0xb88d09a9c5f39817,
		0xd20b909fee733a8e)
}
