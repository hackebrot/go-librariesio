# go-librariesio

[![GoDoc Reference][godoc_badge]][godoc]
[![Build Status][travis_badge]][travis]

go-librariesio is a Go client library for accessing the
[libraries.io][libraries.io] API.


## Installation

``go get github.com/hackebrot/go-librariesio``


## libraries.io API

Connecting to the [libraries.io API][api] with **go-librariesio** requires
a [private API key][api_key].

## Usage

```go
// Create new API client with your API key
c := librariesio.NewClient("... your API key ...")

// Create a new context (with a timeout if you want)
ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
defer cancel()

// Request information about a project using the client
project, _, err := c.GetProject(ctx, "pypi", "cookiecutter")

if err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
}

// All structs for API resources use pointer values.
// If you expect fields to not be returned by the API
// make sure to check for nil values before dereferencing.
fmt.Fprintf(
    os.Stdout,
    "name: %s\nversion: %s\nlanguage: %s\n",
    *project.Name,
    *project.LatestReleaseNumber,
    *project.Language,
)
```

## License

Distributed under the terms of the [MIT License][MIT], **go-librariesio** is
free and open source software.


## Contributing

Contributions are welcome, and they are greatly appreciated! Every
little bit helps, and credit will always be given.

Please check out this [guide][contributing] to get started!


## Code of Conduct

Please note that this project is released with a
[Contributor Code of Conduct][Code of Conduct].

By participating in this project you agree to abide by its terms.


[api_key]: https://libraries.io/account
[api]: https://libraries.io/api
[Code of Conduct]: CODE_OF_CONDUCT.md
[contributing]: CONTRIBUTING.md
[godoc_badge]: https://img.shields.io/badge/go-documentation-blue.svg?style=flat
[godoc]: https://godoc.org/github.com/hackebrot/go-librariesio (See GoDoc Reference)
[libraries.io]: https://libraries.io/
[MIT]: LICENSE
[travis_badge]: https://img.shields.io/travis/hackebrot/go-librariesio.svg?style=flat
[travis]: https://travis-ci.org/hackebrot/go-librariesio (See Build Status on Travis CI)
