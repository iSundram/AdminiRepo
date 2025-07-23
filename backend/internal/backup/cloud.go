
package backup

import (
    "context"
    "io"
    "log"
)

type CloudProvider interface {
    Upload(ctx context.Context, key string, data io.Reader) error
    Download(ctx context.Context, key string) (io.ReadCloser, error)
    Delete(ctx context.Context, key string) error
    List(ctx context.Context, prefix string) ([]string, error)
}

type CloudBackup struct {
    Provider CloudProvider
    Bucket   string
    Prefix   string
}

func NewCloudBackup(provider CloudProvider, bucket, prefix string) *CloudBackup {
    return &CloudBackup{
        Provider: provider,
        Bucket:   bucket,
        Prefix:   prefix,
    }
}

func (cb *CloudBackup) UploadBackup(filename string, data io.Reader) error {
    ctx := context.Background()
    key := cb.Prefix + "/" + filename
    
    log.Printf("Uploading backup to cloud: %s", key)
    return cb.Provider.Upload(ctx, key, data)
}

func (cb *CloudBackup) DownloadBackup(filename string) (io.ReadCloser, error) {
    ctx := context.Background()
    key := cb.Prefix + "/" + filename
    
    log.Printf("Downloading backup from cloud: %s", key)
    return cb.Provider.Download(ctx, key)
}

func (cb *CloudBackup) ListBackups() ([]string, error) {
    ctx := context.Background()
    return cb.Provider.List(ctx, cb.Prefix)
}
