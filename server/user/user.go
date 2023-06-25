package user

import (
	"context"
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type User struct {
	UUID   string
	UserID string
	Ctx    context.Context
	Cancel context.CancelFunc
}

var UserMap sync.Map

func UpdateUserUUID(user *User) {
	newUUID := uuid.NewV4().String()
	user.UUID = newUUID
	UserMap.Store(newUUID, *user)

	if user.Cancel != nil {
		user.Cancel() // 取消之前的删除等待
	}

	// 为新的用户创建一个新的 context
	ctx, cancel := context.WithCancel(context.Background())
	user.Ctx = ctx
	user.Cancel = cancel

	// 启动新的 Goroutine 来在 5 秒后删除用户
	go deleteUserAuto(user)
}

func NewUser(userID string) User {
	user := User{
		UUID:   uuid.NewV4().String(),
		UserID: userID,
	}

	ctx, cancel := context.WithCancel(context.Background())
	user.Ctx = ctx
	user.Cancel = cancel

	UserMap.Range(func(key, value interface{}) bool {
		existingUser := value.(*User)
		if existingUser.UserID == userID {
			UpdateUserUUID(existingUser)
			return false
		}
		return true
	})

	UserMap.Store(user.UUID, &user)

	// 启动新的 Goroutine 来在 5 秒后删除用户
	go deleteUserAuto(&user)

	return user
}

func GetUser(uuid string) (*User, bool) {
	user, ok := UserMap.Load(uuid)
	if !ok {
		return nil, false
	}

	return user.(*User), true
}

// 通过UUID删除用户
func DeleteUser(uuid string) {
	UserMap.Delete(uuid)
}

// 启动新的 Goroutine 来在 5 秒后删除用户
func deleteUserAuto(user *User) {
	select {
	case <-time.After(5 * time.Second):
		DeleteUser(user.UUID)
	case <-user.Ctx.Done():
		// 如果接收到了取消信号，就不删除用户
	}
}
