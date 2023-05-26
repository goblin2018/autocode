package main

import (
	"auto/api"
	"fmt"
)

func main() {
	scs := api.Schemas
	baseDir := api.BaseDir
	fmt.Println(baseDir)

	for _, sc := range scs {
		for _, s := range sc.Structs {
			println(s)
		}
		for _, g := range sc.Groups {
			println(g)
		}
	}
}
