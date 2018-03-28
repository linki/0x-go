package cmd

import (
	"io/ioutil"

	"github.com/stretchr/testify/require"
)

func setupTestKeystoreFile(require *require.Assertions) string {
	someKeyStoreJSON := []byte(`{
	  "version": 3,
	  "id": "65c02e28-d30b-488f-aaac-877d50a9e908",
		  "address": "d3942019eb7006d519945ff2fdd431d7c7217e89",
	  "Crypto": {
	    "ciphertext": "de2b750d221e5afab297ad32c29a32e10eaaacc54f23f1d62a5b50818396e618",
	    "cipherparams": {
	      "iv": "9386486c28d50e9c1952df22463a535f"
	    },
	    "cipher": "aes-128-ctr",
	    "kdf": "scrypt",
	    "kdfparams": {
	      "dklen": 32,
	      "salt": "9b1d44da48bde629f18a807032c9a1a43c404f867f0ebbb2782bcde1b7435037",
	      "n": 8192,
	      "r": 8,
	      "p": 1
	    },
	    "mac": "638fd8f08c5a310c308e5ee08f0367756cb4fd3bb8569c9c907cb3a367b1382c"
	  }
	}`)

	tmpfile, err := ioutil.TempFile("", "eth-test-wallet")
	require.NoError(err)
	defer tmpfile.Close()

	_, err = tmpfile.Write(someKeyStoreJSON)
	require.NoError(err)

	return tmpfile.Name()
}
