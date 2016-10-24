package persistent

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

const PersistentAnnotation = uint64(0xf622595091cafb67)

type Persistent struct{ Client capnp.Client }

func (c Persistent) Save(ctx context.Context, params func(PersistentSaveParams) error, opts ...capnp.CallOption) PersistentSaveResults_Promise {
	if c.Client == nil {
		return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(PersistentSaveParams{Struct: s}) }
	}
	return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Persistent_Server interface {
	Save(Persistent_save) error
}

func Persistent_ServerToClient(s Persistent_Server) Persistent {
	c, _ := s.(server.Closer)
	return Persistent{Client: server.New(Persistent_Methods(nil, s), c)}
}

func Persistent_Methods(methods []server.Method, s Persistent_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc8cb212fcd9f5691,
			MethodID:      0,
			InterfaceName: "persistent.capnp:Persistent",
			MethodName:    "save",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Persistent_save{c, opts, PersistentSaveParams{Struct: p}, PersistentSaveResults{Struct: r}}
			return s.Save(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Persistent_save holds the arguments for a server call to Persistent.save.
type Persistent_save struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  PersistentSaveParams
	Results PersistentSaveResults
}

type PersistentSaveParams struct{ capnp.Struct }

// PersistentSaveParams_TypeID is the unique identifier for the type PersistentSaveParams.
const PersistentSaveParams_TypeID = 0xf76fba59183073a5

func NewPersistentSaveParams(s *capnp.Segment) (PersistentSaveParams, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PersistentSaveParams{st}, err
}

func NewRootPersistentSaveParams(s *capnp.Segment) (PersistentSaveParams, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PersistentSaveParams{st}, err
}

func ReadRootPersistentSaveParams(msg *capnp.Message) (PersistentSaveParams, error) {
	root, err := msg.RootPtr()
	return PersistentSaveParams{root.Struct()}, err
}

func (s PersistentSaveParams) String() string {
	str, _ := text.Marshal(0xf76fba59183073a5, s.Struct)
	return str
}

func (s PersistentSaveParams) SealFor(e *capnp.ErrorSet) capnp.Pointer {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		*e = append(*e, err)
		return nil
	}
	return p
}

func (s PersistentSaveParams) HasSealFor() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PersistentSaveParams) SealForPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s PersistentSaveParams) SetSealFor(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s PersistentSaveParams) SetSealForPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// PersistentSaveParams_List is a list of PersistentSaveParams.
type PersistentSaveParams_List struct{ capnp.List }

// NewPersistentSaveParams creates a new list of PersistentSaveParams.
func NewPersistentSaveParams_List(s *capnp.Segment, sz int32) (PersistentSaveParams_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return PersistentSaveParams_List{l}, err
}

func (s PersistentSaveParams_List) At(i int) PersistentSaveParams {
	return PersistentSaveParams{s.List.Struct(i)}
}

func (s PersistentSaveParams_List) Set(i int, v PersistentSaveParams) error {
	return s.List.SetStruct(i, v.Struct)
}

// PersistentSaveParams_Promise is a wrapper for a PersistentSaveParams promised by a client call.
type PersistentSaveParams_Promise struct{ *capnp.Pipeline }

func (p PersistentSaveParams_Promise) Struct() (PersistentSaveParams, error) {
	s, err := p.Pipeline.Struct()
	return PersistentSaveParams{s}, err
}

func (p PersistentSaveParams_Promise) SealFor() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type PersistentSaveResults struct{ capnp.Struct }

// PersistentSaveResults_TypeID is the unique identifier for the type PersistentSaveResults.
const PersistentSaveResults_TypeID = 0xb76848c18c40efbf

func NewPersistentSaveResults(s *capnp.Segment) (PersistentSaveResults, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PersistentSaveResults{st}, err
}

func NewRootPersistentSaveResults(s *capnp.Segment) (PersistentSaveResults, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PersistentSaveResults{st}, err
}

func ReadRootPersistentSaveResults(msg *capnp.Message) (PersistentSaveResults, error) {
	root, err := msg.RootPtr()
	return PersistentSaveResults{root.Struct()}, err
}

func (s PersistentSaveResults) String() string {
	str, _ := text.Marshal(0xb76848c18c40efbf, s.Struct)
	return str
}

func (s PersistentSaveResults) SturdyRef(e *capnp.ErrorSet) capnp.Pointer {
	p, err := s.Struct.Pointer(0)
	if err != nil {
		*e = append(*e, err)
		return nil
	}
	return p
}

func (s PersistentSaveResults) HasSturdyRef() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PersistentSaveResults) SturdyRefPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s PersistentSaveResults) SetSturdyRef(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s PersistentSaveResults) SetSturdyRefPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// PersistentSaveResults_List is a list of PersistentSaveResults.
type PersistentSaveResults_List struct{ capnp.List }

// NewPersistentSaveResults creates a new list of PersistentSaveResults.
func NewPersistentSaveResults_List(s *capnp.Segment, sz int32) (PersistentSaveResults_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return PersistentSaveResults_List{l}, err
}

func (s PersistentSaveResults_List) At(i int) PersistentSaveResults {
	return PersistentSaveResults{s.List.Struct(i)}
}

func (s PersistentSaveResults_List) Set(i int, v PersistentSaveResults) error {
	return s.List.SetStruct(i, v.Struct)
}

// PersistentSaveResults_Promise is a wrapper for a PersistentSaveResults promised by a client call.
type PersistentSaveResults_Promise struct{ *capnp.Pipeline }

func (p PersistentSaveResults_Promise) Struct() (PersistentSaveResults, error) {
	s, err := p.Pipeline.Struct()
	return PersistentSaveResults{s}, err
}

func (p PersistentSaveResults_Promise) SturdyRef() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type RealmGateway struct{ Client capnp.Client }

func (c RealmGateway) Import(ctx context.Context, params func(RealmGatewayimportParams) error, opts ...capnp.CallOption) PersistentSaveResults_Promise {
	if c.Client == nil {
		return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x84ff286cd00a3ed4,
			MethodID:      0,
			InterfaceName: "persistent.capnp:RealmGateway",
			MethodName:    "import",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(RealmGatewayimportParams{Struct: s}) }
	}
	return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c RealmGateway) Export(ctx context.Context, params func(RealmGatewayexportParams) error, opts ...capnp.CallOption) PersistentSaveResults_Promise {
	if c.Client == nil {
		return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0x84ff286cd00a3ed4,
			MethodID:      1,
			InterfaceName: "persistent.capnp:RealmGateway",
			MethodName:    "export",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		call.ParamsFunc = func(s capnp.Struct) error { return params(RealmGatewayexportParams{Struct: s}) }
	}
	return PersistentSaveResults_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type RealmGateway_Server interface {
	Import(RealmGateway_import) error

	Export(RealmGateway_export) error
}

func RealmGateway_ServerToClient(s RealmGateway_Server) RealmGateway {
	c, _ := s.(server.Closer)
	return RealmGateway{Client: server.New(RealmGateway_Methods(nil, s), c)}
}

func RealmGateway_Methods(methods []server.Method, s RealmGateway_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x84ff286cd00a3ed4,
			MethodID:      0,
			InterfaceName: "persistent.capnp:RealmGateway",
			MethodName:    "import",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := RealmGateway_import{c, opts, RealmGatewayimportParams{Struct: p}, PersistentSaveResults{Struct: r}}
			return s.Import(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x84ff286cd00a3ed4,
			MethodID:      1,
			InterfaceName: "persistent.capnp:RealmGateway",
			MethodName:    "export",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := RealmGateway_export{c, opts, RealmGatewayexportParams{Struct: p}, PersistentSaveResults{Struct: r}}
			return s.Export(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// RealmGateway_import holds the arguments for a server call to RealmGateway.import.
type RealmGateway_import struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  RealmGatewayimportParams
	Results PersistentSaveResults
}

// RealmGateway_export holds the arguments for a server call to RealmGateway.export.
type RealmGateway_export struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  RealmGatewayexportParams
	Results PersistentSaveResults
}

type RealmGatewayimportParams struct{ capnp.Struct }

// RealmGatewayimportParams_TypeID is the unique identifier for the type RealmGatewayimportParams.
const RealmGatewayimportParams_TypeID = 0xf0c2cc1d3909574d

func NewRealmGatewayimportParams(s *capnp.Segment) (RealmGatewayimportParams, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return RealmGatewayimportParams{st}, err
}

func NewRootRealmGatewayimportParams(s *capnp.Segment) (RealmGatewayimportParams, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return RealmGatewayimportParams{st}, err
}

func ReadRootRealmGatewayimportParams(msg *capnp.Message) (RealmGatewayimportParams, error) {
	root, err := msg.RootPtr()
	return RealmGatewayimportParams{root.Struct()}, err
}

func (s RealmGatewayimportParams) String() string {
	str, _ := text.Marshal(0xf0c2cc1d3909574d, s.Struct)
	return str
}

func (s RealmGatewayimportParams) Cap() Persistent {
	p, _ := s.Struct.Ptr(0)
	return Persistent{Client: p.Interface().Client()}
}

func (s RealmGatewayimportParams) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RealmGatewayimportParams) SetCap(v Persistent) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s RealmGatewayimportParams) Params(e *capnp.ErrorSet) PersistentSaveParams {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return PersistentSaveParams{}
	}
	return PersistentSaveParams{Struct: p.Struct()}
}

func (s RealmGatewayimportParams) HasParams() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RealmGatewayimportParams) SetParams(v PersistentSaveParams) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewParams sets the params field to a newly
// allocated PersistentSaveParams struct, preferring placement in s's segment.
func (s RealmGatewayimportParams) NewParams() (PersistentSaveParams, error) {
	ss, err := NewPersistentSaveParams(s.Struct.Segment())
	if err != nil {
		return PersistentSaveParams{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// RealmGatewayimportParams_List is a list of RealmGatewayimportParams.
type RealmGatewayimportParams_List struct{ capnp.List }

// NewRealmGatewayimportParams creates a new list of RealmGatewayimportParams.
func NewRealmGatewayimportParams_List(s *capnp.Segment, sz int32) (RealmGatewayimportParams_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return RealmGatewayimportParams_List{l}, err
}

func (s RealmGatewayimportParams_List) At(i int) RealmGatewayimportParams {
	return RealmGatewayimportParams{s.List.Struct(i)}
}

func (s RealmGatewayimportParams_List) Set(i int, v RealmGatewayimportParams) error {
	return s.List.SetStruct(i, v.Struct)
}

// RealmGatewayimportParams_Promise is a wrapper for a RealmGatewayimportParams promised by a client call.
type RealmGatewayimportParams_Promise struct{ *capnp.Pipeline }

func (p RealmGatewayimportParams_Promise) Struct() (RealmGatewayimportParams, error) {
	s, err := p.Pipeline.Struct()
	return RealmGatewayimportParams{s}, err
}

func (p RealmGatewayimportParams_Promise) Cap() Persistent {
	return Persistent{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p RealmGatewayimportParams_Promise) Params() PersistentSaveParams_Promise {
	return PersistentSaveParams_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type RealmGatewayexportParams struct{ capnp.Struct }

// RealmGatewayexportParams_TypeID is the unique identifier for the type RealmGatewayexportParams.
const RealmGatewayexportParams_TypeID = 0xecafa18b482da3aa

func NewRealmGatewayexportParams(s *capnp.Segment) (RealmGatewayexportParams, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return RealmGatewayexportParams{st}, err
}

func NewRootRealmGatewayexportParams(s *capnp.Segment) (RealmGatewayexportParams, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return RealmGatewayexportParams{st}, err
}

func ReadRootRealmGatewayexportParams(msg *capnp.Message) (RealmGatewayexportParams, error) {
	root, err := msg.RootPtr()
	return RealmGatewayexportParams{root.Struct()}, err
}

func (s RealmGatewayexportParams) String() string {
	str, _ := text.Marshal(0xecafa18b482da3aa, s.Struct)
	return str
}

func (s RealmGatewayexportParams) Cap() Persistent {
	p, _ := s.Struct.Ptr(0)
	return Persistent{Client: p.Interface().Client()}
}

func (s RealmGatewayexportParams) HasCap() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RealmGatewayexportParams) SetCap(v Persistent) error {
	if v.Client == nil {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s RealmGatewayexportParams) Params(e *capnp.ErrorSet) PersistentSaveParams {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		*e = append(*e, err)
		return PersistentSaveParams{}
	}
	return PersistentSaveParams{Struct: p.Struct()}
}

func (s RealmGatewayexportParams) HasParams() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RealmGatewayexportParams) SetParams(v PersistentSaveParams) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewParams sets the params field to a newly
// allocated PersistentSaveParams struct, preferring placement in s's segment.
func (s RealmGatewayexportParams) NewParams() (PersistentSaveParams, error) {
	ss, err := NewPersistentSaveParams(s.Struct.Segment())
	if err != nil {
		return PersistentSaveParams{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// RealmGatewayexportParams_List is a list of RealmGatewayexportParams.
type RealmGatewayexportParams_List struct{ capnp.List }

// NewRealmGatewayexportParams creates a new list of RealmGatewayexportParams.
func NewRealmGatewayexportParams_List(s *capnp.Segment, sz int32) (RealmGatewayexportParams_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return RealmGatewayexportParams_List{l}, err
}

func (s RealmGatewayexportParams_List) At(i int) RealmGatewayexportParams {
	return RealmGatewayexportParams{s.List.Struct(i)}
}

func (s RealmGatewayexportParams_List) Set(i int, v RealmGatewayexportParams) error {
	return s.List.SetStruct(i, v.Struct)
}

// RealmGatewayexportParams_Promise is a wrapper for a RealmGatewayexportParams promised by a client call.
type RealmGatewayexportParams_Promise struct{ *capnp.Pipeline }

func (p RealmGatewayexportParams_Promise) Struct() (RealmGatewayexportParams, error) {
	s, err := p.Pipeline.Struct()
	return RealmGatewayexportParams{s}, err
}

func (p RealmGatewayexportParams_Promise) Cap() Persistent {
	return Persistent{Client: p.Pipeline.GetPipeline(0).Client()}
}

func (p RealmGatewayexportParams_Promise) Params() PersistentSaveParams_Promise {
	return PersistentSaveParams_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

const schema_b8630836983feed7 = "x\xda\xbcSKh3U\x14>\xdf\xdcy$\xd2\x9f" +
	"?7\xf3\x97Vh\x8d--\xb5b_\x0ab\x8b\x98" +
	"L\xa1\xb6]\xa8\x99\x14\x94\x0a\x827u\xaa\x85d2" +
	"\xccL_\x0bQA\x10\xd1.\xbas'\"n\x0a\xe2" +
	"R\xebF,\x0a\xe2\x03\x14\x94\xee\xbb\x16\xad.t\xe1" +
	"\xe2\xca\x9d\xbc\xc66\xa5H\xc5\xcd\xc0\x9c{\xee\xfd\xce" +
	"\xf9\x1e\xb3\xe7(is\xc6\xa9N\xe4>b\x98\xf2\xc7" +
	"\xc7\xee\xfa\xbev\x9f|\x9d8g\xf2\xf4\x97\xe2;\x0f" +
	"g6>!\xa2\x1clW\xfb\xdd~N\xb3\x88\xecu" +
	"\xed\x0d[0\xcb\x16lB~\xf6k\xe9\xed\xcfW^" +
	"\xfa\x98\xf8\x10\xe4\xe1\xd3\xef~73\xf2\xcdWd\xc0" +
	"\xca\xe1\xa1\x97\xd9\"\xec\x03\xa6\xae\xbc\xc9\x8a\x94:\xbf" +
	"\xfc\xfa\x87\xec\xcc>f\x13D\xf6Ol\xd9\x1e\xd7-" +
	"{\\\x1f\x90G\xefO\xad\xbc\xf5\xdeG?\x13\xbf\x07" +
	"D\x86\xa6^\x1d\xd7\xab \xd8s\xfa.A>\xf1L" +
	"v~\xf8\xdb\x93\xf3t\xc3\xb0\x914L\x1a\xaa\xe1\xc5" +
	"\xbf\xbe>,\xaf\x8f\xfeA?p\xe3^\xa4@a\xf7" +
	"\x9bg\xf6\x88i\xd9#fam\xc5d \xc8\x0f\xa2" +
	"\xd9\xc1\xf5O\x1b\x7f\xf6\xdaf\xde\\\x80\xbdj\xaam" +
	"\x96\xcc\"=)\x03/\x8c\xb6\xa2\xd8\xd3\xfcxzC" +
	"\x04~\xb0P\xf1D\xad\xbe,\x0a\xb1\xb7+\xf6\xcb\x80" +
	"\x9ba\x06QgH\xb4\xc9\xe2s\x0bD\xce\x03p\x1e" +
	"\x05\x7f\xcd\x02:{v;\xb6UG\x00\xe7U\xf0/" +
	"\xad\xe2V=h\x841G\xc1\xd55tE\x02\x88\xda" +
	"\xc5\xce\xb4\xaa4\xe1f\x80\x04_}s\x0c\xc8#u" +
	"\x8f\x08y\x03Z\xbaPB\xd1\xdb\xbb1\x88q\x19\x85" +
	"\xfd\x13\xa5\x0c89\xf0[U\xce\xab\xbc?\xe4w\x87" +
	"r\xd5\x8f\xbd\xd0\x175\xb2*\xde\xa6\\\xdaK\xffu" +
	"\xce\x0aO\xed\xfa^\xd8=m\xfd\xb75\xd0;\x1a\x94" +
	"[\x15?\x9e^\x13;^\xc5\x8b\xb6kqDJ\x0d" +
	"\x9d\xe9D\xbaZ\xe7V\x85\xc8\xedcp\x075\xc8(" +
	"\xde\x0e_\xd8\xafx\x84\xcd\x84\xa6\xd4\x92\xc8\x13z\xe8" +
	"\xdc\xc2\xb0<?V\x1c\xa4l\x93}6\x95\x88lU" +
	"\xaa\x11\xca\"\x14\xc4\xea\x91l\xcfCV-\x8e\\=" +
	"\xf1F\xfbjWy~?\x91\xd3\x07g\x08|\xca\xba" +
	"\x1d\x89\x1d\xef\x12\xfb=%Q\xc5&\xbf\x19p\xa3\xc2" +
	"\xb3\x0f\xca\xb5\xeefW\xf2\xd5\xf2lb\xd9\xe9\xa6\x05" +
	"\xc6\xca\"\xb4D=r3\x1d\xc2&G\x89\xdc1\x06" +
	"wO\x03\x07\xee$\x13(\x8f\xba\x01\x83\xfb\x85\x06k" +
	"C\x04\xe0i\xeeJ\xf8\x8f\xac\x09N(\x06\"\x14\xf5" +
	"\x08\xb9.\xd7\xff\x06\xe1:_\"\x97\xd2\xf9\x0an\x9a" +
	"\x19\xfc\xbf\xb9\xb9v\xf2\x9b\x93s\x1d\xfd\xb9\x9e\x19\x08" +
	"\xba\x19H\xb2\x05M\xfev03\x90\x7f\xfe\xf8\x84\x14" +
	"\xac3\x08\xf4)\xc4#\xd9\x8e$\xfc\xd8\xf1\xfdF," +
	"n\xc7[\x0d\x9f\xf4\xce\xab\xec\xaa\xf4\x16\xcb\xc9f\x17" +
	"\xc2\xbbH\xa4\xe6w\xefhx%\xf2D\xed\xf1F\xd8" +
	"$\xeaBr\xff\x0e\x00\x00\xff\xff\x1ar\xf0\xc2"

func init() {
	schemas.Register(schema_b8630836983feed7,
		0x84ff286cd00a3ed4,
		0xb76848c18c40efbf,
		0xc8cb212fcd9f5691,
		0xecafa18b482da3aa,
		0xf0c2cc1d3909574d,
		0xf622595091cafb67,
		0xf76fba59183073a5)
}
