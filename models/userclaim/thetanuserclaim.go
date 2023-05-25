package userclaim

import (
	"github.com/WolffunService/thetan-shared-common/enums/userrole"
	"github.com/kataras/jwt"
)

type ThetanUserClaims struct {
	jwt.Claims `json:",inline"`
	UserId     string            `json:"user_id"`
	Role       userrole.UserRole `json:"role"`
}

func (u ThetanUserClaims) GetID() string {
	return u.UserId
}

func (u ThetanUserClaims) GetUsername() string {
	return u.Subject
}

func (u ThetanUserClaims) GetRoles() []string {
	return []string{u.Role.String()}
}
