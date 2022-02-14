package main

import "fmt"

// before {} you must specifie which type the return value should have. (here float64).
// If all parameters have the same type, it is enough to specifie the last. (only first will result in error)
func calcColor(x, y float64) float64 {
	return (x * y) / 10.0
}

func main() {
	result := calcColor(19, 14.5)
	//  %f is the format of an float. %.2f (or %0.2f) says that there is no padding before the "." and that the float is round to 2 decimal places.
	fmt.Printf("%0.2f liters needed.\n", result)
}

/*Verbs:

%f floating point
%d decimal integer
%s string
%t Boolean
%v any value
%#v any value, formatte as it would appear in Go
%T Type of the supplied value
%% a literal percent sign

*/
