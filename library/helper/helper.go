package helper

import (
	"os"
	"strings"
	"path/filepath"
	"log"
	"github.com/satori/go.uuid"
	"archive/zip"
	"io"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha256"
	"reflect"
	"unicode"
	"fmt"
)

func GetCurrentPath() (string,error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	path := strings.Replace(dir, "\\", "/", -1) //将\替换成/

	return path,err
}

func Guid() string {
	uuidString,err := uuid.NewV4()
	if err != nil {

	}

	return uuidString.String()
}

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}

	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

/**
判断文件或文件夹是否存在
 */
func FileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }

	return true, err
}

func CryptMd5Encode(pwdText string) string {
	const BlockSize = 64

	h := md5.New()
	h.Write([]byte(pwdText))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func CryptSha256Encode(content string) string {
	const BlockSize = 64

	h := sha256.New()
	h.Write([]byte(content))

	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

/**
字符串写入指定文件
 */
func WriteFile(fileName string,content string) error {
	if _, err := os.Stat(fileName); err != nil {
		file, err := os.Create(fileName)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}


		err2 := putContentIntoFile(file,content)

		return err2
	}else {
		file,err := os.Open(fileName)
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}
		err2 := putContentIntoFile(file,content)

		return err2
	}

	return nil
}

func putContentIntoFile(file io.Writer,content string) error {
	n, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println(n, err)
	}

	return err
}