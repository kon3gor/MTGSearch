package mtgsearch

type cardsQuery struct {
	Cmc       string    `url:"cmc,omitempty"`
	Oracle    parameter `url:"text,omitempty"`
	Type      parameter `url:"type,omitempty"`
	Supertype parameter `url:"types,omitempty"`
	Subtype   parameter `url:"subtypes,omitempty"`
	Colors    parameter `url:"colors,omitempty"`
	Rarity    parameter `url:"rarity,omitempty"`
}

var do parameter = empty{}

var dq cardsQuery = cardsQuery{
	Type:      do,
	Oracle:    do,
	Subtype:   do,
	Supertype: do,
	Rarity:    do,
	Colors:    do,
}

func simpleQuery(key, v string) cardsQuery {
	q := dq
	switch key {
	case cmc:
		q.Cmc = replaceRelativeCmc(v)
	case tipe:
		q.Type = value{v}
	case oracle:
		q.Oracle = value{v}
	case supertype:
		q.Supertype = value{v}
	case subtype:
		q.Subtype = value{v}
	case rarity:
		q.Rarity = value{v}
	case color:
		q.Colors = value{v}
	}
	return q
}

func (q cardsQuery) combine(o cardsQuery) cardsQuery {
	if o.Cmc != dq.Cmc && q.Cmc == dq.Cmc {
		q.Cmc = o.Cmc
	}

	if v := q.Type.combine(o.Type); v != emptyParameter {
		q.Type = v
	}

	if v := q.Oracle.combine(o.Oracle); v != emptyParameter {
		q.Oracle = v
	}

	if v := q.Supertype.combine(o.Supertype); v != emptyParameter {
		q.Supertype = v
	}

	if v := q.Subtype.combine(o.Subtype); v != emptyParameter {
		q.Subtype = v
	}

	if v := q.Colors.combine(o.Colors); v != emptyParameter {
		q.Colors = v
	}

	if v := q.Rarity.combine(o.Rarity); v != emptyParameter {
		q.Rarity = v
	}

	return q
}

func (q cardsQuery) pipe(o cardsQuery) cardsQuery {
	if o.Cmc != dq.Cmc && q.Cmc == dq.Cmc {
		q.Cmc = o.Cmc
	}

	if v := q.Type.pipe(o.Type); v != emptyParameter {
		q.Type = v
	}

	if v := q.Oracle.pipe(o.Oracle); v != emptyParameter {
		q.Oracle = v
	}

	if v := q.Supertype.pipe(o.Supertype); v != emptyParameter {
		q.Supertype = v
	}

	if v := q.Subtype.pipe(o.Subtype); v != emptyParameter {
		q.Subtype = v
	}

	if v := q.Colors.pipe(o.Colors); v != emptyParameter {
		q.Colors = v
	}

	if v := q.Rarity.pipe(o.Rarity); v != emptyParameter {
		q.Rarity = v
	}

	return q
}
