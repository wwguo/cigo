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

// Package ecm provides interfaces and types for evolutionary computation.
package ecm

// ga.go implements genetic algorighms.


// Choose within crossover operators.
//
// Parameters:
//     crossover_type: string
// Outputs:
//     crossover function
//
// Special cases:
//     None
func Crossover (crossover_type) (func (float64) float64) {
	switch crossover_type {
    case "one_point": 
		return OnePointCrossover;
    case "two_point":
		return TwoPointCrossover;
    case "uniform":
		return UniformCrossover;
	}
}

type GABinary struct {
	
}
