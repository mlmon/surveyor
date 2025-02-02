package source

type Fn func() (*Records, error)

type Records struct {
	Source  string  `json:"source"`
	Entries Entries `json:"entries"`
}

type Entries []Record

type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
