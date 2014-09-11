//creat a new type
package main
import "fmt"
type NameAge struct {
	name string
	age int
}
func main() {
	a := new(NameAge)
	a.name = "pete" ; a.age = 42

	fmt.Printf("%v \n", a)
}

