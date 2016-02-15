package areena

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Config - client settings
type Config struct {
	YleAppID   string `envconfig:"APP_ID" required:"true"`
	YleAppKey  string `envconfig:"APP_KEY" required:"true"`
	YleSecret  string `envconfig:"SECRET" required:"true"`
	APIBaseURL string `envconfig:"BASE_URL" default:"https://external.api.yle.fi/v1/"`
	Debug      bool   `envconfig:"DEBUG" default:"false"`
	HTTPClient *http.Client
}

// NewConfig init config based on env
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}
	err = envconfig.Process("YLE", cfg)
	return
}
