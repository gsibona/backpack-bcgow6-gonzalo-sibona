package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// en una situacion normal se haria un mock del db para probar la relacion repository - service - handler, pero dado que el repository dado en si es un mock, decidi usar el repositorio base
/*type MockRepository struct {}

func NewMockRepository() Repository{
	return &MockRepository{}
}

func (r *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}*/

var s = createServer();
func createServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	/*repository := NewMockRepository()*/
	repository := NewRepository()
	service := NewService(repository)
	p := NewHandler(service)

	r := gin.Default()

	r.GET("/products", p.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func Test_GetProducts_Ok(t *testing.T){
	req, rw := createRequestTest(http.MethodGet, "/products", "")
	fmt.Println(req.URL.String())
	q := req.URL.Query()
    q.Add("seller_id", "1")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	s.ServeHTTP(rw, req)
	
	objRes := &[]Product{}
	assert.Equal(t, 200, rw.Code)
	err := json.Unmarshal(rw.Body.Bytes(), &objRes)

	data := reflect.ValueOf(*objRes).Len() // Obteniendo la cantidad de transactions de Data
	assert.Nil(t, err)
	assert.True(t, data > 0)
}