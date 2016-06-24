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
	// "github.com/rtuin/go-plalib"
	check "gopkg.in/check.v1"
)

type TargetRunnerSuite struct{}

var _ = check.Suite(&TargetRunnerSuite{})

// var findTests = []struct {
// 	name     string
// 	targets  []pla.Target
// 	expected pla.Target
// 	err      check.Checker
// }{
// 	{"foo", make([]pla.Target, 0), pla.Target{}, check.NotNil},
// 	{"Foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{}, check.NotNil},
// 	{"foo", []pla.Target{pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, check.IsNil},
// 	{"foo", []pla.Target{pla.Target{Name: "bar"}, pla.Target{Name: "baz"}}, pla.Target{}, check.NotNil},
// 	{"foo", []pla.Target{pla.Target{Name: "oof"}, pla.Target{Name: "foo"}}, pla.Target{Name: "foo"}, check.IsNil},
// }

// func (s *TargetRunnerSuite) TestFindTargetByName(c *check.C) {
// 	for i := range findTests {
// 		target, err := pla.FindTargetByTargetName(findTests[i].name, findTests[i].targets)
// 		c.Assert(err, findTests[i].err)
// 		c.Assert(target, check.DeepEquals, findTests[i].expected)
// 	}
// }
