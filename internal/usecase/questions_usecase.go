package usecase

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository"
)

type QuestionsUsecaseProperty struct {
	Repository repository.QuestionsRepository
}

type QuestionsCommandUsecase interface {
	Create(ctx context.Context, payload model.Questions) (entity.Questions, interface{}, error)
	Update(ctx context.Context, id string, payload model.Questions) (entity.Questions, interface{}, error)
	Delete(ctx context.Context, id string) error
}

type QuestionsQueryUsecase interface {
	GetOne(ctx context.Context, id string) (entity.Questions, interface{}, error)
	Get(ctx context.Context) ([]entity.Questions, interface{}, error)
}
