// Each Go source code begins with a package declaration.  The package keyword helps you
// define the name of a new package, which can be anything you want with just one exception:
// if you are creating an executable application and not just a package that will be shared
// by other applications or packages, you should name your package main.
package main

// The import keyword is used for importing other Go packages in your Go programs
// in order to use some or all of their functionality. A Go package can either be a part
// of the rich Standard Go library or come from an external source. Packages of the
// standard Go library are imported by name (os) without the need for a hostname and
// a path, whereas external packages are imported using their full internet paths, like
// github.com/spf13/cobra.
import (
	"fmt"
)

// Go considers a main() function the entry point to the application and begins
// the execution of the application with the code found in the main() function of the
// main package.
// Each Go function definition begins with the func keyword followed by its name,
// signature and implementation. There is a global Go rule that also applies to
// function and variable names and is valid for all packages except main: everything
// that begins with a lowercase letter is considered private and is accessible in the
// current package only.
func main() {
	fmt.Println("Hello World!")
}
