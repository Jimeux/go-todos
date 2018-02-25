package article

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"gin-test/common"
	"strings"
)

// TODO: Desperately need a mocking library
type MockArticleRepository struct{}

func (m *MockArticleRepository) Create(title string, content string) (*Model, error) {
	return &Model{1, "Create title", "Create content"}, nil
}
func (m *MockArticleRepository) FindAll() (*[]Model, error) {
	return &[]Model{
		{1, "FindAll title", "FindAll content"},
	}, nil
}
func (m *MockArticleRepository) FindById(id int64) (*Model, bool, error) {
	return &Model{1, "FindById title", "FindById content"}, true, nil
}

var handler = Handler{&MockArticleRepository{}}

func TestIndex(t *testing.T) {
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
