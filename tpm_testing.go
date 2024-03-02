package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/google/go-tpm/tpmutil"
	"io"
	"os"

	"github.com/google/go-tpm/legacy/tpm2"
)

const (
	emptyPassword   = ""
	defaultPassword = ""
)

// # tpm2_evictcontrol -C o -c0x81008000
// persistent-handle: 0x81008000
// action: evicted

var handleNames = map[string][]tpm2.HandleType{
	"all":       {tpm2.HandleTypeLoadedSession, tpm2.HandleTypeSavedSession, tpm2.HandleTypeTransient},
	"loaded":    {tpm2.HandleTypeLoadedSession},
	"saved":     {tpm2.HandleTypeSavedSession},
	"transient": {tpm2.HandleTypeTransient},
}

var (
	tpmPath          = flag.String("tpm-path", "/dev/tpmrm0", "Path to the TPM device (character device or a Unix socket).")
	persistentHandle = flag.Uint("persistentHandle", 0x81008000, "Handle value")
	evict            = flag.Bool("evict", false, "Evict prior handle")
	flush            = flag.String("flush", "all", "Flush contexts, must be oneof transient|saved|loaded|all")
	mode             = flag.String("mode", "encrypt", "create or encrypt or decrypt")
	encryptedS       = flag.String("enc", "", "")

	defaultKeyParams = tpm2.Public{
		Type:       tpm2.AlgRSA,
		NameAlg:    tpm2.AlgSHA256,
		Attributes: tpm2.FlagDecrypt | tpm2.FlagUserWithAuth | tpm2.FlagFixedParent | tpm2.FlagFixedTPM | tpm2.FlagSensitiveDataOrigin,
		RSAParameters: &tpm2.RSAParams{
			Sign: &tpm2.SigScheme{
				Alg:  tpm2.AlgNull,
				Hash: tpm2.AlgNull,
			},
			KeyBits: 2048,
		},
	}
	//
	//rsaKeyParams = tpm2.Public{
	//	Type:    tpm2.AlgRSA,
	//	NameAlg: tpm2.AlgSHA256,
	//	Attributes: tpm2.FlagFixedTPM | tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin |
	//		tpm2.FlagUserWithAuth | tpm2.FlagDecrypt,
	//	AuthPolicy: []byte{},
	//	RSAParameters: &tpm2.RSAParams{
	//		Sign: &tpm2.SigScheme{
	//			Alg:  tpm2.AlgNull,
	//			Hash: tpm2.AlgNull,
	//		},
	//		KeyBits: 2048,
	//	},
	//}
)

func createAndPersistKeysInTPM() {
	flag.Parse()
	if *mode == "create" {
		rwc, err := tpm2.OpenTPM(*tpmPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't open TPM %s: %v", *tpmPath, err)
			return
		}

		pcrList := []int{0}
		pcrSelection := tpm2.PCRSelection{Hash: tpm2.AlgSHA256, PCRs: pcrList}

		pkh, _, err := tpm2.CreatePrimary(rwc, tpm2.HandleOwner, pcrSelection, emptyPassword, emptyPassword, defaultKeyParams)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating Primary %v\n", err)
			return
		}
		defer tpm2.FlushContext(rwc, pkh)

		pHandle := tpmutil.Handle(uint32(*persistentHandle))
		defer tpm2.FlushContext(rwc, pHandle)
		err = tpm2.EvictControl(rwc, "", tpm2.HandleOwner, pkh, pHandle)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Persisting the key %v\n", err)
			return
		}
		fmt.Printf("======= Key persisted ========\n")
	}
}

func encryptUsingTPM(dataToSeal []byte) string {
	rwc, err := tpm2.OpenTPM(*tpmPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open TPM %s: %v", *tpmPath, err)
		return ""
	}
	pHandle := tpmutil.Handle(uint32(*persistentHandle))
	fmt.Printf("======= Evicting Handles ========\n")

	//client.Handles(rwc, tpm2.HandleTypePersistent)
	if !handleExists(rwc, pHandle) {
		createAndPersistKeysInTPM()
	}
	encrypted, err := tpm2.RSAEncrypt(rwc, pHandle, dataToSeal, &tpm2.AsymScheme{Alg: tpm2.AlgOAEP, Hash: tpm2.AlgSHA256}, "label")
	if err != nil {
		glog.Fatalf("Error Encrypting: %v", err)
	}

	encryptedString := hex.EncodeToString(encrypted)
	fmt.Printf("Encrypted data " + encryptedString)
	glog.V(2).Infof("Encrypted Data %v", encryptedString)
	return encryptedString
}

func handleExists(rwc io.ReadWriter, handle tpmutil.Handle) bool {
	// Use tpm2.ReadPublic to check if the specified handle exists
	_, _, _, err := tpm2.ReadPublic(rwc, handle)
	return err == nil
}

func decryptUsingTPM(encrypted []byte) []byte {

	//encrypted, _ := base64.RawStdEncoding.DecodeString("a/UxW1gRb6w6cPIQ5ADSv32MJa6S03FjixOviE69VP7REJ5ahN/+Lo0uzESt5E7kKlyDj5+dNmhwYVw2sS0CH2Qi9yrmtzFr/MgHODa7UbGuPMKcmvanZQ9T92X1vNXEsy+JfKA80pHYD6byNsU0gqlJiMD7WTCiM8m6uNi90heplqwfZZKD7SJ/VB6aEtvu4W8W2IPWS3UYqvabP/sYfQDPHRjI25HuFGthOJ+cIWrXfaFzm+rXvvk4Yd+yAVFK8W7CzxttlVe4d2ISZd8LctisTNmm3T0azm28+I1eerDlMiLHRvnJnWcTT82IfA66mgre2Z1ga4pop6yaRIH62g")
	rwc, err := tpm2.OpenTPM(*tpmPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open TPM %s: %v", *tpmPath, err)
		return nil
	}

	pHandle := tpmutil.Handle(uint32(*persistentHandle))
	if !handleExists(rwc, pHandle) {
		createAndPersistKeysInTPM()
	}
	decrypted, err := tpm2.RSADecrypt(rwc, pHandle, "", encrypted, &tpm2.AsymScheme{Alg: tpm2.AlgOAEP, Hash: tpm2.AlgSHA256}, "label")
	if err != nil {
		glog.Fatalf("Error Decrypting: %v", err)
	}
	fmt.Printf("Decrypted : " + string(decrypted))
	return decrypted
}
