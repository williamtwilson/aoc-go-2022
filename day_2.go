package main

import (
	"fmt"
	"williamthewilson/aoc-go-2022/io"
)

func Day_2_1() {
	io.ProcessLines("./assets/day_2.txt", func() (func(string), func()) {
		score := 0
		return func(line string) {
				if line == "" {
					return
				}

				if line[2] == 'X' {
					score += 1
					if line[0] == 'C' {
						score += 6
					} else if line[0] == 'A' {
						score += 3
					}
				} else if line[2] == 'Y' {
					score += 2
					if line[0] == 'A' {
						score += 6
					} else if line[0] == 'B' {
						score += 3
					}
				} else if line[2] == 'Z' {
					score += 3
					if line[0] == 'B' {
						score += 6
					} else if line[0] == 'C' {
						score += 3
					}
				}
			}, func() {
				fmt.Printf("%v\n", score)
			}
	})
}

func Day_2_2() {
	io.ProcessLines("./assets/day_2.txt", func() (func(string), func()) {
		score := 0
		return func(line string) {
				if line == "" {
					return
				}

				if line[0] == 'A' {
					if line[2] == 'X' {
						score += 3
					} else if line[2] == 'Y' {
						score += 4
					} else {
						score += 8
					}
				} else if line[0] == 'B' {
					if line[2] == 'X' {
						score += 1
					} else if line[2] == 'Y' {
						score += 5
					} else {
						score += 9
					}
				} else {
					if line[2] == 'X' {
						score += 2
					} else if line[2] == 'Y' {
						score += 6
					} else {
						score += 7
					}
				}
			}, func() {
				fmt.Printf("%v\n", score)
			}
	})
}
