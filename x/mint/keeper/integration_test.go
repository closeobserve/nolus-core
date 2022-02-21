package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"gitlab-nomo.credissimo.net/nomo/cosmzone/app/params"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	params.SetAddressPrefixes()
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})

	// TODO
	//app.MintKeeper.SetParams(ctx, types.DefaultParams())
	//app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}
