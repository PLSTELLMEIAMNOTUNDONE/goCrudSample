package handlers

import (
	"encoding/json"

	repository "maintmp/internal/repos"
	"maintmp/internal/taskService"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
var ts = taskservice.New()


// @Summary get task with id
// @Description get task with certain id
// @Produce json
// @ID getTask
// @Param id query integer true "id"
// @Success 200 {object} error
// @Failure 400 {string} error
// @Router /task [get]
func HandleGet (c echo.Context) error {
      id, er := strconv.Atoi(c.QueryParam("id"))
	  if er != nil {
		return c.String(http.StatusBadRequest, c.QueryParam("id") + " is not number")
	  }
	  
	  t, er := ts.GetById(id)
      if er != nil {
          return c.String(http.StatusBadRequest, er.Error()) 
	  }
	  return c.JSON(http.StatusOK, t)
}
type TaskWithNoId struct {
	Text string    `json:"text" db:"text"`
	Tags repository.Tags  `json:"tags" db:"tags"`
}

// @Summary making task
// @Description make task
// @Accept json 
// @Produce json
// @ID posttask
// @Param task body TaskWithNoId true "task"
// @Success 200 {object} error
// @Failure 400 {string} error
// @Router /task [post]
func HandlePost(c echo.Context) error {
	task := &repository.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(task)
	
	if(err != nil){
		c.String(http.StatusBadRequest, err.Error())
	}
	ts.Create(task.Text, task.Tags)
	return c.JSON(http.StatusAccepted, map[string]string{
		"message": "okk"})
}


// @Summary GetAll
// @Description shoing all tasks
// @Produce json
// @ID getall
// @Success 200 {object} error
// @Failure 400 {string} error
// @Router /task/all [get]
func HandleGetAll(c echo.Context) error {
	t, e := ts.GetAll()
	if(e != nil) {
		c.String(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK,t)
}
// @Summary delete task with id
// @Description delete task with certain id
// @Produce json
// @ID deleteTask
// @Param id query integer true "id"
// @Success 200 {string} error
// @Failure 400 {string} error
// @Router /task [delete]
func HandleDelete(c echo.Context) error {
	id, er := strconv.Atoi(c.QueryParam("id"))
	if er != nil {
		return c.String(http.StatusBadRequest, c.QueryParam("id") + " is not number")
	}
	er = ts.DeleteById(id)
	if er != nil {
		return c.String(http.StatusBadRequest, er.Error())
	}
	return c.String(http.StatusAccepted, "task deleted")
}

// @Summary updating task
// @Description update task
// @Accept json 
// @Produce json
// @ID updateTask
// @Param text body repository.Task true "upd"
// @Success 200 {object} error
// @Failure 400 {string} error
// @Router /task [put]
func HandleUpdate (c echo.Context) error {
	task := &repository.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(task)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
    if task.Id == 0 {
		return c.String(http.StatusBadRequest, "no id")
	}
	err, err1 := ts.UpdateTask(task.Id, task.Text, task.Tags)
	if err != nil  {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if err1 != nil {
		return c.String(http.StatusBadRequest, err1.Error())
	}
	return c.String(http.StatusAccepted, "task changed")
} 