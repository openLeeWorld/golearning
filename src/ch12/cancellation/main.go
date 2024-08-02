package main

import (
	"cancellation/client"
	"cancellation/servers"
	"context"
	"os"
)

func main() {
	ss := servers.SlowServer()
	defer ss.Close()
	fs := servers.FastServer()
	defer fs.Close()

	ctx := context.Background()
	client.CallBoth(ctx, os.Args[1], ss.URL, fs.URL) // true나 false 넣음
}
