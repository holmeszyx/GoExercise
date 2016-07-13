package str

import "testing"

func Test_convert(t *testing.T) {
	assertEqual(convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR", t)
	assertEqual(convert("ABC", 2), "ACB", t)
	assertEqual(convert("ABCD", 2), "ACBD", t)
	assertEqual(convert("ABCDE", 4), "ABCED", t)
}

func assertEqual(a interface{}, b interface{}, t *testing.T) {
	if a != b {
		t.Error(a, "not", b)
	}
}
