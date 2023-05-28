# medium.rip
This is an alternative frontend for [Medium](https://medium.com) written in Go. I was inspired by the [Scribe](https://scribe.rip) project, but wanted a few different things, and I did not know Crystal.

## Building
Please feel free to self host and run this on your own. I only ask that you contribute any changes back upstream.

### Dependencies
 - [Go](https://go.dev/dl/) (at least v1.20)
 - [Node.js](https://nodejs.org/en/download)
 - [PNPM](https://pnpm.io/installation)

### Building
First, build the frontend
```sh
cd frontend
pnpm i
pnpm run build
```

Then, build the binary. The frontend static files will be embedded in the binary using `go:embed`.
```sh
go mod download
go build .
```

You should now have a static binary called `medium.rip` that is self contained.

### Dockerfile
You can alternately build and run via `docker`
```sh
docker build -t medium-rip .
docker run -p 3000:3000 -e PORT=3000 medium-rip
```

## Licensing
Dual licensed under Apache 2.0 or MIT at your option.
