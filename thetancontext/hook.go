package thetancontext

import "context"

// Gắn ở mấy hàm tìm, trong trường hợp data context không có thì tự nó đi tìm bằng hàm này
var (
	//HookGetUser = userService.FindByID
	HookGetUser func(ctx context.Context, userID string) (UserData, error) //

	//HookGetUserRanking = userRankingService.GetUserRanking
	HookGetUserRanking func(ctx context.Context, userID string) (UserRanking, error)
)

// - Gắn mấy hàm này ở mấy service viết hàm Find, nếu context dùng cho hàm đó là thetancontext, thì tự set data lúc đó luôn
//
// - vd:
//
//   - func (service userService) FindByID(ctx context.Context, userId string) (*datamodels.Users, error) {
//     // Find user
//     user, err := service.FindById(userId)
//     if err != nil {
//     return nil, err
//     }
//
//     // - Tui ở đây nè
//     thetancontext.TrySetUserData(ctx, user)
//
//     return user, nil
//     }
func TrySetUserData(ctx context.Context, userData UserData) {
	//type assert context.Context to thetancontext.Context
	if tctx, ok := ctx.(Context); ok {
		tctx.SetUserData(userData)
	}
}

// same TrySetUserData
func TrySetUserUserRanking(ctx context.Context, userRanking UserRanking) {
	//type assert context.Context to thetancontext.Context
	if tctx, ok := ctx.(Context); ok {
		tctx.SetUserRanking(userRanking)
	}
}
