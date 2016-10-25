package rpc

import (
	"errors"
	"fmt"

	"zombiezen.com/go/capnproto2"
	rpccapnp "zombiezen.com/go/capnproto2/std/capnp/rpc"
)

// An Exception is a Cap'n Proto RPC error.
type Exception struct {
	rpccapnp.Exception
}

// Error returns the exception's reason.
func (e Exception) Error() string {
	var es capnp.ErrorSet
	r := e.Reason(&es)
	if es != nil {
		return "rpc exception"
	}
	return "rpc exception: " + r
}

// An Abort is a hang-up by a remote vat.
type Abort Exception

func copyAbort(m rpccapnp.Message) (Abort, error) {
	var es capnp.ErrorSet
	ma := m.Abort(&es)
	if es != nil {
		return Abort{}, es
	}
	msg, _, _ := capnp.NewMessage(capnp.SingleSegment(nil))
	if err := msg.SetRootPtr(ma.ToPtr()); err != nil {
		return Abort{}, err
	}
	p, err := msg.RootPtr()
	if err != nil {
		return Abort{}, err
	}
	return Abort{rpccapnp.Exception{Struct: p.Struct()}}, nil
}

// Error returns the exception's reason.
func (a Abort) Error() string {
	var es capnp.ErrorSet
	r := a.Reason(&es)
	if es != nil {
		return "rpc: aborted by remote"
	}
	return "rpc: aborted by remote: " + r
}

// toException sets fields on exc to match err.
func toException(exc rpccapnp.Exception, err error) {
	var es capnp.ErrorSet
	if ee, ok := err.(Exception); ok {
		// TODO(light): copy struct
		r := ee.Reason(&es)
		if es != nil {
			exc.SetReason(r)
		}
		exc.SetType(ee.Type())
		return
	}

	exc.SetReason(err.Error())
	exc.SetType(rpccapnp.EXCEPTIONTYPE_FAILED)
}

// Errors
var (
	ErrConnClosed = errors.New("rpc: connection closed")
)

// Internal errors
var (
	errQuestionReused  = errors.New("rpc: question ID reused")
	errNoMainInterface = errors.New("rpc: no bootstrap interface")
	errBadTarget       = errors.New("rpc: target not found")
	errShutdown        = errors.New("rpc: shutdown")
	errUnimplemented   = errors.New("rpc: remote used unimplemented protocol feature")
)

type bootstrapError struct {
	err error
}

func (e bootstrapError) Error() string {
	return "rpc bootstrap:" + e.err.Error()
}

type questionError struct {
	id     questionID
	method *capnp.Method // nil if this is bootstrap
	err    error
}

func (qe *questionError) Error() string {
	if qe.method == nil {
		return fmt.Sprintf("bootstrap call id=%d: %v", qe.id, qe.err)
	}
	return fmt.Sprintf("%v call id=%d: %v", qe.method, qe.id, qe.err)
}
