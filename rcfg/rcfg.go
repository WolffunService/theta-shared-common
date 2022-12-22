package rcfg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"
)

type UserContext struct {
	UserID     string
	Attributes map[string]any
}

type Option struct {
	PreventPushEvent bool
	Country          string
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

const remoteCfgBaseUrl = "https://thetan-support.thetanarena.com/api/remote-config"

func GetLatest(env Environment, name string) ([]byte, error) {
	name = strings.ToLower(name)
	url := fmt.Sprintf("%s/view", remoteCfgBaseUrl)
	client := req.C()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("env", env.String()).
		SetQueryParam("name", name).
		SetQueryParam("revision", "0").
		SetQueryParam("raw", "true").
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsSuccess() {
		return resp.ToBytes()
	}

	return nil, ErrUnknownRequest
}

func GetByUser[T any](env Environment, name string, userCtx UserContext, option Option) (*T, error) {
	name = strings.ToLower(name)
	url := fmt.Sprintf("%s/config", remoteCfgBaseUrl)

	attribute, _ := json.Marshal(userCtx.Attributes)
	client := req.C()
	resp, err := client.R().
		SetQueryParam("env", env.String()).
		SetQueryParam("name", name).
		SetQueryParam("userId", userCtx.UserID).
		SetQueryParam("attribute", string(attribute)).
		SetQueryParam("raw", "true").
		SetQueryParam("preventPushEvent", strconv.FormatBool(option.PreventPushEvent)).
		SetQueryParam("country", option.Country).
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
