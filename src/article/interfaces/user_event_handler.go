// Package interfaces 定义各个接口。
package interfaces

import "encoding/json"

//
const (
	UserEventTypeCreate = "create"
	UserEventTypeUpdate = "update"
	UserEventTypeDelete = "delete"
)

// UserInfo 用户信息。
type UserInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// UserEvent 用户相关的时间。
type UserEvent struct {
	EventType string
	UserInfo  *UserInfo
}

// UserEventHandler 用户事件处理。
type UserEventHandler interface {
}

// ToJSON 输出为JSON。
func (u *UserEvent) ToJSON() ([]byte, error) {
	bytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// FromJSON 从JSON中解析。
func (u *UserEvent) FromJSON(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, u)
}
