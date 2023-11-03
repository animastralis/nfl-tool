package main

import (
	"encoding/csv"
	"log"
	"os"

	t "github.com/animastralis/nfl-tool/team"
)

func main() {
	teamManager := *t.InitTeamManager()

	// CSV
	file, err := os.Create("SOS.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()
	for _, team := range teamManager.Teams {
		row1 := []string{team.Name}
		row2 := []string{""}
		for _, game := range team.Schedule.Games {
			opponent := teamManager.Teams[game.OpponentId]
			opponentName := opponent.Name
			row1 = append(row1, opponentName)
			if game.HasWinner {
				if game.Winner {
					row2 = append(row2, "WIN")
				} else {
					row2 = append(row2, "LOSE")
				}
			} else {
				opponentRecord := opponent.Record
				row2 = append(row2, opponentRecord)
			}
		}
		if err := w.Write(row1); err != nil {
			log.Fatalln("error writing line to file", err)
		}

		if err = w.Write(row2); err != nil {
			log.Fatalln("error writing line to file", err)
		}
	}
}
