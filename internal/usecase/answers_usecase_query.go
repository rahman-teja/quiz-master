package usecase

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/repository"
)

type AnswersUsecaseQuery struct {
	repo repository.AnswersQueryRepository
}

func NewAnswersUsecaseQuery(prop AnswersUsecaseProperty) *AnswersUsecaseQuery {
	return &AnswersUsecaseQuery{
		repo: prop.Repository,
	}
}

func (q AnswersUsecaseQuery) Get(ctx context.Context) ([]entity.Answers, interface{}, error) {
	ans, err := q.repo.Get(ctx)
	if err != nil {
		return nil, nil, err
	}

	return ans, nil, nil
}
