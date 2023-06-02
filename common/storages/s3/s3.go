package s3

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/NULL-SoftwareMeisterHighSchool/Project_FileServer/common/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
)

type s3Client struct {
	svc    *s3.S3
	bucket string
}

func Get() *s3Client {

	client := new(s3Client)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApNortheast2RegionID),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS,
			config.AWS_SECRET,
			"",
		),
	}))

	client.svc = s3.New(sess)
	client.bucket = "team-null-images"

	return client
}

func (client *s3Client) DeleteImages(urls []string) {
	for _, url := range urls {
		suffix := getSuffixFromURL(url)
		client.svc.DeleteObject(&s3.DeleteObjectInput{
			Bucket: &client.bucket,
			Key:    &suffix,
		})
	}
}

func (client *s3Client) UploadImage(name, extension string, fileHeader *multipart.FileHeader) (string, *fiber.Error) {
	fullName := strings.Join([]string{name, extension}, ".")
	file, _ := fileHeader.Open()
	defer file.Close()

	if _, err := client.svc.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: &client.bucket,
		Key:    &fullName,
	}); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return strings.Join([]string{config.IMAGE_HOST_ENDPOINT, fullName}, "/"), nil
}

func getSuffixFromURL(url string) string {
	suffix, _ := strings.CutPrefix(url, fmt.Sprintf("%s/", config.IMAGE_HOST_ENDPOINT))
	return suffix
}
