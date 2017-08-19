package main 
import "fmt"

func main() {
	var a int = 100;

	if(a == 100) {
		fmt.Printf("Main dekha teri photo %d %d vaar kudde\n", a, a);
	} else {
		fmt.Printf("a is not 100\n");
	}
	fmt.Printf("Value of a is : %d\n", a);
}