package param

type TbItemParam struct {
	Id        int    `json:"id"`
	ItemCatId int    `json:"itemCatId"`
	ParamData string `json:"paramData"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}

type TbItemParamCat struct {
	TbItemParam
	CatName string `json:"catName"`
}
