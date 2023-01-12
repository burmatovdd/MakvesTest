package methods

type Result struct {
	Result []Item `json:"result"`
}

type Item struct {
	Id   int      `json:"id"`
	Info []string `json:"info"`
}
