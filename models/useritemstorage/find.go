package useritemstorage

import (
	"context"
	"fmt"
	"github.com/WolffunService/thetan-shared-common/database/mongodb"
	"github.com/WolffunService/thetan-shared-common/database/mongodb/utils"
	"github.com/WolffunService/thetan-shared-common/models/useritemmodel"
)

func FindUserItemsById(ctx context.Context, userId string) (*useritemmodel.UserItems, error) {
	userItems := &useritemmodel.UserItems{}
	userObjectId := utils.ObjectIDFromHex(userId)
	col := mongodb.CollRead(userItems)
	err := col.FindByIDWithCtx(ctx, userObjectId, userItems)
	if err != nil {
		userItems.ID = userObjectId
		return userItems, err
	}
	fmt.Printf("Found a single document: %+v\n", userItems)
	return userItems, nil
}
