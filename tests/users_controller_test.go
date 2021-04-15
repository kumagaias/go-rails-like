package tests

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"opb/models"
	"testing"
	"unsafe"
)

func TestIndex(t *testing.T) {
	var response []models.User
	w := request("GET", "/users", nil)
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &response)
	assert.Equal(t, http.StatusOK, w.Code)
	// assert.Greater(t, unsafe.Sizeof(response), 0)
}

func TestCreate(t *testing.T) {
	var beforeCount, afterCount int64
	user := models.User{Email: "test1@example.com"}

	DB.Table("users").Where(user).Count(&beforeCount)
	w := request("POST", "/users", user)
	DB.Table("users").Where(user).Count(&afterCount)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, beforeCount+1, afterCount)
}

func TestUpdate(t *testing.T) {
	var beforeCount, afterCount int64
	user := models.User{Email: "test2@example.com"}

	DB.Table("users").Where(user).Count(&beforeCount)
	w := request("PUT", "/users/1", user)
	DB.Table("users").Where(user).Count(&afterCount)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, beforeCount+1, afterCount)
}

func TestDelete(t *testing.T) {
	var beforeCount, afterCount int64

	DB.Model(&models.User{}).Where("id = ?", 1).Count(&beforeCount)
	w := request("DELETE", "/users/1", nil)
	DB.Model(&models.User{}).Where("id = ?", 1).Count(&afterCount)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Equal(t, beforeCount-1, afterCount)
}
