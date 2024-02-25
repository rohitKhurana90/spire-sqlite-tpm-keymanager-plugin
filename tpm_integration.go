package main

//import (
//	"bytes"
//	"flag"
//	"github.com/google/go-tpm/tpmutil"
//	"io"
//	"log"
//
//	"encoding/base64"
//	"encoding/hex"
//
//	//"github.com/google/go-tpm/tpm2"
//	"github.com/google/go-tpm/legacy/tpm2"
//)
//
//const (
//	emptyPassword = "1234"
//)
//
//var (
//	tpmPath            = flag.String("tpm-path", "/dev/tpm0", "Path to the TPM device (character device or a Unix socket).")
//	tpmHandle          io.ReadWriteCloser
//	symmetricKeyHandle tpmutil.Handle
//
//	defaultKeyParams = tpm2.Public{
//		Type:    tpm2.AlgRSA,
//		NameAlg: tpm2.AlgSHA256,
//		Attributes: tpm2.FlagDecrypt | tpm2.FlagRestricted | tpm2.FlagFixedTPM |
//			tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth,
//		AuthPolicy: []byte{},
//		RSAParameters: &tpm2.RSAParams{
//			Symmetric: &tpm2.SymScheme{
//				Alg:     tpm2.AlgAES,
//				KeyBits: 128,
//				Mode:    tpm2.AlgCFB,
//			},
//			KeyBits: 2048,
//		},
//	}
//
//	symKeyParams = tpm2.Public{
//		Type:    tpm2.AlgSymCipher,
//		NameAlg: tpm2.AlgSHA256,
//		Attributes: tpm2.FlagDecrypt | tpm2.FlagSign | tpm2.FlagUserWithAuth |
//			tpm2.FlagFixedParent | tpm2.FlagFixedTPM | tpm2.FlagSensitiveDataOrigin,
//		SymCipherParameters: &tpm2.SymCipherParams{
//			Symmetric: &tpm2.SymScheme{
//				Alg:     tpm2.AlgAES,
//				KeyBits: 128,
//				Mode:    tpm2.AlgCFB,
//			},
//		},
//	}
//)
//
//func encryptUsingTPM(data []byte) string {
//	log.Println("======= Encrypting using TPM  ========")
//	rwc, err := tpm2.OpenTPM(*tpmPath)
//	if err != nil {
//		log.Fatalf("can't open TPM %q: %v", tpmPath, err)
//	}
//	defer func() {
//		if err := rwc.Close(); err != nil {
//			log.Fatalf("%v\ncan't close TPM %q: %v", tpmPath, err)
//		}
//	}()
//	iv := bytes.Repeat([]byte("1234"), 16)
//	encrypted, err := tpm2.EncryptSymmetric(rwc, emptyPassword, symmetricKeyHandle, iv, data)
//	if err != nil {
//		log.Fatalf("EncryptSymmetric failed: %s", err)
//	}
//	log.Printf("Encrypted %s", base64.StdEncoding.EncodeToString(encrypted))
//	return hex.EncodeToString(encrypted)
//}
//
//func decryptUsingTPM(encryptedData []byte) []byte {
//
//	log.Println("======= Decrypting using TPM  ========")
//	rwc, err := tpm2.OpenTPM(*tpmPath)
//	if err != nil {
//		log.Fatalf("can't open TPM %q: %v", tpmPath, err)
//	}
//	defer func() {
//		if err := rwc.Close(); err != nil {
//			log.Fatalf("%v\ncan't close TPM %q: %v", tpmPath, err)
//		}
//	}()
//	iv := bytes.Repeat([]byte("1234"), 16)
//	decrypted, err := tpm2.DecryptSymmetric(rwc, emptyPassword, symmetricKeyHandle, iv, encryptedData)
//	if err != nil {
//		log.Fatalf("DecryptSymmetric failed: %s", err)
//	}
//
//	log.Printf("Decrypted %s", string(decrypted))
//	return decrypted
//}
//
//func openTpmAndInitialize() {
//	flag.Parse()
//	log.Println("======= Init  ========")
//
//	rwc, err := tpm2.OpenTPM(*tpmPath)
//	if err != nil {
//		log.Fatalf("can't open TPM %q: %v", tpmPath, err)
//	}
//	defer func() {
//		if err := rwc.Close(); err != nil {
//			log.Fatalf("%v\ncan't close TPM %q: %v", tpmPath, err)
//		}
//	}()
//	log.Println("======= open tpm finished  ========")
//
//	pcrList := []int{}
//	pcrSelection := tpm2.PCRSelection{Hash: tpm2.AlgSHA256, PCRs: pcrList}
//
//	log.Printf("======= createPrimary ========")
//	pkh, _, err := tpm2.CreatePrimary(rwc, tpm2.HandleEndorsement, pcrSelection, emptyPassword, emptyPassword, defaultKeyParams)
//	if err != nil {
//		log.Fatalf("Error creating EK: %v", err)
//	}
//	defer tpm2.FlushContext(rwc, pkh)
//
//	log.Printf("======= CreateKey ========")
//
//	symPriv, symPub, _, _, _, err := tpm2.CreateKey(rwc, pkh, pcrSelection, emptyPassword, emptyPassword, symKeyParams)
//	if err != nil {
//		log.Fatalf("Create SymKey failed: %s", err)
//	}
//	log.Printf("symPub: %v,", hex.EncodeToString(symPub))
//	log.Printf("symPriv: %v,", hex.EncodeToString(symPriv))
//
//	tPub, err := tpm2.DecodePublic(symPub)
//	if err != nil {
//		log.Fatalf("Error DecodePublic AK %v", tPub)
//	}
//
//	symkeyHandle, keyName, err := tpm2.Load(rwc, pkh, "", symPub, symPriv)
//	//defer tpm2.FlushContext(rwc, symkeyHandle)
//	if err != nil {
//		log.Fatalf("Load symkh failed: %s", err)
//	}
//	log.Printf("SYM keyName: %v,", hex.EncodeToString(keyName))
//	symmetricKeyHandle = symkeyHandle
//	//iv := make([]byte, 16)
//}
