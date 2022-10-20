package usecase

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository"
)

type AnswersUsecaseProperty struct {
	Repository         repository.AnswersRepository
	QuestionRepository repository.QuestionsRepository
}

type AnswersCommandUsecase interface {
	Create(ctx context.Context, payload model.Answers) (entity.Answers, interface{}, error)
}

type AnswersQueryUsecase interface {
	Get(ctx context.Context) ([]entity.Answers, interface{}, error)
}
