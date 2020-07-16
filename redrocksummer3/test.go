package main

import (
	"fmt"
	"strings"
)

func main()  {
	str  := `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"Name":"电影","Pname":"summer","Num":1,"Where":"红岩网校"}`
	str0 :=strings.Split(str, `"`)
	for i := 0; i < len(str0); i++ {
		fmt.Println(i, str0[i])
	}
}

