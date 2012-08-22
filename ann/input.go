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

// Package ann provides interfaces and types for artifical neural networks.
package ann

// input.go defines functions caculating the net input single for a neuron.


import (
	"math"
	"fmt"
)

// SU is the summation unit for computing net input signal.
// 
// Parameters:
//     I: indicate total number of neuron input signals
//     z: a vector of all signals, augmented with -1 for representing θ
//     v: a vector of weights, augmented by the bias θ
// Output:
//     Single number
func SU (I int) (f func (z,v []float64) (net float64)) {
	return func (z,v []float64) (net float64) {
		net = 0
		if len(z) == I && len(v) == I {
			for i := 0; i < I; i++ {
				net += z[i]*v[i]
			}
		} else {
			fmt.Printf("Sum lengths do not fit. z length: %g, v length: %g, and specified length: %g\n", len(z), len(v), I)
		}
		return net
	}
}
	
// PU is the prudoct unit for computing net input signal.
// 
// Parameters:
//     I: indicate total number of neuron input signals
//     z: a vector of all signals, augmented with -1 for representing θ
//     v: a vector of weights, augmented by the distortion factor θ
// Output:
//     Single number
func PU (I int) (f func (z,v []float64) (net float64)) {
	return func (z,v []float64) (net float64) {
		net = 1
		if len(z) == I && len(v) == I {
			for i := 0; i < I; i++ {
				net *= math.Pow(z[i],v[i])
			}
		} else {
			fmt.Printf("Lengths do not fit. z length: %g, v length: %g, and specified length: %g\n", len(z), len(v), I)
		}
		return net
	}
}


// Zero is used for computing augmented bias's input.
// 
// Parameters:
//     I: indicate total number of neuron input signals
//     z: a vector of all signals, augmented with -1 for representing θ
//     v: a vector of weights, augmented by the bias θ
// Output:
//     Single number
func Zero (z,v []float64) (net float64) {
	net = 0
	return net
}
