# rp.go

An implementation of Discord's rich presence in Golang for Linux, macOS and Windows. Fork of [hugolgst/rich-go](https://github.com/hugolgst/rich-go)

## Installation

Install `github.com/x6r/rp`:

```sh
$ go get -u github.com/x6r/rp
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

	if err := c.SetActivity(rpc.Activity{
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
