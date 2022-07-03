package entity

import "time"

type TestCase struct {
	Id        int64         `json:"id"`
	Name      string        `json:"name"`
	Status    string        `json:"status"`
	Processes []CaseProcess `json:"processes"`
	Result    []CaseResult  `json:"result"`
}

type CaseProcess struct {
	Id      int64                  `json:"id"`
	Name    string                 `json:"name"`
	Order   int                    `json:"status"`
	Param   map[string]interface{} `json:"param"`
	EffDate time.Time              `json:"eff_date"`
}

type CaseResult struct {
	ResultType string `json:"result_type"`
	Expected   string `json:"expected"`
	Result     string `json:"result"`
}
