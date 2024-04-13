package calcbuilder

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildCalcFunc(t *testing.T) {

	t.Run("true", func(t *testing.T) {
		file, err := os.Open("test_data/true_tests.json")
		require.Nil(t, err)

		decoder := json.NewDecoder(file)
		var data []DataSet
		require.Nil(t, decoder.Decode(&data))

		truTest(t, data)
	})

	// TO DO add error tests
}

type DataSet struct {
	Formula string    `json:"formula"`
	Vars    []float64 `json:"vars"`
	Result  float64   `json:"result"`
}

func truTest(t *testing.T, data []DataSet) {
	t.Helper()
	for _, d := range data {
		f, err := BuildСalcFunc(d.Formula, func(i int) float64 { return d.Vars[i] })
		require.Nil(t, err, d)
		require.Equal(t, d.Result, f(), d.Formula)
	}
}
