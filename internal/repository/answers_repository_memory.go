package repository

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"gitlab.com/rteja-library3/rapperror"
	"gitlab.com/rteja-library3/rhelper"
)

type AnswersRepositoryMemory struct {
	collection memorydb.Collection
}

type AnswersRepositoryMemoryProp struct {
	DB memorydb.Database
}

func NewAnswersRepositoryMemory(prop AnswersRepositoryMemoryProp) *AnswersRepositoryMemory {
	return &AnswersRepositoryMemory{
		collection: prop.DB.Collection("answers"),
	}
}

func (r *AnswersRepositoryMemory) convertError(err error) *rapperror.AppError {
	if err == memorydb.ErrDuplicate {
		return rapperror.ErrConflict(
			"",
			"Answer already created",
			"",
			nil,
		)
	}

	if err == memorydb.ErrNotFound {
		return rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		)
	}

	return rapperror.ErrInternalServerError(
		"",
		"something went wrong on question",
		"",
		nil,
	)
}

func (r *AnswersRepositoryMemory) Create(ctx context.Context, answers entity.Answers) error {
	if answers.ID == "" {
		answers.ID = rhelper.GenerateID()
	}

	err := r.collection.Create(ctx, answers)
	if err != nil {
		return r.convertError(err)
	}

	return nil
}

func (r *AnswersRepositoryMemory) Get(ctx context.Context) ([]entity.Answers, error) {
	res := make([]entity.Answers, 0)

	csr, err := r.collection.Find(ctx)
	if err != nil {
		return nil, r.convertError(err)
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		qst := entity.Answers{}

		err := csr.Decode(&qst)
		if err != nil {
			return nil, r.convertError(err)
		}

		res = append(res, qst)
	}

	if len(res) == 0 {
		return nil, rapperror.ErrNotFound(
			"",
			"Answer not found",
			"",
			nil,
		)
	}

	return res, nil
}
