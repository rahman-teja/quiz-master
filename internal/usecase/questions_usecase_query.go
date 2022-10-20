package usecase

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/repository"
)

type QuestionsUsecaseQuery struct {
	repo repository.QuestionsQueryRepository
}

func NewQuestionsUsecaseQuery(prop QuestionsUsecaseProperty) *QuestionsUsecaseQuery {
	return &QuestionsUsecaseQuery{
		repo: prop.Repository,
	}
}

func (q QuestionsUsecaseQuery) GetOne(ctx context.Context, id string) (entity.Questions, interface{}, error) {
	qst, err := q.repo.GetOne(ctx, id)
	if err != nil {
		return entity.Questions{}, nil, err
	}

	return qst, nil, nil
}

func (q QuestionsUsecaseQuery) Get(ctx context.Context) ([]entity.Questions, interface{}, error) {
	qst, err := q.repo.Get(ctx)
	if err != nil {
		return nil, nil, err
	}

	return qst, nil, nil
}
