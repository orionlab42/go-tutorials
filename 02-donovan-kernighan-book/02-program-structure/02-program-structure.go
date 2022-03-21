package _2_program_structure

import "fmt"

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

func Simple() {
	fmt.Println("simple")
}
