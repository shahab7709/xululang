package xulu

import "testing"

func TestXuluLang(t *testing.T) {
	//init
	input := "abcd abcd aabbc ab a c ccd dede cccd cd"
	xuluProvider := NewXuluLang(input)
	sentece, err := xuluProvider.create()

	//execute
	if err != nil {
		t.Error(err)
	}
	result, err := sentece.compute()
	if err != nil {
		t.Error(err)
	}

	//exercise

	if result != 861 {
		t.Errorf("result not ok %v = %v", input, result)
	}
	//teardown
	t.Logf("succeed completely!!")

}
