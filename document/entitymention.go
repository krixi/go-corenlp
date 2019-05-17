package document

type EntityMention struct {
	DocTokenBegin        int    `json:"docTokenBegin"`
	DocTokenEnd          int    `json:"docTokenEnd"`
	TokenBegin           int    `json:"tokenBegin"`
	TokenEnd             int    `json:"tokenEnd"`
	Text                 string `json:"text"`
	CharacterOffsetBegin int    `json:"characterOffsetBegin"`
	CharacterOffsetEnd   int    `json:"characterOffsetEnd"`
	EntityType           string `json:"ner"`
	NormalizedEntityType string `json:"normalizedNER"`
}
