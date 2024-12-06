package hello

import "strings"

//func Say(name string) string {
//	return fmt.Sprintf("Hello, %s!", name)
//}

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"World"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
