package modelcinemacity

type Request struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}
