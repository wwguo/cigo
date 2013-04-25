// Copyright ©2012 Wei-Wei Guo <wwguocn@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package num provides interfaces and algorithms for numeric analysis.
package num

import (
	// "math"
	// "fmt"
)

// vector.go defines algorithms for slice vector manipulations.

type vector []float64
type index  []int

// Assign a vector.
//   n : length of the vector
//   a : `nil' for zeor vector
//       slice for recursive assignment
func Vector(n int, a []float64) (v vector) {
	v = make(vector, n, n)
	if a != nil {
		l := len(a)
		t := n/l
		r := n - l*t
		// Do nothing when n < len(a)
		for i := 0; i < t; i++ {
			copy(v[i*l:(i+1)*l], a)
		}
		if r != 0  {
			copy(v[t*l:t*l+r], a[0:r])
		}
	}
	return
}

// Return length of a vector.
func (v vector) Size() int {
	return len(v)
}

// // Resize a vector.
// func (v *vector) Resize(s,e int) {
// 	v = &v[s:e]
// }
