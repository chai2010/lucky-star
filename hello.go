// Copyright 2018 chaishushan@gmail.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	. "github.com/chai2010/lucky-star"
)

func main() {
	luckyStar := MakeLuckyStarFromList("wasm",
		[2]string{"chai2010", "chai2010"},
		[2]string{"ending", "ending"},
	)

	fmt.Println(luckyStar)
	// output: ending

	luckyStar = MakeLuckyStar("wasm", map[string]string{
		"chai2010": "chai2010",
		"ending":   "ending",
	})

	fmt.Println(luckyStar)
	// output: ending
}
