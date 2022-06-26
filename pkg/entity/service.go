package entity

// Service type
type Service struct {
	ID            string        `xml:"id,attr" faker:"-" json:"id"`
	ServiceNumber string        `xml:"service_number" faker:"-" json:"service_number"`
	Value         string        `xml:"value,attr" faker:"-" json:"value"`
	ServiceObject ServiceObject `xml:"service_object" faker:"-" json:"service_object"`
	Error         Error         `xml:"error" faker:"-" json:"error"`
}

// ServiceObject type

type ServiceObject struct {
	RefID          string         `xml:"ref_id,attr" faker:"-" json:"ref_id"`
	StartDate      string         `xml:"start_date" faker:"-" json:"start_date"`
	EndDate        string         `xml:"end_date" faker:"-" json:"end_date"`
	AttributeNum   AttributeNum   `xml:"attribute_num" faker:"-" json:"attribute_num"`
	AttributeChar  AttributeChar  `xml:"attribute_char" faker:"-" json:"attribute_char"`
	AttributeDate  AttributeDate  `xml:"attribute_date" faker:"-" json:"attribute_date"`
	AttributeFee   AttributeFee   `xml:"attribute_fee" faker:"-" json:"attribute_fee"`
	AttributeCycle AttributeCycle `xml:"attribute_cycle" faker:"-" json:"attribute_cycle"`
	AttributeLimit AttributeLimit `xml:"attribute_limit" faker:"-" json:"attribute_limit"`
	Error          Error          `xml:"error" faker:"-" json:"error"`
}

type AttributeChar struct {
	Text               string `xml:",chardata" json:"text"`
	Value              string `xml:"value,attr" json:"value"`
	ID                 string `xml:"id,attr" json:"id"`
	AttributeValueChar string `xml:"attribute_value_char" json:"attribute_value_char"`
	StartDate          string `xml:"start_date" json:"start_date"`
	EndDate            string `xml:"end_date" json:"end_date"`
	ModCondition       string `xml:"mod_condition" json:"mod_condition"`
	ModName            string `xml:"mod_name" json:"mod_name"`
	ModID              string `xml:"mod_id" json:"mod_id"`
	Error              Error  `xml:"error" json:"error"`
}

type AttributeDate struct {
	Value              string `xml:"value,attr" json:"value"`
	ID                 string `xml:"id,attr" json:"id"`
	AttributeValueDate string `xml:"attribute_value_date" json:"attribute_value_date"`
	StartDate          string `xml:"start_date" json:"start_date"`
	EndDate            string `xml:"end_date" json:"end_date"`
	ModCondition       string `xml:"mod_condition" json:"mod_condition"`
	ModName            string `xml:"mod_name" json:"mod_name"`
	ModID              string `xml:"mod_id" json:"mod_id"`
	Error              Error  `xml:"error" json:"error"`
}

type FeeTier struct {
	SumThreshold    string `xml:"sum_threshold" json:"sum_threshold"`
	CountThreshold  string `xml:"count_threshold" json:"count_threshold"`
	FeeFixedValue   string `xml:"fee_fixed_value" json:"fee_fixed_value"`
	FeePercentValue string `xml:"fee_percent_value" json:"fee_percent_value"`
	FeeMinValue     string `xml:"fee_min_value" json:"fee_min_value"`
	FeeMaxValue     string `xml:"fee_max_value" json:"fee_max_value"`
}

type AttributeFee struct {
	Value               string         `xml:"value,attr" json:"value"`
	ID                  string         `xml:"id,attr" json:"id"`
	FeeRateCalc         string         `xml:"fee_rate_calc" json:"fee_rate_calc"`
	FeeBaseCalc         string         `xml:"fee_base_calc" json:"fee_base_calc"`
	FeeFixedValue       string         `xml:"fee_fixed_value" json:"fee_fixed_value"`
	FeePercentValue     string         `xml:"fee_percent_value" json:"fee_percent_value"`
	FeeMinValue         string         `xml:"fee_min_value" json:"fee_min_value"`
	FeeMaxValue         string         `xml:"fee_max_value" json:"fee_max_value"`
	Currency            string         `xml:"currency" json:"currency"`
	StartDate           string         `xml:"start_date" json:"start_date"`
	EndDate             string         `xml:"end_date" json:"end_date"`
	CycleLengthType     string         `xml:"cycle_length_type" json:"cycle_length_type"`
	ModCondition        string         `xml:"mod_condition" json:"mod_condition"`
	ModName             string         `xml:"mod_name" json:"mod_name"`
	ModID               string         `xml:"mod_id" json:"mod_id"`
	LengthTypeAlgorithm string         `xml:"length_type_algorithm" json:"length_type_algorithm"`
	FeeTier             FeeTier        `xml:"fee_tier" json:"fee_tier"`
	Cycle               AttributeCycle `xml:"cycle" json:"cycle"`
	Error               Error          `xml:"error" json:"error"`
}

type Shift struct {
	ShiftType       string `xml:"shift_type" json:"shift_type"`
	ShiftPriority   string `xml:"shift_priority" json:"shift_priority"`
	ShiftSign       string `xml:"shift_sign" json:"shift_sign"`
	ShiftLengthType string `xml:"shift_length_type" json:"shift_length_type"`
	ShiftLength     string `xml:"shift_length" json:"shift_length"`
}

type AttributeCycle struct {
	Value           string `xml:"value,attr" json:"value"`
	ID              string `xml:"id,attr" json:"id"`
	CycleStartDate  string `xml:"cycle_start_date" json:"cycle_start_date"`
	CycleLengthType string `xml:"cycle_length_type" json:"cycle_length_type"`
	CycleLength     string `xml:"cycle_length" json:"cycle_length"`
	CycleExclType   string `xml:"cycle_excl_type" json:"cycle_excl_type"`
	Shift           Shift  `xml:"shift" json:"shift"`
	StartDate       string `xml:"start_date" json:"start_date"`
	EndDate         string `xml:"end_date" json:"end_date"`
	ModCondition    string `xml:"mod_condition" json:"mod_condition"`
	ModName         string `xml:"mod_name" json:"mod_name"`
	ModID           string `xml:"mod_id" json:"mod_id"`
	Error           Error  `xml:"error" json:"error"`
}

type AttributeLimit struct {
	Value            string `xml:"value,attr" json:"value"`
	ID               string `xml:"id,attr" json:"id"`
	LimitCountValue  string `xml:"limit_count_value" json:"limit_count_value"`
	LimitSumValue    string `xml:"limit_sum_value" json:"limit_sum_value"`
	LimitCheckType   string `xml:"limit_check_type" json:"limit_check_type"`
	Currency         string `xml:"currency" json:"currency"`
	CounterAlgorithm string `xml:"counter_algorithm" json:"counter_algorithm"`
	StartDate        string `xml:"start_date" json:"start_date"`
	EndDate          string `xml:"end_date" json:"end_date"`
	ModCondition     string `xml:"mod_condition" json:"mod_condition"`
	ModName          string `xml:"mod_name" json:"mod_name"`
	ModID            string `xml:"mod_id" json:"mod_id"`
	Error            Error  `xml:"error" json:"error"`
}

type AttributeNum struct {
	Value             string `xml:"value,attr" json:"value"`
	ID                string `xml:"id,attr" json:"id"`
	AttributeValueNum string `xml:"attribute_value_num" json:"attribute_value_num"`
	StartDate         string `xml:"start_date" json:"start_date"`
	EndDate           string `xml:"end_date" json:"end_date"`
	ModCondition      string `xml:"mod_condition" json:"mod_condition"`
	ModName           string `xml:"mod_name" json:"mod_name"`
	ModID             string `xml:"mod_id" json:"mod_id"`
	Error             Error  `xml:"error" json:"error"`
}
