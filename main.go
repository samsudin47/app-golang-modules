package main

import (
	"fmt"

	say_hello "github.com/samsudin47/go-module/v3"
)

func main() {
	fmt.Println(say_hello.SayHello("Samsudin"))
}

// cara menggunakan module dari github
// go get github.com/samsudin47/go-module

// upgrade module
// go get -u github.com/samsudin47/go-module