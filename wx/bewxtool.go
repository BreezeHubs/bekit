package wx

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"unsafe"

	"errors"

	"golang.org/x/sync/errgroup"
)

var (
	IV_LENTH_ERROR                     = errors.New("The length of iv is not equal to 24")
	SESSION_KEY_LENTH_ERROR            = errors.New("The length of sessionKey is not equal to 24")
	IV_DECRYPT_BASE64_ERROR            = errors.New("iv Failed to decrypt base64")
	SESSION_KEY_DECRYPT_BASE64_ERROR   = errors.New("sessionKey Failed to decrypt base64")
	ENCRYPTEDDATA_DECRYPT_BASE64_ERROR = errors.New("encryptedData Failed to decrypt base64")
	AESDECRYPT_ERROR                   = errors.New("AesDecrypt fail")
)

type beWXTools struct {
	appId      string
	sessionKey string
}

func NewBeWXTools(appId, sessionKey string) (*beWXTools, error) {
	if len(sessionKey) != 24 {
		return nil, SESSION_KEY_LENTH_ERROR
	}

	return &beWXTools{
		appId:      appId,
		sessionKey: sessionKey,
	}, nil
}

// 小程序 根据前端提供的encryptedData、iv解密用户信息 返回string
func (b *beWXTools) DecryptWXUserInfoString(encryptedData, iv string) (string, error) {
	result, err := decryptWXUserInfo(b.sessionKey, encryptedData, iv)
	if err != nil {
		return "", err
	}

	return *(*string)(unsafe.Pointer(&result)), nil
}

type WXUserInfo struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	Watermark struct {
		Timestamp int    `json:"timestamp"`
		Appid     string `json:"appid"`
	} `json:"watermark"`
}

// 小程序 根据前端提供的encryptedData、iv解密用户信息 返回结构体
func (b *beWXTools) DecryptWXUserInfo(encryptedData, iv string) (*WXUserInfo, error) {
	result, err := decryptWXUserInfo(b.sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	var data WXUserInfo
	if err := json.Unmarshal(result, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func decryptWXUserInfo(sessionKey, encryptedData, iv string) ([]byte, error) {
	if len(iv) != 24 {
		return nil, IV_LENTH_ERROR
	}

	var aesKeyByte, aesIVByte, aesCipherByte []byte
	eg := errgroup.Group{}
	eg.Go(func() (err error) {
		aesKeyByte, err = base64.StdEncoding.DecodeString(sessionKey)
		if err != nil {
			return SESSION_KEY_DECRYPT_BASE64_ERROR
		}
		return
	})

	eg.Go(func() (err error) {
		aesIVByte, err = base64.StdEncoding.DecodeString(iv)
		if err != nil {
			return IV_DECRYPT_BASE64_ERROR
		}
		return
	})

	eg.Go(func() (err error) {
		aesCipherByte, err = base64.StdEncoding.DecodeString(encryptedData)
		if err != nil {
			return ENCRYPTEDDATA_DECRYPT_BASE64_ERROR
		}
		return
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	result, err := aesDecrypt(aesCipherByte, aesKeyByte, aesIVByte)
	if err != nil {
		return nil, AESDECRYPT_ERROR
	}
	return result, nil
}

// openssl_decrypt AES-128-CBC
func aesDecrypt(src, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	decData := make([]byte, len(src))
	blockMode.CryptBlocks(decData, src)

	// pkcs7UnPadding
	length := len(decData)
	unpadding := int(decData[length-1])
	decData = decData[:(length - unpadding)]

	return decData, nil
}

// auth.code2Session 小程序授权登录
