package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stamp), nil
}

type UserResponse struct {
	ID       int32    `json:"id"`
	Mobile   string   `json:"mobile"`
	NickName string   `json:"nick_name"`
	Gender   string   `json:"gender"`
	Birthday JsonTime `json:"birthday"`
}
