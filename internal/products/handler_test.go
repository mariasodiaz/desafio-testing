package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	products := []Product{{ID: "1", SellerID: "123", Description: "una descripcion", Price: 1000.5}, {ID: "2", SellerID: "AC234", Description: "otra descripcion", Price: 100.12}}

	mock := MockRepository{Products: products}
	service := NewService(mock)
	p := NewHandler(service)

	r := gin.Default()

	r.GET("/products/", p.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetAllHandlerWithoutParam(t *testing.T) {
	server := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/", "")

	server.ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Code)
}

func TestGetAllHandler(t *testing.T) {
	server := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/?seller_id=123", "")

	server.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestGetAllNotFoundID(t *testing.T) {
	server := createServer()
	req, rr := createRequestTest(http.MethodGet, "/products/?seller_id=444", "")

	server.ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Code)
}
