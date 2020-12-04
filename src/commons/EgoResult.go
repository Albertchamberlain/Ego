package commons

//客户端服务器端数据的交互模版 ，大写可以被访问到
type EgoResult struct {
	Status int         //状态码,为200表示成功
	Data   interface{} //返回的数据
	Msg    string      //返回的消息
}
