package taskcontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/blowmind99/go-crud/entities"
	"github.com/blowmind99/go-crud/models/taskmodel"
)

// memanggil model
var taskModel = taskmodel.New()

func Index(w http.ResponseWriter, _ *http.Request){
	// memanggil fungsi getdata dalam fungsi index
	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	// memanggil view untuk menampilkan data mahasiswa
	temp, _:=template.ParseFiles("views/task/index.html")
	temp.Execute(w, data)
}


// memanggil method FindAll
// tujuannya karena method GetData akan di panggil di Index
// juga akan dipakai saat proses simpan data dan hapus data
func GetData() string{
	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/task/data.html")

	var task []entities.Task
	err := taskModel.FindAll(&task)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"task": task,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()

}

// membuat function getform
func GetForm(w http.ResponseWriter, r *http.Request){
	temp, _:=template.ParseFiles("views/task/form.html")
	temp.Execute(w, nil)
}

// membuat function store
func Store(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost {
		// simpan data ke databse
		r.ParseForm()

		// membuat variable task dengan struct task
		var task entities.Task

		// mengambil inputan yang di terima dari form
		task.Task = r.Form.Get("task")
		task.Assignee = r.Form.Get("assignee")
		task.Deadline = r.Form.Get("deadline")

		err := taskModel.Create(&task)

		if err != nil {
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		// jika tidak gagal
		data := map[string]interface{}{
			"message": "Task Has Been Created!",
			// mengembalikan data dari func GetData
			"data": template.HTML(GetData()),
		}
		ResponseJson(w, http.StatusOK, data)
	}
}

// response error
func ResponseError(w http.ResponseWriter, code int, message string){
	ResponseJson (w, code, map[string]string{"error": message})
}



// response json
func ResponseJson(w http.ResponseWriter, code int, payload interface{}){
	response, _:= json.Marshal(payload)
	// set header content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// membuat function rubah button mark as done