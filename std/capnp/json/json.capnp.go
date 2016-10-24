package json

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type JsonValue struct{ capnp.Struct }
type JsonValue_Which uint16

const (
	JsonValue_Which_null    JsonValue_Which = 0
	JsonValue_Which_boolean JsonValue_Which = 1
	JsonValue_Which_number  JsonValue_Which = 2
	JsonValue_Which_string_ JsonValue_Which = 3
	JsonValue_Which_array   JsonValue_Which = 4
	JsonValue_Which_object  JsonValue_Which = 5
	JsonValue_Which_call    JsonValue_Which = 6
)

func (w JsonValue_Which) String() string {
	const s = "nullbooleannumberstring_arrayobjectcall"
	switch w {
	case JsonValue_Which_null:
		return s[0:4]
	case JsonValue_Which_boolean:
		return s[4:11]
	case JsonValue_Which_number:
		return s[11:17]
	case JsonValue_Which_string_:
		return s[17:24]
	case JsonValue_Which_array:
		return s[24:29]
	case JsonValue_Which_object:
		return s[29:35]
	case JsonValue_Which_call:
		return s[35:39]

	}
	return "JsonValue_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// JsonValue_TypeID is the unique identifier for the type JsonValue.
const JsonValue_TypeID = 0x8825ffaa852cda72

func NewJsonValue(s *capnp.Segment) (JsonValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return JsonValue{st}, err
}

func NewRootJsonValue(s *capnp.Segment) (JsonValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return JsonValue{st}, err
}

func ReadRootJsonValue(msg *capnp.Message) (JsonValue, error) {
	root, err := msg.RootPtr()
	return JsonValue{root.Struct()}, err
}

func (s JsonValue) String() string {
	str, _ := text.Marshal(0x8825ffaa852cda72, s.Struct)
	return str
}

func (s JsonValue) Which() JsonValue_Which {
	return JsonValue_Which(s.Struct.Uint16(0))
}
func (s JsonValue) SetNull() {
	s.Struct.SetUint16(0, 0)

}

func (s JsonValue) Boolean() bool {
	return s.Struct.Bit(16)
}

func (s JsonValue) SetBoolean(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s JsonValue) Number() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s JsonValue) SetNumber(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s JsonValue) String_(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s JsonValue) HasString_() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) String_Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValue) SetString_(v string) error {
	s.Struct.SetUint16(0, 3)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s JsonValue) Array(e *capnp.ErrorSet) JsonValue_List {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return JsonValue_List{}
	}
	return JsonValue_List{List: p.List()}
}

func (s JsonValue) HasArray() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetArray(v JsonValue_List) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewArray sets the array field to a newly
// allocated JsonValue_List, preferring placement in s's segment.
func (s JsonValue) NewArray(n int32) (JsonValue_List, error) {
	s.Struct.SetUint16(0, 4)
	l, err := NewJsonValue_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValue_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s JsonValue) Object(e *capnp.ErrorSet) JsonValueField_List {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return JsonValueField_List{}
	}
	return JsonValueField_List{List: p.List()}
}

func (s JsonValue) HasObject() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetObject(v JsonValueField_List) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewObject sets the object field to a newly
// allocated JsonValueField_List, preferring placement in s's segment.
func (s JsonValue) NewObject(n int32) (JsonValueField_List, error) {
	s.Struct.SetUint16(0, 5)
	l, err := NewJsonValueField_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValueField_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s JsonValue) Call(e *capnp.ErrorSet) JsonValueCall {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return JsonValueCall{}
	}
	return JsonValueCall{Struct: p.Struct()}
}

func (s JsonValue) HasCall() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetCall(v JsonValueCall) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCall sets the call field to a newly
// allocated JsonValueCall struct, preferring placement in s's segment.
func (s JsonValue) NewCall() (JsonValueCall, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewJsonValueCall(s.Struct.Segment())
	if err != nil {
		return JsonValueCall{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// JsonValue_List is a list of JsonValue.
type JsonValue_List struct{ capnp.List }

// NewJsonValue creates a new list of JsonValue.
func NewJsonValue_List(s *capnp.Segment, sz int32) (JsonValue_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return JsonValue_List{l}, err
}

func (s JsonValue_List) At(i int) JsonValue { return JsonValue{s.List.Struct(i)} }

func (s JsonValue_List) Set(i int, v JsonValue) error { return s.List.SetStruct(i, v.Struct) }

// JsonValue_Promise is a wrapper for a JsonValue promised by a client call.
type JsonValue_Promise struct{ *capnp.Pipeline }

func (p JsonValue_Promise) Struct() (JsonValue, error) {
	s, err := p.Pipeline.Struct()
	return JsonValue{s}, err
}

func (p JsonValue_Promise) Call() JsonValueCall_Promise {
	return JsonValueCall_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type JsonValueField struct{ capnp.Struct }

// JsonValueField_TypeID is the unique identifier for the type JsonValueField.
const JsonValueField_TypeID = 0xc27855d853a937cc

func NewJsonValueField(s *capnp.Segment) (JsonValueField, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValueField{st}, err
}

func NewRootJsonValueField(s *capnp.Segment) (JsonValueField, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValueField{st}, err
}

func ReadRootJsonValueField(msg *capnp.Message) (JsonValueField, error) {
	root, err := msg.RootPtr()
	return JsonValueField{root.Struct()}, err
}

func (s JsonValueField) String() string {
	str, _ := text.Marshal(0xc27855d853a937cc, s.Struct)
	return str
}

func (s JsonValueField) Name(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s JsonValueField) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValueField) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValueField) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s JsonValueField) Value(e *capnp.ErrorSet) JsonValue {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return JsonValue{}
	}
	return JsonValue{Struct: p.Struct()}
}

func (s JsonValueField) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s JsonValueField) SetValue(v JsonValue) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated JsonValue struct, preferring placement in s's segment.
func (s JsonValueField) NewValue() (JsonValue, error) {
	ss, err := NewJsonValue(s.Struct.Segment())
	if err != nil {
		return JsonValue{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// JsonValueField_List is a list of JsonValueField.
type JsonValueField_List struct{ capnp.List }

// NewJsonValueField creates a new list of JsonValueField.
func NewJsonValueField_List(s *capnp.Segment, sz int32) (JsonValueField_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return JsonValueField_List{l}, err
}

func (s JsonValueField_List) At(i int) JsonValueField { return JsonValueField{s.List.Struct(i)} }

func (s JsonValueField_List) Set(i int, v JsonValueField) error { return s.List.SetStruct(i, v.Struct) }

// JsonValueField_Promise is a wrapper for a JsonValueField promised by a client call.
type JsonValueField_Promise struct{ *capnp.Pipeline }

func (p JsonValueField_Promise) Struct() (JsonValueField, error) {
	s, err := p.Pipeline.Struct()
	return JsonValueField{s}, err
}

func (p JsonValueField_Promise) Value() JsonValue_Promise {
	return JsonValue_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type JsonValueCall struct{ capnp.Struct }

// JsonValueCall_TypeID is the unique identifier for the type JsonValueCall.
const JsonValueCall_TypeID = 0x9bbf84153dd4bb60

func NewJsonValueCall(s *capnp.Segment) (JsonValueCall, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValueCall{st}, err
}

func NewRootJsonValueCall(s *capnp.Segment) (JsonValueCall, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValueCall{st}, err
}

func ReadRootJsonValueCall(msg *capnp.Message) (JsonValueCall, error) {
	root, err := msg.RootPtr()
	return JsonValueCall{root.Struct()}, err
}

func (s JsonValueCall) String() string {
	str, _ := text.Marshal(0x9bbf84153dd4bb60, s.Struct)
	return str
}

func (s JsonValueCall) Function(e *capnp.ErrorSet) string {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		*e = append(*e, err)
		return ""
	}
	return p.Text()
}

func (s JsonValueCall) HasFunction() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValueCall) FunctionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValueCall) SetFunction(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s JsonValueCall) Params(e *capnp.ErrorSet) JsonValue_List {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return JsonValue_List{}
	}
	return JsonValue_List{List: p.List()}
}

func (s JsonValueCall) HasParams() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s JsonValueCall) SetParams(v JsonValue_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewParams sets the params field to a newly
// allocated JsonValue_List, preferring placement in s's segment.
func (s JsonValueCall) NewParams(n int32) (JsonValue_List, error) {
	l, err := NewJsonValue_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValue_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// JsonValueCall_List is a list of JsonValueCall.
type JsonValueCall_List struct{ capnp.List }

// NewJsonValueCall creates a new list of JsonValueCall.
func NewJsonValueCall_List(s *capnp.Segment, sz int32) (JsonValueCall_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return JsonValueCall_List{l}, err
}

func (s JsonValueCall_List) At(i int) JsonValueCall { return JsonValueCall{s.List.Struct(i)} }

func (s JsonValueCall_List) Set(i int, v JsonValueCall) error { return s.List.SetStruct(i, v.Struct) }

// JsonValueCall_Promise is a wrapper for a JsonValueCall promised by a client call.
type JsonValueCall_Promise struct{ *capnp.Pipeline }

func (p JsonValueCall_Promise) Struct() (JsonValueCall, error) {
	s, err := p.Pipeline.Struct()
	return JsonValueCall{s}, err
}

const schema_8ef99297a43a5e34 = "x\xdat\x92AHTQ\x14\x86\xff\xff\xde7\xa3\xe2" +
	"L\xf3^\xf3\xa4\x16\x89\x9b\xa2\x1235!\x18\x88)" +
	"-\x09\x17\xe1m\xa8e\xf9\x1c_1r}O\xde8" +
	"Y+7\x05\xb5\xa8(\x88\x16\xad\x826\xbajQ\x90" +
	"\x94\x94\xd2\xb2U\xb4\x08W-\xda\x04m\xdaee7" +
	"n\x92O$w\x87s\xce=\xff\xcf\xf7\xdf\x9e\x93<" +
	".z33\x12P\x072Y\x93\xact\xdd\x987\xfb" +
	"nB\xb5R\x98\xfe\x0b\xa5'\x0f\xef\xaf\xde\xc1)6" +
	"5\x01\xc5\xc7\x9c/\xceq?pd\x91w\x09\x9a\xd1" +
	"\x97\x1f\x8e\xb5]\x7f\xfd\x08^\x1b\xd3\xb7\x19a\x97o" +
	"\xc9\xf7\xc5\x07\xd2V\xf7\xe4\x0ch\xde\x1d\x9d\xab|<" +
	"wu\xf9\x7f\xbbtV\x8ay\xc7V-\xce\x0c\xce\x98" +
	"\x89z\x1cuW\x83)FS\xa5\xe1z\x1c\x9d/\x04" +
	"\xba\x11\xaafn>\xd3\xd2\xb7I?\xd3\xd91T\x0b" +
	"\xf5xa0\xd0Z\xed\x91N\xce\x18\x87\x80\xf7\xbc\x13" +
	"PO%\xd5+\xc1v\xfe6\xaeO\xdb^\x18\x00\xd4" +
	"3I\xf5F\xb0]\xac\x19\xfa\x14\x80\xb7X\x02\xd4\x0b" +
	"I\xf5V0/\x7f\x19\x9f\x12\xf0\x96J\xdeR\x87\xfa" +
	"$\xa9\xbe\x0a\xe6\x9d\x9f\xc6\xa7\x03x_\xfa\x00\xf5Y" +
	"\xf2,\x05\xf3\x99\x1f\xc6g\x06\xf0\xd6\xec\x89\xef\x92\x15" +
	"\xdf\xb6\xb3\xab\xc6g\x16(z\xec\x04*9JVv" +
	"S\xb0\x105\xb4Fvv,\x8eu\x18D$\x04\x09" +
	"\x96\xa3\xc6\xe4X\x98\xb0\x15\x82\xad`\xb9>\x9d\xd4\xa2" +
	"\xcb\xca\xa10\xdfn\x1f\xde\xb5sta\x19\xca\x11<" +
	"\xe1\x929\xc0\xe3\xc0\xec\xfa\xcaE\x809\x08\xe6\xc0\x8e" +
	" I\x82k\xdc\x01\x8eH\xd2MY\x83\xb6Y\x8e\xc7" +
	"&\xc2\xeat:\xdf \xba>/T\x03\xad\xe9\xa6l" +
	"A\xba\xe0F&\xe2_&6\x92\xee\xc1@S\x8f\x90" +
	"\xaaY:\xc0_\xe2\x07\x87\xedg\x92T\xfd\x82\x1e\xb9" +
	"\xce\xbb\xd7R\xe9\x92T\xa7\x05\xcd\xa5FT\x9d\xae\xc5" +
	"\x11R\xd3\xe5\xa9 \x09&\xeb\xdb\xba\xdeF~\xa8\x16" +
	"J=\xbeE\xdf&\xbeWR\xf5l\xd2?\xd4\x97\x9a" +
	"*D\xc1d\xb8A\xeb\x8a=\xb4E\xd0\x05\xff\x04\x00" +
	"\x00\xff\xff\x9f\xdb\xc7\x93"

func init() {
	schemas.Register(schema_8ef99297a43a5e34,
		0x8825ffaa852cda72,
		0x9bbf84153dd4bb60,
		0xc27855d853a937cc)
}
