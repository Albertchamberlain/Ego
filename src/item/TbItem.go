package item

//商品
type TbItem struct {
	Id        int
	Title     string
	SellPoint string
	Price     int
	Num       int
	Barcode   string
	Image     string
	Cid       int
	Status    int8
	Created   string
	Updated   string
}

type TbItemChild struct {
	TbItem
	CategoryName string
}

type TbItemDescChild struct {
	TbItem
	CategoryName string
	Desc         string
}
