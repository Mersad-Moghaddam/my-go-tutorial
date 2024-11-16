package main

import (
	"bufio"
	"fmt"
	f "fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// seedSec = time.Now().Unix()
	// r := rand.New(rand.NewSource(seedSec))
	// randomNum := r.Intn(100)
	// fmt.Println()

	fmt.Println("Enter Your Age : ")
	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal((err))
	}
	age = strings.TrimSpace(age)
	myAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	Vote(myAge)

	var p = fmt.Println

	p("Enter Your Age :")
	input, err := reader.ReadString('\n')
	iage, err := strconv.Atoi(strings.TrimSpace(input))

	if err == nil {
		if iage < 5 {
			p("too young for school")
		} else if iage == 5 {
			p("Go to Kindergarten")
		} else if iage > 5 && iage < 18 {
			p("Go to school")
		} else if iage == 18 {
			p("Go to College")
		} else {
			p("Go to University")
		}
	} else {
		log.Fatal(err)
	}

	// Getting Start
	f.Println("Hello Go")
	f.Println("What is your age?")

	f.Print("What is your name?")
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
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mySlice)

	mySlice2 := []string{"mersad", "ahmed", "mohamed"}
	for i := 0; i < len(mySlice2); i++ {
		fmt.Println(mySlice2[i])
	}

	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
}

func Vote(iAge int) {
	myAge := iAge
	if myAge >= 18 {
		f.Println("You can vote")
	} else {
		f.Println("You can't vote")
	}
}
