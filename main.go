/*
	The file doesn't need to be called main, but .go is the extension
	There are a few options for how to build, complile and run, but we want to run 'go run . ' in our command line
*/

// Start with the package identifier
package main

// If you start working with a library (and save the file) it seems to add it automatically
import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"

	// Adding my own package
	"demo/coffeeshop"
)

// MARK: Main
// Entry point is determined by a function called main
func main() {
	println("Hello, Gophers!")     // This is a built in print function, but isn't what we usually are going to use (good for keeping less dependencies, debugging)
	fmt.Println("Hello, Gophers!") // This is the imported version we typically want to use, for formatted strings and such
	coffeeshop.Operate()
}

// MARK: Simple Data Types
func mod3() {
	// Variables are strongly typed
	fmt.Println("\n-- Variables --")
	var myName string          // Declare variable
	var yrName string = "Mike" // Declare and initialize
	var hsName = "Steve"       // Initialize with inferred type (this is the least common usage, the others all have their times)
	hrName := "Sarah"          // Shorthand declaration

	// A variable that is declared but not used is a runtime error
	fmt.Println(myName, yrName, hsName, hrName)

	// Type conversions (Go's goal is clear over clever, it never allows for assumptions that might be wrong)
	var i int = 32
	var f float32
	// We can't do f = i because that's an implicit conversion and would error
	f = float32(i) // This is valid because it's explicit
	fmt.Println(f)
	// If you use the shorthand declaration and there are 2 possible data types, Go will choose a default (like float64 for floats)

	// Arithmetic operators
	fmt.Println("\n-- Operations and Comparisons --")
	a, b := 10, 5  // Go allows for multiple declarations on one line!!
	c := a + b     // 15 - Addition => c will be the same data type as a and b
	c = a - b      // 5 - Subtraction
	c = a * b      // 50 - Multiplication
	c = a / b      // 2 - Division
	c = a / 3      // 3 - Integer division => it will lose the remainder
	c = a % 3      // 1 - Modulus => gives the remainder
	d := 7.0 / 2.0 // 3.5
	fmt.Println(c, d)
	fmt.Println(a + b)                   // You can do all of this like normal
	fmt.Println(float32(a) / float32(b)) // Would cast our values and give us a proper decimal result like we may have expected (if a and b were different values)

	// Comparisons
	// c = a == b didn't work because c was an int and now we're asking it to be a bool
	z := a == b // false - Equality => Errors don't support the equality operator
	z = a != b  // true - Inequality
	// The following typically don't work for non-numeric types
	z = a < b  // false - less than
	z = a <= b // false - less than or equal to
	z = a > b  // true - greater than
	z = a >= b // true - greater than or equal to
	fmt.Println(z)

	// Constants
	fmt.Println("\n-- Constants --")
	const e = 42 // Implicitly typed, has to be declared and initialized at the same time because it can't change
	const g string = "hello, world"
	const h = e // A constant can be assigned to another constant, but a variable can't be assigned to a constant
	// We can group constants to reduce redundancy!
	const (
		j = true
		k = 3.14
		l         // This looks weird but if a value isn't provided to a constant, the constant inherits the value above it
		m = 2 * 5 // Constant expression
		n = "hello, " + "world"
		// You can't assign a function result as a constant because it can't be evaluated at compile time
	)
	const o = iota // 0 => Special to assigning constants, specifically in groups
	const (
		p = iota     // 0
		q            // 1
		r = 3 * iota // 6
	)
	// Iota increments within a const group and resets when you create a new group.
	// Iota is relative to the constants position in the group, not the usage of iota
	fmt.Println(e, g, h, j, k, l, m, n, o, p, q, r)

	// Pointers
	fmt.Println("\n-- Pointers --")
	s := "foo"
	t := &s         // Create a pointer that only points to strings
	fmt.Println(t)  // Prints the memory address
	fmt.Println(*t) // Prints the value
	*t = "bar"      // Dereference the pointer, here we're updating s to hold this new value
	fmt.Println(*t) // Prints the updated value!
	u := new(int)   // Creates pointer to anonymous variable
	fmt.Println(*u)
	// Use copies whenever possible, sharing memory risks race conditions
	t = new(string) // A second valid way to create a pointer
}

// MARK: Creating Programs
func mod4() {
	// Module 4 CLI Application
	fmt.Println("\n --- Module 4 ---")
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin) // Stdin alone reads one character at a time. What we've done here is wrapped it with bufio to read more useful chunks
	st, _ := in.ReadString('\n')    // Read until you reach your specified delimiter, single quotes specify that we're delimiting a single character, 2 returned variables (string and error)
	st = strings.TrimSpace(st)
	st = strings.ToUpper(st)
	fmt.Println(st + "!")

	// Module 4 Web Service
	http.HandleFunc("/", Handler)              // Register the function as the back controller
	http.ListenAndServe("localhost:3000", nil) // Start the web service to be listening, normally you need to give the IP and the port (local host can be assumed). Nil is the second arg because Go will provide the front handler for us

}

// Module 4 Web Service (this is the back controller)
func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt") // Open reads a file for reading, again 2 returned variables (file object and error)
	io.Copy(w, f)                 // Copy lets us copy from a read source (like a file) to a write source
}

// MARK: Aggregate Data Types
func mod5() {
	// Module 5
	fmt.Println("\n --- Module 5 ---")
	var arr [3]int   // Array of 3 ints
	fmt.Println(arr) // [0 0 0]
	arr = [3]int{1, 2, 3}
	fmt.Println(arr[1]) // Print an index
	arr[1] = 99         // Update the value at an index
	fmt.Println(arr)
	fmt.Print(len(arr)) // Prints the length
	// The following shows that arrays copy data, not share memory
	arr2 := [3]string{"foo", "bar", "baz"}
	arr3 := arr2
	fmt.Println(arr3)
	arr2[0] = "quux"
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr2 == arr3) // false - arrays are comparable

	// Slices
	var sl []int // Notice how there's no size initialized
	fmt.Println(sl)
	sl = []int{1, 2, 3}
	fmt.Println(sl[1])
	sl[1] = 99
	fmt.Println(sl)            // Creating, accessing, printing, updating are all the same as arrays
	sl = append(sl, 5, 10, 15) // Make the slice bigger (don't forget to add the original chunk in!)
	fmt.Println(sl)
	sl = slices.Delete(sl, 1, 3) // The slices package is part of the experimental standard package (golang.org/x/exp/slices)
	fmt.Println(sl)
	// The key difference is how updating a value that's been "copied". Because you're referencing an array structure, when the value is updated the change affect both variables

	// Maps
	var m1 map[string]int
	fmt.Println(m1)
	m1 = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m1)
	fmt.Println(m1["foo"]) // Look up value
	m1["bar"] = 99         // Update value
	delete(m1, "foo")      // Remove entry from map
	m1["baz"] = 418        // Add entry to map, simply use a not-yet-used key
	fmt.Println(m1)
	fmt.Println(m1["foo"]) // Now that foo has been deleted, this will print out 0 because that's the default value
	v, ok := m1["foo"]     // If it's important to know whether that's a non-existant value or an important 0, we use this. If the value v comes out, ok will be true or false
	fmt.Println(v, ok)
	// Maps are copied by reference, but if you need them to be independednt you can use the experimental package maps.Clone
	// Maps are not comparable
	menu := map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai", "Chai Latte"}, // Notice here that we do want the comma line endings, because it's a multi line declaration
	}
	fmt.Println(menu)
	fmt.Println(menu["coffee"])
	menu["other"] = []string{"Hot Chocolate"}
	fmt.Println(menu)
	// Notice that maps are NOT ORDERED, the ordering is non-deterministic and not done by order of entry or alphabetical
	delete(menu, "tea")
	fmt.Println(menu)
	fmt.Println(menu["tea"])
	va, ok := menu["tea"]
	fmt.Println(va, ok)
	m2 := menu
	m2["coffee"] = []string{"Coffee"}
	menu["tea"] = []string{"Hot Tea"}
	fmt.Println(menu)
	fmt.Println(m2)

	// Structs
	var ah struct { // Declare anonymous struct
		name string
		id   int
	}
	fmt.Println(ah)
	ah.name = "Arthur"     // Update value of a field
	fmt.Println(ah.name)   // Access value of a field
	type myStruct struct { // Custom type
		name string
		id   int
	}
	var ba myStruct
	fmt.Println(ba)
	ba = myStruct{
		name: "Bob",
		id:   42}
	ma := ba
	ma.name = "Celine"
	fmt.Println(ba, ma) // Notice that when we changed the value in ma it didn't change ba, so it's a copy not a share
	fact := ba == ma    // false - structs are comparable
	fmt.Println(fact)
}

// MARK: Looping
func mod6() {
	// Infinite loop for {...}
	i := 1
	for {
		fmt.Println(i)
		i += 1
		break
	}

	// Loop til condition for condition {...}
	j := 1
	for j < 3 {
		fmt.Println(j)
		j++
	}
	fmt.Println("Done!")

	// Counter-based loop for initializer; test; post clause {...}
	for k := 1; k < 3; k++ {
		fmt.Println(k)
	}
	fmt.Println("Done!")

	// Looping through collections (there are 3 versions of this)
	arr := [3]int{101, 102, 103}
	// for key, value := range collection {...} (collection can be array, slice or map)
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Done!")

	// for key := range collection {...} (returns the key/index)
	// for _, value := range collection {...} (ignores the key/index and only returns the value)
}

// MARK: Branching Logic
func mod7() {
	// Goto statements
myLabel:
	fmt.Println("Back here")

	// if test {...}  else if test {...}  else {...}
	// if initializer; test {...}
	i := 5
	if i < 5 { // This line can also be i := 5; i < 5
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else { // Notice how it shares a line with the end bracket before it
		fmt.Println("i is at least 10")
	}
	fmt.Println("After the if statement")

	/* switch test {
	case expression1:
		...
	case expression2, epxression3:
		...
	default:
		...
	}
	*/
	i = 5
	switch i { // we can also do switch i = 5; i {
	case 1:
		fmt.Println("First case")
	case 2 + 3, 2*i + 3:
		fmt.Println("Second case")
	default:
		fmt.Println("Default case")
	}

	// The only difference between a switch and a logical switch is the logic of the condiiton
	switch i := 8; true { // true can also be left as implied
	case i < 5:
		fmt.Println("i is less than 5")
	case i < 10:
		fmt.Println("i is less than 10")
	default:
		fmt.Println("i is greater than 10")
	}

	// Deferred Functions
	fmt.Println("first")
	defer fmt.Println("defer this")
	fmt.Println("second")
	defer fmt.Println("defer that") // Deferred functions are LIFO (last in first out)
	/* The above will print as
		first
		second
		defer that
		defer this
	It prints in this order because of the typical use cases of defered statements, for example databases
	A database is a resource, we need to make sure release resources when we're done with that
	Now all of that is relevant because when we release resources we are doing it in the opposite order to how we opened them (like closing brackets),
	and what we can do is defer the closing immediately after we open it so that we don't forget to do it
	*/
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	// Go statements
	goto myLabel
}

// Panic demo
func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

// MARK: Interfaces
/*
type Reader interface { // Annonymous interface
	Read([]byte) (int, error)
}

type File struct {
}
func (f File) Read(b []byte) (n int, err error)

type TCPConn struct {
}
func (t TCPConn) Read(b []byte) (n int, err error)

var f File
var t TCPConn

var r Reader
r = f
r.Read()
r = t
r.Read()

// Type Assertion
var f2 File = r // This wouldn't work because Go can't guaruntee the underlying type of the interface
f2 = r.(File) // This is a type assertion, and if you're wrong it will panic
f2, ok := r.(File) // This is also a type assertion, but if you are unsure it won't panic and will tell if or if not a file was present and f2 will not be changed here

// Type Switch (if you want to check multiple types, as opposed to just one)
switch v := r.(type) {
case File:
	// v is now a file
case TCPConn:
	// v is now a TCPConn object
default:
	// If no type was matched, make v this type of object
}
*/

// Should go at type of project
type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

func mod9() {
	var p printer
	p = user{username: "adent", id: 42}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee",
		prices: map[string]float64{"small": 1.65,
			"medium": 1.80,
			"large":  1.95,
		},
	}
	fmt.Println(p.Print())

	u, ok := p.(user)
	fmt.Println(u, ok) // This triggers a panic, well not when we add the ok to control the panic
	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	switch v := p.(type) {
	case user:
		fmt.Println("Found a user:", v)
	case menuItem:
		fmt.Println("Found a menu item:", v)
	default:
		fmt.Println("I'm not sure what this is")
	}
}

// MARK: Generics
func mod9_2() {
	testScores := []float64{
		87.3,
		105,
		63.5,
		27,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c) // Should have different memory addresses but same result
	// This is great and all, but what if we change from Float64 to Float32, or what if we want to clone strings...
	// We'd need a whole new function and then have to call the right function based on the input and ugh
}

// V is the name we're giving our generic, and that square bracket after clone is telling Go what is considered V
func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

// Basically you can't compare any type to any type, so K can be anything as long as it's comparable
func compClone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Create custome type constraints
func mod9_3() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

// Make our own custom constraints (ie we can't use any or comparable, so create our own set)
type addable interface {
	int | float64 | string
}

// If we added something that would break our function (like bool) it would flag

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v // Problem is Go doesn't know how to use + for literally any type
	}

	return result
}

// MARK: Errors
func mod10() {
	// Errors are easy to create
	err := errors.New("This is an error")
	fmt.Println(err)
	err2 := fmt.Errorf("This error wraps the first one: %w", err)
	fmt.Println(err2)

	// Don't handle error
	// fmt.Println(divNoHandle(10,0))

	// Handle with error
	// result, err := divError(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
	// fmt.Println("Result:", result)

	// Handle with panic
	// result, err := divPanic(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return 
	// }
	// fmt.Println("Result:", result)

}

func divNoHandle(l, r int) int {
	return l / r
}

func divError(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("division by zero")
	}
	return l / r, nil
}

// We want to return to normal program flow as quickly as possible
// This is a good use case for named return values!
func divPanic(l, r int) (result int, err error){
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}

// MARK: Notes
/*
Module 3
	Simple Data Types
		Not primatives, we're talking about data types that can contain only one value (strings, numbers, booleans, errors)

		Strings come in 2 varieties: quoted strings ("this is an interpreted string") and backtick strings (`this is a raw string`)
		If you put a \n for example in the middle of an interpreted string, it would behave like I expect (ie it would make a newline).
		If you put a \n in the middle of a raw string, it will print as a normal character. Raw strings also don't interpret whitespace
		and will print exactly how you write it.

		Numbers come as int, uint (lowest value is 0), float32, float64, complex64, complex128 (there are others, but these are the
		most common)

		Booleans (they work the same everywhere)

		Errors "the error build-in interface type is the conventional interface for representing an error condition, with the nil value
		representing no error". The error type is probably the most complicated, for example there is no discrete list of types because
		it's an interface, but basically it tells you that an error has occured.

	pkg.go.dev

	Pointers and Values
		When I create a variable foo and assign a value, I assign a place in memory and store the value there
		If I create a variable bar and assign foo as it's value, I copy what is currently in foo into the memory of bar
		If I change foo later on and don't reassign bar, bar's value doesn't change
		Now, I can create a variable baz and assign &foo, the value I've assigned to baz is the memory address of foo
		Dereferencing the pointer (*baz) shows me the value being stored in the address assigned to baz
		And now if I change foo and print *baz I get the updated value of foo because we're printing the same memory address
		Pointers are sharing memory, values are used to copy memory

Module 4
	Creating a CLI (Command Line Interface) Application
		CLI Application connects the user and computer through the monitor (display output) and keyboard (recieve input).
		Standard library tools for CLI: the OS package (stdin, stdout, stderr) and fmt package (string management => scan functions, print functions) and
		bufio package (buffered i/o => group text into useful chunks)

	Creating a Web Service
		A web service is 2 computers (the client and the server) that are talking to each other with requests and responses
		There is one comprehensive package for web in Go: net/http

Module 5
	Aggregate Data Types
		AKA collections

		Arrays are fixed sized collections of data that are all the same type

		Slices always refer to data stored in some array. Slices aren't fixed in size. Slices are referenced data types, which means if you update the value
		in the array it changes the value in the slice, and if you change the value in the slice it changes the value in the array.
		The slices package is technically part of the standard library but is really an experimental portion, so we need to import it from the terminal with
		go get golang.org/x/exp/slices. It creates a go.sum file the next time we run it, which keeps tracks of our dependencies.

		Maps let us use our own index, using value and key instead of value and (strictly numeric) index. Maps are also not fixed size. Maps are also reference
		types like slices. Maps can have 2 data types (one for the keys and one for the values)
		Experimental package: golang.org/x/exp/maps

		Structs are special because they can contain different data types at the same time. Structs are fixed sized. The fields do need to be hardcoded, they
		can't be determined at runtime.

Module 6
	Looping (control of execution)
		Every type of loop is a for loop, but there are different types
		Infinite loop

		Loop til condition

		Counter-based loop

		Looping over collections

Module 7
	Panic is a built in function that tell us that the function is no longer stable and destroys the function and returns to its calling function but it
	doesn't know how to handle the panic so it also gets destroyed and returns to its caller. That goes all the way up the call stack until the program exits
	We can also use deferred functions to react to a panic
	recover() is the paired built in function that goes with panic()

	goto statements
	1. Can leave a block (anything contained in curly braces)
	2. Can jump to a containing block
	3. Cannot jump after variable declarations
	4. Cannot jump into another block

Module 8
	func funcName (parameters string, are int, comma, delimited sting) (return values) { function body }
	Above, comma is type string because it inherits the next specified datatype
	Variadic parameters pass a comma delimited list as a splice. It has to be the last parameter, and there can be only one per function.
	Passing pointers you add *dataType in the function signature, *variableName when we use it in the function, and &passedVariable
	Pointers should only be used when we want to share memory. Every other instance should use values

	Values can be returned as a single value (single value or expression) and is put back into the variable assigned to the function call.
	Don't need () in the function signature
	Values can be returned as multiple values (values or expressions) and is put back into the variables assigned to the function call, in order.
	Do need () in the function signature (not the return line)
	Values can be returned as a Named Return value. The function signature looks almost the same except we've included variable names with the types.
	And instead of explicitly listing the variable names on the return line, we just write return. The return statement will then return the current
	value of the variables we listed in the function signature. These aren't super common, but it's just interesting to think about

	A package is a directory within a module that contians at least one source file. Keep in mind that all members are visible to other package members
	In your source file, the first line is the package declaration. Every file in the package will have the same name there, and that's the name of the folder it's in.
	Then, you'll have the import statement.
	After this is your package level member (variables, constants, functions, etc)
	There are 2 levels of visability
		Package level field (has lowercase first letter)
		Public field (has uppercase first letter)

Module 9
	Go isn't by definition object oriented, but it can be used that way

	A method starts with a custom type, something we control. It doesn't have to be a struct like before.
	Once we have our type, we create what looks like a function signature, but there is a method receiver too
	Function: func isEven(i int) bool
	Method:   func (i myInt) isEven() bool
	The invocation is also slightly different
	Function: ans := isEven(i)
	Method:   ans := var.isEven()
	The method receiver tells us if we're passing a value or a pointer. And generally to decide which to use, pointer receivers are used to share variables
	between the caller and method
	In terms of using methods vs functions... use whatever is most readable. Sometimes binding data to a function is beneficial (ie methods) but both are
	valid and common

	Interfaces let us generalize behaviours, for example we want a reader interface but we don't always know what type we're reading from (file vs tcp connection).
	Interfaces are great, but types lose their identitiy and that's not necessarily what we want.

	Generic programming helps solve the lost identity by temporarily changing the identity but after using the generic function the object goes back to its original type.
	✨Transient Polymorphism✨
	We can use generic functionality without sacrificing what we know about the object

Module 10
	Error management
		Errors shouldn't be surprising or unexpected, when running on a big enough scale we should expect to run into errors eventually
		Errors in Go are values, like any other function return value. It's not a result we want, but it's still a result that we can do something with
		Check if an error has occured immediately after it may have happened... handle things first, then continue on
	Errors vs panics
		We want to avoid panicing as often as possible. An error is just the result of an operation, a panic changes the flow of the program. 
		Errors are part of the function signature, for panics we rely on documentation or understanding code to know if it's a possibility
		Errors indicate a deviation from the plan, panics indicate unstability 

*/
