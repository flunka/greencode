package endpoints

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

const ATMEndpoint = "/atms/calculateOrder"

type Task struct {
	Region      int32  `json:"region" binding:"required,min=1,max=9999"`
	RequestType string `json:"requestType" binding:"required,oneof='STANDARD' 'PRIORITY' 'SIGNAL_LOW' 'FAILURE_RESTART'"`
	AtmId       int32  `json:"atmId" binding:"required,min=1,max=9999"`
}

var taskPriorities = map[string]int{
	"FAILURE_RESTART": 0,
	"PRIORITY":        1,
	"SIGNAL_LOW":      2,
	"STANDARD":        3,
}

type ATM struct {
	Region int32 `json:"region" binding:"required,min=1,max=9999"`
	AtmId  int32 `json:"atmId" binding:"required,min=1,max=9999"`
}

type ByRules []Task

func (a ByRules) Len() int      { return len(a) }
func (a ByRules) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRules) Less(i, j int) bool {
	// sort by region first
	if a[i].Region != a[j].Region {
		return a[i].Region < a[j].Region
	}
	// sort by task priorities
	return taskPriorities[a[i].RequestType] < taskPriorities[a[j].RequestType]
}

func getAtms(tasks []Task) []ATM {
	seen := make(map[ATM]bool)
	atms := make([]ATM, 0, len(tasks))
	for _, task := range tasks {
		atm := ATM{AtmId: task.AtmId, Region: task.Region}
		if seen[atm] {
			continue
		}
		atms = append(atms, atm)
		seen[atm] = true
	}
	return atms
}

func Order(c *gin.Context) {
	var tasks []Task
	if err := c.ShouldBindJSON(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sort.Sort(ByRules(tasks))
	atms := getAtms(tasks)
	c.JSON(http.StatusOK, atms)
}
