package usecase

import (
	"context"
	"strings"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository"
)

type AnswersUsecaseCommand struct {
	repo               repository.AnswersCommandRepository
	questionRepository repository.QuestionsQueryRepository
}

func NewAnswersUsecaseCommand(prop AnswersUsecaseProperty) *AnswersUsecaseCommand {
	return &AnswersUsecaseCommand{
		repo:               prop.Repository,
		questionRepository: prop.QuestionRepository,
	}
}

func (q AnswersUsecaseCommand) Create(ctx context.Context, payload model.Answers) (entity.Answers, interface{}, error) {
	qst, err := q.questionRepository.GetOne(ctx, payload.QuestionId)
	if err != nil {
		return entity.Answers{}, nil, err
	}

	var isCorrect bool
	var point int64

	ansLower := strings.ToLower(payload.Answer)
	for _, ans := range qst.Answers {
		if ans == ansLower {
			isCorrect = true
			point = qst.Point
		}
	}

	ans := entity.Answers{
		Questions: qst,
		Answer:    payload.Answer,
		IsCorrect: isCorrect,
		Point:     point,
	}

	err = q.repo.Create(ctx, ans)
	if err != nil {
		return entity.Answers{}, nil, err
	}

	return ans, nil, nil
}
