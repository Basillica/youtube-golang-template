package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"go-template/api"
	"go-template/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
	)
}

type ComponentTest struct {
	router *gin.Engine
	t      *testing.T
	w      *httptest.ResponseRecorder
}

func TestComponent(t *testing.T) {
	router := api.New("v1").Engine
	w := httptest.NewRecorder()
	newC := ComponentTest{
		router: router,
		t:      t,
		w:      w,
	}
	newC.TestPingRoute()
	newC.TestInitMigratesDB()
	// test cases
}

func (c *ComponentTest) TestCreateComponentRoute() {

	w := httptest.NewRecorder()
	postBody, _ := json.Marshal(map[string]any{
		"component_name":      "another component_name",
		"is_parent_component": true,
		"machine_name":        "machine_name",
		"machine_id":          "592b44cd-62f5-4be1-900e-ce6c3fe5c876",
		"plant_name":          "some plant_name",
		"plant_id":            "592b44cd-62f5-4be1-900e-ce6c3fe5c876",
		"organisation_path":   "/",
		"organisation_id":     "organisation_id",
		"organisation_name":   "organisation_name",
		"scope":               "/jenoptik/europe",
	})

	responseBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "/v1/components/create", responseBody)
	c.router.ServeHTTP(w, req)

	assert.Equal(c.t, 200, w.Code)
	assert.Equal(c.t, "pong", w.Body.String())
}

func (c *ComponentTest) TestPingRoute() {

	req, _ := http.NewRequest("GET", "/v1/misc/health", nil)
	c.router.ServeHTTP(c.w, req)

	assert.Equal(c.t, 200, c.w.Code)
	assert.Equal(c.t, "{\"message\":\"pong\"}", c.w.Body.String())
}

func (c *ComponentTest) TestInitMigratesDB() {
	db, mock, sqlDB := NewDatabase()
	defer sqlDB.Close()

	mock.ExpectExec("CREATE TABLE dbo.users(.*)")
	mock.ExpectCommit()

	// tests here
	Init(db)
	//
	req := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "doe@example.com",
	}

	mm := models.NewUserManager(db)
	if ok := mm.Create(*req); !ok {
		fmt.Println("create error = ; want nil")
	}
	// if got != nil {
	// 	c.t.Errorf("create error = %d; want nil", got)
	// }
	// req.Get(db, *user)

	postBody, _ := json.Marshal(map[string]any{
		"component_name":      "another component_name",
		"is_parent_component": true,
		"machine_name":        "machine_name",
		"machine_id":          "592b44cd-62f5-4be1-900e-ce6c3fe5c876",
		"plant_name":          "some plant_name",
		"plant_id":            "592b44cd-62f5-4be1-900e-ce6c3fe5c876",
		"organisation_path":   "/",
		"organisation_id":     "organisation_id",
		"organisation_name":   "organisation_name",
		"scope":               "/jenoptik/europe",
	})
	responseBody := bytes.NewBuffer(postBody)
	reqq, _ := http.NewRequest("POST", "/v1/users/create", responseBody)
	reqq.Header.Set("Authorization", "Bearer token")
	c.router.ServeHTTP(c.w, reqq)
	fmt.Println(reqq)

	assert.Equal(c.t, 201, c.w.Code)
	assert.Equal(c.t, "{\"message\":\"pong\"}", c.w.Body.String())

	mock.ExpectCommit()
	mock.ExpectExec("CREATE TABLE dbo.users(.*)")
	mock.ExpectExec("SELECT * from dbo.users(.*)")
	mock.ExpectExec("INSERT INTO FROM dbo.users(.*)")
}
