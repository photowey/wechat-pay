/*
 * @Author: photowey
 * @Date: 2022-02-01 17:54:57
 * @LastEditTime: 2022-02-01 22:00:13
 * @LastEditors: photowey
 * @Description: key.go
 * @FilePath: /wechat-pay/internal/crypto/key.go
 * Copyright (c) 2022 by photowey<photowey@gmail.com>, All Rights Reserved.
 */

package crypto

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "fmt"
)

type RSAKEY string

var (
    PRIVATE RSAKEY = "PRIVATE KEY"
    PUBLIC  RSAKEY = "PUBLIC KEY"
)

func LoadPublicKeyPem(publicKeyByte []byte) (publicKey *rsa.PublicKey, err error) {
    block, _ := pem.Decode(publicKeyByte)
    if block == nil {
        return nil, errors.New("pem.Decode public key error")
    }
    if block.Type != string(PUBLIC) {
        return nil, fmt.Errorf("the kind of PEM should be: %s", PUBLIC)
    }

    return LoadPublicKey(block.Bytes)
}

func LoadPrivateKeyPem(privateKeyByte []byte) (privateKey *rsa.PrivateKey, err error) {
    block, _ := pem.Decode(privateKeyByte)
    if block == nil {
        return nil, fmt.Errorf("pem.Decode private key err")
    }
    if block.Type != string(PRIVATE) {
        return nil, fmt.Errorf("the kind of PEM should be: %s", PRIVATE)
    }
    return LoadPrivateKey(block.Bytes)
}

func LoadPublicKey(publicKeyByte []byte) (publicKey *rsa.PublicKey, err error) {
    pub, err := x509.ParsePKIXPublicKey(publicKeyByte)
    if err != nil {
        return nil, fmt.Errorf("parse rsa public key err:%s", err.Error())
    }
    publicKey, ok := pub.(*rsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("the input parameter is not rsa public key")
    }

    return publicKey, nil
}

func LoadPrivateKey(privateKeyByte []byte) (privateKey *rsa.PrivateKey, err error) {
    key, err := x509.ParsePKCS8PrivateKey(privateKeyByte)
    if err != nil {
        return nil, fmt.Errorf("parse rsa private key err:%s", err.Error())
    }
    privateKey, ok := key.(*rsa.PrivateKey)
    if !ok {
        return nil, fmt.Errorf("the input parameter is not rsa private key")
    }
    return privateKey, nil
}
