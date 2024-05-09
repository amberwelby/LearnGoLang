/*
	The file doesn't need to be called main, but .go is the extension
	There are a few options for how to build, complile and run, but we want to run 'go run . ' in our command line
*/

// Start with the package identifier
package main

// If you start working with a library it seems to add it automatically
import "fmt"

// Entry point is determined by a function called main
func main() {
	println("Hello, Gophers!")     // This is a built in print function, but isn't what we usually are going to use (good for keeping less dependencies, debugging)
	fmt.Println("Hello, Gophers!") // This is the imported version we typically want to use, for formatted strings and such

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

/*
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
*/