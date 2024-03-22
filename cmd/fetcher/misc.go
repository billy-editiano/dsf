package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/billy-editiano/dsfetch/internal/pkg/datasaur"
	"github.com/billy-editiano/dsfetch/internal/pkg/utils"
)

type HealthResponseWithName struct {
	Name   string                   `json:"name"`
	Health *datasaur.HealthResponse `json:"health"`
}

func fetchHost(response chan HealthResponseWithName, wg *sync.WaitGroup, host datasaur.Host) {
	defer wg.Done()

	resp, err := fetchData(host)

	if err != nil {
		fmt.Println("Error fetching data from", host.Name, ":", err)
		response <- HealthResponseWithName{
			Name: host.Name,
			Health: &datasaur.HealthResponse{
				Status:  "down",
				Version: "-",
			},
		}
	}

	fmt.Println(host.Name, "=>", resp.Version, resp.Status)
	response <- HealthResponseWithName{
		Name:   host.Name,
		Health: resp,
	}
}

func fetchData(host datasaur.Host) (*datasaur.HealthResponse, error) {
	url := host.Host

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	parsedResp := &datasaur.HealthResponse{}
	err = json.Unmarshal(body, parsedResp)
	if err != nil {
		return nil, err
	}

	return parsedResp, nil
}

func writeResult(response chan HealthResponseWithName) {
	allResult := map[string]*datasaur.HealthResponse{}
	for hr := range response {
		allResult[hr.Name] = hr.Health
	}
	jsonResult, _ := json.MarshalIndent(allResult, "", "  ")
	err := writeToFile(utils.GetResultPath(true), jsonResult)

	if err != nil {
		fmt.Println("Error writing result to file:", err)
	}
	fmt.Println("Result written to", utils.GetResultPath(false))
}

func writeToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}
