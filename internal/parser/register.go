package parser

import (
	"strings"

	"github.com/jadson-medeiros/quake_log_parser/internal/pkg/data"
)

// RegisterKill registers a kill in the match data.
func RegisterKill(match *data.MatchData, tokens []string) {
	match.TotalKills++

	killer, victim, killMean := parseKill(tokens)

	if killer != "<world>" {
		match.KillCount[killer]++
	} else {
		match.KillCount[victim]--
	}

	if killer == victim {
		match.KillMeans["MOD_SUICIDE"]++
	} else if _, ok := match.KillMeans[killMean]; !ok {
		match.KillMeans["MOD_UNKNOWN"]++
	} else {
		match.KillMeans[killMean]++
	}
}

// parseKill parses the kill information from tokens and returns killer, victim, and kill mean.
func parseKill(tokens []string) (string, string, string) {
	killer, victim, killMean := strings.Builder{}, strings.Builder{}, strings.Builder{}
	i := 5
	for tokens[i+1] != "killed" {
		killer.WriteString(tokens[i] + " ")
		i++
	}
	killer.WriteString(tokens[i])
	i += 2
	for tokens[i+1] != "by" {
		victim.WriteString(tokens[i] + " ")
		i++
	}
	victim.WriteString(tokens[i])
	i += 2
	for i < len(tokens) {
		killMean.WriteString(tokens[i])
		i++
	}
	return strings.TrimSpace(killer.String()), strings.TrimSpace(victim.String()), killMean.String()
}

// RegisterPlayer registers a new player in the match data.
func RegisterPlayer(match *data.MatchData, tokens []string) {
	player := extractPlayerName(tokens)
	if len(player) > 1 && !sliceContainsString(match.Players, player) {
		match.Players = append(match.Players, player)
		match.KillCount[player] = 0
	}
}

// extractPlayerName extracts the player name from tokens.
func extractPlayerName(tokens []string) string {
	start := false
	builder := strings.Builder{}
	for _, token := range tokens {
		if strings.Contains(token, "\\n") {
			start = true
		} else if start {
			builder.WriteString(token + " ")
		}
	}
	return strings.TrimSpace(builder.String())
}

// sliceContainsString checks if a string exists in a slice.
func sliceContainsString(array []string, find string) bool {
	for _, aux := range array {
		if aux == find {
			return true
		}
	}
	return false
}
