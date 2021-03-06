// Code generated by radius-dict-gen. DO NOT EDIT.

package infonet

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Infonet_VendorID = 4453
)

func _Infonet_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Infonet_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Infonet_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Infonet_VendorID {
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

func _Infonet_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Infonet_VendorID {
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

func _Infonet_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Infonet_VendorID {
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
	return _Infonet_AddVendor(p, typ, attr)
}

func _Infonet_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Infonet_VendorID {
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

func InfonetProxy_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 238, a)
}

func InfonetProxy_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 238, a)
}

func InfonetProxy_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetProxy_Lookup(p)
	return
}

func InfonetProxy_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetProxy_LookupString(p)
	return
}

func InfonetProxy_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 238) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetProxy_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 238) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetProxy_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 238)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetProxy_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 238)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetProxy_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 238, a)
}

func InfonetProxy_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 238, a)
}

func InfonetProxy_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 238)
}

func InfonetConfig_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 239, a)
}

func InfonetConfig_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 239, a)
}

func InfonetConfig_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetConfig_Lookup(p)
	return
}

func InfonetConfig_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetConfig_LookupString(p)
	return
}

func InfonetConfig_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 239) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetConfig_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 239) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetConfig_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 239)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetConfig_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 239)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetConfig_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 239, a)
}

func InfonetConfig_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 239, a)
}

func InfonetConfig_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 239)
}

func InfonetMCSCountry_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 240, a)
}

func InfonetMCSCountry_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 240, a)
}

func InfonetMCSCountry_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSCountry_Lookup(p)
	return
}

func InfonetMCSCountry_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSCountry_LookupString(p)
	return
}

func InfonetMCSCountry_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 240) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSCountry_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 240) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSCountry_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 240)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSCountry_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 240)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSCountry_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 240, a)
}

func InfonetMCSCountry_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 240, a)
}

func InfonetMCSCountry_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 240)
}

func InfonetMCSRegion_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 241, a)
}

func InfonetMCSRegion_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 241, a)
}

func InfonetMCSRegion_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSRegion_Lookup(p)
	return
}

func InfonetMCSRegion_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSRegion_LookupString(p)
	return
}

func InfonetMCSRegion_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 241) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSRegion_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 241) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSRegion_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 241)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSRegion_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 241)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSRegion_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 241, a)
}

func InfonetMCSRegion_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 241, a)
}

func InfonetMCSRegion_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 241)
}

func InfonetMCSOffPeak_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 242, a)
}

func InfonetMCSOffPeak_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 242, a)
}

func InfonetMCSOffPeak_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSOffPeak_Lookup(p)
	return
}

func InfonetMCSOffPeak_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSOffPeak_LookupString(p)
	return
}

func InfonetMCSOffPeak_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 242) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSOffPeak_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 242) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSOffPeak_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 242)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSOffPeak_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 242)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSOffPeak_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 242, a)
}

func InfonetMCSOffPeak_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 242, a)
}

func InfonetMCSOffPeak_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 242)
}

func InfonetMCSOverflow_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 243, a)
}

func InfonetMCSOverflow_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 243, a)
}

func InfonetMCSOverflow_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSOverflow_Lookup(p)
	return
}

func InfonetMCSOverflow_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSOverflow_LookupString(p)
	return
}

func InfonetMCSOverflow_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 243) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSOverflow_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 243) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSOverflow_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 243)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSOverflow_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 243)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSOverflow_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 243, a)
}

func InfonetMCSOverflow_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 243, a)
}

func InfonetMCSOverflow_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 243)
}

func InfonetMCSPort_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 244, a)
}

func InfonetMCSPort_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 244, a)
}

func InfonetMCSPort_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSPort_Lookup(p)
	return
}

func InfonetMCSPort_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSPort_LookupString(p)
	return
}

func InfonetMCSPort_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 244) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSPort_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 244) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSPort_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 244)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSPort_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 244)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSPort_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 244, a)
}

func InfonetMCSPort_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 244, a)
}

func InfonetMCSPort_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 244)
}

func InfonetMCSPortCount_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 245, a)
}

func InfonetMCSPortCount_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 245, a)
}

func InfonetMCSPortCount_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetMCSPortCount_Lookup(p)
	return
}

func InfonetMCSPortCount_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetMCSPortCount_LookupString(p)
	return
}

func InfonetMCSPortCount_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 245) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSPortCount_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 245) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetMCSPortCount_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 245)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetMCSPortCount_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 245)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetMCSPortCount_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 245, a)
}

func InfonetMCSPortCount_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 245, a)
}

func InfonetMCSPortCount_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 245)
}

func InfonetAccountNumber_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 247, a)
}

func InfonetAccountNumber_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 247, a)
}

func InfonetAccountNumber_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetAccountNumber_Lookup(p)
	return
}

func InfonetAccountNumber_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetAccountNumber_LookupString(p)
	return
}

func InfonetAccountNumber_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 247) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetAccountNumber_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 247) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetAccountNumber_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 247)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetAccountNumber_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 247)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetAccountNumber_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 247, a)
}

func InfonetAccountNumber_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 247, a)
}

func InfonetAccountNumber_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 247)
}

func InfonetType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 248, a)
}

func InfonetType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 248, a)
}

func InfonetType_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetType_Lookup(p)
	return
}

func InfonetType_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetType_LookupString(p)
	return
}

func InfonetType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 248) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 248) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 248)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 248)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 248, a)
}

func InfonetType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 248, a)
}

func InfonetType_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 248)
}

func InfonetPoolRequest_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 252, a)
}

func InfonetPoolRequest_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 252, a)
}

func InfonetPoolRequest_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetPoolRequest_Lookup(p)
	return
}

func InfonetPoolRequest_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetPoolRequest_LookupString(p)
	return
}

func InfonetPoolRequest_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 252) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetPoolRequest_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 252) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetPoolRequest_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 252)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetPoolRequest_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 252)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetPoolRequest_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 252, a)
}

func InfonetPoolRequest_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 252, a)
}

func InfonetPoolRequest_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 252)
}

type InfonetSurchargeType uint32

var InfonetSurchargeType_Strings = map[InfonetSurchargeType]string{}

func (a InfonetSurchargeType) String() string {
	if str, ok := InfonetSurchargeType_Strings[a]; ok {
		return str
	}
	return "InfonetSurchargeType(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func InfonetSurchargeType_Add(p *radius.Packet, value InfonetSurchargeType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Infonet_AddVendor(p, 254, a)
}

func InfonetSurchargeType_Get(p *radius.Packet) (value InfonetSurchargeType) {
	value, _ = InfonetSurchargeType_Lookup(p)
	return
}

func InfonetSurchargeType_Gets(p *radius.Packet) (values []InfonetSurchargeType, err error) {
	var i uint32
	for _, attr := range _Infonet_GetsVendor(p, 254) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, InfonetSurchargeType(i))
	}
	return
}

func InfonetSurchargeType_Lookup(p *radius.Packet) (value InfonetSurchargeType, err error) {
	a, ok := _Infonet_LookupVendor(p, 254)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = InfonetSurchargeType(i)
	return
}

func InfonetSurchargeType_Set(p *radius.Packet, value InfonetSurchargeType) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Infonet_SetVendor(p, 254, a)
}

func InfonetSurchargeType_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 254)
}

func InfonetNASLocation_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 255, a)
}

func InfonetNASLocation_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 255, a)
}

func InfonetNASLocation_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetNASLocation_Lookup(p)
	return
}

func InfonetNASLocation_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetNASLocation_LookupString(p)
	return
}

func InfonetNASLocation_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 255) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetNASLocation_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 255) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetNASLocation_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 255)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetNASLocation_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 255)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetNASLocation_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 255, a)
}

func InfonetNASLocation_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 255, a)
}

func InfonetNASLocation_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 255)
}

func InfonetRandomIPPool_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 246, a)
}

func InfonetRandomIPPool_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 246, a)
}

func InfonetRandomIPPool_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetRandomIPPool_Lookup(p)
	return
}

func InfonetRandomIPPool_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetRandomIPPool_LookupString(p)
	return
}

func InfonetRandomIPPool_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 246) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetRandomIPPool_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 246) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetRandomIPPool_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 246)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetRandomIPPool_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 246)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetRandomIPPool_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 246, a)
}

func InfonetRandomIPPool_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 246, a)
}

func InfonetRandomIPPool_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 246)
}

func InfonetRealmType_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 249, a)
}

func InfonetRealmType_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 249, a)
}

func InfonetRealmType_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetRealmType_Lookup(p)
	return
}

func InfonetRealmType_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetRealmType_LookupString(p)
	return
}

func InfonetRealmType_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 249) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetRealmType_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 249) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetRealmType_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 249)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetRealmType_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 249)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetRealmType_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 249, a)
}

func InfonetRealmType_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 249, a)
}

func InfonetRealmType_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 249)
}

func InfonetLoginHostDest_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 250, a)
}

func InfonetLoginHostDest_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 250, a)
}

func InfonetLoginHostDest_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetLoginHostDest_Lookup(p)
	return
}

func InfonetLoginHostDest_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetLoginHostDest_LookupString(p)
	return
}

func InfonetLoginHostDest_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 250) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetLoginHostDest_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 250) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetLoginHostDest_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 250)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetLoginHostDest_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 250)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetLoginHostDest_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 250, a)
}

func InfonetLoginHostDest_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 250, a)
}

func InfonetLoginHostDest_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 250)
}

func InfonetTunnelDecisionIP_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 251, a)
}

func InfonetTunnelDecisionIP_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_AddVendor(p, 251, a)
}

func InfonetTunnelDecisionIP_Get(p *radius.Packet) (value []byte) {
	value, _ = InfonetTunnelDecisionIP_Lookup(p)
	return
}

func InfonetTunnelDecisionIP_GetString(p *radius.Packet) (value string) {
	value, _ = InfonetTunnelDecisionIP_LookupString(p)
	return
}

func InfonetTunnelDecisionIP_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Infonet_GetsVendor(p, 251) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetTunnelDecisionIP_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Infonet_GetsVendor(p, 251) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func InfonetTunnelDecisionIP_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Infonet_LookupVendor(p, 251)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func InfonetTunnelDecisionIP_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Infonet_LookupVendor(p, 251)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func InfonetTunnelDecisionIP_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 251, a)
}

func InfonetTunnelDecisionIP_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Infonet_SetVendor(p, 251, a)
}

func InfonetTunnelDecisionIP_Del(p *radius.Packet) {
	_Infonet_DelVendor(p, 251)
}
