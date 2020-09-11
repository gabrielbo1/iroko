package msg

import (
	"github.com/gabrielbo1/iroko/domain"
	"testing"
)

func TestMessage(t *testing.T) {
	err := domain.NewErr().WithCode("hello_word")
	pt := "hello_word - Hello Word. Mensagem de teste."
	en := "hello_word - Hello Word. Test message."
	if pt != Message(err, "", "").OnError().Error() {
		t.Error("Msg internationalization portuguese error!")
	}
	if pt != Message(err, "pt", "").OnError().Error() {
		t.Error("Msg internationalization portuguese error!")
	}
	if en != Message(err, "en", "").OnError().Error() {
		t.Error("Msg internationalization english error!")
	}

	err = domain.NewErr().
		WithCode("hello_word_param").
		WithMsgParam(map[string]string{"Name": "gabriel"})

	pt = "hello_word_param - Hello Word. Mensagem de teste gabriel."
	en = "hello_word_param - Hello Word. Test message gabriel."
	if pt != Message(err, "pt", "").OnError().Error() {
		t.Error("Msg internationalization portuguese error!")
	}
	if en != Message(err, "en", "").OnError().Error() {
		t.Error("Msg internationalization english error!")
	}
}
