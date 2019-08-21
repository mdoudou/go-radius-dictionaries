// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc7055

import (
	"layeh.com/radius"
)

const (
	GSSAcceptorServiceName_Type      radius.Type = 164
	GSSAcceptorHostName_Type         radius.Type = 165
	GSSAcceptorServiceSpecifics_Type radius.Type = 166
	GSSAcceptorRealmName_Type        radius.Type = 167
)

func GSSAcceptorServiceName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorServiceName_Type, a)
	return
}

func GSSAcceptorServiceName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorServiceName_Type, a)
	return
}

func GSSAcceptorServiceName_Get(p *radius.Packet) (value []byte) {
	value, _ = GSSAcceptorServiceName_Lookup(p)
	return
}

func GSSAcceptorServiceName_GetString(p *radius.Packet) (value string) {
	value, _ = GSSAcceptorServiceName_LookupString(p)
	return
}

func GSSAcceptorServiceName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[GSSAcceptorServiceName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorServiceName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[GSSAcceptorServiceName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorServiceName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(GSSAcceptorServiceName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GSSAcceptorServiceName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(GSSAcceptorServiceName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GSSAcceptorServiceName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorServiceName_Type, a)
	return
}

func GSSAcceptorServiceName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorServiceName_Type, a)
	return
}

func GSSAcceptorServiceName_Del(p *radius.Packet) {
	p.Attributes.Del(GSSAcceptorServiceName_Type)
}

func GSSAcceptorHostName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorHostName_Type, a)
	return
}

func GSSAcceptorHostName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorHostName_Type, a)
	return
}

func GSSAcceptorHostName_Get(p *radius.Packet) (value []byte) {
	value, _ = GSSAcceptorHostName_Lookup(p)
	return
}

func GSSAcceptorHostName_GetString(p *radius.Packet) (value string) {
	value, _ = GSSAcceptorHostName_LookupString(p)
	return
}

func GSSAcceptorHostName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[GSSAcceptorHostName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorHostName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[GSSAcceptorHostName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorHostName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(GSSAcceptorHostName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GSSAcceptorHostName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(GSSAcceptorHostName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GSSAcceptorHostName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorHostName_Type, a)
	return
}

func GSSAcceptorHostName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorHostName_Type, a)
	return
}

func GSSAcceptorHostName_Del(p *radius.Packet) {
	p.Attributes.Del(GSSAcceptorHostName_Type)
}

func GSSAcceptorServiceSpecifics_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorServiceSpecifics_Type, a)
	return
}

func GSSAcceptorServiceSpecifics_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorServiceSpecifics_Type, a)
	return
}

func GSSAcceptorServiceSpecifics_Get(p *radius.Packet) (value []byte) {
	value, _ = GSSAcceptorServiceSpecifics_Lookup(p)
	return
}

func GSSAcceptorServiceSpecifics_GetString(p *radius.Packet) (value string) {
	value, _ = GSSAcceptorServiceSpecifics_LookupString(p)
	return
}

func GSSAcceptorServiceSpecifics_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[GSSAcceptorServiceSpecifics_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorServiceSpecifics_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[GSSAcceptorServiceSpecifics_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorServiceSpecifics_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(GSSAcceptorServiceSpecifics_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GSSAcceptorServiceSpecifics_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(GSSAcceptorServiceSpecifics_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GSSAcceptorServiceSpecifics_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorServiceSpecifics_Type, a)
	return
}

func GSSAcceptorServiceSpecifics_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorServiceSpecifics_Type, a)
	return
}

func GSSAcceptorServiceSpecifics_Del(p *radius.Packet) {
	p.Attributes.Del(GSSAcceptorServiceSpecifics_Type)
}

func GSSAcceptorRealmName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorRealmName_Type, a)
	return
}

func GSSAcceptorRealmName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(GSSAcceptorRealmName_Type, a)
	return
}

func GSSAcceptorRealmName_Get(p *radius.Packet) (value []byte) {
	value, _ = GSSAcceptorRealmName_Lookup(p)
	return
}

func GSSAcceptorRealmName_GetString(p *radius.Packet) (value string) {
	value, _ = GSSAcceptorRealmName_LookupString(p)
	return
}

func GSSAcceptorRealmName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[GSSAcceptorRealmName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorRealmName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[GSSAcceptorRealmName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func GSSAcceptorRealmName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(GSSAcceptorRealmName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func GSSAcceptorRealmName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(GSSAcceptorRealmName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func GSSAcceptorRealmName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorRealmName_Type, a)
	return
}

func GSSAcceptorRealmName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(GSSAcceptorRealmName_Type, a)
	return
}

func GSSAcceptorRealmName_Del(p *radius.Packet) {
	p.Attributes.Del(GSSAcceptorRealmName_Type)
}
