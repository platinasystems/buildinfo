This package provides a [runtime/debug.BuildInfo] [Formatter].
Usage,

```golang
	if bi := buildinfo.New(); bi.Version() != buildinfo.Unavailable {
		fmt.Println(bi)
	} else {
		fmt.Println(buildinfo.Unavailable)
	}
```

---

*&copy; 2018-2019 Platina Systems, Inc. All rights reserved.
Use of this source code is governed by this BSD-style [LICENSE].*

[LICENSE]: LICENSE
[runtime/debug.BuildInfo]: https://godoc.org/runtime/debug#BuildInfo
[Formatter]: https://golang.org/pkg/fmt/#Formatter
