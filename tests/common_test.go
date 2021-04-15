package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"opb/config"
	"opb/lib"
	"opb/routes"
	"os"
	"testing"
)

var router *gin.Engine
var DB *gorm.DB

func TestMain(m *testing.M) {
	beforeTest()
	status := m.Run()
	afterTest(status)
}

func beforeTest() {
	config.Init("test")
	// drop database
	if lib.FileExists(config.Get("db.database")) {
		os.Remove(config.Get("db.database"))
	}
	DB = lib.GetDB()
	lib.InitDB(DB)

	router = routes.Init(DB)
	lib.SeedDB(DB)
}

func afterTest(status int) {
	os.Exit(status)
}

func request(method, path string, body interface{}) *httptest.ResponseRecorder {
	value, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, path, bytes.NewReader(value))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	return w
}
