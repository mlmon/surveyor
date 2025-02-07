package source

type Fn func() (*Records, error)

type RecordSet struct {
	Records []*Records
}

type Records struct {
	Source  string  `json:"source"`
	Entries Entries `json:"entries"`
}

type Entries []Record

func (a Entries) Len() int           { return len(a) }
func (a Entries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Entries) Less(i, j int) bool { return a[i].Key < a[j].Key }

type Record struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
