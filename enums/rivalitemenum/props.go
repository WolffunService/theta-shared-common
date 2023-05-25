package rivalitemenum

import "strings"

type props struct {
	Season       string
	ExclusiveFor string
	SSPassType   string
	Rarity       string
	SSPassId     string
	SSPPageID    string
	SSPPrice     string
}

var Props = props{
	Season:       strings.ToLower("season"),
	ExclusiveFor: strings.ToLower("exclusiveFor"),
	SSPassType:   strings.ToLower("ssPassType"),
	Rarity:       strings.ToLower("rarity"),
	SSPassId:     strings.ToLower("ssPassId"),
	SSPPageID:    strings.ToLower("sspPageId"),
	SSPPrice:     strings.ToLower("sspPrice"),
}

type exclusiveFor struct {
	SeasonPass string
}

var ExclusiveFor = exclusiveFor{
	SeasonPass: "SeasonPass",
}

type ssPassType struct {
	Free    string
	Premium string
}

var SSPassType = ssPassType{
	Free:    "free",
	Premium: "premium",
}
