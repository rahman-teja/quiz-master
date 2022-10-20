package repository

import (
	"context"

	"github.com/rahman-teja/quiz-master/internal/entity"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"gitlab.com/rteja-library3/rapperror"
)

type QuestionsRepositoryMemory struct {
	collection memorydb.Collection
}

type QuestionsRepositoryMemoryProp struct {
	DB memorydb.Database
}

func NewQuestionsRepositoryMemory(prop QuestionsRepositoryMemoryProp) *QuestionsRepositoryMemory {
	return &QuestionsRepositoryMemory{
		collection: prop.DB.Collection("questions"),
	}
}

func (r *QuestionsRepositoryMemory) convertError(err error, id string) error {
	if err == memorydb.ErrDuplicate {
		return rapperror.ErrConflict(
			"",
			"Question no "+id+" already existed!",
			"",
			nil,
		)
	}

	if err == memorydb.ErrNotFound {
		return rapperror.ErrNotFound(
			"",
			"Question no "+id+" is not found",
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

func (r *QuestionsRepositoryMemory) Create(ctx context.Context, questions entity.Questions) error {
	err := r.collection.Create(ctx, questions)
	if err != nil {
		return r.convertError(err, questions.ID)
	}

	return nil
}

func (r *QuestionsRepositoryMemory) Update(ctx context.Context, id string, questions entity.Questions) error {
	err := r.collection.Update(ctx, id, questions)
	if err != nil {
		return r.convertError(err, questions.ID)
	}

	return nil
}

func (r *QuestionsRepositoryMemory) Delete(ctx context.Context, id string) error {
	err := r.collection.Delete(ctx, id)
	if err != nil {
		return r.convertError(err, id)
	}

	return nil
}

func (r *QuestionsRepositoryMemory) Get(ctx context.Context) ([]entity.Questions, error) {
	res := make([]entity.Questions, 0)

	csr, err := r.collection.Find(ctx)
	if err != nil {
		return nil, r.convertError(err, "")
	}
	defer csr.Close(ctx)

	for csr.Next(ctx) {
		qst := entity.Questions{}

		err := csr.Decode(&qst)
		if err != nil {
			return nil, r.convertError(err, "")
		}

		res = append(res, qst)
	}

	if len(res) == 0 {
		return nil, rapperror.ErrNotFound(
			"",
			"Questions is not found",
			"",
			nil,
		)
	}

	return res, nil
}

func (r *QuestionsRepositoryMemory) GetOne(ctx context.Context, id string) (entity.Questions, error) {
	res := entity.Questions{}

	err := r.collection.FindOne(ctx, id).Decode(&res)
	if err != nil {
		return entity.Questions{}, r.convertError(err, id)
	}

	return res, nil
}
