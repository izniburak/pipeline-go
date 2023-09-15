# pipeline-go [![](https://github.com/izniburak/pipeline-go/actions/workflows/go.yml/badge.svg)](https://github.com/izniburak/pipeline-go/actions) [![PkgGoDev](https://pkg.go.dev/badge/github.com/izniburak/pipeline-go)](https://pkg.go.dev/github.com/izniburak/pipeline-go)

This package allows you to use the Pipeline pattern in your processes, and it's built upon the Chain of Responsibility (CoR) design pattern.

CoR is a behavioral design pattern that processes given data through a series of handlers. When a request reaches the pipe class, it processes the data and then forwards it to the next handler. The principle behind this pattern is straightforward.

![pipeline](/.github/static/pipeline.png)

> In summary, the Pipeline is a design pattern tailored for managing sequential modifications to an object. Imagine it as an assembly line: each station represents a pipe, and by the end of the line, you're left with a transformed object.

## Install

```bash
go get github.com/izniburak/pipeline-go
```

## Examples
```go
package main

import (
	"fmt"
	"strings"
	"github.com/izniburak/pipeline-go"
)

type UpperCasePipe struct{}

func (u *UpperCasePipe) Handle(value pipeline.PipeValue, next pipeline.PipeNext) pipeline.PipeValue {
	// get value
	text := value.(string)
	capitalized := strings.ToUpper(text)
	return next(capitalized)
}

type TrimSpacePipe struct{}

func (t *TrimSpacePipe) Handle(value pipeline.PipeValue, next pipeline.PipeNext) pipeline.PipeValue {
	// get value
	text := value.(string)
	trimmed := strings.Trim(text, " ")
	return next(trimmed)
}

func main() {
	text := "   buki.dev   "

	pipes := []pipeline.PipeInterface{
		new(UpperCasePipe),
		new(TrimSpacePipe),
	}
	result := pipeline.Send(text).Through(pipes).ThenReturn()

	fmt.Println(result) // BUKI.DEV
}
```

## Contributing

1. Fork the repo (https://github.com/izniburak/pipeline-go/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [izniburak](https://github.com/izniburak) İzni Burak Demirtaş - creator, maintainer

## License
The MIT License (MIT) - see [`license.md`](https://github.com/izniburak/pipeline-go/blob/main/license.md) for more details
