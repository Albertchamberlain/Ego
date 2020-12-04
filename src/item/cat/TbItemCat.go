package cat

type TbItemCat struct {
	Id        int
	ParentId  int
	Name      string
	Status    byte
	SortOrder int8
	IsParent  bool
	Created   string
	Updated   string
}
