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
func (err *Err) WithCode(code string) *Err {
	err.code = code
	return err
}

//GetCode - Code internationalization message.
func (err *Err) GetCode() string {
	return err.code
}

//WithMessage - Set Message.
func (err *Err) WithMessage(msg string) *Err {
	err.message = msg
	return err
}

//WithError - Set Err.
func (err *Err) WithError(error error) *Err {
	err.err = error
	return err
}

//WithMsgParam - Set params.
func (err *Err) WithMsgParam(param map[string]string) *Err {
	err.msgParam = param
	return err
}

//GetMsgParam - Return msg params.
func (err *Err) GetMsgParam() map[string]string {
	return err.msgParam
}

func (err *Err) WithErr(e *Err) *Err {
	err.code = e.code
	err.message = e.message
	err.msgParam = e.msgParam
	err.err = e.err
	return err
}

//OnError - Convert Err type to error type.
func (e *Err) OnError() error {
	if e == nil {
		return nil
	}
	return errors.New(fmt.Sprintf("%s - %s", e.code, e.message))
}
