package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/joho/godotenv"
)

// ModuleBasics is a mock module basic manager for testing
var ModuleBasics = module.NewBasicManager()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Key added to your test keyring
	user := os.Getenv("WALLET_KEY")
	// Message to be signed
	msg := os.Getenv("MSG")
	appName := os.Getenv("APP_NAME")
	appRootDir := os.Getenv("APP_ROOT_DIR")
	backend := os.Getenv("BACKEND")

	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	cdc := codec.NewProtoCodec(interfaceRegistry)

	buf := bufio.NewReader(os.Stdin)

	kb, err := cKeys.New(appName, backend, appRootDir, buf, cdc)
	if err != nil {
		log.Fatalf("%v", err)
	}

	bites := []byte(msg)

	signature, pk, err := kb.Sign(user, bites)
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Verify signature
	if !pk.VerifySignature(bites, signature) {
		log.Fatal("bad signature")
	}

	// Print signature
	fmt.Println("Signature: ", hex.EncodeToString(signature[:]))
	// Print public key
	fmt.Println("Public Key: ", pk)
}
