package jstype

type JSResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data,omitempty"`
}

func Success(data any) (resp JSResp) {
	resp.Success = true
	resp.Data = data
	return resp
}

func Error(error error) (resp JSResp) {
	resp.Msg = error.Error()
	return resp
}

type CryptoRequest struct {
	Env  string `json:"env"`
	Data string `json:"data"`
}
