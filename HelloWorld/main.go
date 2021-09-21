// So here this package main tells that this is an executable package and this will lead to a file that will be executed.
// If the name of package was anything other then main then it would have been reusable package , that we can import and use.
package main

// Here we are importing other files
import "fmt"

// This main function tells that this is the entry point of execution for this main package.
func main(){
	fmt.Println("Hello world!!");
}