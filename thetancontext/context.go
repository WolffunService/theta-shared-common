package thetancontext

import (
	"context"
	"time"
)

type (
	Context interface {
		context.Context
		IUserData

		WithValue(key, val interface{})
	}

	userContext struct {
		context.Context
		*baseUserData
	}

	//context key
	userDataKey struct {
	}
)

// NewContext creates a new context that combines the provided context with user data, represented by a userID.
//   - The returned context implements both the context.Context interface and the IUserData interface.
//   - If the provided context is nil, a new context with a timeout of 10 seconds will be created using context.Background().
//   - The user data is stored in a struct that implements the IUserData interface. The userContext struct embeds the provided context and this user data struct.
//   - The function returns a Context interface that contains the new context and user data.
//   - This code is useful when you need to pass around a context that contains additional user data.
func NewContext(ctx context.Context, userID string) Context {
	if ctx == nil {
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second) //?huh
	}
	userData := newBase(userID)
	return &userContext{
		Context: context.WithValue(ctx, userDataKey{}, userData),
		//user data
		baseUserData: userData,
	}
}

func (u *userContext) WithValue(key, val interface{}) {
	u.Context = context.WithValue(u.Context, key, val)
}
