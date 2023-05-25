package serverenum

import "strings"

const (
	EnvDevelopment Environment = "development"
	EnvStaging     Environment = "staging"
	EnvUAT         Environment = "uat"
	EnvProduction  Environment = "production"
)

type Environment string

func (env Environment) String() string {
	return string(env)
}

func (env Environment) IsProduction() bool {
	return env == EnvProduction
}

func (env Environment) IsCloudNative() bool {
	return env != EnvDevelopment
}

func (env Environment) IsDev() bool {
	return env == EnvDevelopment
}

func (env Environment) ToRemote() string {
	rewriteEnv := env
	switch env {
	case EnvStaging, EnvUAT, EnvProduction:
	default:
		rewriteEnv = EnvStaging
	}
	return strings.ToUpper(string(rewriteEnv))
}
