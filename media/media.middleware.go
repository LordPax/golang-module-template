package media

import (
	"fmt"
	"golang-api/core"
	"golang-api/log"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaMiddleware struct {
	*core.Provider
	mediaService     *MediaService
	openstackService *OpenstackService
	logService       *log.LogService
}

func NewMediaMiddleware(module *MediaModule) *MediaMiddleware {
	return &MediaMiddleware{
		Provider:         core.NewProvider("MediaMiddleware"),
		mediaService:     module.Get("MediaService").(*MediaService),
		openstackService: module.Get("OpenstackService").(*OpenstackService),
		logService:       module.Get("LogService").(*log.LogService),
	}
}

func (cm *MediaMiddleware) FileUploader(availableType []string, availabeSize int, containerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var medias []*Media
		tags := []string{"MediaMiddleware", "FileUploader"}

		if err := cm.openstackService.CreateContainerIfNotExist(containerName); err != nil {
			cm.logService.Errorf(tags, "%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		form, err := c.MultipartForm()
		if err != nil {
			cm.logService.Errorf(tags, "%v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		files := form.File["upload[]"]

		for _, file := range files {
			media, err := cm.uploadFile(file, availableType, availabeSize, containerName)
			if err != nil {
				cm.logService.Errorf(tags, "%v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			medias = append(medias, media)
		}

		if len(medias) == 0 {
			cm.logService.Errorf(tags, "Failed to upload file")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
			c.Abort()
			return
		}

		c.Set("medias", medias)
		c.Next()
	}
}

func (cm *MediaMiddleware) uploadFile(
	fileHeader *multipart.FileHeader,
	availableType []string,
	availableSize int,
	containerName string,
) (*Media, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return &Media{}, err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return &Media{}, err
	}

	if fileHeader.Size > int64(availableSize) {
		return &Media{}, fmt.Errorf("File size is too large, maximum %s", FormatSize(availableSize))
	}

	if !IsFileType(buf, availableType) {
		return &Media{}, fmt.Errorf("File type is not allowed, only %v are accepted", availableType)
	}

	return cm.mediaService.UploadMedia(
		file,
		fileHeader.Filename,
		fileHeader.Header.Get("Content-Type"),
		fileHeader.Size,
		containerName,
	)
}
