package chat

import (
	AppConfig "curio/app"
)

type SearchQuestionProcessor struct {
	searchQuestionDbGateway *SearchQuestionDbGateway
}

func NewSearchQuestionProcessor(config *AppConfig.AppConfig) *SearchQuestionProcessor {
	return &SearchQuestionProcessor{
		searchQuestionDbGateway: NewSearchQuestionDbGateway(config),
	}
}
