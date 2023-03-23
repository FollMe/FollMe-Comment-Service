package serializer

type MetaRes struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
}

type HttpRes struct {
	Meta MetaRes     `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func NewSuccessHttpRes(mess string, data interface{}) HttpRes {
	return HttpRes{
		Meta: MetaRes{
			Ok:      true,
			Message: mess,
		},
		Data: data,
	}
}

func NewFailHttpRes(mess string) HttpRes {
	if mess == "" {
		mess = "Có lỗi xảy ra, vui lòng thử lại!"
	}
	return HttpRes{
		Meta: MetaRes{
			Ok:      false,
			Message: mess,
		},
	}
}
