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

// Package base provides basic interfaces and algorithms.
package base

// sort.go defines algorithms for slice sorting.

type Vector []float64

func (a Vector) Sort () {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for i > 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
}
