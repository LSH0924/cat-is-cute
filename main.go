package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// 귀여운 고양이 사진을 보내주는 api
	r := gin.Default()
	r.GET("/tuxedo", func(c *gin.Context) {
		img, _ := getRandomImageIn("./assets/images/tuxedo")
		c.String(200, img)
	})
	r.GET("/cheese", func(c *gin.Context) {
		img, _ := getRandomImageIn("./assets/images/cheese")
		c.String(200, img)
	})
	r.Run()
}

func getRandomImageIn(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	info := files[rand.Int31n(int32(len(files)))]
	return convertToBase64(filepath.Join(path, info.Name()))
}

func convertToBase64(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	r := bufio.NewReader(f)
	content, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(content), nil
}
