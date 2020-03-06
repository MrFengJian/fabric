/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package csp

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/bccsp/signer"
	"github.com/pkg/errors"
)

// LoadPrivateKey loads a private key from file in keystorePath
func LoadPrivateKey(keystorePath string, algo string) (bccsp.Key, crypto.Signer, error) {
	var err error
	var priv bccsp.Key
	var s crypto.Signer

	opts := &factory.FactoryOpts{
		ProviderName: "GM",
		SwOpts: &factory.SwOpts{
			HashFamily: bccsp.SM3,
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorePath,
			},
		},
	}

	csp, err := factory.GetBCCSPFromOpts(opts)
	if err != nil {
		return nil, nil, err
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "_sk") {
			rawKey, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			block, _ := pem.Decode(rawKey)
			if block == nil {
				return errors.Errorf("%s: wrong PEM encoding", path)
			}
			//priv, err = csp.KeyImport(block.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: true})
			var opt bccsp.KeyImportOpts
			if bccsp.RSA == strings.ToUpper(algo) {
				opt = &bccsp.RSA2048PrivateKeyImportOpts{Temporary: true}
			} else {
				opt = &bccsp.SM2PrivateKeyImportOpts{Temporary: true}
			}
			priv, err = csp.KeyImport(block.Bytes, opt)
			if err != nil {
				return err
			}

			s, err = signer.New(csp, priv)
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	}

	err = filepath.Walk(keystorePath, walkFunc)
	if err != nil {
		return nil, nil, err
	}

	return priv, s, err
}

// GeneratePrivateKey creates a private key and stores it in keystorePath
func GeneratePrivateKey(keystorePath string, algo string) (bccsp.Key,
	crypto.Signer, error) {

	var err error
	var priv bccsp.Key
	var s crypto.Signer

	opts := &factory.FactoryOpts{
		ProviderName: "GM",
		SwOpts: &factory.SwOpts{
			HashFamily: bccsp.SM3,
			SecLevel:   256,

			FileKeystore: &factory.FileKeystoreOpts{
				KeyStorePath: keystorePath,
			},
		},
	}
	csp, err := factory.GetBCCSPFromOpts(opts)
	if err == nil {
		// 根据算法参数设置私钥生成参数，目前只使用RSA2048
		var opt bccsp.KeyGenOpts
		opt = &bccsp.SM2KeyGenOpts{Temporary: false}
		if strings.ToUpper(algo) == bccsp.RSA {
			opt = &bccsp.RSAKeyGenOpts{Temporary: false}
		}
		// generate a key
		priv, err = csp.KeyGen(opt)
		if err == nil {
			// create a crypto.Signer
			s, err = signer.New(csp, priv)
		}
	}
	return priv, s, err
}

func GetPublicKey(priv bccsp.Key) (interface{}, error) {

	// get the public key
	pubKey, err := priv.PublicKey()
	if err != nil {
		return nil, err
	}
	// marshal to bytes
	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		return nil, err
	}
	// unmarshal using pkix
	lowLevelKey, err := sm2.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		// 可能是RSA算法
		lowLevelKey, err = x509.ParsePKIXPublicKey(pubKeyBytes)
		if err != nil {
			return nil, err
		}
	}
	return lowLevelKey, nil
}

func GetSM2PublicKey(priv bccsp.Key) (*sm2.PublicKey, error) {

	// get the public key
	pubKey, err := priv.PublicKey()
	if err != nil {
		return nil, err
	}

	// marshal to bytes
	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		return nil, err
	}
	// unmarshal using pkix
	sm2PubKey, err := sm2.ParseSm2PublicKey(pubKeyBytes)
	//ecPubKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, err
	}
	return sm2PubKey, nil
}
