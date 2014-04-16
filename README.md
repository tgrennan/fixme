Package `fixme` provides conditional printing for benignly littering code under
development with debug information and simple reminders.  Unless enabled, each
closure is a return stub; once enabled, these log:

	FILE:LINE FIXME ...

Fetch, build and install `fixme` with GO tool like this:

	go get gopkg.in/tgrennan/fixme.v0

Import this package with:

	import "gopkg.in/tgrennan/fixme.v0"

[![GoDoc](https://godoc.org/gopkg.in/tgrennan/fixme.v0?status.png)](https://godoc.org/gopkg.in/tgrennan/fixme.v0)
