package keymanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSKeyLoader struct {
	KeyID string
}

func (l *KMSKeyLoader) LoadKey() ([]byte, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	svc := kms.NewFromConfig(cfg)
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(l.KeyID),
	}

	result, err := svc.Decrypt(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}
