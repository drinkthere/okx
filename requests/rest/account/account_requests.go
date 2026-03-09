package account

import "github.com/drinkthere/okx"

type (
	GetBalance struct {
		Ccy []string `json:"ccy,omitempty"`
	}
	GetPositions struct {
		InstID   []string           `json:"instId,omitempty"`
		PosID    []string           `json:"posId,omitempty"`
		InstType okx.InstrumentType `json:"instType,omitempty"`
	}
	GetAccountAndPositionRisk struct {
		InstType okx.InstrumentType `json:"instType,omitempty"`
	}
	GetBills struct {
		Ccy      string             `json:"ccy,omitempty"`
		After    int64              `json:"after,omitempty,string"`
		Before   int64              `json:"before,omitempty,string"`
		Limit    int64              `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType,omitempty"`
		MgnMode  okx.MarginMode     `json:"mgnMode,omitempty"`
		CtType   okx.ContractType   `json:"ctType,omitempty"`
		Type     okx.BillType       `json:"type,omitempty,string"`
		SubType  okx.BillSubType    `json:"subType,omitempty,string"`
	}
	SetPositionMode struct {
		PosMode okx.PositionType `json:"posMode"`
	}
	SetLeverage struct {
		Lever   int64            `json:"lever,string"`
		InstID  string           `json:"instId,omitempty"`
		Ccy     string           `json:"ccy,omitempty"`
		MgnMode okx.MarginMode   `json:"mgnMode"`
		PosSide okx.PositionSide `json:"posSide,omitempty"`
	}
	GetMaxBuySellAmount struct {
		Ccy    string        `json:"ccy,omitempty"`
		Px     float64       `json:"px,string,omitempty"`
		InstID []string      `json:"instId"`
		TdMode okx.TradeMode `json:"tdMode"`
	}
	GetMaxAvailableTradeAmount struct {
		Ccy        string        `json:"ccy,omitempty"`
		InstID     string        `json:"instId"`
		ReduceOnly bool          `json:"reduceOnly,omitempty"`
		TdMode     okx.TradeMode `json:"tdMode"`
	}
	IncreaseDecreaseMargin struct {
		InstID     string           `json:"instId"`
		Amt        float64          `json:"amt,string"`
		PosSide    okx.PositionSide `json:"posSide"`
		ActionType okx.CountAction  `json:"actionType"`
	}
	GetLeverage struct {
		InstID  []string       `json:"instId"`
		MgnMode okx.MarginMode `json:"mgnMode"`
	}
	GetMaxLoan struct {
		InstID  string         `json:"instId"`
		MgnCcy  string         `json:"mgnCcy,omitempty"`
		MgnMode okx.MarginMode `json:"mgnMode"`
	}
	GetFeeRates struct {
		InstID   string             `json:"instId,omitempty"`
		Uly      string             `json:"uly,omitempty"`
		Category okx.FeeCategory    `json:"category,omitempty,string"`
		InstType okx.InstrumentType `json:"instType"`
	}
	GetAccountInstruments struct {
		InstType okx.InstrumentType `json:"instType"`
		InstId   string             `json:"instId,omitempty"`
		Uly      string             `json:"uly,omitempty"`
	}
	GetInterestAccrued struct {
		InstID  string         `json:"instId,omitempty"`
		Ccy     string         `json:"ccy,omitempty"`
		After   int64          `json:"after,omitempty,string"`
		Before  int64          `json:"before,omitempty,string"`
		Limit   int64          `json:"limit,omitempty,string"`
		MgnMode okx.MarginMode `json:"mgnMode,omitempty"`
	}
	SetGreeks struct {
		GreeksType okx.GreekType `json:"greeksType"`
	}
	SetAutoLoan struct {
		AutoLoan bool `json:"autoLoan,omitempty"`
	}
	GetInterestLimits struct {
		Type string `json:"type,omitempty"`
		Ccy  string `json:"ccy,omitempty"`
	}
	SetAccountLevel struct {
		AcctLv string `json:"acctLv"`
	}
)
