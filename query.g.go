package mtgsearch

const (
	name          = "name"
	layout        = "layout"
	cmc           = "cmc"
	color         = "color"
	colorIdentity = "colorIdentity"
	fullType      = "fullType"
	supertype     = "supertype"
	tipe          = "type"
	subType       = "subType"
	rarity        = "rarity"
	setCode       = "setCode"
	setName       = "setName"
	oracle        = "oracle"
	flavor        = "flavor"
	artist        = "artist"
	number        = "number"
	power         = "power"
	toughness     = "toughness"
	loyalty       = "loyalty"
	language      = "language"
	gameFormat    = "gameFormat"
	legality      = "legality"
)

type cardsQuery struct {
	Name          parameter `url:"name,omitempty"`
	Layout        parameter `url:"layout,omitempty"`
	Cmc           string    `url:"cmc,omitempty"`
	Color         parameter `url:"colors,omitempty"`
	Coloridentity parameter `url:"colorIdentity,omitempty"`
	Fulltype      parameter `url:"type,omitempty"`
	Supertype     parameter `url:"supertypes,omitempty"`
	Tipe          parameter `url:"types,omitempty"`
	Subtype       parameter `url:"subtypes,omitempty"`
	Rarity        parameter `url:"rarity,omitempty"`
	Setcode       parameter `url:"set,omitempty"`
	Setname       parameter `url:"setName,omitempty"`
	Oracle        parameter `url:"text,omitempty"`
	Flavor        parameter `url:"flavor,omitempty"`
	Artist        parameter `url:"artist,omitempty"`
	Number        parameter `url:"number,omitempty"`
	Power         parameter `url:"power,omitempty"`
	Toughness     parameter `url:"toughness,omitempty"`
	Loyalty       parameter `url:"loyalty,omitempty"`
	Language      parameter `url:"language,omitempty"`
	Gameformat    parameter `url:"gameFormat,omitempty"`
	Legality      parameter `url:"legality,omitempty"`
}

var dp parameter = empty{}

var dq cardsQuery = cardsQuery{
	Name:          dp,
	Layout:        dp,
	Color:         dp,
	Coloridentity: dp,
	Fulltype:      dp,
	Supertype:     dp,
	Tipe:          dp,
	Subtype:       dp,
	Rarity:        dp,
	Setcode:       dp,
	Setname:       dp,
	Oracle:        dp,
	Flavor:        dp,
	Artist:        dp,
	Number:        dp,
	Power:         dp,
	Toughness:     dp,
	Loyalty:       dp,
	Language:      dp,
	Gameformat:    dp,
	Legality:      dp,
}

func simpleCardsQuery(k, v string) cardsQuery {
	q := dq
	switch k {
	case name:
		q.Name = value{v}
	case layout:
		q.Layout = value{v}
	case cmc:
		q.Cmc = v
	case color:
		q.Color = value{v}
	case colorIdentity:
		q.Coloridentity = value{v}
	case fullType:
		q.Fulltype = value{v}
	case supertype:
		q.Supertype = value{v}
	case tipe:
		q.Tipe = value{v}
	case subType:
		q.Subtype = value{v}
	case rarity:
		q.Rarity = value{v}
	case setCode:
		q.Setcode = value{v}
	case setName:
		q.Setname = value{v}
	case oracle:
		q.Oracle = value{v}
	case flavor:
		q.Flavor = value{v}
	case artist:
		q.Artist = value{v}
	case number:
		q.Number = value{v}
	case power:
		q.Power = value{v}
	case toughness:
		q.Toughness = value{v}
	case loyalty:
		q.Loyalty = value{v}
	case language:
		q.Language = value{v}
	case gameFormat:
		q.Gameformat = value{v}
	case legality:
		q.Legality = value{v}
	}
	return q
}

func (q cardsQuery) combine(o cardsQuery) cardsQuery {
	if v := q.Name.combine(o.Name); v != emptyParameter {
		q.Name = v
	}

	if v := q.Layout.combine(o.Layout); v != emptyParameter {
		q.Layout = v
	}

	if v := q.Color.combine(o.Color); v != emptyParameter {
		q.Color = v
	}

	if v := q.Coloridentity.combine(o.Coloridentity); v != emptyParameter {
		q.Coloridentity = v
	}

	if v := q.Fulltype.combine(o.Fulltype); v != emptyParameter {
		q.Fulltype = v
	}

	if v := q.Supertype.combine(o.Supertype); v != emptyParameter {
		q.Supertype = v
	}

	if v := q.Tipe.combine(o.Tipe); v != emptyParameter {
		q.Tipe = v
	}

	if v := q.Subtype.combine(o.Subtype); v != emptyParameter {
		q.Subtype = v
	}

	if v := q.Rarity.combine(o.Rarity); v != emptyParameter {
		q.Rarity = v
	}

	if v := q.Setcode.combine(o.Setcode); v != emptyParameter {
		q.Setcode = v
	}

	if v := q.Setname.combine(o.Setname); v != emptyParameter {
		q.Setname = v
	}

	if v := q.Oracle.combine(o.Oracle); v != emptyParameter {
		q.Oracle = v
	}

	if v := q.Flavor.combine(o.Flavor); v != emptyParameter {
		q.Flavor = v
	}

	if v := q.Artist.combine(o.Artist); v != emptyParameter {
		q.Artist = v
	}

	if v := q.Number.combine(o.Number); v != emptyParameter {
		q.Number = v
	}

	if v := q.Power.combine(o.Power); v != emptyParameter {
		q.Power = v
	}

	if v := q.Toughness.combine(o.Toughness); v != emptyParameter {
		q.Toughness = v
	}

	if v := q.Loyalty.combine(o.Loyalty); v != emptyParameter {
		q.Loyalty = v
	}

	if v := q.Language.combine(o.Language); v != emptyParameter {
		q.Language = v
	}

	if v := q.Gameformat.combine(o.Gameformat); v != emptyParameter {
		q.Gameformat = v
	}

	if v := q.Legality.combine(o.Legality); v != emptyParameter {
		q.Legality = v
	}

	return q
}

func (q cardsQuery) pipe(o cardsQuery) cardsQuery {
	if v := q.Name.pipe(o.Name); v != emptyParameter {
		q.Name = v
	}

	if v := q.Layout.pipe(o.Layout); v != emptyParameter {
		q.Layout = v
	}

	if v := q.Color.pipe(o.Color); v != emptyParameter {
		q.Color = v
	}

	if v := q.Coloridentity.pipe(o.Coloridentity); v != emptyParameter {
		q.Coloridentity = v
	}

	if v := q.Fulltype.pipe(o.Fulltype); v != emptyParameter {
		q.Fulltype = v
	}

	if v := q.Supertype.pipe(o.Supertype); v != emptyParameter {
		q.Supertype = v
	}

	if v := q.Tipe.pipe(o.Tipe); v != emptyParameter {
		q.Tipe = v
	}

	if v := q.Subtype.pipe(o.Subtype); v != emptyParameter {
		q.Subtype = v
	}

	if v := q.Rarity.pipe(o.Rarity); v != emptyParameter {
		q.Rarity = v
	}

	if v := q.Setcode.pipe(o.Setcode); v != emptyParameter {
		q.Setcode = v
	}

	if v := q.Setname.pipe(o.Setname); v != emptyParameter {
		q.Setname = v
	}

	if v := q.Oracle.pipe(o.Oracle); v != emptyParameter {
		q.Oracle = v
	}

	if v := q.Flavor.pipe(o.Flavor); v != emptyParameter {
		q.Flavor = v
	}

	if v := q.Artist.pipe(o.Artist); v != emptyParameter {
		q.Artist = v
	}

	if v := q.Number.pipe(o.Number); v != emptyParameter {
		q.Number = v
	}

	if v := q.Power.pipe(o.Power); v != emptyParameter {
		q.Power = v
	}

	if v := q.Toughness.pipe(o.Toughness); v != emptyParameter {
		q.Toughness = v
	}

	if v := q.Loyalty.pipe(o.Loyalty); v != emptyParameter {
		q.Loyalty = v
	}

	if v := q.Language.pipe(o.Language); v != emptyParameter {
		q.Language = v
	}

	if v := q.Gameformat.pipe(o.Gameformat); v != emptyParameter {
		q.Gameformat = v
	}

	if v := q.Legality.pipe(o.Legality); v != emptyParameter {
		q.Legality = v
	}

	return q
}
