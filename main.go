package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
	"xatu/excel"
)

func getImg(url string,name string) (n int64, err error) {
	imgPath := "./photo/"
	fileName := path.Base(url)
	fileName=name+fileName
	fmt.Println(fileName,"downloading")
	out, err := os.Create(imgPath + fileName)
	defer out.Close()
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	fmt.Println(fileName,"finished")
	return

}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func test() {
	//imgUrl := "http://jwgl.xatu.edu.cn/photo/160205171/16020517128.jpg"
	for i:=1;i>=0;i++{
		imgUrl := "http://jwgl.xatu.edu.cn/photo/"
		m:=RandInt64(2400,2600)
		if m!=2790{
			imgUrl+=excel.GetClass(int(m))+"/"+excel.GetNum(int(m))+".jpg"
			_, _ = getImg(imgUrl, excel.GetName(int(m)))
		}
	}
}

func main(){
	for i :=0 ; i<5 ;i++{
		go test()  //启
	}
	time.Sleep(time.Second * 600)
}