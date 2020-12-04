package desc

func Insert(t TbItemDesc) int {
	return insertDescDao(t)
}

func SelByIdService(id int) *TbItemDesc {
	return selByIdDao(id)
}
