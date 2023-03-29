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

// UserEvent 用户相关的事件。
type UserEvent struct {
	EventType string
	UserInfo  *UserInfo
}

// UserEventPublisher 发布用户事件的接口。
type UserEventPublisher interface {
	Publish(userEvent *UserEvent) error
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
