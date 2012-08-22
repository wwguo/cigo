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

// rule.go defines learning rule functions for a neural network.

import (
	"math"
	"fmt"
)


// Gradient function is to caculate the sum of squared error for 
// gradient descent learning rule.  
//
// Parameters:
//     t: expected outputs
//     o: model outputs
// Outputs:
//     Single number
//
// Special cases:
//     None
func Gradient (t,o []float64) (epsilon float64) {
	if len(t) == len(o) {
		for i := 0; i < len(t); i++ {
			epsilon += math.Pow(t[i]-o[i],2)
		}
	} else {
		fmt.Printf("Vector lengths do not fit. t length: %g, o length: %g\n", len(t), len(o))
	}
	return epsilon
}


// update function is the criterion function.
//
// Parameters:
//     v: weight vector
//     eta: speed factor
//     delta: derivative function of the 
// Outputs:
//     Single number
//
// Special cases:
//     None
func WeightUpdate (d func (float64) float64) (func (v,t,o,z,eta float64) float64) {
	return func (v,t,o,z,eta float64) float64 {
		//  ∂E                   ∂f
		// ---- = −2(t_p − o_p)------z_{i,p}
		// ∂v_i                ∂net_p
		partialv := -2*(t - o)*d(z)*z
		//                ∂E
		// ∆v_i(t) = η(− ----)
		//               ∂v_i
		delta := eta*(-partialv)
		// vi(t) = vi(t − 1) + ∆vi(t)   
		newv := v + delta
		return newv
	}
}

// LMS(Widrow-Hoff) learning rule
// Madaline network: layered neural networks with multiple adaptive linear neurons. 
func LMS (v,t,o,z,eta float64) float64 {
	// least-means-square (LMS) algorithm
	// f = net_p  linear function, lambda=1
    //  ∂E
    // ---- = −2(t_p − o_p)z_{i,p}
    // ∂v_i
	// TODO  vi(t) = vi(t − 1) + 2η(t_p − o_p)z_{i,p}
	partialv := -2*(t - o)*z
	delta := eta*(-partialv)
	newv := v + delta
	return newv
}

// Generalized Delta learning rule
func GeneralDelta (v,t,o,z,eta float64) float64 {
    //  ∂E
    // ---- = −2(t_p − o_p)o_p(1-o_p)z_{i,p}
    // ∂v_i
	partialv := -2*(t - o)*o*(1-o)*z
	delta := eta*(-partialv)
	newv := v + delta
	return newv
}

