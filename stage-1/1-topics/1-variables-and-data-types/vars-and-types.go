package main

import "fmt"

var (
	// some useful data types at their "default" values (some implicit types depends on the platform e.g.,
	//int can be int32 or 64) and is possible to omit the Type and go do it implicit

	// explicit types
	str          string     // = "just a text" //string
	anyInteger   int        // = -10           //int (int8..int64)
	onlyPositive uint       // = 10            //uint (uint8...uint64)
	decimal      float64    // = 3.1415        //float (float32 or float64)
	boolean      bool       // = true          // bool (true or false)
	complexN     complex128 // = 5 + 12i       //complex (complex64 or complex128)
	runes        rune       //  alias for int32
	bytes        byte       //  alias for uint8
	uptr         uintptr    //  uintptr (holds pointer values as unsigned integers

	// implicit types
	str1 = "just a text" //string
)

// "enums"
const (
	enum0 = iota //0
	enum1        //1
	enum2        //2
)

func main() {
	fmt.Printf(
		"string: %s\nint: %v\nuint: %v\nfloat: %f\nboolean %v\ncomplex: %v\nrune :%v\nbyte :%v\nuintptr :%v\n", str,
		anyInteger,
		onlyPositive, decimal, boolean, complexN, runes, bytes, uptr)
}
