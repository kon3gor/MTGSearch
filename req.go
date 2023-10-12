package mtgsearch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

const (
	apiUrl   = "https://api.magicthegathering.io/v1/cards"
	pageSize = 100
)

func fetchCards(q cardsQuery, page int) (CardsCollection, int) {
	qs, _ := query.Values(q)
	reqUrl := fmt.Sprintf("%s?%s&page=%d", apiUrl, qs.Encode(), page)

	res, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}

	tcnt, _ := strconv.Atoi(res.Header.Get("Total-Count"))
	var resp CardsCollection
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		panic(err)
	}

	return resp, tcnt - pageSize*page

}
