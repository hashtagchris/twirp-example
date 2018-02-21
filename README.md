# twirp-example

This is an exmaple Twirp service for educational purposes.

## Try it out

First, download this repo with the Go tool:
```
go get github.com/twitchtv/twirp-example
cd $GOPATH/src/github.com/twitchtv/twirp-example
```

Next, try building the client and server binaries:
```
go build ./cmd/client
go build ./cmd/server
```

And run them. In one terminal session:
```
./server
```

And in another:
```
./client
```

In the client, you should see something like this:
```
-> % ./client
size:12 color:"red" name:"baseball cap"
```

In the server, something like this:
```% ./server
received req svc="Haberdasher" method="MakeHat"
response sent svc="Haberdasher" method="MakeHat" time="109.01µs"
```