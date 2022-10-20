package main

import (
	"fmt"
	"time"

	"github.com/rahman-teja/quiz-master/internal/clihandler"
	"github.com/rahman-teja/quiz-master/internal/repository"
	"github.com/rahman-teja/quiz-master/internal/usecase"
	"github.com/rahman-teja/quiz-master/pkg/cli"
	"github.com/rahman-teja/quiz-master/pkg/memorydb"
	"gitlab.com/rteja-library3/rhelper"
)

func main() {
	fmt.Println("Welcome to Quiz Master!")
	fmt.Println("Crafted by : Rahman Teja Wicksono | rahman.9h.23@gmail.com")
	defer fmt.Println("Goodbye Quiz Master!")

	loop := rhelper.GenerateRandom(5, 6)

	fmt.Print("Loading")
	for i := 0; i < loop; i++ {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("")

	fmt.Println("Try `help`")
	fmt.Println("")

	db := memorydb.NewMemoryDB()

	qstProp := usecase.QuestionsUsecaseProperty{
		Repository: repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
			DB: db,
		}),
	}

	ucQstQuery := usecase.NewQuestionsUsecaseQuery(qstProp)
	ucQstCommand := usecase.NewQuestionsUsecaseCommand(qstProp)

	ansProp := usecase.AnswersUsecaseProperty{
		Repository: repository.NewAnswersRepositoryMemory(repository.AnswersRepositoryMemoryProp{
			DB: db,
		}),
		QuestionRepository: repository.NewQuestionsRepositoryMemory(repository.QuestionsRepositoryMemoryProp{
			DB: db,
		}),
	}

	ucAnsQuery := usecase.NewAnswersUsecaseQuery(ansProp)
	ucAnsCommand := usecase.NewAnswersUsecaseCommand(ansProp)

	cli := cli.NewCli()
	cli.Register("create_question", clihandler.QuestionCreateCLIHandler{
		Usecase: ucQstCommand,
	})
	cli.Register("update_question", clihandler.QuestionUpdateCLIHandler{
		Usecase: ucQstCommand,
	})
	cli.Register("delete_question", clihandler.QuestionDeleteCLIHandler{
		Usecase: ucQstCommand,
	})
	cli.Register("question", clihandler.QuestionGetOneCLIHandler{
		Usecase: ucQstQuery,
	})
	cli.Register("questions", clihandler.QuestionGetCLIHandler{
		Usecase: ucQstQuery,
	})
	cli.Register("answer_question", clihandler.AnswersCreateCLIHandler{
		Usecase: ucAnsCommand,
	})
	cli.Register("answers", clihandler.AnswersGetCLIHandler{
		Usecase: ucAnsQuery,
	})

	cli.Start()

	cli.Handle("answers")
	fmt.Println("Thank you")
}
