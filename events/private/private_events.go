package private

import (
	"github.com/drinkthere/okx/events"
	"github.com/drinkthere/okx/models/account"
	"github.com/drinkthere/okx/models/trade"
)

type (
	Account struct {
		Arg      *events.Argument   `json:"arg"`
		Balances []*account.Balance `json:"data"`
	}
	Position struct {
		Arg       *events.Argument    `json:"arg"`
		Positions []*account.Position `json:"data"`
	}
	BalanceAndPosition struct {
		Arg                 *events.Argument              `json:"arg"`
		BalanceAndPositions []*account.BalanceAndPosition `json:"data"`
	}
	Order struct {
		Arg    *events.Argument `json:"arg"`
		Orders []*trade.Order   `json:"data"`
	}
	AlgoOrder struct {
		Arg    *events.Argument   `json:"arg"`
		Orders []*trade.AlgoOrder `json:"data"`
	}
)
