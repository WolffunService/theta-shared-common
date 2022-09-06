package rcfg

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"strings"
)

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
