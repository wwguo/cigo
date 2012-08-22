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

// activation.go defines activation functions for a neuron.

import (
	"math"
)


// Linear function is a type of activation function.
// 
// Parameters:
//     λ: parameter of function
// Outputs:
//     linear function
//
// Special cases:
//     linear(±0) = ±0
func Linear (lambda float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		f = lambda * x
		return f
	}
	f[1] = func (x float64) (f float64) {
		f = lambda
		return f
	}
	return
}

// Step function (Heaviside step function) is a type of activation function.
// 
// Parameters:
//     γ1,γ2: parameters of function
// Outputs:
//     step function
//
// Special cases:
//     step(±0) = γ1
func Step (gamma1,gamma2 float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		switch {
		case x >= 0:
			f = gamma1
		case x < 0:
			f = gamma2
		}
		return
	}
	f[1] = func (x float64) (f float64) {
		switch {
		case x == 0:
			f = math.Inf(1)
		default:
			f = 0
		}
		return
	}
	return
}

// Ramp function is a type of activation function.
//
// Parameters:
//     γ: parameter of function
//     ε: criterion for sections
// Outputs:
//     ramp function
//
// Special cases:
//     ramp( ε) =  γ
//     ramp(-ε) = -γ
func Ramp (gamma,epsilon float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		switch {
		case x >= epsilon:
			f = gamma
		case x <= epsilon:
			f = -gamma
		default:
			f = x
		}
		return
	}
	f[1] = func (x float64) (f float64) {
		switch {
		case x >= epsilon:
			f = 0
		case x <= epsilon:
			f = 0
		default:
			f = 1
		}
		return
	}
	return
}

// Sigmoid function is a type of activation function.
//
// Parameters:
//     λ: parameter of function
// Outputs:
//     sigmoid function
//
// Special cases:
//     ramp(+Inf) -> 1
//     ramp(-Inf) -> 0
func Sigmoid (lambda float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		f = 1/(1+math.Exp(-lambda*x))
		return
	}
	f[1] = func (x float64) (f float64) {
		f = 1/(1+math.Exp(-lambda*x))*(1 - 1/(1+math.Exp(-lambda*x))) // s(1-s)
		return
	}
	return
}

// Hyperbolic function is a type of activation function.
//
// Parameters:
//     λ: parameter of function
// Outputs:
//     hyperbolic function
//
// Special cases:
//     ramp(+Inf) ->  1
//     ramp(-Inf) -> -1
func Hyperbolic (lambda float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		a := math.Exp(lambda*x)
		b := math.Exp(-lambda*x)
		f = (a-b)/(a+b) // tanh(λx)
		return
	}
	f[1] = func (x float64) (f float64) {
		a := math.Exp(lambda*x)
		b := math.Exp(-lambda*x)
		f = lambda*math.Pow((a-b)/2,2) // λ[sinh(λx)]^2
		return
	}
	return
}

// Gaussian function is a type of activation function.
//
// Parameters:
//     σ: parameter of function
// Outputs:
//     gaussian function
//
// Special cases:
//     ramp(+Inf) -> 0
//     ramp(-Inf) -> 0
func Gaussian (sigma float64) (f [2]func (x float64) (f float64)) {
	f[0] = func (x float64) (f float64) {
		f = math.Exp(-math.Pow(x/sigma,2))
		return
	}
	f[1] = func (x float64) (f float64) {
		f = math.Exp(-math.Pow(x/sigma,2))*(-2*x/(sigma*sigma))
		return
	}
	return
}


// Bias function is used the activation function of bias.
func Bias () (f [2]func (float64) (float64)) {
	f[0] = func (x float64) (f float64) {
		f = -1
		return
	}
	f[1] = func (x float64) (f float64) {
		f = 0
		return
	}
	return
}
