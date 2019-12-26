package aliyunOss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
	"os"
	"foss/library/helper"
)

/**
根据文件名上传
 */
func PutObjectFromFile(ossFileName string, localFilePath string) (string, error) {
	_, err := helper.FileExist(localFilePath)
	if err != nil {
		panic(err.Error())
	}

	client, err := oss.New("oss.aliyuncs.com", "z7w1bYwRgqniP1ul", "o0R3EhEtCxGttr6EKdvHw8DVUBV7Jd")
	if err != nil {
		handleError(err)
	}

	bucket, err := client.Bucket("gsp-fs")
	if err != nil {
		handleError(err)
	}

	err = bucket.PutObjectFromFile(ossFileName, localFilePath)
	if err != nil {
		handleError(err)
	}

	return "http://" + bucket.BucketName + ".oss.aliyuncs.com/"+ossFileName,err
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
