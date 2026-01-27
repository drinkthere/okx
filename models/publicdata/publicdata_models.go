package publicdata

import "github.com/drinkthere/okx"

type (
	Instrument struct {
		InstID     string              `json:"instId"`
		Uly        string              `json:"uly,omitempty"`
		BaseCcy    string              `json:"baseCcy,omitempty"`
		QuoteCcy   string              `json:"quoteCcy,omitempty"`
		SettleCcy  string              `json:"settleCcy,omitempty"`
		CtValCcy   string              `json:"ctValCcy,omitempty"`
		CtVal      okx.JSONFloat64     `json:"ctVal,omitempty"`
		CtMult     okx.JSONFloat64     `json:"ctMult,omitempty"`
		Stk        okx.JSONFloat64     `json:"stk,omitempty"`
		TickSz     okx.JSONFloat64     `json:"tickSz,omitempty"`
		LotSz      okx.JSONFloat64     `json:"lotSz,omitempty"`
		MinSz      okx.JSONFloat64     `json:"minSz,omitempty"`
		Lever      okx.JSONFloat64     `json:"lever"`
		InstType   okx.InstrumentType  `json:"instType"`
		Category   okx.FeeCategory     `json:"category,string"`
		OptType    okx.OptionType      `json:"optType,omitempty"`
		ListTime   okx.JSONTime        `json:"listTime"`
		ExpTime    okx.JSONTime        `json:"expTime,omitempty"`
		CtType     okx.ContractType    `json:"ctType,omitempty"`
		Alias      okx.AliasType       `json:"alias,omitempty"`
		State      okx.InstrumentState `json:"state"`
		InstIDCode string              `json:"instIdCode"`
	}
	DeliveryExerciseHistory struct {
		Details []*DeliveryExerciseHistoryDetails `json:"details"`
		TS      okx.JSONTime                      `json:"ts"`
	}
	DeliveryExerciseHistoryDetails struct {
		InstID string                   `json:"instId"`
		Px     okx.JSONFloat64          `json:"px"`
		Type   okx.DeliveryExerciseType `json:"type"`
	}
	OpenInterest struct {
		InstID   string             `json:"instId"`
		Oi       okx.JSONFloat64    `json:"oi"`
		OiCcy    okx.JSONFloat64    `json:"oiCcy"`
		InstType okx.InstrumentType `json:"instType"`
		TS       okx.JSONTime       `json:"ts"`
	}
	FundingRate struct {
		InstID          string             `json:"instId"`
		InstType        okx.InstrumentType `json:"instType"`
		FundingRate     okx.JSONFloat64    `json:"fundingRate"`
		NextFundingRate okx.JSONFloat64    `json:"nextFundingRate"`
		FundingTime     okx.JSONTime       `json:"fundingTime"`
		NextFundingTime okx.JSONTime       `json:"nextFundingTime"`
	}
	FundingRateRest struct {
		InstID          string             `json:"instId"`
		InstType        okx.InstrumentType `json:"instType"`
		Method          string             `json:"method"`
		FundingRate     okx.JSONFloat64    `json:"fundingRate"`
		NextFundingRate string             `json:"nextFundingRate"`
		FundingTime     okx.JSONTime       `json:"fundingTime"`
		NextFundingTime okx.JSONTime       `json:"nextFundingTime"`
		MinFundingRate  okx.JSONFloat64    `json:"minFundingRate"`
		MaxFundingRate  okx.JSONFloat64    `json:"maxFundingRate"`
		SettState       string             `json:"settState"`
		SettFundingRate okx.JSONFloat64    `json:"settFundingRate"`
		Premium         okx.JSONFloat64    `json:"premium"`
		Ts              okx.JSONTime       `json:"ts"`
	}
	LimitPrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		BuyLmt   okx.JSONFloat64    `json:"buyLmt"`
		SellLmt  okx.JSONFloat64    `json:"sellLmt"`
		TS       okx.JSONTime       `json:"ts"`
	}
	EstimatedDeliveryExercisePrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		SettlePx okx.JSONFloat64    `json:"settlePx"`
		TS       okx.JSONTime       `json:"ts"`
	}
	OptionMarketData struct {
		InstID   string             `json:"instId"`
		Uly      string             `json:"uly"`
		InstType okx.InstrumentType `json:"instType"`
		Delta    okx.JSONFloat64    `json:"delta"`
		Gamma    okx.JSONFloat64    `json:"gamma"`
		Vega     okx.JSONFloat64    `json:"vega"`
		Theta    okx.JSONFloat64    `json:"theta"`
		DeltaBS  okx.JSONFloat64    `json:"deltaBS"`
		GammaBS  okx.JSONFloat64    `json:"gammaBS"`
		VegaBS   okx.JSONFloat64    `json:"vegaBS"`
		ThetaBS  okx.JSONFloat64    `json:"thetaBS"`
		Lever    okx.JSONFloat64    `json:"lever"`
		MarkVol  okx.JSONFloat64    `json:"markVol"`
		BidVol   okx.JSONFloat64    `json:"bidVol"`
		AskVol   okx.JSONFloat64    `json:"askVol"`
		RealVol  okx.JSONFloat64    `json:"realVol"`
		TS       okx.JSONTime       `json:"ts"`
	}
	GetDiscountRateAndInterestFreeQuota struct {
		Ccy          string          `json:"ccy"`
		Amt          okx.JSONFloat64 `json:"amt"`
		DiscountLv   okx.JSONInt64   `json:"discountLv"`
		DiscountInfo []*DiscountInfo `json:"discountInfo"`
	}
	DiscountInfo struct {
		DiscountRate okx.JSONInt64 `json:"discountRate"`
		MaxAmt       okx.JSONInt64 `json:"maxAmt"`
		MinAmt       okx.JSONInt64 `json:"minAmt"`
	}
	SystemTime struct {
		TS okx.JSONTime `json:"ts"`
	}
	LiquidationOrder struct {
		InstID    string                    `json:"instId"`
		Uly       string                    `json:"uly,omitempty"`
		InstType  okx.InstrumentType        `json:"instType"`
		TotalLoss okx.JSONFloat64           `json:"totalLoss"`
		Details   []*LiquidationOrderDetail `json:"details"`
	}
	LiquidationOrderDetail struct {
		Ccy     string           `json:"ccy,omitempty"`
		Side    okx.OrderSide    `json:"side"`
		OosSide okx.PositionSide `json:"posSide"`
		BkPx    okx.JSONFloat64  `json:"bkPx"`
		Sz      okx.JSONFloat64  `json:"sz"`
		BkLoss  okx.JSONFloat64  `json:"bkLoss"`
		TS      okx.JSONTime     `json:"ts"`
	}
	MarkPrice struct {
		InstID   string             `json:"instId"`
		InstType okx.InstrumentType `json:"instType"`
		MarkPx   okx.JSONFloat64    `json:"markPx"`
		TS       okx.JSONTime       `json:"ts"`
	}
	PositionTier struct {
		InstID       string             `json:"instId"`
		Uly          string             `json:"uly,omitempty"`
		InstType     okx.InstrumentType `json:"instType"`
		Tier         okx.JSONInt64      `json:"tier"`
		MinSz        okx.JSONFloat64    `json:"minSz"`
		MaxSz        okx.JSONFloat64    `json:"maxSz"`
		Mmr          okx.JSONFloat64    `json:"mmr"`
		Imr          okx.JSONFloat64    `json:"imr"`
		OptMgnFactor okx.JSONFloat64    `json:"optMgnFactor,omitempty"`
		QuoteMaxLoan okx.JSONFloat64    `json:"quoteMaxLoan,omitempty"`
		BaseMaxLoan  okx.JSONFloat64    `json:"baseMaxLoan,omitempty"`
		MaxLever     okx.JSONFloat64    `json:"maxLever"`
		TS           okx.JSONTime       `json:"ts"`
	}
	InterestRateAndLoanQuota struct {
		Basic   []*InterestRateAndLoanBasic `json:"basic"`
		Vip     []*InterestRateAndLoanUser  `json:"vip"`
		Regular []*InterestRateAndLoanUser  `json:"regular"`
	}
	InterestRateAndLoanBasic struct {
		Ccy   string          `json:"ccy"`
		Rate  okx.JSONFloat64 `json:"rate"`
		Quota okx.JSONFloat64 `json:"quota"`
	}
	InterestRateAndLoanUser struct {
		Level         string          `json:"level"`
		IrDiscount    okx.JSONFloat64 `json:"irDiscount"`
		LoanQuotaCoef int             `json:"loanQuotaCoef,string"`
	}
	State struct {
		Title       string       `json:"title"`
		State       string       `json:"state"`
		Href        string       `json:"href"`
		ServiceType string       `json:"serviceType"`
		System      string       `json:"system"`
		ScheDesc    string       `json:"scheDesc"`
		Begin       okx.JSONTime `json:"begin"`
		End         okx.JSONTime `json:"end"`
	}
	UnitConvert struct {
		Type   okx.ConvertType `json:"type,string,omitempty"`
		InstID string          `json:"instId"`
		Px     okx.JSONFloat64 `json:"px,omitempty"`
		Sz     okx.JSONFloat64 `json:"sz,omitempty"`
		Unit   string          `json:"unit,omitempty"`
	}
)
