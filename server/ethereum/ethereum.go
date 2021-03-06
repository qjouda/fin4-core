package ethereum

import (
	"errors"
	"math/big"

	"github.com/FuturICT2/fin4-core/server/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lytics/logrus"
)

// Ethereum ethereum struct to implement crypto interface
type Ethereum struct {
	rpc  *ethclient.Client
	sim  *backends.SimulatedBackend
	auth *bind.TransactOpts
}

// MustNewEthereum create new Ethereum interface, panic if no connection
func MustNewEthereum() *Ethereum {
	rawKey := util.MustGetenv("ETH_KEY_RAW")
	conn, err := ethclient.Dial("https://rinkeby.infura.io/")
	if err != nil {
		logrus.Fatal("Failed to connect to the Ethereum client: %v", err)
	}
	rawKeyECDSA, err := crypto.HexToECDSA(rawKey)
	if err != nil {
		logrus.Fatal("Something wrong with server private key.", err)
	}
	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(rawKeyECDSA)
	if err != nil {
		logrus.Fatal("Failed to create transactor: %v", err)
	}
	// Setup blockchain simmulator
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(10000000000)},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 10393939)
	return &Ethereum{
		rpc:  conn,
		sim:  sim,
		auth: auth,
	}
}

// GetBlockNumber returns best blocknumber in the blockchain
func (b *Ethereum) GetBlockNumber() (int, error) {
	return 0, nil
}

// DeployNewToken deployes new token to Ethereum from server account
func (b *Ethereum) DeployNewToken(
	initialSupply *big.Int,
	tokenName string,
	decimals uint8,
	tokenSymbol string,
) (common.Address, *types.Transaction, error) {
	address, tx, _, err := DeployToken(
		b.auth, b.rpc, initialSupply, tokenName, decimals, tokenSymbol)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("ethereum:DeployNewToken:e1")
		return address, nil, errors.New("Error deploying token contract to Ethereum")
	}
	return address, tx, nil
}
