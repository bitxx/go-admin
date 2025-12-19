package response

type Response struct {
	// 数据集
	RequestId string `json:"requestId,omitempty"`
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"msg,omitempty"`
}

type response struct {
	Response
	Data interface{} `json:"data"`
}

type Page struct {
	Count     int64 `json:"count"`
	PageIndex int   `json:"pageIndex"`
	PageSize  int   `json:"pageSize"`
}

type page struct {
	Page
	Extend interface{} `json:"extend"`
	List   interface{} `json:"list"`
}

func (e *response) SetData(data interface{}) {
	e.Data = data
}

func (e response) Clone() Responses {
	return &e
}

func (e *response) SetTraceID(id string) {
	e.RequestId = id
}

func (e *response) SetMsg(s string) {
	e.Msg = s
}

func (e *response) SetCode(code int) {
	e.Code = code
}

func (e *response) SetSuccess(success bool) {
	/*if !success {
		e.PkgType = "error"
	}*/
}
