#Ch2 - Program Structure

***

# 2.1 Names 
* A name begins with a letter or an underscore and may have any number of additional letters, digits, and underscores.
* Case matters: heapSort and Heapsort are different names.
* If an entity is declared within a function, it is local to that function, if declared outside a function, however, it 
is visible in all files of the package to which it belongs.
* The case of the first letter of a name determines its visibility across package boundaries. If the name begins with 
an upper-case letter, it is exported, which means that it is visible and accessible outside its own package and may be
referred to by other parts of the program.
* When forming names by combining words ‘camel case’ is used.
* ASCII and HTML are always rendered in the same case, so a function might be called htmlEscape, HTMLEscape, or
escapeHTML, but not escapeHtml.

# 2.2 Declarations
* Declaration names a program entity and specifies some or all of its properties.
* There are four major kinds of declarations: var, const, type, and func.
* Each file begins with a package declaration that says what package the file is part of, followed by any import
declarations, and then a sequence of package-level declarations of types, variables, constants, and functions, in any order.
* The name of each package-level entity is visible not only throughout the source file that contains its declaration,
but throughout all the files of the package, if it's declared right after the import declaration, outside any function

```go
const boilingF = 212.0
```

#2.3 Variables
* Var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value.

**var name type = expression**

```go
var i, j, k int                   // int, int, int
var b, t, s = true, 2.3, "four"   // bool, float64, string
var c, err = os.Open(s)           // os.Open returns a file and an error
```

#2.3.1 Short Variable Declaration
* Within a function, an alternate form called a short variable declaration may be used to declare and initialize local 
variables

**name := expression**
```go
func shortDeclaration(name string) error {
  hello := "Hello"
  i, j := 0, 1
  i, j = j, i              // swap values of i and j
  fmt.Println(hello, i, j)
  _, err := os.Open(hello)
  f, err := os.Open(name) // here err is only assigned, not declared, because it was already declared before
  if err != nil {
    return err
  }
  return nil
}
```

* Keep in mind that := is a declaration, whereas = is an assignment.
* One subtle but important point: a short variable declaration does not necessarily declare all the variables on its 
left-hand side. If some of them were already declared in the same lexical block, then the short variable declaration acts 
like an assignment to those variables.

#2.3.2 Pointers
* A variable is a piece of storage containing a value.
* A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every
value has an address, but every variable does. With a pointer, we can read or update the value of a variable indirectly,
without using or even knowing the name of the variable, if indeed it has a name.

```go
func pointers() {
  x := 1          // the expression &x("address of x") yields a pointer to an integer variable
  p := &x         // p is a value of type *int(pronounced "pointer to int"); "p points to x" or equivalently "p contains the address of x"
  fmt.Println(*p) // "1", the value stored at this pointer,
  *p = 2          // equivalent to x = 2,  the expression *p yields the value of that variable
  fmt.Println(x)  // "2"
}
```

* Variables are sometimes described as addressable values. Expressions that denote variables are the only expressions
to which the address-of operator & may be applied.
* The zero value for a pointer of any type is nil. The test p != nil is true if p points to a variable. Pointers are 
comparable; two pointers are equal if and only if they point to the same variable or both are nil.

```go
var p = f()

func f() *int {
    v := 1
return &v
}
```
* Each call of f returns a distinct value: 

```go
fmt.Println(f() == f())     // "false"

func incr(p *int) int {
    *p++                    // increments what p points to; does not change p
    return *p
}

func incrUse() {
    v := 1
    incr(&v)              // side effect: v is now 2
    fmt.Println(incr(&v)) // "3" (and v is 3)
}
```

* Each time we take the address of a variable or copy a pointer, we create new aliases or ways to identify the same 
variable. For example, *p is an alias for v. Pointer aliasing is useful because it allows us to access a variable 
without using its name, but this is a double-edged sword: to find all the statements that access a variable, we have to 
know all its aliases.
* It’s not just pointers that create aliases; aliasing also occurs when we copy values of other reference types like
slices, maps, and channels, and even structs, arrays, and interfaces that contain these types.

```go
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
```

#2.3.3 The _new_ Function
* The expression new(T) creates an unnamed variable of type T, initializes it to the zero value of T, and returns its
address, which is a value of type *T.

```go
func NewFunction() {
    p := new(int)       // creates p, of type *int, points to an unnamed int variable
    fmt.Println(p)      // "0xc00000e114" its memory address
    fmt.Println(*p)     // "0"
    *p = 2              // sets the unnamed int to 2
    fmt.Println(*p)     // "2"
}
```
* A variable created with new it's just like an ordinary local variable, "new" is only a syntactic convenience, not a 
fundamental notion. Each call to _new_ returns a distinct variable with a unique address.

```go
type bum struct {
    Name string
}

func exceptionUniqueAddress() {
    p := bum{}
    q := bum{}
    fmt.Println(p == q) // this returns true
}
```

* Exception: two variables whose type carries no information and is therefore of size zero, such as struct{} or [0]int,
may, depending on the implementation, have the same address.

#2.3.4 Lifetime of Variables
```go
var global *int

func lifetimeVariablesF() {
    var x int
    x = 1
    global = &x
}
```

* Must be heap-al located because it is still reachable from the variable global after lifetimeVariablesF() has returned,
despite being declared as a local variable; we say x escapes from lifetimeVariablesF function

```go
func lifetimeVariablesG() {
    y := new(int)
    *y = 1
}
```

* By contrast when lifetimeVariablesG() returns, the variable *y becomes unreachable and can be recycled.
* Since *y does not escape from lifetimeVariablesG(), it’s safe for the compiler to allocate *y on the stack, even though it
was allocated with new.
* It’s good to keep in mind during performance optimization, since each variable that escapes requires an extra memory allocation
* Keeping unnecessary pointers to short-lived objects within long-lived objects, especially global variables, will prevent
the garbage collector from reclaiming the short-lived objects.

#2.4 Assignments
* Ex:
```go
count[x] *= scale
v++		// same as v = v + 1; v becomes 2
v--		// same as v = v 1; v becomes 1 again
```

#2.4.1. Tuple Assignment
```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```

* When computing the greatest common divisor (GCD) of two integers:

```go
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
```

* Examples:
```go
f, err = os.Open("foo.txt")	// function call returns two values
v, ok = m[key]				// map lookup, produces an additional boolean
v, ok = x.(T)				// type assertion, produces an additional boolean
v, ok = <ch					// channel receive, produces an additional boolean
_, err = io.Copy(dst, src)	// discard byte count
_, ok = x.(T)				// check type but discard result
```

#2.4.2. Assignability
* Assignment statements, like above are an explicit form of assignment, but there are many places in a
program where an assignment occurs implicitly: ex. a function call implicitly assigns the argument
values to the corresponding parameter variables;
* Ex. of implicit assignment:  `medals := []string{"gold", "silver", "bronze"}`
* The assignment is legal only if the value is assignable to the type of the variable
* Nil may be assigned to any variable of interface or reference type

#2.5 Type Declaration
* A type declaration defines a new named type that has the same underlying type as an existing type. The named type
provides a way to separate different and perhaps incompatible uses of the underlying type so that they can’t be mixed unintentionally.
* Syntax: **type name underlying-type**
* Type declarations most often appear at package level, where the named type is visible throughout the package, and if
the name is exported (it starts with an upper-case letter), it’s accessible from other packages as well.

```go
type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    Boiling       Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
* This package defines two types, Celsius and Fahrenheit, for the two units of temperature. Even though  both have the
same underlying type, float64, they are not the same type, so they cannot be compared or combined in arithmetic expressions.

```go
var x = 3.5
var y = Celsius(x)
```

* Celsius(t) and Fahrenheit(t) are explicit type conversions, not function calls.
* For every type T, there is a corresponding conversion operation T(x) that converts the value x
to type T. A conversion from one type to another is allowed if both have the same underlying
type, or if both are unnamed pointer types that point to variables of the same underlying type;
these conversions change the type but not the representation of the value.

```go
var celsius Celsius
var fahrenheit Fahrenheit
fmt.Println(celsius == 0) // "true"
fmt.Println(fahrenheit >= 0) // "true"
fmt.Println(celsius == fahrenheit) // compile error: type mismatch
fmt.Println(celsius == Celsius(fahrenheit)) // "true"!
```
* Named types also make it possible to define new behaviors for values of the type. These
behaviors are expressed as a set of functions associated with the type, called the type’s methods.

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

#2.6 Packages and Files
* Packages in Go serve the same purposes as libraries or modules in other languages, supporting
modularity, encapsulation, separate compilation, and reuse.

* Packages also let us hide information by controlling which names are visible outside the package,
or exported. In Go, a simple rule governs which identifiers are exported and which are
not: exported identifiers start with an upper-case letter.

```go
Package tempconv performs Celsius and Fahrenheit conversions. (**)
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
....
```
* The doc comment(**) immediately preceding the package declaration documents the
package as a whole. Conventionally, it should start with a summary sentence in the style
illustrated. Only one file in each package should have a package doc comment. Extensive doc
comments are often placed in a file of their own, conventionally called doc.go.

#2.6.1 Imports
* Within a Go program, every package is identified by a unique string called its import path.
These are the strings that appear in an import declaration.
* In addition to its import path, each package has a package name, which is the short (and not
necessarily unique) name that appears in its package declaration. By convention, a package’s
name matches the last segment of its import path.
* The import declaration binds a short name to the imported package that may be used to refer
to its contents throughout the file.
* By default, the short name is the package name but an import declaration may
specify an alternative name to avoid a conflict

#2.6.2 Package Initialization

* Package initialization begins by initializing package-level variables in the order in which they
are declared, except that dependencies are resolved first:
```go
var a = b + c // a initialized third, to 3
var b = f() // b initialized second, to 2, by calling f
var c = 1 // c initialized first, to 1
func f() int { return c + 1 }

func init() { /* ... */ }
```
* init functions can’t be called or referenced, but otherwise they are normal functions.
* Within each file, init functions are automatically executed when the program starts, in the
order in which they are declared, like this we can initialize some variables, like tables of data

* One package is initialized at a time, in the order of imports in the program, dependencies first,
initialization proceeds from the bottom up; the main package is the last to be initialized. In
this manner, all packages are fully initialized before the application’s main function begins.

#2.7 Scope
* A declaration associates a name with a program entity, such as a function or a variable. The
scope of a declaration is the part of the source code where a use of the declared name refers to
that declaration.

* There is a lexical block for the entire source code, called the universe block; for each package; for each file; for each for, if,
and switch statement; for each case in a switch or select statement; and, of course, for each explicit syntactic block.
* A declaration’s lexical block determines its scope, which may be large or small.

* The declarations of built-in types, functions, and constants like int, len, and true are in the universe
block and can be referred to throughout the entire program.

* If a name is declared in both an outer block and an inner block, the inner declaration will be found first. In that case, the
inner declaration is said to shadow or hide the outer one, making it inaccessible.

```go
func f() {}

func randomFunction() {
	f := "f"
	fmt.Println(f) // "f"; local var f shadows package level
}

var cwd string

func init() {
    cwd, err := os.Getwd() // compile error: unused: cwd
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}

func init() {
    var err error
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

* By declaring err in a separate var declaration we avoid making the outer cwd inaccessible, if no the statement does
not update the package-level cwd variable as intended(in the first init)
