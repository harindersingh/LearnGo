/*
if(boolean_expression) {
   // statement(s) will execute if the boolean expression is true 
}
*/

package main

import "fmt"

func main() {
	//local variable definition
	var a int = 100

	/* check the boolean condition using if statement */
   if( a < 200 ) {
      /* if condition is true then print the following */
      fmt.Printf("a is less than 200\n" )
   }
   fmt.Printf("value of a is : %d\n", a)	
}
