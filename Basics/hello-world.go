package main

import (
	"fmt"
	"time"
)

const GoCon1 int16 = 255    // globally exported
const goCon2 float32 = 34.3 // local to file

// STRUCTS
// type Career struct {
// 	job    string
// 	salary uint32
// }
// type Person struct {
// 	name      string
// 	age       uint8
// 	gender    string
// 	isMarried bool
// 	Career
// }
// func (person *Person) updateName(name string) {
// 	person.name = name
// }
// func (person Person) yearOfBirth() uint32 {
// 	return 2021 - uint32(person.age)
// }

// INTERFACES
// type Figure interface {
// 	Area() uint
// 	Perimeter() uint
// }
// type Square struct {
// 	length uint
// }
// type Rectangle struct {
// 	length uint
// 	width  uint
// }
// func (s *Square) Area() uint {
// 	return s.length * s.length
// }
// func (s *Square) Perimeter() uint {
// 	return 4 * s.length
// }
// func (r *Rectangle) Area() uint {
// 	return r.length * r.width
// }
// func (r *Rectangle) Perimeter() uint {
// 	return 2 * (r.length + r.width)
// }

// function to read from CHANNEL
func sumChan(ch chan uint16) {
	var vals uint16 = 0
	for val := range ch {
		vals += val
	}
	fmt.Println("Sum of points:", vals)
	fmt.Println("...Done!")
}

func main() {

	// CHANNEL Example
	myChan := make(chan uint16)
	// my3D := [5][2]int{{54, 45}, {67, 76}, {89, 98}, {21, 12}, {43, 34}}
	myArr := []uint16{76, 80, 88, 90, 67}
	fmt.Println("Writing to Channel...")
	go sumChan(myChan)
	func() {
		for i := 0; i < len(myArr); i++ {
			// fmt.Printf("\n%d - %d", i+1, my3D[i])
			// fmt.Println(myArr[i])
			myChan <- myArr[i]
		}
		close(myChan)
	}()

	time.Sleep(100 * time.Millisecond)

	// GOROUTINE Example
	// fmt.Println("Starting...")
	// go countMe(3, "Alex")
	// countMe(3, "Kitaa")
	// time.Sleep(100 * time.Millisecond)
	// fmt.Println("...Done!")

	// // INTERFACES Usage
	// mySquare := Square{3}
	// myRect := Rectangle{3, 4}
	// myFigures := []Figure{&mySquare, &myRect}
	// for _, Fig := range myFigures {
	// 	fmt.Println(Fig.Area(), Fig.Perimeter())
	// }
	// fmt.Printf("Area of Square: %d\nArea of Rectangle: %d\n", mySquare.Area(), myRect.Area())
	// fmt.Printf("Perimeter of Square: %d\nPerimeter of Rectangle: %d\n", myFigures[0].Perimeter(), myFigures[1].Perimeter())

	// STRUCTS Usage
	// var personsList []Person = []Person{}
	// me := Person{"Alex", 25, "M", false, Career{"Sys Admin", 50000}}
	// personsList = append(personsList, me)
	// fmt.Println(me)
	// fmt.Println(personsList)
	// newName := "Kitaa"
	// me.updateName(newName)
	// personsList[0] = me
	// fmt.Println(me)
	// fmt.Println(personsList)
	// fmt.Println(me.job)
	// year := me.yearOfBirth()
	// fmt.Println(year)

	// CONSTANTS
	// fmt.Printf("Constant goCon2 contains %.1f", goCon2)
	// // internal local context
	// const goCon2 float32 = 32.3
	// fmt.Printf("\nNow constant goCon2 contains %.1f\n", goCon2)

	// BIT ABOUT FUNCTIONS
	// anonymous function
	// func(name string) {
	// 	fmt.Printf("Hello %q", name)
	// }(name)
	// function as variable
	// hello := func(name string) {
	// 	fmt.Printf("Hello %q", name)
	// }
	// hello(name)
	// fmt.Printf("\nhello is a %T", hello)

	// COMMON DATA TYPES
	// var fName string = "Kitaa"
	// var age uint8 = 25
	// var weight float32 = 64.7
	// var gender byte = 'M'
	// var isAlive bool = true
	// whoami := "sudo"
	// fmt.Println("Hello World!!")
	// fmt.Println("My name is ", fName, " i am ", age, " years old ")
	// fmt.Printf("I weigh %.1f kgs and my gender is %c, am i alive? %t", weight, gender, isAlive)
	// var root = fmt.Sprintf("I am %s %q", whoami, "root")
	// fmt.Println("\n", root)

	// IO READING USER INPUT FROM STDIN
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Username: ")
	// scanner.Scan()
	// var name string = scanner.Text()
	// fmt.Println("Hello", name)
	// BUFIO READER - 4k per single read
	// userIn := bufio.NewReader(os.Stdin)
	// fmt.Println("Age: ")
	// value, err := userIn.ReadString('\n')
	// if err != nil {
	// 	log.Fatal("Unable to read input")
	// }
	// toRoot, err := strconv.ParseInt(strings.TrimSpace(value), 10, 8)
	// if err != nil {
	// 	log.Fatal("Input was not a number")
	// }
	// fmt.Println(toRoot + 10)
	// BUFIO SCANNER - 64k per single read
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Username: ")
	// scanner.Scan()
	// name := scanner.Text()

	// CONVERTING, FORMATTING & PRINTING DATA TO STDOUT
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Port to scan: ")
	// scanner.Scan()
	// var input = scanner.Text()
	// port, err := strconv.ParseUint(input, 10, 16)
	// if err != nil {
	// 	log.Fatalf("Invalid value for port :%q", err)
	// 	return
	// }
	// fmt.Printf("Scanning %T, %d", port, port)

	// ARITHMETICS
	// var num1 int32 = 24
	// var num2 int64 = 78
	// results := num2 != int64(num1)
	// fmt.Println(results)
	// var l8r1 string = "z"
	// var l8r2 string = "Z"
	// comp := l8r1 >= l8r2
	// fmt.Printf("is (%x)%q greater than (%x)%q ? %t", l8r1, l8r1, l8r2, l8r2, comp)
	// logic := !false && !!true || 1 > 'b'
	// fmt.Printf("!false && !!true || %d > %q(%x) ? %t", 1, 'b', 'b', logic)

	// DATA STRUCTURES
	// myArr := [6]int{23, 34, 55, 56, 32, 67}
	// fmt.Printf("myArr is of type %T and contains %v", myArr, myArr)
	// for i := 0; i < len(myArr); i++ {
	// 	fmt.Printf("\n%d - %d", i+1, myArr[i])
	// }
	// my3D := [5][2]int{{54, 45}, {67, 76}, {89, 98}, {21, 12}, {43, 34}}
	// fmt.Printf("my3D Array of type %T and length %d", my3D, len(my3D))
	// fmt.Printf("\n\t\tTotal elements: %d", len(my3D)*len(my3D[0]))
	// fmt.Println("\n\t\tContents: ")
	// for i := 0; i < len(my3D); i++ {
	// 	fmt.Printf("\n%d - %d", i+1, my3D[i])
	// }
	// var sum int
	// for i := 0; i < len(my3D); i++ {
	// 	for j := 0; j < len(my3D[i]); j++ {
	// 		sum += my3D[i][j]
	// 	}
	// }
	// fmt.Printf("\n\t\tSum of Elements: %d", sum)
	// var mySlice []int = []int{1, 2, 3, 4, 5}
	// fmt.Println(mySlice)
	// fmt.Printf("Capacity: %d & Length: %d of type %T", cap(mySlice), len(mySlice), mySlice)
	// for _, el := range mySlice {
	// 	fmt.Printf("\n%d", el)
	// }
	// myMap := make(map[string]int32)
	// myMap = map[string]int32{
	// 	"Age":    25,
	// 	"Height": 760,
	// 	"Weight": 68,
	// }
	// fmt.Println(myMap["Age"])
	// myMap["Age"] = 24
	// fmt.Println(myMap["Age"])
	// fmt.Println(myMap)
	// myMap["Salary"] = 50000
	// fmt.Println(myMap)
	// delete(myMap, "Salary")
	// fmt.Println(myMap)

	// POINTERS & DEREFERENCING
	// myName := "Kitaa"
	// var point *string = &myName
	// fmt.Printf("Variable myName at address %p contains %q  ", point, myName)
	// fmt.Printf("\nPointer point at address %p which contains address %p that contains %q", &point, point, *point)

	// BUFIO TO OTHER FUNCTION
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Username: ")
	// scanner.Scan()
	// name := scanner.Text()
	// newUser(name)

}

// SAMPLE FUNCTION
// func newUser(name string) {
// 	fmt.Printf("Hello %q", name)
// }

// GOROUTINE FUNCTION
// func countMe(times uint, text string) {
// 	for i := 0; i < int(times); i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println("\t", text, "-", i+1)
// 	}
// }
