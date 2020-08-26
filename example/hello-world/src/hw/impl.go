package hw

import (
	"fmt"
	//"math"
)

func do_add(a int64, b int64) int64
func get_fcw() uint16
func get_mxcsr() uint16

func HelloWorld(done chan bool) {
	fmt.Println("Hello World!")

        fmt.Printf("[ENTRY] Control Word: %#x\n", get_fcw())
        fmt.Printf("[ENTRY] MXCSR: %#x\n", get_mxcsr())

	fmt.Println("\nAdd(1,2) is", do_add(1,2))

        fmt.Printf("[EXIT] Control Word: %#x\n", get_fcw())
        fmt.Printf("[EXIT] MXCSR: %#x\n", get_mxcsr())

	done <- true
}
