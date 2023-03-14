package poster

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"image/jpeg"
	"os"
	"testing"
)

func Test(t *testing.T) {
	file, err := os.Create("test.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	str := ""
	resp, err := GetQRImage(str, qrcode.Medium, 100)
	if err != nil {
		fmt.Println(err)
	} else {
		err = jpeg.Encode(file, resp, nil)
		if err != nil {
			fmt.Println(err)
		}
	}

}
