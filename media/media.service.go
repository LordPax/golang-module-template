package media

import (
	"golang-api/core"
	"golang-api/query"
	"io"
	"path/filepath"

	"github.com/google/uuid"
)

// type IMediaService interface {
// 	core.IProvider
// 	FindAll(query.QueryFilter) ([]*Media, error)
// 	FindByID(string) (*Media, error)
// 	FindOneBy(string, any) (*Media, error)
// 	Create(*Media) error
// }

type MediaService struct {
	*core.Provider
	mediaModel       *MediaModel
	openstackService *OpenstackService
}

func NewMediaService(module *MediaModule) *MediaService {
	return &MediaService{
		Provider:         core.NewProvider("MediaService"),
		mediaModel:       module.Get("MediaModel").(*MediaModel),
		openstackService: module.Get("OpenstackService").(*OpenstackService),
	}
}

func (us *MediaService) FindAll(query query.QueryFilter) ([]*Media, error) {
	return us.mediaModel.FindAll(query)
}

func (us *MediaService) FindByID(id string) (*Media, error) {
	return us.mediaModel.FindByID(id)
}

func (us *MediaService) FindOneBy(field string, value any) (*Media, error) {
	return us.mediaModel.FindOneBy(field, value)
}

func (us *MediaService) Create(media *Media) error {
	return us.mediaModel.Create(media)
}

func (cs *MediaService) UploadMedia(
	file io.Reader,
	fileName string,
	fileType string,
	fileSize int64,
	containerName string,
) (*Media, error) {
	objecName := uuid.New().String() + filepath.Ext(fileName)
	url, err := cs.openstackService.UploadFile(file, objecName, containerName)
	if err != nil {
		return &Media{}, err
	}

	media := &Media{
		Name:      objecName,
		Size:      fileSize,
		Type:      fileType,
		Url:       url,
		Container: containerName,
	}
	if err := cs.Create(media); err != nil {
		return &Media{}, err
	}

	return media, nil
}
