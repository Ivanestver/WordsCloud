package questionnaire

import (
	"html/template"
	"net/http"
)

type ClientData struct {
	Title string
}

type ClientAnswer struct {
	Text string
	Freq int
}

type OwnerData struct {
	Title   string
	Answers []ClientAnswer
}

func HandleQuestion(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("questionnaire/question.html"))
	data := ClientData{Title: GetQuestionTarget()}
	tmpl.Execute(w, data)
}

func HandleQuestionAdmin(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("questionnaire/question_admin.html"))
	options := GetOptions()
	data := OwnerData{Title: GetQuestionTarget(), Answers: make([]ClientAnswer, 0)}
	for option, freq := range options {
		data.Answers = append(data.Answers, ClientAnswer{Text: option, Freq: freq})
	}
	tmpl.Execute(w, data)
}

func HandleNewOption(w http.ResponseWriter, req *http.Request) {
	option := req.FormValue("answer")
	AddOption(option)
	http.Redirect(w, req, "/thank_page", http.StatusSeeOther)
}

func HandleThankPage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "questionnaire/thank_page.html")
}
