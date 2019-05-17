package document

// OpenIE contains records parsed by the OpenIE annotator: https://stanfordnlp.github.io/CoreNLP/openie.html
// Also see https://nlp.stanford.edu/software/openie.html
type OpenIE struct {
	Subject      string `json:"subject"`
	SubjectSpan  []int  `json:"subjectSpan"`
	Relation     string `json:"relation"`
	RelationSpan []int  `json:"relationSpan"`
	Object       string `json:"object"`
	ObjectSpan   []int  `json:"objectSpan"`
}
