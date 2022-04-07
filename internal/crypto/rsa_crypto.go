/*
 * @Author: photowey
 * @Date: 2022-02-01 17:44:08
 * @LastEditTime: 2022-02-01 18:02:34
 * @LastEditors: photowey
 * @Description: rsa_crypto.go
 * @FilePath: /wechat-pay/internal/crypto/rsa_crypto.go
 * Copyright (c) 2022 by photowey<photowey@gmail.com>, All Rights Reserved.
 */

package crypto

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha1"
    "encoding/base64"
    "fmt"
)

// EncryptOAEP RSA 公钥加密
func EncryptOAEP(data string, publicKey *rsa.PublicKey) (encryptString string, err error) {
    if publicKey == nil {
        return "", fmt.Errorf("the *rsa.PublicKey can't be nil")
    }
    encryptByte, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, []byte(data), nil)
    if err != nil {
        return "", fmt.Errorf("rsa:encrypt data err:%s", err.Error())
    }
    encryptString = base64.StdEncoding.EncodeToString(encryptByte)

    return encryptString, nil
}

// DecryptOAEP RSA 私钥解密
func DecryptOAEP(encryptString string, privateKey *rsa.PrivateKey) (data string, err error) {
    if privateKey == nil {
        return "", fmt.Errorf("the *rsa.PrivateKey can't be nil")
    }
    decodedString, err := base64.StdEncoding.DecodeString(encryptString)
    if err != nil {
        return "", fmt.Errorf("base64 decode decrypt text failed, error=%s", err.Error())
    }
    dataBytes, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privateKey, decodedString, nil)
    if err != nil {
        return "", fmt.Errorf("decrypt err:%s", err)
    }

    return string(dataBytes), nil
}
