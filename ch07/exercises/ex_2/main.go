package main

import (
	"fmt"
	"sort"
)

type Team struct {
	TeamName    string
	PlayerNames []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		l.Wins[team1]++
	}
	if score2 > score1 {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	type TeamWin struct {
		TeamName string
		Wins     int
	}
	teams := make([]TeamWin, 0, len(l.Teams))
	for teamName := range l.Teams {
		teams = append(teams, TeamWin{teamName, l.Wins[teamName]})
	}
	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Wins > teams[j].Wins
	})
	rankings := make([]string, 0, len(l.Teams))
	for _, team := range teams {
		rankings = append(rankings, team.TeamName)
	}
	return rankings
}

func main() {
	teamNames := []string{
		"England",
		"France",
		"Ireland",
		"Italy",
		"Scotland",
		"Wales",
	}
	league := League{
		Teams: make(map[string]Team, len(teamNames)),
		Wins:  map[string]int{},
	}
	for _, teamName := range teamNames {
		league.Teams[teamName] = Team{TeamName: teamName}
	}

	league.MatchResult("Italy", 10, "France", 50)
	league.MatchResult("England", 6, "Scotland", 11)
	league.MatchResult("Wales", 21, "Ireland", 16)

	league.MatchResult("England", 41, "Italy", 18)
	league.MatchResult("Scotland", 24, "Wales", 25)
	league.MatchResult("Ireland", 13, "France", 15)

	league.MatchResult("Italy", 10, "Ireland", 48)
	league.MatchResult("Wales", 40, "England", 24)
	league.MatchResult("France", 23, "Scotland", 27)

	league.MatchResult("Italy", 7, "Wales", 48)
	league.MatchResult("England", 23, "France", 20)
	league.MatchResult("Scotland", 24, "Ireland", 27)

	league.MatchResult("Scotland", 52, "Ireland", 10)
	league.MatchResult("Ireland", 32, "England", 18)
	league.MatchResult("France", 32, "Wales", 30)

	fmt.Println(league.Ranking())
}
