package areena

import "testing"

func TestSlugify(t *testing.T) {
	testData := map[string]string{"Katti Matikainen - 4 - Päivä": "katti-matikainen-4-paiva"}
	for tstStr, expected := range testData {
		s := Slugify(tstStr)
		if s != expected {
			t.Errorf("expected %s, got %s", expected, s)
		}
	}
}
