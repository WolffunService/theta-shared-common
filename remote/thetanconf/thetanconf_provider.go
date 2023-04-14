package thetanconf

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/buger/jsonparser"
	"github.com/cockroachdb/errors"
	"github.com/spf13/viper"

	"github.com/WolffunService/theta-shared-common/rcfg"
	"github.com/WolffunService/theta-shared-common/remote"
	"github.com/WolffunService/theta-shared-common/thetalog"
)

const (
	ProviderName        = "thetan-conf"
	defaultTickInterval = 10 * time.Second
)

func init() {
	remote.RegisterConfigProvider(ProviderName, NewConfigProvider())
}

// ConfigProvider implements reads configuration from Hashicorp Vault.
type ConfigProvider struct {
}

// NewConfigProvider returns a new ConfigProvider.
func NewConfigProvider() *ConfigProvider {
	return &ConfigProvider{}
}

func (p ConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	env := rcfg.Environment(rp.Endpoint())
	raw, err := rcfg.GetConfig(env, rp.Path())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get config from remote server")
	}

	success, err := jsonparser.GetBoolean(raw, "success")
	if err != nil && err != jsonparser.KeyPathNotFoundError {
		return nil, errors.Wrap(err, "error parse config")
	}

	if err == nil && !success {
		code, _ := jsonparser.GetInt(raw, "code")
		msg, _ := jsonparser.GetString(raw, "message")
		return nil, fmt.Errorf("fail to get remote config - code: %d - msg: \"%s\"", code, msg)
	}

	return bytes.NewReader(raw), nil
}

func (p ConfigProvider) Watch(rp viper.RemoteProvider) (io.Reader, error) {
	return p.Get(rp)
}

func (p ConfigProvider) WatchChannel(rp viper.RemoteProvider) (<-chan *viper.RemoteResponse, chan bool) {
	resp := make(chan *viper.RemoteResponse)

	fetchRemoteCfg := func(rp viper.RemoteProvider, resp chan *viper.RemoteResponse) {
		defer func() {
			if err := recover(); err != nil {
				thetalog.NewLogger().Error().
					Op("thetanconf.WatchChannel").
					Var("error", err).Send()
			}
		}()

		data, err := p.Watch(rp)
		result := &viper.RemoteResponse{}
		_, err = data.Read(result.Value)
		result.Error = err
		resp <- result
	}

	go func() {
		for range time.Tick(defaultTickInterval) {
			fetchRemoteCfg(rp, resp)
		}
	}()

	return resp, nil
}
