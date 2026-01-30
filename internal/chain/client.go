package chain

import (
    "context"
    "fmt"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/nekowawolf/MSC/internal/config"
    "github.com/nekowawolf/MSC/internal/wallet"
)

type ChainClient struct {
    Client  *ethclient.Client
    Config  config.ChainConfig
    Wallet  *wallet.Wallet
    chainId *big.Int
}

func NewChainClient(cfg config.ChainConfig, w *wallet.Wallet) (*ChainClient, error) {
    client, err := ethclient.Dial(cfg.RpcURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to RPC: %w", err)
    }

    return &ChainClient{
        Client:  client,
        Config:  cfg,
        Wallet:  w,
        chainId: big.NewInt(cfg.ChainID),
    }, nil
}

func (c *ChainClient) GetBalance(ctx context.Context) (*big.Float, error) {
    balance, err := c.Client.BalanceAt(ctx, common.HexToAddress(c.Wallet.Address), nil)
    if err != nil {
        return nil, err
    }
    
    fBalance := new(big.Float).SetInt(balance)
    ethValue := new(big.Float).Quo(fBalance, big.NewFloat(1e18))
    return ethValue, nil
}

func (c *ChainClient) SendTransaction(ctx context.Context, to string, amount float64) (string, error) {
    nonce, err := c.Client.PendingNonceAt(ctx, common.HexToAddress(c.Wallet.Address))
    if err != nil {
        return "", fmt.Errorf("failed to get nonce: %w", err)
    }

    value := new(big.Int)
    
    fAmount := new(big.Float).SetFloat64(amount)
    fWei := new(big.Float).Mul(fAmount, big.NewFloat(1e18))
    fWei.Int(value)

    gasLimit := uint64(21000)
    gasPrice, err := c.Client.SuggestGasPrice(ctx)
    if err != nil {
        return "", fmt.Errorf("failed to suggest gas price: %w", err)
    }

    toAddress := common.HexToAddress(to)
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.chainId), c.Wallet.PrivateKey)
    if err != nil {
        return "", fmt.Errorf("failed to sign tx: %w", err)
    }

    err = c.Client.SendTransaction(ctx, signedTx)
    if err != nil {
        return "", fmt.Errorf("failed to send tx: %w", err)
    }

    return signedTx.Hash().Hex(), nil
}

func (c *ChainClient) Close() {
    c.Client.Close()
}