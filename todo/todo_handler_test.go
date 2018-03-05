package todo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"gin-test/common"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
)

type MockArticleRepository struct{ mock.Mock }
func (m *MockArticleRepository) Create(title string, content string) (*Model, error) {
	args := m.Called(title, content)
	return args.Get(0).(*Model), args.Error(1)
}
func (m *MockArticleRepository) FindAll() (*[]Model, error) {
	args := m.Called()
	return args.Get(0).(*[]Model), args.Error(1)
}
func (m *MockArticleRepository) FindById(id int64) (*Model, bool, error) {
	args := m.Called(id)
	return args.Get(0).(*Model), args.Bool(1), args.Error(2)
}

func TestIndex(t *testing.T) {
	repo := new(MockArticleRepository)
	repo.On("FindAll").Return(
		&[]Model{
			{1, "FindAll title", "FindAll content"},
		},
		nil,
	)
	var handler = Handler{repo}

	r := common.GetRouter(true)
	r.GET("/", handler.Index)
	req, _ := http.NewRequest("GET", "/", nil)

	common.RunTestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Articles</title>") > 0
		return statusOK && pageOK
	})
}

func TestShow(t *testing.T) {
	model := Model{1, "FindById title", "FindById content"}
	repo := new(MockArticleRepository)
	repo.On("FindById", int64(1)).Return(&model, true, nil)
	var handler = Handler{repo}

	r := common.GetRouter(true)
	r.GET("/article/:id", handler.Show)
	req, _ := http.NewRequest("GET", "/article/1", nil)

	common.RunTestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>FindById title</title>") > 0
		return statusOK && pageOK
	})
}

func TestArticleListJSON(t *testing.T) {
	repo := new(MockArticleRepository)
	repo.On("FindAll").Return(
		&[]Model{
			{1, "FindAll title", "FindAll content"},
		},
		nil,
	)
	var handler = Handler{repo}

	r := common.GetRouter(false)
	r.GET("/", handler.Index)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	common.RunTestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []Model
		err = json.Unmarshal(p, &articles)

		return err == nil && len(articles) >= 1 && statusOK
	})
}
