package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type teststruct struct {
	a int
}

func main() {
	bytes, err := json.Marshal(nil)
	if err != nil {
		fmt.Println("NIL")
		return
	}
	fmt.Println(string(bytes))

	var test teststruct
	err = json.Unmarshal(bytes, &test)
	if err != nil {
		fmt.Println("Unmarshal error")
		return
	}
	fmt.Printf("output ok, %#v\n", test)

	sizeof := unsafe.Sizeof(test)
	fmt.Printf("%d", sizeof )

}
