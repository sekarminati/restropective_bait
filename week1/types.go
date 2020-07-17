package main

import (
	"fmt"
	//"math/cmplx"
	// "time"
)

// 3.1
// var (
// 	goIsFun bool       = true
// 	maxInt  uint64     = 1<<64 - 1
// 	complex complex128 = cmplx.Sqrt(-5 + 12i)
// )

// 3.3
// func timeMap(y interface{}) {
// 	z, ok := y.(map[string]interface{})
// 	if ok {
// 		z["updated_at"] = time.Now()
// 	}
// }

// 3.3.2
type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// function used to implement the Stringer interface
func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())

	}
}

// func main() {
// 	// 3.1
// 	// const f = "%T(%v)\n"
// 	// fmt.Printf(f, goIsFun, goIsFun)
// 	// fmt.Printf(f, maxInt, maxInt)
// 	// fmt.Printf(f, complex, complex)

// 	//3.3
// 	// foo := map[string]interface{}{
// 	// 	"Matt": 42,
// 	// 	"Aca":  true,
// 	// }
// 	// timeMap(foo)
// 	// fmt.Println(foo)

// 	s := &fakeString{"Ceci n'est pas un string"}
// 	printString(s)
// 	printString("Hello, Gophers")
// }
