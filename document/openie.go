package document

type OpenIE struct {
	Subject      string `json:"subject"`
	SubjectSpan  []int  `json:"subjectSpan"`
	Relation     string `json:"relation"`
	RelationSpan []int  `json:"relationSpan"`
	Object       string `json:"object"`
	ObjectSpan   []int  `json:"objectSpan"`
}
