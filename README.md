This package provides a [WriterTo] wrapper to [runtime/debug.BuildInfo].
Usage,

```golang
	bi, err := buildinfo.New()
	if err == nil {
		_, err = bi.WriteTo(os.Stdout)
	}
```

---

*&copy; 2018-2019 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: LICENSE
[WriterTo]: https://godoc.org/io#WriterTo
[runtime/debug.BuildInfo]: https://godoc.org/runtime/debug#BuildInfo
