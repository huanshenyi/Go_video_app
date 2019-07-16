package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`  //内部ディバグ用
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed =
		ErrResponse{HttpSC:400, Error: Err{Error: "Request body is not correct ", ErrorCode:"001"}}

	ErrorNoAuthUser =
	 	ErrResponse{HttpSC:401,Error:Err{Error:"User authentication failed.", ErrorCode:"002"}}
)