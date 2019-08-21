// Code generated by radius-dict-gen. DO NOT EDIT.

package colubris

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Colubris_VendorID = 8744
)

func _Colubris_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Colubris_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Colubris_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Colubris_VendorID {
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

func _Colubris_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Colubris_VendorID {
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

func _Colubris_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Colubris_VendorID {
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
	return _Colubris_AddVendor(p, typ, attr)
}

func _Colubris_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Colubris_VendorID {
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

func ColubrisAVPair_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Colubris_AddVendor(p, 0, a)
}

func ColubrisAVPair_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Colubris_AddVendor(p, 0, a)
}

func ColubrisAVPair_Get(p *radius.Packet) (value []byte) {
	value, _ = ColubrisAVPair_Lookup(p)
	return
}

func ColubrisAVPair_GetString(p *radius.Packet) (value string) {
	value, _ = ColubrisAVPair_LookupString(p)
	return
}

func ColubrisAVPair_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Colubris_GetsVendor(p, 0) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ColubrisAVPair_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Colubris_GetsVendor(p, 0) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ColubrisAVPair_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Colubris_LookupVendor(p, 0)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ColubrisAVPair_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Colubris_LookupVendor(p, 0)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ColubrisAVPair_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Colubris_SetVendor(p, 0, a)
}

func ColubrisAVPair_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Colubris_SetVendor(p, 0, a)
}

func ColubrisAVPair_Del(p *radius.Packet) {
	_Colubris_DelVendor(p, 0)
}

type ColubrisIntercept uint32

var ColubrisIntercept_Strings = map[ColubrisIntercept]string{}

func (a ColubrisIntercept) String() string {
	if str, ok := ColubrisIntercept_Strings[a]; ok {
		return str
	}
	return "ColubrisIntercept(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ColubrisIntercept_Add(p *radius.Packet, value ColubrisIntercept) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Colubris_AddVendor(p, 1, a)
}

func ColubrisIntercept_Get(p *radius.Packet) (value ColubrisIntercept) {
	value, _ = ColubrisIntercept_Lookup(p)
	return
}

func ColubrisIntercept_Gets(p *radius.Packet) (values []ColubrisIntercept, err error) {
	var i uint32
	for _, attr := range _Colubris_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ColubrisIntercept(i))
	}
	return
}

func ColubrisIntercept_Lookup(p *radius.Packet) (value ColubrisIntercept, err error) {
	a, ok := _Colubris_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ColubrisIntercept(i)
	return
}

func ColubrisIntercept_Set(p *radius.Packet, value ColubrisIntercept) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Colubris_SetVendor(p, 1, a)
}

func ColubrisIntercept_Del(p *radius.Packet) {
	_Colubris_DelVendor(p, 1)
}
