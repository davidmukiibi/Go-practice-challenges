package twofer

import "fmt"

// ShareWith has no comment.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	statement := fmt.Sprintf("One for %s, one for me.", name)
	return statement
}

func main() {
	ShareWith("Jesus")
}
