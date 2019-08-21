// Code generated by radius-dict-gen. DO NOT EDIT.

package merit

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Merit_VendorID = 61
)

func _Merit_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Merit_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Merit_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Merit_VendorID {
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

func _Merit_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Merit_VendorID {
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

func _Merit_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Merit_VendorID {
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
	return _Merit_AddVendor(p, typ, attr)
}

func _Merit_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Merit_VendorID {
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

func MeritProxyAction_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 211, a)
}

func MeritProxyAction_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 211, a)
}

func MeritProxyAction_Get(p *radius.Packet) (value []byte) {
	value, _ = MeritProxyAction_Lookup(p)
	return
}

func MeritProxyAction_GetString(p *radius.Packet) (value string) {
	value, _ = MeritProxyAction_LookupString(p)
	return
}

func MeritProxyAction_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Merit_GetsVendor(p, 211) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritProxyAction_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Merit_GetsVendor(p, 211) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritProxyAction_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Merit_LookupVendor(p, 211)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MeritProxyAction_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Merit_LookupVendor(p, 211)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MeritProxyAction_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 211, a)
}

func MeritProxyAction_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 211, a)
}

func MeritProxyAction_Del(p *radius.Packet) {
	_Merit_DelVendor(p, 211)
}

func MeritUserID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 222, a)
}

func MeritUserID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 222, a)
}

func MeritUserID_Get(p *radius.Packet) (value []byte) {
	value, _ = MeritUserID_Lookup(p)
	return
}

func MeritUserID_GetString(p *radius.Packet) (value string) {
	value, _ = MeritUserID_LookupString(p)
	return
}

func MeritUserID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Merit_GetsVendor(p, 222) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritUserID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Merit_GetsVendor(p, 222) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritUserID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Merit_LookupVendor(p, 222)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MeritUserID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Merit_LookupVendor(p, 222)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MeritUserID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 222, a)
}

func MeritUserID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 222, a)
}

func MeritUserID_Del(p *radius.Packet) {
	_Merit_DelVendor(p, 222)
}

func MeritUserRealm_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 223, a)
}

func MeritUserRealm_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_AddVendor(p, 223, a)
}

func MeritUserRealm_Get(p *radius.Packet) (value []byte) {
	value, _ = MeritUserRealm_Lookup(p)
	return
}

func MeritUserRealm_GetString(p *radius.Packet) (value string) {
	value, _ = MeritUserRealm_LookupString(p)
	return
}

func MeritUserRealm_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Merit_GetsVendor(p, 223) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritUserRealm_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Merit_GetsVendor(p, 223) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MeritUserRealm_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Merit_LookupVendor(p, 223)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MeritUserRealm_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Merit_LookupVendor(p, 223)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MeritUserRealm_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 223, a)
}

func MeritUserRealm_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Merit_SetVendor(p, 223, a)
}

func MeritUserRealm_Del(p *radius.Packet) {
	_Merit_DelVendor(p, 223)
}
