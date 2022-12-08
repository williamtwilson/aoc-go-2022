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

func Day_1_2() {
	io.ProcessLines("./assets/day_1.txt", func() (func(string), func()) {
		longest_total := 0
		second_longest := 0
		third_longest := 0
		current_total := 0
		return func(line string) {
				if line == "" {
					if current_total > third_longest {
						if current_total > longest_total {
							third_longest = second_longest
							second_longest = longest_total
							longest_total = current_total
						} else if current_total > second_longest {
							third_longest = second_longest
							second_longest = current_total
						} else {
							third_longest = current_total
						}
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
				fmt.Printf("%v\n%v\n%v\n", longest_total, second_longest, third_longest)
				fmt.Printf("%v\n", longest_total+second_longest+third_longest)
			}
	})
}
