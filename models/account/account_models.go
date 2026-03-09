package account

import "github.com/drinkthere/okx"

type (
	Balance struct {
		TotalEq     okx.JSONFloat64   `json:"totalEq"`
		IsoEq       okx.JSONFloat64   `json:"isoEq"`
		AdjEq       okx.JSONFloat64   `json:"adjEq,omitempty"`
		OrdFroz     okx.JSONFloat64   `json:"ordFroz,omitempty"`
		Imr         okx.JSONFloat64   `json:"imr,omitempty"`
		Mmr         okx.JSONFloat64   `json:"mmr,omitempty"`
		MgnRatio    okx.JSONFloat64   `json:"mgnRatio,omitempty"`
		NotionalUsd okx.JSONFloat64   `json:"notionalUsd,omitempty"`
		Details     []*BalanceDetails `json:"details,omitempty"`
		UTime       okx.JSONTime      `json:"uTime"`
	}
	BalancesFunding struct {
		Ccy       string          `json:"ccy"`
		Bal       okx.JSONFloat64 `json:"bal"`
		FrozenBal okx.JSONFloat64 `json:"frozenBal"`
		AvailBal  okx.JSONFloat64 `json:"availBal"`
	}
	BalanceDetails struct {
		Ccy           string          `json:"ccy"`
		Eq            okx.JSONFloat64 `json:"eq"`
		CashBal       okx.JSONFloat64 `json:"cashBal"`
		IsoEq         okx.JSONFloat64 `json:"isoEq,omitempty"`
		AvailEq       okx.JSONFloat64 `json:"availEq,omitempty"`
		DisEq         okx.JSONFloat64 `json:"disEq"`
		AvailBal      okx.JSONFloat64 `json:"availBal"`
		FrozenBal     okx.JSONFloat64 `json:"frozenBal"`
		OrdFrozen     okx.JSONFloat64 `json:"ordFrozen"`
		Liab          okx.JSONFloat64 `json:"liab,omitempty"`
		Upl           okx.JSONFloat64 `json:"upl,omitempty"`
		UplLib        okx.JSONFloat64 `json:"uplLib,omitempty"`
		CrossLiab     okx.JSONFloat64 `json:"crossLiab,omitempty"`
		IsoLiab       okx.JSONFloat64 `json:"isoLiab,omitempty"`
		MgnRatio      okx.JSONFloat64 `json:"mgnRatio,omitempty"`
		Interest      okx.JSONFloat64 `json:"interest,omitempty"`
		Twap          okx.JSONFloat64 `json:"twap,omitempty"`
		MaxLoan       okx.JSONFloat64 `json:"maxLoan,omitempty"`
		EqUsd         okx.JSONFloat64 `json:"eqUsd"`
		NotionalLever okx.JSONFloat64 `json:"notionalLever,omitempty"`
		StgyEq        okx.JSONFloat64 `json:"stgyEq"`
		IsoUpl        okx.JSONFloat64 `json:"isoUpl,omitempty"`
		UTime         okx.JSONTime    `json:"uTime"`
	}
	Position struct {
		InstID      string             `json:"instId"`
		PosCcy      string             `json:"posCcy,omitempty"`
		LiabCcy     string             `json:"liabCcy,omitempty"`
		OptVal      string             `json:"optVal,omitempty"`
		Ccy         string             `json:"ccy"`
		PosID       string             `json:"posId"`
		TradeID     string             `json:"tradeId"`
		Pos         okx.JSONFloat64    `json:"pos"`
		AvailPos    okx.JSONFloat64    `json:"availPos,omitempty"`
		AvgPx       okx.JSONFloat64    `json:"avgPx"`
		BePx        okx.JSONFloat64    `json:"bePx"`
		MarkPx      okx.JSONFloat64    `json:"markPx"`
		IdxPx       okx.JSONFloat64    `json:"idxPx"`
		UsdPx       okx.JSONFloat64    `json:"usdPx"`
		Upl         okx.JSONFloat64    `json:"upl"`
		UplRatio    okx.JSONFloat64    `json:"uplRatio"`
		Lever       okx.JSONFloat64    `json:"lever"`
		LiqPx       okx.JSONFloat64    `json:"liqPx,omitempty"`
		Imr         okx.JSONFloat64    `json:"imr,omitempty"`
		Margin      okx.JSONFloat64    `json:"margin,omitempty"`
		MgnRatio    okx.JSONFloat64    `json:"mgnRatio"`
		Mmr         okx.JSONFloat64    `json:"mmr"`
		Liab        okx.JSONFloat64    `json:"liab,omitempty"`
		Interest    okx.JSONFloat64    `json:"interest"`
		NotionalUsd okx.JSONFloat64    `json:"notionalUsd"`
		ADL         okx.JSONFloat64    `json:"adl"`
		Last        okx.JSONFloat64    `json:"last"`
		DeltaBS     okx.JSONFloat64    `json:"deltaBS"`
		DeltaPA     okx.JSONFloat64    `json:"deltaPA"`
		GammaBS     okx.JSONFloat64    `json:"gammaBS"`
		GammaPA     okx.JSONFloat64    `json:"gammaPA"`
		ThetaBS     okx.JSONFloat64    `json:"thetaBS"`
		ThetaPA     okx.JSONFloat64    `json:"thetaPA"`
		VegaBS      okx.JSONFloat64    `json:"vegaBS"`
		VegaPA      okx.JSONFloat64    `json:"vegaPA"`
		PosSide     okx.PositionSide   `json:"posSide"`
		MgnMode     okx.MarginMode     `json:"mgnMode"`
		InstType    okx.InstrumentType `json:"instType"`
		CTime       okx.JSONTime       `json:"cTime"`
		UTime       okx.JSONTime       `json:"uTime"`
	}
	BalanceAndPosition struct {
		EventType okx.EventType     `json:"eventType"`
		PTime     okx.JSONTime      `json:"pTime"`
		UTime     okx.JSONTime      `json:"uTime"`
		PosData   []*Position       `json:"posData"`
		BalData   []*BalanceDetails `json:"balData"`
	}
	PositionAndAccountRisk struct {
		AdjEq   okx.JSONFloat64                      `json:"adjEq,omitempty"`
		BalData []*PositionAndAccountRiskBalanceData `json:"balData"`
		PosData []*PositionAndAccountRiskBalanceData `json:"posData"`
		TS      okx.JSONTime                         `json:"ts"`
	}
	PositionAndAccountRiskBalanceData struct {
		Ccy   string          `json:"ccy"`
		Eq    okx.JSONFloat64 `json:"eq"`
		DisEq okx.JSONFloat64 `json:"disEq"`
	}
	PositionAndAccountRiskPositionData struct {
		InstID      string             `json:"instId"`
		PosCcy      string             `json:"posCcy,omitempty"`
		Ccy         string             `json:"ccy"`
		NotionalCcy okx.JSONFloat64    `json:"notionalCcy"`
		Pos         okx.JSONFloat64    `json:"pos"`
		NotionalUsd okx.JSONFloat64    `json:"notionalUsd"`
		PosSide     okx.PositionSide   `json:"posSide"`
		InstType    okx.InstrumentType `json:"instType"`
		MgnMode     okx.MarginMode     `json:"mgnMode"`
	}
	Bill struct {
		Ccy       string             `json:"ccy"`
		InstID    string             `json:"instId"`
		Notes     string             `json:"notes"`
		BillID    string             `json:"billId"`
		OrdID     string             `json:"ordId"`
		BalChg    okx.JSONFloat64    `json:"balChg"`
		PosBalChg okx.JSONFloat64    `json:"posBalChg"`
		Bal       okx.JSONFloat64    `json:"bal"`
		PosBal    okx.JSONFloat64    `json:"posBal"`
		Sz        okx.JSONFloat64    `json:"sz"`
		Pnl       okx.JSONFloat64    `json:"pnl"`
		Fee       okx.JSONFloat64    `json:"fee"`
		From      okx.AccountType    `json:"from,omitempty"`
		To        okx.AccountType    `json:"to,omitempty"`
		InstType  okx.InstrumentType `json:"instType"`
		MgnMode   okx.MarginMode     `json:"MgnMode"`
		Type      okx.BillType       `json:"type,string"`
		SubType   okx.BillSubType    `json:"subType,string"`
		TS        okx.JSONTime       `json:"ts"`
	}
	Config struct {
		Level       string           `json:"level"`
		LevelTmp    string           `json:"levelTmp"`
		AcctLv      string           `json:"acctLv"`
		AutoLoan    bool             `json:"autoLoan"`
		UID         string           `json:"uid"`
		MainUID     string           `json:"mainUid"`
		Label       string           `json:"label"`
		IP          string           `json:"ip"`
		Permissions string           `json:"perm"`
		GreeksType  okx.GreekType    `json:"greeksType"`
		PosMode     okx.PositionType `json:"posMode"`
	}
	PositionMode struct {
		PosMode okx.PositionType `json:"posMode"`
	}
	Leverage struct {
		InstID  string           `json:"instId"`
		Lever   okx.JSONFloat64  `json:"lever"`
		MgnMode okx.MarginMode   `json:"mgnMode"`
		PosSide okx.PositionSide `json:"posSide"`
	}
	MaxBuySellAmount struct {
		InstID  string          `json:"instId"`
		Ccy     string          `json:"ccy"`
		MaxBuy  okx.JSONFloat64 `json:"maxBuy"`
		MaxSell okx.JSONFloat64 `json:"maxSell"`
	}
	MaxAvailableTradeAmount struct {
		InstID    string          `json:"instId"`
		AvailBuy  okx.JSONFloat64 `json:"availBuy"`
		AvailSell okx.JSONFloat64 `json:"availSell"`
	}
	MarginBalanceAmount struct {
		InstID  string           `json:"instId"`
		Amt     okx.JSONFloat64  `json:"amt"`
		PosSide okx.PositionSide `json:"posSide,string"`
		Type    okx.CountAction  `json:"type,string"`
	}
	Loan struct {
		InstID  string          `json:"instId"`
		MgnCcy  string          `json:"mgnCcy"`
		Ccy     string          `json:"ccy"`
		MaxLoan okx.JSONFloat64 `json:"maxLoan"`
		MgnMode okx.MarginMode  `json:"mgnMode"`
		Side    okx.OrderSide   `json:"side"`
	}
	FeeGroupItem struct {
		GroupId   string          `json:"groupId"`
		Maker     okx.JSONFloat64 `json:"maker"`
		Taker     okx.JSONFloat64 `json:"taker"`
		MakerU    okx.JSONFloat64 `json:"makerU"`
		TakerU    okx.JSONFloat64 `json:"takerU"`
		MakerUSDC okx.JSONFloat64 `json:"makerUSDC"`
		TakerUSDC okx.JSONFloat64 `json:"takerUSDC"`
	}
	Fee struct {
		Level     string             `json:"level"`
		Taker     okx.JSONFloat64    `json:"taker"`
		Maker     okx.JSONFloat64    `json:"maker"`
		Delivery  okx.JSONFloat64    `json:"delivery,omitempty"`
		Exercise  okx.JSONFloat64    `json:"exercise,omitempty"`
		Category  okx.FeeCategory    `json:"category,string"`
		InstType  okx.InstrumentType `json:"instType"`
		TakerU    okx.JSONFloat64    `json:"takerU"`
		MakerU    okx.JSONFloat64    `json:"makerU"`
		TakerUSDC okx.JSONFloat64    `json:"takerUSDC"`
		MakerUSDC okx.JSONFloat64    `json:"makerUSDC"`
		FeeGroup  []*FeeGroupItem    `json:"feeGroup,omitempty"`
		TS        okx.JSONTime       `json:"ts"`
	}
	AccountInstrument struct {
		InstType okx.InstrumentType `json:"instType"`
		InstId   string             `json:"instId"`
		GroupId  string             `json:"groupId"`
	}
	InterestAccrued struct {
		InstID       string          `json:"instId"`
		Ccy          string          `json:"ccy"`
		Interest     okx.JSONFloat64 `json:"interest"`
		InterestRate okx.JSONFloat64 `json:"interestRate"`
		Liab         okx.JSONFloat64 `json:"liab"`
		MgnMode      okx.MarginMode  `json:"mgnMode"`
		TS           okx.JSONTime    `json:"ts"`
	}
	InterestRate struct {
		Ccy          string          `json:"ccy"`
		InterestRate okx.JSONFloat64 `json:"interestRate"`
	}
	Greek struct {
		GreeksType string `json:"greeksType"`
	}
	MaxWithdrawal struct {
		Ccy   string          `json:"ccy"`
		MaxWd okx.JSONFloat64 `json:"maxWd"`
	}
	AutoLoan struct {
		AutoLoan bool `json:"autoLoan"`
	}
	AcctLevel struct {
		AcctLv string `json:"acctLv"`
	}
	InterestLimitsRecordDetail struct {
		AllAcctRemainingQuota okx.JSONFloat64 `json:"allAcctRemainingQuota"`
		CurAcctRemainingQuota okx.JSONFloat64 `json:"curAcctRemainingQuota,omitempty"`
		PlatRemainingQuota    string          `json:"platRemainingQuota"`
	}
	InterestLimitsRecord struct {
		Ccy               string                     `json:"ccy"`
		Rate              okx.JSONFloat64            `json:"rate"`
		LoanQuota         okx.JSONFloat64            `json:"loanQuota"`
		SurplusLmt        okx.JSONFloat64            `json:"surplusLmt"`
		SurplusLmtDetails InterestLimitsRecordDetail `json:"surplusLmtDetails,omitempty"`
		UsedLmt           okx.JSONFloat64            `json:"usedLmt"`
		Interest          okx.JSONFloat64            `json:"interest,omitempty"`
		PosLoan           okx.JSONFloat64            `json:"posLoan,omitempty"`
		AvailLoan         okx.JSONFloat64            `json:"availLoan,omitempty"`
		UsedLoan          okx.JSONFloat64            `json:"usedLoan,omitempty"`
		AvgRate           okx.JSONFloat64            `json:"avgRate,omitempty"`
	}
	InterestLimits struct {
		Debt             okx.JSONFloat64         `json:"debt"`
		Interest         okx.JSONFloat64         `json:"interest,omitempty"`
		NextDiscountTime okx.JSONTime            `json:"nextDiscountTime"`
		NextInterestTime okx.JSONTime            `json:"nextInterestTime"`
		LoanAlloc        string                  `json:"loanAlloc"`
		Records          []*InterestLimitsRecord `json:"records"`
	}
)
