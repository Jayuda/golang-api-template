package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetJsonFieldsString(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	GetJSONFields("Hello")
}

func TestMessageDataFoundOrNotSliceInt(t *testing.T) {
	message := MessageDataFoundOrNot([]int{1, 3, 4, 5, 6, 5, 3})
	assert.Equal(t, "Record found", message, "Result must be 'Record found'")

	message2 := MessageDataFoundOrNot([]int{})
	assert.Equal(t, "Record not found", message2, "Result must be 'Record not found'")
}

func TestMessageDataFoundOrNotStruct(t *testing.T) {
	// message := MessageDataFoundOrNot(domain.City{
	// 	ID:   "JKT",
	// 	Name: "Jakarta",
	// })
	// assert.Equal(t, "Record found", message, "Result must be 'Record found'")

	message2 := MessageDataFoundOrNot(nil)
	assert.Equal(t, "Record not found", message2, "Result must be 'Record not found'")
}

func TestMessageDataFoundOrNotString(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	MessageDataFoundOrNot("Hello")
}

func TestMessageDataFoundOrNotInt(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	MessageDataFoundOrNot(34)
}

func TestMessageDataFoundOrNotBoolean(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	MessageDataFoundOrNot(true)
}
