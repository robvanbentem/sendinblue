package sib

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestNewEmail(t *testing.T) {

	from := [2]string{"1", "2"}
	stringMap := make(map[string]string)

	email := NewEmail()

	if reflect.TypeOf(email.From) != reflect.TypeOf(from) {
		t.Error("The From Type is not correct.")
	}

	if reflect.TypeOf(email.Attachment) != reflect.TypeOf(stringMap) {
		t.Error("The Attachment Type is not correct.")
	}

	if reflect.TypeOf(email.Bcc) != reflect.TypeOf(stringMap) {
		t.Error("The BCC Type is not correct.")
	}

	if reflect.TypeOf(email.CC) != reflect.TypeOf(stringMap) {
		t.Error("The CC Type is not correct.")
	}

	if reflect.TypeOf(email.Headers) != reflect.TypeOf(stringMap) {
		t.Error("The Headers Type is not correct.")
	}

	if reflect.TypeOf(email.To) != reflect.TypeOf(stringMap) {
		t.Error("The To Type is not correct.")
	}

	if reflect.TypeOf(email.Inline_image) != reflect.TypeOf(stringMap) {
		t.Error("The Inline_image Type is not correct.")
	}
}

func TestNewEmailOptions(t *testing.T) {

	myEmails := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
	}
	ccjoin := strings.Join(myEmails, "|")
	bccjoin := strings.Join(myEmails, "|")
	stringMap := make(map[string]string)
	testString := "test"

	options := NewEmailOptions("reply", "attach.ment", myEmails, myEmails)

	if reflect.TypeOf(options.Cc) != reflect.TypeOf(testString) {
		t.Error("The CC Type is not correct.")
	}

	if options.Cc != ccjoin {
		t.Error("Array of email strings is not being properly converted to a single string.")
	}

	if reflect.TypeOf(options.Bcc) != reflect.TypeOf(testString) {
		t.Error("The BCC Type is not correct.")
	}

	if options.Bcc != bccjoin {
		t.Error("Array of email strings is not being properly converted to a single string.")
	}

	if reflect.TypeOf(options.Attr) != reflect.TypeOf(stringMap) {
		t.Error("The Attr Type is not correct.")
	}

	if reflect.TypeOf(options.Attachment) != reflect.TypeOf(stringMap) {
		t.Error("The Attachment Type is not correct.")
	}

	if reflect.TypeOf(options.Headers) != reflect.TypeOf(stringMap) {
		t.Error("The Headers Type is not correct.")
	}

	if options.ReplyTo != "reply" {
		t.Error("The ReplyTo field is not being set.")
	}

	if options.Attachment_url != "attach.ment" {
		t.Error("The Attachment field is not being set.")
	}
}

func TestEmailAddImage(t *testing.T) {
	email := NewEmail()

	f, _ := os.Open("./test/myimage.jpg")
	name := email.AddImage(f)

	if email.Inline_image[name] == "" {
		t.Error("Image is not being added in AddImage Email method.")
	}

	f.Close()
}

func TestEmailOptionsAddAttachment(t *testing.T) {
	options := NewEmailOptions("reply", "attach.ment", nil, nil)

	f, _ := os.Open("./test/attachment.pdf")
	options.AddAttachment(f)

	if options.Attachment[f.Name()] == "" {
		t.Error("Attachments are not being added with EmailOptions AddAttachment method.")
	}

	f.Close()
}
