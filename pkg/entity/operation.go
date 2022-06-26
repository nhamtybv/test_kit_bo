package entity

type OperationRequest struct {
	Operation Operation `xml:"operation" json:"operation"`
}

type Operation struct {
	OperID                string              `xml:"oper_id" json:"oper_id"`
	OperType              string              `xml:"oper_type" json:"oper_type"`
	MsgType               string              `xml:"msg_type" json:"msg_type"`
	SttlType              string              `xml:"sttl_type" json:"sttl_type"`
	OriginalID            string              `xml:"original_id" json:"original_id"`
	DisputeID             string              `xml:"dispute_id" json:"dispute_id"`
	ReconciliationType    string              `xml:"reconciliation_type" json:"reconciliation_type"`
	OperDate              string              `xml:"oper_date" json:"oper_date"`
	HostDate              string              `xml:"host_date" json:"host_date"`
	OperCount             string              `xml:"oper_count" json:"oper_count"`
	OperAmount            OperAmount          `xml:"oper_amount" json:"oper_amount"`
	OperRequestAmount     OperRequestAmount   `xml:"oper_request_amount" json:"oper_request_amount"`
	OperSurchargeAmount   OperSurchargeAmount `xml:"oper_surcharge_amount" json:"oper_surcharge_amount"`
	OperCashbackAmount    OperCashbackAmount  `xml:"oper_cashback_amount" json:"oper_cashback_amount"`
	SttlAmount            SttlAmount          `xml:"sttl_amount" json:"sttl_amount"`
	InterchangeFee        InterchangeFee      `xml:"interchange_fee" json:"interchange_fee"`
	OriginatorRefnum      string              `xml:"originator_refnum" json:"originator_refnum"`
	NetworkRefnum         string              `xml:"network_refnum" json:"network_refnum"`
	AcqInstBin            string              `xml:"acq_inst_bin" json:"acq_inst_bin"`
	ForwardingInstBin     string              `xml:"forwarding_inst_bin" json:"forwarding_inst_bin"`
	OperReason            string              `xml:"oper_reason" json:"oper_reason"`
	Status                string              `xml:"status" json:"status"`
	StatusReason          string              `xml:"status_reason" json:"status_reason"`
	IsReversal            string              `xml:"is_reversal" json:"is_reversal"`
	MerchantNumber        string              `xml:"merchant_number" json:"merchant_number"`
	Mcc                   string              `xml:"mcc" json:"mcc"`
	MerchantName          string              `xml:"merchant_name" json:"merchant_name"`
	MerchantStreet        string              `xml:"merchant_street" json:"merchant_street"`
	MerchantCity          string              `xml:"merchant_city" json:"merchant_city"`
	MerchantRegion        string              `xml:"merchant_region" json:"merchant_region"`
	MerchantCountry       string              `xml:"merchant_country" json:"merchant_country"`
	MerchantPostcode      string              `xml:"merchant_postcode" json:"merchant_postcode"`
	TerminalType          string              `xml:"terminal_type" json:"terminal_type"`
	TerminalNumber        string              `xml:"terminal_number" json:"terminal_number"`
	SttlDate              string              `xml:"sttl_date" json:"sttl_date"`
	AcqSttlDate           string              `xml:"acq_sttl_date" json:"acq_sttl_date"`
	MatchStatus           string              `xml:"match_status" json:"match_status"`
	PaymentOrder          PaymentOrder        `xml:"payment_order" json:"payment_order"`
	Issuer                Issuer              `xml:"issuer" json:"issuer"`
	Acquirer              Acquirer            `xml:"acquirer" json:"acquirer"`
	Destination           Destination         `xml:"destination" json:"destination"`
	Aggregator            Aggregator          `xml:"aggregator" json:"aggregator"`
	ServiceProvider       ServiceProvider     `xml:"service_provider" json:"service_provider"`
	Participant           Participant         `xml:"participant" json:"participant"`
	AuthData              AuthData            `xml:"auth_data" json:"auth_data"`
	AdditionalAmount      AdditionalAmount    `xml:"additional_amount" json:"additional_amount"`
	ProcessingStage       ProcessingStage     `xml:"processing_stage" json:"processing_stage"`
	FlexibleData          []FlexibleData      `xml:"flexible_data" json:"flexible_data"`
	ClearingSequenceNum   string              `xml:"clearing_sequence_num" json:"clearing_sequence_num"`
	ClearingSequenceCount string              `xml:"clearing_sequence_count" json:"clearing_sequence_count"`
}

type FlexibleData struct {
	FieldName  string `xml:"field_name" json:"field_name"`
	FieldValue string `xml:"field_value" json:"field_value"`
}

type ProcessingStage struct {
	ProcStage string `xml:"proc_stage" json:"proc_stage"`
	Status    string `xml:"status" json:"status"`
}

type AdditionalAmount struct {
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
	AmountType  string `xml:"amount_type" json:"amount_type"`
}

type AuthTag struct {
	TagID     string `xml:"tag_id" json:"tag_id"`
	TagName   string `xml:"tag_name" json:"tag_name"`
	TagValue  string `xml:"tag_value" json:"tag_value"`
	SeqNumber string `xml:"seq_number" json:"seq_number"`
}

type AuthData struct {
	RespCode               string  `xml:"resp_code" json:"resp_code"`
	ProcType               string  `xml:"proc_type" json:"proc_type"`
	ProcMode               string  `xml:"proc_mode" json:"proc_mode"`
	IsAdvice               string  `xml:"is_advice" json:"is_advice"`
	IsRepeat               string  `xml:"is_repeat" json:"is_repeat"`
	BinAmount              string  `xml:"bin_amount" json:"bin_amount"`
	BinCurrency            string  `xml:"bin_currency" json:"bin_currency"`
	BinCnvtRate            string  `xml:"bin_cnvt_rate" json:"bin_cnvt_rate"`
	NetworkAmount          string  `xml:"network_amount" json:"network_amount"`
	NetworkCurrency        string  `xml:"network_currency" json:"network_currency"`
	NetworkCnvtDate        string  `xml:"network_cnvt_date" json:"network_cnvt_date"`
	NetworkCnvtRate        string  `xml:"network_cnvt_rate" json:"network_cnvt_rate"`
	AccountCnvtRate        string  `xml:"account_cnvt_rate" json:"account_cnvt_rate"`
	AddrVerifResult        string  `xml:"addr_verif_result" json:"addr_verif_result"`
	AcqRespCode            string  `xml:"acq_resp_code" json:"acq_resp_code"`
	AcqDeviceProcResult    string  `xml:"acq_device_proc_result" json:"acq_device_proc_result"`
	CardDataInputCap       string  `xml:"card_data_input_cap" json:"card_data_input_cap"`
	CrdhAuthCap            string  `xml:"crdh_auth_cap" json:"crdh_auth_cap"`
	CardCaptureCap         string  `xml:"card_capture_cap" json:"card_capture_cap"`
	TerminalOperatingEnv   string  `xml:"terminal_operating_env" json:"terminal_operating_env"`
	CrdhPresence           string  `xml:"crdh_presence" json:"crdh_presence"`
	CardPresence           string  `xml:"card_presence" json:"card_presence"`
	CardDataInputMode      string  `xml:"card_data_input_mode" json:"card_data_input_mode"`
	CrdhAuthMethod         string  `xml:"crdh_auth_method" json:"crdh_auth_method"`
	CrdhAuthEntity         string  `xml:"crdh_auth_entity" json:"crdh_auth_entity"`
	CardDataOutputCap      string  `xml:"card_data_output_cap" json:"card_data_output_cap"`
	TerminalOutputCap      string  `xml:"terminal_output_cap" json:"terminal_output_cap"`
	PinCaptureCap          string  `xml:"pin_capture_cap" json:"pin_capture_cap"`
	PinPresence            string  `xml:"pin_presence" json:"pin_presence"`
	Cvv2Presence           string  `xml:"cvv2_presence" json:"cvv_2_presence"`
	CvcIndicator           string  `xml:"cvc_indicator" json:"cvc_indicator"`
	PosEntryMode           string  `xml:"pos_entry_mode" json:"pos_entry_mode"`
	PosCondCode            string  `xml:"pos_cond_code" json:"pos_cond_code"`
	EmvData                string  `xml:"emv_data" json:"emv_data"`
	AddlData               string  `xml:"addl_data" json:"addl_data"`
	ServiceCode            string  `xml:"service_code" json:"service_code"`
	DeviceDate             string  `xml:"device_date" json:"device_date"`
	Cvv2Result             string  `xml:"cvv2_result" json:"cvv_2_result"`
	IsCompleted            string  `xml:"is_completed" json:"is_completed"`
	Amounts                string  `xml:"amounts" json:"amounts"`
	SystemTraceAuditNumber string  `xml:"system_trace_audit_number" json:"system_trace_audit_number"`
	AuthTransactionID      string  `xml:"auth_transaction_id" json:"auth_transaction_id"`
	ExternalAuthID         string  `xml:"external_auth_id" json:"external_auth_id"`
	ExternalOrigID         string  `xml:"external_orig_id" json:"external_orig_id"`
	AgentUniqueID          string  `xml:"agent_unique_id" json:"agent_unique_id"`
	NativeRespCode         string  `xml:"native_resp_code" json:"native_resp_code"`
	AuthPurposeID          string  `xml:"auth_purpose_id" json:"auth_purpose_id"`
	AuthTag                AuthTag `xml:"auth_tag" json:"auth_tag"`
}

type Participant struct {
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type ServiceProvider struct {
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type Aggregator struct {
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type Destination struct {
	Text            string `xml:",chardata" json:"text"`
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type Acquirer struct {
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type Issuer struct {
	ParticipantType string `xml:"participant_type" json:"participant_type"`
	ClientIDType    string `xml:"client_id_type" json:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value" json:"client_id_value"`
	CardNumber      string `xml:"card_number" json:"card_number"`
	CardID          string `xml:"card_id" json:"card_id"`
	CardInstanceID  string `xml:"card_instance_id" json:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number" json:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date" json:"card_expir_date"`
	CardCountry     string `xml:"card_country" json:"card_country"`
	InstID          string `xml:"inst_id" json:"inst_id"`
	NetworkID       string `xml:"network_id" json:"network_id"`
	AuthCode        string `xml:"auth_code" json:"auth_code"`
	AccountNumber   string `xml:"account_number" json:"account_number"`
	AccountAmount   string `xml:"account_amount" json:"account_amount"`
	AccountCurrency string `xml:"account_currency" json:"account_currency"`
}

type SttlAmount struct {
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}

type InterchangeFee struct {
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}

type OperAmount struct {
	Text        string `xml:",chardata" json:"text"`
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}

type OperRequestAmount struct {
	Text        string `xml:",chardata" json:"text"`
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}

type OperSurchargeAmount struct {
	Text        string `xml:",chardata" json:"text"`
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}

type OperCashbackAmount struct {
	Text        string `xml:",chardata" json:"text"`
	AmountValue string `xml:"amount_value" json:"amount_value"`
	Currency    string `xml:"currency" json:"currency"`
}
