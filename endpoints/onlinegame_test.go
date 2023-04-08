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

func TestCaltulateOrder(t *testing.T) {
	cases := []struct {
		payload  []byte
		response []byte
	}{
		{[]byte(`{
			"groupCount": 6,
			"clans": [
			  {
				"numberOfPlayers": 4,
				"points": 50
			  },
			  {
				"numberOfPlayers": 2,
				"points": 70
			  },
			  {
				"numberOfPlayers": 6,
				"points": 60
			  },
			  {
				"numberOfPlayers": 1,
				"points": 15
			  },
			  {
				"numberOfPlayers": 5,
				"points": 40
			  },
			  {
				"numberOfPlayers": 3,
				"points": 45
			  },
			  {
				"numberOfPlayers": 1,
				"points": 12
			  },
			  {
				"numberOfPlayers": 4,
				"points": 40
			  }
			]
		  }		  
		`), []byte(`[
			[
			  {
				"numberOfPlayers": 2,
				"points": 70
			  },
			  {
				"numberOfPlayers": 4,
				"points": 50
			  }
			],
			[
			  {
				"numberOfPlayers": 6,
				"points": 60
			  }
			],
			[
			  {
				"numberOfPlayers": 3,
				"points": 45
			  },
			  {
				"numberOfPlayers": 1,
				"points": 15
			  },
			  {
				"numberOfPlayers": 1,
				"points": 12
			  }
			],
			[
			  {
				"numberOfPlayers": 4,
				"points": 40
			  }
			],
			[
			  {
				"numberOfPlayers": 5,
				"points": 40
			  }
			]
		  ]		  
		  `)},
	}
	for _, c := range cases {
		router := gin.Default()
		router.POST(OnlineGameEndpoint, CaltulateOrder)

		w := httptest.NewRecorder()
		var jsonData = c.payload
		request, _ := http.NewRequest("POST", OnlineGameEndpoint, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		router.ServeHTTP(w, request)

		want_bytes := c.response
		var got [][]Clan

		var want [][]Clan
		_ = json.Unmarshal(want_bytes, &want)
		_ = json.Unmarshal(w.Body.Bytes(), &got)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, want, got)
	}

}
