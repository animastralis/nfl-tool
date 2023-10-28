package main

import (
	"fmt"

	t "github.com/animastralis/nfl-tool/team"
)

func main() {
	teams := t.GetTeams()
	for _, team := range teams {
		fmt.Println(team)
	}
}
