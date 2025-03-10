package tx

import (
	"fmt"
	"strconv"
)

// Flag constants for transaction-related flags
const (
	defaultGasLimit = 200000
	gasFlagAuto     = "auto"

	FlagTimeoutTimestamp = "timeout-timestamp"
	FlagChainID          = "chain-id"
	FlagNote             = "note"
	FlagSignMode         = "sign-mode"
	FlagAccountNumber    = "account-number"
	FlagSequence         = "sequence"
	FlagFrom             = "from"
	FlagDryRun           = "dry-run"
	FlagGas              = "gas"
	FlagGasAdjustment    = "gas-adjustment"
	FlagGasPrices        = "gas-prices"
	FlagFees             = "fees"
	FlagFeePayer         = "fee-payer"
	FlagFeeGranter       = "fee-granter"
	FlagUnordered        = "unordered"
	FlagOffline          = "offline"
	FlagGenerateOnly     = "generate-only"
)

// parseGasSetting parses a string gas value. The value may either be 'auto',
// which indicates a transaction should be executed in simulate mode to
// automatically find a sufficient gas value, or a string integer. It returns an
// error if a string integer is provided which cannot be parsed.
func parseGasSetting(gasStr string) (bool, uint64, error) {
	switch gasStr {
	case "":
		return false, defaultGasLimit, nil

	case gasFlagAuto:
		return true, 0, nil

	default:
		gas, err := strconv.ParseUint(gasStr, 10, 64)
		if err != nil {
			return false, 0, fmt.Errorf("gas must be either integer or %s", gasFlagAuto)
		}

		return false, gas, nil
	}
}
