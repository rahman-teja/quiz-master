package repository

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
)

type QuestionsCommandRepository interface {
	Create(ctx context.Context, questions entity.Questions) error
	Update(ctx context.Context, id string, questions entity.Questions) error
	Delete(ctx context.Context, id string) error
}

type QuestionsQueryRepository interface {
	Get(ctx context.Context) ([]entity.Questions, error)
	GetOne(ctx context.Context, id string) (entity.Questions, error)
}

type QuestionsRepository interface {
	QuestionsCommandRepository
	QuestionsQueryRepository
}
