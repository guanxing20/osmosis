package v13

import (
	"context"
	"embed"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"

	"github.com/osmosis-labs/osmosis/v30/app/keepers"
	"github.com/osmosis-labs/osmosis/v30/app/upgrades"
	ibcratelimittypes "github.com/osmosis-labs/osmosis/v30/x/ibc-rate-limit/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v30/x/lockup/types"
)

//go:embed rate_limiter.wasm
var embedFs embed.FS

func setupRateLimiting(ctx sdk.Context, keepers *keepers.AppKeepers) error {
	govModule := keepers.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	code, err := embedFs.ReadFile("rate_limiter.wasm")
	if err != nil {
		return err
	}
	contractKeeper := wasmkeeper.NewGovPermissionKeeper(keepers.WasmKeeper)
	instantiateConfig := wasmtypes.AccessConfig{Permission: wasmtypes.AccessTypeAnyOfAddresses, Addresses: []string{govModule.String()}}
	codeID, _, err := contractKeeper.Create(ctx, govModule, code, &instantiateConfig)
	if err != nil {
		return err
	}

	transferModule := keepers.AccountKeeper.GetModuleAddress(transfertypes.ModuleName)

	initMsgBz := []byte(fmt.Sprintf(`{
           "gov_module":  "%s",
           "ibc_module":"%s",
           "paths": []
        }`,
		govModule, transferModule))

	addr, _, err := contractKeeper.Instantiate(ctx, codeID, govModule, govModule, initMsgBz, "rate limiting contract", nil)
	if err != nil {
		return err
	}
	addrStr, err := sdk.Bech32ifyAddressBytes("osmo", addr)
	if err != nil {
		return err
	}
	params, err := ibcratelimittypes.NewParams(addrStr)
	if err != nil {
		return err
	}
	paramSpace, ok := keepers.ParamsKeeper.GetSubspace(ibcratelimittypes.ModuleName)
	if !ok {
		//nolint:staticcheck
		return errorsmod.New("rate-limiting-upgrades", 2, "can't create paramspace")
	}
	paramSpace.SetParamSet(ctx, &params)
	return nil
}

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)
		keepers.LockupKeeper.SetParams(ctx, lockuptypes.DefaultParams())
		if err := setupRateLimiting(ctx, keepers); err != nil {
			return nil, err
		}
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
