// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opt

import "fmt"

const (
	//The dimensions of the inputs do not make sense for this operation.
	errorDimensionMismatch = iota + 1
	//The indices provided are out of bounds.
	errorIllegalIndex
)

type error_ int

func (e error_) Error() string {
	switch e {
	case errorDimensionMismatch:
		return "Input dimensions do not match"
	case errorIllegalIndex:
		return "Index out of bounds"
	}
	return fmt.Sprintf("Unknown error code %d", e)
}
func (e error_) String() string {
	return e.Error()
}

var (
	//The dimensions of the inputs do not make sense for this operation.
	ErrorDimensionMismatch error_ = error_(errorDimensionMismatch)
	//The indices provided are out of bounds.
	ErrorIllegalIndex error_ = error_(errorIllegalIndex)
)
