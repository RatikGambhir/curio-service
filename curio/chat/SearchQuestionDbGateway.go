package chat

import (
	AppConfig "curio/app"
)

type SearchQuestionDbGateway struct {
	dbConfig *AppConfig.PostgresConfig
}

func NewSearchQuestionDbGateway(config *AppConfig.AppConfig) *SearchQuestionDbGateway {
	return &SearchQuestionDbGateway{
		dbConfig: config.PostgresConfig,
	}

}
