package usecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/divan/num2words"
	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/internal/model"
	"github.com/rahman-teja/quiz-master/internal/repository"
	"github.com/rahman-teja/quiz-master/pkg/words"
)

type QuestionsUsecaseCommand struct {
	repo repository.QuestionsCommandRepository
}

func NewQuestionsUsecaseCommand(prop QuestionsUsecaseProperty) *QuestionsUsecaseCommand {
	return &QuestionsUsecaseCommand{
		repo: prop.Repository,
	}
}

func (q QuestionsUsecaseCommand) buildAnswers(ans string) []string {
	nAns, err := strconv.Atoi(ans)
	if err == nil { // is number
		cvt := num2words.Convert(nAns)
		return []string{ans, cvt, strings.ToLower(cvt)}
	}

	iWordAns, err := words.ToNumber(strings.ToLower(ans))
	if err == nil { // is word number
		return []string{ans, strconv.Itoa(iWordAns), strings.ToLower(ans)}
	}

	return []string{ans}
}

func (q QuestionsUsecaseCommand) Create(ctx context.Context, payload model.Questions) (entity.Questions, interface{}, error) {
	qst := entity.Questions{
		ID:        payload.ID,
		Questions: payload.Questions,
		Answers:   q.buildAnswers(payload.Answer),
		Point:     1,
	}

	err := q.repo.Create(ctx, qst)
	if err != nil {
		return entity.Questions{}, nil, err
	}

	return qst, nil, nil
}

func (q QuestionsUsecaseCommand) Update(ctx context.Context, id string, payload model.Questions) (entity.Questions, interface{}, error) {
	qst := entity.Questions{
		ID:        id,
		Questions: payload.Questions,
		Answers:   q.buildAnswers(payload.Answer),
		Point:     1,
	}

	err := q.repo.Update(ctx, id, qst)
	if err != nil {
		return entity.Questions{}, nil, err
	}

	return qst, nil, nil
}

func (q QuestionsUsecaseCommand) Delete(ctx context.Context, id string) error {
	err := q.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
