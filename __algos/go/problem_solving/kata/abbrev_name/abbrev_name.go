package abbrev_name

import (
	"strings"
)

func abbrevName(name string) string {
	if len(name) == 3 {
		return strings.ToUpper(string(name[0])) + "." + strings.ToUpper(string(name[2]))
	}
	words := strings.Split(name, " ")
	return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}
