package questionnaire

type Question struct {
	Title   string
	Options map[string]int
}

var currentQuestion Question = Question{
	Title:   "",
	Options: make(map[string]int),
}

func MakeNewQuestion(Title string) {
	currentQuestion.Title = Title
	currentQuestion.Options = make(map[string]int)
}

func GetQuestionTarget() string {
	return currentQuestion.Title
}

func GetOptions() map[string]int {
	return currentQuestion.Options
}

func AddOption(option string) {
	currentQuestion.Options[option] += 1
}
