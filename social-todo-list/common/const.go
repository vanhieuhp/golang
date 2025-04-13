package common

import "log"

type DbType int

const (
	CurrentUser         = "current_user"
	DbTypeItem   DbType = 1
	DbTypeUser   DbType = 2
	PluginDBMain        = "mysql"
)

func Recovery() {
	if err := recover(); err != nil {
		log.Println("Recovered: ", err)
	}
}

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin" || requester.GetRole() == "mod"
}
