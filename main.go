package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	word_list := UpdateWordleList()
	for {
		top10Word := SortByExpectation(word_list)
		fmt.Println(top10Word)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Word : ")
		word, _ := reader.ReadString('\n')
		fmt.Print("Enter non perfect matches as a string : ")
		matches, _ := reader.ReadString('\n')
		fmt.Print("Enter non matches as a string : ")
		nonmatches, _ := reader.ReadString('\n')
		fmt.Print("Perfect matches : (y/n) ")
		flag, _ := reader.ReadString('\n')
		perfMatches := make(map[int]string)
		fmt.Println(flag)
		for strings.Contains(flag, "y") {
			fmt.Print("position of match (1 2 3 4 5) : ")
			pos, _ := reader.ReadString('\n')
			i, err := strconv.Atoi(pos[:1])
			if err != nil {
				panic(err)
			}
			fmt.Print("char of match : ")
			ch, _ := reader.ReadString('\n')

			perfMatches[i] = ch
			fmt.Print("More Perfect matches? : (y/n) ")
			flag, _ = reader.ReadString('\n')
		}
		word_list = ReduceList(word, matches, nonmatches, perfMatches, word_list)
	}

	// fmt.Println(ReduceList("pears", "e", "pars", make(map[int]string), word_list))

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Perfect matches : (y/n) ")
	// flag, _ := reader.ReadString('\n')
	// fmt.Println(flag)
	// for strings.Contains(flag, "y") {
	// 	flag, _ = reader.ReadString('\n')
	// }
}
