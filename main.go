package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/joho/godotenv/autoload"
)

var (
	Key      = os.Getenv("S3_KEY")
	Secret   = os.Getenv("S3_SECRET")
	Endpoint = os.Getenv("S3_ENDPOINT")
	Bucket   = os.Getenv("S3_BUCKET")
	Region   = os.Getenv("S3_REGION")
)

func main() {
	Uploader()
}

func getAWSConfig() *aws.Config {
	return &aws.Config{
		Region:      aws.String("sgp1"),
		Credentials: credentials.NewStaticCredentials(Key, Secret, ""),
		Endpoint:    aws.String(Endpoint),
	}
}

func Uploader() error {
	config := getAWSConfig()
	sess, err := session.NewSession(config)
	if err != nil {
		return fmt.Errorf("failed to create a new session, %v", err)
	}

	client := s3.New(sess)

	file, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(Bucket),
		Key:    aws.String("id_card/TMaWd_FHgAChPy09tbgEP"),
	})
	if err != nil {
		return fmt.Errorf("failed to get the object, %v", err)
	}
	defer file.Body.Close()

	tmpFile, err := os.CreateTemp("", "video.*")
	if err != nil {
		return fmt.Errorf("failed to create a temp file, %v", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, file.Body)
	if err != nil {
		return fmt.Errorf("failed to copy the file, %v", err)
	}

	err = videoToM3U8(tmpFile.Name(), "TMaWd_FHgAChPy09tbgEP")
	if err != nil {
		return fmt.Errorf("failed to convert video to m3u8, %v", err)
	}

	return nil
}

func videoToM3U8(videoFilePath string, folderName string) error {
	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create output folder: %v", err)
	}

	cmd := exec.Command("ffmpeg", "-i", videoFilePath,
		"-c:v", "libx264", "-c:a", "aac", "-f", "hls",
		"-hls_time", "6",
		"-hls_list_size", "0", folderName+"/index.m3u8")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run FFmpeg command: %v", err)
	}

	err = saveM3U8File(folderName)
	if err != nil {
		return fmt.Errorf("failed to save m3u8 file to s3, %v", err)
	}

	return nil
}

func saveM3U8File(folderName string) error {
	config := getAWSConfig()
	sess, err := session.NewSession(config)
	if err != nil {
		return fmt.Errorf("failed to create a new session, %v", err)
	}

	client := s3.New(sess)

	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == folderName {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open the file, %v", err)
		}
		defer file.Close()

		Key := "HLS/" + folderName + "/" + info.Name()
		Ext := filepath.Ext(info.Name())
		CType := "application/vnd.apple.mpegurl"

		if Ext == ".ts" {
			CType = "video/mp2t"
		}

		_, err = client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(Bucket),
			Key:         aws.String(Key),
			Body:        file,
			ContentType: aws.String(CType),
			ACL:         aws.String("public-read"),
		})
		if err != nil {
			return fmt.Errorf("failed to put the object, %v", err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk through the folder, %v", err)
	}
	defer os.RemoveAll(folderName)

	return nil
}
