package common

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"github.com/shifty11/cosmos-notifier/log"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ImageManager struct {
	Name          string
	Description   string
	ThumbnailPath string
	ThumbnailUrl  string
	DownloadUrl   string
	Width         int
	Height        int
}

func isSVG(url string) bool {
	split := strings.Split(url, ".")
	return split[len(split)-1] == "svg"
}

func NewImageManager(name string, description string, url string, weblocation string, imagePath string, width int, height int) *ImageManager {
	extension := ".png"
	//if isSVG(url) {
	//	extension = ".svg"
	//}
	return &ImageManager{
		Name:          name,
		Description:   description,
		ThumbnailPath: filepath.Join(weblocation, imagePath, name+extension),
		ThumbnailUrl:  filepath.Join(imagePath, name+extension),
		DownloadUrl:   url,
		Width:         width,
		Height:        height,
	}
}

func (im *ImageManager) DoesExist() bool {
	if _, err := os.Stat(im.ThumbnailPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (im *ImageManager) isImageFiletype(data []byte) bool {
	filetype := http.DetectContentType(data)
	return strings.Split(filetype, "/")[0] == "image"
}

func ensureDir(path string) {
	dirName := filepath.Dir(path)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}

func (im *ImageManager) createThumbnail(data []byte) error {
	srcImage, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}
	ensureDir(im.ThumbnailPath)
	thumb := imaging.Thumbnail(srcImage, im.Width, im.Height, imaging.Lanczos)
	err = imaging.Save(thumb, im.ThumbnailPath)
	if err != nil {
		return err
	}
	return nil
}

func (im *ImageManager) saveAsSVG(data []byte) error {
	ensureDir(im.ThumbnailPath)
	file, err := os.Create(im.ThumbnailPath)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (im *ImageManager) saveAsPNG(data []byte) error {
	w, h := 512, 512
	tmpFilePath := "/tmp/out.png"
	icon, _ := oksvg.ReadIconStream(bytes.NewReader(data))
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		return err
	}

	pngData, err := os.ReadFile(tmpFilePath)
	if err != nil {
		return err
	}

	if im.isImageFiletype(pngData) {
		return im.createThumbnail(pngData)
	}
	return errors.New("not an image filetype")
}

func (im *ImageManager) DownloadAndCreateThumbnail() error {
	log.Sugar.Debugf("downloading image for %v (%v): %v", im.Description, im.Name, im.DownloadUrl)
	resp, err := http.Get(im.DownloadUrl)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if isSVG(im.DownloadUrl) {
		return im.saveAsPNG(data)
	}

	if im.isImageFiletype(data) {
		return im.createThumbnail(data)
	}
	return errors.New("not an image filetype")
}
