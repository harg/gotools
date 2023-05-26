// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the append checker.

package b

func append(args ...interface{}) []int {
	println(args)
	return []int{0}
}

func userdefine() {
	sli := []int{1, 2, 3}
	sli = append(sli, 4, 5, 6)
	sli = append(sli)
}
