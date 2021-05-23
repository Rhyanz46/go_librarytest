package golib

import "fmt"

func Greeting(nama string) string{
	return fmt.Sprintf("hello %v", nama)
}