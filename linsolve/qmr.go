// Copyright ©2018 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linsolve

type QMR struct {
	d  []float64
	p  []float64
	pt []float64
	q  []float64
	s  []float64
	v  []float64
	vt []float64
	w  []float64
	y  []float64
	z  []float64

	resume int
}

// Init implements the Method interface.
func (q *QMR) Init(dim int) {
	q.resume = 1
}

// Iterate implements the Method interface. It will command the following
// operations:
//  MulVec
//  PreconSolve
//  CheckResidualNorm
//  ComputeResidual
//  MajorIteration
//  NoOperation
func (q *QMR) Iterate(ctx *Context) (Operation, error) {
	switch q.resume {
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:

	default:
		panic("linsolve: QMR.Init not called")
	}
}
