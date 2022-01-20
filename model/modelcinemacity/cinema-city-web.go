package modelcinemacity

type Request struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}

type Response struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
	Zip   string `json:"zip"`
}
