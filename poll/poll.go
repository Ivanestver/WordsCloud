package poll

import (
	"net/http"
	"wordscloud/questionnaire"
)

func HandleNewPoll(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "poll/new_poll.html")
}

func HandleCreatePoll(w http.ResponseWriter, req *http.Request) {
	questionTitle := req.FormValue("question")
	questionnaire.MakeNewQuestion(questionTitle)
	http.Redirect(w, req, "/question", http.StatusSeeOther)
}
