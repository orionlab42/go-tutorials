package _2_program_structure

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//2.1 Names
// - a name begins with a letter or an underscore and may have any number of additional letters, digits, and underscores
// - case matters: heapSort and Heapsort are different names
// - if an entity is declared within a function, it is local to that function
// - if declared outside a function, however, it is visible in all files of the package to which it belongs
// - the case of the first letter of a name determines its visibility across package boundaries. If the name begins with
// an upper-case letter, it is exported, which means that it is visible and accessible outside its own package and may be
// referred to by other parts of the program
// - when forming names by combining words ‘camel case’ is used
// - ASCII and HTML are always rendered in the same case, so a function might be called htmlEscape, HTMLEscape, or
// escapeHTML, but not escapeHtml

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//2.2 Declarations
// - declaration names a program entity and specifies some or all of its properties
// - there are four major kinds of declarations: var, const, type, and func
// - Each file begins with a package declaration that says what package the file is part of, followed by any import
// declarations, and then a sequence of package-level declarations of types, variables, constants, and functions, in any order.
// - the name of each package-level entity is visible not only throughout the source file that contains its declaration,
// but throughout all the files of the package. ex. const boilingF = 212.0 if it's declared right after the import
// declaration, outside any function

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//2.3 Variables
// -var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value
// var name type = expression
var i, j, k int                 // int, int, int
var b, t, s = true, 2.3, "four" // bool, float64, string
var c, err = os.Open(s)         // os.Open returns a file and an error

//////////////////////////////////////////////////////////////////
//2.3.1 Short Variable Declaration
// -within a function, an alternate form called a short variable declaration may be used to declare and initialize local variables
// name:= expression

func shortDeclaration(name string) error {
	hello := "Hello"
	i, j := 0, 1
	i, j = j, i // swap values of i and j
	fmt.Println(hello, i, j)
	_, err := os.Open(hello)
	f, err := os.Open(name) // here err is only assigned, not declared, because it was already declared before
	if err != nil {
		return err
	}
	fmt.Println(f)
	return nil
}

// -keep in mind that := is a declaration, whereas = is an assignment
// -one subtle but important point: a short variable declaration does not necessarily declare all the
// variables on its left-hand side. If some of them were already declared in the same lexical block
// ,then the short variable declaration acts like an assignment to those variables.

//////////////////////////////////////////////////////////////////
//2.3.2 Pointers
// -a variable is a piece of storage containing a value
// -a pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every
// value has an address, but every variable does. With a pointer, we can read or update the value of a variable indirectly,
// without using or even knowing the name of the variable, if indeed it has a name.

func pointers() {
	x := 1          // the expression &x("address of x") yields a pointer to an integer variable
	p := &x         // p is a value of type *int(pronounced "pointer to int"); "p points to x" or equivalently "p contains the address of x"
	fmt.Println(*p) // "1", the value stored at this pointer,
	*p = 2          // equivalent to x = 2,  the expression *p yields the value of that variable
	fmt.Println(x)  // "2"
}

// -Variables are sometimes described as addressable values. Expressions that denote variables are the only expressions
// to which the address-of operator & may be applied.
// -The zero value for a pointer of any type is nil. The test p != nil is true if p points to a variable.
// Pointers are comparable; two pointers are equal if and only if they point to the same variable or both are nil.

var p = f()

func f() *int {
	v := 1
	return &v
}

// -Each call of f returns a distinct value: fmt.Println(f() == f()) // "false"
func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}
func incrUse() {
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

// -Each time we take the address of a variable or copy a pointer, we create new aliases or ways to
// identify the same variable. For example, *p is an alias for v. Pointer aliasing is useful because
// it allows us to access a variable without using its name, but this is a double-edged sword: to
// find all the statements that access a variable, we have to know all its aliases.
// -It’s not just pointers that create aliases; aliasing also occurs when we copy values of other reference types like
// slices, maps, and channels, and even structs, arrays, and interfaces that contain these types.

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// Echo4 prints its commandline arguments.
func Echo4() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("bla")
	}
	fmt.Printf("%+v", n)
}

//////////////////////////////////////////////////////////////////
//2.3.3 The new Function
// - The expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns its
// address, which is a value of type *T.

func NewFunction() {
	p := new(int)   // creates p, of type *int, points to an unnamed int variable
	fmt.Println(p)  // "0xc00000e114" its memory address
	fmt.Println(*p) // "0"
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p) // "2"
}

// -a variable created with new it's just like an ordinary local variable, "new" is only a syntactic convenience, not a fundamental notion
// -Each cal l to new returns a distinct variable with a unique address

type bum struct {
	Name string
}

func exceptionUniqueAddress() {
	p := bum{}
	q := bum{}
	fmt.Println(p == q) // this returns true
}

// -exception: two variables whose typ e carries no information and is therefore of size zero, such as struct{} or [0]int,
// may, dep ending on the implementation, have the same address.

//////////////////////////////////////////////////////////////////
//2.3.4 Lifetime of Variables
var global *int

func lifetimeVariablesF() {
	var x int
	x = 1
	global = &x
}

// -must be heap-al located because it is still reachable from the variable global after lifetimeVariablesF() has returned,
// despite being declared as a local variable; we say x escapes from lifetimeVariablesF function

func lifetimeVariablesG() {
	y := new(int)
	*y = 1
}

// -by contrast when lifetimeVariablesG() returns, the variable *y becomes unreachable and can be recycled
// -Since *y does not escape from lifetimeVariablesG(), it’s safe for the compiler to allocate *y on the stack, even though it
// was allocated with new.
// -it’s good to keep in mind dur ing performance optimization, since each variable that escapes requires an extra memory allocation

// -keeping unnecessary pointers to short-lived objects within long-lived objects, especially global variables, will prevent
// the garbage collector from reclaiming the short-lived objects.

//////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////
//2.4 Assignments
// -ex:
// count[x] *= scale
// v++		// same as v = v + 1; v becomes 2
// v--		// same as v = v 1; v becomes 1 again

//////////////////////////////////////////////////////////////////
//2.4.1. Tuple Assignment
//	x, y = y, x
//	a[i], a[j] = a[j], a[i]

// -when computing the greatest common divisor (GCD) of two integers:

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

// -Examples:
// f, err = os.Open("foo.txt")	// function call returns two values
// v, ok = m[key]				// map lookup, produces an additional boolean
// v, ok = x.(T)				// type assertion, produces an additional boolean
// v, ok = <ch					// channel receive, produces an additional boolean
// _, err = io.Copy(dst, src)	// discard byte count
// _, ok = x.(T)				// check type but discard result
