package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func getTeams() []Team {
	const TeamsBaseURL string = "https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2023/teams"

	// Get Team Links
	var result map[string]interface{}

	resp, err := http.Get(TeamsBaseURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("ERROR: %s", err)
		return nil
	}

	var teamLinks []string
	fmt.Println(result["items"])
	for _, linkMap := range result["items"].([]interface{}) {
		fmt.Println(linkMap)
		teamLinks = append(teamLinks, linkMap.(map[string]interface{})["$ref"].(string))
	}

	// Get Teams

	var teams []Team
	for _, link := range teamLinks {
		url := fmt.Sprintf("%s/%s", TeamsBaseURL, link)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			fmt.Println("ERROR: %s", err)
			return nil
		}

		items := result["items"].([]interface{})
		var team Team
		for k, v := range items[0].(map[string]interface{}) {
			switch k {
			case "id":
				team.Id = v.(string)
			case "displayName":
				team.FullName = v.(string)
			case "nickname":
				team.Name = v.(string)
			case "abbreviation":
				team.Abbreviation = v.(string)
			}
			teams = append(teams, team)
		}
	}

	return teams
}

func getTeamRecord(id string) *TeamRecord {
	var result map[string]interface{}
	url := fmt.Sprintf("https://sports.core.api.espn.com/v2/sports/football/leagues/nfl/seasons/2023/types/2/teams/%s/record", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("ERROR: %s", err)
		return nil
	}

	var items []interface{}

	items = result["items"].([]interface{})
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

func main() {
	teams := getTeams()
	fmt.Println(teams)
	tr := getTeamRecord("9")
	fmt.Println(tr.Value)
}
