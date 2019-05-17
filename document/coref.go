package document

// Coref represents the co-references parsed from the text. See: https://stanfordnlp.github.io/CoreNLP/coref.html
type Coref struct {
	ID                      int    `json:"id"`
	Text                    string `json:"text"`
	Type                    string `json:"type"`
	Number                  string `json:"number"`
	Gender                  string `json:"gender"`
	Animacy                 string `json:"animacy"`
	StartIndex              int    `json:"startIndex"`
	EndIndex                int    `json:"endIndex"`
	HeadIndex               int    `json:"headIndex"`
	SentenceNum             int    `json:"sentNum"`
	Position                []int  `json:"position"`
	IsRepresentativeMention bool   `json:"isRepresentativeMention"`
}
