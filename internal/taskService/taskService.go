package taskservice

import (
	"fmt"
	"io"
	repository "maintmp/internal/repos"
	"strings"
	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/charmap"
	
)


// -d  TT

type TaskService struct {
	
	rep *repository.Repository
	nextId int
}

func New() *TaskService {
	ts := &TaskService{}
	db, er := repository.New(repository.Config{
		Host : "localhost",
		Port: "5432",
		Username: "postgres",
		Password: "postgres",
		DBName: "postgres",
		SSLMode: "disable",
	})
    if er != nil {
		a, _ := io.ReadAll(charmap.Windows1251.NewDecoder().Reader(strings.NewReader(er.Error())))
		fmt.Print(string(a))
	}
	ts.rep = db
    ts.nextId = 0
	return ts
}


func (ts *TaskService) Create(text string,tags repository.Tags) int {
	task := repository.Task {
		Id : ts.nextId,
		Text: text,
		Tags: tags,
	}


	ts.rep.InsertTask(task.Id,task.Text, task.Tags)
	ts.nextId++
	return task.Id
}

func (ts *TaskService) GetById(id int) (repository.Task, error) {
	t, er  := ts.rep.GetById(id)
	if er == nil {
		return t, nil
	} 
	return repository.Task{}, er
}

func (ts *TaskService) DeleteById(id int) error {
	
	
	return fmt.Errorf("no Task with id = %d", id)
}

func (ts *TaskService) GetAll() ([]repository.Task,error) {
	
	t, e := ts.rep.GetAllTasks()
	if(e != nil){
		return nil,e
	}
	return t, nil
}