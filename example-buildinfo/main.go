// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"

	"github.com/platinasystems/buildinfo"
)

func main() {
	if bi := buildinfo.New(); bi.Version() != buildinfo.Unavailable {
		fmt.Println(bi)
	} else {
		fmt.Println(buildinfo.Unavailable)
	}
}
