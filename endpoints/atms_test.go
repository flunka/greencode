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

func TestOrderATM(t *testing.T) {
	cases := []struct {
		payload  []byte
		response []byte
	}{
		{[]byte(`[
			{
			"region": 4,
			"requestType": "STANDARD",
			"atmId": 1
			},
			{
			"region": 1,
			"requestType": "STANDARD",
			"atmId": 1
			},
			{
			"region": 2,
			"requestType": "STANDARD",
			"atmId": 1
			},
			{
			"region": 3,
			"requestType": "PRIORITY",
			"atmId": 2
			},
			{
			"region": 3,
			"requestType": "STANDARD",
			"atmId": 1
			},
			{
			"region": 2,
			"requestType": "SIGNAL_LOW",
			"atmId": 1
			},
			{
			"region": 5,
			"requestType": "STANDARD",
			"atmId": 2
			},
			{
			"region": 5,
			"requestType": "FAILURE_RESTART",
			"atmId": 1
			}
		]
		`), []byte(`[
			{
			"region": 1,
			"atmId": 1
			},
			{
			"region": 2,
			"atmId": 1
			},
			{
			"region": 3,
			"atmId": 2
			},
			{
			"region": 3,
			"atmId": 1
			},
			{
			"region": 4,
			"atmId": 1
			},
			{
			"region": 5,
			"atmId": 1
			},
			{
			"region": 5,
			"atmId": 2
			}
		]`)},
		{[]byte(`[
			{
			  "region": 1,
			  "requestType": "STANDARD",
			  "atmId": 2
			},
			{
			  "region": 1,
			  "requestType": "STANDARD",
			  "atmId": 1
			},
			{
			  "region": 2,
			  "requestType": "PRIORITY",
			  "atmId": 3
			},
			{
			  "region": 3,
			  "requestType": "STANDARD",
			  "atmId": 4
			},
			{
			  "region": 4,
			  "requestType": "STANDARD",
			  "atmId": 5
			},
			{
			  "region": 5,
			  "requestType": "PRIORITY",
			  "atmId": 2
			},
			{
			  "region": 5,
			  "requestType": "STANDARD",
			  "atmId": 1
			},
			{
			  "region": 3,
			  "requestType": "SIGNAL_LOW",
			  "atmId": 2
			},
			{
			  "region": 2,
			  "requestType": "SIGNAL_LOW",
			  "atmId": 1
			},
			{
			  "region": 3,
			  "requestType": "FAILURE_RESTART",
			  "atmId": 1
			}
		  ]
		  `), []byte(`[
			{
			  "region": 1,
			  "atmId": 2
			},
			{
			  "region": 1,
			  "atmId": 1
			},
			{
			  "region": 2,
			  "atmId": 3
			},
			{
			  "region": 2,
			  "atmId": 1
			},
			{
			  "region": 3,
			  "atmId": 1
			},
			{
			  "region": 3,
			  "atmId": 2
			},
			{
			  "region": 3,
			  "atmId": 4
			},
			{
			  "region": 4,
			  "atmId": 5
			},
			{
			  "region": 5,
			  "atmId": 2
			},
			{
			  "region": 5,
			  "atmId": 1
			}
		  ]
		  `)},
	}
	for _, c := range cases {
		router := gin.Default()
		router.POST(ATMEndpoint, Order)

		w := httptest.NewRecorder()
		var jsonData = c.payload
		request, _ := http.NewRequest("POST", ATMEndpoint, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, request)

		want_bytes := c.response
		var got []ATM

		var want []ATM
		_ = json.Unmarshal(want_bytes, &want)
		_ = json.Unmarshal(w.Body.Bytes(), &got)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, want, got)
	}

}
