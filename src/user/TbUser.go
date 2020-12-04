package user

//属性跨包访问 1 需要转换为json  2 可能出现跨包访问
type TbUser struct {
	Id       int64
	Username string
	Password string
	Phone    string
	Email    string
	Created  string
	Updated  string
}
