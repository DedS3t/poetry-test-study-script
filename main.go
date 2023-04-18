package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/deds3t/poem-study/app/models"
	"github.com/deds3t/poem-study/pkg"
)

func removeSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	file, err := ioutil.ReadFile("poems.json")
	if err != nil {
		panic(err)
	}

	var poem []models.PoemDto
	err = json.Unmarshal(file, &poem)
	if err != nil {
		panic(err)
	}

	poemList := []*pkg.Poem{}

	for _, p := range poem {
		poemList = append(poemList, pkg.CreatePoem(p))
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// clearInputBuffer()
		var response string

		p := poemList[rand.Intn(len(poemList))]
		stanza := ""
		for stanza == "" {
			stanza = p.GetRandomPart()
		}

		fmt.Println(stanza)
		fmt.Println("\nIdentify name and author of excerpt above")

		scanner.Scan()
		response = scanner.Text()

		if response == "exit" {
			break
		}

		acc := pkg.CosineSimilarity(removeSpaces(response), removeSpaces(p.Name+p.Author))
		var res string
		if acc > .75 {
			res = "correct"
		} else {
			res = "incorrect"
		}

		fmt.Printf("You are %s with an accuracy of %f\n", res, acc)
		if res == "incorrect" {
			fmt.Printf("It was %s by %s", p.Name, p.Author)
		}

		fmt.Print("\n\n\n")
		time.Sleep(2 * time.Second)
	}
}
