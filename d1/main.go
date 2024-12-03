package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func part1() (total int) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1, list2 []int

	var num1, num2 int
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &num1, &num2)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	slices.Sort(list1)
	slices.Sort(list2)

	total = 0
	for i:= 0; i < len(list1); i++ {
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return
}

func main() {
	fmt.Println(part1())
}