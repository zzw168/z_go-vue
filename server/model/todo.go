package model

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Msg_Data struct {
	ID     int `json:"id"`
	Vistor int `json:"vistor"`
	Member int `json:"member"`
	Qnum   int `json:"qnum"`
}
