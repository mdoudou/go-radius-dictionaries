// Code generated by radius-dict-gen. DO NOT EDIT.

package microsemi

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Microsemi_VendorID = 40676
)

func _Microsemi_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Microsemi_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Microsemi_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Microsemi_VendorID {
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

func _Microsemi_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Microsemi_VendorID {
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

func _Microsemi_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Microsemi_VendorID {
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
	return _Microsemi_AddVendor(p, typ, attr)
}

func _Microsemi_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Microsemi_VendorID {
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

func MicrosemiUserFullName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 1, a)
}

func MicrosemiUserFullName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 1, a)
}

func MicrosemiUserFullName_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiUserFullName_Lookup(p)
	return
}

func MicrosemiUserFullName_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiUserFullName_LookupString(p)
	return
}

func MicrosemiUserFullName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserFullName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserFullName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiUserFullName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiUserFullName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 1, a)
}

func MicrosemiUserFullName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 1, a)
}

func MicrosemiUserFullName_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 1)
}

func MicrosemiUserName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 2, a)
}

func MicrosemiUserName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 2, a)
}

func MicrosemiUserName_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiUserName_Lookup(p)
	return
}

func MicrosemiUserName_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiUserName_LookupString(p)
	return
}

func MicrosemiUserName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiUserName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiUserName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 2, a)
}

func MicrosemiUserName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 2, a)
}

func MicrosemiUserName_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 2)
}

func MicrosemiUserInitials_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 3, a)
}

func MicrosemiUserInitials_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 3, a)
}

func MicrosemiUserInitials_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiUserInitials_Lookup(p)
	return
}

func MicrosemiUserInitials_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiUserInitials_LookupString(p)
	return
}

func MicrosemiUserInitials_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserInitials_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserInitials_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiUserInitials_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiUserInitials_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 3, a)
}

func MicrosemiUserInitials_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 3, a)
}

func MicrosemiUserInitials_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 3)
}

func MicrosemiUserEmail_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 4, a)
}

func MicrosemiUserEmail_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 4, a)
}

func MicrosemiUserEmail_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiUserEmail_Lookup(p)
	return
}

func MicrosemiUserEmail_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiUserEmail_LookupString(p)
	return
}

func MicrosemiUserEmail_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserEmail_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserEmail_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiUserEmail_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiUserEmail_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 4, a)
}

func MicrosemiUserEmail_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 4, a)
}

func MicrosemiUserEmail_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 4)
}

func MicrosemiUserGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 5, a)
}

func MicrosemiUserGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 5, a)
}

func MicrosemiUserGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiUserGroup_Lookup(p)
	return
}

func MicrosemiUserGroup_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiUserGroup_LookupString(p)
	return
}

func MicrosemiUserGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiUserGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiUserGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiUserGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 5, a)
}

func MicrosemiUserGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 5, a)
}

func MicrosemiUserGroup_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 5)
}

func MicrosemiFallbackUserGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 6, a)
}

func MicrosemiFallbackUserGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 6, a)
}

func MicrosemiFallbackUserGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiFallbackUserGroup_Lookup(p)
	return
}

func MicrosemiFallbackUserGroup_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiFallbackUserGroup_LookupString(p)
	return
}

func MicrosemiFallbackUserGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiFallbackUserGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiFallbackUserGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiFallbackUserGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiFallbackUserGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 6, a)
}

func MicrosemiFallbackUserGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 6, a)
}

func MicrosemiFallbackUserGroup_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 6)
}

func MicrosemiNetworkElementGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 7, a)
}

func MicrosemiNetworkElementGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_AddVendor(p, 7, a)
}

func MicrosemiNetworkElementGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = MicrosemiNetworkElementGroup_Lookup(p)
	return
}

func MicrosemiNetworkElementGroup_GetString(p *radius.Packet) (value string) {
	value, _ = MicrosemiNetworkElementGroup_LookupString(p)
	return
}

func MicrosemiNetworkElementGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Microsemi_GetsVendor(p, 7) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiNetworkElementGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Microsemi_GetsVendor(p, 7) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MicrosemiNetworkElementGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Microsemi_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MicrosemiNetworkElementGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Microsemi_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MicrosemiNetworkElementGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 7, a)
}

func MicrosemiNetworkElementGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Microsemi_SetVendor(p, 7, a)
}

func MicrosemiNetworkElementGroup_Del(p *radius.Packet) {
	_Microsemi_DelVendor(p, 7)
}