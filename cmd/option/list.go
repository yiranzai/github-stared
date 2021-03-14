package option

import (
	"fmt"

	"github.com/spf13/viper"
)

// ListCmdConfig is config for list command
type ListCmdConfig struct {
	Username string
	Output   string
	Write    bool
	All      bool
}

// NewListCmdConfigFromViper generate config for list command from viper
func NewListCmdConfigFromViper() (*ListCmdConfig, error) {
	var conf ListCmdConfig
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config from viper: %w", err)
	}
	return &conf, nil
}
