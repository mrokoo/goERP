package app_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/internal/share/account/domain"
	"github.com/mrokoo/goERP/routes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupRouter() *gin.Engine {
	connectionString := "mongodb://localhost:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	routes.AccountRoutes(r, client)
	return r
}

func TestHandler_GetAccountList(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/account/getAccountList", nil)
	router.ServeHTTP(w, req)
	assert := assert.New(t)
	fmt.Println(w.Result())
	assert.Equal(200, w.Code)
}

func TestHandler_AddAccountAndDeleteAccount(t *testing.T) {
	router := setupRouter()
	rec := httptest.NewRecorder()
	accountReq := domain.Account{
		ID:      "A001",
		Name:    "账号1",
		Type:    3,
		Holder:  "张三",
		Number:  "402901000226",
		Note:    "测试实例",
		State:   2,
		Balance: 2000,
	}
	account, _ := json.Marshal(&accountReq)
	reader := bytes.NewReader(account)
	fmt.Printf("%#v", rec.Result())
	req, _ := http.NewRequest("POST", "/account/addAccount", reader)
	req.Header.Set("Content-Type", "application/json")
	defer t.Cleanup(func() {
		idreq := struct {
			ID string `json:"id"`
		}{
			ID: "A001",
		}
		idjson, _ := json.Marshal(&idreq)
		reader := bytes.NewReader(idjson)
		req, _ := http.NewRequest("Delete", "/account/deleteAccount", reader)
		req.Header.Set("Content-Type", "application/json")
	})
	router.ServeHTTP(rec, req)
	assert := assert.New(t)

	assert.Equal(200, rec.Code)
}

// func TestHandler_UpdateAccount(t *testing.T) {
// 	router := setupRouter()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/account/updateAccount", nil)
// 	router.ServeHTTP(w, req)
// 	assert := assert.New(t)
// 	fmt.Println(w.Result())
// 	assert.Equal(200, w.Code)
// }
