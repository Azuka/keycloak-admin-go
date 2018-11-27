package keycloak

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var timeMarshalCases = []struct {
	Time          UnixTime
	ExpectedValue []byte
}{
	{
		Time:          UnixTime(time.Unix(1257894000, 78000000)),
		ExpectedValue: []byte("1257894000078"),
	},
	{
		Time:          UnixTime(time.Unix(1530029597, 8000000)),
		ExpectedValue: []byte("1530029597008"),
	},
}

func TestTimeUnmarshalFails(t *testing.T) {
	a := assert.New(t)

	var v2 UnixTime
	v := "hogehoge"
	err := json.Unmarshal([]byte(v), &v2)

	a.Error(err)
}

func TestTimeMarshalUnmarshal(t *testing.T) {
	a := assert.New(t)

	for _, tt := range timeMarshalCases {
		v, err := json.Marshal(tt.Time)

		a.NoError(err)
		a.Equal(tt.ExpectedValue, v)

		var v2 UnixTime

		err = json.Unmarshal(v, &v2)

		a.NoError(err)
		a.NotNil(v2)
		a.Equal(tt.Time.String(), v2.String())
	}
}
