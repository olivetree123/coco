package coco

// Result 返回结果
type Result struct {
	Type       string      `json:"-"` // object or binary
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	BinaryData []byte      `json:"-"`
	Message    string      `json:"message"`
	FileName   string      `json:"-"`
}

func APIResponse(data interface{}) Result {
	return Result{
		Type:       "object",
		Code:       0,
		Data:       data,
		Message:    "",
		BinaryData: nil,
		FileName:   "",
	}
}

func ErrorResponse(code int) Result {
	return Result{
		Type:       "object",
		Code:       code,
		Data:       nil,
		Message:    "",
		BinaryData: nil,
		FileName:   "",
	}
}

func FileResponse(fileName string, content []byte) Result {
	return Result{
		Type:       "binary",
		Code:       0,
		Data:       nil,
		Message:    "",
		BinaryData: content,
		FileName:   fileName,
	}
}
