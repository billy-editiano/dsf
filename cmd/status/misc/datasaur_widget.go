package misc

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/billy-editiano/dsfetch/internal/pkg/datasaur"
	"github.com/billy-editiano/dsfetch/internal/pkg/utils"
)

const (
	RED   = "#EF4444"
	GREEN = "#A3E635"
)

type Widget struct {
	Name     string `json:"name"`
	Color    string `json:"color"`
	FullText string `json:"full_text"`
}

func GetDatasaurWidgets() []Widget {
	keys, allResult := loadSortedDatasaurWidgets()
	widgets := []Widget{}
	for _, key := range keys {
		health := allResult[key]
		color := RED
		if health.Status == "up" {
			color = GREEN
		}
		widget := Widget{
			Name:     key,
			Color:    color,
			FullText: fmt.Sprintf("%s %s", key, health.Version),
		}
		widgets = append(widgets, widget)
	}
	return widgets
}

func loadDsfetchFile() map[string]*datasaur.HealthResponse {
	resultPath := utils.GetResultPath(false)
	content, err := os.ReadFile(resultPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return map[string]*datasaur.HealthResponse{}
	}
	allResult := map[string]*datasaur.HealthResponse{}

	err = json.Unmarshal(content, &allResult)
	if err != nil {
		fmt.Println("Error unmarshalling content:", err)
		return map[string]*datasaur.HealthResponse{}
	}
	return allResult
}

func loadSortedDatasaurWidgets() ([]string, map[string]*datasaur.HealthResponse) {
	allResult := loadDsfetchFile()
	keys := make([]string, 0, len(allResult))
	for k := range allResult {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys, allResult
}
