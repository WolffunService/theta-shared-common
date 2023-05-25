package dynamiclinkservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/api/firebasedynamiclinks/v1"
	"google.golang.org/api/option"
)

var cacheSrv *dynamicLinkService

type dynamicLinkService struct {
	//logger  thetalog.Logger
	service *firebasedynamiclinks.Service
}

func Init(apiKeyFireBase string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if cacheSrv != nil {
		return errors.New("already initialized")
	}

	if apiKeyFireBase == "" {
		return errors.New("API key FireBase must not empty")
	}

	//logger := thetalog.NewBizLogger("DynamicLinkService")
	service, err := firebasedynamiclinks.NewService(ctx, option.WithAPIKey(apiKeyFireBase))
	if err != nil {
		return err
	}

	cacheSrv = &dynamicLinkService{
		//logger:  logger,
		service: service,
	}

	return nil
}

func GenShortLink(ctx context.Context, dynamicLinkInfo firebasedynamiclinks.DynamicLinkInfo) (string, error) {
	//validate service
	if err := validate(cacheSrv); err != nil {
		return "", err
	}

	req := &firebasedynamiclinks.CreateShortDynamicLinkRequest{
		DynamicLinkInfo: &dynamicLinkInfo,
		Suffix: &firebasedynamiclinks.Suffix{
			Option: "SHORT",
		},
	}

	caller := cacheSrv.service.ShortLinks.Create(req)
	res, err := caller.Context(ctx).Do()
	if err != nil {
		return "", fmt.Errorf("Failed to gen short link : %w ", err)
	}

	return res.ShortLink, nil
}

func validate(cacheSrv *dynamicLinkService) error {
	if cacheSrv != nil {
		return nil
	}
	return fmt.Errorf("You MUST call dynamiclinkservice.Init() prior to using it ")
}
