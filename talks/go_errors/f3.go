package go_errors

import (
	"errors"
	"fmt"
)

func main() {
	const fn = "temp.cpp"
	var ln = 19
	text := fmt.Sprintf("compile problem with %q: %d", fn, ln)
	err := errors.New(text)
	if err != nil {
		fmt.Println(err)
	}
}
