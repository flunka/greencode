package endpoints

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

const OnlineGameEndpoint = "/onlinegame/calculate"

type Clan struct {
	NumberOfPlayers int `json:"numberOfPlayers"  binding:"required,min=1,max=1000"`
	Points          int `json:"points"  binding:"required,min=1,max=1000000"`
}

type Players struct {
	GroupCount int    `json:"groupCount"  binding:"required,min=1,max=1000"`
	Clans      []Clan `json:"clans"  binding:"required,max=20000"`
}

type ByGameRules []Clan

func (a ByGameRules) Len() int      { return len(a) }
func (a ByGameRules) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByGameRules) Less(i, j int) bool {
	// sort by number of points first
	if a[i].Points != a[j].Points {
		return a[i].Points > a[j].Points
	}
	// sort by number of players
	return a[i].NumberOfPlayers < a[j].NumberOfPlayers
}

func getGroup(players Players, seen []bool) []Clan {
	group := make([]Clan, 0, len(players.Clans))
	groupSize := 0
	for i := 0; i < len(players.Clans) && groupSize < players.GroupCount; i++ {
		if seen[i] {
			continue
		}
		if players.Clans[i].NumberOfPlayers+groupSize > players.GroupCount {
			continue
		}
		group = append(group, players.Clans[i])
		groupSize += players.Clans[i].NumberOfPlayers
		seen[i] = true
	}
	return group
}

func isAllSeen(seen []bool) bool {
	for i := 0; i < len(seen); i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}

func getGroupOrder(players Players) [][]Clan {
	clans := make([][]Clan, 0, len(players.Clans))
	seen := make([]bool, len(players.Clans))
	for !isAllSeen(seen) {
		clans = append(clans, getGroup(players, seen))
	}
	return clans
}

func CaltulateOrder(c *gin.Context) {
	var players Players
	if err := c.ShouldBindJSON(&players); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sort.Sort(ByGameRules(players.Clans))
	order := getGroupOrder(players)
	c.JSON(http.StatusOK, order)
}
