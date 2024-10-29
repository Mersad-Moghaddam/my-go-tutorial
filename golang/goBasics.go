package main

import (
	"bufio"
	f "fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	// Getting Start
	f.Println("Hello Go")
	f.Print("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		f.Println("Hello", name)
	} else {
		log.Fatal(err)
	}

	// Variables
	var a, b = 1, 2.56
	var c = "this is String"
	d := 4
	f.Println(a, b, c, d)

	g := 1.25

	z := int(g)
	f.Println(z)
	// convert to int
	cV3 := "50000"
	f.Println(reflect.TypeOf(cV3))
	cV4, err := strconv.Atoi(cV3)
	f.Println(reflect.TypeOf(cV4))

	f.Println(cV4)
	// convert to String
	cV5 := 5000
	cV6 := strconv.Itoa(cV5)
	f.Println(reflect.TypeOf(cV6))
	// convert to float
	cV7 := "5000"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		f.Println(reflect.TypeOf(cV8))
	}

}
