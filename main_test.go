package mysqlfsencoding

import (
	"testing"
)

func TestEncDec(t *testing.T) {
	var strings = []string{
		"stathis-2021",
		"stathis@002d2021",

		"greek-αΑβΒγΓ",
		"greek@002d@6l@6L@6m@7W@6n@6N",
	}

	for i := 0; i < len(strings); i += 2 {
		if strings[i + 1] != encode(strings[i]) {
			t.Errorf("Error encoding %v", strings[i])
		}
		if strings[i] != decode(strings[i + 1]) {
			t.Errorf("Error decoding %v", strings[i + 1])
		}
	}
}

