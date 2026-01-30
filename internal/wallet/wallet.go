package wallet

import (
	"fmt"
	"os"
    "crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
    PrivateKey *ecdsa.PrivateKey
    Address    string
}

func LoadWallet() (*Wallet, error) {
    pkHex := os.Getenv("PRIVATE_KEY")
    if pkHex == "" {
        return nil, fmt.Errorf("PRIVATE_KEY environment variable not set")
    }

    if len(pkHex) > 2 && pkHex[:2] == "0x" {
        pkHex = pkHex[2:]
    }

    privateKey, err := crypto.HexToECDSA(pkHex)
    if err != nil {
        return nil, fmt.Errorf("invalid private key: %w", err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }

    address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

    return &Wallet{
        PrivateKey: privateKey,
        Address:    address,
    }, nil
}