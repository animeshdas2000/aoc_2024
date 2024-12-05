package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func InputReader() ([]int, []int, error) {
	// Open the TSV file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	r := bufio.NewScanner(file)
	Larr := make([]int, 0, 1000)
	Rarr := make([]int, 0, 1000)

	for r.Scan() {
		lines := r.Text()
		nums := strings.Fields(lines)
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error converting to int")
			return nil, nil, err
		}
		right, err := strconv.Atoi(nums[1])

		if err != nil {
			fmt.Println("Error converting to int")
			return nil, nil, err
		}

		Larr = append(Larr, left)
		Rarr = append(Rarr, right)
	}

	return Larr, Rarr, nil
}

func Day1A() {
	Llist, Rlist, err := InputReader()
	if err != nil {
		fmt.Println("Error reading TSV file")
	}

	sort.Ints(Llist)
	sort.Ints(Rlist)

	res := 0
	for i := range Llist {
		diff := Rlist[i] - Llist[i]
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}

	fmt.Printf("Final Answer: %d", res)

}

func Day1B() {
	Llist, Rlist, err := InputReader()
	if err != nil {
		fmt.Println("Error reading TSV file")
	}

	res := 0

	for i := range Llist {
		temp := 0
		cnt := 0
		for j := range Rlist {
			if Llist[i] == Rlist[j] {
				cnt++
			}
		}
		temp = Llist[i] * cnt
		res += temp
	}
	fmt.Printf("Final Answer: %d", res)

}

func main() {
	Day1A()
	Day1B()
}
