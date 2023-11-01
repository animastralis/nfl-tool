package main

import (
	t "github.com/animastralis/nfl-tool/team"
)

func main() {
	teamManager := *t.InitTeamManager()
	teamManager.PrintTeams()
}
