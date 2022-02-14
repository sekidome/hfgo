package main

import "fmt"

func calcColor(x float64, y float64) {
	result := (x * y) / 10.0
	//  %f is the format of an float. %.2f (or %0.2f) says that there is no padding before the "." and that the float is round to 2 decimal places.
	fmt.Printf("%0.2f liters needed.\n", result)
}

func main() {
	calcColor(19, 14.5)
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
