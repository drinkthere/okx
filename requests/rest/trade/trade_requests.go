package trade

import (
	"github.com/drinkthere/okx"
	"github.com/drinkthere/okx/models/trade"
)

type (
	PlaceOrder struct {
		ID              string           `json:"-"`
		InstID          string           `json:"instId"`
		Ccy             string           `json:"ccy,omitempty"`
		ClOrdID         string           `json:"clOrdId,omitempty"`
		Tag             string           `json:"tag,omitempty"`
		QuickMgnType    string           `json:"quickMgnType,omitempty"`
		ReduceOnly      bool             `json:"reduceOnly,string,omitempty"`
		Sz              float64          `json:"sz,string"`
		Px              float64          `json:"px,omitempty,string"`
		TdMode          okx.TradeMode    `json:"tdMode"`
		Side            okx.OrderSide    `json:"side"`
		PosSide         okx.PositionSide `json:"posSide,omitempty"`
		OrdType         okx.OrderType    `json:"ordType"`
		TgtCcy          okx.QuantityType `json:"tgtCcy,omitempty"`
		TpTriggerPx     float64          `json:"tpTriggerPx,string,omitempty"`
		TpOrdPx         float64          `json:"tpOrdPx,string,omitempty"`
		TpTriggerPxType string           `json:"tpTriggerPxType,omitempty"`
		SlTriggerPx     float64          `json:"slTriggerPx,string,omitempty"`
		SlOrdPx         float64          `json:"slOrdPx,string,omitempty"`
		SlTriggerPxType string           `json:"slTriggerPxType,omitempty"`
	}
	CancelOrder struct {
		ID      string `json:"-"`
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	AmendOrder struct {
		ID        string  `json:"-"`
		InstID    string  `json:"instId"`
		OrdID     string  `json:"ordId,omitempty"`
		ClOrdID   string  `json:"clOrdId,omitempty"`
		ReqID     string  `json:"reqId,omitempty"`
		NewSz     float64 `json:"newSz,omitempty,string"`
		NewPx     float64 `json:"newPx,omitempty,string"`
		CxlOnFail bool    `json:"cxlOnFail,omitempty"`
	}
	ClosePosition struct {
		InstID  string           `json:"instId"`
		Ccy     string           `json:"ccy,omitempty"`
		PosSide okx.PositionSide `json:"posSide,omitempty"`
		MgnMode okx.MarginMode   `json:"mgnMode"`
	}
	OrderDetails struct {
		InstID  string `json:"instId"`
		OrdID   string `json:"ordId,omitempty"`
		ClOrdID string `json:"clOrdId,omitempty"`
	}
	OrderList struct {
		Uly      string             `json:"uly,omitempty"`
		InstID   string             `json:"instId,omitempty"`
		After    float64            `json:"after,omitempty,string"`
		Before   float64            `json:"before,omitempty,string"`
		Limit    float64            `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType,omitempty"`
		OrdType  okx.OrderType      `json:"ordType,omitempty"`
		State    okx.OrderState     `json:"state,omitempty"`
	}
	TransactionDetails struct {
		Uly      string             `json:"uly,omitempty"`
		InstID   string             `json:"instId,omitempty"`
		OrdID    string             `json:"ordId,omitempty"`
		After    float64            `json:"after,omitempty,string"`
		Before   float64            `json:"before,omitempty,string"`
		Limit    float64            `json:"limit,omitempty,string"`
		InstType okx.InstrumentType `json:"instType,omitempty"`
	}
	PlaceAlgoOrder struct {
		InstID        string            `json:"instId"`
		TdMode        okx.TradeMode     `json:"tdMode"`
		ClOrdID       string            `json:"clOrdId,omitempty"`
		Tag           string            `json:"tag,omitempty"`
		Ccy           string            `json:"ccy,omitempty"`
		Side          okx.OrderSide     `json:"side"`
		PosSide       okx.PositionSide  `json:"posSide,omitempty"`
		OrdType       okx.AlgoOrderType `json:"ordType"`
		Sz            float64           `json:"sz,string,omitempty"`
		ReduceOnly    bool              `json:"reduceOnly,omitempty"`
		QuickMgnType  string            `json:"quickMgnType,omitempty"`
		TgtCcy        okx.QuantityType  `json:"tgtCcy,omitempty"`
		CloseFraction string            `json:"closeFraction,omitempty"`
		StopOrder
		TriggerOrder
		IcebergOrder
		TWAPOrder
		TrailingStopOrder
	}
	StopOrder struct {
		TpTriggerPx     float64 `json:"tpTriggerPx,string,omitempty"`
		TpOrdPx         float64 `json:"tpOrdPx,string,omitempty"`
		TpTriggerPxType string  `json:"tpTriggerPxType,omitempty"`
		SlTriggerPx     float64 `json:"slTriggerPx,string,omitempty"`
		SlOrdPx         float64 `json:"slOrdPx,string,omitempty"`
		SlTriggerPxType string  `json:"slTriggerPxType,omitempty"`
	}
	TriggerOrder struct {
		TriggerPx     float64 `json:"triggerPx,string,omitempty"`
		TriggerPxType string  `json:"triggerPxType,omitempty"`
		OrdPx         float64 `json:"orderPx,string,omitempty"`
	}
	TrailingStopOrder struct {
		CallbackRatio  float64 `json:"callbackRatio,string,omitempty"`
		CallbackSpread float64 `json:"callbackSpread,string,omitempty"`
		ActivePx       float64 `json:"activePx,string,omitempty"`
	}
	IcebergOrder struct {
		PxVar    float64 `json:"pxVar,string,omitempty"`
		PxSpread float64 `json:"pxSpread,string,omitempty"`
		SzLimit  float64 `json:"szLimit,string,omitempty"`
		PxLimit  float64 `json:"pxLimit,string,omitempty"`
	}
	TWAPOrder struct {
		IcebergOrder
		TimeInterval string `json:"timeInterval,omitempty"`
	}
	CancelAlgoOrder struct {
		InstID string `json:"instId"`
		AlgoID string `json:"algoId"`
	}
	AlgoOrderList struct {
		InstType okx.InstrumentType `json:"instType,omitempty"`
		Uly      string             `json:"uly,omitempty"`
		AlgoID   string             `json:"algoId,omitempty"`
		InstID   string             `json:"instId,omitempty"`
		ClOrdID  string             `json:"clOrdId,omitempty"`
		After    float64            `json:"after,omitempty,string"`
		Before   float64            `json:"before,omitempty,string"`
		Limit    float64            `json:"limit,omitempty,string"`
		OrdType  okx.AlgoOrderType  `json:"ordType,omitempty"`
		State    okx.OrderState     `json:"state,omitempty"`
	}

	EasyConvertCurrencyList struct {
		Source trade.EasyConvertSource `json:"source,omitempty"`
	}

	EasyConvert struct {
		FromCcy []string                `json:"fromCcy"`
		ToCcy   string                  `json:"toCcy"`
		Source  trade.EasyConvertSource `json:"source,omitempty"`
	}
)
