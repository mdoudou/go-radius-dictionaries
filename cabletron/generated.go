// Code generated by radius-dict-gen. DO NOT EDIT.

package cabletron

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Cabletron_VendorID = 52
)

func _Cabletron_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Cabletron_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Cabletron_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Cabletron_VendorID {
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

func _Cabletron_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Cabletron_VendorID {
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

func _Cabletron_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Cabletron_VendorID {
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
	return _Cabletron_AddVendor(p, typ, attr)
}

func _Cabletron_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Cabletron_VendorID {
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

type CabletronProtocolEnable uint32

const (
	CabletronProtocolEnable_Value_IPEnable      CabletronProtocolEnable = 1
	CabletronProtocolEnable_Value_BridgeEnable  CabletronProtocolEnable = 2
	CabletronProtocolEnable_Value_IPBREnable    CabletronProtocolEnable = 3
	CabletronProtocolEnable_Value_BRIPXEnable   CabletronProtocolEnable = 6
	CabletronProtocolEnable_Value_IPBRIPXEnable CabletronProtocolEnable = 7
)

var CabletronProtocolEnable_Strings = map[CabletronProtocolEnable]string{
	CabletronProtocolEnable_Value_IPEnable:      "IP-Enable",
	CabletronProtocolEnable_Value_BridgeEnable:  "Bridge-Enable",
	CabletronProtocolEnable_Value_IPBREnable:    "IP-BR-Enable",
	CabletronProtocolEnable_Value_BRIPXEnable:   "BR-IPX-Enable",
	CabletronProtocolEnable_Value_IPBRIPXEnable: "IP-BR-IPX-Enable",
}

func (a CabletronProtocolEnable) String() string {
	if str, ok := CabletronProtocolEnable_Strings[a]; ok {
		return str
	}
	return "CabletronProtocolEnable(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CabletronProtocolEnable_Add(p *radius.Packet, value CabletronProtocolEnable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Cabletron_AddVendor(p, 201, a)
}

func CabletronProtocolEnable_Get(p *radius.Packet) (value CabletronProtocolEnable) {
	value, _ = CabletronProtocolEnable_Lookup(p)
	return
}

func CabletronProtocolEnable_Gets(p *radius.Packet) (values []CabletronProtocolEnable, err error) {
	var i uint32
	for _, attr := range _Cabletron_GetsVendor(p, 201) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CabletronProtocolEnable(i))
	}
	return
}

func CabletronProtocolEnable_Lookup(p *radius.Packet) (value CabletronProtocolEnable, err error) {
	a, ok := _Cabletron_LookupVendor(p, 201)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CabletronProtocolEnable(i)
	return
}

func CabletronProtocolEnable_Set(p *radius.Packet, value CabletronProtocolEnable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Cabletron_SetVendor(p, 201, a)
}

func CabletronProtocolEnable_Del(p *radius.Packet) {
	_Cabletron_DelVendor(p, 201)
}

type CabletronProtocolCallable uint32

const (
	CabletronProtocolCallable_Value_IPCallable      CabletronProtocolCallable = 1
	CabletronProtocolCallable_Value_BridgeCallable  CabletronProtocolCallable = 2
	CabletronProtocolCallable_Value_IPBRCallable    CabletronProtocolCallable = 3
	CabletronProtocolCallable_Value_BRIPXCallable   CabletronProtocolCallable = 6
	CabletronProtocolCallable_Value_IPBRIPXCallable CabletronProtocolCallable = 7
)

var CabletronProtocolCallable_Strings = map[CabletronProtocolCallable]string{
	CabletronProtocolCallable_Value_IPCallable:      "IP-Callable",
	CabletronProtocolCallable_Value_BridgeCallable:  "Bridge-Callable",
	CabletronProtocolCallable_Value_IPBRCallable:    "IP-BR-Callable",
	CabletronProtocolCallable_Value_BRIPXCallable:   "BR-IPX-Callable",
	CabletronProtocolCallable_Value_IPBRIPXCallable: "IP-BR-IPX-Callable",
}

func (a CabletronProtocolCallable) String() string {
	if str, ok := CabletronProtocolCallable_Strings[a]; ok {
		return str
	}
	return "CabletronProtocolCallable(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CabletronProtocolCallable_Add(p *radius.Packet, value CabletronProtocolCallable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Cabletron_AddVendor(p, 202, a)
}

func CabletronProtocolCallable_Get(p *radius.Packet) (value CabletronProtocolCallable) {
	value, _ = CabletronProtocolCallable_Lookup(p)
	return
}

func CabletronProtocolCallable_Gets(p *radius.Packet) (values []CabletronProtocolCallable, err error) {
	var i uint32
	for _, attr := range _Cabletron_GetsVendor(p, 202) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CabletronProtocolCallable(i))
	}
	return
}

func CabletronProtocolCallable_Lookup(p *radius.Packet) (value CabletronProtocolCallable, err error) {
	a, ok := _Cabletron_LookupVendor(p, 202)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CabletronProtocolCallable(i)
	return
}

func CabletronProtocolCallable_Set(p *radius.Packet, value CabletronProtocolCallable) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Cabletron_SetVendor(p, 202, a)
}

func CabletronProtocolCallable_Del(p *radius.Packet) {
	_Cabletron_DelVendor(p, 202)
}
