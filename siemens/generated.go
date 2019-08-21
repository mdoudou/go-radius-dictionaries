// Code generated by radius-dict-gen. DO NOT EDIT.

package siemens

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Siemens_VendorID = 4329
)

func _Siemens_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Siemens_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Siemens_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Siemens_VendorID {
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

func _Siemens_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Siemens_VendorID {
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

func _Siemens_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Siemens_VendorID {
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
	return _Siemens_AddVendor(p, typ, attr)
}

func _Siemens_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Siemens_VendorID {
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

func SiemensURLRedirection_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 1, a)
}

func SiemensURLRedirection_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 1, a)
}

func SiemensURLRedirection_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensURLRedirection_Lookup(p)
	return
}

func SiemensURLRedirection_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensURLRedirection_LookupString(p)
	return
}

func SiemensURLRedirection_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensURLRedirection_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensURLRedirection_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensURLRedirection_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensURLRedirection_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 1, a)
}

func SiemensURLRedirection_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 1, a)
}

func SiemensURLRedirection_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 1)
}

func SiemensAPName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 2, a)
}

func SiemensAPName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 2, a)
}

func SiemensAPName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensAPName_Lookup(p)
	return
}

func SiemensAPName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensAPName_LookupString(p)
	return
}

func SiemensAPName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensAPName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensAPName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensAPName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensAPName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 2, a)
}

func SiemensAPName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 2, a)
}

func SiemensAPName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 2)
}

func SiemensAPSerial_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 3, a)
}

func SiemensAPSerial_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 3, a)
}

func SiemensAPSerial_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensAPSerial_Lookup(p)
	return
}

func SiemensAPSerial_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensAPSerial_LookupString(p)
	return
}

func SiemensAPSerial_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensAPSerial_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensAPSerial_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensAPSerial_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensAPSerial_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 3, a)
}

func SiemensAPSerial_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 3, a)
}

func SiemensAPSerial_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 3)
}

func SiemensVNSName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 4, a)
}

func SiemensVNSName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 4, a)
}

func SiemensVNSName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensVNSName_Lookup(p)
	return
}

func SiemensVNSName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensVNSName_LookupString(p)
	return
}

func SiemensVNSName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 4) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensVNSName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 4) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensVNSName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensVNSName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensVNSName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 4, a)
}

func SiemensVNSName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 4, a)
}

func SiemensVNSName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 4)
}

func SiemensSSID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 5, a)
}

func SiemensSSID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 5, a)
}

func SiemensSSID_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensSSID_Lookup(p)
	return
}

func SiemensSSID_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensSSID_LookupString(p)
	return
}

func SiemensSSID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensSSID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensSSID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensSSID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensSSID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 5, a)
}

func SiemensSSID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 5, a)
}

func SiemensSSID_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 5)
}

func SiemensBSSMAC_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 6, a)
}

func SiemensBSSMAC_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 6, a)
}

func SiemensBSSMAC_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensBSSMAC_Lookup(p)
	return
}

func SiemensBSSMAC_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensBSSMAC_LookupString(p)
	return
}

func SiemensBSSMAC_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensBSSMAC_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensBSSMAC_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensBSSMAC_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensBSSMAC_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 6, a)
}

func SiemensBSSMAC_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 6, a)
}

func SiemensBSSMAC_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 6)
}

func SiemensPolicyName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 7, a)
}

func SiemensPolicyName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 7, a)
}

func SiemensPolicyName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensPolicyName_Lookup(p)
	return
}

func SiemensPolicyName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensPolicyName_LookupString(p)
	return
}

func SiemensPolicyName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 7) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensPolicyName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 7) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensPolicyName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensPolicyName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensPolicyName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 7, a)
}

func SiemensPolicyName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 7, a)
}

func SiemensPolicyName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 7)
}

func SiemensTopologyName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 8, a)
}

func SiemensTopologyName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 8, a)
}

func SiemensTopologyName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensTopologyName_Lookup(p)
	return
}

func SiemensTopologyName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensTopologyName_LookupString(p)
	return
}

func SiemensTopologyName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 8) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensTopologyName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 8) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensTopologyName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensTopologyName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensTopologyName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 8, a)
}

func SiemensTopologyName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 8, a)
}

func SiemensTopologyName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 8)
}

func SiemensIngressRCName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 9, a)
}

func SiemensIngressRCName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 9, a)
}

func SiemensIngressRCName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensIngressRCName_Lookup(p)
	return
}

func SiemensIngressRCName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensIngressRCName_LookupString(p)
	return
}

func SiemensIngressRCName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 9) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensIngressRCName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 9) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensIngressRCName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensIngressRCName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 9)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensIngressRCName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 9, a)
}

func SiemensIngressRCName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 9, a)
}

func SiemensIngressRCName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 9)
}

func SiemensEgressRCName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 10, a)
}

func SiemensEgressRCName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_AddVendor(p, 10, a)
}

func SiemensEgressRCName_Get(p *radius.Packet) (value []byte) {
	value, _ = SiemensEgressRCName_Lookup(p)
	return
}

func SiemensEgressRCName_GetString(p *radius.Packet) (value string) {
	value, _ = SiemensEgressRCName_LookupString(p)
	return
}

func SiemensEgressRCName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Siemens_GetsVendor(p, 10) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensEgressRCName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Siemens_GetsVendor(p, 10) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func SiemensEgressRCName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Siemens_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func SiemensEgressRCName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Siemens_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func SiemensEgressRCName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 10, a)
}

func SiemensEgressRCName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Siemens_SetVendor(p, 10, a)
}

func SiemensEgressRCName_Del(p *radius.Packet) {
	_Siemens_DelVendor(p, 10)
}
