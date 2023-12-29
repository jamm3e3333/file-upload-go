package gcs

import (
	"context"
	"fmt"
	"jamm3e3333/file-upload/cmd/app/config"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type StorageWriter interface {
	NewWriter(ctx context.Context, path string, fileName string) *storage.Writer
}

type FilesInfo interface {
	Files(ctx context.Context) ([]StorageFile, error)
}

type CloudStorage struct {
	cl                       *storage.Client
	bucketName               string
	basePath                 string
	signedUrlExpirationHours time.Duration
}

func NewCloudStorage(cfg *config.GcStorage) *CloudStorage {
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cfg.ServiceAccountKeyAbsFilePath)
	if err != nil {
		log.Fatalf("failed to load google service account key %v", err)
	}

	var client *storage.Client
	client, err = storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	return &CloudStorage{
		cl:                       client,
		bucketName:               cfg.BucketName,
		basePath:                 cfg.BasePath,
		signedUrlExpirationHours: time.Duration(cfg.SignedUrlExpirationHours),
	}
}

func (s *CloudStorage) NewWriter(ctx context.Context, path string, fileName string) *storage.Writer {
	filePath := fmt.Sprintf("%s/%s/%s", s.basePath, path, fileName)
	return s.cl.Bucket(s.bucketName).Object(filePath).NewWriter(ctx)
}

type StorageFile struct {
	Path      string
	SignedUrl string
}

// Files TODO: Determine the path by unique identifier
func (s *CloudStorage) Files(ctx context.Context) ([]StorageFile, error) {
	path := fmt.Sprintf("%s", s.basePath)
	query := &storage.Query{
		Projection:               storage.ProjectionNoACL,
		Prefix:                   path,
		IncludeTrailingDelimiter: false,
	}
	it := s.cl.Bucket(s.bucketName).Objects(ctx, query)

	files := []StorageFile{}
	for {
		attrs, err := it.Next()
		// when reading is done -> returns error `Done`
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		signedUrl, signedUrlErr := s.fileSignedUrl(attrs.Name)
		if signedUrlErr != nil {
			return nil, err
		}
		files = append(files, StorageFile{
			Path:      attrs.Name,
			SignedUrl: signedUrl,
		})
	}

	return files, nil
}

func (s *CloudStorage) fileSignedUrl(filePath string) (string, error) {
	signedUrl, err := s.cl.Bucket(s.bucketName).SignedURL(filePath, &storage.SignedURLOptions{
		Expires: time.Now().Add(time.Hour * s.signedUrlExpirationHours),
		Method:  "GET",
	})
	if err != nil {
		return "", err
	}

	return signedUrl, nil
}
