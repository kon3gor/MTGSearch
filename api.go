package mtgsearch

type Card struct {
	Name     string  `json:"name"`
	ManaCost string  `json:"manaCost"`
	Cmc      float32 `json:"cmc"`
	Type     string  `json:"type"`
}

type CardsCollection struct {
	Cards []Card `json:"cards"`
}

var EmptyCollection = CardsCollection{}

type CardsCollector struct {
	mergeSame bool
	page      int
	remains   int
	query     cardsQuery
}

// NewCollector creates new collector for cards with specified settings
// mergeSame determines if cards with the same name should be merged
func NewCollector(mergeSame bool) *CardsCollector {
	return &CardsCollector{mergeSame: mergeSame, page: 0, remains: 42}
}

func (cc *CardsCollector) Collect(expr string) CardsCollection {
	rpn := parseExpression(expr)
	q := buildQuery(rpn)
	return cc.fetch(q, 1)
}

func (cc *CardsCollector) NextPage() CardsCollection {
	if cc.HasMore() && cc.page > 0 {
		return cc.fetch(cc.query, cc.page+1)
	} else {
		return EmptyCollection
	}
}

func (cc *CardsCollector) HasMore() bool {
	return cc.remains > 0
}

func (cc *CardsCollector) fetch(q cardsQuery, page int) CardsCollection {
	cards, remains := fetchCards(q, page)
	cc.page = page
	cc.remains = remains
	cc.query = q
	return cards
}
