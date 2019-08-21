// Code generated by radius-dict-gen. DO NOT EDIT.

package propel

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Propel_VendorID = 14895
)

func _Propel_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Propel_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Propel_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Propel_VendorID {
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

func _Propel_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Propel_VendorID {
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

func _Propel_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Propel_VendorID {
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
	return _Propel_AddVendor(p, typ, attr)
}

func _Propel_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Propel_VendorID {
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

type PropelAccelerate uint32

var PropelAccelerate_Strings = map[PropelAccelerate]string{}

func (a PropelAccelerate) String() string {
	if str, ok := PropelAccelerate_Strings[a]; ok {
		return str
	}
	return "PropelAccelerate(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PropelAccelerate_Add(p *radius.Packet, value PropelAccelerate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_AddVendor(p, 1, a)
}

func PropelAccelerate_Get(p *radius.Packet) (value PropelAccelerate) {
	value, _ = PropelAccelerate_Lookup(p)
	return
}

func PropelAccelerate_Gets(p *radius.Packet) (values []PropelAccelerate, err error) {
	var i uint32
	for _, attr := range _Propel_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PropelAccelerate(i))
	}
	return
}

func PropelAccelerate_Lookup(p *radius.Packet) (value PropelAccelerate, err error) {
	a, ok := _Propel_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PropelAccelerate(i)
	return
}

func PropelAccelerate_Set(p *radius.Packet, value PropelAccelerate) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_SetVendor(p, 1, a)
}

func PropelAccelerate_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 1)
}

func PropelDialedDigits_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Propel_AddVendor(p, 2, a)
}

func PropelDialedDigits_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Propel_AddVendor(p, 2, a)
}

func PropelDialedDigits_Get(p *radius.Packet) (value []byte) {
	value, _ = PropelDialedDigits_Lookup(p)
	return
}

func PropelDialedDigits_GetString(p *radius.Packet) (value string) {
	value, _ = PropelDialedDigits_LookupString(p)
	return
}

func PropelDialedDigits_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Propel_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PropelDialedDigits_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Propel_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PropelDialedDigits_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Propel_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func PropelDialedDigits_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Propel_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func PropelDialedDigits_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Propel_SetVendor(p, 2, a)
}

func PropelDialedDigits_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Propel_SetVendor(p, 2, a)
}

func PropelDialedDigits_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 2)
}

func PropelClientIPAddress_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Propel_AddVendor(p, 3, a)
}

func PropelClientIPAddress_Get(p *radius.Packet) (value net.IP) {
	value, _ = PropelClientIPAddress_Lookup(p)
	return
}

func PropelClientIPAddress_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Propel_GetsVendor(p, 3) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PropelClientIPAddress_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Propel_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func PropelClientIPAddress_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Propel_SetVendor(p, 3, a)
}

func PropelClientIPAddress_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 3)
}

func PropelClientNASIPAddress_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Propel_AddVendor(p, 4, a)
}

func PropelClientNASIPAddress_Get(p *radius.Packet) (value net.IP) {
	value, _ = PropelClientNASIPAddress_Lookup(p)
	return
}

func PropelClientNASIPAddress_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Propel_GetsVendor(p, 4) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func PropelClientNASIPAddress_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Propel_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func PropelClientNASIPAddress_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Propel_SetVendor(p, 4, a)
}

func PropelClientNASIPAddress_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 4)
}

type PropelClientSourceID uint32

var PropelClientSourceID_Strings = map[PropelClientSourceID]string{}

func (a PropelClientSourceID) String() string {
	if str, ok := PropelClientSourceID_Strings[a]; ok {
		return str
	}
	return "PropelClientSourceID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PropelClientSourceID_Add(p *radius.Packet, value PropelClientSourceID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_AddVendor(p, 5, a)
}

func PropelClientSourceID_Get(p *radius.Packet) (value PropelClientSourceID) {
	value, _ = PropelClientSourceID_Lookup(p)
	return
}

func PropelClientSourceID_Gets(p *radius.Packet) (values []PropelClientSourceID, err error) {
	var i uint32
	for _, attr := range _Propel_GetsVendor(p, 5) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PropelClientSourceID(i))
	}
	return
}

func PropelClientSourceID_Lookup(p *radius.Packet) (value PropelClientSourceID, err error) {
	a, ok := _Propel_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PropelClientSourceID(i)
	return
}

func PropelClientSourceID_Set(p *radius.Packet, value PropelClientSourceID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_SetVendor(p, 5, a)
}

func PropelClientSourceID_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 5)
}

type PropelContentFilterID uint32

var PropelContentFilterID_Strings = map[PropelContentFilterID]string{}

func (a PropelContentFilterID) String() string {
	if str, ok := PropelContentFilterID_Strings[a]; ok {
		return str
	}
	return "PropelContentFilterID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PropelContentFilterID_Add(p *radius.Packet, value PropelContentFilterID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_AddVendor(p, 6, a)
}

func PropelContentFilterID_Get(p *radius.Packet) (value PropelContentFilterID) {
	value, _ = PropelContentFilterID_Lookup(p)
	return
}

func PropelContentFilterID_Gets(p *radius.Packet) (values []PropelContentFilterID, err error) {
	var i uint32
	for _, attr := range _Propel_GetsVendor(p, 6) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PropelContentFilterID(i))
	}
	return
}

func PropelContentFilterID_Lookup(p *radius.Packet) (value PropelContentFilterID, err error) {
	a, ok := _Propel_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PropelContentFilterID(i)
	return
}

func PropelContentFilterID_Set(p *radius.Packet, value PropelContentFilterID) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Propel_SetVendor(p, 6, a)
}

func PropelContentFilterID_Del(p *radius.Packet) {
	_Propel_DelVendor(p, 6)
}
