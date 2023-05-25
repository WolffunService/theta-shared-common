package jwtauth

import (
	"github.com/WolffunService/thetan-shared-common/common"
	"github.com/kataras/iris/v12/middleware/jwt"
)

type ThetanUserClaims struct {
	jwt.Claims `json:",inline"`
	UserId     string          `json:"user_id"`
	Role       common.UserRole `json:"role"`
}

func (u ThetanUserClaims) GetID() string {
	return u.UserId
}

func (u ThetanUserClaims) GetUsername() string {
	return u.Subject
}

func (u ThetanUserClaims) GetRoles() []string {
	// TODO: Refactor with enum
	role := "User"
	switch u.Role {
	case common.ADMIN:
		role = "Admin"
	case common.ROOT:
		role = "Root"
	}
	return []string{role}
}
