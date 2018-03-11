package common

import (
	_ "github.com/lib/pq"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"testing"
	"net/http"
	"github.com/go-xorm/xorm"
)

func GetRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("../views/*")
	}
	return r
}

func RunTestHTTPResponse(
	t *testing.T,
	r *gin.Engine,
	req *http.Request,
	f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

const driverName = "postgres"
const dataSourceName = "postgres://localhost:5432/xorm_test_test?sslmode=disable"

func InitTestDb() *xorm.Engine {
	db, engineErr := xorm.NewEngine(driverName, dataSourceName)

	if engineErr != nil {
		panic("couldn't connect to database: \n" + engineErr.Error())
	} else {
		return db
	}
}
