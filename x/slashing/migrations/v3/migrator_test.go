package v3_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	"github.com/cosmos/cosmos-sdk/testutil"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
// 	"github.com/iqlusioninc/liquidity-staking-module/x/slashing"
// 	"github.com/iqlusioninc/liquidity-staking-module/x/slashing/exported"
// 	v3 "github.com/iqlusioninc/liquidity-staking-module/x/slashing/migrations/v3"
// 	"github.com/iqlusioninc/liquidity-staking-module/x/slashing/types"
// )

// type mockSubspace struct {
// 	ps types.Params
// }

// func newMockSubspace(ps types.Params) mockSubspace {
// 	return mockSubspace{ps: ps}
// }

// func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
// 	*ps.(*types.Params) = ms.ps
// }

// func TestMigrate(t *testing.T) {
// 	encCfg := moduletestutil.MakeTestEncodingConfig(slashing.AppModuleBasic{})
// 	cdc := encCfg.Codec

// 	storeKey := sdk.NewKVStoreKey(v3.ModuleName)
// 	tKey := sdk.NewTransientStoreKey("transient_test")
// 	ctx := testutil.DefaultContext(storeKey, tKey)
// 	store := ctx.KVStore(storeKey)

// 	legacySubspace := newMockSubspace(types.DefaultParams())
// 	require.NoError(t, v3.Migrate(ctx, store, legacySubspace, cdc))

// 	var res types.Params
// 	bz := store.Get(v3.ParamsKey)
// 	require.NoError(t, cdc.Unmarshal(bz, &res))
// 	require.Equal(t, legacySubspace.ps, res)
// }
