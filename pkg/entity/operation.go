package entity

type OperationRequest struct {
	Operation Operation `xml:"operation"`
}

type Operation struct {
	OperID                string              `xml:"oper_id"`
	OperType              string              `xml:"oper_type"`
	MsgType               string              `xml:"msg_type"`
	SttlType              string              `xml:"sttl_type"`
	OriginalID            string              `xml:"original_id"`
	DisputeID             string              `xml:"dispute_id"`
	ReconciliationType    string              `xml:"reconciliation_type"`
	OperDate              string              `xml:"oper_date"`
	HostDate              string              `xml:"host_date"`
	OperCount             string              `xml:"oper_count"`
	OperAmount            OperAmount          `xml:"oper_amount"`
	OperRequestAmount     OperRequestAmount   `xml:"oper_request_amount"`
	OperSurchargeAmount   OperSurchargeAmount `xml:"oper_surcharge_amount"`
	OperCashbackAmount    OperCashbackAmount  `xml:"oper_cashback_amount"`
	SttlAmount            SttlAmount          `xml:"sttl_amount"`
	InterchangeFee        InterchangeFee      `xml:"interchange_fee"`
	OriginatorRefnum      string              `xml:"originator_refnum"`
	NetworkRefnum         string              `xml:"network_refnum"`
	AcqInstBin            string              `xml:"acq_inst_bin"`
	ForwardingInstBin     string              `xml:"forwarding_inst_bin"`
	OperReason            string              `xml:"oper_reason"`
	Status                string              `xml:"status"`
	StatusReason          string              `xml:"status_reason"`
	IsReversal            string              `xml:"is_reversal"`
	MerchantNumber        string              `xml:"merchant_number"`
	Mcc                   string              `xml:"mcc"`
	MerchantName          string              `xml:"merchant_name"`
	MerchantStreet        string              `xml:"merchant_street"`
	MerchantCity          string              `xml:"merchant_city"`
	MerchantRegion        string              `xml:"merchant_region"`
	MerchantCountry       string              `xml:"merchant_country"`
	MerchantPostcode      string              `xml:"merchant_postcode"`
	TerminalType          string              `xml:"terminal_type"`
	TerminalNumber        string              `xml:"terminal_number"`
	SttlDate              string              `xml:"sttl_date"`
	AcqSttlDate           string              `xml:"acq_sttl_date"`
	MatchStatus           string              `xml:"match_status"`
	PaymentOrder          PaymentOrder        `xml:"payment_order"`
	Issuer                Issuer              `xml:"issuer"`
	Acquirer              Acquirer            `xml:"acquirer"`
	Destination           Destination         `xml:"destination"`
	Aggregator            Aggregator          `xml:"aggregator"`
	ServiceProvider       ServiceProvider     `xml:"service_provider"`
	Participant           Participant         `xml:"participant"`
	AuthData              AuthData            `xml:"auth_data"`
	AdditionalAmount      AdditionalAmount    `xml:"additional_amount"`
	ProcessingStage       ProcessingStage     `xml:"processing_stage"`
	FlexibleData          FlexibleData        `xml:"flexible_data"`
	ClearingSequenceNum   string              `xml:"clearing_sequence_num"`
	ClearingSequenceCount string              `xml:"clearing_sequence_count"`
}

type FlexibleData struct {
	FieldName  string `xml:"field_name"`
	FieldValue string `xml:"field_value"`
}

type ProcessingStage struct {
	ProcStage string `xml:"proc_stage"`
	Status    string `xml:"status"`
}

type AdditionalAmount struct {
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
	AmountType  string `xml:"amount_type"`
}

type AuthTag struct {
	TagID     string `xml:"tag_id"`
	TagName   string `xml:"tag_name"`
	TagValue  string `xml:"tag_value"`
	SeqNumber string `xml:"seq_number"`
}

type AuthData struct {
	RespCode               string  `xml:"resp_code"`
	ProcType               string  `xml:"proc_type"`
	ProcMode               string  `xml:"proc_mode"`
	IsAdvice               string  `xml:"is_advice"`
	IsRepeat               string  `xml:"is_repeat"`
	BinAmount              string  `xml:"bin_amount"`
	BinCurrency            string  `xml:"bin_currency"`
	BinCnvtRate            string  `xml:"bin_cnvt_rate"`
	NetworkAmount          string  `xml:"network_amount"`
	NetworkCurrency        string  `xml:"network_currency"`
	NetworkCnvtDate        string  `xml:"network_cnvt_date"`
	NetworkCnvtRate        string  `xml:"network_cnvt_rate"`
	AccountCnvtRate        string  `xml:"account_cnvt_rate"`
	AddrVerifResult        string  `xml:"addr_verif_result"`
	AcqRespCode            string  `xml:"acq_resp_code"`
	AcqDeviceProcResult    string  `xml:"acq_device_proc_result"`
	CardDataInputCap       string  `xml:"card_data_input_cap"`
	CrdhAuthCap            string  `xml:"crdh_auth_cap"`
	CardCaptureCap         string  `xml:"card_capture_cap"`
	TerminalOperatingEnv   string  `xml:"terminal_operating_env"`
	CrdhPresence           string  `xml:"crdh_presence"`
	CardPresence           string  `xml:"card_presence"`
	CardDataInputMode      string  `xml:"card_data_input_mode"`
	CrdhAuthMethod         string  `xml:"crdh_auth_method"`
	CrdhAuthEntity         string  `xml:"crdh_auth_entity"`
	CardDataOutputCap      string  `xml:"card_data_output_cap"`
	TerminalOutputCap      string  `xml:"terminal_output_cap"`
	PinCaptureCap          string  `xml:"pin_capture_cap"`
	PinPresence            string  `xml:"pin_presence"`
	Cvv2Presence           string  `xml:"cvv2_presence"`
	CvcIndicator           string  `xml:"cvc_indicator"`
	PosEntryMode           string  `xml:"pos_entry_mode"`
	PosCondCode            string  `xml:"pos_cond_code"`
	EmvData                string  `xml:"emv_data"`
	AddlData               string  `xml:"addl_data"`
	ServiceCode            string  `xml:"service_code"`
	DeviceDate             string  `xml:"device_date"`
	Cvv2Result             string  `xml:"cvv2_result"`
	IsCompleted            string  `xml:"is_completed"`
	Amounts                string  `xml:"amounts"`
	SystemTraceAuditNumber string  `xml:"system_trace_audit_number"`
	AuthTransactionID      string  `xml:"auth_transaction_id"`
	ExternalAuthID         string  `xml:"external_auth_id"`
	ExternalOrigID         string  `xml:"external_orig_id"`
	AgentUniqueID          string  `xml:"agent_unique_id"`
	NativeRespCode         string  `xml:"native_resp_code"`
	AuthPurposeID          string  `xml:"auth_purpose_id"`
	AuthTag                AuthTag `xml:"auth_tag"`
}

type Participant struct {
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type ServiceProvider struct {
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type Aggregator struct {
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type Destination struct {
	Text            string `xml:",chardata"`
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type Acquirer struct {
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type Issuer struct {
	ParticipantType string `xml:"participant_type"`
	ClientIDType    string `xml:"client_id_type"`
	ClientIDValue   string `xml:"client_id_value"`
	CardNumber      string `xml:"card_number"`
	CardID          string `xml:"card_id"`
	CardInstanceID  string `xml:"card_instance_id"`
	CardSeqNumber   string `xml:"card_seq_number"`
	CardExpirDate   string `xml:"card_expir_date"`
	CardCountry     string `xml:"card_country"`
	InstID          string `xml:"inst_id"`
	NetworkID       string `xml:"network_id"`
	AuthCode        string `xml:"auth_code"`
	AccountNumber   string `xml:"account_number"`
	AccountAmount   string `xml:"account_amount"`
	AccountCurrency string `xml:"account_currency"`
}

type PaymentOrder struct {
	PaymentOrderID     string           `xml:"payment_order_id"`
	PaymentOrderStatus string           `xml:"payment_order_status"`
	PaymentOrderNumber string           `xml:"payment_order_number"`
	PurposeID          string           `xml:"purpose_id"`
	PurposeNumber      string           `xml:"purpose_number"`
	PaymentAmount      PaymentAmount    `xml:"payment_amount"`
	PaymentDate        string           `xml:"payment_date"`
	PaymentParameter   PaymentParameter `xml:"payment_parameter"`
	ParticipantType    string           `xml:"participant_type"`
}

type PaymentAmount struct {
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type PaymentParameter struct {
	PaymentParameterName  string `xml:"payment_parameter_name"`
	PaymentParameterValue string `xml:"payment_parameter_value"`
}

type SttlAmount struct {
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type InterchangeFee struct {
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type OperAmount struct {
	Text        string `xml:",chardata"`
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type OperRequestAmount struct {
	Text        string `xml:",chardata"`
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type OperSurchargeAmount struct {
	Text        string `xml:",chardata"`
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}

type OperCashbackAmount struct {
	Text        string `xml:",chardata"`
	AmountValue string `xml:"amount_value"`
	Currency    string `xml:"currency"`
}
