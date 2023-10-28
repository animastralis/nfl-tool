package team

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/animastralis/nfl-tool/util"
)

type Team struct {
	Id           string `json:"id"`
	FullName     string `json:"displayName"`
	Name         string `json:"nickname"`
	Abbreviation string `json:"abbreviation"`
	TeamRecord   TeamRecord
}

type TeamRecord struct {
	Name    string  // "overall"
	Type    string  // "total"
	Summary string  // "1-6"
	Value   float64 // 0.123456789
}

func httpGet(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		util.PrintError(err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		util.PrintError(err)
		return nil
	}

	return result
}

func GetTeams() []Team {
	const teamsBaseUrl = "https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2023/teams"
	const teamRequestLimit = 50

	// Get Team Links
	result := httpGet(fmt.Sprintf("%s?limit=%d", teamsBaseUrl, teamRequestLimit))

	var teamLinks []string
	for _, linkMap := range result["items"].([]interface{}) {
		teamLinks = append(teamLinks, linkMap.(map[string]interface{})["$ref"].(string))
	}

	// Get Teams
	var teams []Team
	for _, link := range teamLinks {
		result = httpGet(link)

		var team Team
		team.Id = result["id"].(string)
		team.FullName = result["displayName"].(string)
		team.Name = result["nickname"].(string)
		team.Abbreviation = result["abbreviation"].(string)
		teams = append(teams, team)
	}

	return teams
}

func GetTeamRecord(id string) *TeamRecord {
	url := fmt.Sprintf("https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2023/types/2/teams/%s/record", id)
	result := httpGet(url)

	items := result["items"].([]interface{})
	var tr TeamRecord

	for k, v := range items[0].(map[string]interface{}) {
		switch k {
		case "name":
			tr.Name = v.(string)
		case "type":
			tr.Type = v.(string)
		case "summary":
			tr.Summary = v.(string)
		case "value":
			tr.Value = v.(float64)
		}
	}

	return &tr
}
