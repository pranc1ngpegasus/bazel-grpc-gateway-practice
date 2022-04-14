package configuration

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Http struct {
			ServerPort string `required:"true" envconfig:"HTTP_SERVER_PORT" default:"3000"`
		}

		Grpc struct {
			ServerPort                     string        `required:"true" envconfig:"GRPC_SERVER_PORT" default:"5020"`
			HealthCheckInterval            time.Duration `envconfig:"GRPC_HEALTHCHECK_INTERVAL" default:"5s"`
			EnforceMentPolicyMinTime       time.Duration `envconfig:"GRPC_ENFORCEMENT_POLICY_MINTIME" default:"10s"`
			EnforceMentPermitWithoutStream bool          `envconfig:"GRPC_ENFORCEMENT_POLICY_PERMIT_WITHOUT_STREAM" default:"true"`
			MaxConnectionIdle              time.Duration `envconfig:"GRPC_SERVER_PARAM_MAX_CONNECTION_IDLE" default:"0"`
			MaxConnectionAge               time.Duration `envconfig:"GRPC_SERVER_PARAM_MAX_CONNECTION_AGE" default:"10m"`
			MaxConnectionAgeGrace          time.Duration `envconfig:"GRPC_SERVER_PARAM_MAX_CONNECTION_AGE_GRACE" default:"0"`
			Time                           time.Duration `envconfig:"GRPC_SERVER_PARAM_TIME" default:"20s"`
			Timeout                        time.Duration `envconfig:"GRPC_SERVER_PARAM_TIMEOUT" default:"6s"`
		}

		Database struct {
			User     string `envconfig:"DATABASE_USER" default:"lqsym"`
			Password string `envconfig:"DATABASE_PASSWORD" default:"lqsym"`
			Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
			Port     int    `envconfig:"DATABASE_PORT" default:"3306"`
			Name     string `envconfig:"DATABASE_NAME" default:"bazel-grpc-gateway-practice"`
		}
	}
)

var globalConfig Config

func Load() {
	envconfig.MustProcess("", &globalConfig)
}

func Get() Config {
	return globalConfig
}
