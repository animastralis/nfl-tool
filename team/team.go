package team

import (
	"encoding/json"
	"fmt"

	"github.com/animastralis/nfl-tool/util"
)

const NUM_TEAMS = 32
const TEAM_URL_BASE = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams"

type Team struct {
	Id           string `json:"id"`
	FullName     string `json:"displayName"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Record       string
	Schedule     TeamSchedule
}

type TeamRecord struct {
	Name    string  // "overall"
	Type    string  // "total"
	Summary string  // "1-6"
	Value   float64 // 0.123456789
}

type TeamGame struct {
	Week       uint8
	OpponentId string
	HasWinner  bool
	Winner     bool
}

type TeamSchedule struct {
	Games []TeamGame
}

func initTeams() *map[string]Team {
	teams := make(map[string]Team)
	result := util.GetApiData(TEAM_URL_BASE)
	sports := result["sports"].([]interface{})
	leagues := sports[0].(map[string]interface{})["leagues"].([]interface{})
	teamsData := leagues[0].(map[string]interface{})["teams"].([]interface{})

	for _, teamData := range teamsData {
		teamString, _ := json.Marshal(teamData.(map[string]interface{})["team"].(map[string]interface{}))
		var t Team
		if err := json.Unmarshal(teamString, &t); err != nil {
			fmt.Println(err)
		}

		teams[t.Id] = t
	}

	for id := range teams {
		t := teams[id]
		t.Record, t.Schedule = t.generateSchedule()
		teams[id] = t
	}

	return &teams
}

func (t Team) generateSchedule() (string, TeamSchedule) {
	url := fmt.Sprintf("%s/%v/schedule", TEAM_URL_BASE, t.Id)
	result := util.GetApiData(url)

	// Set Team Record
	team := result["team"].(map[string]interface{})
	record := team["recordSummary"].(string)

	// Parse schedule
	var ts TeamSchedule
	events := result["events"].([]interface{})
	for _, event := range events {
		var tg TeamGame
		tg.Week = uint8(event.(map[string]interface{})["week"].(map[string]any)["number"].(float64))

		competition := event.(map[string]interface{})["competitions"].([]interface{})[0].(map[string]interface{})
		competitors := competition["competitors"].([]interface{})
		for _, c := range competitors {
			competitorId := c.(map[string]interface{})["id"].(string)
			if competitorId == t.Id {
				if c.(map[string]interface{})["winner"] != nil {
					tg.HasWinner = true
					tg.Winner = c.(map[string]interface{})["winner"].(bool)
				} else {
					tg.HasWinner = false
				}
			} else {
				tg.OpponentId = competitorId
			}
		}
		ts.Games = append(ts.Games, tg)
	}
	return record, ts
}
