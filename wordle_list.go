package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

var (
	link              = "https://www.nytimes.com/games/wordle/main.4d41d2be.js"
	list_var_start    = "var Ma=["
	list_var_end      = "],"
	list_var_interval = "],Oa=["
)

func UpdateWordleList() []string {
	st := getWordleListFile()
	return parseWordleString(st)
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func getWordleListFile() string {
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	r := bufio.NewReader(resp.Body)
	if s, e := Readln(r); e != nil {
		panic(e)
	} else {
		return s
	}
}

func parseWordleString(st string) []string {
	i := strings.Index(st, list_var_start)
	st = st[i+8 : len(st)-1]

	st = strings.Replace(st, list_var_interval, ",", 1)

	i = strings.Index(st, list_var_end)
	st = st[:i-1]

	st = strings.ReplaceAll(st, "\"", "")

	word_list := strings.Split(st, ",")

	return word_list
}
