package document

// A Document contains a result of annotation.
type Document struct {
	Sentences Sentences          `json:"sentences"`
	Corefs    map[string][]Coref `json:"corefs"`
}
