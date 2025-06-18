package questionnaire

import "strings"

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
	begin := 0
	for option[begin] == ' ' {
		begin++
	}
	end := len(option) - 1
	for option[end] == ' ' {
		end--
	}
	option_ready := option[begin : end+1]
	currentQuestion.Options[strings.ToLower(option_ready)] += 1
}
