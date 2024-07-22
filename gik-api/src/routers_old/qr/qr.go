package qr

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

func GetQRCodes(c *gin.Context) {

	textIn := c.Query("labels")

	var text []string

	for _, word := range strings.Fields(textIn) {
		text = append(text, word)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	var err error

	count := 0.0
	max := float64(len(text))

	err = pdf.AddTTFFont("roboto", "assets/ttf/Roboto-Light.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("roboto", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	for {
		for i := 0.0; i < 10.0; i++ {
			for k := 0.0; k < 5.0; k++ {
				var png []byte

				qr, _ := qrcode.New(text[int(count)], qrcode.Medium)

				qr.DisableBorder = true

				png, _ = qr.PNG(120)

				img, _ := gopdf.ImageHolderByBytes(png)

				pdf.ImageByHolder(img, (2+3.8*k)*72.0/2.54, (1.37+2.543*i)*72.0/2.54, nil) //print image

				pdf.SetXY((2.1+2.543*i)*72.0/2.54, -1*(-2.15+3.8*k)*72.0/2.54) //move current location
				pdf.Rotate(270.0, 100.0, 100.0)
				pdf.Cell(nil, text[int(count)]) //print text
				pdf.RotateReset()

				count += 1
				if count >= max {
					break
				}
			}
			if count >= max {
				break
			}
		}
		if count >= max {
			break
		}
		pdf.AddPage()
	}

	if err != nil {
		log.Print(err.Error())
		return
	}

	name := uuid.New().String() + ".pdf"

	pdf.WritePdf(name)

	c.File(name)

	os.Remove(name)

}
