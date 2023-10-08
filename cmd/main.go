package main

import (
	"bufio"
	"fmt"
	"os"

	mtgsearch "github.com/ko3gor/MTGSearch"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')

	col := mtgsearch.NewCollector(false)
	res := col.Collect(expr)
	for _, c := range res.Cards {
		fmt.Println(c.Name)
	}
	for col.HasMore() {
		res = col.NextPage()
		for _, c := range res.Cards {
			fmt.Println(c.Name)
		}
	}
}
