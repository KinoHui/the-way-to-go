package eth

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreatKs() {
	ks := keystore.NewKeyStore("./eth-wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address: ", account.Address.Hex())
}

func ImportKs() {
	dirPath := "./eth-wallets"
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	var filename string
	for _, file := range files {
		if !file.IsDir() {
			filename = file.Name()
		}
	}
	file := dirPath + "/" + filename
	fmt.Println(file)

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password) // 2025/03/08 16:16:57 account already exists
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("address: ", account.Address.Hex())

	// if err := os.Remove(file); err != nil {
	// 	log.Fatal(err)
	// }

}
