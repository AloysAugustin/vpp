// Code generated by GoVPP binapi-generator. DO NOT EDIT.
//  source: /usr/share/vpp/api/plugins/stn.api.json

/*
 Package stn is a generated from VPP binary API module 'stn'.

 It contains following objects:
	  2 services
	  4 messages
*/
package stn

import api "git.fd.io/govpp.git/api"
import struc "github.com/lunixbochs/struc"
import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = struc.Pack
var _ = bytes.NewBuffer

// Services represents VPP binary API services:
type Services interface {
	DumpStnRules(*StnRulesDump) ([]*StnRulesDetails, error)
	StnAddDelRule(*StnAddDelRule) (*StnAddDelRuleReply, error)
}

/* Messages */

// StnAddDelRule represents VPP binary API message 'stn_add_del_rule':
type StnAddDelRule struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
	IsAdd     uint8
}

func (*StnAddDelRule) GetMessageName() string {
	return "stn_add_del_rule"
}
func (*StnAddDelRule) GetCrcString() string {
	return "9f0bbe21"
}
func (*StnAddDelRule) GetMessageType() api.MessageType {
	return api.RequestMessage
}

// StnAddDelRuleReply represents VPP binary API message 'stn_add_del_rule_reply':
type StnAddDelRuleReply struct {
	Retval int32
}

func (*StnAddDelRuleReply) GetMessageName() string {
	return "stn_add_del_rule_reply"
}
func (*StnAddDelRuleReply) GetCrcString() string {
	return "e8d4e804"
}
func (*StnAddDelRuleReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// StnRulesDetails represents VPP binary API message 'stn_rules_details':
type StnRulesDetails struct {
	IsIP4     uint8
	IPAddress []byte `struc:"[16]byte"`
	SwIfIndex uint32
}

func (*StnRulesDetails) GetMessageName() string {
	return "stn_rules_details"
}
func (*StnRulesDetails) GetCrcString() string {
	return "5eafa31e"
}
func (*StnRulesDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

// StnRulesDump represents VPP binary API message 'stn_rules_dump':
type StnRulesDump struct{}

func (*StnRulesDump) GetMessageName() string {
	return "stn_rules_dump"
}
func (*StnRulesDump) GetCrcString() string {
	return "51077d14"
}
func (*StnRulesDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn.StnAddDelRule")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn.StnAddDelRuleReply")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn.StnRulesDetails")
	api.RegisterMessage((*StnRulesDump)(nil), "stn.StnRulesDump")
}

var Messages = []api.Message{
	(*StnAddDelRule)(nil),
	(*StnAddDelRuleReply)(nil),
	(*StnRulesDetails)(nil),
	(*StnRulesDump)(nil),
}
