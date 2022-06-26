package entity

import "encoding/xml"

type OperationRequest struct {
	XMLName   xml.Name  `xml:"ins:operationRequest"`
	Operation Operation `xml:"operation" json:"operation"`
}

type Operation struct {
	OperID                string               `xml:"iss:oper_id,omitempty" json:"oper_id"`
	OperType              string               `xml:"iss:oper_type,omitempty" json:"oper_type"`
	MsgType               string               `xml:"iss:msg_type,omitempty" json:"msg_type"`
	SttlType              string               `xml:"iss:sttl_type,omitempty" json:"sttl_type"`
	OriginalID            string               `xml:"iss:original_id,omitempty" json:"original_id"`
	DisputeID             string               `xml:"iss:dispute_id,omitempty" json:"dispute_id"`
	ReconciliationType    string               `xml:"iss:reconciliation_type,omitempty" json:"reconciliation_type"`
	OperDate              string               `xml:"iss:oper_date,omitempty" json:"oper_date"`
	HostDate              string               `xml:"iss:host_date,omitempty" json:"host_date"`
	OperCount             string               `xml:"iss:oper_count,omitempty" json:"oper_count"`
	OperAmount            *OperAmount          `xml:"iss:oper_amount,omitempty" json:"oper_amount"`
	OperRequestAmount     *OperRequestAmount   `xml:"iss:oper_request_amount,omitempty" json:"oper_request_amount"`
	OperSurchargeAmount   *OperSurchargeAmount `xml:"iss:oper_surcharge_amount,omitempty" json:"oper_surcharge_amount"`
	OperCashbackAmount    *OperCashbackAmount  `xml:"iss:oper_cashback_amount,omitempty" json:"oper_cashback_amount"`
	SttlAmount            *SttlAmount          `xml:"iss:sttl_amount,omitempty" json:"sttl_amount"`
	InterchangeFee        *InterchangeFee      `xml:"iss:interchange_fee,omitempty" json:"interchange_fee"`
	OriginatorRefnum      string               `xml:"iss:originator_refnum,omitempty" json:"originator_refnum"`
	NetworkRefnum         string               `xml:"iss:network_refnum,omitempty" json:"network_refnum"`
	AcqInstBin            string               `xml:"iss:acq_inst_bin,omitempty" json:"acq_inst_bin"`
	ForwardingInstBin     string               `xml:"iss:forwarding_inst_bin,omitempty" json:"forwarding_inst_bin"`
	OperReason            string               `xml:"iss:oper_reason,omitempty" json:"oper_reason"`
	Status                string               `xml:"iss:status,omitempty" json:"status"`
	StatusReason          string               `xml:"iss:status_reason,omitempty" json:"status_reason"`
	IsReversal            string               `xml:"iss:is_reversal,omitempty" json:"is_reversal"`
	MerchantNumber        string               `xml:"iss:merchant_number,omitempty" json:"merchant_number"`
	Mcc                   string               `xml:"iss:mcc,omitempty" json:"mcc"`
	MerchantName          string               `xml:"iss:merchant_name,omitempty" json:"merchant_name"`
	MerchantStreet        string               `xml:"iss:merchant_street,omitempty" json:"merchant_street"`
	MerchantCity          string               `xml:"iss:merchant_city,omitempty" json:"merchant_city"`
	MerchantRegion        string               `xml:"iss:merchant_region,omitempty" json:"merchant_region"`
	MerchantCountry       string               `xml:"iss:merchant_country,omitempty" json:"merchant_country"`
	MerchantPostcode      string               `xml:"iss:merchant_postcode,omitempty" json:"merchant_postcode"`
	TerminalType          string               `xml:"iss:terminal_type,omitempty" json:"terminal_type"`
	TerminalNumber        string               `xml:"iss:terminal_number,omitempty" json:"terminal_number"`
	SttlDate              string               `xml:"iss:sttl_date,omitempty" json:"sttl_date"`
	AcqSttlDate           string               `xml:"iss:acq_sttl_date,omitempty" json:"acq_sttl_date"`
	MatchStatus           string               `xml:"iss:match_status,omitempty" json:"match_status"`
	PaymentOrder          *PaymentOrder        `xml:"iss:payment_order,omitempty" json:"payment_order"`
	Issuer                *Issuer              `xml:"iss:issuer,omitempty" json:"issuer"`
	Acquirer              *Acquirer            `xml:"iss:acquirer,omitempty" json:"acquirer"`
	Destination           *Destination         `xml:"iss:destination,omitempty" json:"destination"`
	Aggregator            *Aggregator          `xml:"iss:aggregator,omitempty" json:"aggregator"`
	ServiceProvider       *ServiceProvider     `xml:"iss:service_provider,omitempty" json:"service_provider"`
	Participant           *Participant         `xml:"iss:participant,omitempty" json:"participant"`
	AuthData              *AuthData            `xml:"iss:auth_data,omitempty" json:"auth_data"`
	AdditionalAmount      *AdditionalAmount    `xml:"iss:additional_amount,omitempty" json:"additional_amount"`
	ProcessingStage       *ProcessingStage     `xml:"iss:processing_stage,omitempty" json:"processing_stage"`
	FlexibleData          []FlexibleData       `xml:"iss:flexible_data,omitempty" json:"flexible_data"`
	ClearingSequenceNum   string               `xml:"iss:clearing_sequence_num,omitempty" json:"clearing_sequence_num"`
	ClearingSequenceCount string               `xml:"iss:clearing_sequence_count,omitempty" json:"clearing_sequence_count"`
}

type FlexibleData struct {
	FieldName  string `xml:"iss:field_name,omitempty" json:"field_name"`
	FieldValue string `xml:"iss:field_value,omitempty" json:"field_value"`
}

type ProcessingStage struct {
	ProcStage string `xml:"iss:proc_stage,omitempty" json:"proc_stage"`
	Status    string `xml:"iss:status,omitempty" json:"status"`
}

type AdditionalAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
	AmountType  string `xml:"iss:amount_type,omitempty" json:"amount_type"`
}

type AuthTag struct {
	TagID     string `xml:"iss:tag_id,omitempty" json:"tag_id"`
	TagName   string `xml:"iss:tag_name,omitempty" json:"tag_name"`
	TagValue  string `xml:"iss:tag_value,omitempty" json:"tag_value"`
	SeqNumber string `xml:"iss:seq_number,omitempty" json:"seq_number"`
}

type AuthData struct {
	RespCode               string  `xml:"iss:resp_code,omitempty" json:"resp_code"`
	ProcType               string  `xml:"iss:proc_type,omitempty" json:"proc_type"`
	ProcMode               string  `xml:"iss:proc_mode,omitempty" json:"proc_mode"`
	IsAdvice               string  `xml:"iss:is_advice,omitempty" json:"is_advice"`
	IsRepeat               string  `xml:"iss:is_repeat,omitempty" json:"is_repeat"`
	BinAmount              string  `xml:"iss:bin_amount,omitempty" json:"bin_amount"`
	BinCurrency            string  `xml:"iss:bin_currency,omitempty" json:"bin_currency"`
	BinCnvtRate            string  `xml:"iss:bin_cnvt_rate,omitempty" json:"bin_cnvt_rate"`
	NetworkAmount          string  `xml:"iss:network_amount,omitempty" json:"network_amount"`
	NetworkCurrency        string  `xml:"iss:network_currency,omitempty" json:"network_currency"`
	NetworkCnvtDate        string  `xml:"iss:network_cnvt_date,omitempty" json:"network_cnvt_date"`
	NetworkCnvtRate        string  `xml:"iss:network_cnvt_rate,omitempty" json:"network_cnvt_rate"`
	AccountCnvtRate        string  `xml:"iss:account_cnvt_rate,omitempty" json:"account_cnvt_rate"`
	AddrVerifResult        string  `xml:"iss:addr_verif_result,omitempty" json:"addr_verif_result"`
	AcqRespCode            string  `xml:"iss:acq_resp_code,omitempty" json:"acq_resp_code"`
	AcqDeviceProcResult    string  `xml:"iss:acq_device_proc_result,omitempty" json:"acq_device_proc_result"`
	CardDataInputCap       string  `xml:"iss:card_data_input_cap,omitempty" json:"card_data_input_cap"`
	CrdhAuthCap            string  `xml:"iss:crdh_auth_cap,omitempty" json:"crdh_auth_cap"`
	CardCaptureCap         string  `xml:"iss:card_capture_cap,omitempty" json:"card_capture_cap"`
	TerminalOperatingEnv   string  `xml:"iss:terminal_operating_env,omitempty" json:"terminal_operating_env"`
	CrdhPresence           string  `xml:"iss:crdh_presence,omitempty" json:"crdh_presence"`
	CardPresence           string  `xml:"iss:card_presence,omitempty" json:"card_presence"`
	CardDataInputMode      string  `xml:"iss:card_data_input_mode,omitempty" json:"card_data_input_mode"`
	CrdhAuthMethod         string  `xml:"iss:crdh_auth_method,omitempty" json:"crdh_auth_method"`
	CrdhAuthEntity         string  `xml:"iss:crdh_auth_entity,omitempty" json:"crdh_auth_entity"`
	CardDataOutputCap      string  `xml:"iss:card_data_output_cap,omitempty" json:"card_data_output_cap"`
	TerminalOutputCap      string  `xml:"iss:terminal_output_cap,omitempty" json:"terminal_output_cap"`
	PinCaptureCap          string  `xml:"iss:pin_capture_cap,omitempty" json:"pin_capture_cap"`
	PinPresence            string  `xml:"iss:pin_presence,omitempty" json:"pin_presence"`
	Cvv2Presence           string  `xml:"iss:cvv2_presence,omitempty" json:"cvv_2_presence"`
	CvcIndicator           string  `xml:"iss:cvc_indicator,omitempty" json:"cvc_indicator"`
	PosEntryMode           string  `xml:"iss:pos_entry_mode,omitempty" json:"pos_entry_mode"`
	PosCondCode            string  `xml:"iss:pos_cond_code,omitempty" json:"pos_cond_code"`
	EmvData                string  `xml:"iss:emv_data,omitempty" json:"emv_data"`
	AddlData               string  `xml:"iss:addl_data,omitempty" json:"addl_data"`
	ServiceCode            string  `xml:"iss:service_code,omitempty" json:"service_code"`
	DeviceDate             string  `xml:"iss:device_date,omitempty" json:"device_date"`
	Cvv2Result             string  `xml:"iss:cvv2_result,omitempty" json:"cvv_2_result"`
	IsCompleted            string  `xml:"iss:is_completed,omitempty" json:"is_completed"`
	Amounts                string  `xml:"iss:amounts,omitempty" json:"amounts"`
	SystemTraceAuditNumber string  `xml:"iss:system_trace_audit_number,omitempty" json:"system_trace_audit_number"`
	AuthTransactionID      string  `xml:"iss:auth_transaction_id,omitempty" json:"auth_transaction_id"`
	ExternalAuthID         string  `xml:"iss:external_auth_id,omitempty" json:"external_auth_id"`
	ExternalOrigID         string  `xml:"iss:external_orig_id,omitempty" json:"external_orig_id"`
	AgentUniqueID          string  `xml:"iss:agent_unique_id,omitempty" json:"agent_unique_id"`
	NativeRespCode         string  `xml:"iss:native_resp_code,omitempty" json:"native_resp_code"`
	AuthPurposeID          string  `xml:"iss:auth_purpose_id,omitempty" json:"auth_purpose_id"`
	AuthTag                AuthTag `xml:"iss:auth_tag,omitempty" json:"auth_tag"`
}

type Participant struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type ServiceProvider struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type Aggregator struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type Destination struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type Acquirer struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type Issuer struct {
	ParticipantType string `xml:"iss:participant_type,omitempty" json:"participant_type"`
	ClientIDType    string `xml:"iss:client_id_type,omitempty" json:"client_id_type"`
	ClientIDValue   string `xml:"iss:client_id_value,omitempty" json:"client_id_value"`
	CardNumber      string `xml:"iss:card_number,omitempty" json:"card_number"`
	CardID          string `xml:"iss:card_id,omitempty" json:"card_id"`
	CardInstanceID  string `xml:"iss:card_instance_id,omitempty" json:"card_instance_id"`
	CardSeqNumber   string `xml:"iss:card_seq_number,omitempty" json:"card_seq_number"`
	CardExpirDate   string `xml:"iss:card_expir_date,omitempty" json:"card_expir_date"`
	CardCountry     string `xml:"iss:card_country,omitempty" json:"card_country"`
	InstID          string `xml:"iss:inst_id,omitempty" json:"inst_id"`
	NetworkID       string `xml:"iss:network_id,omitempty" json:"network_id"`
	AuthCode        string `xml:"iss:auth_code,omitempty" json:"auth_code"`
	AccountNumber   string `xml:"iss:account_number,omitempty" json:"account_number"`
	AccountAmount   string `xml:"iss:account_amount,omitempty" json:"account_amount"`
	AccountCurrency string `xml:"iss:account_currency,omitempty" json:"account_currency"`
}

type SttlAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}

type InterchangeFee struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}

type OperAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}

type OperRequestAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}

type OperSurchargeAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}

type OperCashbackAmount struct {
	AmountValue string `xml:"iss:amount_value,omitempty" json:"amount_value"`
	Currency    string `xml:"iss:currency,omitempty" json:"currency"`
}
