package questionnaire

import (
	"html/template"
	"net/http"
)

type Data struct {
	Name string
}

func HandleQuestion(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("questionnaire/question.html"))
	data := Data{Name: GetQuestionTarget()}
	tmpl.Execute(w, data)
}

func HandleQuestionAdmin(w http.ResponseWriter, req *http.Request) {

}
