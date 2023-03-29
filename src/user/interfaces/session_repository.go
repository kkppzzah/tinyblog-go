// Package interfaces 定义各个接口。
package interfaces

// SessionRepository 用来操作会话存储的接口。
type SessionRepository interface {
	Create(sessionID string, userID int64) error
	Get(sessionID string) (int64, error)
	Delete(sessionID string) error
}
