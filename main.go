package main

import (
	"fmt"

	say_hello "github.com/samsudin47/go-module"
)

func main() {
	fmt.Println(say_hello.SayHello())
}

// cara menggunakan module dari github
// go get github.com/samsudin47/go-module

// upgrade module
// go get -u github.com/samsudin47/go-module