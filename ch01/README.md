# Exercises

1. Take the "Hello, World!" program and run it on the Go Playground. Share a link to the code in the playground with a
coworker who would love to learn about go.

* https://go.dev/play/p/xY2BPIIxBbk

2. Add a target to the Makefile called `clean` that removes the `hello_world` binary and any other temporary files
created by `go build`. Take a look at the Go command documentation (https://oreil.ly/uqsMy) to find a go command to help
implement this.


3. Experiment with modifying the formatting in the "Hello, world!" program. Add blank lines, spaces, change indentation,
insert newlines. After making a modification, run `go fmt` to see if the formatting change is undone. Also, run
`go build` to see if the code still compiles. You can also add additional `fmt.Println` calls so you can see what
happens if you put blank lines in the middle of a function.
