package domain

import (
	"errors"
	"fmt"
)

//Err - Define pattern error return.
type Err struct {
	code     string `json:"code"`
	message  string `json:"message"`
	err      error  `json:"err"`
	msgParam map[string]string
}

//NewErr - New Err set properties with builder function.
func NewErr() *Err {
	return &Err{}
}

//WithCode - Set Code.
func (e *Err) WithCode(code string) *Err {
	e.code = code
	return e
}

//GetCode - Code internationalization message.
func (e *Err) GetCode() string {
	return e.code
}

//WithMessage - Set Message.
func (e *Err) WithMessage(msg string) *Err {
	e.message = msg
	return e
}

//WithError - Set Err.
func (e *Err) WithError(error error) *Err {
	e.err = error
	return e
}

//WithMsgParam - Set params.
func (e *Err) WithMsgParam(param map[string]string) *Err {
	e.msgParam = param
	return e
}

//GetMsgParam - Return msg params.
func (e *Err) GetMsgParam() map[string]string {
	return e.msgParam
}

//WithErr - Build Err  with Err.
func (e *Err) WithErr(k *Err) *Err {
	e.code = k.code
	e.message = k.message
	e.msgParam = k.msgParam
	e.err = k.err
	return e
}

//OnError - Convert Err type to error type.
func (e *Err) OnError() error {
	if e == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s - %s", e.code, e.message))
}
