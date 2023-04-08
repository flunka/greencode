package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestReportTransactions(t *testing.T) {
	cases := []struct {
		payload  []byte
		response []byte
	}{
		{[]byte(`[
			{
			  "debitAccount": "32309111922661937852684864",
			  "creditAccount": "06105023389842834748547303",
			  "amount": 10.90
			},
			{
			  "debitAccount": "31074318698137062235845814",
			  "creditAccount": "66105036543749403346524547",
			  "amount": 200.90
			},
			{
			  "debitAccount": "66105036543749403346524547",
			  "creditAccount": "32309111922661937852684864",
			  "amount": 50.10
			}
		  ]
		  
		`), []byte(`[
			{
			  "account": "06105023389842834748547303",
			  "debitCount": 0,
			  "creditCount": 1,
			  "balance": 10.90
			},
			{
			  "account": "31074318698137062235845814",
			  "debitCount": 1,
			  "creditCount": 0,
			  "balance": -200.90
			},
			{
			  "account": "32309111922661937852684864",
			  "debitCount": 1,
			  "creditCount": 1,
			  "balance": 39.20
			},
			{
			  "account": "66105036543749403346524547",
			  "debitCount": 1,
			  "creditCount": 1,
			  "balance": 150.80
			}
		  ]
		  `)},
	}
	for _, c := range cases {
		router := gin.Default()
		router.POST(TransactionsEndpoint, Report)

		w := httptest.NewRecorder()
		var jsonData = c.payload
		request, _ := http.NewRequest("POST", TransactionsEndpoint, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, request)

		want_bytes := c.response
		var got []Account

		var want []Account
		_ = json.Unmarshal(want_bytes, &want)
		_ = json.Unmarshal(w.Body.Bytes(), &got)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, want, got)
	}

}
