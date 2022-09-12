package crawler

import (
	"bytes"
	"errors"
	imaging "github.com/disintegration/imaging"
	"github.com/shifty11/dao-dao-notifier/log"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ImageManager struct {
	ContractAddress string
	ThumbnailPath   string
	ThumbnailUrl    string
	Width           int
	Height          int
}

func NewImageManager(contractAddress string, weblocation string, imagePath string, width int, height int) *ImageManager {
	return &ImageManager{
		ThumbnailPath:   filepath.Join(weblocation, imagePath, contractAddress+".png"),
		ThumbnailUrl:    filepath.Join(imagePath, contractAddress+".png"),
		ContractAddress: contractAddress,
		Width:           width,
		Height:          height,
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

func (im *ImageManager) downloadAndCreateThumbnail(url string) error {
	log.Sugar.Debugf("downloading image for %v: %v", im.ContractAddress, url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if im.isImageFiletype(data) {
		return im.createThumbnail(data)
	}
	return errors.New("not an image filetype")
}
