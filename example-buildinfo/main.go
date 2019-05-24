// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/platinasystems/buildinfo"
)

func main() {
	bi, err := buildinfo.New()
	if err == nil {
		_, err = bi.WriteTo(os.Stdout)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
