package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var RPC = "wss://eth-sepolia.g.alchemy.com/v2/5Jk48tzhV8aBBl8C1uoO8ZNyRVWi-Uf4"
var addr = "0xdE333a88A341f91Af4a3e94bE47BFa544bfD711C"
var contractAddr = "0x2FAc160c6c2642f3d748f406A6aAaE5D52e6eCaA"

func Test() error {
	err := CreateDistribution()
	if err != nil {
		log.Printf("create distribution error  %v ", err)
		return err
	}
	return nil
}

func FullTest() error {
	cl, err := ethclient.DialContext(context.Background(), RPC)
	if err != nil {
		log.Printf("client error %v ", err)
		return err
	}
	client := cl.Client()

	block, err := cl.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("block error %v", err)
		return err
	}
	address := common.HexToAddress(addr)
	fmt.Printf("block number %v", block.Number())

	balance, err := cl.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("balance err %v", err)
		return err
	}

	privateKey, publicKey, new_addr, err := GenerateKeys()
	if err != nil {
		log.Fatalf("Failed to generate keys: %v", err)
	}

	fmt.Printf("Private Key: %s\n", privateKey)
	fmt.Printf("Public Key: %s\n", publicKey)
	fmt.Printf("Ethereum Address: %s\n", new_addr)

	ether := new(big.Float).SetInt(balance)
	etherValue := new(big.Float).Quo(ether, big.NewFloat(1e18))

	fmt.Printf("Account %s has balance: %v ETH\n", address.Hex(), etherValue)

	err = ExecuteTransaction()
	if err != nil {
		log.Printf("balance err %v", err)
		return err
	}

	defer client.Close()

	return nil
}

func GenerateKeys() (string, string, string, error) {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	// Convert the private key to bytes for storage or transfer
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:] // [2:] to remove '0x' prefix
	log.Printf("Private key: %s", privateKeyHex)

	// Derive the public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	// Convert the public key to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeyBytes)[4:] // [4:] to remove '0x04' prefix that indicates an uncompressed public key
	log.Printf("Public key: %s", publicKeyHex)

	// Derive the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressHex := address.Hex()
	log.Printf("Ethereum address: %s", addressHex)

	return privateKeyHex, publicKeyHex, addressHex, nil
}

func ExecuteTransaction() error {
	client, err := ethclient.Dial(RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}
	defer client.Close()

	// Load private keys for both wallets
	privateKey2, privateKey1 := os.Getenv("DEV_WALLET1"), os.Getenv("DEV_WALLET2")
	key1, err := crypto.HexToECDSA(privateKey1)
	if err != nil {
		return fmt.Errorf("failed to parse private key 1: %w", err)
	}
	key2, err := crypto.HexToECDSA(privateKey2)
	if err != nil {
		return fmt.Errorf("failed to parse private key 2: %w", err)
	}

	fromAddress := crypto.PubkeyToAddress(key1.PublicKey)
	toAddress := crypto.PubkeyToAddress(key2.PublicKey)

	// Log balances before transaction
	LogBalance(client, fromAddress, "From address before")
	LogBalance(client, toAddress, "To address before")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch nonce: %w", err)
	}

	value := big.NewInt(1e18) // 1 ETH in Wei
	gasLimit := uint64(21000) // In units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %w", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(11155111)), key1)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}
	bind.WaitMined(context.Background(), client, signedTx)

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// Log balances after transaction
	LogBalance(client, fromAddress, "From address after")
	LogBalance(client, toAddress, "To address after")

	return nil
}

func CreateDistribution() error {
	client, err := InitializeEthClient(RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// Load the private key from an environment variable
	privateKeyHex := os.Getenv("DEV_WALLET1")
	if privateKeyHex == "" {
		log.Fatal("DEV_WALLET1 not set in environment variables")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}
	privateKey2, privateKey1 := os.Getenv("DEV_WALLET1"), os.Getenv("DEV_WALLET2")
	key1, err := crypto.HexToECDSA(privateKey1)
	if err != nil {
		return fmt.Errorf("failed to parse private key 1: %w", err)
	}
	key2, err := crypto.HexToECDSA(privateKey2)
	if err != nil {
		return fmt.Errorf("failed to parse private key 2: %w", err)
	}

	fromAddress := crypto.PubkeyToAddress(key1.PublicKey)
	toAddress := crypto.PubkeyToAddress(key2.PublicKey)

	name := "Global Tech Summit"
	ticketPrice := big.NewInt(100) // E.g., 100 wei for simplicity

	// Beneficiaries addresses
	beneficiaries := []common.Address{
		fromAddress, // Organizer's address as a beneficiary
		toAddress,   // Another beneficiary
	}

	// Splits
	splits := []*big.Int{
		big.NewInt(70), // 70% to DEV_WALLET1
		big.NewInt(30), // 30% to DEV_WALLET2
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia testnet chain ID
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddr)

	contract, err := NewEventManagement(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to load contract: %v", err)
	}

	// Subscribe to the EventCreated event
	eventCh := make(chan *EventManagementEventCreated)
	sub, err := contract.WatchEventCreated(nil, eventCh, nil)
	if err != nil {
		return fmt.Errorf("failed to watch EventCreated events: %w", err)
	}
	defer sub.Unsubscribe()

	// Create the event using the contract
	tx, err := contract.CreateEvent(auth, name, ticketPrice, beneficiaries, splits)
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}
	bind.WaitMined(context.Background(), client, tx)

	fmt.Printf("Tx sent: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Println("Waiting for transaction to be mined...")
	_, err = bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("error waiting for transaction to be mined: %w", err)
	}

	// Listen for the event
	select {
	case event := <-eventCh:
		fmt.Printf("Event created: id=%d, name=%s, ticketPrice=%s, beneficiaries=%v, splits=%v\n",
			event.Id, event.Name, event.TicketPrice, event.Beneficiaries, event.Splits)
	case err := <-sub.Err():
		return fmt.Errorf("error while watching for events: %w", err)
	case <-ctx.Done():
		return fmt.Errorf("timeout waiting for event")
	}
	// Get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// Log the transaction receipt data
	log.Printf("Receipt Status: %v\n", receipt.Status)
	log.Printf("Gas Used: %v\n", receipt.GasUsed)
	log.Printf("Cumulative Gas Used: %v\n", receipt.CumulativeGasUsed)
	log.Printf("Logs: %v\n", receipt.Logs)

	// Print each log in a readable format
	for i, vLog := range receipt.Logs {
		log.Printf("Log %d: %+v\n", i, vLog)
		log.Printf("Log %d - Address: %s\n", i, vLog.Address.Hex())
		log.Printf("Log %d - Topics: %v\n", i, vLog.Topics)
		log.Printf("Log %d - Data: %s\n", i, vLog.Data)
	}

	return nil
}

// InitializeEthClient creates a new Ethereum client connected to the specified endpoint.
func InitializeEthClient(rpcURL string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetNonce retrieves the nonce for the given address.
func GetNonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// Helper to convert Wei to ETH
func FromWei(value *big.Int) *big.Float {
	fvalue := new(big.Float)
	fvalue.SetString(value.String())
	return new(big.Float).Quo(fvalue, big.NewFloat(1e18))
}

func LogBalance(client *ethclient.Client, address common.Address, msg string) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("Failed to retrieve balance for %s: %v", msg, err)
		return
	}

	fmt.Printf("%v: %v has balance: %v ETH\n", msg, address.Hex(), new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18)))
}
