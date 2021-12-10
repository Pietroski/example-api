package binders

import (
	"encoding/json"
	"encoding/xml"
)

type Binder interface {
	Bind(body string, obj interface{}) error
}

type (
	bind struct {
		body string
		obj  interface{}
	}

	bindWithJSON struct{}
	bindWithXML  struct{}
)

var (
	JSON Binder = &bindWithJSON{}
	XML  Binder = &bindWithXML{}
)

func (b *bind) shouldBindWith(bindingStruct Binder) error {
	return bindingStruct.Bind(b.body, b.obj)
}

func ShouldBindJSON(body string, obj interface{}) error {
	b := bind{
		body: body,
		obj:  obj,
	}

	return b.shouldBindWith(JSON)
}

func ShouldBindXML(body string, obj interface{}) error {
	b := bind{
		body: body,
		obj:  obj,
	}

	return b.shouldBindWith(XML)
}

func (bJSON *bindWithJSON) Bind(body string, obj interface{}) error {
	reqBody := []byte(body)
	if err := json.Unmarshal(reqBody, obj); err != nil {
		return err
	}

	return nil
}

func (bJSON *bindWithXML) Bind(body string, obj interface{}) error {
	reqBody := []byte(body)
	if err := xml.Unmarshal(reqBody, obj); err != nil {
		return err
	}

	return nil
}
