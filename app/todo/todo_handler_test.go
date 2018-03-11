package todo

import (
	"testing"
	"gin-todos/common"
	"net/http/httptest"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestIndex(t *testing.T) {
	var handler = Handler{TestRepository{}}
	router := common.GetRouter(true)
	router.GET("/", handler.Index)
	request, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := ioutil.ReadAll(recorder.Body)

	assert.Contains(t, string(body),"<title>Todo</title>")
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestListValidDefault(t *testing.T) {
	todoList := []Model{{1, "A todo 1", false, time.Now()}}
	var handler = Handler{TestRepository{TodoList: &todoList}}
	router := common.GetRouter(false)
	router.GET("/todo", handler.List)
	request, _ := http.NewRequest("GET", "/todo", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	js, _ := ioutil.ReadAll(recorder.Body)
	var todos []Model
	json.Unmarshal(js, &todos)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, len(todoList), len(todos))
}

func TestListValidWithoutComplete(t *testing.T) {
	var handler = Handler{TestRepository{}}
	router := common.GetRouter(false)
	router.GET("/todo", handler.List)
	request, _ := http.NewRequest("GET", "/todo", nil)
	recorder := httptest.NewRecorder()

	q := request.URL.Query()
	q.Add(HideCompleteParam, "true")
	request.URL.RawQuery = q.Encode()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestListParseError(t *testing.T) {
	var handler = Handler{TestRepository{}}
	router := common.GetRouter(false)
	router.GET("/todo", handler.List)
	request, _ := http.NewRequest("GET", "/todo", nil)

	q := request.URL.Query()
	q.Add(HideCompleteParam, "oops")
	request.URL.RawQuery = q.Encode()

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotAcceptable, recorder.Code)
}

func TestListRepoError(t *testing.T) {
	var handler = Handler{TestRepository{Error: errors.New("oops")}}
	router := common.GetRouter(false)
	router.GET("/todo", handler.List)
	request, _ := http.NewRequest("GET", "/todo", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}


type TestRepository struct {
	Todo         *Model
	TodoList     *[]Model
	SetCompleted int64
	Error        error
}

func (t TestRepository) Create(title string) (*Model, error) {
	return t.Todo, t.Error
}
func (t TestRepository) FindAll(hideCompleted bool) (*[]Model, error) {
	return t.TodoList, t.Error
}
func (t TestRepository) SetComplete(id int64, complete bool) (int64, error) {
	return t.SetCompleted, t.Error
}
func (t TestRepository) Delete(id int64) error {
	return t.Error
}
