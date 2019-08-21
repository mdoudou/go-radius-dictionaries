// Code generated by radius-dict-gen. DO NOT EDIT.

package telkom

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Telkom_VendorID = 1431
)

func _Telkom_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Telkom_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Telkom_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Telkom_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Telkom_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Telkom_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Telkom_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Telkom_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(p.Attributes[rfc2865.VendorSpecific_Type][i][4:], vsa)
			i++
		} else {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+i:]...)
		}
	}
	return _Telkom_AddVendor(p, typ, attr)
}

func _Telkom_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Telkom_VendorID {
			continue
		}
		offset := 0
		for len(vsa[offset:]) >= 3 {
			vsaTyp, vsaLen := vsa[offset], vsa[offset+1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				continue vsaLoop
			}
			if vsaTyp == typ {
				copy(vsa[offset:], vsa[offset+int(vsaLen):])
				vsa = vsa[:len(vsa)-int(vsaLen)]
			} else {
				offset += int(vsaLen)
			}
		}
		if offset == 0 {
			p.Attributes[rfc2865.VendorSpecific_Type] = append(p.Attributes[rfc2865.VendorSpecific_Type][:i], p.Attributes[rfc2865.VendorSpecific_Type][i+1:]...)
		} else {
			i++
		}
	}
	return
}

func TelkomAccessType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 1, a)
}

func TelkomAccessType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 1, a)
}

func TelkomAccessType_Get(p *radius.Packet) (value []byte) {
	value, _ = TelkomAccessType_Lookup(p)
	return
}

func TelkomAccessType_GetString(p *radius.Packet) (value string) {
	value, _ = TelkomAccessType_LookupString(p)
	return
}

func TelkomAccessType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Telkom_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomAccessType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Telkom_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomAccessType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Telkom_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TelkomAccessType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Telkom_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TelkomAccessType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 1, a)
}

func TelkomAccessType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 1, a)
}

func TelkomAccessType_Del(p *radius.Packet) {
	_Telkom_DelVendor(p, 1)
}

func TelkomServiceType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 2, a)
}

func TelkomServiceType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 2, a)
}

func TelkomServiceType_Get(p *radius.Packet) (value []byte) {
	value, _ = TelkomServiceType_Lookup(p)
	return
}

func TelkomServiceType_GetString(p *radius.Packet) (value string) {
	value, _ = TelkomServiceType_LookupString(p)
	return
}

func TelkomServiceType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Telkom_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomServiceType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Telkom_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomServiceType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Telkom_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TelkomServiceType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Telkom_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TelkomServiceType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 2, a)
}

func TelkomServiceType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 2, a)
}

func TelkomServiceType_Del(p *radius.Packet) {
	_Telkom_DelVendor(p, 2)
}

func TelkomDegradeToken_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 200, a)
}

func TelkomDegradeToken_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_AddVendor(p, 200, a)
}

func TelkomDegradeToken_Get(p *radius.Packet) (value []byte) {
	value, _ = TelkomDegradeToken_Lookup(p)
	return
}

func TelkomDegradeToken_GetString(p *radius.Packet) (value string) {
	value, _ = TelkomDegradeToken_LookupString(p)
	return
}

func TelkomDegradeToken_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Telkom_GetsVendor(p, 200) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomDegradeToken_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Telkom_GetsVendor(p, 200) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func TelkomDegradeToken_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Telkom_LookupVendor(p, 200)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func TelkomDegradeToken_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Telkom_LookupVendor(p, 200)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func TelkomDegradeToken_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 200, a)
}

func TelkomDegradeToken_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Telkom_SetVendor(p, 200, a)
}

func TelkomDegradeToken_Del(p *radius.Packet) {
	_Telkom_DelVendor(p, 200)
}
