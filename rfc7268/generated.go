// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc7268

import (
	"errors"
	"strconv"

	"layeh.com/radius"
)

const (
	AllowedCalledStationID_Type radius.Type = 174
	EAPPeerID_Type              radius.Type = 175
	EAPServerID_Type            radius.Type = 176
	MobilityDomainID_Type       radius.Type = 177
	PreauthTimeout_Type         radius.Type = 178
	NetworkIDName_Type          radius.Type = 179
	EAPoLAnnouncement_Type      radius.Type = 180
	WLANHESSID_Type             radius.Type = 181
	WLANVenueInfo_Type          radius.Type = 182
	WLANVenueLanguage_Type      radius.Type = 183
	WLANVenueName_Type          radius.Type = 184
	WLANReasonCode_Type         radius.Type = 185
	WLANPairwiseCipher_Type     radius.Type = 186
	WLANGroupCipher_Type        radius.Type = 187
	WLANAKMSuite_Type           radius.Type = 188
	WLANGroupMgmtCipher_Type    radius.Type = 189
	WLANRFBand_Type             radius.Type = 190
)

func AllowedCalledStationID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(AllowedCalledStationID_Type, a)
	return
}

func AllowedCalledStationID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(AllowedCalledStationID_Type, a)
	return
}

func AllowedCalledStationID_Get(p *radius.Packet) (value []byte) {
	value, _ = AllowedCalledStationID_Lookup(p)
	return
}

func AllowedCalledStationID_GetString(p *radius.Packet) (value string) {
	value, _ = AllowedCalledStationID_LookupString(p)
	return
}

func AllowedCalledStationID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[AllowedCalledStationID_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AllowedCalledStationID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[AllowedCalledStationID_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func AllowedCalledStationID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(AllowedCalledStationID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func AllowedCalledStationID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(AllowedCalledStationID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func AllowedCalledStationID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(AllowedCalledStationID_Type, a)
	return
}

func AllowedCalledStationID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(AllowedCalledStationID_Type, a)
	return
}

func AllowedCalledStationID_Del(p *radius.Packet) {
	p.Attributes.Del(AllowedCalledStationID_Type)
}

func EAPPeerID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(EAPPeerID_Type, a)
	return
}

func EAPPeerID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(EAPPeerID_Type, a)
	return
}

func EAPPeerID_Get(p *radius.Packet) (value []byte) {
	value, _ = EAPPeerID_Lookup(p)
	return
}

func EAPPeerID_GetString(p *radius.Packet) (value string) {
	value, _ = EAPPeerID_LookupString(p)
	return
}

func EAPPeerID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[EAPPeerID_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPPeerID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[EAPPeerID_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPPeerID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(EAPPeerID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func EAPPeerID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(EAPPeerID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func EAPPeerID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(EAPPeerID_Type, a)
	return
}

func EAPPeerID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(EAPPeerID_Type, a)
	return
}

func EAPPeerID_Del(p *radius.Packet) {
	p.Attributes.Del(EAPPeerID_Type)
}

func EAPServerID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(EAPServerID_Type, a)
	return
}

func EAPServerID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(EAPServerID_Type, a)
	return
}

func EAPServerID_Get(p *radius.Packet) (value []byte) {
	value, _ = EAPServerID_Lookup(p)
	return
}

func EAPServerID_GetString(p *radius.Packet) (value string) {
	value, _ = EAPServerID_LookupString(p)
	return
}

func EAPServerID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[EAPServerID_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPServerID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[EAPServerID_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPServerID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(EAPServerID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func EAPServerID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(EAPServerID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func EAPServerID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(EAPServerID_Type, a)
	return
}

func EAPServerID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(EAPServerID_Type, a)
	return
}

func EAPServerID_Del(p *radius.Packet) {
	p.Attributes.Del(EAPServerID_Type)
}

type MobilityDomainID uint32

var MobilityDomainID_Strings = map[MobilityDomainID]string{}

func (a MobilityDomainID) String() string {
	if str, ok := MobilityDomainID_Strings[a]; ok {
		return str
	}
	return "MobilityDomainID(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func MobilityDomainID_Add(p *radius.Packet, value MobilityDomainID) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(MobilityDomainID_Type, a)
	return
}

func MobilityDomainID_Get(p *radius.Packet) (value MobilityDomainID) {
	value, _ = MobilityDomainID_Lookup(p)
	return
}

func MobilityDomainID_Gets(p *radius.Packet) (values []MobilityDomainID, err error) {
	var i uint32
	for _, attr := range p.Attributes[MobilityDomainID_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, MobilityDomainID(i))
	}
	return
}

func MobilityDomainID_Lookup(p *radius.Packet) (value MobilityDomainID, err error) {
	a, ok := p.Lookup(MobilityDomainID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = MobilityDomainID(i)
	return
}

func MobilityDomainID_Set(p *radius.Packet, value MobilityDomainID) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(MobilityDomainID_Type, a)
	return
}

func MobilityDomainID_Del(p *radius.Packet) {
	p.Attributes.Del(MobilityDomainID_Type)
}

type PreauthTimeout uint32

var PreauthTimeout_Strings = map[PreauthTimeout]string{}

func (a PreauthTimeout) String() string {
	if str, ok := PreauthTimeout_Strings[a]; ok {
		return str
	}
	return "PreauthTimeout(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PreauthTimeout_Add(p *radius.Packet, value PreauthTimeout) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(PreauthTimeout_Type, a)
	return
}

func PreauthTimeout_Get(p *radius.Packet) (value PreauthTimeout) {
	value, _ = PreauthTimeout_Lookup(p)
	return
}

func PreauthTimeout_Gets(p *radius.Packet) (values []PreauthTimeout, err error) {
	var i uint32
	for _, attr := range p.Attributes[PreauthTimeout_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PreauthTimeout(i))
	}
	return
}

func PreauthTimeout_Lookup(p *radius.Packet) (value PreauthTimeout, err error) {
	a, ok := p.Lookup(PreauthTimeout_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PreauthTimeout(i)
	return
}

func PreauthTimeout_Set(p *radius.Packet, value PreauthTimeout) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(PreauthTimeout_Type, a)
	return
}

func PreauthTimeout_Del(p *radius.Packet) {
	p.Attributes.Del(PreauthTimeout_Type)
}

func NetworkIDName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(NetworkIDName_Type, a)
	return
}

func NetworkIDName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(NetworkIDName_Type, a)
	return
}

func NetworkIDName_Get(p *radius.Packet) (value []byte) {
	value, _ = NetworkIDName_Lookup(p)
	return
}

func NetworkIDName_GetString(p *radius.Packet) (value string) {
	value, _ = NetworkIDName_LookupString(p)
	return
}

func NetworkIDName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[NetworkIDName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NetworkIDName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[NetworkIDName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NetworkIDName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(NetworkIDName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NetworkIDName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(NetworkIDName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NetworkIDName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(NetworkIDName_Type, a)
	return
}

func NetworkIDName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(NetworkIDName_Type, a)
	return
}

func NetworkIDName_Del(p *radius.Packet) {
	p.Attributes.Del(NetworkIDName_Type)
}

func EAPoLAnnouncement_Get(p *radius.Packet) (value []byte) {
	value, _ = EAPoLAnnouncement_Lookup(p)
	return
}

func EAPoLAnnouncement_GetString(p *radius.Packet) (value string) {
	value, _ = EAPoLAnnouncement_LookupString(p)
	return
}

func EAPoLAnnouncement_Lookup(p *radius.Packet) (value []byte, err error) {
	var i []byte
	var valid bool
	for _, attr := range p.Attributes[EAPoLAnnouncement_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		value = append(value, i...)
		valid = true
	}
	if !valid {
		err = radius.ErrNoAttribute
	}
	return
}

func EAPoLAnnouncement_LookupString(p *radius.Packet) (value string, err error) {
	var i string
	var valid bool
	for _, attr := range p.Attributes[EAPoLAnnouncement_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		value += i
		valid = true
	}
	if !valid {
		err = radius.ErrNoAttribute
	}
	return
}

func EAPoLAnnouncement_Set(p *radius.Packet, value []byte) (err error) {
	const maximumChunkSize = 253
	var attrs []radius.Attribute
	for len(value) > 0 {
		var a radius.Attribute
		n := len(value)
		if n > maximumChunkSize {
			n = maximumChunkSize
		}
		a, err = radius.NewBytes(value[:n])
		if err != nil {
			return
		}
		attrs = append(attrs, a)
		value = value[n:]
	}
	p.Attributes[EAPoLAnnouncement_Type] = attrs
	return
}

func EAPoLAnnouncement_SetString(p *radius.Packet, value string) (err error) {
	const maximumChunkSize = 253
	var attrs []radius.Attribute
	for len(value) > 0 {
		var a radius.Attribute
		n := len(value)
		if n > maximumChunkSize {
			n = maximumChunkSize
		}
		a, err = radius.NewString(value[:n])
		if err != nil {
			return
		}
		attrs = append(attrs, a)
		value = value[n:]
	}
	p.Attributes[EAPoLAnnouncement_Type] = attrs
	return
}

func EAPoLAnnouncement_Del(p *radius.Packet) {
	p.Attributes.Del(EAPoLAnnouncement_Type)
}

func WLANHESSID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(WLANHESSID_Type, a)
	return
}

func WLANHESSID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(WLANHESSID_Type, a)
	return
}

func WLANHESSID_Get(p *radius.Packet) (value []byte) {
	value, _ = WLANHESSID_Lookup(p)
	return
}

func WLANHESSID_GetString(p *radius.Packet) (value string) {
	value, _ = WLANHESSID_LookupString(p)
	return
}

func WLANHESSID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[WLANHESSID_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANHESSID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[WLANHESSID_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANHESSID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(WLANHESSID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WLANHESSID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(WLANHESSID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WLANHESSID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(WLANHESSID_Type, a)
	return
}

func WLANHESSID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(WLANHESSID_Type, a)
	return
}

func WLANHESSID_Del(p *radius.Packet) {
	p.Attributes.Del(WLANHESSID_Type)
}

type WLANVenueInfo uint32

var WLANVenueInfo_Strings = map[WLANVenueInfo]string{}

func (a WLANVenueInfo) String() string {
	if str, ok := WLANVenueInfo_Strings[a]; ok {
		return str
	}
	return "WLANVenueInfo(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANVenueInfo_Add(p *radius.Packet, value WLANVenueInfo) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANVenueInfo_Type, a)
	return
}

func WLANVenueInfo_Get(p *radius.Packet) (value WLANVenueInfo) {
	value, _ = WLANVenueInfo_Lookup(p)
	return
}

func WLANVenueInfo_Gets(p *radius.Packet) (values []WLANVenueInfo, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANVenueInfo_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANVenueInfo(i))
	}
	return
}

func WLANVenueInfo_Lookup(p *radius.Packet) (value WLANVenueInfo, err error) {
	a, ok := p.Lookup(WLANVenueInfo_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANVenueInfo(i)
	return
}

func WLANVenueInfo_Set(p *radius.Packet, value WLANVenueInfo) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANVenueInfo_Type, a)
	return
}

func WLANVenueInfo_Del(p *radius.Packet) {
	p.Attributes.Del(WLANVenueInfo_Type)
}

func WLANVenueLanguage_Add(p *radius.Packet, value []byte) (err error) {
	if len(value) != 3 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(WLANVenueLanguage_Type, a)
	return
}

func WLANVenueLanguage_AddString(p *radius.Packet, value string) (err error) {
	if len(value) != 3 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(WLANVenueLanguage_Type, a)
	return
}

func WLANVenueLanguage_Get(p *radius.Packet) (value []byte) {
	value, _ = WLANVenueLanguage_Lookup(p)
	return
}

func WLANVenueLanguage_GetString(p *radius.Packet) (value string) {
	value, _ = WLANVenueLanguage_LookupString(p)
	return
}

func WLANVenueLanguage_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[WLANVenueLanguage_Type] {
		i = radius.Bytes(attr)
		if err == nil && len(i) != 3 {
			err = errors.New("invalid value length")
		}
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANVenueLanguage_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[WLANVenueLanguage_Type] {
		i = radius.String(attr)
		if err == nil && len(i) != 3 {
			err = errors.New("invalid value length")
		}
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANVenueLanguage_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(WLANVenueLanguage_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	if err == nil && len(value) != 3 {
		err = errors.New("invalid value length")
	}
	return
}

func WLANVenueLanguage_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(WLANVenueLanguage_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	if err == nil && len(value) != 3 {
		err = errors.New("invalid value length")
	}
	return
}

func WLANVenueLanguage_Set(p *radius.Packet, value []byte) (err error) {
	if len(value) != 3 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(WLANVenueLanguage_Type, a)
	return
}

func WLANVenueLanguage_SetString(p *radius.Packet, value string) (err error) {
	if len(value) != 3 {
		err = errors.New("invalid value length")
		return
	}
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(WLANVenueLanguage_Type, a)
	return
}

func WLANVenueLanguage_Del(p *radius.Packet) {
	p.Attributes.Del(WLANVenueLanguage_Type)
}

func WLANVenueName_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(WLANVenueName_Type, a)
	return
}

func WLANVenueName_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(WLANVenueName_Type, a)
	return
}

func WLANVenueName_Get(p *radius.Packet) (value []byte) {
	value, _ = WLANVenueName_Lookup(p)
	return
}

func WLANVenueName_GetString(p *radius.Packet) (value string) {
	value, _ = WLANVenueName_LookupString(p)
	return
}

func WLANVenueName_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[WLANVenueName_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANVenueName_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[WLANVenueName_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func WLANVenueName_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(WLANVenueName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func WLANVenueName_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(WLANVenueName_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func WLANVenueName_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(WLANVenueName_Type, a)
	return
}

func WLANVenueName_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(WLANVenueName_Type, a)
	return
}

func WLANVenueName_Del(p *radius.Packet) {
	p.Attributes.Del(WLANVenueName_Type)
}

type WLANReasonCode uint32

var WLANReasonCode_Strings = map[WLANReasonCode]string{}

func (a WLANReasonCode) String() string {
	if str, ok := WLANReasonCode_Strings[a]; ok {
		return str
	}
	return "WLANReasonCode(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANReasonCode_Add(p *radius.Packet, value WLANReasonCode) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANReasonCode_Type, a)
	return
}

func WLANReasonCode_Get(p *radius.Packet) (value WLANReasonCode) {
	value, _ = WLANReasonCode_Lookup(p)
	return
}

func WLANReasonCode_Gets(p *radius.Packet) (values []WLANReasonCode, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANReasonCode_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANReasonCode(i))
	}
	return
}

func WLANReasonCode_Lookup(p *radius.Packet) (value WLANReasonCode, err error) {
	a, ok := p.Lookup(WLANReasonCode_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANReasonCode(i)
	return
}

func WLANReasonCode_Set(p *radius.Packet, value WLANReasonCode) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANReasonCode_Type, a)
	return
}

func WLANReasonCode_Del(p *radius.Packet) {
	p.Attributes.Del(WLANReasonCode_Type)
}

type WLANPairwiseCipher uint32

var WLANPairwiseCipher_Strings = map[WLANPairwiseCipher]string{}

func (a WLANPairwiseCipher) String() string {
	if str, ok := WLANPairwiseCipher_Strings[a]; ok {
		return str
	}
	return "WLANPairwiseCipher(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANPairwiseCipher_Add(p *radius.Packet, value WLANPairwiseCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANPairwiseCipher_Type, a)
	return
}

func WLANPairwiseCipher_Get(p *radius.Packet) (value WLANPairwiseCipher) {
	value, _ = WLANPairwiseCipher_Lookup(p)
	return
}

func WLANPairwiseCipher_Gets(p *radius.Packet) (values []WLANPairwiseCipher, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANPairwiseCipher_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANPairwiseCipher(i))
	}
	return
}

func WLANPairwiseCipher_Lookup(p *radius.Packet) (value WLANPairwiseCipher, err error) {
	a, ok := p.Lookup(WLANPairwiseCipher_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANPairwiseCipher(i)
	return
}

func WLANPairwiseCipher_Set(p *radius.Packet, value WLANPairwiseCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANPairwiseCipher_Type, a)
	return
}

func WLANPairwiseCipher_Del(p *radius.Packet) {
	p.Attributes.Del(WLANPairwiseCipher_Type)
}

type WLANGroupCipher uint32

var WLANGroupCipher_Strings = map[WLANGroupCipher]string{}

func (a WLANGroupCipher) String() string {
	if str, ok := WLANGroupCipher_Strings[a]; ok {
		return str
	}
	return "WLANGroupCipher(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANGroupCipher_Add(p *radius.Packet, value WLANGroupCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANGroupCipher_Type, a)
	return
}

func WLANGroupCipher_Get(p *radius.Packet) (value WLANGroupCipher) {
	value, _ = WLANGroupCipher_Lookup(p)
	return
}

func WLANGroupCipher_Gets(p *radius.Packet) (values []WLANGroupCipher, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANGroupCipher_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANGroupCipher(i))
	}
	return
}

func WLANGroupCipher_Lookup(p *radius.Packet) (value WLANGroupCipher, err error) {
	a, ok := p.Lookup(WLANGroupCipher_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANGroupCipher(i)
	return
}

func WLANGroupCipher_Set(p *radius.Packet, value WLANGroupCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANGroupCipher_Type, a)
	return
}

func WLANGroupCipher_Del(p *radius.Packet) {
	p.Attributes.Del(WLANGroupCipher_Type)
}

type WLANAKMSuite uint32

var WLANAKMSuite_Strings = map[WLANAKMSuite]string{}

func (a WLANAKMSuite) String() string {
	if str, ok := WLANAKMSuite_Strings[a]; ok {
		return str
	}
	return "WLANAKMSuite(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANAKMSuite_Add(p *radius.Packet, value WLANAKMSuite) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANAKMSuite_Type, a)
	return
}

func WLANAKMSuite_Get(p *radius.Packet) (value WLANAKMSuite) {
	value, _ = WLANAKMSuite_Lookup(p)
	return
}

func WLANAKMSuite_Gets(p *radius.Packet) (values []WLANAKMSuite, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANAKMSuite_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANAKMSuite(i))
	}
	return
}

func WLANAKMSuite_Lookup(p *radius.Packet) (value WLANAKMSuite, err error) {
	a, ok := p.Lookup(WLANAKMSuite_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANAKMSuite(i)
	return
}

func WLANAKMSuite_Set(p *radius.Packet, value WLANAKMSuite) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANAKMSuite_Type, a)
	return
}

func WLANAKMSuite_Del(p *radius.Packet) {
	p.Attributes.Del(WLANAKMSuite_Type)
}

type WLANGroupMgmtCipher uint32

var WLANGroupMgmtCipher_Strings = map[WLANGroupMgmtCipher]string{}

func (a WLANGroupMgmtCipher) String() string {
	if str, ok := WLANGroupMgmtCipher_Strings[a]; ok {
		return str
	}
	return "WLANGroupMgmtCipher(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANGroupMgmtCipher_Add(p *radius.Packet, value WLANGroupMgmtCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANGroupMgmtCipher_Type, a)
	return
}

func WLANGroupMgmtCipher_Get(p *radius.Packet) (value WLANGroupMgmtCipher) {
	value, _ = WLANGroupMgmtCipher_Lookup(p)
	return
}

func WLANGroupMgmtCipher_Gets(p *radius.Packet) (values []WLANGroupMgmtCipher, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANGroupMgmtCipher_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANGroupMgmtCipher(i))
	}
	return
}

func WLANGroupMgmtCipher_Lookup(p *radius.Packet) (value WLANGroupMgmtCipher, err error) {
	a, ok := p.Lookup(WLANGroupMgmtCipher_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANGroupMgmtCipher(i)
	return
}

func WLANGroupMgmtCipher_Set(p *radius.Packet, value WLANGroupMgmtCipher) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANGroupMgmtCipher_Type, a)
	return
}

func WLANGroupMgmtCipher_Del(p *radius.Packet) {
	p.Attributes.Del(WLANGroupMgmtCipher_Type)
}

type WLANRFBand uint32

var WLANRFBand_Strings = map[WLANRFBand]string{}

func (a WLANRFBand) String() string {
	if str, ok := WLANRFBand_Strings[a]; ok {
		return str
	}
	return "WLANRFBand(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func WLANRFBand_Add(p *radius.Packet, value WLANRFBand) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(WLANRFBand_Type, a)
	return
}

func WLANRFBand_Get(p *radius.Packet) (value WLANRFBand) {
	value, _ = WLANRFBand_Lookup(p)
	return
}

func WLANRFBand_Gets(p *radius.Packet) (values []WLANRFBand, err error) {
	var i uint32
	for _, attr := range p.Attributes[WLANRFBand_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, WLANRFBand(i))
	}
	return
}

func WLANRFBand_Lookup(p *radius.Packet) (value WLANRFBand, err error) {
	a, ok := p.Lookup(WLANRFBand_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = WLANRFBand(i)
	return
}

func WLANRFBand_Set(p *radius.Packet, value WLANRFBand) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(WLANRFBand_Type, a)
	return
}

func WLANRFBand_Del(p *radius.Packet) {
	p.Attributes.Del(WLANRFBand_Type)
}
