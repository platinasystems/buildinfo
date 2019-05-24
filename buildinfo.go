// Copyright © 2019 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// +build go1.12

// This package provides a WriterTo wrapper to runtime/debug.BuildInfo.
// Usage,
//
//	bi, err := buildinfo.New()
//	if err == nil {
//		_, err = bi.WriteTo(os.Stdout)
//	}
package buildinfo

import (
	"errors"
	"fmt"
	"io"
	"runtime/debug"
)

type BuildInfo struct {
	*debug.BuildInfo
}

func New() (BuildInfo, error) {
	if bi, ok := debug.ReadBuildInfo(); ok {
		return BuildInfo{bi}, nil
	}
	return BuildInfo{}, errors.New("can't read BuildInfo")
}

func (bi BuildInfo) WriteTo(w io.Writer) (n int64, err error) {
	acc := func(i int, nextErr error) {
		n += int64(i)
		if err == nil {
			err = nextErr
		}
	}
	pmod := func(m *debug.Module) {
		acc(fmt.Fprint(w, m.Path, "@"))
		if m.Replace != nil {
			acc(fmt.Fprint(w, m.Replace.Path))
			if len(m.Replace.Version) > 0 {
				acc(fmt.Fprint(w, "@", m.Replace.Version))
			}
		} else {
			acc(fmt.Fprint(w, m.Version))
		}
		acc(fmt.Fprintln(w))
	}
	if bi.BuildInfo != nil {
		pmod(&bi.Main)
		for _, dep := range bi.Deps {
			acc(fmt.Fprint(w, "\t"))
			pmod(dep)
		}
	} else {
		err = errors.New("empty BuildInfo")
	}
	return
}
