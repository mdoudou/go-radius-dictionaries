// Code generated by radius-dict-gen. DO NOT EDIT.

package audiocodes

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_AudioCodes_VendorID = 5003
)

func _AudioCodes_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_AudioCodes_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _AudioCodes_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _AudioCodes_VendorID {
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

func _AudioCodes_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _AudioCodes_VendorID {
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

func _AudioCodes_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _AudioCodes_VendorID {
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
	return _AudioCodes_AddVendor(p, typ, attr)
}

func _AudioCodes_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _AudioCodes_VendorID {
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

type ACLAuthLevel uint32

const (
	ACLAuthLevel_Value_ACLAuthUserLevel          ACLAuthLevel = 50
	ACLAuthLevel_Value_ACLAuthAdminLevel         ACLAuthLevel = 100
	ACLAuthLevel_Value_ACLAuthSecurityAdminLevel ACLAuthLevel = 200
)

var ACLAuthLevel_Strings = map[ACLAuthLevel]string{
	ACLAuthLevel_Value_ACLAuthUserLevel:          "ACL-Auth-UserLevel",
	ACLAuthLevel_Value_ACLAuthAdminLevel:         "ACL-Auth-AdminLevel",
	ACLAuthLevel_Value_ACLAuthSecurityAdminLevel: "ACL-Auth-SecurityAdminLevel",
}

func (a ACLAuthLevel) String() string {
	if str, ok := ACLAuthLevel_Strings[a]; ok {
		return str
	}
	return "ACLAuthLevel(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ACLAuthLevel_Add(p *radius.Packet, value ACLAuthLevel) (err error) {
	a := radius.NewInteger(uint32(value))
	return _AudioCodes_AddVendor(p, 35, a)
}

func ACLAuthLevel_Get(p *radius.Packet) (value ACLAuthLevel) {
	value, _ = ACLAuthLevel_Lookup(p)
	return
}

func ACLAuthLevel_Gets(p *radius.Packet) (values []ACLAuthLevel, err error) {
	var i uint32
	for _, attr := range _AudioCodes_GetsVendor(p, 35) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ACLAuthLevel(i))
	}
	return
}

func ACLAuthLevel_Lookup(p *radius.Packet) (value ACLAuthLevel, err error) {
	a, ok := _AudioCodes_LookupVendor(p, 35)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ACLAuthLevel(i)
	return
}

func ACLAuthLevel_Set(p *radius.Packet, value ACLAuthLevel) (err error) {
	a := radius.NewInteger(uint32(value))
	return _AudioCodes_SetVendor(p, 35, a)
}

func ACLAuthLevel_Del(p *radius.Packet) {
	_AudioCodes_DelVendor(p, 35)
}