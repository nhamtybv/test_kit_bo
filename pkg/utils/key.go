package utils

type conextId string

const (
	OracleConnectionKey conextId = "oracle_connection"
	SettingTable        string   = "SETTINGS"
	ProductTable        string   = "PRODUCTS"
	AgentTable          string   = "AGENTS"
	TestCaseTable       string   = "TESTCASES_DATA"
	CardTable           string   = "CARDS_DATA"
	MainBucket          string   = "CONFIG_BUCKET"
	DatabaseName        string   = "config.db"
	WebserviceAddress   string   = "webservice_address"
)

const (
	AccountMaintenance           string = "Account maintenance"
	DirectDebeting               string = "Direct debeting"
	CustomerNotification         string = "Customer notification"
	CardNotification             string = "Card notification"
	CustomerMaintenanceService   string = "Customer maintenance service"
	CustomerReferralService      string = "Customer referral service"
	CardMaintenance              string = "Card maintenance"
	IssuingContractServicing     string = "Issuing contract servicing"
	MobilePaymentServicing       string = "Mobile payment servicing"
	MerchantCardMaintenance      string = "Merchant card maintenance"
	MastercardRewardService      string = "MasterCard reward service"
	Credit                       string = "Credit"
	DeferredPaymentPlan          string = "Deferred payment plan"
	InsuranceCompanySettlement   string = "Insurance company settlement"
	CreditInsurance              string = "Credit insurance"
	LoyaltyForCard               string = "Loyalty for card"
	LoyaltyForAccount            string = "Loyalty for account"
	LoyaltyForCustomer           string = "Loyalty for customer"
	ServiceProviderSettlement    string = "Service provider settlement"
	S3DSecureProgram             string = "3D secure program"
	SettlementAccountMaintenance string = "Settlement account maintenance"
	AntiVAUParticipation         string = "Anti-VAU participation"
)

const (
	ErrorNoDataFound string = "no_data_found"
)

const (
	CREATE_CARD         string = "flow_1001"
	ADD_SUB_CAR         string = "flow_1003"
	REISSUE_CARD        string = "flow_0005"
	APPLICATION_SERVICE string = "ApplicationService"
	INSTANT_ISSUE       string = "InstantIssue"
	CLEARING_WS         string = "ClearingWS"
	CARD_NUMBER         string = "card_number"
	CARD_ID             string = "card_id"
	APPLICATION_ID      string = "application_id"
	CHANGE_CARD_STATE   string = "process_ibg_data_request"
)

const (
	CHANGE_CARD_STATUS string = "OPTP0171"
	PURCHASE           string = "OPTP0000"
	CASH_WDR           string = "OPTP0001"
	PAYMENT            string = "OPTP0422"
	CARD_ISSUED        string = "CSTE0100"
	CARD_ACTIVATED     string = "CSTE0200"
	DEFAULT_PVV        string = "000"
	CARD_CATEGORY_MAIN string = "CRCG0800"
	CARD_CATEGORY_SUB  string = "CRCG0400"
)
