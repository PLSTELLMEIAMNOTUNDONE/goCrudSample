package taskservice

import (
	"errors"
	"fmt"
	"io"
	repository "maintmp/internal/repos"
	"strings"
	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/charmap"
)


type TaskService struct {
	nextId int
	
	rep *repository.Repository
	
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
func (ts *TaskService) GetNextId() int {
      ts.nextId++
	  return ts.nextId
}
func (ts *TaskService) Create(text string,tags repository.Tags) int {
	task := repository.Task {
		Id: ts.GetNextId(),
		Text: text,
		Tags: tags,
	}


	ts.rep.InsertTask(task.Id,task.Text, task.Tags)
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
	ts.rep.DeleteTask(id)
	var er error
	er = nil
	defer func()  {
		if  recover() != nil {
			 er = errors.New("sql error")
		}
	}()
	return er
}

func (ts *TaskService) GetAll() ([]repository.Task,error) {
	
	t, e := ts.rep.GetAllTasks()
	if(e != nil){
		return nil,e
	}
	return t, nil
}
func (ts *TaskService) UpdateTaskText(id int, text string) error {
    ts.rep.UpdateTaskText(id, text)
	var er error
	er = nil
	defer func()  {
		if  recover() != nil {
			 er = errors.New("sql error")
		}
	}()
	return er
}
func (ts *TaskService) UpdateTaskTags(id int, tags repository.Tags) error {
    ts.rep.UpdateTaskTags(id, tags)
	var er error
	er = nil
	defer func()  {
		if  recover() != nil {
			 er = errors.New("sql error")
		}
	}()
	return er
	
}
func (ts *TaskService) UpdateTask (id int, text string,  tags repository.Tags) (error, error) {
	 var er1 error 
	 var er2 error 
	 er1, er2 = nil, nil
	 if text != "" {
		er1 = ts.UpdateTaskText(id, text)
	 }
	 if tags != nil {
		er2 = ts.UpdateTaskTags(id, tags)
	 }
	 return er1, er2
	 
}