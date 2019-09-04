package proto


// rpc_proto: 接口
type tp interface {
	SayHello(HelloReq) HelloRes
	SayBye(ByeReq) ByeRes
}

// rpc_proto: 状态值
type RetStatus int

// rpc_proto: 状态值
const (
	RetStatus_Failure RetStatus = 0
	RetStatus_Success RetStatus = 1
)

//rpc_proto:返回状态
type CommonRes struct {
	Status RetStatus
}

// rpc_proto: helloReq
type HelloReq struct {
	Name string
	Content string
}
// rpc_proto: HelloRes
type HelloRes struct {
	Status RetStatus
	Content string
}

// rpc_proto: ByeReq
type ByeReq struct {
	Name string
	Content string
}
//rpc_proto: ByeRes
type ByeRes struct {
	Status RetStatus
	Content string
}