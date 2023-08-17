package model

type Config struct {
	ProviderAURL string `mapstructure:"provider_a_url"`
	ProviderBURL string `mapstructure:"provider_b_url"`
}
