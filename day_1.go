package main

import (
	"fmt"
	"log"
	"strconv"
	"williamthewilson/aoc-go-2022/io"
)

func Day_1_1() {
	io.ProcessLines("./assets/day_1.txt", func() (func(string), func()) {
		longest_total := 0
		current_total := 0
		return func(line string) {
				if line == "" {
					if current_total > longest_total {
						longest_total = current_total
					}
					current_total = 0
					return
				}
				intVal, err := strconv.Atoi(line)
				if err != nil {
					log.Println(err)
				}
				current_total += intVal
			}, func() {
				fmt.Printf("%v\n", longest_total)
			}
	})
}
