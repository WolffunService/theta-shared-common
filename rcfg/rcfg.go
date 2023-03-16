package rcfg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/WolffunGame/theta-shared-common/common"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"
)

type UserContext struct {
	UserID     string
	Attributes map[string]any
}

type Option struct {
	DisablePushEvent bool
	Country          string
	EventName        string
	TopicName        string
}

type GetByUserRequest struct {
	User   UserContext
	Option Option
}

var ErrUnknownRequest = errors.New("unknown request error")

type Environment string

const (
	Production Environment = "PRODUCTION"
	UAT        Environment = "UAT"
	Staging    Environment = "STAGING"
)

func (env Environment) String() string {
	return string(env)
}

var (
	remoteCfgBaseInternalUrl = "http://thetan-support.default.svc.cluster.local:1706/api/remote-config"
	remoteCfgBaseUrl         = map[Environment]string{
		Staging:    "https://thetan-support.staging.thetanarena.com/api/remote-config",
		UAT:        "https://thetan-support.uat.thetanarena.com/api/remote-config",
		Production: "https://thetan-support.thetanarena.com/api/remote-config",
	}
)

func GetRemoteCfgBaseUrl(env Environment, route string) (string, error) {
	if common.OnCloud() {
		return remoteCfgBaseInternalUrl, nil
	}

	if url, ok := remoteCfgBaseUrl[env]; ok {
		urlFull := url + route
		return urlFull, nil
	}
	return "", fmt.Errorf("unknown env: %s", env)
}

func GetLatest(env Environment, name string) ([]byte, error) {
	name = strings.ToLower(name)
	url, err := GetRemoteCfgBaseUrl(env, "/view")
	if err != nil {
		return nil, err
	}
	client := req.C()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("name", name).
		SetQueryParam("revision", "0").
		SetQueryParam("raw", "true").
		SetQueryParam("checkFlag", "true").
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess() {
		return resp.ToBytes()
	}

	return nil, ErrUnknownRequest
}

func GetConfig(env Environment, name string) ([]byte, error) {
	name = strings.ToLower(name)
	url, err := GetRemoteCfgBaseUrl(env, "/config")
	if err != nil {
		return nil, err
	}
	client := req.C()
	resp, err := client.R().
		SetQueryParam("name", name).
		SetQueryParam("raw", "true").
		SetQueryParam("viewOnly", "true").
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess() {
		return resp.ToBytes()
	}

	return nil, ErrUnknownRequest
}

func GetByUser[T any](env Environment, name string, request GetByUserRequest) (*T, error) {
	name = strings.ToLower(name)
	url, err := GetRemoteCfgBaseUrl(env, "/config")
	if err != nil {
		return nil, err
	}
	attribute, _ := json.Marshal(request.User.Attributes)
	client := req.C()
	resp, err := client.R().
		SetQueryParam("name", name).
		SetQueryParam("userId", request.User.UserID).
		SetQueryParam("attribute", string(attribute)).
		SetQueryParam("raw", "true").
		SetQueryParam("disablePushEvent", strconv.FormatBool(request.Option.DisablePushEvent)).
		SetQueryParam("country", request.Option.Country).
		SetQueryParam("eventName", request.Option.EventName).
		SetQueryParam("topicName", request.Option.TopicName).
		Get(url)
	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, ErrUnknownRequest
	}

	data, err := resp.ToBytes()
	if err != nil {
		return nil, err
	}

	var res T
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
