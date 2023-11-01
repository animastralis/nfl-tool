package team

import (
	"fmt"
)

type TeamManager struct {
	Teams []Team
}

func InitTeamManager() *TeamManager {
	teamManager := TeamManager{}
	teamManager.InitTeams()
	return &teamManager
}

func (tm *TeamManager) InitTeams() {
	teams := initTeams()
	for k, _ := range *teams {
		(*teams)[k].fetchTeamRecord()
	}

	tm.Teams = *teams
}

func (tm *TeamManager) PrintTeams() {
	for _, team := range tm.Teams {
		fmt.Println(team.Record)
	}
}
