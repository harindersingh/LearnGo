/*v
var variable_list optional_data_type;
var  i, j, k int;
var  c, ch byte;
var  f, salary float32;

variable_name = value;
d = 3, f = 5;    // declaration of d and f. Here d and f are int 
*/

package main 

import "fmt"

func main() {
	var x float32
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T \n", x)	
}