// nake is a simple makefile target interpreter.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	// data contains the full contents from $PWD/Makefile
	data    string
	// targets are the targets filtered down from the data in the makefile.
	targets []string
)

// init will read in the data from the makefile for processing.
func init() {
	file, err := os.Open("Makefile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	data = string(b)
}

// process will filter the file contents down to populate the targets variable.
func process() {
	// Check each line for a target.
	for _, line := range strings.Split(data, "\n") {
		// For the purpose of this tool, we're searching for explicitly declared
		// PHONY targets, so we need to match a regex and filter out target prerequisites.
		if check, _ := regexp.MatchString("^.PHONY: .*", line); check {
			// This line matches our regex, so we need to filter it down for accuracy.
			target := strings.TrimLeft(line, ".PHONY: ")
			if !strings.HasPrefix(target, "$") {
				// Add our target to the list of targets.
				targets = append(targets, target)
			}
		}
	}
	// Sort the targets.
	sort.Strings(targets)
}

// list prints our list of targets.
func list() {
	for i := range targets {
		fmt.Println(targets[i])
	}
}

// main is the main application entrypoint.
func main() {
	process()
	list()
}
