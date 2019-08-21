// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc4072

import (
	"layeh.com/radius"
)

const (
	EAPKeyName_Type radius.Type = 102
)

func EAPKeyName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(EAPKeyName_Type, a)
	return
}

func EAPKeyName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(EAPKeyName_Type, a)
	return
}

func EAPKeyName_Get(p *radius.Packet) (value []byte) {
	value, _ = EAPKeyName_Lookup(p)
	return
}

func EAPKeyName_GetString(p *radius.Packet) (value string) {
	value, _ = EAPKeyName_LookupString(p)
	return
}

func EAPKeyName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[EAPKeyName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPKeyName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[EAPKeyName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPKeyName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(EAPKeyName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func EAPKeyName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(EAPKeyName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func EAPKeyName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(EAPKeyName_Type, a)
	return
}

func EAPKeyName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(EAPKeyName_Type, a)
	return
}

func EAPKeyName_Del(p *radius.Packet) {
	p.Attributes.Del(EAPKeyName_Type)
}
