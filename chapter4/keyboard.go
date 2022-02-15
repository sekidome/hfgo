package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin) // reading loop
	//fmt.Println(reader) --> This will keep on printing 0s
	input, err := reader.ReadString('\n') // ReadString method requires an argument with a rune that marks the end of the input
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
