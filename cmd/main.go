package main

import (
	"fmt"
	"github.com/nanwct/testGolangThirdModule"
	"time"
)

func main() {
	fmt.Println(testGolangThirdModule.GetTime())
	go func() {
		time.Sleep(10)
	}()
	fmt.Println("after select!")
}

func f1()  {
	a := "main branch"
	fmt.Println(a)
}

func f2()  {

}