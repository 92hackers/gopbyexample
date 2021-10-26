# Go+ by Tutorials

Content and build toolchain for [Go+ by Tutorials](https://github.com/goplus/gop#tutorials), a site that teaches Go+ via annotated tutorials.

## Overview

The Go+ by Tutorials site is built by extracting code and
comments from source files in `tutorials` and rendering
them via the `templates` into a static `public`
directory. The programs implementing this build process
are in `tools`, along with dependencies specified in
the `go.mod`file.

The built `public` directory can be served by any
static content system.

## Building

To build the site you'll need Go installed. Run:

```console
$ tools/build
```

To build continuously in a loop:

```console
$ tools/build-loop
```

To see the site locally:

```
$ tools/serve
```

and open `http://127.0.0.1:8000/` in your browser.

## License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](http://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

## Thanks

1. Thanks to [gobyexample](https://github.com/mmcgrana/gobyexample) project, this project is built upon it.

2. Thanks to [Jeremy Ashkenas](https://github.com/jashkenas) for [Docco](http://jashkenas.github.com/docco/), which
inspired this project.
