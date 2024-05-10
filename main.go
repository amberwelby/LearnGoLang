/*
	The file doesn't need to be called main, but .go is the extension
	There are a few options for how to build, complile and run, but we want to run 'go run . ' in our command line
*/

// Start with the package identifier
package main

// If you start working with a library (and save the file) it seems to add it automatically
import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"strings"
)

// MARK: Main
// Entry point is determined by a function called main
func main() {
	println("Hello, Gophers!")     // This is a built in print function, but isn't what we usually are going to use (good for keeping less dependencies, debugging)
	fmt.Println("Hello, Gophers!") // This is the imported version we typically want to use, for formatted strings and such
}

// MARK: Module 3
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

// MARK: Module 4
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

// MARK: Module 5
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

// MARK: Module 6
func mod6(){
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
		j ++
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

// MARK: Demo App
func menu() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	// in := bufio.NewReader(os.Stdin)
	// choice, _ := in.ReadString('\n')
	// choice = strings.TrimSpace(choice)

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices{
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
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
*/

/*
	Get queries on sprint 6
*/
