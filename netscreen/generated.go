// Code generated by radius-dict-gen. DO NOT EDIT.

package netscreen

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Netscreen_VendorID = 3224
)

func _Netscreen_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Netscreen_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Netscreen_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Netscreen_VendorID {
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

func _Netscreen_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Netscreen_VendorID {
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

func _Netscreen_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Netscreen_VendorID {
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
	return _Netscreen_AddVendor(p, typ, attr)
}

func _Netscreen_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Netscreen_VendorID {
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

type NSAdminPrivilege uint32

const (
	NSAdminPrivilege_Value_RootAdmin         NSAdminPrivilege = 1
	NSAdminPrivilege_Value_AllVSYSRootAdmin  NSAdminPrivilege = 2
	NSAdminPrivilege_Value_VSYSAdmin         NSAdminPrivilege = 3
	NSAdminPrivilege_Value_ReadOnlyAdmin     NSAdminPrivilege = 4
	NSAdminPrivilege_Value_ReadOnlyVSYSAdmin NSAdminPrivilege = 5
)

var NSAdminPrivilege_Strings = map[NSAdminPrivilege]string{
	NSAdminPrivilege_Value_RootAdmin:         "Root-Admin",
	NSAdminPrivilege_Value_AllVSYSRootAdmin:  "All-VSYS-Root-Admin",
	NSAdminPrivilege_Value_VSYSAdmin:         "VSYS-Admin",
	NSAdminPrivilege_Value_ReadOnlyAdmin:     "Read-Only-Admin",
	NSAdminPrivilege_Value_ReadOnlyVSYSAdmin: "Read-Only-VSYS-Admin",
}

func (a NSAdminPrivilege) String() string {
	if str, ok := NSAdminPrivilege_Strings[a]; ok {
		return str
	}
	return "NSAdminPrivilege(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func NSAdminPrivilege_Add(p *radius.Packet, value NSAdminPrivilege) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Netscreen_AddVendor(p, 1, a)
}

func NSAdminPrivilege_Get(p *radius.Packet) (value NSAdminPrivilege) {
	value, _ = NSAdminPrivilege_Lookup(p)
	return
}

func NSAdminPrivilege_Gets(p *radius.Packet) (values []NSAdminPrivilege, err error) {
	var i uint32
	for _, attr := range _Netscreen_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, NSAdminPrivilege(i))
	}
	return
}

func NSAdminPrivilege_Lookup(p *radius.Packet) (value NSAdminPrivilege, err error) {
	a, ok := _Netscreen_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = NSAdminPrivilege(i)
	return
}

func NSAdminPrivilege_Set(p *radius.Packet, value NSAdminPrivilege) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Netscreen_SetVendor(p, 1, a)
}

func NSAdminPrivilege_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 1)
}

func NSVSYSName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 2, a)
}

func NSVSYSName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 2, a)
}

func NSVSYSName_Get(p *radius.Packet) (value []byte) {
	value, _ = NSVSYSName_Lookup(p)
	return
}

func NSVSYSName_GetString(p *radius.Packet) (value string) {
	value, _ = NSVSYSName_LookupString(p)
	return
}

func NSVSYSName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Netscreen_GetsVendor(p, 2) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSVSYSName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Netscreen_GetsVendor(p, 2) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSVSYSName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Netscreen_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NSVSYSName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Netscreen_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NSVSYSName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 2, a)
}

func NSVSYSName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 2, a)
}

func NSVSYSName_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 2)
}

func NSUserGroup_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 3, a)
}

func NSUserGroup_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 3, a)
}

func NSUserGroup_Get(p *radius.Packet) (value []byte) {
	value, _ = NSUserGroup_Lookup(p)
	return
}

func NSUserGroup_GetString(p *radius.Packet) (value string) {
	value, _ = NSUserGroup_LookupString(p)
	return
}

func NSUserGroup_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Netscreen_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSUserGroup_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Netscreen_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSUserGroup_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Netscreen_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NSUserGroup_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Netscreen_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NSUserGroup_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 3, a)
}

func NSUserGroup_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 3, a)
}

func NSUserGroup_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 3)
}

func NSPrimaryDNS_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 4, a)
}

func NSPrimaryDNS_Get(p *radius.Packet) (value net.IP) {
	value, _ = NSPrimaryDNS_Lookup(p)
	return
}

func NSPrimaryDNS_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Netscreen_GetsVendor(p, 4) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSPrimaryDNS_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Netscreen_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func NSPrimaryDNS_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 4, a)
}

func NSPrimaryDNS_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 4)
}

func NSSecondaryDNS_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 5, a)
}

func NSSecondaryDNS_Get(p *radius.Packet) (value net.IP) {
	value, _ = NSSecondaryDNS_Lookup(p)
	return
}

func NSSecondaryDNS_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Netscreen_GetsVendor(p, 5) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSSecondaryDNS_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Netscreen_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func NSSecondaryDNS_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 5, a)
}

func NSSecondaryDNS_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 5)
}

func NSPrimaryWINS_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 6, a)
}

func NSPrimaryWINS_Get(p *radius.Packet) (value net.IP) {
	value, _ = NSPrimaryWINS_Lookup(p)
	return
}

func NSPrimaryWINS_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Netscreen_GetsVendor(p, 6) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSPrimaryWINS_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Netscreen_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func NSPrimaryWINS_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 6, a)
}

func NSPrimaryWINS_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 6)
}

func NSSecondaryWINS_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 7, a)
}

func NSSecondaryWINS_Get(p *radius.Packet) (value net.IP) {
	value, _ = NSSecondaryWINS_Lookup(p)
	return
}

func NSSecondaryWINS_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Netscreen_GetsVendor(p, 7) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSSecondaryWINS_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Netscreen_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func NSSecondaryWINS_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 7, a)
}

func NSSecondaryWINS_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 7)
}

func NSNSMUserDomainName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 220, a)
}

func NSNSMUserDomainName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 220, a)
}

func NSNSMUserDomainName_Get(p *radius.Packet) (value []byte) {
	value, _ = NSNSMUserDomainName_Lookup(p)
	return
}

func NSNSMUserDomainName_GetString(p *radius.Packet) (value string) {
	value, _ = NSNSMUserDomainName_LookupString(p)
	return
}

func NSNSMUserDomainName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Netscreen_GetsVendor(p, 220) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSNSMUserDomainName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Netscreen_GetsVendor(p, 220) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSNSMUserDomainName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Netscreen_LookupVendor(p, 220)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NSNSMUserDomainName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Netscreen_LookupVendor(p, 220)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NSNSMUserDomainName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 220, a)
}

func NSNSMUserDomainName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 220, a)
}

func NSNSMUserDomainName_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 220)
}

func NSNSMUserRoleMapping_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 221, a)
}

func NSNSMUserRoleMapping_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_AddVendor(p, 221, a)
}

func NSNSMUserRoleMapping_Get(p *radius.Packet) (value []byte) {
	value, _ = NSNSMUserRoleMapping_Lookup(p)
	return
}

func NSNSMUserRoleMapping_GetString(p *radius.Packet) (value string) {
	value, _ = NSNSMUserRoleMapping_LookupString(p)
	return
}

func NSNSMUserRoleMapping_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Netscreen_GetsVendor(p, 221) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSNSMUserRoleMapping_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Netscreen_GetsVendor(p, 221) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NSNSMUserRoleMapping_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Netscreen_LookupVendor(p, 221)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NSNSMUserRoleMapping_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Netscreen_LookupVendor(p, 221)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NSNSMUserRoleMapping_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 221, a)
}

func NSNSMUserRoleMapping_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Netscreen_SetVendor(p, 221, a)
}

func NSNSMUserRoleMapping_Del(p *radius.Packet) {
	_Netscreen_DelVendor(p, 221)
}
