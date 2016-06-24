// The MIT License (MIT)

// Copyright (c) 2016 Richard Tuin

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plalib_test

import (
	pla "github.com/rtuin/go-plalib"
	check "gopkg.in/check.v1"
)

type TargetSuite struct{}

var _ = check.Suite(&TargetSuite{})

var runTests = []struct {
	target pla.Target
	params []string
	rVal   bool
}{
	{pla.Target{Commands: make([]pla.Runnable, 0)}, make([]string, 0), false},
	{pla.Target{Commands: ([]pla.Runnable{pla.Command{"echo 1"}})}, make([]string, 0), false},
	{pla.Target{Commands: ([]pla.Runnable{pla.Command{"exit 1"}})}, make([]string, 0), true},
	{
		pla.Target{Commands: []pla.Runnable{pla.Target{Commands: []pla.Runnable{pla.Command{"echo foo"}}}}},
		make([]string, 0),
		false,
	},
	{
		pla.Target{Commands: []pla.Runnable{pla.Target{Commands: []pla.Runnable{pla.Command{"exit 1"}}}}},
		make([]string, 0),
		true,
	},
}

func (s *TargetSuite) TestRunTarget(c *check.C) {
	for i := range runTests {
		c.Assert(runTests[i].target.Run(runTests[i].params, false), check.Equals, runTests[i].rVal)
	}
}
