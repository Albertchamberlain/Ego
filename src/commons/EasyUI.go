package commons

type Datagrid struct {
	Rows  interface{} `json:"rows"`
	Total int         `json:"total"`
}

type EasyUITree struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}
