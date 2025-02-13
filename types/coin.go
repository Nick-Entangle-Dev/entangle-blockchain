// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/Entangle-Protocol/entangle-blockchain/blob/main/LICENSE
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoNGL defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	AttoNGL string = "aNGL"

	// BaseDenomUnit defines the base denomination unit for NGLs.
	// 1 NGL = 1x10^{BaseDenomUnit} aNGL
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewNGLCoin is a utility function that returns an "aNGL" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewNGLCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoNGL, amount)
}

// NewNGLDecCoin is a utility function that returns an "aNGL" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewNGLDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoNGL, amount)
}

// NewNGLCoinInt64 is a utility function that returns an "aNGL" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewNGLCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoNGL, amount)
}
