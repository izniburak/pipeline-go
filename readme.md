# pipeline-go [![](https://github.com/izniburak/pipeline-go/workflows/build/badge.svg)](https://github.com/izniburak/pipeline-go/actions) [![PkgGoDev](https://pkg.go.dev/badge/github.com/izniburak/pipeline-go)](https://pkg.go.dev/github.com/izniburak/pipeline-go)

This package allows you to use the Pipeline pattern in your processes, and it's built upon the Chain of Responsibility (CoR) design pattern.

CoR is a behavioral design pattern that processes given data through a series of handlers. When a request reaches the pipe class, it processes the data and then forwards it to the next handler. The principle behind this pattern is straightforward.

![pipeline](/.github/static/pipeline.png)

> In summary, the Pipeline is a design pattern tailored for managing sequential modifications to an object. Imagine it as an assembly line: each station represents a pipe, and by the end of the line, you're left with a transformed object.

## Install

```bash
go get github.com/izniburak/pipeline-go
```

## Examples


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
