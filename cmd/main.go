package main

import (
	"bufio"
	"fmt"
	"os"

	mtgsearch "github.com/kon3gor/MTGSearch"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expr, _ := reader.ReadString('\n')

	col := mtgsearch.NewCollector(false)
	res := col.Collect(expr)
	for _, c := range res.Cards {
		fmt.Println(c.Name)
	}
	var i int
	for col.HasMore() && i < 2 {
		res = col.NextPage()
		for _, c := range res.Cards {
			fmt.Println(c.Name)
		}
		i++
	}
}
