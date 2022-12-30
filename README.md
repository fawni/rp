# rp [![Build Status](https://img.shields.io/github/actions/workflow/status/fawni/rp/build.yml?logo=github&branch=master)](https://github.com/fawni/rp/actions)


An implementation of Discord's rich presence for Linux, MacOS and Windows in Go. Fork of [hugolgst/rich-go](https://github.com/hugolgst/rich-go).

## Disclaimer

This fork is not meant to be a replacement or a production-ready library. The sole purpose of this fork is opinionated changes to suit my own workflow with no gurantee for stablility.

## Installation

Install `github.com/fawni/rp`:

```sh
$ go get -u github.com/fawni/rp
```

## Usage

create a new client

```go
c, err := rp.NewClient("DISCORD_APP_ID")
if err != nil {
	panic(err)
}
```

set the rich presence activity

```go
if err := c.SetActivity(&rpc.Activity{
	State:      "Hey!",
	Details:    "Running on rp.go!",
	LargeImage: "largeimageid",
	LargeText:  "This is the large image",
	SmallImage: "smallimageid",
	SmallText:  "And this is the small image",
}) err != nil {
	panic(err)
}
```

more details in the [example](example/main.go)

## License

[ISC](LICENSE)
