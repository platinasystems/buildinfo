// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build !go1.12

package buildinfo

import (
	"fmt"
	"io"
)

type BuildInfo struct{}

func New() (BuildInfo, error) {
	return BuildInfo{}, nil
}

func (bi BuildInfo) WriteTo(w io.Writer) (int64, error) {
	n, err := fmt.Fprintln(w, "unavailable")
	return int64(n), err
}
