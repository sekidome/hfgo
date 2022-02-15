package main

import (
	"fmt"
)

// before {} you must specifie which type the return value should have. (here float64).
// If all parameters have the same type, it is enough to specifie the last. (only first will result in error)
func calcColor(x, y float64) (neededColor float64, ltZero error) { // neededColor is an optional name for the return value
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("Sum should not be 0 or less.") // Even if there was an error, u need to return a primary value (here 0), but it can be ignored.
	}
	return (x * y) / 10.0, nil
}

func main() {
	result, err := calcColor(-19, 14.5)
	if err != nil {
		fmt.Println(err)
	}
	if result != 0 {
		//  %f is the format of an float. %.2f (or %0.2f) says that there is no padding before the "." and that the float is round to 2 decimal places.
		fmt.Printf("%0.2f liters needed.\n", result)
	}
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
