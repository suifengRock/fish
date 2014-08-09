package main

import (
	"fmt"
	//"github.com/go-martini/martini"
	"os"
)

// func main() {
// 	m := martini.Classic()
// 	m.Get("/", func() string {
// 		return "Hello world!"
// 	})
// 	m.Run()
// }
//
func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic("open failed!")
	}
	defer f.Close()

	buff := make([]byte, 1024)
	for n, err := f.Read(buff); err == nil; n, err = f.Read(buff) {
		fmt.Print(string(buff[:n]))
	}
	if err != nil {
		panic(fmt.Sprintf("Read occurs error: %s", err))
	}
}
