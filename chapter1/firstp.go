package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(2+2 != 5)
	fmt.Println(reflect.TypeOf(4.0))
}

/*Go is statically typed -> knows types of vars before running*/
