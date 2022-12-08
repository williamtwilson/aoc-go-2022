package io

import (
	"bufio"
	"log"
	"os"
)

func ProcessLines(file string, lineProcessorGetter func() (func(string), func())) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lineProcessor, finalFunc := lineProcessorGetter()

	for scanner.Scan() {
		lineProcessor(scanner.Text())
	}

	finalFunc()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
