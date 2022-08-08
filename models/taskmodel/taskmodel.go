package taskmodel

import (
	"database/sql"
	
	"github.com/blowmind99/go-crud/config"
	"github.com/blowmind99/go-crud/entities"
)

// membuat struct TaskModel
type TaskModel struct {
	db *sql.DB
}

// mengembalikan struct Task dalam bentuk pointer
func New() *TaskModel {
	// melakukan koneksi atau memanggil koneksi database
	db, err := config.DBConnector()
	if err != nil {
		panic(err)
	}
	return &TaskModel{db:db}
}


// struct method untuk menampilkan semua data Task
func (t *TaskModel) FindAll(task *[]entities.Task) error{
	// membuat query
	rows, err := t.db.Query("select * from task")
	if err != nil {
		return err
	}
	// close koneksi ke database
	defer rows.Close()

	// iterasi rows yang di dapat dari tabel task
	for rows.Next(){
		// membuat variable untuk menampung sebelum di append ke slash task
		var data entities.Task
		rows.Scan(
			&data.Id, 
			&data.Task, 
			&data.Assignee, 
			&data.Deadline)

		// append data entities.Task
		*task = append(*task, data)
	}
	return nil
}

// func string untuk menyimpan data
func (m *TaskModel) Create(task *entities.Task) error {
	result, err := m.db.Exec("insert into task (task, assignee, deadline) values (?,?,?)",
	task.Task, task.Assignee, task.Deadline)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	task.Id = lastInsertId
	return nil

}
	
	
	
	





