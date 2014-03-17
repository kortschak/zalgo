// Copyright ©2013 Dan Kortschak. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zalgo_test

import (
	"fmt"
	"github.com/kortschak/zalgo"
	"os"
)

const invoke = `To invoke the hive-mind representing chaos.

Invoking the feeling of chaos.

With out order.

The Nezperdian hive-mind of chaos. Zalgo.

He who Waits Behind The Wall.

ZALGO!
`

func Example_1() {
	z := zalgo.NewCorrupter(os.Stdout)

	z.Zalgo = func(n int, r rune, z *zalgo.Corrupter) bool {
		z.Up += 0.1
		z.Middle += complex(0.01, 0.01)
		z.Down += complex(real(z.Down)*0.1, 0)
		return false
	}

	z.Up = complex(0, 0.2)
	z.Middle = complex(0, 0.2)
	z.Down = complex(0.001, 0.3)

	fmt.Fprintln(z, invoke)

	// Output:
	// Eternal happiness.
}
