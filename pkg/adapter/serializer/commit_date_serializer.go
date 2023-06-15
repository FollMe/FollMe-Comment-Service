package serializer

import "time"

type UpdateCommitDateReq struct {
	Date time.Time `json:"date" validate:"required"`
}
