package esmodel

type TestIndexMapping struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (TestIndexMapping) Index() string {
	return "test-index"
}
