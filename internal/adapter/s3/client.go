// Package s3
package s3

import (
	"context"
	"fmt"
	"main/pkg"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(ctx context.Context, cfg pkg.S3Config) (*minio.Client, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.SecretKey, cfg.AccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}
	return client, nil
}
