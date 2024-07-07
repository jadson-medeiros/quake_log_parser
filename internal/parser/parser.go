package parser

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jadson-medeiros/quake_log_parser/internal/pkg/data"
)

// NewMatch initializes a new match in the provided map and returns the match data.
func NewMatch(matches map[string]*data.MatchData, matchNumber int) *data.MatchData {
	newMatch := &data.MatchData{
		TotalKills:  0,
		Players:     []string{},
		KillCount:   map[string]int{},
		Leaderboard: map[int]string{},
		KillMeans:   map[string]int{},
	}

	fillKillMeans(newMatch.KillMeans)

	matchName := fmt.Sprintf("game_%02d", matchNumber)
	matches[matchName] = newMatch

	return newMatch
}

// fillKillMeans initializes the kill means map with default values.
func fillKillMeans(means map[string]int) {
	meansOfDeath := []string{
		"MOD_UNKNOWN", "MOD_SHOTGUN", "MOD_GAUNTLET", "MOD_MACHINEGUN", "MOD_GRENADE",
		"MOD_GRENADE_SPLASH", "MOD_ROCKET", "MOD_ROCKET_SPLASH", "MOD_PLASMA", "MOD_PLASMA_SPLASH",
		"MOD_RAILGUN", "MOD_LIGHTNING", "MOD_BFG", "MOD_BFG_SPLASH", "MOD_WATER",
		"MOD_SLIME", "MOD_LAVA", "MOD_CRUSH", "MOD_TELEFRAG", "MOD_FALLING",
		"MOD_SUICIDE", "MOD_TARGET_LASER", "MOD_TRIGGER_HURT", "MOD_NAIL", "MOD_CHAINGUN",
		"MOD_PROXIMITY_MINE", "MOD_KAMIKAZE", "MOD_JUICED", "MOD_GRAPPLE",
	}
	for _, mod := range meansOfDeath {
		means[mod] = 0
	}
}

// NewLeaderboard creates a leaderboard based on kill counts of players in a match.
func NewLeaderboard(match *data.MatchData) {
	// Copy players to leaderboard
	leaderboard := make([]string, len(match.Players))
	copy(leaderboard, match.Players)

	// Sort players based on kill count
	for i := 1; i < len(leaderboard); i++ {
		key := leaderboard[i]
		j := i - 1
		for j >= 0 && match.KillCount[leaderboard[j]] < match.KillCount[key] {
			leaderboard[j+1] = leaderboard[j]
			j--
		}
		leaderboard[j+1] = key
	}

	// Assign positions to leaderboard
	for i, player := range leaderboard {
		match.Leaderboard[i+1] = player
	}
}

// Parse parses the game log and returns a map of match data.
func Parse(log string) map[string]*data.MatchData {
	var waitgroup sync.WaitGroup
	matches := make(map[string]*data.MatchData)
	matchNumber := 0

	lines := strings.Split(log, "\n")

	for lineNumber, line := range lines {
		line = strings.TrimSpace(line)
		tokens := strings.Split(line, " ")

		if len(tokens) > 2 && tokens[1] == "InitGame:" {
			matchNumber++
			waitgroup.Add(1)
			newMatch := NewMatch(matches, matchNumber)

			go ExtractMatchData(newMatch, lines, lineNumber+1, &waitgroup)
		}
	}

	waitgroup.Wait()
	return matches
}

// ExtractMatchData extracts match data from log lines and updates the match object.
func ExtractMatchData(match *data.MatchData, lines []string, lineNumber int, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	for lineNumber < len(lines) {
		line := strings.TrimSpace(lines[lineNumber])
		tokens := strings.Split(line, " ")

		if len(tokens) > 1 {
			switch tokens[1] {
			case "Kill:":
				RegisterKill(match, tokens)
			case "ClientUserinfoChanged:":
				RegisterPlayer(match, tokens)
			case "InitGame:":
				NewLeaderboard(match)
				return
			}
		}

		lineNumber++
	}

	NewLeaderboard(match)
}
