package team

import (
	"fmt"
)

type TeamManager struct {
	Teams map[string]Team
}

func InitTeamManager() *TeamManager {
	teamManager := TeamManager{}
	teamManager.InitTeams()
	return &teamManager
}

func (tm *TeamManager) InitTeams() {
	teams := initTeams()
	tm.Teams = *teams
}

func (tm *TeamManager) PrintTeams() {
	for _, team := range tm.Teams {
		fmt.Println(team)
	}
}
