// Code generated by radius-dict-gen. DO NOT EDIT.

package xedia

import (
	"net"
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Xedia_VendorID = 838
)

func _Xedia_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Xedia_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Xedia_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Xedia_VendorID {
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

func _Xedia_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _Xedia_VendorID {
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

func _Xedia_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _Xedia_VendorID {
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
	return _Xedia_AddVendor(p, typ, attr)
}

func _Xedia_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Xedia_VendorID {
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

func XediaDNSServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 1, a)
}

func XediaDNSServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = XediaDNSServer_Lookup(p)
	return
}

func XediaDNSServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Xedia_GetsVendor(p, 1) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaDNSServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Xedia_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func XediaDNSServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 1, a)
}

func XediaDNSServer_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 1)
}

func XediaNetBiosServer_Add(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 2, a)
}

func XediaNetBiosServer_Get(p *radius.Packet) (value net.IP) {
	value, _ = XediaNetBiosServer_Lookup(p)
	return
}

func XediaNetBiosServer_Gets(p *radius.Packet) (values []net.IP, err error) {
	var i net.IP
	for _, attr := range _Xedia_GetsVendor(p, 2) {
		i, err = radius.IPAddr(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaNetBiosServer_Lookup(p *radius.Packet) (value net.IP, err error) {
	a, ok := _Xedia_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.IPAddr(a)
	return
}

func XediaNetBiosServer_Set(p *radius.Packet, value net.IP) (err error) {
	var a radius.Attribute
	a, err = radius.NewIPAddr(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 2, a)
}

func XediaNetBiosServer_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 2)
}

func XediaAddressPool_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 3, a)
}

func XediaAddressPool_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 3, a)
}

func XediaAddressPool_Get(p *radius.Packet) (value []byte) {
	value, _ = XediaAddressPool_Lookup(p)
	return
}

func XediaAddressPool_GetString(p *radius.Packet) (value string) {
	value, _ = XediaAddressPool_LookupString(p)
	return
}

func XediaAddressPool_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Xedia_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaAddressPool_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Xedia_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaAddressPool_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Xedia_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func XediaAddressPool_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Xedia_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func XediaAddressPool_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 3, a)
}

func XediaAddressPool_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 3, a)
}

func XediaAddressPool_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 3)
}

type XediaPPPEchoInterval uint32

var XediaPPPEchoInterval_Strings = map[XediaPPPEchoInterval]string{}

func (a XediaPPPEchoInterval) String() string {
	if str, ok := XediaPPPEchoInterval_Strings[a]; ok {
		return str
	}
	return "XediaPPPEchoInterval(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func XediaPPPEchoInterval_Add(p *radius.Packet, value XediaPPPEchoInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_AddVendor(p, 4, a)
}

func XediaPPPEchoInterval_Get(p *radius.Packet) (value XediaPPPEchoInterval) {
	value, _ = XediaPPPEchoInterval_Lookup(p)
	return
}

func XediaPPPEchoInterval_Gets(p *radius.Packet) (values []XediaPPPEchoInterval, err error) {
	var i uint32
	for _, attr := range _Xedia_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, XediaPPPEchoInterval(i))
	}
	return
}

func XediaPPPEchoInterval_Lookup(p *radius.Packet) (value XediaPPPEchoInterval, err error) {
	a, ok := _Xedia_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = XediaPPPEchoInterval(i)
	return
}

func XediaPPPEchoInterval_Set(p *radius.Packet, value XediaPPPEchoInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_SetVendor(p, 4, a)
}

func XediaPPPEchoInterval_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 4)
}

type XediaSSHPrivileges uint32

var XediaSSHPrivileges_Strings = map[XediaSSHPrivileges]string{}

func (a XediaSSHPrivileges) String() string {
	if str, ok := XediaSSHPrivileges_Strings[a]; ok {
		return str
	}
	return "XediaSSHPrivileges(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func XediaSSHPrivileges_Add(p *radius.Packet, value XediaSSHPrivileges) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_AddVendor(p, 5, a)
}

func XediaSSHPrivileges_Get(p *radius.Packet) (value XediaSSHPrivileges) {
	value, _ = XediaSSHPrivileges_Lookup(p)
	return
}

func XediaSSHPrivileges_Gets(p *radius.Packet) (values []XediaSSHPrivileges, err error) {
	var i uint32
	for _, attr := range _Xedia_GetsVendor(p, 5) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, XediaSSHPrivileges(i))
	}
	return
}

func XediaSSHPrivileges_Lookup(p *radius.Packet) (value XediaSSHPrivileges, err error) {
	a, ok := _Xedia_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = XediaSSHPrivileges(i)
	return
}

func XediaSSHPrivileges_Set(p *radius.Packet, value XediaSSHPrivileges) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_SetVendor(p, 5, a)
}

func XediaSSHPrivileges_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 5)
}

func XediaClientAccessNetwork_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 6, a)
}

func XediaClientAccessNetwork_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Xedia_AddVendor(p, 6, a)
}

func XediaClientAccessNetwork_Get(p *radius.Packet) (value []byte) {
	value, _ = XediaClientAccessNetwork_Lookup(p)
	return
}

func XediaClientAccessNetwork_GetString(p *radius.Packet) (value string) {
	value, _ = XediaClientAccessNetwork_LookupString(p)
	return
}

func XediaClientAccessNetwork_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Xedia_GetsVendor(p, 6) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaClientAccessNetwork_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Xedia_GetsVendor(p, 6) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func XediaClientAccessNetwork_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Xedia_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func XediaClientAccessNetwork_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Xedia_LookupVendor(p, 6)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func XediaClientAccessNetwork_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 6, a)
}

func XediaClientAccessNetwork_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Xedia_SetVendor(p, 6, a)
}

func XediaClientAccessNetwork_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 6)
}

type XediaClientFirewallSetting uint32

var XediaClientFirewallSetting_Strings = map[XediaClientFirewallSetting]string{}

func (a XediaClientFirewallSetting) String() string {
	if str, ok := XediaClientFirewallSetting_Strings[a]; ok {
		return str
	}
	return "XediaClientFirewallSetting(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func XediaClientFirewallSetting_Add(p *radius.Packet, value XediaClientFirewallSetting) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_AddVendor(p, 7, a)
}

func XediaClientFirewallSetting_Get(p *radius.Packet) (value XediaClientFirewallSetting) {
	value, _ = XediaClientFirewallSetting_Lookup(p)
	return
}

func XediaClientFirewallSetting_Gets(p *radius.Packet) (values []XediaClientFirewallSetting, err error) {
	var i uint32
	for _, attr := range _Xedia_GetsVendor(p, 7) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, XediaClientFirewallSetting(i))
	}
	return
}

func XediaClientFirewallSetting_Lookup(p *radius.Packet) (value XediaClientFirewallSetting, err error) {
	a, ok := _Xedia_LookupVendor(p, 7)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = XediaClientFirewallSetting(i)
	return
}

func XediaClientFirewallSetting_Set(p *radius.Packet, value XediaClientFirewallSetting) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_SetVendor(p, 7, a)
}

func XediaClientFirewallSetting_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 7)
}

type XediaSavePassword uint32

var XediaSavePassword_Strings = map[XediaSavePassword]string{}

func (a XediaSavePassword) String() string {
	if str, ok := XediaSavePassword_Strings[a]; ok {
		return str
	}
	return "XediaSavePassword(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func XediaSavePassword_Add(p *radius.Packet, value XediaSavePassword) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_AddVendor(p, 8, a)
}

func XediaSavePassword_Get(p *radius.Packet) (value XediaSavePassword) {
	value, _ = XediaSavePassword_Lookup(p)
	return
}

func XediaSavePassword_Gets(p *radius.Packet) (values []XediaSavePassword, err error) {
	var i uint32
	for _, attr := range _Xedia_GetsVendor(p, 8) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, XediaSavePassword(i))
	}
	return
}

func XediaSavePassword_Lookup(p *radius.Packet) (value XediaSavePassword, err error) {
	a, ok := _Xedia_LookupVendor(p, 8)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = XediaSavePassword(i)
	return
}

func XediaSavePassword_Set(p *radius.Packet, value XediaSavePassword) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Xedia_SetVendor(p, 8, a)
}

func XediaSavePassword_Del(p *radius.Packet) {
	_Xedia_DelVendor(p, 8)
}
