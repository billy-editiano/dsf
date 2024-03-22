package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/billy-editiano/dsfetch/cmd/status/misc"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	readFirst2Lines(scanner)

	for scanner.Scan() {
		line := scanner.Bytes()

		var prefix rune
		if line[0] == ',' {
			prefix = ','
			line = line[1:]
		}
		_ = prefix

		jsondata := []map[string]interface{}{}
		err := json.Unmarshal(line, &jsondata)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		widgets := []interface{}{}

		// append datasaur widgets
		datasaurWidgets := misc.GetDatasaurWidgets()
		for _, data := range datasaurWidgets {
			widgets = append(widgets, data)
		}

		// append default widgets
		for _, data := range jsondata {
			widgets = append(widgets, data)
			continue
		}

		// compile all widgets into json string array of widgets
		unifiedOutput, err := json.Marshal(widgets)
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		result := strings.Builder{}
		result.WriteString(string(unifiedOutput))
		result.WriteString(",")
		fmt.Println(result.String())
	}
}

func readFirst2Lines(scanner *bufio.Scanner) {
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
		if count == 2 {
			break
		}
	}
}
