# golang

First steps:

https://golang.org/doc/code.html

Workspaces: https://golang.org/doc/code.html#Workspaces

Each folder inside src, compiles as a different executable, with the folder name 
as it's name. For example:

    // installs with 'go install src/hello'
    src/hello/hello.go

    // installs with 'go install src/golang'
    src/golang/hello/hello.go



We need to export some paths:

export GOPATH=$PATH:~/dev/golang
export PATH=$PATH:$GOPATH/bin
export GOBIN=~/dev/golang/bin