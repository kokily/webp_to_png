package main

import (
	"fmt"
	"image/png"
	"os"
	"regexp"

	"golang.org/x/image/webp"
)

func main() {
	var folderDir = "./"

	fileInfoList, err := os.ReadDir(folderDir)

	if err != nil {
		fmt.Println(err)
	}

	for _, fileInfo := range fileInfoList {
		var fileName = fileInfo.Name()

		fmt.Println(fileName)

		re := regexp.MustCompile("webp$")

		matched := re.MatchString(fileName)

		if !matched {
			continue
		}

		wepbFilename := folderDir + fileName

		f0, err := os.Open(wepbFilename)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer f0.Close()

		img0, err := webp.Decode(f0)

		if err != nil {
			fmt.Println(err)
			return
		}

		newFileName := fileName[:4]

		pngFile, err := os.Create("./" + newFileName + ".png")

		if err != nil {
			fmt.Println(err)
		}

		err = png.Encode(pngFile, img0)

		if err != nil {
			fmt.Println(err)
		}
	}
}
