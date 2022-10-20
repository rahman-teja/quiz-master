package repository

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
)

type AnswersCommandRepository interface {
	Create(ctx context.Context, questions entity.Answers) error
}

type AnswersQueryRepository interface {
	Get(ctx context.Context) ([]entity.Answers, error)
}

type AnswersRepository interface {
	AnswersCommandRepository
	AnswersQueryRepository
}
