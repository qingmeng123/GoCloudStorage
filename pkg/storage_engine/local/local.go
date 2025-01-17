package local

import (
	"context"
	"fmt"
	"github.com/GoCloudstorage/GoCloudstorage/pkg/db/redis"
	"github.com/GoCloudstorage/GoCloudstorage/pkg/storage_engine"
	redis2 "github.com/redis/go-redis/v9"
	"path"
	"time"
)

var (
	FileAlreadyExist = fmt.Errorf("file already exist")
)

type StorageEngine struct {
	rootPath string
	uploader chunkUploader
}

func (s *StorageEngine) UploadChunk(request storage_engine.UploadChunkRequest) error {
	dirPath := s.getFileDir(request.FileMD5)
	return s.uploader.saveChunk(dirPath, request.PartNum, request.Data)
}

func (s *StorageEngine) getFileDir(fileMD5 string) string {
	return path.Join(s.rootPath, fileMD5)
}

// GenerateObjectURL 获取文件存储位置
func (s *StorageEngine) GenerateObjectURL(key string, expire time.Duration) (string, error) {
	filePath := path.Join(s.getFileDir(key), "data")
	cmd := redis.Client.Get(context.Background(), key)
	if cmd.Err() == redis2.Nil {
		redis.Client.SetEx(context.Background(), key, filePath, expire)
	} else if cmd.Err() != nil {
		return "", fmt.Errorf("redis failed get key, err: %v", cmd.Err())
	}
	return filePath, nil
}

func (s *StorageEngine) Init(config storage_engine.InitConfig) {
	s.rootPath = path.Join(config.Endpoint, config.BucketName)
}

func (s *StorageEngine) MergeChunk(fileMD5 string, partSize int, dataSize int) error {
	dirPath := s.getFileDir(fileMD5)
	return s.uploader.mergeChunk(dirPath, partSize, dataSize)
}

func (s *StorageEngine) GetObjectURL(key string) (string, error) {
	cmd := redis.Client.Get(context.Background(), key)
	if cmd.Err() != nil {
		return "", cmd.Err()
	}
	return cmd.Result()
}
