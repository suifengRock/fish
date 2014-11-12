package main

import (
	"fmt"
	"tools"
)

func main() {
	fmt.Println(tools.CapitalOnly(10))
	fmt.Println(tools.LowercaseOnly(10))
	fmt.Println(tools.NumberOnly(10))
	fmt.Println(tools.SpeStrOnly(10))
	fmt.Println(tools.NoSpeStr(10))
}
