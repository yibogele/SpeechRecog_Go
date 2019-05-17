package util

import (
	"bufio"
	"os"
	"strings"
)

const (
	ActionConfFile = "/config/verbs.conf"
)

func ReadConf(filename string) (actions []string) {
	file, err := os.Open(filename)
	if err != nil {
		Log.Fatal(err)
	}
	defer file.Close()

	actions = []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		actionstr := strings.TrimSpace(scanner.Text())
		actions = append(actions, actionstr)
		//fmt.Println(actionstr)
	}

	if err := scanner.Err(); err != nil {
		Log.Fatal(err)
	}

	return
}
