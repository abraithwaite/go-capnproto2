package aircraftlib

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

// Constants defined in aircraft.capnp.
const (
	ConstEnum = Airport_jfk
)

// Constants defined in aircraft.capnp.
var (
	ConstDate = Zdate{Struct: capnp.MustUnmarshalRootPtr(x_832bcc6686a26d56[0:24]).Struct()}
	ConstList = Zdate_List{List: capnp.MustUnmarshalRootPtr(x_832bcc6686a26d56[24:64]).List()}
)

type Zdate struct{ capnp.Struct }

// Zdate_TypeID is the unique identifier for the type Zdate.
const Zdate_TypeID = 0xde50aebbad57549d

func NewZdate(s *capnp.Segment) (Zdate, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Zdate{st}, err
}

func NewRootZdate(s *capnp.Segment) (Zdate, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Zdate{st}, err
}

func ReadRootZdate(msg *capnp.Message) (Zdate, error) {
	root, err := msg.RootPtr()
	return Zdate{root.Struct()}, err
}

func (s Zdate) String() string {
	str, _ := text.Marshal(0xde50aebbad57549d, s.Struct)
	return str
}

func (s Zdate) Copy(seg *capnp.Segment) (Zdate, error) {
	var err error
	t, err := NewZdate(seg)
	if err != nil {
		return t, err
	}
	t.SetYear(s.Year())
	t.SetMonth(s.Month())
	t.SetDay(s.Day())
	return t, nil
}

func (s Zdate) Year() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s Zdate) SetYear(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Zdate) Month() uint8 {
	return s.Struct.Uint8(2)
}

func (s Zdate) SetMonth(v uint8) {
	s.Struct.SetUint8(2, v)
}

func (s Zdate) Day() uint8 {
	return s.Struct.Uint8(3)
}

func (s Zdate) SetDay(v uint8) {
	s.Struct.SetUint8(3, v)
}

// Zdate_List is a list of Zdate.
type Zdate_List struct{ capnp.List }

// NewZdate creates a new list of Zdate.
func NewZdate_List(s *capnp.Segment, sz int32) (Zdate_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Zdate_List{l}, err
}

func (s Zdate_List) At(i int) Zdate { return Zdate{s.List.Struct(i)} }

func (s Zdate_List) Set(i int, v Zdate) error { return s.List.SetStruct(i, v.Struct) }

func (s Zdate_List) Copy(seg *capnp.Segment) (Zdate_List, error) {
	var err error
	t, err := NewZdate_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Zdate_Promise is a wrapper for a Zdate promised by a client call.
type Zdate_Promise struct{ *capnp.Pipeline }

func (p Zdate_Promise) Struct() (Zdate, error) {
	s, err := p.Pipeline.Struct()
	return Zdate{s}, err
}

type Zdata struct{ capnp.Struct }

// Zdata_TypeID is the unique identifier for the type Zdata.
const Zdata_TypeID = 0xc7da65f9a2f20ba2

func NewZdata(s *capnp.Segment) (Zdata, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zdata{st}, err
}

func NewRootZdata(s *capnp.Segment) (Zdata, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zdata{st}, err
}

func ReadRootZdata(msg *capnp.Message) (Zdata, error) {
	root, err := msg.RootPtr()
	return Zdata{root.Struct()}, err
}

func (s Zdata) String() string {
	str, _ := text.Marshal(0xc7da65f9a2f20ba2, s.Struct)
	return str
}

func (s Zdata) Copy(seg *capnp.Segment) (Zdata, error) {
	var err error
	t, err := NewZdata(seg)
	if err != nil {
		return t, err
	}
	{
		v, err := s.Data()
		if err != nil {
			return t, err
		}
		t.SetData(v)
	}
	return t, nil
}

func (s Zdata) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Zdata) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zdata) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

// Zdata_List is a list of Zdata.
type Zdata_List struct{ capnp.List }

// NewZdata creates a new list of Zdata.
func NewZdata_List(s *capnp.Segment, sz int32) (Zdata_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Zdata_List{l}, err
}

func (s Zdata_List) At(i int) Zdata { return Zdata{s.List.Struct(i)} }

func (s Zdata_List) Set(i int, v Zdata) error { return s.List.SetStruct(i, v.Struct) }

func (s Zdata_List) Copy(seg *capnp.Segment) (Zdata_List, error) {
	var err error
	t, err := NewZdata_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Zdata_Promise is a wrapper for a Zdata promised by a client call.
type Zdata_Promise struct{ *capnp.Pipeline }

func (p Zdata_Promise) Struct() (Zdata, error) {
	s, err := p.Pipeline.Struct()
	return Zdata{s}, err
}

type Airport uint16

// Values of Airport.
const (
	Airport_none Airport = 0
	Airport_jfk  Airport = 1
	Airport_lax  Airport = 2
	Airport_sfo  Airport = 3
	Airport_luv  Airport = 4
	Airport_dfw  Airport = 5
	Airport_test Airport = 6
)

// String returns the enum's constant name.
func (c Airport) String() string {
	switch c {
	case Airport_none:
		return "none"
	case Airport_jfk:
		return "jfk"
	case Airport_lax:
		return "lax"
	case Airport_sfo:
		return "sfo"
	case Airport_luv:
		return "luv"
	case Airport_dfw:
		return "dfw"
	case Airport_test:
		return "test"

	default:
		return ""
	}
}

// AirportFromString returns the enum value with a name,
// or the zero value if there's no such value.
func AirportFromString(c string) Airport {
	switch c {
	case "none":
		return Airport_none
	case "jfk":
		return Airport_jfk
	case "lax":
		return Airport_lax
	case "sfo":
		return Airport_sfo
	case "luv":
		return Airport_luv
	case "dfw":
		return Airport_dfw
	case "test":
		return Airport_test

	default:
		return 0
	}
}

type Airport_List struct{ capnp.List }

func NewAirport_List(s *capnp.Segment, sz int32) (Airport_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Airport_List{l.List}, err
}

func (l Airport_List) At(i int) Airport {
	ul := capnp.UInt16List{List: l.List}
	return Airport(ul.At(i))
}

func (l Airport_List) Set(i int, v Airport) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

func (l Airport_List) Copy(seg *capnp.Segment) (Airport_List, error) {
	ul := capnp.UInt16List{List: l.List}
	t, err := ul.Copy(seg)
	return Airport_List{t.List}, err
}

type PlaneBase struct{ capnp.Struct }

// PlaneBase_TypeID is the unique identifier for the type PlaneBase.
const PlaneBase_TypeID = 0xd8bccf6e60a73791

func NewPlaneBase(s *capnp.Segment) (PlaneBase, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return PlaneBase{st}, err
}

func NewRootPlaneBase(s *capnp.Segment) (PlaneBase, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2})
	return PlaneBase{st}, err
}

func ReadRootPlaneBase(msg *capnp.Message) (PlaneBase, error) {
	root, err := msg.RootPtr()
	return PlaneBase{root.Struct()}, err
}

func (s PlaneBase) String() string {
	str, _ := text.Marshal(0xd8bccf6e60a73791, s.Struct)
	return str
}

func (s PlaneBase) Copy(seg *capnp.Segment) (PlaneBase, error) {
	var err error
	t, err := NewPlaneBase(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Homes()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetHomes(v)
	}
	{
		v, err := s.Name()
		if err != nil {
			return t, err
		}
		t.SetName(v)
	}
	t.SetRating(s.Rating())
	t.SetCanFly(s.CanFly())
	t.SetCapacity(s.Capacity())
	t.SetMaxSpeed(s.MaxSpeed())
	return t, nil
}

func (s PlaneBase) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s PlaneBase) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PlaneBase) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s PlaneBase) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s PlaneBase) Homes() (Airport_List, error) {
	p, err := s.Struct.Ptr(1)
	return Airport_List{List: p.List()}, err
}

func (s PlaneBase) HasHomes() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PlaneBase) SetHomes(v Airport_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewHomes sets the homes field to a newly
// allocated Airport_List, preferring placement in s's segment.
func (s PlaneBase) NewHomes(n int32) (Airport_List, error) {
	l, err := NewAirport_List(s.Struct.Segment(), n)
	if err != nil {
		return Airport_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s PlaneBase) Rating() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s PlaneBase) SetRating(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s PlaneBase) CanFly() bool {
	return s.Struct.Bit(64)
}

func (s PlaneBase) SetCanFly(v bool) {
	s.Struct.SetBit(64, v)
}

func (s PlaneBase) Capacity() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s PlaneBase) SetCapacity(v int64) {
	s.Struct.SetUint64(16, uint64(v))
}

func (s PlaneBase) MaxSpeed() float64 {
	return math.Float64frombits(s.Struct.Uint64(24))
}

func (s PlaneBase) SetMaxSpeed(v float64) {
	s.Struct.SetUint64(24, math.Float64bits(v))
}

// PlaneBase_List is a list of PlaneBase.
type PlaneBase_List struct{ capnp.List }

// NewPlaneBase creates a new list of PlaneBase.
func NewPlaneBase_List(s *capnp.Segment, sz int32) (PlaneBase_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 2}, sz)
	return PlaneBase_List{l}, err
}

func (s PlaneBase_List) At(i int) PlaneBase { return PlaneBase{s.List.Struct(i)} }

func (s PlaneBase_List) Set(i int, v PlaneBase) error { return s.List.SetStruct(i, v.Struct) }

func (s PlaneBase_List) Copy(seg *capnp.Segment) (PlaneBase_List, error) {
	var err error
	t, err := NewPlaneBase_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// PlaneBase_Promise is a wrapper for a PlaneBase promised by a client call.
type PlaneBase_Promise struct{ *capnp.Pipeline }

func (p PlaneBase_Promise) Struct() (PlaneBase, error) {
	s, err := p.Pipeline.Struct()
	return PlaneBase{s}, err
}

type B737 struct{ capnp.Struct }

// B737_TypeID is the unique identifier for the type B737.
const B737_TypeID = 0xccb3b2e3603826e0

func NewB737(s *capnp.Segment) (B737, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return B737{st}, err
}

func NewRootB737(s *capnp.Segment) (B737, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return B737{st}, err
}

func ReadRootB737(msg *capnp.Message) (B737, error) {
	root, err := msg.RootPtr()
	return B737{root.Struct()}, err
}

func (s B737) String() string {
	str, _ := text.Marshal(0xccb3b2e3603826e0, s.Struct)
	return str
}

func (s B737) Copy(seg *capnp.Segment) (B737, error) {
	var err error
	t, err := NewB737(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Base()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBase(v)
	}
	return t, nil
}

func (s B737) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s B737) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s B737) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s B737) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// B737_List is a list of B737.
type B737_List struct{ capnp.List }

// NewB737 creates a new list of B737.
func NewB737_List(s *capnp.Segment, sz int32) (B737_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return B737_List{l}, err
}

func (s B737_List) At(i int) B737 { return B737{s.List.Struct(i)} }

func (s B737_List) Set(i int, v B737) error { return s.List.SetStruct(i, v.Struct) }

func (s B737_List) Copy(seg *capnp.Segment) (B737_List, error) {
	var err error
	t, err := NewB737_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// B737_Promise is a wrapper for a B737 promised by a client call.
type B737_Promise struct{ *capnp.Pipeline }

func (p B737_Promise) Struct() (B737, error) {
	s, err := p.Pipeline.Struct()
	return B737{s}, err
}

func (p B737_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type A320 struct{ capnp.Struct }

// A320_TypeID is the unique identifier for the type A320.
const A320_TypeID = 0xd98c608877d9cb8d

func NewA320(s *capnp.Segment) (A320, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return A320{st}, err
}

func NewRootA320(s *capnp.Segment) (A320, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return A320{st}, err
}

func ReadRootA320(msg *capnp.Message) (A320, error) {
	root, err := msg.RootPtr()
	return A320{root.Struct()}, err
}

func (s A320) String() string {
	str, _ := text.Marshal(0xd98c608877d9cb8d, s.Struct)
	return str
}

func (s A320) Copy(seg *capnp.Segment) (A320, error) {
	var err error
	t, err := NewA320(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Base()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBase(v)
	}
	return t, nil
}

func (s A320) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s A320) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s A320) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s A320) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// A320_List is a list of A320.
type A320_List struct{ capnp.List }

// NewA320 creates a new list of A320.
func NewA320_List(s *capnp.Segment, sz int32) (A320_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return A320_List{l}, err
}

func (s A320_List) At(i int) A320 { return A320{s.List.Struct(i)} }

func (s A320_List) Set(i int, v A320) error { return s.List.SetStruct(i, v.Struct) }

func (s A320_List) Copy(seg *capnp.Segment) (A320_List, error) {
	var err error
	t, err := NewA320_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// A320_Promise is a wrapper for a A320 promised by a client call.
type A320_Promise struct{ *capnp.Pipeline }

func (p A320_Promise) Struct() (A320, error) {
	s, err := p.Pipeline.Struct()
	return A320{s}, err
}

func (p A320_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type F16 struct{ capnp.Struct }

// F16_TypeID is the unique identifier for the type F16.
const F16_TypeID = 0xe1c9eac512335361

func NewF16(s *capnp.Segment) (F16, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return F16{st}, err
}

func NewRootF16(s *capnp.Segment) (F16, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return F16{st}, err
}

func ReadRootF16(msg *capnp.Message) (F16, error) {
	root, err := msg.RootPtr()
	return F16{root.Struct()}, err
}

func (s F16) String() string {
	str, _ := text.Marshal(0xe1c9eac512335361, s.Struct)
	return str
}

func (s F16) Copy(seg *capnp.Segment) (F16, error) {
	var err error
	t, err := NewF16(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Base()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBase(v)
	}
	return t, nil
}

func (s F16) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s F16) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s F16) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s F16) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// F16_List is a list of F16.
type F16_List struct{ capnp.List }

// NewF16 creates a new list of F16.
func NewF16_List(s *capnp.Segment, sz int32) (F16_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return F16_List{l}, err
}

func (s F16_List) At(i int) F16 { return F16{s.List.Struct(i)} }

func (s F16_List) Set(i int, v F16) error { return s.List.SetStruct(i, v.Struct) }

func (s F16_List) Copy(seg *capnp.Segment) (F16_List, error) {
	var err error
	t, err := NewF16_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// F16_Promise is a wrapper for a F16 promised by a client call.
type F16_Promise struct{ *capnp.Pipeline }

func (p F16_Promise) Struct() (F16, error) {
	s, err := p.Pipeline.Struct()
	return F16{s}, err
}

func (p F16_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Regression struct{ capnp.Struct }

// Regression_TypeID is the unique identifier for the type Regression.
const Regression_TypeID = 0xb1f0385d845e367f

func NewRegression(s *capnp.Segment) (Regression, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Regression{st}, err
}

func NewRootRegression(s *capnp.Segment) (Regression, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Regression{st}, err
}

func ReadRootRegression(msg *capnp.Message) (Regression, error) {
	root, err := msg.RootPtr()
	return Regression{root.Struct()}, err
}

func (s Regression) String() string {
	str, _ := text.Marshal(0xb1f0385d845e367f, s.Struct)
	return str
}

func (s Regression) Copy(seg *capnp.Segment) (Regression, error) {
	var err error
	t, err := NewRegression(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Base()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBase(v)
	}
	{
		o, err := s.Beta()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBeta(v)
	}
	{
		o, err := s.Planes()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPlanes(v)
	}
	t.SetB0(s.B0())
	t.SetYmu(s.Ymu())
	t.SetYsd(s.Ysd())
	return t, nil
}

func (s Regression) Base() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s Regression) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Regression) SetBase(v PlaneBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Regression) NewBase() (PlaneBase, error) {
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Regression) B0() float64 {
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s Regression) SetB0(v float64) {
	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s Regression) Beta() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Float64List{List: p.List()}, err
}

func (s Regression) HasBeta() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Regression) SetBeta(v capnp.Float64List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewBeta sets the beta field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s Regression) NewBeta(n int32) (capnp.Float64List, error) {
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s Regression) Planes() (Aircraft_List, error) {
	p, err := s.Struct.Ptr(2)
	return Aircraft_List{List: p.List()}, err
}

func (s Regression) HasPlanes() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Regression) SetPlanes(v Aircraft_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewPlanes sets the planes field to a newly
// allocated Aircraft_List, preferring placement in s's segment.
func (s Regression) NewPlanes(n int32) (Aircraft_List, error) {
	l, err := NewAircraft_List(s.Struct.Segment(), n)
	if err != nil {
		return Aircraft_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Regression) Ymu() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Regression) SetYmu(v float64) {
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Regression) Ysd() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s Regression) SetYsd(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// Regression_List is a list of Regression.
type Regression_List struct{ capnp.List }

// NewRegression creates a new list of Regression.
func NewRegression_List(s *capnp.Segment, sz int32) (Regression_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return Regression_List{l}, err
}

func (s Regression_List) At(i int) Regression { return Regression{s.List.Struct(i)} }

func (s Regression_List) Set(i int, v Regression) error { return s.List.SetStruct(i, v.Struct) }

func (s Regression_List) Copy(seg *capnp.Segment) (Regression_List, error) {
	var err error
	t, err := NewRegression_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Regression_Promise is a wrapper for a Regression promised by a client call.
type Regression_Promise struct{ *capnp.Pipeline }

func (p Regression_Promise) Struct() (Regression, error) {
	s, err := p.Pipeline.Struct()
	return Regression{s}, err
}

func (p Regression_Promise) Base() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Aircraft struct{ capnp.Struct }
type Aircraft_Which uint16

const (
	Aircraft_Which_void Aircraft_Which = 0
	Aircraft_Which_b737 Aircraft_Which = 1
	Aircraft_Which_a320 Aircraft_Which = 2
	Aircraft_Which_f16  Aircraft_Which = 3
)

func (w Aircraft_Which) String() string {
	const s = "voidb737a320f16"
	switch w {
	case Aircraft_Which_void:
		return s[0:4]
	case Aircraft_Which_b737:
		return s[4:8]
	case Aircraft_Which_a320:
		return s[8:12]
	case Aircraft_Which_f16:
		return s[12:15]

	}
	return "Aircraft_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Aircraft_TypeID is the unique identifier for the type Aircraft.
const Aircraft_TypeID = 0xe54e10aede55c7b1

func NewAircraft(s *capnp.Segment) (Aircraft, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Aircraft{st}, err
}

func NewRootAircraft(s *capnp.Segment) (Aircraft, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Aircraft{st}, err
}

func ReadRootAircraft(msg *capnp.Message) (Aircraft, error) {
	root, err := msg.RootPtr()
	return Aircraft{root.Struct()}, err
}

func (s Aircraft) String() string {
	str, _ := text.Marshal(0xe54e10aede55c7b1, s.Struct)
	return str
}

func (s Aircraft) Copy(seg *capnp.Segment) (Aircraft, error) {
	var err error
	t, err := NewAircraft(seg)
	if err != nil {
		return t, err
	}
	if s.HasB737() {
		v, err := s.B737()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetB737(c)
	}
	if s.HasA320() {
		v, err := s.A320()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetA320(c)
	}
	if s.HasF16() {
		v, err := s.F16()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetF16(c)
	}
	return t, nil
}

func (s Aircraft) Which() Aircraft_Which {
	return Aircraft_Which(s.Struct.Uint16(0))
}
func (s Aircraft) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Aircraft) B737() (B737, error) {
	p, err := s.Struct.Ptr(0)
	return B737{Struct: p.Struct()}, err
}

func (s Aircraft) HasB737() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetB737(v B737) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Aircraft) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Aircraft) A320() (A320, error) {
	p, err := s.Struct.Ptr(0)
	return A320{Struct: p.Struct()}, err
}

func (s Aircraft) HasA320() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetA320(v A320) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Aircraft) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Aircraft) F16() (F16, error) {
	p, err := s.Struct.Ptr(0)
	return F16{Struct: p.Struct()}, err
}

func (s Aircraft) HasF16() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Aircraft) SetF16(v F16) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Aircraft) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Aircraft_List is a list of Aircraft.
type Aircraft_List struct{ capnp.List }

// NewAircraft creates a new list of Aircraft.
func NewAircraft_List(s *capnp.Segment, sz int32) (Aircraft_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Aircraft_List{l}, err
}

func (s Aircraft_List) At(i int) Aircraft { return Aircraft{s.List.Struct(i)} }

func (s Aircraft_List) Set(i int, v Aircraft) error { return s.List.SetStruct(i, v.Struct) }

func (s Aircraft_List) Copy(seg *capnp.Segment) (Aircraft_List, error) {
	var err error
	t, err := NewAircraft_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Aircraft_Promise is a wrapper for a Aircraft promised by a client call.
type Aircraft_Promise struct{ *capnp.Pipeline }

func (p Aircraft_Promise) Struct() (Aircraft, error) {
	s, err := p.Pipeline.Struct()
	return Aircraft{s}, err
}

func (p Aircraft_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Aircraft_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Z struct{ capnp.Struct }
type Z_grp Z
type Z_Which uint16

const (
	Z_Which_void        Z_Which = 0
	Z_Which_zz          Z_Which = 1
	Z_Which_f64         Z_Which = 2
	Z_Which_f32         Z_Which = 3
	Z_Which_i64         Z_Which = 4
	Z_Which_i32         Z_Which = 5
	Z_Which_i16         Z_Which = 6
	Z_Which_i8          Z_Which = 7
	Z_Which_u64         Z_Which = 8
	Z_Which_u32         Z_Which = 9
	Z_Which_u16         Z_Which = 10
	Z_Which_u8          Z_Which = 11
	Z_Which_bool        Z_Which = 12
	Z_Which_text        Z_Which = 13
	Z_Which_blob        Z_Which = 14
	Z_Which_f64vec      Z_Which = 15
	Z_Which_f32vec      Z_Which = 16
	Z_Which_i64vec      Z_Which = 17
	Z_Which_i32vec      Z_Which = 18
	Z_Which_i16vec      Z_Which = 19
	Z_Which_i8vec       Z_Which = 20
	Z_Which_u64vec      Z_Which = 21
	Z_Which_u32vec      Z_Which = 22
	Z_Which_u16vec      Z_Which = 23
	Z_Which_u8vec       Z_Which = 24
	Z_Which_boolvec     Z_Which = 39
	Z_Which_datavec     Z_Which = 40
	Z_Which_textvec     Z_Which = 41
	Z_Which_zvec        Z_Which = 25
	Z_Which_zvecvec     Z_Which = 26
	Z_Which_zdate       Z_Which = 27
	Z_Which_zdata       Z_Which = 28
	Z_Which_aircraftvec Z_Which = 29
	Z_Which_aircraft    Z_Which = 30
	Z_Which_regression  Z_Which = 31
	Z_Which_planebase   Z_Which = 32
	Z_Which_airport     Z_Which = 33
	Z_Which_b737        Z_Which = 34
	Z_Which_a320        Z_Which = 35
	Z_Which_f16         Z_Which = 36
	Z_Which_zdatevec    Z_Which = 37
	Z_Which_zdatavec    Z_Which = 38
	Z_Which_grp         Z_Which = 42
)

func (w Z_Which) String() string {
	const s = "voidzzf64f32i64i32i16i8u64u32u16u8booltextblobf64vecf32veci64veci32veci16veci8vecu64vecu32vecu16vecu8vecboolvecdatavectextveczveczvecveczdatezdataaircraftvecaircraftregressionplanebaseairportb737a320f16zdateveczdatavecgrp"
	switch w {
	case Z_Which_void:
		return s[0:4]
	case Z_Which_zz:
		return s[4:6]
	case Z_Which_f64:
		return s[6:9]
	case Z_Which_f32:
		return s[9:12]
	case Z_Which_i64:
		return s[12:15]
	case Z_Which_i32:
		return s[15:18]
	case Z_Which_i16:
		return s[18:21]
	case Z_Which_i8:
		return s[21:23]
	case Z_Which_u64:
		return s[23:26]
	case Z_Which_u32:
		return s[26:29]
	case Z_Which_u16:
		return s[29:32]
	case Z_Which_u8:
		return s[32:34]
	case Z_Which_bool:
		return s[34:38]
	case Z_Which_text:
		return s[38:42]
	case Z_Which_blob:
		return s[42:46]
	case Z_Which_f64vec:
		return s[46:52]
	case Z_Which_f32vec:
		return s[52:58]
	case Z_Which_i64vec:
		return s[58:64]
	case Z_Which_i32vec:
		return s[64:70]
	case Z_Which_i16vec:
		return s[70:76]
	case Z_Which_i8vec:
		return s[76:81]
	case Z_Which_u64vec:
		return s[81:87]
	case Z_Which_u32vec:
		return s[87:93]
	case Z_Which_u16vec:
		return s[93:99]
	case Z_Which_u8vec:
		return s[99:104]
	case Z_Which_boolvec:
		return s[104:111]
	case Z_Which_datavec:
		return s[111:118]
	case Z_Which_textvec:
		return s[118:125]
	case Z_Which_zvec:
		return s[125:129]
	case Z_Which_zvecvec:
		return s[129:136]
	case Z_Which_zdate:
		return s[136:141]
	case Z_Which_zdata:
		return s[141:146]
	case Z_Which_aircraftvec:
		return s[146:157]
	case Z_Which_aircraft:
		return s[157:165]
	case Z_Which_regression:
		return s[165:175]
	case Z_Which_planebase:
		return s[175:184]
	case Z_Which_airport:
		return s[184:191]
	case Z_Which_b737:
		return s[191:195]
	case Z_Which_a320:
		return s[195:199]
	case Z_Which_f16:
		return s[199:202]
	case Z_Which_zdatevec:
		return s[202:210]
	case Z_Which_zdatavec:
		return s[210:218]
	case Z_Which_grp:
		return s[218:221]

	}
	return "Z_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Z_TypeID is the unique identifier for the type Z.
const Z_TypeID = 0xea26e9973bd6a0d9

func NewZ(s *capnp.Segment) (Z, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Z{st}, err
}

func NewRootZ(s *capnp.Segment) (Z, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Z{st}, err
}

func ReadRootZ(msg *capnp.Message) (Z, error) {
	root, err := msg.RootPtr()
	return Z{root.Struct()}, err
}

func (s Z) String() string {
	str, _ := text.Marshal(0xea26e9973bd6a0d9, s.Struct)
	return str
}

func (s Z) Copy(seg *capnp.Segment) (Z, error) {
	var err error
	t, err := NewZ(seg)
	if err != nil {
		return t, err
	}
	if s.HasZz() {
		v, err := s.Zz()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZz(c)
	}
	if s.HasF64vec() {
		v, err := s.F64vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetF64vec(c)
	}
	if s.HasF32vec() {
		v, err := s.F32vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetF32vec(c)
	}
	if s.HasI64vec() {
		v, err := s.I64vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetI64vec(c)
	}
	if s.HasI32vec() {
		v, err := s.I32vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetI32vec(c)
	}
	if s.HasI16vec() {
		v, err := s.I16vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetI16vec(c)
	}
	if s.HasI8vec() {
		v, err := s.I8vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetI8vec(c)
	}
	if s.HasU64vec() {
		v, err := s.U64vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetU64vec(c)
	}
	if s.HasU32vec() {
		v, err := s.U32vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetU32vec(c)
	}
	if s.HasU16vec() {
		v, err := s.U16vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetU16vec(c)
	}
	if s.HasU8vec() {
		v, err := s.U8vec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetU8vec(c)
	}
	if s.HasBoolvec() {
		v, err := s.Boolvec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBoolvec(c)
	}
	if s.HasDatavec() {
		v, err := s.Datavec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetDatavec(c)
	}
	if s.HasTextvec() {
		v, err := s.Textvec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetTextvec(c)
	}
	if s.HasZvec() {
		v, err := s.Zvec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZvec(c)
	}
	if s.HasZvecvec() {
		v, err := s.Zvecvec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZvecvec(c)
	}
	if s.HasZdate() {
		v, err := s.Zdate()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZdate(c)
	}
	if s.HasZdata() {
		v, err := s.Zdata()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZdata(c)
	}
	if s.HasAircraftvec() {
		v, err := s.Aircraftvec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetAircraftvec(c)
	}
	if s.HasAircraft() {
		v, err := s.Aircraft()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetAircraft(c)
	}
	if s.HasRegression() {
		v, err := s.Regression()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetRegression(c)
	}
	if s.HasPlanebase() {
		v, err := s.Planebase()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPlanebase(c)
	}
	if s.HasB737() {
		v, err := s.B737()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetB737(c)
	}
	if s.HasA320() {
		v, err := s.A320()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetA320(c)
	}
	if s.HasF16() {
		v, err := s.F16()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetF16(c)
	}
	if s.HasZdatevec() {
		v, err := s.Zdatevec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZdatevec(c)
	}
	if s.HasZdatavec() {
		v, err := s.Zdatavec()
		if err != nil {
			return t, err
		}
		c, err := v.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetZdatavec(c)
	}
	if s.Which() == Z_Which_text {
		v, err := s.Text()
		if err != nil {
			return t, err
		}
		t.SetText(v)
	}
	if s.Which() == Z_Which_blob {
		v, err := s.Blob()
		if err != nil {
			return t, err
		}
		t.SetBlob(v)
	}
	if s.Which() == Z_Which_f64 {
		t.SetF64(s.F64())
	}
	if s.Which() == Z_Which_f32 {
		t.SetF32(s.F32())
	}
	if s.Which() == Z_Which_i64 {
		t.SetI64(s.I64())
	}
	if s.Which() == Z_Which_i32 {
		t.SetI32(s.I32())
	}
	if s.Which() == Z_Which_i16 {
		t.SetI16(s.I16())
	}
	if s.Which() == Z_Which_i8 {
		t.SetI8(s.I8())
	}
	if s.Which() == Z_Which_u64 {
		t.SetU64(s.U64())
	}
	if s.Which() == Z_Which_u32 {
		t.SetU32(s.U32())
	}
	if s.Which() == Z_Which_u16 {
		t.SetU16(s.U16())
	}
	if s.Which() == Z_Which_u8 {
		t.SetU8(s.U8())
	}
	if s.Which() == Z_Which_bool {
		t.SetBool(s.Bool())
	}
	if s.Which() == Z_Which_airport {
		t.SetAirport(s.Airport())
	}
	return t, nil
}

func (s Z) Which() Z_Which {
	return Z_Which(s.Struct.Uint16(0))
}
func (s Z) SetVoid() {
	s.Struct.SetUint16(0, 0)

}

func (s Z) Zz() (Z, error) {
	p, err := s.Struct.Ptr(0)
	return Z{Struct: p.Struct()}, err
}

func (s Z) HasZz() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZz(v Z) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZz sets the zz field to a newly
// allocated Z struct, preferring placement in s's segment.
func (s Z) NewZz() (Z, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewZ(s.Struct.Segment())
	if err != nil {
		return Z{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) F64() float64 {
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s Z) SetF64(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s Z) F32() float32 {
	return math.Float32frombits(s.Struct.Uint32(8))
}

func (s Z) SetF32(v float32) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint32(8, math.Float32bits(v))
}

func (s Z) I64() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Z) SetI64(v int64) {
	s.Struct.SetUint16(0, 4)
	s.Struct.SetUint64(8, uint64(v))
}

func (s Z) I32() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s Z) SetI32(v int32) {
	s.Struct.SetUint16(0, 5)
	s.Struct.SetUint32(8, uint32(v))
}

func (s Z) I16() int16 {
	return int16(s.Struct.Uint16(8))
}

func (s Z) SetI16(v int16) {
	s.Struct.SetUint16(0, 6)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) I8() int8 {
	return int8(s.Struct.Uint8(8))
}

func (s Z) SetI8(v int8) {
	s.Struct.SetUint16(0, 7)
	s.Struct.SetUint8(8, uint8(v))
}

func (s Z) U64() uint64 {
	return s.Struct.Uint64(8)
}

func (s Z) SetU64(v uint64) {
	s.Struct.SetUint16(0, 8)
	s.Struct.SetUint64(8, v)
}

func (s Z) U32() uint32 {
	return s.Struct.Uint32(8)
}

func (s Z) SetU32(v uint32) {
	s.Struct.SetUint16(0, 9)
	s.Struct.SetUint32(8, v)
}

func (s Z) U16() uint16 {
	return s.Struct.Uint16(8)
}

func (s Z) SetU16(v uint16) {
	s.Struct.SetUint16(0, 10)
	s.Struct.SetUint16(8, v)
}

func (s Z) U8() uint8 {
	return s.Struct.Uint8(8)
}

func (s Z) SetU8(v uint8) {
	s.Struct.SetUint16(0, 11)
	s.Struct.SetUint8(8, v)
}

func (s Z) Bool() bool {
	return s.Struct.Bit(64)
}

func (s Z) SetBool(v bool) {
	s.Struct.SetUint16(0, 12)
	s.Struct.SetBit(64, v)
}

func (s Z) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Z) HasText() bool {
	if s.Struct.Uint16(0) != 13 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Z) SetText(v string) error {
	s.Struct.SetUint16(0, 13)
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Z) Blob() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Z) HasBlob() bool {
	if s.Struct.Uint16(0) != 14 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetBlob(v []byte) error {
	s.Struct.SetUint16(0, 14)
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Z) F64vec() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Float64List{List: p.List()}, err
}

func (s Z) HasF64vec() bool {
	if s.Struct.Uint16(0) != 15 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF64vec(v capnp.Float64List) error {
	s.Struct.SetUint16(0, 15)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewF64vec sets the f64vec field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s Z) NewF64vec(n int32) (capnp.Float64List, error) {
	s.Struct.SetUint16(0, 15)
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) F32vec() (capnp.Float32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Float32List{List: p.List()}, err
}

func (s Z) HasF32vec() bool {
	if s.Struct.Uint16(0) != 16 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF32vec(v capnp.Float32List) error {
	s.Struct.SetUint16(0, 16)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewF32vec sets the f32vec field to a newly
// allocated capnp.Float32List, preferring placement in s's segment.
func (s Z) NewF32vec(n int32) (capnp.Float32List, error) {
	s.Struct.SetUint16(0, 16)
	l, err := capnp.NewFloat32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I64vec() (capnp.Int64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int64List{List: p.List()}, err
}

func (s Z) HasI64vec() bool {
	if s.Struct.Uint16(0) != 17 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI64vec(v capnp.Int64List) error {
	s.Struct.SetUint16(0, 17)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI64vec sets the i64vec field to a newly
// allocated capnp.Int64List, preferring placement in s's segment.
func (s Z) NewI64vec(n int32) (capnp.Int64List, error) {
	s.Struct.SetUint16(0, 17)
	l, err := capnp.NewInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I32vec() (capnp.Int32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int32List{List: p.List()}, err
}

func (s Z) HasI32vec() bool {
	if s.Struct.Uint16(0) != 18 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI32vec(v capnp.Int32List) error {
	s.Struct.SetUint16(0, 18)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI32vec sets the i32vec field to a newly
// allocated capnp.Int32List, preferring placement in s's segment.
func (s Z) NewI32vec(n int32) (capnp.Int32List, error) {
	s.Struct.SetUint16(0, 18)
	l, err := capnp.NewInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I16vec() (capnp.Int16List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int16List{List: p.List()}, err
}

func (s Z) HasI16vec() bool {
	if s.Struct.Uint16(0) != 19 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI16vec(v capnp.Int16List) error {
	s.Struct.SetUint16(0, 19)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI16vec sets the i16vec field to a newly
// allocated capnp.Int16List, preferring placement in s's segment.
func (s Z) NewI16vec(n int32) (capnp.Int16List, error) {
	s.Struct.SetUint16(0, 19)
	l, err := capnp.NewInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) I8vec() (capnp.Int8List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.Int8List{List: p.List()}, err
}

func (s Z) HasI8vec() bool {
	if s.Struct.Uint16(0) != 20 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetI8vec(v capnp.Int8List) error {
	s.Struct.SetUint16(0, 20)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewI8vec sets the i8vec field to a newly
// allocated capnp.Int8List, preferring placement in s's segment.
func (s Z) NewI8vec(n int32) (capnp.Int8List, error) {
	s.Struct.SetUint16(0, 20)
	l, err := capnp.NewInt8List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int8List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U64vec() (capnp.UInt64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt64List{List: p.List()}, err
}

func (s Z) HasU64vec() bool {
	if s.Struct.Uint16(0) != 21 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU64vec(v capnp.UInt64List) error {
	s.Struct.SetUint16(0, 21)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU64vec sets the u64vec field to a newly
// allocated capnp.UInt64List, preferring placement in s's segment.
func (s Z) NewU64vec(n int32) (capnp.UInt64List, error) {
	s.Struct.SetUint16(0, 21)
	l, err := capnp.NewUInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U32vec() (capnp.UInt32List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt32List{List: p.List()}, err
}

func (s Z) HasU32vec() bool {
	if s.Struct.Uint16(0) != 22 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU32vec(v capnp.UInt32List) error {
	s.Struct.SetUint16(0, 22)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU32vec sets the u32vec field to a newly
// allocated capnp.UInt32List, preferring placement in s's segment.
func (s Z) NewU32vec(n int32) (capnp.UInt32List, error) {
	s.Struct.SetUint16(0, 22)
	l, err := capnp.NewUInt32List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt32List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U16vec() (capnp.UInt16List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt16List{List: p.List()}, err
}

func (s Z) HasU16vec() bool {
	if s.Struct.Uint16(0) != 23 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU16vec(v capnp.UInt16List) error {
	s.Struct.SetUint16(0, 23)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU16vec sets the u16vec field to a newly
// allocated capnp.UInt16List, preferring placement in s's segment.
func (s Z) NewU16vec(n int32) (capnp.UInt16List, error) {
	s.Struct.SetUint16(0, 23)
	l, err := capnp.NewUInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) U8vec() (capnp.UInt8List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt8List{List: p.List()}, err
}

func (s Z) HasU8vec() bool {
	if s.Struct.Uint16(0) != 24 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetU8vec(v capnp.UInt8List) error {
	s.Struct.SetUint16(0, 24)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewU8vec sets the u8vec field to a newly
// allocated capnp.UInt8List, preferring placement in s's segment.
func (s Z) NewU8vec(n int32) (capnp.UInt8List, error) {
	s.Struct.SetUint16(0, 24)
	l, err := capnp.NewUInt8List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt8List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Boolvec() (capnp.BitList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.BitList{List: p.List()}, err
}

func (s Z) HasBoolvec() bool {
	if s.Struct.Uint16(0) != 39 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetBoolvec(v capnp.BitList) error {
	s.Struct.SetUint16(0, 39)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBoolvec sets the boolvec field to a newly
// allocated capnp.BitList, preferring placement in s's segment.
func (s Z) NewBoolvec(n int32) (capnp.BitList, error) {
	s.Struct.SetUint16(0, 39)
	l, err := capnp.NewBitList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.BitList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Datavec() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.DataList{List: p.List()}, err
}

func (s Z) HasDatavec() bool {
	if s.Struct.Uint16(0) != 40 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetDatavec(v capnp.DataList) error {
	s.Struct.SetUint16(0, 40)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewDatavec sets the datavec field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Z) NewDatavec(n int32) (capnp.DataList, error) {
	s.Struct.SetUint16(0, 40)
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Textvec() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s Z) HasTextvec() bool {
	if s.Struct.Uint16(0) != 41 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetTextvec(v capnp.TextList) error {
	s.Struct.SetUint16(0, 41)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTextvec sets the textvec field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Z) NewTextvec(n int32) (capnp.TextList, error) {
	s.Struct.SetUint16(0, 41)
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zvec() (Z_List, error) {
	p, err := s.Struct.Ptr(0)
	return Z_List{List: p.List()}, err
}

func (s Z) HasZvec() bool {
	if s.Struct.Uint16(0) != 25 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZvec(v Z_List) error {
	s.Struct.SetUint16(0, 25)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZvec sets the zvec field to a newly
// allocated Z_List, preferring placement in s's segment.
func (s Z) NewZvec(n int32) (Z_List, error) {
	s.Struct.SetUint16(0, 25)
	l, err := NewZ_List(s.Struct.Segment(), n)
	if err != nil {
		return Z_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zvecvec() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.PointerList{List: p.List()}, err
}

func (s Z) HasZvecvec() bool {
	if s.Struct.Uint16(0) != 26 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZvecvec(v capnp.PointerList) error {
	s.Struct.SetUint16(0, 26)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZvecvec sets the zvecvec field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s Z) NewZvecvec(n int32) (capnp.PointerList, error) {
	s.Struct.SetUint16(0, 26)
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zdate() (Zdate, error) {
	p, err := s.Struct.Ptr(0)
	return Zdate{Struct: p.Struct()}, err
}

func (s Z) HasZdate() bool {
	if s.Struct.Uint16(0) != 27 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdate(v Zdate) error {
	s.Struct.SetUint16(0, 27)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZdate sets the zdate field to a newly
// allocated Zdate struct, preferring placement in s's segment.
func (s Z) NewZdate() (Zdate, error) {
	s.Struct.SetUint16(0, 27)
	ss, err := NewZdate(s.Struct.Segment())
	if err != nil {
		return Zdate{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Zdata() (Zdata, error) {
	p, err := s.Struct.Ptr(0)
	return Zdata{Struct: p.Struct()}, err
}

func (s Z) HasZdata() bool {
	if s.Struct.Uint16(0) != 28 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdata(v Zdata) error {
	s.Struct.SetUint16(0, 28)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewZdata sets the zdata field to a newly
// allocated Zdata struct, preferring placement in s's segment.
func (s Z) NewZdata() (Zdata, error) {
	s.Struct.SetUint16(0, 28)
	ss, err := NewZdata(s.Struct.Segment())
	if err != nil {
		return Zdata{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Aircraftvec() (Aircraft_List, error) {
	p, err := s.Struct.Ptr(0)
	return Aircraft_List{List: p.List()}, err
}

func (s Z) HasAircraftvec() bool {
	if s.Struct.Uint16(0) != 29 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetAircraftvec(v Aircraft_List) error {
	s.Struct.SetUint16(0, 29)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewAircraftvec sets the aircraftvec field to a newly
// allocated Aircraft_List, preferring placement in s's segment.
func (s Z) NewAircraftvec(n int32) (Aircraft_List, error) {
	s.Struct.SetUint16(0, 29)
	l, err := NewAircraft_List(s.Struct.Segment(), n)
	if err != nil {
		return Aircraft_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Aircraft() (Aircraft, error) {
	p, err := s.Struct.Ptr(0)
	return Aircraft{Struct: p.Struct()}, err
}

func (s Z) HasAircraft() bool {
	if s.Struct.Uint16(0) != 30 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetAircraft(v Aircraft) error {
	s.Struct.SetUint16(0, 30)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAircraft sets the aircraft field to a newly
// allocated Aircraft struct, preferring placement in s's segment.
func (s Z) NewAircraft() (Aircraft, error) {
	s.Struct.SetUint16(0, 30)
	ss, err := NewAircraft(s.Struct.Segment())
	if err != nil {
		return Aircraft{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Regression() (Regression, error) {
	p, err := s.Struct.Ptr(0)
	return Regression{Struct: p.Struct()}, err
}

func (s Z) HasRegression() bool {
	if s.Struct.Uint16(0) != 31 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetRegression(v Regression) error {
	s.Struct.SetUint16(0, 31)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRegression sets the regression field to a newly
// allocated Regression struct, preferring placement in s's segment.
func (s Z) NewRegression() (Regression, error) {
	s.Struct.SetUint16(0, 31)
	ss, err := NewRegression(s.Struct.Segment())
	if err != nil {
		return Regression{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Planebase() (PlaneBase, error) {
	p, err := s.Struct.Ptr(0)
	return PlaneBase{Struct: p.Struct()}, err
}

func (s Z) HasPlanebase() bool {
	if s.Struct.Uint16(0) != 32 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetPlanebase(v PlaneBase) error {
	s.Struct.SetUint16(0, 32)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPlanebase sets the planebase field to a newly
// allocated PlaneBase struct, preferring placement in s's segment.
func (s Z) NewPlanebase() (PlaneBase, error) {
	s.Struct.SetUint16(0, 32)
	ss, err := NewPlaneBase(s.Struct.Segment())
	if err != nil {
		return PlaneBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Airport() Airport {
	return Airport(s.Struct.Uint16(8))
}

func (s Z) SetAirport(v Airport) {
	s.Struct.SetUint16(0, 33)
	s.Struct.SetUint16(8, uint16(v))
}

func (s Z) B737() (B737, error) {
	p, err := s.Struct.Ptr(0)
	return B737{Struct: p.Struct()}, err
}

func (s Z) HasB737() bool {
	if s.Struct.Uint16(0) != 34 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetB737(v B737) error {
	s.Struct.SetUint16(0, 34)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB737 sets the b737 field to a newly
// allocated B737 struct, preferring placement in s's segment.
func (s Z) NewB737() (B737, error) {
	s.Struct.SetUint16(0, 34)
	ss, err := NewB737(s.Struct.Segment())
	if err != nil {
		return B737{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) A320() (A320, error) {
	p, err := s.Struct.Ptr(0)
	return A320{Struct: p.Struct()}, err
}

func (s Z) HasA320() bool {
	if s.Struct.Uint16(0) != 35 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetA320(v A320) error {
	s.Struct.SetUint16(0, 35)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewA320 sets the a320 field to a newly
// allocated A320 struct, preferring placement in s's segment.
func (s Z) NewA320() (A320, error) {
	s.Struct.SetUint16(0, 35)
	ss, err := NewA320(s.Struct.Segment())
	if err != nil {
		return A320{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) F16() (F16, error) {
	p, err := s.Struct.Ptr(0)
	return F16{Struct: p.Struct()}, err
}

func (s Z) HasF16() bool {
	if s.Struct.Uint16(0) != 36 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetF16(v F16) error {
	s.Struct.SetUint16(0, 36)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewF16 sets the f16 field to a newly
// allocated F16 struct, preferring placement in s's segment.
func (s Z) NewF16() (F16, error) {
	s.Struct.SetUint16(0, 36)
	ss, err := NewF16(s.Struct.Segment())
	if err != nil {
		return F16{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Z) Zdatevec() (Zdate_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zdate_List{List: p.List()}, err
}

func (s Z) HasZdatevec() bool {
	if s.Struct.Uint16(0) != 37 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdatevec(v Zdate_List) error {
	s.Struct.SetUint16(0, 37)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZdatevec sets the zdatevec field to a newly
// allocated Zdate_List, preferring placement in s's segment.
func (s Z) NewZdatevec(n int32) (Zdate_List, error) {
	s.Struct.SetUint16(0, 37)
	l, err := NewZdate_List(s.Struct.Segment(), n)
	if err != nil {
		return Zdate_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Zdatavec() (Zdata_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zdata_List{List: p.List()}, err
}

func (s Z) HasZdatavec() bool {
	if s.Struct.Uint16(0) != 38 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Z) SetZdatavec(v Zdata_List) error {
	s.Struct.SetUint16(0, 38)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewZdatavec sets the zdatavec field to a newly
// allocated Zdata_List, preferring placement in s's segment.
func (s Z) NewZdatavec(n int32) (Zdata_List, error) {
	s.Struct.SetUint16(0, 38)
	l, err := NewZdata_List(s.Struct.Segment(), n)
	if err != nil {
		return Zdata_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s Z) Grp() Z_grp { return Z_grp(s) }

func (s Z) SetGrp() {
	s.Struct.SetUint16(0, 42)
}

func (s Z_grp) First() uint64 {
	return s.Struct.Uint64(8)
}

func (s Z_grp) SetFirst(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Z_grp) Second() uint64 {
	return s.Struct.Uint64(16)
}

func (s Z_grp) SetSecond(v uint64) {
	s.Struct.SetUint64(16, v)
}

// Z_List is a list of Z.
type Z_List struct{ capnp.List }

// NewZ creates a new list of Z.
func NewZ_List(s *capnp.Segment, sz int32) (Z_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	return Z_List{l}, err
}

func (s Z_List) At(i int) Z { return Z{s.List.Struct(i)} }

func (s Z_List) Set(i int, v Z) error { return s.List.SetStruct(i, v.Struct) }

func (s Z_List) Copy(seg *capnp.Segment) (Z_List, error) {
	var err error
	t, err := NewZ_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Z_Promise is a wrapper for a Z promised by a client call.
type Z_Promise struct{ *capnp.Pipeline }

func (p Z_Promise) Struct() (Z, error) {
	s, err := p.Pipeline.Struct()
	return Z{s}, err
}

func (p Z_Promise) Zz() Z_Promise {
	return Z_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdate() Zdate_Promise {
	return Zdate_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Zdata() Zdata_Promise {
	return Zdata_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Aircraft() Aircraft_Promise {
	return Aircraft_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Regression() Regression_Promise {
	return Regression_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Planebase() PlaneBase_Promise {
	return PlaneBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) B737() B737_Promise {
	return B737_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) A320() A320_Promise {
	return A320_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) F16() F16_Promise {
	return F16_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Z_Promise) Grp() Z_grp_Promise { return Z_grp_Promise{p.Pipeline} }

// Z_grp_Promise is a wrapper for a Z_grp promised by a client call.
type Z_grp_Promise struct{ *capnp.Pipeline }

func (p Z_grp_Promise) Struct() (Z_grp, error) {
	s, err := p.Pipeline.Struct()
	return Z_grp{s}, err
}

type Counter struct{ capnp.Struct }

// Counter_TypeID is the unique identifier for the type Counter.
const Counter_TypeID = 0x8748bc095e10cb5d

func NewCounter(s *capnp.Segment) (Counter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Counter{st}, err
}

func NewRootCounter(s *capnp.Segment) (Counter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Counter{st}, err
}

func ReadRootCounter(msg *capnp.Message) (Counter, error) {
	root, err := msg.RootPtr()
	return Counter{root.Struct()}, err
}

func (s Counter) String() string {
	str, _ := text.Marshal(0x8748bc095e10cb5d, s.Struct)
	return str
}

func (s Counter) Copy(seg *capnp.Segment) (Counter, error) {
	var err error
	t, err := NewCounter(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Wordlist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetWordlist(v)
	}
	{
		v, err := s.Words()
		if err != nil {
			return t, err
		}
		t.SetWords(v)
	}
	t.SetSize(s.Size())
	return t, nil
}

func (s Counter) Size() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Counter) SetSize(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Counter) Words() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Counter) HasWords() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Counter) WordsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Counter) SetWords(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Counter) Wordlist() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Counter) HasWordlist() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Counter) SetWordlist(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewWordlist sets the wordlist field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Counter) NewWordlist(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Counter_List is a list of Counter.
type Counter_List struct{ capnp.List }

// NewCounter creates a new list of Counter.
func NewCounter_List(s *capnp.Segment, sz int32) (Counter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Counter_List{l}, err
}

func (s Counter_List) At(i int) Counter { return Counter{s.List.Struct(i)} }

func (s Counter_List) Set(i int, v Counter) error { return s.List.SetStruct(i, v.Struct) }

func (s Counter_List) Copy(seg *capnp.Segment) (Counter_List, error) {
	var err error
	t, err := NewCounter_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Counter_Promise is a wrapper for a Counter promised by a client call.
type Counter_Promise struct{ *capnp.Pipeline }

func (p Counter_Promise) Struct() (Counter, error) {
	s, err := p.Pipeline.Struct()
	return Counter{s}, err
}

type Bag struct{ capnp.Struct }

// Bag_TypeID is the unique identifier for the type Bag.
const Bag_TypeID = 0xd636fba4f188dabe

func NewBag(s *capnp.Segment) (Bag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bag{st}, err
}

func NewRootBag(s *capnp.Segment) (Bag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bag{st}, err
}

func ReadRootBag(msg *capnp.Message) (Bag, error) {
	root, err := msg.RootPtr()
	return Bag{root.Struct()}, err
}

func (s Bag) String() string {
	str, _ := text.Marshal(0xd636fba4f188dabe, s.Struct)
	return str
}

func (s Bag) Copy(seg *capnp.Segment) (Bag, error) {
	var err error
	t, err := NewBag(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Counter()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetCounter(v)
	}
	return t, nil
}

func (s Bag) Counter() (Counter, error) {
	p, err := s.Struct.Ptr(0)
	return Counter{Struct: p.Struct()}, err
}

func (s Bag) HasCounter() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Bag) SetCounter(v Counter) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCounter sets the counter field to a newly
// allocated Counter struct, preferring placement in s's segment.
func (s Bag) NewCounter() (Counter, error) {
	ss, err := NewCounter(s.Struct.Segment())
	if err != nil {
		return Counter{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Bag_List is a list of Bag.
type Bag_List struct{ capnp.List }

// NewBag creates a new list of Bag.
func NewBag_List(s *capnp.Segment, sz int32) (Bag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Bag_List{l}, err
}

func (s Bag_List) At(i int) Bag { return Bag{s.List.Struct(i)} }

func (s Bag_List) Set(i int, v Bag) error { return s.List.SetStruct(i, v.Struct) }

func (s Bag_List) Copy(seg *capnp.Segment) (Bag_List, error) {
	var err error
	t, err := NewBag_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Bag_Promise is a wrapper for a Bag promised by a client call.
type Bag_Promise struct{ *capnp.Pipeline }

func (p Bag_Promise) Struct() (Bag, error) {
	s, err := p.Pipeline.Struct()
	return Bag{s}, err
}

func (p Bag_Promise) Counter() Counter_Promise {
	return Counter_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Zserver struct{ capnp.Struct }

// Zserver_TypeID is the unique identifier for the type Zserver.
const Zserver_TypeID = 0xcc4411e60ba9c498

func NewZserver(s *capnp.Segment) (Zserver, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zserver{st}, err
}

func NewRootZserver(s *capnp.Segment) (Zserver, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Zserver{st}, err
}

func ReadRootZserver(msg *capnp.Message) (Zserver, error) {
	root, err := msg.RootPtr()
	return Zserver{root.Struct()}, err
}

func (s Zserver) String() string {
	str, _ := text.Marshal(0xcc4411e60ba9c498, s.Struct)
	return str
}

func (s Zserver) Copy(seg *capnp.Segment) (Zserver, error) {
	var err error
	t, err := NewZserver(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Waitingjobs()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetWaitingjobs(v)
	}
	return t, nil
}

func (s Zserver) Waitingjobs() (Zjob_List, error) {
	p, err := s.Struct.Ptr(0)
	return Zjob_List{List: p.List()}, err
}

func (s Zserver) HasWaitingjobs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zserver) SetWaitingjobs(v Zjob_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewWaitingjobs sets the waitingjobs field to a newly
// allocated Zjob_List, preferring placement in s's segment.
func (s Zserver) NewWaitingjobs(n int32) (Zjob_List, error) {
	l, err := NewZjob_List(s.Struct.Segment(), n)
	if err != nil {
		return Zjob_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Zserver_List is a list of Zserver.
type Zserver_List struct{ capnp.List }

// NewZserver creates a new list of Zserver.
func NewZserver_List(s *capnp.Segment, sz int32) (Zserver_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Zserver_List{l}, err
}

func (s Zserver_List) At(i int) Zserver { return Zserver{s.List.Struct(i)} }

func (s Zserver_List) Set(i int, v Zserver) error { return s.List.SetStruct(i, v.Struct) }

func (s Zserver_List) Copy(seg *capnp.Segment) (Zserver_List, error) {
	var err error
	t, err := NewZserver_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Zserver_Promise is a wrapper for a Zserver promised by a client call.
type Zserver_Promise struct{ *capnp.Pipeline }

func (p Zserver_Promise) Struct() (Zserver, error) {
	s, err := p.Pipeline.Struct()
	return Zserver{s}, err
}

type Zjob struct{ capnp.Struct }

// Zjob_TypeID is the unique identifier for the type Zjob.
const Zjob_TypeID = 0xddd1416669fb7613

func NewZjob(s *capnp.Segment) (Zjob, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Zjob{st}, err
}

func NewRootZjob(s *capnp.Segment) (Zjob, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Zjob{st}, err
}

func ReadRootZjob(msg *capnp.Message) (Zjob, error) {
	root, err := msg.RootPtr()
	return Zjob{root.Struct()}, err
}

func (s Zjob) String() string {
	str, _ := text.Marshal(0xddd1416669fb7613, s.Struct)
	return str
}

func (s Zjob) Copy(seg *capnp.Segment) (Zjob, error) {
	var err error
	t, err := NewZjob(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Args()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetArgs(v)
	}
	{
		v, err := s.Cmd()
		if err != nil {
			return t, err
		}
		t.SetCmd(v)
	}
	return t, nil
}

func (s Zjob) Cmd() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Zjob) HasCmd() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Zjob) CmdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Zjob) SetCmd(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Zjob) Args() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s Zjob) HasArgs() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Zjob) SetArgs(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewArgs sets the args field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Zjob) NewArgs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Zjob_List is a list of Zjob.
type Zjob_List struct{ capnp.List }

// NewZjob creates a new list of Zjob.
func NewZjob_List(s *capnp.Segment, sz int32) (Zjob_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Zjob_List{l}, err
}

func (s Zjob_List) At(i int) Zjob { return Zjob{s.List.Struct(i)} }

func (s Zjob_List) Set(i int, v Zjob) error { return s.List.SetStruct(i, v.Struct) }

func (s Zjob_List) Copy(seg *capnp.Segment) (Zjob_List, error) {
	var err error
	t, err := NewZjob_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Zjob_Promise is a wrapper for a Zjob promised by a client call.
type Zjob_Promise struct{ *capnp.Pipeline }

func (p Zjob_Promise) Struct() (Zjob, error) {
	s, err := p.Pipeline.Struct()
	return Zjob{s}, err
}

type VerEmpty struct{ capnp.Struct }

// VerEmpty_TypeID is the unique identifier for the type VerEmpty.
const VerEmpty_TypeID = 0x93c99951eacc72ff

func NewVerEmpty(s *capnp.Segment) (VerEmpty, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return VerEmpty{st}, err
}

func NewRootVerEmpty(s *capnp.Segment) (VerEmpty, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return VerEmpty{st}, err
}

func ReadRootVerEmpty(msg *capnp.Message) (VerEmpty, error) {
	root, err := msg.RootPtr()
	return VerEmpty{root.Struct()}, err
}

func (s VerEmpty) String() string {
	str, _ := text.Marshal(0x93c99951eacc72ff, s.Struct)
	return str
}

func (s VerEmpty) Copy(seg *capnp.Segment) (VerEmpty, error) {
	var err error
	t, err := NewVerEmpty(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

// VerEmpty_List is a list of VerEmpty.
type VerEmpty_List struct{ capnp.List }

// NewVerEmpty creates a new list of VerEmpty.
func NewVerEmpty_List(s *capnp.Segment, sz int32) (VerEmpty_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return VerEmpty_List{l}, err
}

func (s VerEmpty_List) At(i int) VerEmpty { return VerEmpty{s.List.Struct(i)} }

func (s VerEmpty_List) Set(i int, v VerEmpty) error { return s.List.SetStruct(i, v.Struct) }

func (s VerEmpty_List) Copy(seg *capnp.Segment) (VerEmpty_List, error) {
	var err error
	t, err := NewVerEmpty_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerEmpty_Promise is a wrapper for a VerEmpty promised by a client call.
type VerEmpty_Promise struct{ *capnp.Pipeline }

func (p VerEmpty_Promise) Struct() (VerEmpty, error) {
	s, err := p.Pipeline.Struct()
	return VerEmpty{s}, err
}

type VerOneData struct{ capnp.Struct }

// VerOneData_TypeID is the unique identifier for the type VerOneData.
const VerOneData_TypeID = 0xfca3742893be4cde

func NewVerOneData(s *capnp.Segment) (VerOneData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VerOneData{st}, err
}

func NewRootVerOneData(s *capnp.Segment) (VerOneData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VerOneData{st}, err
}

func ReadRootVerOneData(msg *capnp.Message) (VerOneData, error) {
	root, err := msg.RootPtr()
	return VerOneData{root.Struct()}, err
}

func (s VerOneData) String() string {
	str, _ := text.Marshal(0xfca3742893be4cde, s.Struct)
	return str
}

func (s VerOneData) Copy(seg *capnp.Segment) (VerOneData, error) {
	var err error
	t, err := NewVerOneData(seg)
	if err != nil {
		return t, err
	}
	t.SetVal(s.Val())
	return t, nil
}

func (s VerOneData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerOneData) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

// VerOneData_List is a list of VerOneData.
type VerOneData_List struct{ capnp.List }

// NewVerOneData creates a new list of VerOneData.
func NewVerOneData_List(s *capnp.Segment, sz int32) (VerOneData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return VerOneData_List{l}, err
}

func (s VerOneData_List) At(i int) VerOneData { return VerOneData{s.List.Struct(i)} }

func (s VerOneData_List) Set(i int, v VerOneData) error { return s.List.SetStruct(i, v.Struct) }

func (s VerOneData_List) Copy(seg *capnp.Segment) (VerOneData_List, error) {
	var err error
	t, err := NewVerOneData_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerOneData_Promise is a wrapper for a VerOneData promised by a client call.
type VerOneData_Promise struct{ *capnp.Pipeline }

func (p VerOneData_Promise) Struct() (VerOneData, error) {
	s, err := p.Pipeline.Struct()
	return VerOneData{s}, err
}

type VerTwoData struct{ capnp.Struct }

// VerTwoData_TypeID is the unique identifier for the type VerTwoData.
const VerTwoData_TypeID = 0xf705dc45c94766fd

func NewVerTwoData(s *capnp.Segment) (VerTwoData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return VerTwoData{st}, err
}

func NewRootVerTwoData(s *capnp.Segment) (VerTwoData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0})
	return VerTwoData{st}, err
}

func ReadRootVerTwoData(msg *capnp.Message) (VerTwoData, error) {
	root, err := msg.RootPtr()
	return VerTwoData{root.Struct()}, err
}

func (s VerTwoData) String() string {
	str, _ := text.Marshal(0xf705dc45c94766fd, s.Struct)
	return str
}

func (s VerTwoData) Copy(seg *capnp.Segment) (VerTwoData, error) {
	var err error
	t, err := NewVerTwoData(seg)
	if err != nil {
		return t, err
	}
	t.SetVal(s.Val())
	t.SetDuo(s.Duo())
	return t, nil
}

func (s VerTwoData) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoData) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoData) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoData) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

// VerTwoData_List is a list of VerTwoData.
type VerTwoData_List struct{ capnp.List }

// NewVerTwoData creates a new list of VerTwoData.
func NewVerTwoData_List(s *capnp.Segment, sz int32) (VerTwoData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 0}, sz)
	return VerTwoData_List{l}, err
}

func (s VerTwoData_List) At(i int) VerTwoData { return VerTwoData{s.List.Struct(i)} }

func (s VerTwoData_List) Set(i int, v VerTwoData) error { return s.List.SetStruct(i, v.Struct) }

func (s VerTwoData_List) Copy(seg *capnp.Segment) (VerTwoData_List, error) {
	var err error
	t, err := NewVerTwoData_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerTwoData_Promise is a wrapper for a VerTwoData promised by a client call.
type VerTwoData_Promise struct{ *capnp.Pipeline }

func (p VerTwoData_Promise) Struct() (VerTwoData, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoData{s}, err
}

type VerOnePtr struct{ capnp.Struct }

// VerOnePtr_TypeID is the unique identifier for the type VerOnePtr.
const VerOnePtr_TypeID = 0x94bf7df83408218d

func NewVerOnePtr(s *capnp.Segment) (VerOnePtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VerOnePtr{st}, err
}

func NewRootVerOnePtr(s *capnp.Segment) (VerOnePtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return VerOnePtr{st}, err
}

func ReadRootVerOnePtr(msg *capnp.Message) (VerOnePtr, error) {
	root, err := msg.RootPtr()
	return VerOnePtr{root.Struct()}, err
}

func (s VerOnePtr) String() string {
	str, _ := text.Marshal(0x94bf7df83408218d, s.Struct)
	return str
}

func (s VerOnePtr) Copy(seg *capnp.Segment) (VerOnePtr, error) {
	var err error
	t, err := NewVerOnePtr(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Ptr()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr(v)
	}
	return t, nil
}

func (s VerOnePtr) Ptr() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerOnePtr) HasPtr() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerOnePtr) SetPtr(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr sets the ptr field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerOnePtr) NewPtr() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// VerOnePtr_List is a list of VerOnePtr.
type VerOnePtr_List struct{ capnp.List }

// NewVerOnePtr creates a new list of VerOnePtr.
func NewVerOnePtr_List(s *capnp.Segment, sz int32) (VerOnePtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return VerOnePtr_List{l}, err
}

func (s VerOnePtr_List) At(i int) VerOnePtr { return VerOnePtr{s.List.Struct(i)} }

func (s VerOnePtr_List) Set(i int, v VerOnePtr) error { return s.List.SetStruct(i, v.Struct) }

func (s VerOnePtr_List) Copy(seg *capnp.Segment) (VerOnePtr_List, error) {
	var err error
	t, err := NewVerOnePtr_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerOnePtr_Promise is a wrapper for a VerOnePtr promised by a client call.
type VerOnePtr_Promise struct{ *capnp.Pipeline }

func (p VerOnePtr_Promise) Struct() (VerOnePtr, error) {
	s, err := p.Pipeline.Struct()
	return VerOnePtr{s}, err
}

func (p VerOnePtr_Promise) Ptr() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VerTwoPtr struct{ capnp.Struct }

// VerTwoPtr_TypeID is the unique identifier for the type VerTwoPtr.
const VerTwoPtr_TypeID = 0xc95babe3bd394d2d

func NewVerTwoPtr(s *capnp.Segment) (VerTwoPtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerTwoPtr{st}, err
}

func NewRootVerTwoPtr(s *capnp.Segment) (VerTwoPtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return VerTwoPtr{st}, err
}

func ReadRootVerTwoPtr(msg *capnp.Message) (VerTwoPtr, error) {
	root, err := msg.RootPtr()
	return VerTwoPtr{root.Struct()}, err
}

func (s VerTwoPtr) String() string {
	str, _ := text.Marshal(0xc95babe3bd394d2d, s.Struct)
	return str
}

func (s VerTwoPtr) Copy(seg *capnp.Segment) (VerTwoPtr, error) {
	var err error
	t, err := NewVerTwoPtr(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Ptr1()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr1(v)
	}
	{
		o, err := s.Ptr2()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr2(v)
	}
	return t, nil
}

func (s VerTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoPtr) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoPtr) SetPtr1(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr1() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Ptr(1)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoPtr) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoPtr) SetPtr2(v VerOneData) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoPtr) NewPtr2() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerTwoPtr_List is a list of VerTwoPtr.
type VerTwoPtr_List struct{ capnp.List }

// NewVerTwoPtr creates a new list of VerTwoPtr.
func NewVerTwoPtr_List(s *capnp.Segment, sz int32) (VerTwoPtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return VerTwoPtr_List{l}, err
}

func (s VerTwoPtr_List) At(i int) VerTwoPtr { return VerTwoPtr{s.List.Struct(i)} }

func (s VerTwoPtr_List) Set(i int, v VerTwoPtr) error { return s.List.SetStruct(i, v.Struct) }

func (s VerTwoPtr_List) Copy(seg *capnp.Segment) (VerTwoPtr_List, error) {
	var err error
	t, err := NewVerTwoPtr_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerTwoPtr_Promise is a wrapper for a VerTwoPtr promised by a client call.
type VerTwoPtr_Promise struct{ *capnp.Pipeline }

func (p VerTwoPtr_Promise) Struct() (VerTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoPtr{s}, err
}

func (p VerTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type VerTwoDataTwoPtr struct{ capnp.Struct }

// VerTwoDataTwoPtr_TypeID is the unique identifier for the type VerTwoDataTwoPtr.
const VerTwoDataTwoPtr_TypeID = 0xb61ee2ecff34ca73

func NewVerTwoDataTwoPtr(s *capnp.Segment) (VerTwoDataTwoPtr, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return VerTwoDataTwoPtr{st}, err
}

func NewRootVerTwoDataTwoPtr(s *capnp.Segment) (VerTwoDataTwoPtr, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return VerTwoDataTwoPtr{st}, err
}

func ReadRootVerTwoDataTwoPtr(msg *capnp.Message) (VerTwoDataTwoPtr, error) {
	root, err := msg.RootPtr()
	return VerTwoDataTwoPtr{root.Struct()}, err
}

func (s VerTwoDataTwoPtr) String() string {
	str, _ := text.Marshal(0xb61ee2ecff34ca73, s.Struct)
	return str
}

func (s VerTwoDataTwoPtr) Copy(seg *capnp.Segment) (VerTwoDataTwoPtr, error) {
	var err error
	t, err := NewVerTwoDataTwoPtr(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Ptr1()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr1(v)
	}
	{
		o, err := s.Ptr2()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr2(v)
	}
	t.SetVal(s.Val())
	t.SetDuo(s.Duo())
	return t, nil
}

func (s VerTwoDataTwoPtr) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoDataTwoPtr) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoDataTwoPtr) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoDataTwoPtr) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoDataTwoPtr) Ptr1() (VerOneData, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoDataTwoPtr) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoDataTwoPtr) SetPtr1(v VerOneData) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr1() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoDataTwoPtr) Ptr2() (VerOneData, error) {
	p, err := s.Struct.Ptr(1)
	return VerOneData{Struct: p.Struct()}, err
}

func (s VerTwoDataTwoPtr) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoDataTwoPtr) SetPtr2(v VerOneData) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerOneData struct, preferring placement in s's segment.
func (s VerTwoDataTwoPtr) NewPtr2() (VerOneData, error) {
	ss, err := NewVerOneData(s.Struct.Segment())
	if err != nil {
		return VerOneData{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// VerTwoDataTwoPtr_List is a list of VerTwoDataTwoPtr.
type VerTwoDataTwoPtr_List struct{ capnp.List }

// NewVerTwoDataTwoPtr creates a new list of VerTwoDataTwoPtr.
func NewVerTwoDataTwoPtr_List(s *capnp.Segment, sz int32) (VerTwoDataTwoPtr_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return VerTwoDataTwoPtr_List{l}, err
}

func (s VerTwoDataTwoPtr_List) At(i int) VerTwoDataTwoPtr { return VerTwoDataTwoPtr{s.List.Struct(i)} }

func (s VerTwoDataTwoPtr_List) Set(i int, v VerTwoDataTwoPtr) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s VerTwoDataTwoPtr_List) Copy(seg *capnp.Segment) (VerTwoDataTwoPtr_List, error) {
	var err error
	t, err := NewVerTwoDataTwoPtr_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerTwoDataTwoPtr_Promise is a wrapper for a VerTwoDataTwoPtr promised by a client call.
type VerTwoDataTwoPtr_Promise struct{ *capnp.Pipeline }

func (p VerTwoDataTwoPtr_Promise) Struct() (VerTwoDataTwoPtr, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoDataTwoPtr{s}, err
}

func (p VerTwoDataTwoPtr_Promise) Ptr1() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoDataTwoPtr_Promise) Ptr2() VerOneData_Promise {
	return VerOneData_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsVerEmptyList struct{ capnp.Struct }

// HoldsVerEmptyList_TypeID is the unique identifier for the type HoldsVerEmptyList.
const HoldsVerEmptyList_TypeID = 0xde9ed43cfaa83093

func NewHoldsVerEmptyList(s *capnp.Segment) (HoldsVerEmptyList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerEmptyList{st}, err
}

func NewRootHoldsVerEmptyList(s *capnp.Segment) (HoldsVerEmptyList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerEmptyList{st}, err
}

func ReadRootHoldsVerEmptyList(msg *capnp.Message) (HoldsVerEmptyList, error) {
	root, err := msg.RootPtr()
	return HoldsVerEmptyList{root.Struct()}, err
}

func (s HoldsVerEmptyList) String() string {
	str, _ := text.Marshal(0xde9ed43cfaa83093, s.Struct)
	return str
}

func (s HoldsVerEmptyList) Copy(seg *capnp.Segment) (HoldsVerEmptyList, error) {
	var err error
	t, err := NewHoldsVerEmptyList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerEmptyList) Mylist() (VerEmpty_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerEmpty_List{List: p.List()}, err
}

func (s HoldsVerEmptyList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerEmptyList) SetMylist(v VerEmpty_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerEmpty_List, preferring placement in s's segment.
func (s HoldsVerEmptyList) NewMylist(n int32) (VerEmpty_List, error) {
	l, err := NewVerEmpty_List(s.Struct.Segment(), n)
	if err != nil {
		return VerEmpty_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerEmptyList_List is a list of HoldsVerEmptyList.
type HoldsVerEmptyList_List struct{ capnp.List }

// NewHoldsVerEmptyList creates a new list of HoldsVerEmptyList.
func NewHoldsVerEmptyList_List(s *capnp.Segment, sz int32) (HoldsVerEmptyList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerEmptyList_List{l}, err
}

func (s HoldsVerEmptyList_List) At(i int) HoldsVerEmptyList {
	return HoldsVerEmptyList{s.List.Struct(i)}
}

func (s HoldsVerEmptyList_List) Set(i int, v HoldsVerEmptyList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerEmptyList_List) Copy(seg *capnp.Segment) (HoldsVerEmptyList_List, error) {
	var err error
	t, err := NewHoldsVerEmptyList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerEmptyList_Promise is a wrapper for a HoldsVerEmptyList promised by a client call.
type HoldsVerEmptyList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerEmptyList_Promise) Struct() (HoldsVerEmptyList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerEmptyList{s}, err
}

type HoldsVerOneDataList struct{ capnp.Struct }

// HoldsVerOneDataList_TypeID is the unique identifier for the type HoldsVerOneDataList.
const HoldsVerOneDataList_TypeID = 0xabd055422a4d7df1

func NewHoldsVerOneDataList(s *capnp.Segment) (HoldsVerOneDataList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOneDataList{st}, err
}

func NewRootHoldsVerOneDataList(s *capnp.Segment) (HoldsVerOneDataList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOneDataList{st}, err
}

func ReadRootHoldsVerOneDataList(msg *capnp.Message) (HoldsVerOneDataList, error) {
	root, err := msg.RootPtr()
	return HoldsVerOneDataList{root.Struct()}, err
}

func (s HoldsVerOneDataList) String() string {
	str, _ := text.Marshal(0xabd055422a4d7df1, s.Struct)
	return str
}

func (s HoldsVerOneDataList) Copy(seg *capnp.Segment) (HoldsVerOneDataList, error) {
	var err error
	t, err := NewHoldsVerOneDataList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerOneDataList) Mylist() (VerOneData_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerOneData_List{List: p.List()}, err
}

func (s HoldsVerOneDataList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerOneDataList) SetMylist(v VerOneData_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerOneData_List, preferring placement in s's segment.
func (s HoldsVerOneDataList) NewMylist(n int32) (VerOneData_List, error) {
	l, err := NewVerOneData_List(s.Struct.Segment(), n)
	if err != nil {
		return VerOneData_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerOneDataList_List is a list of HoldsVerOneDataList.
type HoldsVerOneDataList_List struct{ capnp.List }

// NewHoldsVerOneDataList creates a new list of HoldsVerOneDataList.
func NewHoldsVerOneDataList_List(s *capnp.Segment, sz int32) (HoldsVerOneDataList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerOneDataList_List{l}, err
}

func (s HoldsVerOneDataList_List) At(i int) HoldsVerOneDataList {
	return HoldsVerOneDataList{s.List.Struct(i)}
}

func (s HoldsVerOneDataList_List) Set(i int, v HoldsVerOneDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerOneDataList_List) Copy(seg *capnp.Segment) (HoldsVerOneDataList_List, error) {
	var err error
	t, err := NewHoldsVerOneDataList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerOneDataList_Promise is a wrapper for a HoldsVerOneDataList promised by a client call.
type HoldsVerOneDataList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerOneDataList_Promise) Struct() (HoldsVerOneDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOneDataList{s}, err
}

type HoldsVerTwoDataList struct{ capnp.Struct }

// HoldsVerTwoDataList_TypeID is the unique identifier for the type HoldsVerTwoDataList.
const HoldsVerTwoDataList_TypeID = 0xcbdc765fd5dff7ba

func NewHoldsVerTwoDataList(s *capnp.Segment) (HoldsVerTwoDataList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoDataList{st}, err
}

func NewRootHoldsVerTwoDataList(s *capnp.Segment) (HoldsVerTwoDataList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoDataList{st}, err
}

func ReadRootHoldsVerTwoDataList(msg *capnp.Message) (HoldsVerTwoDataList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoDataList{root.Struct()}, err
}

func (s HoldsVerTwoDataList) String() string {
	str, _ := text.Marshal(0xcbdc765fd5dff7ba, s.Struct)
	return str
}

func (s HoldsVerTwoDataList) Copy(seg *capnp.Segment) (HoldsVerTwoDataList, error) {
	var err error
	t, err := NewHoldsVerTwoDataList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerTwoDataList) Mylist() (VerTwoData_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoData_List{List: p.List()}, err
}

func (s HoldsVerTwoDataList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoDataList) SetMylist(v VerTwoData_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoData_List, preferring placement in s's segment.
func (s HoldsVerTwoDataList) NewMylist(n int32) (VerTwoData_List, error) {
	l, err := NewVerTwoData_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoData_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoDataList_List is a list of HoldsVerTwoDataList.
type HoldsVerTwoDataList_List struct{ capnp.List }

// NewHoldsVerTwoDataList creates a new list of HoldsVerTwoDataList.
func NewHoldsVerTwoDataList_List(s *capnp.Segment, sz int32) (HoldsVerTwoDataList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoDataList_List{l}, err
}

func (s HoldsVerTwoDataList_List) At(i int) HoldsVerTwoDataList {
	return HoldsVerTwoDataList{s.List.Struct(i)}
}

func (s HoldsVerTwoDataList_List) Set(i int, v HoldsVerTwoDataList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerTwoDataList_List) Copy(seg *capnp.Segment) (HoldsVerTwoDataList_List, error) {
	var err error
	t, err := NewHoldsVerTwoDataList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerTwoDataList_Promise is a wrapper for a HoldsVerTwoDataList promised by a client call.
type HoldsVerTwoDataList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoDataList_Promise) Struct() (HoldsVerTwoDataList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoDataList{s}, err
}

type HoldsVerOnePtrList struct{ capnp.Struct }

// HoldsVerOnePtrList_TypeID is the unique identifier for the type HoldsVerOnePtrList.
const HoldsVerOnePtrList_TypeID = 0xe508a29c83a059f8

func NewHoldsVerOnePtrList(s *capnp.Segment) (HoldsVerOnePtrList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOnePtrList{st}, err
}

func NewRootHoldsVerOnePtrList(s *capnp.Segment) (HoldsVerOnePtrList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerOnePtrList{st}, err
}

func ReadRootHoldsVerOnePtrList(msg *capnp.Message) (HoldsVerOnePtrList, error) {
	root, err := msg.RootPtr()
	return HoldsVerOnePtrList{root.Struct()}, err
}

func (s HoldsVerOnePtrList) String() string {
	str, _ := text.Marshal(0xe508a29c83a059f8, s.Struct)
	return str
}

func (s HoldsVerOnePtrList) Copy(seg *capnp.Segment) (HoldsVerOnePtrList, error) {
	var err error
	t, err := NewHoldsVerOnePtrList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerOnePtrList) Mylist() (VerOnePtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerOnePtr_List{List: p.List()}, err
}

func (s HoldsVerOnePtrList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerOnePtrList) SetMylist(v VerOnePtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerOnePtr_List, preferring placement in s's segment.
func (s HoldsVerOnePtrList) NewMylist(n int32) (VerOnePtr_List, error) {
	l, err := NewVerOnePtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerOnePtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerOnePtrList_List is a list of HoldsVerOnePtrList.
type HoldsVerOnePtrList_List struct{ capnp.List }

// NewHoldsVerOnePtrList creates a new list of HoldsVerOnePtrList.
func NewHoldsVerOnePtrList_List(s *capnp.Segment, sz int32) (HoldsVerOnePtrList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerOnePtrList_List{l}, err
}

func (s HoldsVerOnePtrList_List) At(i int) HoldsVerOnePtrList {
	return HoldsVerOnePtrList{s.List.Struct(i)}
}

func (s HoldsVerOnePtrList_List) Set(i int, v HoldsVerOnePtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerOnePtrList_List) Copy(seg *capnp.Segment) (HoldsVerOnePtrList_List, error) {
	var err error
	t, err := NewHoldsVerOnePtrList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerOnePtrList_Promise is a wrapper for a HoldsVerOnePtrList promised by a client call.
type HoldsVerOnePtrList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerOnePtrList_Promise) Struct() (HoldsVerOnePtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerOnePtrList{s}, err
}

type HoldsVerTwoPtrList struct{ capnp.Struct }

// HoldsVerTwoPtrList_TypeID is the unique identifier for the type HoldsVerTwoPtrList.
const HoldsVerTwoPtrList_TypeID = 0xcf9beaca1cc180c8

func NewHoldsVerTwoPtrList(s *capnp.Segment) (HoldsVerTwoPtrList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoPtrList{st}, err
}

func NewRootHoldsVerTwoPtrList(s *capnp.Segment) (HoldsVerTwoPtrList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoPtrList{st}, err
}

func ReadRootHoldsVerTwoPtrList(msg *capnp.Message) (HoldsVerTwoPtrList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoPtrList{root.Struct()}, err
}

func (s HoldsVerTwoPtrList) String() string {
	str, _ := text.Marshal(0xcf9beaca1cc180c8, s.Struct)
	return str
}

func (s HoldsVerTwoPtrList) Copy(seg *capnp.Segment) (HoldsVerTwoPtrList, error) {
	var err error
	t, err := NewHoldsVerTwoPtrList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerTwoPtrList) Mylist() (VerTwoPtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoPtr_List{List: p.List()}, err
}

func (s HoldsVerTwoPtrList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoPtrList) SetMylist(v VerTwoPtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoPtr_List, preferring placement in s's segment.
func (s HoldsVerTwoPtrList) NewMylist(n int32) (VerTwoPtr_List, error) {
	l, err := NewVerTwoPtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoPtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoPtrList_List is a list of HoldsVerTwoPtrList.
type HoldsVerTwoPtrList_List struct{ capnp.List }

// NewHoldsVerTwoPtrList creates a new list of HoldsVerTwoPtrList.
func NewHoldsVerTwoPtrList_List(s *capnp.Segment, sz int32) (HoldsVerTwoPtrList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoPtrList_List{l}, err
}

func (s HoldsVerTwoPtrList_List) At(i int) HoldsVerTwoPtrList {
	return HoldsVerTwoPtrList{s.List.Struct(i)}
}

func (s HoldsVerTwoPtrList_List) Set(i int, v HoldsVerTwoPtrList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerTwoPtrList_List) Copy(seg *capnp.Segment) (HoldsVerTwoPtrList_List, error) {
	var err error
	t, err := NewHoldsVerTwoPtrList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerTwoPtrList_Promise is a wrapper for a HoldsVerTwoPtrList promised by a client call.
type HoldsVerTwoPtrList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoPtrList_Promise) Struct() (HoldsVerTwoPtrList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoPtrList{s}, err
}

type HoldsVerTwoTwoList struct{ capnp.Struct }

// HoldsVerTwoTwoList_TypeID is the unique identifier for the type HoldsVerTwoTwoList.
const HoldsVerTwoTwoList_TypeID = 0x95befe3f14606e6b

func NewHoldsVerTwoTwoList(s *capnp.Segment) (HoldsVerTwoTwoList, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoList{st}, err
}

func NewRootHoldsVerTwoTwoList(s *capnp.Segment) (HoldsVerTwoTwoList, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoList{st}, err
}

func ReadRootHoldsVerTwoTwoList(msg *capnp.Message) (HoldsVerTwoTwoList, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoTwoList{root.Struct()}, err
}

func (s HoldsVerTwoTwoList) String() string {
	str, _ := text.Marshal(0x95befe3f14606e6b, s.Struct)
	return str
}

func (s HoldsVerTwoTwoList) Copy(seg *capnp.Segment) (HoldsVerTwoTwoList, error) {
	var err error
	t, err := NewHoldsVerTwoTwoList(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerTwoTwoList) Mylist() (VerTwoDataTwoPtr_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr_List{List: p.List()}, err
}

func (s HoldsVerTwoTwoList) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoTwoList) SetMylist(v VerTwoDataTwoPtr_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoDataTwoPtr_List, preferring placement in s's segment.
func (s HoldsVerTwoTwoList) NewMylist(n int32) (VerTwoDataTwoPtr_List, error) {
	l, err := NewVerTwoDataTwoPtr_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoDataTwoPtr_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoTwoList_List is a list of HoldsVerTwoTwoList.
type HoldsVerTwoTwoList_List struct{ capnp.List }

// NewHoldsVerTwoTwoList creates a new list of HoldsVerTwoTwoList.
func NewHoldsVerTwoTwoList_List(s *capnp.Segment, sz int32) (HoldsVerTwoTwoList_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoTwoList_List{l}, err
}

func (s HoldsVerTwoTwoList_List) At(i int) HoldsVerTwoTwoList {
	return HoldsVerTwoTwoList{s.List.Struct(i)}
}

func (s HoldsVerTwoTwoList_List) Set(i int, v HoldsVerTwoTwoList) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerTwoTwoList_List) Copy(seg *capnp.Segment) (HoldsVerTwoTwoList_List, error) {
	var err error
	t, err := NewHoldsVerTwoTwoList_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerTwoTwoList_Promise is a wrapper for a HoldsVerTwoTwoList promised by a client call.
type HoldsVerTwoTwoList_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoTwoList_Promise) Struct() (HoldsVerTwoTwoList, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoList{s}, err
}

type HoldsVerTwoTwoPlus struct{ capnp.Struct }

// HoldsVerTwoTwoPlus_TypeID is the unique identifier for the type HoldsVerTwoTwoPlus.
const HoldsVerTwoTwoPlus_TypeID = 0x87c33f2330feb3d8

func NewHoldsVerTwoTwoPlus(s *capnp.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoPlus{st}, err
}

func NewRootHoldsVerTwoTwoPlus(s *capnp.Segment) (HoldsVerTwoTwoPlus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return HoldsVerTwoTwoPlus{st}, err
}

func ReadRootHoldsVerTwoTwoPlus(msg *capnp.Message) (HoldsVerTwoTwoPlus, error) {
	root, err := msg.RootPtr()
	return HoldsVerTwoTwoPlus{root.Struct()}, err
}

func (s HoldsVerTwoTwoPlus) String() string {
	str, _ := text.Marshal(0x87c33f2330feb3d8, s.Struct)
	return str
}

func (s HoldsVerTwoTwoPlus) Copy(seg *capnp.Segment) (HoldsVerTwoTwoPlus, error) {
	var err error
	t, err := NewHoldsVerTwoTwoPlus(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Mylist()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMylist(v)
	}
	return t, nil
}

func (s HoldsVerTwoTwoPlus) Mylist() (VerTwoTwoPlus_List, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoTwoPlus_List{List: p.List()}, err
}

func (s HoldsVerTwoTwoPlus) HasMylist() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsVerTwoTwoPlus) SetMylist(v VerTwoTwoPlus_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMylist sets the mylist field to a newly
// allocated VerTwoTwoPlus_List, preferring placement in s's segment.
func (s HoldsVerTwoTwoPlus) NewMylist(n int32) (VerTwoTwoPlus_List, error) {
	l, err := NewVerTwoTwoPlus_List(s.Struct.Segment(), n)
	if err != nil {
		return VerTwoTwoPlus_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// HoldsVerTwoTwoPlus_List is a list of HoldsVerTwoTwoPlus.
type HoldsVerTwoTwoPlus_List struct{ capnp.List }

// NewHoldsVerTwoTwoPlus creates a new list of HoldsVerTwoTwoPlus.
func NewHoldsVerTwoTwoPlus_List(s *capnp.Segment, sz int32) (HoldsVerTwoTwoPlus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return HoldsVerTwoTwoPlus_List{l}, err
}

func (s HoldsVerTwoTwoPlus_List) At(i int) HoldsVerTwoTwoPlus {
	return HoldsVerTwoTwoPlus{s.List.Struct(i)}
}

func (s HoldsVerTwoTwoPlus_List) Set(i int, v HoldsVerTwoTwoPlus) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s HoldsVerTwoTwoPlus_List) Copy(seg *capnp.Segment) (HoldsVerTwoTwoPlus_List, error) {
	var err error
	t, err := NewHoldsVerTwoTwoPlus_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsVerTwoTwoPlus_Promise is a wrapper for a HoldsVerTwoTwoPlus promised by a client call.
type HoldsVerTwoTwoPlus_Promise struct{ *capnp.Pipeline }

func (p HoldsVerTwoTwoPlus_Promise) Struct() (HoldsVerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return HoldsVerTwoTwoPlus{s}, err
}

type VerTwoTwoPlus struct{ capnp.Struct }

// VerTwoTwoPlus_TypeID is the unique identifier for the type VerTwoTwoPlus.
const VerTwoTwoPlus_TypeID = 0xce44aee2d9e25049

func NewVerTwoTwoPlus(s *capnp.Segment) (VerTwoTwoPlus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return VerTwoTwoPlus{st}, err
}

func NewRootVerTwoTwoPlus(s *capnp.Segment) (VerTwoTwoPlus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return VerTwoTwoPlus{st}, err
}

func ReadRootVerTwoTwoPlus(msg *capnp.Message) (VerTwoTwoPlus, error) {
	root, err := msg.RootPtr()
	return VerTwoTwoPlus{root.Struct()}, err
}

func (s VerTwoTwoPlus) String() string {
	str, _ := text.Marshal(0xce44aee2d9e25049, s.Struct)
	return str
}

func (s VerTwoTwoPlus) Copy(seg *capnp.Segment) (VerTwoTwoPlus, error) {
	var err error
	t, err := NewVerTwoTwoPlus(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Ptr1()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr1(v)
	}
	{
		o, err := s.Ptr2()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetPtr2(v)
	}
	{
		o, err := s.Lst3()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetLst3(v)
	}
	t.SetVal(s.Val())
	t.SetDuo(s.Duo())
	t.SetTre(s.Tre())
	return t, nil
}

func (s VerTwoTwoPlus) Val() int16 {
	return int16(s.Struct.Uint16(0))
}

func (s VerTwoTwoPlus) SetVal(v int16) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s VerTwoTwoPlus) Duo() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s VerTwoTwoPlus) SetDuo(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s VerTwoTwoPlus) Ptr1() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s VerTwoTwoPlus) HasPtr1() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetPtr1(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPtr1 sets the ptr1 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr1() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoTwoPlus) Ptr2() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(1)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s VerTwoTwoPlus) HasPtr2() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetPtr2(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPtr2 sets the ptr2 field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewPtr2() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s VerTwoTwoPlus) Tre() int64 {
	return int64(s.Struct.Uint64(16))
}

func (s VerTwoTwoPlus) SetTre(v int64) {
	s.Struct.SetUint64(16, uint64(v))
}

func (s VerTwoTwoPlus) Lst3() (capnp.Int64List, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.Int64List{List: p.List()}, err
}

func (s VerTwoTwoPlus) HasLst3() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s VerTwoTwoPlus) SetLst3(v capnp.Int64List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLst3 sets the lst3 field to a newly
// allocated capnp.Int64List, preferring placement in s's segment.
func (s VerTwoTwoPlus) NewLst3(n int32) (capnp.Int64List, error) {
	l, err := capnp.NewInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Int64List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// VerTwoTwoPlus_List is a list of VerTwoTwoPlus.
type VerTwoTwoPlus_List struct{ capnp.List }

// NewVerTwoTwoPlus creates a new list of VerTwoTwoPlus.
func NewVerTwoTwoPlus_List(s *capnp.Segment, sz int32) (VerTwoTwoPlus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return VerTwoTwoPlus_List{l}, err
}

func (s VerTwoTwoPlus_List) At(i int) VerTwoTwoPlus { return VerTwoTwoPlus{s.List.Struct(i)} }

func (s VerTwoTwoPlus_List) Set(i int, v VerTwoTwoPlus) error { return s.List.SetStruct(i, v.Struct) }

func (s VerTwoTwoPlus_List) Copy(seg *capnp.Segment) (VerTwoTwoPlus_List, error) {
	var err error
	t, err := NewVerTwoTwoPlus_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VerTwoTwoPlus_Promise is a wrapper for a VerTwoTwoPlus promised by a client call.
type VerTwoTwoPlus_Promise struct{ *capnp.Pipeline }

func (p VerTwoTwoPlus_Promise) Struct() (VerTwoTwoPlus, error) {
	s, err := p.Pipeline.Struct()
	return VerTwoTwoPlus{s}, err
}

func (p VerTwoTwoPlus_Promise) Ptr1() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p VerTwoTwoPlus_Promise) Ptr2() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type HoldsText struct{ capnp.Struct }

// HoldsText_TypeID is the unique identifier for the type HoldsText.
const HoldsText_TypeID = 0xe5817f849ff906dc

func NewHoldsText(s *capnp.Segment) (HoldsText, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HoldsText{st}, err
}

func NewRootHoldsText(s *capnp.Segment) (HoldsText, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return HoldsText{st}, err
}

func ReadRootHoldsText(msg *capnp.Message) (HoldsText, error) {
	root, err := msg.RootPtr()
	return HoldsText{root.Struct()}, err
}

func (s HoldsText) String() string {
	str, _ := text.Marshal(0xe5817f849ff906dc, s.Struct)
	return str
}

func (s HoldsText) Copy(seg *capnp.Segment) (HoldsText, error) {
	var err error
	t, err := NewHoldsText(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Lst()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetLst(v)
	}
	{
		o, err := s.Lstlst()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetLstlst(v)
	}
	{
		v, err := s.Txt()
		if err != nil {
			return t, err
		}
		t.SetTxt(v)
	}
	return t, nil
}

func (s HoldsText) Txt() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HoldsText) HasTxt() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HoldsText) TxtBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HoldsText) SetTxt(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s HoldsText) Lst() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s HoldsText) HasLst() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s HoldsText) SetLst(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewLst sets the lst field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s HoldsText) NewLst(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s HoldsText) Lstlst() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.PointerList{List: p.List()}, err
}

func (s HoldsText) HasLstlst() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s HoldsText) SetLstlst(v capnp.PointerList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLstlst sets the lstlst field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s HoldsText) NewLstlst(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// HoldsText_List is a list of HoldsText.
type HoldsText_List struct{ capnp.List }

// NewHoldsText creates a new list of HoldsText.
func NewHoldsText_List(s *capnp.Segment, sz int32) (HoldsText_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return HoldsText_List{l}, err
}

func (s HoldsText_List) At(i int) HoldsText { return HoldsText{s.List.Struct(i)} }

func (s HoldsText_List) Set(i int, v HoldsText) error { return s.List.SetStruct(i, v.Struct) }

func (s HoldsText_List) Copy(seg *capnp.Segment) (HoldsText_List, error) {
	var err error
	t, err := NewHoldsText_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// HoldsText_Promise is a wrapper for a HoldsText promised by a client call.
type HoldsText_Promise struct{ *capnp.Pipeline }

func (p HoldsText_Promise) Struct() (HoldsText, error) {
	s, err := p.Pipeline.Struct()
	return HoldsText{s}, err
}

type WrapEmpty struct{ capnp.Struct }

// WrapEmpty_TypeID is the unique identifier for the type WrapEmpty.
const WrapEmpty_TypeID = 0x9ab599979b02ac59

func NewWrapEmpty(s *capnp.Segment) (WrapEmpty, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WrapEmpty{st}, err
}

func NewRootWrapEmpty(s *capnp.Segment) (WrapEmpty, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return WrapEmpty{st}, err
}

func ReadRootWrapEmpty(msg *capnp.Message) (WrapEmpty, error) {
	root, err := msg.RootPtr()
	return WrapEmpty{root.Struct()}, err
}

func (s WrapEmpty) String() string {
	str, _ := text.Marshal(0x9ab599979b02ac59, s.Struct)
	return str
}

func (s WrapEmpty) Copy(seg *capnp.Segment) (WrapEmpty, error) {
	var err error
	t, err := NewWrapEmpty(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.MightNotBeReallyEmpty()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMightNotBeReallyEmpty(v)
	}
	return t, nil
}

func (s WrapEmpty) MightNotBeReallyEmpty() (VerEmpty, error) {
	p, err := s.Struct.Ptr(0)
	return VerEmpty{Struct: p.Struct()}, err
}

func (s WrapEmpty) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s WrapEmpty) SetMightNotBeReallyEmpty(v VerEmpty) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerEmpty struct, preferring placement in s's segment.
func (s WrapEmpty) NewMightNotBeReallyEmpty() (VerEmpty, error) {
	ss, err := NewVerEmpty(s.Struct.Segment())
	if err != nil {
		return VerEmpty{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// WrapEmpty_List is a list of WrapEmpty.
type WrapEmpty_List struct{ capnp.List }

// NewWrapEmpty creates a new list of WrapEmpty.
func NewWrapEmpty_List(s *capnp.Segment, sz int32) (WrapEmpty_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return WrapEmpty_List{l}, err
}

func (s WrapEmpty_List) At(i int) WrapEmpty { return WrapEmpty{s.List.Struct(i)} }

func (s WrapEmpty_List) Set(i int, v WrapEmpty) error { return s.List.SetStruct(i, v.Struct) }

func (s WrapEmpty_List) Copy(seg *capnp.Segment) (WrapEmpty_List, error) {
	var err error
	t, err := NewWrapEmpty_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// WrapEmpty_Promise is a wrapper for a WrapEmpty promised by a client call.
type WrapEmpty_Promise struct{ *capnp.Pipeline }

func (p WrapEmpty_Promise) Struct() (WrapEmpty, error) {
	s, err := p.Pipeline.Struct()
	return WrapEmpty{s}, err
}

func (p WrapEmpty_Promise) MightNotBeReallyEmpty() VerEmpty_Promise {
	return VerEmpty_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2 struct{ capnp.Struct }

// Wrap2x2_TypeID is the unique identifier for the type Wrap2x2.
const Wrap2x2_TypeID = 0xe1a2d1d51107bead

func NewWrap2x2(s *capnp.Segment) (Wrap2x2, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2{st}, err
}

func NewRootWrap2x2(s *capnp.Segment) (Wrap2x2, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2{st}, err
}

func ReadRootWrap2x2(msg *capnp.Message) (Wrap2x2, error) {
	root, err := msg.RootPtr()
	return Wrap2x2{root.Struct()}, err
}

func (s Wrap2x2) String() string {
	str, _ := text.Marshal(0xe1a2d1d51107bead, s.Struct)
	return str
}

func (s Wrap2x2) Copy(seg *capnp.Segment) (Wrap2x2, error) {
	var err error
	t, err := NewWrap2x2(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.MightNotBeReallyEmpty()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMightNotBeReallyEmpty(v)
	}
	return t, nil
}

func (s Wrap2x2) MightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoDataTwoPtr{Struct: p.Struct()}, err
}

func (s Wrap2x2) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Wrap2x2) SetMightNotBeReallyEmpty(v VerTwoDataTwoPtr) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoDataTwoPtr struct, preferring placement in s's segment.
func (s Wrap2x2) NewMightNotBeReallyEmpty() (VerTwoDataTwoPtr, error) {
	ss, err := NewVerTwoDataTwoPtr(s.Struct.Segment())
	if err != nil {
		return VerTwoDataTwoPtr{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Wrap2x2_List is a list of Wrap2x2.
type Wrap2x2_List struct{ capnp.List }

// NewWrap2x2 creates a new list of Wrap2x2.
func NewWrap2x2_List(s *capnp.Segment, sz int32) (Wrap2x2_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Wrap2x2_List{l}, err
}

func (s Wrap2x2_List) At(i int) Wrap2x2 { return Wrap2x2{s.List.Struct(i)} }

func (s Wrap2x2_List) Set(i int, v Wrap2x2) error { return s.List.SetStruct(i, v.Struct) }

func (s Wrap2x2_List) Copy(seg *capnp.Segment) (Wrap2x2_List, error) {
	var err error
	t, err := NewWrap2x2_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Wrap2x2_Promise is a wrapper for a Wrap2x2 promised by a client call.
type Wrap2x2_Promise struct{ *capnp.Pipeline }

func (p Wrap2x2_Promise) Struct() (Wrap2x2, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2{s}, err
}

func (p Wrap2x2_Promise) MightNotBeReallyEmpty() VerTwoDataTwoPtr_Promise {
	return VerTwoDataTwoPtr_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Wrap2x2plus struct{ capnp.Struct }

// Wrap2x2plus_TypeID is the unique identifier for the type Wrap2x2plus.
const Wrap2x2plus_TypeID = 0xe684eb3aef1a6859

func NewWrap2x2plus(s *capnp.Segment) (Wrap2x2plus, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2plus{st}, err
}

func NewRootWrap2x2plus(s *capnp.Segment) (Wrap2x2plus, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Wrap2x2plus{st}, err
}

func ReadRootWrap2x2plus(msg *capnp.Message) (Wrap2x2plus, error) {
	root, err := msg.RootPtr()
	return Wrap2x2plus{root.Struct()}, err
}

func (s Wrap2x2plus) String() string {
	str, _ := text.Marshal(0xe684eb3aef1a6859, s.Struct)
	return str
}

func (s Wrap2x2plus) Copy(seg *capnp.Segment) (Wrap2x2plus, error) {
	var err error
	t, err := NewWrap2x2plus(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.MightNotBeReallyEmpty()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetMightNotBeReallyEmpty(v)
	}
	return t, nil
}

func (s Wrap2x2plus) MightNotBeReallyEmpty() (VerTwoTwoPlus, error) {
	p, err := s.Struct.Ptr(0)
	return VerTwoTwoPlus{Struct: p.Struct()}, err
}

func (s Wrap2x2plus) HasMightNotBeReallyEmpty() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Wrap2x2plus) SetMightNotBeReallyEmpty(v VerTwoTwoPlus) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewMightNotBeReallyEmpty sets the mightNotBeReallyEmpty field to a newly
// allocated VerTwoTwoPlus struct, preferring placement in s's segment.
func (s Wrap2x2plus) NewMightNotBeReallyEmpty() (VerTwoTwoPlus, error) {
	ss, err := NewVerTwoTwoPlus(s.Struct.Segment())
	if err != nil {
		return VerTwoTwoPlus{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Wrap2x2plus_List is a list of Wrap2x2plus.
type Wrap2x2plus_List struct{ capnp.List }

// NewWrap2x2plus creates a new list of Wrap2x2plus.
func NewWrap2x2plus_List(s *capnp.Segment, sz int32) (Wrap2x2plus_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Wrap2x2plus_List{l}, err
}

func (s Wrap2x2plus_List) At(i int) Wrap2x2plus { return Wrap2x2plus{s.List.Struct(i)} }

func (s Wrap2x2plus_List) Set(i int, v Wrap2x2plus) error { return s.List.SetStruct(i, v.Struct) }

func (s Wrap2x2plus_List) Copy(seg *capnp.Segment) (Wrap2x2plus_List, error) {
	var err error
	t, err := NewWrap2x2plus_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Wrap2x2plus_Promise is a wrapper for a Wrap2x2plus promised by a client call.
type Wrap2x2plus_Promise struct{ *capnp.Pipeline }

func (p Wrap2x2plus_Promise) Struct() (Wrap2x2plus, error) {
	s, err := p.Pipeline.Struct()
	return Wrap2x2plus{s}, err
}

func (p Wrap2x2plus_Promise) MightNotBeReallyEmpty() VerTwoTwoPlus_Promise {
	return VerTwoTwoPlus_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type VoidUnion struct{ capnp.Struct }
type VoidUnion_Which uint16

const (
	VoidUnion_Which_a VoidUnion_Which = 0
	VoidUnion_Which_b VoidUnion_Which = 1
)

func (w VoidUnion_Which) String() string {
	const s = "ab"
	switch w {
	case VoidUnion_Which_a:
		return s[0:1]
	case VoidUnion_Which_b:
		return s[1:2]

	}
	return "VoidUnion_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// VoidUnion_TypeID is the unique identifier for the type VoidUnion.
const VoidUnion_TypeID = 0x8821cdb23640783a

func NewVoidUnion(s *capnp.Segment) (VoidUnion, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VoidUnion{st}, err
}

func NewRootVoidUnion(s *capnp.Segment) (VoidUnion, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return VoidUnion{st}, err
}

func ReadRootVoidUnion(msg *capnp.Message) (VoidUnion, error) {
	root, err := msg.RootPtr()
	return VoidUnion{root.Struct()}, err
}

func (s VoidUnion) String() string {
	str, _ := text.Marshal(0x8821cdb23640783a, s.Struct)
	return str
}

func (s VoidUnion) Copy(seg *capnp.Segment) (VoidUnion, error) {
	var err error
	t, err := NewVoidUnion(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s VoidUnion) Which() VoidUnion_Which {
	return VoidUnion_Which(s.Struct.Uint16(0))
}
func (s VoidUnion) SetA() {
	s.Struct.SetUint16(0, 0)

}

func (s VoidUnion) SetB() {
	s.Struct.SetUint16(0, 1)

}

// VoidUnion_List is a list of VoidUnion.
type VoidUnion_List struct{ capnp.List }

// NewVoidUnion creates a new list of VoidUnion.
func NewVoidUnion_List(s *capnp.Segment, sz int32) (VoidUnion_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return VoidUnion_List{l}, err
}

func (s VoidUnion_List) At(i int) VoidUnion { return VoidUnion{s.List.Struct(i)} }

func (s VoidUnion_List) Set(i int, v VoidUnion) error { return s.List.SetStruct(i, v.Struct) }

func (s VoidUnion_List) Copy(seg *capnp.Segment) (VoidUnion_List, error) {
	var err error
	t, err := NewVoidUnion_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// VoidUnion_Promise is a wrapper for a VoidUnion promised by a client call.
type VoidUnion_Promise struct{ *capnp.Pipeline }

func (p VoidUnion_Promise) Struct() (VoidUnion, error) {
	s, err := p.Pipeline.Struct()
	return VoidUnion{s}, err
}

type Nester1Capn struct{ capnp.Struct }

// Nester1Capn_TypeID is the unique identifier for the type Nester1Capn.
const Nester1Capn_TypeID = 0xf14fad09425d081c

func NewNester1Capn(s *capnp.Segment) (Nester1Capn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Nester1Capn{st}, err
}

func NewRootNester1Capn(s *capnp.Segment) (Nester1Capn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Nester1Capn{st}, err
}

func ReadRootNester1Capn(msg *capnp.Message) (Nester1Capn, error) {
	root, err := msg.RootPtr()
	return Nester1Capn{root.Struct()}, err
}

func (s Nester1Capn) String() string {
	str, _ := text.Marshal(0xf14fad09425d081c, s.Struct)
	return str
}

func (s Nester1Capn) Copy(seg *capnp.Segment) (Nester1Capn, error) {
	var err error
	t, err := NewNester1Capn(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Strs()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetStrs(v)
	}
	return t, nil
}

func (s Nester1Capn) Strs() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.TextList{List: p.List()}, err
}

func (s Nester1Capn) HasStrs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Nester1Capn) SetStrs(v capnp.TextList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewStrs sets the strs field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Nester1Capn) NewStrs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Nester1Capn_List is a list of Nester1Capn.
type Nester1Capn_List struct{ capnp.List }

// NewNester1Capn creates a new list of Nester1Capn.
func NewNester1Capn_List(s *capnp.Segment, sz int32) (Nester1Capn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Nester1Capn_List{l}, err
}

func (s Nester1Capn_List) At(i int) Nester1Capn { return Nester1Capn{s.List.Struct(i)} }

func (s Nester1Capn_List) Set(i int, v Nester1Capn) error { return s.List.SetStruct(i, v.Struct) }

func (s Nester1Capn_List) Copy(seg *capnp.Segment) (Nester1Capn_List, error) {
	var err error
	t, err := NewNester1Capn_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Nester1Capn_Promise is a wrapper for a Nester1Capn promised by a client call.
type Nester1Capn_Promise struct{ *capnp.Pipeline }

func (p Nester1Capn_Promise) Struct() (Nester1Capn, error) {
	s, err := p.Pipeline.Struct()
	return Nester1Capn{s}, err
}

type RWTestCapn struct{ capnp.Struct }

// RWTestCapn_TypeID is the unique identifier for the type RWTestCapn.
const RWTestCapn_TypeID = 0xf7ff4414476c186a

func NewRWTestCapn(s *capnp.Segment) (RWTestCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RWTestCapn{st}, err
}

func NewRootRWTestCapn(s *capnp.Segment) (RWTestCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return RWTestCapn{st}, err
}

func ReadRootRWTestCapn(msg *capnp.Message) (RWTestCapn, error) {
	root, err := msg.RootPtr()
	return RWTestCapn{root.Struct()}, err
}

func (s RWTestCapn) String() string {
	str, _ := text.Marshal(0xf7ff4414476c186a, s.Struct)
	return str
}

func (s RWTestCapn) Copy(seg *capnp.Segment) (RWTestCapn, error) {
	var err error
	t, err := NewRWTestCapn(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.NestMatrix()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetNestMatrix(v)
	}
	return t, nil
}

func (s RWTestCapn) NestMatrix() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.PointerList{List: p.List()}, err
}

func (s RWTestCapn) HasNestMatrix() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RWTestCapn) SetNestMatrix(v capnp.PointerList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewNestMatrix sets the nestMatrix field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s RWTestCapn) NewNestMatrix(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// RWTestCapn_List is a list of RWTestCapn.
type RWTestCapn_List struct{ capnp.List }

// NewRWTestCapn creates a new list of RWTestCapn.
func NewRWTestCapn_List(s *capnp.Segment, sz int32) (RWTestCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return RWTestCapn_List{l}, err
}

func (s RWTestCapn_List) At(i int) RWTestCapn { return RWTestCapn{s.List.Struct(i)} }

func (s RWTestCapn_List) Set(i int, v RWTestCapn) error { return s.List.SetStruct(i, v.Struct) }

func (s RWTestCapn_List) Copy(seg *capnp.Segment) (RWTestCapn_List, error) {
	var err error
	t, err := NewRWTestCapn_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// RWTestCapn_Promise is a wrapper for a RWTestCapn promised by a client call.
type RWTestCapn_Promise struct{ *capnp.Pipeline }

func (p RWTestCapn_Promise) Struct() (RWTestCapn, error) {
	s, err := p.Pipeline.Struct()
	return RWTestCapn{s}, err
}

type ListStructCapn struct{ capnp.Struct }

// ListStructCapn_TypeID is the unique identifier for the type ListStructCapn.
const ListStructCapn_TypeID = 0xb1ac056ed7647011

func NewListStructCapn(s *capnp.Segment) (ListStructCapn, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ListStructCapn{st}, err
}

func NewRootListStructCapn(s *capnp.Segment) (ListStructCapn, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ListStructCapn{st}, err
}

func ReadRootListStructCapn(msg *capnp.Message) (ListStructCapn, error) {
	root, err := msg.RootPtr()
	return ListStructCapn{root.Struct()}, err
}

func (s ListStructCapn) String() string {
	str, _ := text.Marshal(0xb1ac056ed7647011, s.Struct)
	return str
}

func (s ListStructCapn) Copy(seg *capnp.Segment) (ListStructCapn, error) {
	var err error
	t, err := NewListStructCapn(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Vec()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetVec(v)
	}
	return t, nil
}

func (s ListStructCapn) Vec() (Nester1Capn_List, error) {
	p, err := s.Struct.Ptr(0)
	return Nester1Capn_List{List: p.List()}, err
}

func (s ListStructCapn) HasVec() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ListStructCapn) SetVec(v Nester1Capn_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewVec sets the vec field to a newly
// allocated Nester1Capn_List, preferring placement in s's segment.
func (s ListStructCapn) NewVec(n int32) (Nester1Capn_List, error) {
	l, err := NewNester1Capn_List(s.Struct.Segment(), n)
	if err != nil {
		return Nester1Capn_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// ListStructCapn_List is a list of ListStructCapn.
type ListStructCapn_List struct{ capnp.List }

// NewListStructCapn creates a new list of ListStructCapn.
func NewListStructCapn_List(s *capnp.Segment, sz int32) (ListStructCapn_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ListStructCapn_List{l}, err
}

func (s ListStructCapn_List) At(i int) ListStructCapn { return ListStructCapn{s.List.Struct(i)} }

func (s ListStructCapn_List) Set(i int, v ListStructCapn) error { return s.List.SetStruct(i, v.Struct) }

func (s ListStructCapn_List) Copy(seg *capnp.Segment) (ListStructCapn_List, error) {
	var err error
	t, err := NewListStructCapn_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// ListStructCapn_Promise is a wrapper for a ListStructCapn promised by a client call.
type ListStructCapn_Promise struct{ *capnp.Pipeline }

func (p ListStructCapn_Promise) Struct() (ListStructCapn, error) {
	s, err := p.Pipeline.Struct()
	return ListStructCapn{s}, err
}

type Echo struct{ Client capnp.Client }

func (c Echo) Echo(ctx context.Context, params func(Echo_echo_Params) error, opts ...capnp.CallOption) Echo_echo_Results_Promise {
	if c.Client == nil {
		return Echo_echo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Echo_echo_Params{Struct: s}) }
	}
	return Echo_echo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Echo_Server interface {
	Echo(Echo_echo) error
}

func Echo_ServerToClient(s Echo_Server) Echo {
	c, _ := s.(server.Closer)
	return Echo{Client: server.New(Echo_Methods(nil, s), c)}
}

func Echo_Methods(methods []server.Method, s Echo_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x8e5322c1e9282534,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:Echo",
			MethodName:    "echo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Echo_echo{c, opts, Echo_echo_Params{Struct: p}, Echo_echo_Results{Struct: r}}
			return s.Echo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Echo_echo holds the arguments for a server call to Echo.echo.
type Echo_echo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Echo_echo_Params
	Results Echo_echo_Results
}

type Echo_echo_Params struct{ capnp.Struct }

// Echo_echo_Params_TypeID is the unique identifier for the type Echo_echo_Params.
const Echo_echo_Params_TypeID = 0x8a165fb4d71bf3a2

func NewEcho_echo_Params(s *capnp.Segment) (Echo_echo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Params{st}, err
}

func NewRootEcho_echo_Params(s *capnp.Segment) (Echo_echo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Params{st}, err
}

func ReadRootEcho_echo_Params(msg *capnp.Message) (Echo_echo_Params, error) {
	root, err := msg.RootPtr()
	return Echo_echo_Params{root.Struct()}, err
}

func (s Echo_echo_Params) String() string {
	str, _ := text.Marshal(0x8a165fb4d71bf3a2, s.Struct)
	return str
}

func (s Echo_echo_Params) Copy(seg *capnp.Segment) (Echo_echo_Params, error) {
	var err error
	t, err := NewEcho_echo_Params(seg)
	if err != nil {
		return t, err
	}
	{
		v, err := s.In()
		if err != nil {
			return t, err
		}
		t.SetIn(v)
	}
	return t, nil
}

func (s Echo_echo_Params) In() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Echo_echo_Params) HasIn() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Echo_echo_Params) InBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Echo_echo_Params) SetIn(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Echo_echo_Params_List is a list of Echo_echo_Params.
type Echo_echo_Params_List struct{ capnp.List }

// NewEcho_echo_Params creates a new list of Echo_echo_Params.
func NewEcho_echo_Params_List(s *capnp.Segment, sz int32) (Echo_echo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Echo_echo_Params_List{l}, err
}

func (s Echo_echo_Params_List) At(i int) Echo_echo_Params { return Echo_echo_Params{s.List.Struct(i)} }

func (s Echo_echo_Params_List) Set(i int, v Echo_echo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Echo_echo_Params_List) Copy(seg *capnp.Segment) (Echo_echo_Params_List, error) {
	var err error
	t, err := NewEcho_echo_Params_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Echo_echo_Params_Promise is a wrapper for a Echo_echo_Params promised by a client call.
type Echo_echo_Params_Promise struct{ *capnp.Pipeline }

func (p Echo_echo_Params_Promise) Struct() (Echo_echo_Params, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Params{s}, err
}

type Echo_echo_Results struct{ capnp.Struct }

// Echo_echo_Results_TypeID is the unique identifier for the type Echo_echo_Results.
const Echo_echo_Results_TypeID = 0x9b37d729b9dd7b9d

func NewEcho_echo_Results(s *capnp.Segment) (Echo_echo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Results{st}, err
}

func NewRootEcho_echo_Results(s *capnp.Segment) (Echo_echo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Echo_echo_Results{st}, err
}

func ReadRootEcho_echo_Results(msg *capnp.Message) (Echo_echo_Results, error) {
	root, err := msg.RootPtr()
	return Echo_echo_Results{root.Struct()}, err
}

func (s Echo_echo_Results) String() string {
	str, _ := text.Marshal(0x9b37d729b9dd7b9d, s.Struct)
	return str
}

func (s Echo_echo_Results) Copy(seg *capnp.Segment) (Echo_echo_Results, error) {
	var err error
	t, err := NewEcho_echo_Results(seg)
	if err != nil {
		return t, err
	}
	{
		v, err := s.Out()
		if err != nil {
			return t, err
		}
		t.SetOut(v)
	}
	return t, nil
}

func (s Echo_echo_Results) Out() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Echo_echo_Results) HasOut() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Echo_echo_Results) OutBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Echo_echo_Results) SetOut(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Echo_echo_Results_List is a list of Echo_echo_Results.
type Echo_echo_Results_List struct{ capnp.List }

// NewEcho_echo_Results creates a new list of Echo_echo_Results.
func NewEcho_echo_Results_List(s *capnp.Segment, sz int32) (Echo_echo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Echo_echo_Results_List{l}, err
}

func (s Echo_echo_Results_List) At(i int) Echo_echo_Results {
	return Echo_echo_Results{s.List.Struct(i)}
}

func (s Echo_echo_Results_List) Set(i int, v Echo_echo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Echo_echo_Results_List) Copy(seg *capnp.Segment) (Echo_echo_Results_List, error) {
	var err error
	t, err := NewEcho_echo_Results_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Echo_echo_Results_Promise is a wrapper for a Echo_echo_Results promised by a client call.
type Echo_echo_Results_Promise struct{ *capnp.Pipeline }

func (p Echo_echo_Results_Promise) Struct() (Echo_echo_Results, error) {
	s, err := p.Pipeline.Struct()
	return Echo_echo_Results{s}, err
}

type Hoth struct{ capnp.Struct }

// Hoth_TypeID is the unique identifier for the type Hoth.
const Hoth_TypeID = 0xad87da456fb0ebb9

func NewHoth(s *capnp.Segment) (Hoth, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hoth{st}, err
}

func NewRootHoth(s *capnp.Segment) (Hoth, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Hoth{st}, err
}

func ReadRootHoth(msg *capnp.Message) (Hoth, error) {
	root, err := msg.RootPtr()
	return Hoth{root.Struct()}, err
}

func (s Hoth) String() string {
	str, _ := text.Marshal(0xad87da456fb0ebb9, s.Struct)
	return str
}

func (s Hoth) Copy(seg *capnp.Segment) (Hoth, error) {
	var err error
	t, err := NewHoth(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.Base()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetBase(v)
	}
	return t, nil
}

func (s Hoth) Base() (EchoBase, error) {
	p, err := s.Struct.Ptr(0)
	return EchoBase{Struct: p.Struct()}, err
}

func (s Hoth) HasBase() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Hoth) SetBase(v EchoBase) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBase sets the base field to a newly
// allocated EchoBase struct, preferring placement in s's segment.
func (s Hoth) NewBase() (EchoBase, error) {
	ss, err := NewEchoBase(s.Struct.Segment())
	if err != nil {
		return EchoBase{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Hoth_List is a list of Hoth.
type Hoth_List struct{ capnp.List }

// NewHoth creates a new list of Hoth.
func NewHoth_List(s *capnp.Segment, sz int32) (Hoth_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Hoth_List{l}, err
}

func (s Hoth_List) At(i int) Hoth { return Hoth{s.List.Struct(i)} }

func (s Hoth_List) Set(i int, v Hoth) error { return s.List.SetStruct(i, v.Struct) }

func (s Hoth_List) Copy(seg *capnp.Segment) (Hoth_List, error) {
	var err error
	t, err := NewHoth_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Hoth_Promise is a wrapper for a Hoth promised by a client call.
type Hoth_Promise struct{ *capnp.Pipeline }

func (p Hoth_Promise) Struct() (Hoth, error) {
	s, err := p.Pipeline.Struct()
	return Hoth{s}, err
}

func (p Hoth_Promise) Base() EchoBase_Promise {
	return EchoBase_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type EchoBase struct{ capnp.Struct }

// EchoBase_TypeID is the unique identifier for the type EchoBase.
const EchoBase_TypeID = 0xa8bf13fef2674866

func NewEchoBase(s *capnp.Segment) (EchoBase, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EchoBase{st}, err
}

func NewRootEchoBase(s *capnp.Segment) (EchoBase, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return EchoBase{st}, err
}

func ReadRootEchoBase(msg *capnp.Message) (EchoBase, error) {
	root, err := msg.RootPtr()
	return EchoBase{root.Struct()}, err
}

func (s EchoBase) String() string {
	str, _ := text.Marshal(0xa8bf13fef2674866, s.Struct)
	return str
}

func (s EchoBase) Copy(seg *capnp.Segment) (EchoBase, error) {
	var err error
	t, err := NewEchoBase(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s EchoBase) Echo() Echo {
	p, _ := s.Struct.Ptr(0)
	return Echo{Client: p.Interface().Client()}
}

func (s EchoBase) HasEcho() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s EchoBase) SetEcho(v Echo) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// EchoBase_List is a list of EchoBase.
type EchoBase_List struct{ capnp.List }

// NewEchoBase creates a new list of EchoBase.
func NewEchoBase_List(s *capnp.Segment, sz int32) (EchoBase_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return EchoBase_List{l}, err
}

func (s EchoBase_List) At(i int) EchoBase { return EchoBase{s.List.Struct(i)} }

func (s EchoBase_List) Set(i int, v EchoBase) error { return s.List.SetStruct(i, v.Struct) }

func (s EchoBase_List) Copy(seg *capnp.Segment) (EchoBase_List, error) {
	var err error
	t, err := NewEchoBase_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// EchoBase_Promise is a wrapper for a EchoBase promised by a client call.
type EchoBase_Promise struct{ *capnp.Pipeline }

func (p EchoBase_Promise) Struct() (EchoBase, error) {
	s, err := p.Pipeline.Struct()
	return EchoBase{s}, err
}

func (p EchoBase_Promise) Echo() Echo {
	return Echo{Client: p.Pipeline.GetPipeline(0).Client()}
}

type StackingRoot struct{ capnp.Struct }

// StackingRoot_TypeID is the unique identifier for the type StackingRoot.
const StackingRoot_TypeID = 0x8fae7b41c61fc890

func NewStackingRoot(s *capnp.Segment) (StackingRoot, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return StackingRoot{st}, err
}

func NewRootStackingRoot(s *capnp.Segment) (StackingRoot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return StackingRoot{st}, err
}

func ReadRootStackingRoot(msg *capnp.Message) (StackingRoot, error) {
	root, err := msg.RootPtr()
	return StackingRoot{root.Struct()}, err
}

func (s StackingRoot) String() string {
	str, _ := text.Marshal(0x8fae7b41c61fc890, s.Struct)
	return str
}

func (s StackingRoot) Copy(seg *capnp.Segment) (StackingRoot, error) {
	var err error
	t, err := NewStackingRoot(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.A()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetA(v)
	}
	{
		o, err := s.AWithDefault()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetAWithDefault(v)
	}
	return t, nil
}

func (s StackingRoot) A() (StackingA, error) {
	p, err := s.Struct.Ptr(1)
	return StackingA{Struct: p.Struct()}, err
}

func (s StackingRoot) HasA() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s StackingRoot) SetA(v StackingA) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewA sets the a field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewA() (StackingA, error) {
	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s StackingRoot) AWithDefault() (StackingA, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return StackingA{}, err
	}
	ss, err := p.StructDefault(x_832bcc6686a26d56[64:96])
	return StackingA{Struct: ss}, err
}

func (s StackingRoot) HasAWithDefault() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StackingRoot) SetAWithDefault(v StackingA) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAWithDefault sets the aWithDefault field to a newly
// allocated StackingA struct, preferring placement in s's segment.
func (s StackingRoot) NewAWithDefault() (StackingA, error) {
	ss, err := NewStackingA(s.Struct.Segment())
	if err != nil {
		return StackingA{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// StackingRoot_List is a list of StackingRoot.
type StackingRoot_List struct{ capnp.List }

// NewStackingRoot creates a new list of StackingRoot.
func NewStackingRoot_List(s *capnp.Segment, sz int32) (StackingRoot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return StackingRoot_List{l}, err
}

func (s StackingRoot_List) At(i int) StackingRoot { return StackingRoot{s.List.Struct(i)} }

func (s StackingRoot_List) Set(i int, v StackingRoot) error { return s.List.SetStruct(i, v.Struct) }

func (s StackingRoot_List) Copy(seg *capnp.Segment) (StackingRoot_List, error) {
	var err error
	t, err := NewStackingRoot_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// StackingRoot_Promise is a wrapper for a StackingRoot promised by a client call.
type StackingRoot_Promise struct{ *capnp.Pipeline }

func (p StackingRoot_Promise) Struct() (StackingRoot, error) {
	s, err := p.Pipeline.Struct()
	return StackingRoot{s}, err
}

func (p StackingRoot_Promise) A() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p StackingRoot_Promise) AWithDefault() StackingA_Promise {
	return StackingA_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_832bcc6686a26d56[96:128])}
}

type StackingA struct{ capnp.Struct }

// StackingA_TypeID is the unique identifier for the type StackingA.
const StackingA_TypeID = 0x9d3032ff86043b75

func NewStackingA(s *capnp.Segment) (StackingA, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StackingA{st}, err
}

func NewRootStackingA(s *capnp.Segment) (StackingA, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return StackingA{st}, err
}

func ReadRootStackingA(msg *capnp.Message) (StackingA, error) {
	root, err := msg.RootPtr()
	return StackingA{root.Struct()}, err
}

func (s StackingA) String() string {
	str, _ := text.Marshal(0x9d3032ff86043b75, s.Struct)
	return str
}

func (s StackingA) Copy(seg *capnp.Segment) (StackingA, error) {
	var err error
	t, err := NewStackingA(seg)
	if err != nil {
		return t, err
	}
	{
		o, err := s.B()
		if err != nil {
			return t, err
		}
		v, err := o.Copy(seg)
		if err != nil {
			return t, err
		}
		t.SetB(v)
	}
	t.SetNum(s.Num())
	return t, nil
}

func (s StackingA) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingA) SetNum(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

func (s StackingA) B() (StackingB, error) {
	p, err := s.Struct.Ptr(0)
	return StackingB{Struct: p.Struct()}, err
}

func (s StackingA) HasB() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s StackingA) SetB(v StackingB) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewB sets the b field to a newly
// allocated StackingB struct, preferring placement in s's segment.
func (s StackingA) NewB() (StackingB, error) {
	ss, err := NewStackingB(s.Struct.Segment())
	if err != nil {
		return StackingB{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// StackingA_List is a list of StackingA.
type StackingA_List struct{ capnp.List }

// NewStackingA creates a new list of StackingA.
func NewStackingA_List(s *capnp.Segment, sz int32) (StackingA_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return StackingA_List{l}, err
}

func (s StackingA_List) At(i int) StackingA { return StackingA{s.List.Struct(i)} }

func (s StackingA_List) Set(i int, v StackingA) error { return s.List.SetStruct(i, v.Struct) }

func (s StackingA_List) Copy(seg *capnp.Segment) (StackingA_List, error) {
	var err error
	t, err := NewStackingA_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// StackingA_Promise is a wrapper for a StackingA promised by a client call.
type StackingA_Promise struct{ *capnp.Pipeline }

func (p StackingA_Promise) Struct() (StackingA, error) {
	s, err := p.Pipeline.Struct()
	return StackingA{s}, err
}

func (p StackingA_Promise) B() StackingB_Promise {
	return StackingB_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type StackingB struct{ capnp.Struct }

// StackingB_TypeID is the unique identifier for the type StackingB.
const StackingB_TypeID = 0x85257b30d6edf8c5

func NewStackingB(s *capnp.Segment) (StackingB, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return StackingB{st}, err
}

func NewRootStackingB(s *capnp.Segment) (StackingB, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return StackingB{st}, err
}

func ReadRootStackingB(msg *capnp.Message) (StackingB, error) {
	root, err := msg.RootPtr()
	return StackingB{root.Struct()}, err
}

func (s StackingB) String() string {
	str, _ := text.Marshal(0x85257b30d6edf8c5, s.Struct)
	return str
}

func (s StackingB) Copy(seg *capnp.Segment) (StackingB, error) {
	var err error
	t, err := NewStackingB(seg)
	if err != nil {
		return t, err
	}
	t.SetNum(s.Num())
	return t, nil
}

func (s StackingB) Num() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s StackingB) SetNum(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// StackingB_List is a list of StackingB.
type StackingB_List struct{ capnp.List }

// NewStackingB creates a new list of StackingB.
func NewStackingB_List(s *capnp.Segment, sz int32) (StackingB_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return StackingB_List{l}, err
}

func (s StackingB_List) At(i int) StackingB { return StackingB{s.List.Struct(i)} }

func (s StackingB_List) Set(i int, v StackingB) error { return s.List.SetStruct(i, v.Struct) }

func (s StackingB_List) Copy(seg *capnp.Segment) (StackingB_List, error) {
	var err error
	t, err := NewStackingB_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// StackingB_Promise is a wrapper for a StackingB promised by a client call.
type StackingB_Promise struct{ *capnp.Pipeline }

func (p StackingB_Promise) Struct() (StackingB, error) {
	s, err := p.Pipeline.Struct()
	return StackingB{s}, err
}

type CallSequence struct{ Client capnp.Client }

func (c CallSequence) GetNumber(ctx context.Context, params func(CallSequence_getNumber_Params) error, opts ...capnp.CallOption) CallSequence_getNumber_Results_Promise {
	if c.Client == nil {
		return CallSequence_getNumber_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(CallSequence_getNumber_Params{Struct: s}) }
	}
	return CallSequence_getNumber_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type CallSequence_Server interface {
	GetNumber(CallSequence_getNumber) error
}

func CallSequence_ServerToClient(s CallSequence_Server) CallSequence {
	c, _ := s.(server.Closer)
	return CallSequence{Client: server.New(CallSequence_Methods(nil, s), c)}
}

func CallSequence_Methods(methods []server.Method, s CallSequence_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xabaedf5f7817c820,
			MethodID:      0,
			InterfaceName: "aircraft.capnp:CallSequence",
			MethodName:    "getNumber",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := CallSequence_getNumber{c, opts, CallSequence_getNumber_Params{Struct: p}, CallSequence_getNumber_Results{Struct: r}}
			return s.GetNumber(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	return methods
}

// CallSequence_getNumber holds the arguments for a server call to CallSequence.getNumber.
type CallSequence_getNumber struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  CallSequence_getNumber_Params
	Results CallSequence_getNumber_Results
}

type CallSequence_getNumber_Params struct{ capnp.Struct }

// CallSequence_getNumber_Params_TypeID is the unique identifier for the type CallSequence_getNumber_Params.
const CallSequence_getNumber_Params_TypeID = 0xf58782f48a121998

func NewCallSequence_getNumber_Params(s *capnp.Segment) (CallSequence_getNumber_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CallSequence_getNumber_Params{st}, err
}

func NewRootCallSequence_getNumber_Params(s *capnp.Segment) (CallSequence_getNumber_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return CallSequence_getNumber_Params{st}, err
}

func ReadRootCallSequence_getNumber_Params(msg *capnp.Message) (CallSequence_getNumber_Params, error) {
	root, err := msg.RootPtr()
	return CallSequence_getNumber_Params{root.Struct()}, err
}

func (s CallSequence_getNumber_Params) String() string {
	str, _ := text.Marshal(0xf58782f48a121998, s.Struct)
	return str
}

func (s CallSequence_getNumber_Params) Copy(seg *capnp.Segment) (CallSequence_getNumber_Params, error) {
	var err error
	t, err := NewCallSequence_getNumber_Params(seg)
	if err != nil {
		return t, err
	}
	return t, nil
}

// CallSequence_getNumber_Params_List is a list of CallSequence_getNumber_Params.
type CallSequence_getNumber_Params_List struct{ capnp.List }

// NewCallSequence_getNumber_Params creates a new list of CallSequence_getNumber_Params.
func NewCallSequence_getNumber_Params_List(s *capnp.Segment, sz int32) (CallSequence_getNumber_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return CallSequence_getNumber_Params_List{l}, err
}

func (s CallSequence_getNumber_Params_List) At(i int) CallSequence_getNumber_Params {
	return CallSequence_getNumber_Params{s.List.Struct(i)}
}

func (s CallSequence_getNumber_Params_List) Set(i int, v CallSequence_getNumber_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s CallSequence_getNumber_Params_List) Copy(seg *capnp.Segment) (CallSequence_getNumber_Params_List, error) {
	var err error
	t, err := NewCallSequence_getNumber_Params_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// CallSequence_getNumber_Params_Promise is a wrapper for a CallSequence_getNumber_Params promised by a client call.
type CallSequence_getNumber_Params_Promise struct{ *capnp.Pipeline }

func (p CallSequence_getNumber_Params_Promise) Struct() (CallSequence_getNumber_Params, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Params{s}, err
}

type CallSequence_getNumber_Results struct{ capnp.Struct }

// CallSequence_getNumber_Results_TypeID is the unique identifier for the type CallSequence_getNumber_Results.
const CallSequence_getNumber_Results_TypeID = 0xa465f9502fd11e97

func NewCallSequence_getNumber_Results(s *capnp.Segment) (CallSequence_getNumber_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CallSequence_getNumber_Results{st}, err
}

func NewRootCallSequence_getNumber_Results(s *capnp.Segment) (CallSequence_getNumber_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CallSequence_getNumber_Results{st}, err
}

func ReadRootCallSequence_getNumber_Results(msg *capnp.Message) (CallSequence_getNumber_Results, error) {
	root, err := msg.RootPtr()
	return CallSequence_getNumber_Results{root.Struct()}, err
}

func (s CallSequence_getNumber_Results) String() string {
	str, _ := text.Marshal(0xa465f9502fd11e97, s.Struct)
	return str
}

func (s CallSequence_getNumber_Results) Copy(seg *capnp.Segment) (CallSequence_getNumber_Results, error) {
	var err error
	t, err := NewCallSequence_getNumber_Results(seg)
	if err != nil {
		return t, err
	}
	t.SetN(s.N())
	return t, nil
}

func (s CallSequence_getNumber_Results) N() uint32 {
	return s.Struct.Uint32(0)
}

func (s CallSequence_getNumber_Results) SetN(v uint32) {
	s.Struct.SetUint32(0, v)
}

// CallSequence_getNumber_Results_List is a list of CallSequence_getNumber_Results.
type CallSequence_getNumber_Results_List struct{ capnp.List }

// NewCallSequence_getNumber_Results creates a new list of CallSequence_getNumber_Results.
func NewCallSequence_getNumber_Results_List(s *capnp.Segment, sz int32) (CallSequence_getNumber_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return CallSequence_getNumber_Results_List{l}, err
}

func (s CallSequence_getNumber_Results_List) At(i int) CallSequence_getNumber_Results {
	return CallSequence_getNumber_Results{s.List.Struct(i)}
}

func (s CallSequence_getNumber_Results_List) Set(i int, v CallSequence_getNumber_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s CallSequence_getNumber_Results_List) Copy(seg *capnp.Segment) (CallSequence_getNumber_Results_List, error) {
	var err error
	t, err := NewCallSequence_getNumber_Results_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// CallSequence_getNumber_Results_Promise is a wrapper for a CallSequence_getNumber_Results promised by a client call.
type CallSequence_getNumber_Results_Promise struct{ *capnp.Pipeline }

func (p CallSequence_getNumber_Results_Promise) Struct() (CallSequence_getNumber_Results, error) {
	s, err := p.Pipeline.Struct()
	return CallSequence_getNumber_Results{s}, err
}

type Defaults struct{ capnp.Struct }

// Defaults_TypeID is the unique identifier for the type Defaults.
const Defaults_TypeID = 0x97e38948c61f878d

func NewDefaults(s *capnp.Segment) (Defaults, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Defaults{st}, err
}

func NewRootDefaults(s *capnp.Segment) (Defaults, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Defaults{st}, err
}

func ReadRootDefaults(msg *capnp.Message) (Defaults, error) {
	root, err := msg.RootPtr()
	return Defaults{root.Struct()}, err
}

func (s Defaults) String() string {
	str, _ := text.Marshal(0x97e38948c61f878d, s.Struct)
	return str
}

func (s Defaults) Copy(seg *capnp.Segment) (Defaults, error) {
	var err error
	t, err := NewDefaults(seg)
	if err != nil {
		return t, err
	}
	{
		v, err := s.Text()
		if err != nil {
			return t, err
		}
		t.SetText(v)
	}
	{
		v, err := s.Data()
		if err != nil {
			return t, err
		}
		t.SetData(v)
	}
	t.SetFloat(s.Float())
	t.SetInt(s.Int())
	t.SetUint(s.Uint())
	return t, nil
}

func (s Defaults) Text() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextDefault("foo"), err
}

func (s Defaults) HasText() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Defaults) TextBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytesDefault("foo"), err
}

func (s Defaults) SetText(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Defaults) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.DataDefault([]byte{0x62, 0x61, 0x72})), err
}

func (s Defaults) HasData() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Defaults) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s Defaults) Float() float32 {
	return math.Float32frombits(s.Struct.Uint32(0) ^ 0x4048f5c3)
}

func (s Defaults) SetFloat(v float32) {
	s.Struct.SetUint32(0, math.Float32bits(v)^0x4048f5c3)
}

func (s Defaults) Int() int32 {
	return int32(s.Struct.Uint32(4) ^ 4294967173)
}

func (s Defaults) SetInt(v int32) {
	s.Struct.SetUint32(4, uint32(v)^4294967173)
}

func (s Defaults) Uint() uint32 {
	return s.Struct.Uint32(8) ^ 42
}

func (s Defaults) SetUint(v uint32) {
	s.Struct.SetUint32(8, v^42)
}

// Defaults_List is a list of Defaults.
type Defaults_List struct{ capnp.List }

// NewDefaults creates a new list of Defaults.
func NewDefaults_List(s *capnp.Segment, sz int32) (Defaults_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Defaults_List{l}, err
}

func (s Defaults_List) At(i int) Defaults { return Defaults{s.List.Struct(i)} }

func (s Defaults_List) Set(i int, v Defaults) error { return s.List.SetStruct(i, v.Struct) }

func (s Defaults_List) Copy(seg *capnp.Segment) (Defaults_List, error) {
	var err error
	t, err := NewDefaults_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// Defaults_Promise is a wrapper for a Defaults promised by a client call.
type Defaults_Promise struct{ *capnp.Pipeline }

func (p Defaults_Promise) Struct() (Defaults, error) {
	s, err := p.Pipeline.Struct()
	return Defaults{s}, err
}

type BenchmarkA struct{ capnp.Struct }

// BenchmarkA_TypeID is the unique identifier for the type BenchmarkA.
const BenchmarkA_TypeID = 0xde2a1a960863c11c

func NewBenchmarkA(s *capnp.Segment) (BenchmarkA, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return BenchmarkA{st}, err
}

func NewRootBenchmarkA(s *capnp.Segment) (BenchmarkA, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return BenchmarkA{st}, err
}

func ReadRootBenchmarkA(msg *capnp.Message) (BenchmarkA, error) {
	root, err := msg.RootPtr()
	return BenchmarkA{root.Struct()}, err
}

func (s BenchmarkA) String() string {
	str, _ := text.Marshal(0xde2a1a960863c11c, s.Struct)
	return str
}

func (s BenchmarkA) Copy(seg *capnp.Segment) (BenchmarkA, error) {
	var err error
	t, err := NewBenchmarkA(seg)
	if err != nil {
		return t, err
	}
	{
		v, err := s.Name()
		if err != nil {
			return t, err
		}
		t.SetName(v)
	}
	{
		v, err := s.Phone()
		if err != nil {
			return t, err
		}
		t.SetPhone(v)
	}
	t.SetBirthDay(s.BirthDay())
	t.SetSiblings(s.Siblings())
	t.SetSpouse(s.Spouse())
	t.SetMoney(s.Money())
	return t, nil
}

func (s BenchmarkA) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s BenchmarkA) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BenchmarkA) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s BenchmarkA) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s BenchmarkA) BirthDay() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s BenchmarkA) SetBirthDay(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s BenchmarkA) Phone() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s BenchmarkA) HasPhone() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s BenchmarkA) PhoneBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s BenchmarkA) SetPhone(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s BenchmarkA) Siblings() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s BenchmarkA) SetSiblings(v int32) {
	s.Struct.SetUint32(8, uint32(v))
}

func (s BenchmarkA) Spouse() bool {
	return s.Struct.Bit(96)
}

func (s BenchmarkA) SetSpouse(v bool) {
	s.Struct.SetBit(96, v)
}

func (s BenchmarkA) Money() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s BenchmarkA) SetMoney(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// BenchmarkA_List is a list of BenchmarkA.
type BenchmarkA_List struct{ capnp.List }

// NewBenchmarkA creates a new list of BenchmarkA.
func NewBenchmarkA_List(s *capnp.Segment, sz int32) (BenchmarkA_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	return BenchmarkA_List{l}, err
}

func (s BenchmarkA_List) At(i int) BenchmarkA { return BenchmarkA{s.List.Struct(i)} }

func (s BenchmarkA_List) Set(i int, v BenchmarkA) error { return s.List.SetStruct(i, v.Struct) }

func (s BenchmarkA_List) Copy(seg *capnp.Segment) (BenchmarkA_List, error) {
	var err error
	t, err := NewBenchmarkA_List(seg, int32(s.Len()))
	if err != nil {
		return t, err
	}
	for i := 0; i < s.Len(); i++ {
		t.Set(i, s.At(i))
	}
	return t, nil
}

// BenchmarkA_Promise is a wrapper for a BenchmarkA promised by a client call.
type BenchmarkA_Promise struct{ *capnp.Pipeline }

func (p BenchmarkA_Promise) Struct() (BenchmarkA, error) {
	s, err := p.Pipeline.Struct()
	return BenchmarkA{s}, err
}

const schema_832bcc6686a26d56 = "x\xda\xacZ{x\x14U\x96?\xa7\xaa\xbb+\xafN" +
	"w\xe5\x961/\x0cdx\xc6\xd1!\x09\x1b\x18f\xfd" +
	"B0\x19\x19?\xd0\x14\x0d\xb2\xb82RI*Ic" +
	"\xbf\xec\xae\xe6\xa1\xe3\xc7\xcc\x0e\x8e\x8fOv\xe4S\xc6" +
	"';\x0e\x8b\xbb\xa2\xe0\x82#\xbb\xc2\xaa#\x0e8\x10" +
	"aG\xfcD\x85\x05\x14\x14\x15\xd4\x1dq\x86\x15\xc6G" +
	"\xedwnwuU:\xd5\xa2|\xfb\x0f]\xb9\xbfS" +
	"\xe7\x9e{\x1e\xbf{\xee-&^^6Mh\xf2\xfe" +
	"\x9a\x01\xa8\xefz}\xe6\xce3\x1f\xbf>\xf1\xe61\xb7" +
	"\x82\x1a@4\xaf\x89\xae\xfdE\xdf\x9e\x8b\x7f\x0e\x1e\x09" +
	"\x80m)\x1ed\xdb\x8b\xe9\xe9\xb9\xe26@s\xc1\xcb" +
	"\xc1\x1f\x17?;\xe3\xb6<Y\xaf@\"\x87\x8b\xb7\xb1" +
	"\xe3\\\xf8h\xf1\xbf\x01\x9ao\xfe\xf6\xab\x89\xdfi\xfb" +
	"\xfdm \x07\x9c\xb2(\x01\xb4\xac,\xa9@\xb6\xa6\x84" +
	"\x84\x1f(!\xcdS\x97Nk}j\xef\xa8\xdb\xf34" +
	"wH\x02\x00\xdb]2\xc8\xf6s\xe1}%K\x00\xcd" +
	"\xb5\x7f\xae}\xe3\xe9\xeb+\xef\x04YA\xc8j\xbc\xa4" +
	"T@@\xd6TJ\xda&\x8d\x19\x7fb{C\xe8\x1f" +
	"A\x0e\x88\xb62@\xa6\x96\xaee\xf3KI\xd3\xdc\xd2" +
	"+\xd8-\xf4d\xde\xbd\xab\xfe\xa5\xf6\x9b\x9f\xfce\x9e" +
	"\x9d|Mz\xe91v#\x97\x8f\x96\xd2\xccfr\xcf" +
	"I\xf5\x81\xdd\xf7\x0c\x95\xe5\xbe\xda_\xfa\";\\*" +
	"\x81h\x96\xbd\xf6\xab\x1d\x15OL\xbc\x17\xe4\x80g\xc8" +
	"\xec[K\x07\xd9N\xd2\x16z\xa1T\xc4\xd0!2\x19" +
	"\xcc\x95\xa3\x8a&\x9d\xb9\xe5w\xf7\xba\xf8\x89\x1d(\x1d" +
	"d\xc7\xf9\xfcG\xf9\xc2n\x88-T\xda\xbez~\xb5" +
	"\x9bO\xb1\xac\x02\x99\\F\xc2\xfe2\x12^y[\xfd" +
	"K3\xeex\xe7>\xf2\xa9\x90\xbf\xb2\xcb\xca^d\x9d" +
	"$\xdc\xd2^V\x8f\x80\xe6\xfc\x0d\xc2C\xf7=\xb0\xe5" +
	"A73\xe6\xfb\x07\x99\xee\xa7'\xcdO\x9a\xd7\xdc|" +
	"x\xeb\x847&?\xe4\x0c\xc0\x1d\xfe\x12\x0a\xc0*." +
	"\xb0k\xee1\xef\xb6q\xbf|h\x98\x0b6\xfb\x07\xd9" +
	"s\xa4)\xf4\x8c_\xc4\xd0\x0e?wA\xfa\x07\x9e_" +
	"\x98\xcd\x13\xd7\xe4\xe7\x15\x9f|\x8b\x7f\x90m\xe7\x93?" +
	"\xe7\xa7\x18\xdcw\xd1\xbe\xefu\x9d\xd5\xd7\x81Z\x87\x08" +
	"\xdc\xf9-\xd5\xe5I\x9a|L9M\xde7\xa3\xff\xd3" +
	"\xaf\xd8\xef\x1es[Ig\xf9\x8blV9=\xfd\x88" +
	"\xcb\x8e\xdcu\xe1\xd2\xeb\xdfz\xf2\x89a\x99\x12.?" +
	"\xc6\xd2\\\xf0\xc6\xf2+\xd8\x1az2O\xdd2\xabq" +
	"\xfa\xdcW\x9ep\xf3\xfe\xad\xe55\xc8V\xf3\x17Vq" +
	"\xcd[?\xdc\x14\xef<x\xdbF7+\xb6\x97\xafe" +
	"\xbb\xb9\xecN.+'z\xdf\x88y7lv\x93=" +
	"^\xfe);\xc5e?\xe6\xb2\xcb[\x7f\xbcb\xc1\x94" +
	"O6\x93\xaf\x1c&{E\x12\x19\x11x\x95M\x08\x90" +
	"=c\x02\xf3(\xaa\xa9\xc1I\xe6G\xc7.\xfaw\xb7" +
	"\x1chY\x15\x14\x90\xad\x09\xf2*\x0c\xbe\x0fh\xae\xfb" +
	"\x97mc^\x8a^\xfc\x1f\xa0\xca(\x9a\x07\x1ey\xfd" +
	"\x07\xf7\x9d\x18{\x12.@\x09\xa9\x04\xe4\xc7\x01YT" +
	"\xe6%X\xfa\xe9\xda\xb3\xfa\xc1?\xb8Y\xbc[~\x9c" +
	"\xed\x93\xe9i\xafL\x16_2\xeb\xfb\xcf\xbd\xf3\xc4\xdf" +
	"\xefv+\xb0\x8f\xe5Av\x96\xcb\x9e\xe6z\xb7}\xf6" +
	"\xd6\xfe\xeb\x17\x1fz\xd9\xcd\xc5jE\x0d2\xad\x82\x84" +
	"\x17T\x90\xe2\xfbw\xac/}O\xee\xd8\xe3f\xc4\xca" +
	"\x8aml5\x97]\xc5e\xdf\x1e;e\xe1;O\xfd" +
	"\xd6Uvg\xc5Z\xb6\x97\xcb\xee\xe6\xb2?\xea:v" +
	"\xe0\xd8\x93\x1d\xff\xe5\xea\xe2\x13\x15'\xd9i\x12n9" +
	"U\xc1]\xbc\xeb\xa7\xdb\xeb\x06O>\xf4G7\x93g" +
	")\x15\xc8\x16(\xbc\x86\x14R\xfd\xfc\xc1\xdbO\xad\xfb" +
	"\xbc\xf5u73\xeeP\x1ed\xab\xb8\xecJ.\xbbj" +
	"\xf2\xbf.\x8c\xfd\xf1\xd97\xc9\x0cO\xbe\xe36*\x83" +
	"l+\x09\xb7lQ\xb8\x19+_>\xb0\xe4\xf6\x85w" +
	"\x1dp\xd3\\]\xb9\x96\x8d\xaa\xe4\x09RI\x9a\xd9\xe2" +
	"\xcf\xc3}\xed\xfb\x0e\xbbE\xe4\xb2\xca\xb5\xac\x93\xcb\xb6" +
	"WRD\xea\xb6\xf7\x14\xfd\xaa\xa6\xf1H\xbe32V" +
	"T\xbe\xca\xb6Vr+*\xb9\x15k\xe6\xcc\xdb\xf8\x9f" +
	"Ov\x1dq\xdbM\x8a\xab\x1egr\x15\xe7\xa7*\xda" +
	" \xee\x99\xf8\xd8_\xff\xf6\xb5\x7f:\xe2\xe6\xb8G\xab" +
	"J\x90m\xe1\xc2\x9b\xab\xc8\xe4\x8d\xcfK\xf2\xfe}k" +
	"\x8f\xba-\xefh\xd56v\x82\xcb\x1e\xe7\xb2Z\xa8\xa5" +
	"b\xe7\xc9\xdd\xae\xb2\xfe\xea\x07\xd9\x05\xd5\xf4$W\x93" +
	"\xec\x99\xf9\x8f\xfc\xfc\xe1\xb5E\xc7\xdd\x8ch\xaa\xae@" +
	"\xd6\xce\x85/\xe3\xc2\x9b\xff0\xf7\xc8\x93\xc1\xab\x8e\xe7" +
	"\xad\xae\x13%\x0f\xd5G\xf5\x8b,\xcd\xa5o\xac\xa6j" +
	"\x1a\xf5\xbd3\xb5_\xdc\xba\x80T\x0bC\xb8\xa5\xb3f" +
	"\x1b\x9bU\xc3I\xa8\x86\x1cq\xc8w\xf6\xd7+\x96\xff" +
	",\xdf\x06\x9enkj\x06\xd9z.\xfb(\x97\x9d?" +
	"P\xf3\xa7\xa9\x1f\xaex\xcfmm\xd1\xda\x83lY-" +
	"=\xa5k\xf9n\xb1\xff\xf9u\x1bkn|\x7f\x18\x0b" +
	"\xaf\xa9\x1dd\xebI0\xb4\xaeV\xc4\xd0\xa6Z\xce\xc2" +
	"\xb9z\x1f\x1a\xe9N\x94.\xa6\xe9k\xefd\x1b\xe9\x9d" +
	"\x96\xf5\xb5\xf1bJ\x8c\xa2\x05\xd3\x8b7^}\xca\xcd" +
	"\x92\xbdc\x0f\xb2\x03c\xf9\xb68\x96Wju\xc5\x9d" +
	"\x7f\xf9\x87\xdbN\x83\\gQ\xb6w\xdc\"\x04\x8f\xf9" +
	"e\xdf\x15\xbb;\x0fy?\xcb#(\x9e0\xa7\xc6\xbe" +
	"\xca\xbe\xe4Z\xce\x8e\xa5T\\T\x15\xb9B\xe90?" +
	"s\x9bq\xee\xb8W\x996\x8es\xc38\x9a\xf1\xc8\xcc" +
	"\xe7\xef\x19o\xfc\xf3\x17n\x89\xf8\x9bq\xaf\xb2\x8d\\" +
	"v\xfd\xb86\xb8\xc4\xd4\xc2\xc9\x9e\xa4\xd6g\x08\x97\xf6" +
	"h\x89Xbj\xc8\xd0zn\x08\xc7\xfa\xa7\x03t!" +
	"\xaa\x1e\xd1\x03\xe0A\x00\xd9\xdf\x00\xa0\x16\x89\xa8*\x02" +
	"J\xb1t\x14= \xa0\x872.\xab\x01\xb3\x1a.o" +
	"\x8b\xa7c\x86\x9e\xa4\xd7\xcbr\xafw6\x02\xa8\xd3D" +
	"Tg\x0a\x88\xa8\xd0\xee%\xff\xa8\x19@\xed\x10Q\xed" +
	"\x12P\x16PA\x01@\x9eu%\x80:SDu@" +
	"\xc0@*|\x93\x8e^\x10\xd0\x0bX\xbf$\x9e\xecM" +
	"a\x19\x08X\x06h\xd2_\x91p\xca\x00\x00,\x07\xec" +
	"\x12\x91C\xe5\x0e\x8b\xc4\xacE3\xe2\x91\xde\xd45z" +
	"r\xce\x92\xf8\x9c%\xf1\xaeH\x1aSyk\x9b\x9a]" +
	"\xdbh\x01\xdb\xa2\xcbH\xad\xa53h3\"\xe0\x10\xed" +
	"\x96\xc7\xae\x89\x87{\xe7\xc6\xc2\xf1X\xc6cE\xa2\xa7" +
	"\xcc4\xb9\xda\x09\x15\x00\xeah\x11\xd5\x89\x02\xfa\xf1+" +
	"3\xb3\xeaKht\xbc\x88\xea$\x01Q\x03\x1fv\x83" +
	"o\x98\xc9\x9d=\x03\xf1K\xf5\x9e\x81\xf8\xe8.-\xa9" +
	"ES\xe0\xb4\xb6\xc6\x8e\x84\x18\x8e\xe5<\x92\x1f\x88N" +
	"\xa9g \x9eY\xa8\x17 \xd79\xa2\xd5\xc1\xc8r#" +
	"\x08\xb2W\x0a\xd0<\xd3\xb0\x0b\xb1`6\xcc\x96\xe2q" +
	"#\xbb:D\x0f\"\xca\x13\x16\xd9\xcb\x08f\x03\xfa}" +
	"Z\xda$\x11\xd5i\x02\x9a\xda\xbc\xb01\xd0\xa1\xf7A" +
	"@KG\x0c\x0c\xda\x9d\x0e \x06\xb9\xb3\xb0\x11\x00\xb5" +
	"a\x90\x8b\x8f\xf5dg4a,\x037#{\xe2\xb1" +
	"\x941\x93\xa7\x02\xe9\xcc\x05.\xc7\xc7<p \xe3\x85" +
	"r\x11\x06\xde\x92\x8aj\xe9\x9f:\xb7I\xae\x8e\xe9]" +
	"F\xf2kS?a$1h\x97X\x9e\xc1\xee)7" +
	"3\x9cB\xe3[\xa4\\\xaes)\x90r\x1dz\x1f\xf9" +
	"4\xc5\xedT2:\x11\xe5[\xa8\xc6\x96\x8a\xa8\xae\x10" +
	"P\xe6E\x86(\xdfJ\x83?\x15Q\xbdK@\x14\x14" +
	"\x14\x10\xe5\x95Tx\xb7\x8b\xa8\xde+\xa0,\xa2\x82\"" +
	"\xa2\xbc\x8aVy\x97\x88\xea\xfd\x02\xca\x1eA\xa10\xcb" +
	"\xab\xe9\xed\xbbET\x1f\x160`\xe8K\x8dl\xb6\x81" +
	"\x8c\x0dR_<\x1e\xe8\xd5\x0c\x0d\xfd \xa0\x9f\xc6j" +
	"\xa4n-Y\xdf\x17\x89k\x06\x96\x80p\xaa\xe4\xf7\xa7" +
	"gL\x03\x94\xc21\x83\xf8\xe2\x94\xe7V\xd34\x01\x03" +
	"i\x1a(\x02A.jtY\xde\xbc\xa4\x96\xc8\x84;" +
	"?\x10O\x01\xa8A\x11\xd5:\x01\xcdh\xb8\x7f\xc0\xb8" +
	"*n\xe0t}\xb6\xaeE\"\xcb\xea\xf9;\x18\xb4O" +
	"*\x05\x82c\x17\xd7l=\xc5\xfd\x08\x85\xa2\x1dO\x1b" +
	"\xc3\xeakH\xdeu\xc6\xd2\xd1L\xde\x05\xecM\x0f\xd0" +
	"\x1b@(\\N\xed9\xaa\xc8\xce9\xa1\xc1f\x0a\x8b" +
	"\x1d\x9d<\xe1$\\\xec\xc6\xa0}r\xcd[\xa2\xc7\"" +
	"a-\x12\x09\xe97\xa6\xf5X\x8f~i\xbfn\\\x95" +
	"\x8ev\xeb\xc9\xd1\xb3\xf5z\xbe`\xe7r+\xec\xe5b" +
	"\x8c\x82\x82E.\xa6\x93\xcf\xa6k)=?\"\x8d\xf6" +
	"\xdb\x9cGP\xb6O\xa0\x80(\xbb\xa8\xca\xd9&\xc5z" +
	"t\x9b\x9f\xac}\x12\xadC\x8e,\xcf\x06A.\x96L" +
	"\xcb~\xc0\xe4P\x9a\xca\xaf\xb6\xabcz\x87fh3" +
	"\xc3b\xea\xdb\x94\x9b\xb3\x98\xcb]\x88t\x86\x147\x06" +
	"\xbef\xdd\xddZJ\xc7\xa0}\xf6*@bDQ!" +
	"#\x99\xee\xa97.\xd7\x12\xb1\x02\x1c3Z@i\xb1" +
	"\xdec[\x97\xeb5\x0a\x90\xc1l\xbd?\xa9\xa7R\xe1" +
	"8r\x95U9\x95\x0f\x90\x8d\xf7\x8a\xa8>b'\xd5" +
	"\x1a\xda;\xee\x17Q]\xe7\xd8r\x7fC\x82\x0f\x8b\xa8" +
	">Ct ((\x02\xc8[\xc8c\x9bDT\xf7\x10" +
	"\x1d\xa0\x82\x1e\x00y7Y\xb9CD\xf5\x15\x01e\xaf" +
	"\xa0\xa0\x17@\xdeK\x83\xbbDT_\xb3}\x91\xeb\xdf" +
	"3\xbe\x10\xbb'b)\x08X\x0a\x18\xe8\xd6\x0d\xcdZ" +
	"]if\xc7nKD\xb4\x98\x9e\xb2\xd7\x9ck63" +
	"k\x96\x96E\xd3\xd6\xfb\xd2\xb2T\xaf\xf5<,\x112" +
	"\x8cK9@\x1b\xbd\x91\xe4\xd9\x1a\xccyD#K\xaf" +
	"\xcb\xf4\x16\xb2\xe5\x12\x9d\x06\x17\x8a\xa8F\xb2\x04\x09 " +
	"\x87\xc9#\xbd\"\xaa\x09\x8b \x01\xe4(\x0d\x0e\x88\xa8" +
	"\x1a\x14\"-\x82\"\x08(\x02J\xbd\xe9\xb8\xd5\xa4\x04" +
	"\x12F\xb2i\xf8\xfe@\xc3\xcd_\xb3mX\x99vm" +
	"\xe0\xd2\xfed\x82\x88\x81\xb6\xd4Fb\x86f\x9b\x19d" +
	"2\xefb\xa2\x86\xa965\xd4\xf7\x85\x93)\x03\x8bA" +
	"\xc0b\xc0\xb6\x94\xde\x13\x8f\xf5Z\x7f\xbaL@\x84\xfd" +
	"5\xb9\xec\xe0s\xd7m8\xeb\xd7|\xfejtZi" +
	"\x11X\xa3m\xe5y\xfa\xc5e;=\x9f\x02\xcf5\xda" +
	"\x05\x0a\xfc\xda\xb6\x94\x9e\\\x9ciY\x1d*\xbb\x01\xd4" +
	"2\x11\xd5\xf1\x02\x9aK\xb4\xb0\x11\x8e\xf5/\x02)\xde" +
	"\xedH\xd4\xdci\xb2\x80\xe6\xe9\xd2\xe4\x96\xc9\xe7\xa6\x8e" +
	"\xbcr)\xe0\xf89K\xe2\x81\xaeH:\x95W\xe6\x0d" +
	"v\x99\xe7|\xbf\xa6\xc1\xaes+\xa9se\xfe\x98#" +
	"\xa9\x1f\xa5\xc1GDT7X\xbb>\x80\xbc\x9e\xde^" +
	"'\xa2\xba\xc9Q\xe6\x1bI\xf21\x11\xd5\x1d\xe7N\x7f" +
	"g\x03\xe3\x08s\xde\xb0d$s\xdd} \x922Z" +
	",\xbfz\xcf\xdd\xc8w\x19\xc9o\xdbU\xe5\xeeb\x0a" +
	"\xc5J\xd4\xfa\xf3\xd4M\xb7C\xb5\xbc's\xac\xc1\xa0" +
	"}\x15\\ Z]\xc4f|\xb3\x84\xc2\x8cl\x87\xaa" +
	"9\x1b\xaa\xa7\xedPm\xa6ul\xb0\x18yZ\x1e#" +
	"?\xeb\x08\xd5V:.=\x93\x89\x8a\xec\x153\xa1\xda" +
	"N\x83/d\xb8;\x10\xd3\xa2\xba\xd5\xc3\xd4\x0f\xc4\xa3" +
	"6\xcf\x0e\xe9X8\x0f'5Jr+&m=Z" +
	"\xec\x87\x91e\x88  \xb53=ZB\xeb\x09S[" +
	"\x06\x96\x88\x19\xd5\x96\x86\x12\xba\xdeKc\xf9\xacl9" +
	"\xb6]ji\x9ex\xfeE\x90\xabRiQ\xbc\xbbp" +
	"\xdf4\x9cw:\x04\x94z\xa2\xbd\xd6\xe2\x03Z\xb2?" +
	"U\xe8\xach\xc5n\xba\x1e\xeb\x19\x88j\xc9\x1b\xb0\xfd" +
	"\x1b\xec\xa6W:*\xca\xdaM\x1fmvT\x94\xb5\x9b" +
	"\xae\xbf2[<OS\xec\x16fb7$\xcaV\x99" +
	"mi\xb6\xa3<$vfw8i\x0cthN\xf7" +
	"\xd7'\x06\xe21[\"\x15\xee\x8e\x84c\xfd)\x92\xc8" +
	"v\x8bm\xa9D<\x9d\xd2\xad\x18\xd6G\xe31}Y" +
	"\xc1H\xf1\xedA/|r\xe7;\xd0\xf0\xa3\xbb\x98=" +
	"\xbaS4f\x88\xa8\xce\x110\xb0L\xd7\x92\x16E\xd0" +
	"\xac\xc6\x00\xfa@@\x1f\x11\x86\xb6\xccz.X\xe2\xbc" +
	"\x97\xcf\x1d\xea\xbei\x89;;\x7f\xb7\x12\x9f\xd7\x96\xd4" +
	"\x12\xcdK\x9b\xcf\xebX\x91\xc7^\xc3\x94\xffPlj" +
	"=\xff,w\xe9e\xcf\x83\xe3r\x1fT\x0a4\x8b\xed" +
	"\xd9\xbf\xad\xc6\xc8\xba\xaa\xd0\x1a\xed\xce\xc8\xbe\xaa\xd0\x1b" +
	"\xed\xd6\xc8/|i\x0eo\x8e\xfc\xe2\x17f\xb6;j" +
	"\xb0\xbb\xa3\xc0\xe2x\xb8\x17|\x81\xee\xc9-\x931h" +
	"\xdf?gw\x03\xad\xa5y\"\xd9j\xdd\xdafw\x83" +
	"\xbe\xa6V\x0c\xda\x17\x98\x05\xbc\xdc\xde\x16N&\xe2I" +
	"\xee\x94\xbaL\xda5\xf2\x93Fg\x03\x00\x0a\xf2e\xf4" +
	"#\xca\x7fC?\x1e\xf9\x12\xfa\xf1\xcac\xe8\xc7'\x8f" +
	"h\x04\x08\xc4\xe21]Z\xd4w\x83\x14\xd1\x96J\xa9" +
	"\xbe\xb8\x14I/\x96z\xfb\x96\x04\x0c=e\x0cs\x18" +
	"\x0f\xc7\x1c}i6\x0f\x1de\xd1\xe0,\x0b\xebF\xab" +
	"![\x16\x0b\xa9,\xb2d\xbe\x80\x02\xf6w\x99c\xb9" +
	"d\xe4\x8e\xd0(E\xec\xc8e\xe9\xa8-\x922\x1c\xa3" +
	"\xe7\xe0\xaay\x99dND\xc4t\xea\xbc2\xday\xb5" +
	"\x15,t\xc2\xed\xd0\x8c\xccF\x96w\xa5\x12\x04\x08f" +
	"\xaeS\x86\xd3H\xc6Y\x1br\xf9\xc5\x1e\xf54\x02\x84" +
	"\x1e\xf1\x88\x18\xda\xe0q\xa4\x18[\xef\xa9\x01\x08\xad#" +
	"`\x93G\xc0\x11\xc2\x97f\x86G\xd9FO\x03@\xe8" +
	"1B\x9e&D\xfc\xc2\xcc\x90)\xdb\xcc\x91\x0d\x84<" +
	"C\x88\xe7s3s>a[8\xb2\x89\x90g\x09\xf1" +
	"\xfe\xd5\xcc\xd0*\xdb\xca\x91\xa7\x09y\x81\x10\xdfY\xd3" +
	"\xa3\xa0\x0f\x80=\xc7\x91g\x08\xd9A\x88t\xc6,R" +
	"2\x9f\xaa\xb8m\xcf\x12\xb2\x8b\x90\xa2\xcfh\x9e\"\x00" +
	"\xb6\x93\xbf\xf3\x02!{\x08)\xfe_\x9a\xa7\x18\x80\xed" +
	"\xe6\xc8\x0eB^!\xa4\xe44\xcdS\x02\xc0\xf6rd" +
	"\x17!\xaf\x11R\xfa\x17\x9a\xa7\x14\x80\xed\xe3\xf3\xec!" +
	"\xe4MB\xca\xfelNS\xb0\x0c\x80\xed\xe7n{\x85" +
	"\x90C\xe46\xff\xa7\xa6\x82~\x00v\x80\x03\xaf\x11\xf0" +
	"6\x01\xe5\xa7L\x05\xcb\x01\xd8a\x0e\xbcI\xc0\xbb\x04" +
	"\x04>1\x15\x0c\x00\xb0\xa3\x9e\xa9\x00\xa1C\x04\x9c!" +
	" \xf8'S\xa1\x18\xb2\xd3\x1c\xf8\x84\x802\xaf\x80~" +
	"\xf9\x7fL\x05e\x00V\xec%\xc0\xe3\x1514\x92\x80" +
	"\x8a\x8fM\x05+\x00\xd8\x08\x0eT\x110\x89\x00\xf6\x91" +
	"\xa9 \x03`M\x1c\xf8.\x013\x08P>4\x15T" +
	"\x00X\xa7\xb7\x19 4\x8d\x80\xeb\x08\xb8\xe0\xa4\xa9\xe0" +
	"\x05\x00l>\x7fc\x0e\x01\x09\x02*O\x98\x0aV\x02" +
	"\xb0(\x07\x06\x08XA\xc0\x85\x1f\x98\x0a^\x08\xc0~" +
	"\xc6\x81\x9f\x10p?\x01U\xef\x9b\x0aV\x01\xb0\xd5|" +
	"\x8e\xbb\x09\xd8@@\xdd{\xa6\x82\xd5\x94b^r\xc9" +
	":\x02^ `\xc4qS\xc1\x1a\x8a\xbcw:E\x9e" +
	"\x80w\x09\xb8\xe8]S\xc1Z\xf2\x15Wu\x88\x80\x0f" +
	"\x08\xa8\x7f\xc7T\xb0\x0e\x80\x1d\xe7\xc0\xdb\x04|D\xc0" +
	"\xc8c\xa6\x82#\x00\xd8\x09o7@\xe8\x03\x02<>" +
	"\x01\xfd\xa3\x8e\x9a\x0a^\x04\xc0\xd0w%\xc0l\x1f9" +
	"\x97\xc6\x1b\xde6\x15\xac'\xe7\xfa\xae\x05\x08\x15\x11\xa0" +
	"\x10\xf0\x9d\xb7L\x05G\x020\xd97\x1b \x14$\xa0" +
	"\xce'\xe0\x88\xd1G(\x81F\x01\xb0j\x1f\x99\xab\x10" +
	"2\x92^\x19s\xd8T\xb0\x81\xe2\xe1\xa3\x05V\x110" +
	"\x9a\x80\xb1\x87L\x05\xbf\x03\xc0Fq\xa0\x8e\x80\xf1\x04" +
	"\x8c\xfboS\xc1\xd1\x00l\x8c\x8frq$\x01\xdf%" +
	"`\xfcAS\xc11\x00l\x02\x99\x1b\x1aO@\x07\x01" +
	"\x13\x0e\x98\x0a\x8e\x05`\xed\x1c\x98F\xc0u\x04T\xbf" +
	"i*8\x8e\"\xc8\xad\x9aC@\x82\x80\x9a7L\x05" +
	"\xc7S\x0490@\xc0\x0a\x02j_7\x15\x9c@\x11" +
	"\xe4\xc0O\x08\xb8\x9f\x80\xc6\xfd&:\xbe\xbf\xb2\xd5\xbe" +
	"\x06\x10\xb2[\x88x\xd3M\x18\xb4\xbf\xceX;E\xeb" +
	"\xa4\xdcu@_K3\x96\x80\x80%\x80R\xb8u\x92" +
	"\xd5\x19I\xe1\x96f\xab\x07\x92\xc2M\xadV+\"\x86" +
	"\xa7\xa0\x00\x02\x0a\x80R\xbau\x92uP\x96\xd2-\xcd" +
	"\xd6\xdd\x97\x94njE\x09\x04\x94\x00\xc5\xf4\x14\xabU" +
	"\x09t\xc7\xe3\x11\xab\x8fr\xde\x85b\xa0;\x12\xef\xb6" +
	"\x0e\xcdm}\xad\x93\x1c\xf77\xd6\x0dG_K\xb3c" +
	"\xb4$;\x1a\x1e\"\xeb\xb5F\x87\xc8z\xac\xd1\xa6V" +
	"\xc7\xa8\x98\x19\xad\x0fOq\x0c\x0aY\xd1\xf4\x10\xb5\xc5" +
	"\xd6\xe8\x10\xb5E\xd6\xe8\x10\xb5RVm\xda\xa9\xd6\x97" +
	"\x19\x0c\xdc4\xe4^\xca\x19\x94r\xc0\xe5\x84:\x04\x0a" +
	"\xc9\xd5\xdfD-\xe6\xb0\xbd$3\xaea\xd0\xfe\xbc\x9e" +
	"\xb7)\xc1\xd0k\xb1\xbc+\"[\x0c \x0f%%\xc9" +
	"\xec%\x19\x88\xf1\x18\x06\xed\xffH\x90\x85\xf9\xfdS\xb7" +
	"\x96\x02ti\xd1\x96k\x99\x8e#\xef\xbc\x14\xa0\xc0\xff" +
	"?48\xdc\x1f\x8b\xf5\x1e\xc7\x87\xac\xfco\x17\x19!" +
	"-_\xc8\xe9)\x8a\x00%\xa8\xc3G\x98\x09\xdb\xf2\xec" +
	"\xab\xd6\xb0?;L\x19\xec\x18\xce6\x19R\x7f21" +
	"\xac\x0f\xb8JO\x19z\xb2\xe9rM\x1cvm\xd9h" +
	"7\xa3\x81\x94\x91,x\xbe:\xc7\xc5t\x97\x16Hj" +
	"\xd1T\x81\x8b\x8f\x0e\xcd@\xed\x9b\x9c\xfc\x1a\x1cW\xe6" +
	"\x05\xae(\x86\xdf\x9f\xce\x9b\xa3\xa7\x8c\xcb\xb5\x04\xe6\xaf" +
	"\xed\xda\xec\xfd\xcf\x14\x01\xcd\x98\x9e2fiF\x12\xc4" +
	"\xf0\xd2a)~\xae+\xda\xdc\xd54j_\xf3e\xc9" +
	"a\xf0\xff\x05\x00\x00\xff\xffYgL\xc7"

func init() {
	schemas.Register(schema_832bcc6686a26d56,
		0x85257b30d6edf8c5,
		0x8748bc095e10cb5d,
		0x87c33f2330feb3d8,
		0x8821cdb23640783a,
		0x8a165fb4d71bf3a2,
		0x8e5322c1e9282534,
		0x8fae7b41c61fc890,
		0x93c99951eacc72ff,
		0x9430ab12c496d40c,
		0x94bf7df83408218d,
		0x95befe3f14606e6b,
		0x97e38948c61f878d,
		0x9ab599979b02ac59,
		0x9b37d729b9dd7b9d,
		0x9b8f27ba05e255c8,
		0x9d3032ff86043b75,
		0xa465f9502fd11e97,
		0xa8bf13fef2674866,
		0xabaedf5f7817c820,
		0xabd055422a4d7df1,
		0xad87da456fb0ebb9,
		0xb1ac056ed7647011,
		0xb1f0385d845e367f,
		0xb61ee2ecff34ca73,
		0xb72b6dc625baa6a4,
		0xc7da65f9a2f20ba2,
		0xc95babe3bd394d2d,
		0xcbdc765fd5dff7ba,
		0xcc4411e60ba9c498,
		0xccb3b2e3603826e0,
		0xce44aee2d9e25049,
		0xcf9beaca1cc180c8,
		0xd636fba4f188dabe,
		0xd8bccf6e60a73791,
		0xd98c608877d9cb8d,
		0xddd1416669fb7613,
		0xde2a1a960863c11c,
		0xde50aebbad57549d,
		0xde9ed43cfaa83093,
		0xe1a2d1d51107bead,
		0xe1c9eac512335361,
		0xe508a29c83a059f8,
		0xe54e10aede55c7b1,
		0xe55d85fc1bf82f21,
		0xe5817f849ff906dc,
		0xe684eb3aef1a6859,
		0xe7711aada4bed56b,
		0xea26e9973bd6a0d9,
		0xf14fad09425d081c,
		0xf58782f48a121998,
		0xf705dc45c94766fd,
		0xf7ff4414476c186a,
		0xfca3742893be4cde)
}

var x_832bcc6686a26d56 = []byte{
	0, 0, 0, 0, 2, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0,
	223, 7, 8, 27, 0, 0, 0, 0,
	0, 0, 0, 0, 4, 0, 0, 0,
	1, 0, 0, 0, 23, 0, 0, 0,
	8, 0, 0, 0, 1, 0, 0, 0,
	223, 7, 8, 27, 0, 0, 0, 0,
	223, 7, 8, 28, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 3, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 0,
	42, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
