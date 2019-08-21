// Code generated by radius-dict-gen. DO NOT EDIT.

package f5

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_F5_VendorID = 3375
)

func _F5_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_F5_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _F5_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, attr := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _F5_VendorID {
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

func _F5_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, a := range p.Attributes[rfc2865.VendorSpecific_Type] {
		vendorID, vsa, err := radius.VendorSpecific(a)
		if err != nil || vendorID != _F5_VendorID {
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

func _F5_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		vendorID, vsa, err := radius.VendorSpecific(p.Attributes[rfc2865.VendorSpecific_Type][i])
		if err != nil || vendorID != _F5_VendorID {
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
	return _F5_AddVendor(p, typ, attr)
}

func _F5_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes[rfc2865.VendorSpecific_Type]); {
		attr := p.Attributes[rfc2865.VendorSpecific_Type][i]
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _F5_VendorID {
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

type F5LTMUserRole uint32

const (
	F5LTMUserRole_Value_Administrator                       F5LTMUserRole = 0
	F5LTMUserRole_Value_ResourceAdmin                       F5LTMUserRole = 20
	F5LTMUserRole_Value_UserManager                         F5LTMUserRole = 40
	F5LTMUserRole_Value_Auditor                             F5LTMUserRole = 80
	F5LTMUserRole_Value_Manager                             F5LTMUserRole = 100
	F5LTMUserRole_Value_AppEditor                           F5LTMUserRole = 300
	F5LTMUserRole_Value_AdvancedOperator                    F5LTMUserRole = 350
	F5LTMUserRole_Value_Operator                            F5LTMUserRole = 400
	F5LTMUserRole_Value_FirewallManager                     F5LTMUserRole = 450
	F5LTMUserRole_Value_FraudProtectionManager              F5LTMUserRole = 480
	F5LTMUserRole_Value_CertificateManager                  F5LTMUserRole = 500
	F5LTMUserRole_Value_IRuleManager                        F5LTMUserRole = 510
	F5LTMUserRole_Value_Guest                               F5LTMUserRole = 700
	F5LTMUserRole_Value_WebApplicationSecurityAdministrator F5LTMUserRole = 800
	F5LTMUserRole_Value_WebApplicationSecurityEditor        F5LTMUserRole = 810
	F5LTMUserRole_Value_AccelerationPolicyEditor            F5LTMUserRole = 850
	F5LTMUserRole_Value_NoAccess                            F5LTMUserRole = 900
)

var F5LTMUserRole_Strings = map[F5LTMUserRole]string{
	F5LTMUserRole_Value_Administrator:                       "Administrator",
	F5LTMUserRole_Value_ResourceAdmin:                       "Resource-Admin",
	F5LTMUserRole_Value_UserManager:                         "User-Manager",
	F5LTMUserRole_Value_Auditor:                             "Auditor",
	F5LTMUserRole_Value_Manager:                             "Manager",
	F5LTMUserRole_Value_AppEditor:                           "App-Editor",
	F5LTMUserRole_Value_AdvancedOperator:                    "Advanced-Operator",
	F5LTMUserRole_Value_Operator:                            "Operator",
	F5LTMUserRole_Value_FirewallManager:                     "Firewall-Manager",
	F5LTMUserRole_Value_FraudProtectionManager:              "Fraud-Protection-Manager",
	F5LTMUserRole_Value_CertificateManager:                  "Certificate-Manager",
	F5LTMUserRole_Value_IRuleManager:                        "IRule-Manager",
	F5LTMUserRole_Value_Guest:                               "Guest",
	F5LTMUserRole_Value_WebApplicationSecurityAdministrator: "Web-Application-Security-Administrator",
	F5LTMUserRole_Value_WebApplicationSecurityEditor:        "Web-Application-Security-Editor",
	F5LTMUserRole_Value_AccelerationPolicyEditor:            "Acceleration-Policy-Editor",
	F5LTMUserRole_Value_NoAccess:                            "No-Access",
}

func (a F5LTMUserRole) String() string {
	if str, ok := F5LTMUserRole_Strings[a]; ok {
		return str
	}
	return "F5LTMUserRole(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func F5LTMUserRole_Add(p *radius.Packet, value F5LTMUserRole) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_AddVendor(p, 1, a)
}

func F5LTMUserRole_Get(p *radius.Packet) (value F5LTMUserRole) {
	value, _ = F5LTMUserRole_Lookup(p)
	return
}

func F5LTMUserRole_Gets(p *radius.Packet) (values []F5LTMUserRole, err error) {
	var i uint32
	for _, attr := range _F5_GetsVendor(p, 1) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, F5LTMUserRole(i))
	}
	return
}

func F5LTMUserRole_Lookup(p *radius.Packet) (value F5LTMUserRole, err error) {
	a, ok := _F5_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = F5LTMUserRole(i)
	return
}

func F5LTMUserRole_Set(p *radius.Packet, value F5LTMUserRole) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_SetVendor(p, 1, a)
}

func F5LTMUserRole_Del(p *radius.Packet) {
	_F5_DelVendor(p, 1)
}

type F5LTMUserRoleUniversal uint32

const (
	F5LTMUserRoleUniversal_Value_Disabled F5LTMUserRoleUniversal = 0
	F5LTMUserRoleUniversal_Value_Enabled  F5LTMUserRoleUniversal = 1
)

var F5LTMUserRoleUniversal_Strings = map[F5LTMUserRoleUniversal]string{
	F5LTMUserRoleUniversal_Value_Disabled: "Disabled",
	F5LTMUserRoleUniversal_Value_Enabled:  "Enabled",
}

func (a F5LTMUserRoleUniversal) String() string {
	if str, ok := F5LTMUserRoleUniversal_Strings[a]; ok {
		return str
	}
	return "F5LTMUserRoleUniversal(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func F5LTMUserRoleUniversal_Add(p *radius.Packet, value F5LTMUserRoleUniversal) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_AddVendor(p, 2, a)
}

func F5LTMUserRoleUniversal_Get(p *radius.Packet) (value F5LTMUserRoleUniversal) {
	value, _ = F5LTMUserRoleUniversal_Lookup(p)
	return
}

func F5LTMUserRoleUniversal_Gets(p *radius.Packet) (values []F5LTMUserRoleUniversal, err error) {
	var i uint32
	for _, attr := range _F5_GetsVendor(p, 2) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, F5LTMUserRoleUniversal(i))
	}
	return
}

func F5LTMUserRoleUniversal_Lookup(p *radius.Packet) (value F5LTMUserRoleUniversal, err error) {
	a, ok := _F5_LookupVendor(p, 2)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = F5LTMUserRoleUniversal(i)
	return
}

func F5LTMUserRoleUniversal_Set(p *radius.Packet, value F5LTMUserRoleUniversal) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_SetVendor(p, 2, a)
}

func F5LTMUserRoleUniversal_Del(p *radius.Packet) {
	_F5_DelVendor(p, 2)
}

func F5LTMUserPartition_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 3, a)
}

func F5LTMUserPartition_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 3, a)
}

func F5LTMUserPartition_Get(p *radius.Packet) (value []byte) {
	value, _ = F5LTMUserPartition_Lookup(p)
	return
}

func F5LTMUserPartition_GetString(p *radius.Packet) (value string) {
	value, _ = F5LTMUserPartition_LookupString(p)
	return
}

func F5LTMUserPartition_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _F5_GetsVendor(p, 3) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserPartition_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _F5_GetsVendor(p, 3) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserPartition_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _F5_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func F5LTMUserPartition_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _F5_LookupVendor(p, 3)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func F5LTMUserPartition_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 3, a)
}

func F5LTMUserPartition_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 3, a)
}

func F5LTMUserPartition_Del(p *radius.Packet) {
	_F5_DelVendor(p, 3)
}

type F5LTMUserConsole uint32

const (
	F5LTMUserConsole_Value_Disabled F5LTMUserConsole = 0
	F5LTMUserConsole_Value_Enabled  F5LTMUserConsole = 1
)

var F5LTMUserConsole_Strings = map[F5LTMUserConsole]string{
	F5LTMUserConsole_Value_Disabled: "Disabled",
	F5LTMUserConsole_Value_Enabled:  "Enabled",
}

func (a F5LTMUserConsole) String() string {
	if str, ok := F5LTMUserConsole_Strings[a]; ok {
		return str
	}
	return "F5LTMUserConsole(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func F5LTMUserConsole_Add(p *radius.Packet, value F5LTMUserConsole) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_AddVendor(p, 4, a)
}

func F5LTMUserConsole_Get(p *radius.Packet) (value F5LTMUserConsole) {
	value, _ = F5LTMUserConsole_Lookup(p)
	return
}

func F5LTMUserConsole_Gets(p *radius.Packet) (values []F5LTMUserConsole, err error) {
	var i uint32
	for _, attr := range _F5_GetsVendor(p, 4) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, F5LTMUserConsole(i))
	}
	return
}

func F5LTMUserConsole_Lookup(p *radius.Packet) (value F5LTMUserConsole, err error) {
	a, ok := _F5_LookupVendor(p, 4)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = F5LTMUserConsole(i)
	return
}

func F5LTMUserConsole_Set(p *radius.Packet, value F5LTMUserConsole) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_SetVendor(p, 4, a)
}

func F5LTMUserConsole_Del(p *radius.Packet) {
	_F5_DelVendor(p, 4)
}

func F5LTMUserShell_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 5, a)
}

func F5LTMUserShell_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 5, a)
}

func F5LTMUserShell_Get(p *radius.Packet) (value []byte) {
	value, _ = F5LTMUserShell_Lookup(p)
	return
}

func F5LTMUserShell_GetString(p *radius.Packet) (value string) {
	value, _ = F5LTMUserShell_LookupString(p)
	return
}

func F5LTMUserShell_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _F5_GetsVendor(p, 5) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserShell_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _F5_GetsVendor(p, 5) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserShell_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _F5_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func F5LTMUserShell_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _F5_LookupVendor(p, 5)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func F5LTMUserShell_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 5, a)
}

func F5LTMUserShell_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 5, a)
}

func F5LTMUserShell_Del(p *radius.Packet) {
	_F5_DelVendor(p, 5)
}

type F5LTMUserContext1 uint32

var F5LTMUserContext1_Strings = map[F5LTMUserContext1]string{}

func (a F5LTMUserContext1) String() string {
	if str, ok := F5LTMUserContext1_Strings[a]; ok {
		return str
	}
	return "F5LTMUserContext1(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func F5LTMUserContext1_Add(p *radius.Packet, value F5LTMUserContext1) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_AddVendor(p, 10, a)
}

func F5LTMUserContext1_Get(p *radius.Packet) (value F5LTMUserContext1) {
	value, _ = F5LTMUserContext1_Lookup(p)
	return
}

func F5LTMUserContext1_Gets(p *radius.Packet) (values []F5LTMUserContext1, err error) {
	var i uint32
	for _, attr := range _F5_GetsVendor(p, 10) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, F5LTMUserContext1(i))
	}
	return
}

func F5LTMUserContext1_Lookup(p *radius.Packet) (value F5LTMUserContext1, err error) {
	a, ok := _F5_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = F5LTMUserContext1(i)
	return
}

func F5LTMUserContext1_Set(p *radius.Packet, value F5LTMUserContext1) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_SetVendor(p, 10, a)
}

func F5LTMUserContext1_Del(p *radius.Packet) {
	_F5_DelVendor(p, 10)
}

type F5LTMUserContext2 uint32

var F5LTMUserContext2_Strings = map[F5LTMUserContext2]string{}

func (a F5LTMUserContext2) String() string {
	if str, ok := F5LTMUserContext2_Strings[a]; ok {
		return str
	}
	return "F5LTMUserContext2(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func F5LTMUserContext2_Add(p *radius.Packet, value F5LTMUserContext2) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_AddVendor(p, 11, a)
}

func F5LTMUserContext2_Get(p *radius.Packet) (value F5LTMUserContext2) {
	value, _ = F5LTMUserContext2_Lookup(p)
	return
}

func F5LTMUserContext2_Gets(p *radius.Packet) (values []F5LTMUserContext2, err error) {
	var i uint32
	for _, attr := range _F5_GetsVendor(p, 11) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, F5LTMUserContext2(i))
	}
	return
}

func F5LTMUserContext2_Lookup(p *radius.Packet) (value F5LTMUserContext2, err error) {
	a, ok := _F5_LookupVendor(p, 11)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = F5LTMUserContext2(i)
	return
}

func F5LTMUserContext2_Set(p *radius.Packet, value F5LTMUserContext2) (err error) {
	a := radius.NewInteger(uint32(value))
	return _F5_SetVendor(p, 11, a)
}

func F5LTMUserContext2_Del(p *radius.Packet) {
	_F5_DelVendor(p, 11)
}

func F5LTMUserInfo1_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 12, a)
}

func F5LTMUserInfo1_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 12, a)
}

func F5LTMUserInfo1_Get(p *radius.Packet) (value []byte) {
	value, _ = F5LTMUserInfo1_Lookup(p)
	return
}

func F5LTMUserInfo1_GetString(p *radius.Packet) (value string) {
	value, _ = F5LTMUserInfo1_LookupString(p)
	return
}

func F5LTMUserInfo1_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _F5_GetsVendor(p, 12) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserInfo1_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _F5_GetsVendor(p, 12) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserInfo1_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _F5_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func F5LTMUserInfo1_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _F5_LookupVendor(p, 12)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func F5LTMUserInfo1_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 12, a)
}

func F5LTMUserInfo1_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 12, a)
}

func F5LTMUserInfo1_Del(p *radius.Packet) {
	_F5_DelVendor(p, 12)
}

func F5LTMUserInfo2_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 13, a)
}

func F5LTMUserInfo2_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 13, a)
}

func F5LTMUserInfo2_Get(p *radius.Packet) (value []byte) {
	value, _ = F5LTMUserInfo2_Lookup(p)
	return
}

func F5LTMUserInfo2_GetString(p *radius.Packet) (value string) {
	value, _ = F5LTMUserInfo2_LookupString(p)
	return
}

func F5LTMUserInfo2_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _F5_GetsVendor(p, 13) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserInfo2_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _F5_GetsVendor(p, 13) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMUserInfo2_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _F5_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func F5LTMUserInfo2_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _F5_LookupVendor(p, 13)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func F5LTMUserInfo2_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 13, a)
}

func F5LTMUserInfo2_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 13, a)
}

func F5LTMUserInfo2_Del(p *radius.Packet) {
	_F5_DelVendor(p, 13)
}

func F5LTMAuditMsg_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 14, a)
}

func F5LTMAuditMsg_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_AddVendor(p, 14, a)
}

func F5LTMAuditMsg_Get(p *radius.Packet) (value []byte) {
	value, _ = F5LTMAuditMsg_Lookup(p)
	return
}

func F5LTMAuditMsg_GetString(p *radius.Packet) (value string) {
	value, _ = F5LTMAuditMsg_LookupString(p)
	return
}

func F5LTMAuditMsg_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _F5_GetsVendor(p, 14) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMAuditMsg_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _F5_GetsVendor(p, 14) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func F5LTMAuditMsg_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _F5_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func F5LTMAuditMsg_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _F5_LookupVendor(p, 14)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func F5LTMAuditMsg_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 14, a)
}

func F5LTMAuditMsg_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _F5_SetVendor(p, 14, a)
}

func F5LTMAuditMsg_Del(p *radius.Packet) {
	_F5_DelVendor(p, 14)
}
