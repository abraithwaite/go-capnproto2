package capnp

import "fmt"

type ErrorSet []error

func (es ErrorSet) Error() string {
	if es == nil || len(es) == 0 {
		return ""
	}
	return fmt.Sprintf("first error: %v", es[0])
}
