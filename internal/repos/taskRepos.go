package repository

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Tags []string

func (t Tags) Value() (driver.Value, error) {
	if(len(t) == 0) {
		return driver.Value("[]"), nil
	}
	return driver.Value(fmt.Sprintf(`["%s"]`, strings.Join(t, `", "`))), nil
}

func (t *Tags) Scan(src interface{}) (err error) {
	var srcb []byte
	switch src.(type) {
	case string:
		srcb = []byte(src.(string))
	case []byte:
		srcb = src.([]byte)
	default:
		return errors.New("bad type")
	}

	reader := bytes.NewReader(srcb)
	
	b, err:= io.ReadAll(reader)
	s := string(b[:])
	fmt.Println(s)
    ss := strings.Split(s[1 : len(s) - 1], ",")
	fmt.Println(ss)
	if(err != nil) {
		return
	}
	*t = ss
	return nil
}


type Task struct {
	Id   int       `json:"id" db:"id"`
	Text string    `json:"text" db:"text"`
	Tags Tags  `json:"tags" db:"tags"`
	//Due  time.Time `json:"due"`
}
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
type Repository struct {
	Database *sqlx.DB
}

func (r *Repository ) MakeTable(name string) (sql.Result, error){
	 s := fmt.Sprintf( `create table %s (
		id integer primary key not null,
		text varchar,
		tags varchar[]
   );`,name)
    res, e := r.Database.Exec(s)
	if(e != nil) {
		return nil, e
	}
	return res, nil

}
func (r *Repository) GetAllTasks() ([]Task, error){
      tasks := []Task{}
	  s := "select * from tasks;"
	  e := r.Database.Select(&tasks,s)
	  return tasks, e
}
func (r *Repository) GetById(i int) (Task, error){
	task := Task{}
	s := fmt.Sprintf("select * from tasks where id = %d;", i)
	e := r.Database.Get(&task,s)
	return task, e
}
func (r *Repository) InsertTask(id int, text string, tags Tags) {
    var as strings.Builder
	as.WriteString("{")
	for i, tag := range tags {
		as.WriteString(fmt.Sprintf("%s", tag))
		if(i != len(tags) - 1) {
			as.WriteString(", ")
		}
	}
	as.WriteString("}")
	
	r.Database.MustExec("insert into tasks (id, text, tags)  values ($1, $2, $3);", id, text, as.String())
}


func New(cnf Config) (*Repository, error) {
	bd, er := new(cnf)
	r := &Repository{bd}
	if(er != nil) {
		return nil, er
	}
    return r, nil
}
func new(cnf Config) (*sqlx.DB, error) {
	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cnf.Host,cnf.Port,cnf.Username,cnf.Password,cnf.DBName,cnf.SSLMode);
	fmt.Println(s)
	db, er := sqlx.Open("postgres", s)
	if er != nil {
		return nil, er
	}
	er = db.Ping()
	if er != nil {
		return nil, er
	}
	return db, nil
}