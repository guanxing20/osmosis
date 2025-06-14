package v5

import (
	"encoding/csv"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	gammkeeper "github.com/osmosis-labs/osmosis/v30/x/gamm/keeper"
	"github.com/osmosis-labs/osmosis/v30/x/txfees/types"
)

// Every asset with a liquid osmo pairing pool on Osmosis, as of 12/01/21
// Notably, Tick is not on this list because the osmo pool has $76 of liquidity.
// Cheq'd and KRT are also not on this, due to neither having osmo pairings.
// We nolint because these are strings of whitelisted ibc denoms.
//
//nolint:gosec
var feetoken_whitelist_data = `
ion,uion,2
atom,ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2,1
akt,ibc/1480B8FD20AD5FCAE81EA87584D269547DD4D436843C1D20F15E00EB64743EF4,3
xprt,ibc/A0CC0CF735BFB30E730C70019D4218A1244FF383503FF7579C9201AB93CA9293,15
iris,ibc/7C4D60AA95E5A7558B0A364860979CA34B7FF8AAF255B87AF9E879374470CEC0,7
dvpn,ibc/9712DBB13B9631EDFA9BF61B55F1B2D290B2ADB67E3A4EB3A875F3B6081B3B84,5
cro,ibc/E6931F78057F7CC5DA0FD6CEF82FF39373A6E0452BF1FD76910B93292CF356C1,9
regen,ibc/1DCC8A6CB5689018431323953344A9F6CC4D0BFB261E88C9F7777372C10CD076,42
iov,ibc/52B1AA623B34EB78FD767CEA69E8D7FA6C9CFE1FBF49C5406268FD325E2CC2AC,197
ngm,ibc/1DC495FCEFDA068A3820F903EDBD78B942FBD204D7E93D3BA2B432E9669D1A59,463
eeur,ibc/5973C068568365FFF40DEDCF1A1CB7582B6116B731CD31A12231AE25E20B871F,481
bcna,ibc/D805F1DA50D31B96E4282C1D4181EDDFB1A44A598BFF5666F4B43E4B8BEA95A5,571
juno,ibc/46B44899322F3CD854D2D46DEEF881958467CDD4B3B10086DA49296BBED94BED,497
ixo,ibc/F3FF7A84A73B62921538642F9797C423D2B4C4ACB3C7FCFFCE7F12AA69909C4B,557
like,ibc/9989AD6CCA39D1131523DB0617B50F6442081162294B4795E26746292467B525,553
luna,ibc/0EF15DF2F02480ADE0BB6E85D9EBB5DAEA2836D3860E9F97F9AADE4F57A31AA0,561
ust,ibc/BE1BB42D4BE3C30D50B68D7C41DB4DFCE9678E8EF8C539F6E6A9345048894FCC,560
btsg,ibc/4E5444C35610CC76FC94E7F7886B93121175C28262DDFDDE6F84E82BF2425452,573
xki,ibc/B547DC9B897E7C3AA5B824696110B8E3D2C31E3ED3F02FF363DCBAD82457E07E,577
scrt,ibc/0954E1C28EB7AF5B72D24F3BC2B47BBB2FDF91BDDFD57B74B99E133AED40972A,584
med,ibc/3BCCC93AD5DF58D11A6F8A05FA8BC801CBA0BA61A981F57E91B8B598BF8061CB,586
boot,ibc/FE2CD1E6828EC0FAB8AF39BAC45BC25B965BA67CCBC50C13A14BD610B0D1E2C4,597
`

func InitialWhitelistedFeetokens(ctx sdk.Context, gamm *gammkeeper.Keeper) []types.FeeToken {
	r := csv.NewReader(strings.NewReader(feetoken_whitelist_data))
	assets, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	feeTokens := make([]types.FeeToken, 0, len(assets))
	for _, asset := range assets {
		base10 := 10
		bitLen := 64
		poolIdStr := strings.TrimSpace(asset[2])
		poolId, err := strconv.ParseUint(poolIdStr, base10, bitLen)
		if err != nil {
			panic(err)
		}

		pool, poolExistsErr := gamm.GetPoolAndPoke(ctx, poolId)
		_ = pool
		if poolExistsErr != nil {
			continue
		}

		feeToken := types.FeeToken{
			Denom:  asset[1],
			PoolID: poolId,
		}

		feeTokens = append(feeTokens, feeToken)
	}

	return feeTokens
}
