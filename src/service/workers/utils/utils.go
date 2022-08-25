package utils

import (
	"crypto/ecdsa"
	"math"
	"math/big"
	"strconv"

	"github.com/latoken/bridge-backend-service/src/models"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPrivateKey(config *models.WorkerConfig) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	return privKey, nil
}

// QuoBigInt ...
func QuoBigInt(a *big.Int, b *big.Int) *big.Float {
	fl := new(big.Float).SetInt(a)
	fl.Quo(fl, new(big.Float).SetInt(b))
	return fl
}

// GetBigIntForDecimal ...
func GetBigIntForDecimal(decimal int) *big.Int {
	floatDecimal := big.NewFloat(math.Pow10(decimal))
	bigIntDecimal := new(big.Int)
	floatDecimal.Int(bigIntDecimal)

	return bigIntDecimal
}

// CalcActualOutAmount ...
func CalcActualOutAmount(amount *big.Int, ratio *big.Float, fixedFee *big.Int) *big.Int {
	res := new(big.Float).SetInt(amount)
	res.Mul(res, ratio)

	amountInt := new(big.Int)
	res.Int(amountInt)

	amountInt.Sub(amountInt, fixedFee)
	return amountInt
}

func BytesToBytes32(b []byte) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], b)
	return byteArr
}

func StringToBytes32(b string) [32]byte {
	var byteArr [32]byte
	copy(byteArr[:], common.RightPadBytes(common.Hex2Bytes(b), 32))
	return byteArr
}

func StringToBytes8(b string) [8]byte {
	var byteArr [8]byte
	copy(byteArr[:], common.RightPadBytes(common.Hex2Bytes(b), 8))
	return byteArr
}

func BytesToBytes8(b []byte) [8]byte {
	var byteArr [8]byte
	copy(byteArr[:], b)
	return byteArr
}

func StringToBytes(b string) []byte {
	var byteArr []byte
	copy(byteArr[:], common.Hex2Bytes(b))
	return byteArr
}

func CalcutateSwapID(originChainID, destChainID, nonce string) string {
	return originChainID + destChainID + nonce
}

func ConvertDecimals(amount string, inDecimals, outDecimals int64) int64 {
	value, _ := strconv.ParseInt(amount, 10, 32)
	ret := value * outDecimals / inDecimals
	return ret
}

func CalculateLiquidityIndex(balance, scaledBalance *big.Int) *big.Int {
	if scaledBalance.Cmp(big.NewInt(0)) <= 0 {
		return big.NewInt(0)
	}
	exp := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(27), nil)
	divisor := balance.Div(balance, scaledBalance)
	liquidityIndex := big.NewInt(1).Mul(exp, divisor)
	return liquidityIndex
}

//converts decimal outputAmount -> inputAmount (used to convert amount from proposal event)
func ConvertDecimalsForInput(originDecimals, destDecimals uint8, amount string) *big.Int {
	origin := new(big.Int).SetInt64(int64(math.Pow10(int(originDecimals))))
	dest := new(big.Int).SetInt64(int64(math.Pow10(int(destDecimals))))
	amountInFloat, _ := new(big.Int).SetString(amount, 10)
	// conversion := new(big.Int).Quo(origin, dest)
	return new(big.Int).Quo(new(big.Int).Mul(amountInFloat, origin), dest)
}

func CheckGasSwap(resourceID string) bool {
	swapIdentifier := hex.EncodeToString([]byte("swap"))
	if resourceID[:8] == swapIdentifier {
		return true
	} else {
		return false
	}
}
