package models

import "time"

type ConfigurationDetail struct {
}

func NewConfigurationDetail() *ConfigurationDetail {
	return &ConfigurationDetail{}
}

type Configuration struct {
	Id            int
	Configuration ConfigurationDetail
	User          User
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewConfiguration(configuration ConfigurationDetail, user User) *Configuration {
	return &Configuration{
		User:          user,
		Configuration: configuration,
		CreatedAt:     time.Now(),
	}
}
