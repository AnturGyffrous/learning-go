package main

type Team struct {
	TeamName    string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}
