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



func HandleGet (c echo.Context) error {
      id, er := strconv.Atoi(c.QueryParam("id"))
	  if er != nil {
		return c.String(http.StatusBadRequest, c.QueryParam("id") + " is not number")
	  }
	  t, er := ts.GetById(id)
      if er != nil {
          return c.String(http.StatusBadRequest,er.Error()) 
	  }
	  return c.JSON(http.StatusOK,t)
}
func HandlePost(c echo.Context) error {
	task := &repository.Task{}
	err := json.NewDecoder(c.Request().Body).Decode(task)
	
	if(err != nil){
		c.String(http.StatusBadRequest,err.Error())
	}
	ts.Create(task.Text,task.Tags)
	return c.JSON(http.StatusAccepted,map[string]string{
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

func HandleDelete(c echo.Context) error {
	id, er := strconv.Atoi(c.QueryParam("id"))
	  if er != nil {
		return c.String(http.StatusBadRequest, c.QueryParam("id") + " is not number")
	  }
	er = ts.DeleteById(id)
	if(er != nil) {
		return c.String(http.StatusBadRequest,er.Error())
	}
	return c.String(http.StatusAccepted, "task deleted")
}