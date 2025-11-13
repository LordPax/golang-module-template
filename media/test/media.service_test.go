package media_test

import (
	"golang-api/core"
	"golang-api/media"
	"golang-api/query"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMedia(t *testing.T, expected *media.Media, actual *media.Media) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Url, actual.Url)
	assert.Equal(t, expected.Type, actual.Type)
	assert.Equal(t, expected.Size, actual.Size)
	assert.Equal(t, expected.Container, actual.Container)
	assert.Equal(t, expected.UserID, actual.UserID)
}

func NewMediaModuleTest() *media.MediaModule {
	module := &media.MediaModule{
		Module: core.NewModule("MediaTestModule"),
	}

	module.AddProvider(NewMediaModelMock(module))
	module.AddProvider(NewOpenstackServiceMock(module))
	module.AddProvider(media.NewMediaService(module))

	return module
}

func TestSetupMediaModule(t *testing.T) {
	module := NewMediaModuleTest()
	mediaModel := module.Get("MediaModel").(*MediaModelMock)
	openstackService := module.Get("OpenstackService").(*OpenstackServiceMock)
	mediaService := module.Get("MediaService").(*media.MediaService)

	assert.NotNil(t, module, "Media module should be created")
	assert.NotNil(t, mediaModel, "MediaModel should be created")
	assert.NotNil(t, openstackService, "OpenstackService should be created")
	assert.NotNil(t, mediaService, "MediaService should be created")
}

func TestMediaService_FindAll(t *testing.T) {
	module := NewMediaModuleTest()
	mediaService := module.Get("MediaService").(*media.MediaService)
	mediaModel := module.Get("MediaModel").(*MediaModelMock)

	q := query.QueryFilter{}
	nbMedias := 3
	expectedMedias := CreateManyMedias(nbMedias)

	mediaModel.MockMethod("QueryFindAll", func(params ...any) any { return expectedMedias })

	newMedias, _ := mediaService.FindAll(q)

	called := mediaModel.IsMethodCalled("QueryFindAll")
	if !assert.Equal(t, true, called, "QueryFindAll method should be called") {
		return
	}

	if !assert.Len(t, newMedias, nbMedias, "Number of medias should be equal to expected") {
		return
	}
	for i := 0; i < nbMedias; i++ {
		testMedia(t, expectedMedias[i], newMedias[i])
	}
}

func TestMediaService_FindByID(t *testing.T) {
	module := NewMediaModuleTest()
	mediaService := module.Get("MediaService").(*media.MediaService)
	mediaModel := module.Get("MediaModel").(*MediaModelMock)

	expectedMedia := CreateMedia()

	mediaModel.MockMethod("FindByID", func(params ...any) any {
		return expectedMedia
	})

	newMedia, _ := mediaService.FindByID(expectedMedia.ID)

	called := mediaModel.IsMethodCalled("FindByID")
	if !assert.Equal(t, true, called, "FindByID method should be called") {
		return
	}
	params := mediaModel.IsParamsEqual("FindByID", expectedMedia.ID)
	if !assert.Equal(t, true, params, "FindByID parameter should be the media ID") {
		return
	}

	testMedia(t, expectedMedia, newMedia)
}

func TestMediaService_FindOneBy(t *testing.T) {
	module := NewMediaModuleTest()
	mediaService := module.Get("MediaService").(*media.MediaService)
	mediaModel := module.Get("MediaModel").(*MediaModelMock)

	expectedMedia := CreateMedia()

	mediaModel.MockMethod("FindOneBy", func(params ...any) any {
		return expectedMedia
	})

	newMedia, _ := mediaService.FindOneBy("type", expectedMedia.Type)

	called := mediaModel.IsMethodCalled("FindOneBy")
	if !assert.Equal(t, true, called, "FindOneBy method should be called") {
		return
	}
	params := mediaModel.IsParamsEqual("FindOneBy", "type", expectedMedia.Type)
	if !assert.Equal(t, true, params, "FindOneBy parameters should be the field and value") {
		return
	}

	testMedia(t, expectedMedia, newMedia)
}

func TestMediaService_Create(t *testing.T) {
	module := NewMediaModuleTest()
	mediaService := module.Get("MediaService").(*media.MediaService)
	mediaModel := module.Get("MediaModel").(*MediaModelMock)

	expectedMedia := CreateMedia()
	var createdMedia *media.Media

	mediaModel.MockMethod("Create", func(params ...any) any {
		createdMedia = params[0].(*media.Media)
		return nil
	})

	err := mediaService.Create(expectedMedia)
	if !assert.Nil(t, err, "Create should not return an error") {
		return
	}

	called := mediaModel.IsMethodCalled("Create")
	if !assert.Equal(t, true, called, "Create method should be called") {
		return
	}
	params := mediaModel.GetMethodParams("Create")
	if !assert.Len(t, params, 1, "Create should have one parameter") {
		return
	}
	paramMedia, ok := params[0].(*media.Media)
	if !assert.Equal(t, true, ok, "Create parameter should be a media") {
		return
	}

	testMedia(t, expectedMedia, paramMedia)
	testMedia(t, expectedMedia, createdMedia)
}

// func TestMediaService_UploadMedia(t *testing.T) {
// 	module := NewMediaModuleTest()
// 	mediaService := module.Get("MediaService").(*media.MediaService)
// 	mediaModel := module.Get("MediaModel").(*MediaModelMock)
// 	openstackService := module.Get("OpenstackService").(*OpenstackServiceMock)

// 	expectedMedia := CreateMedia()
// 	fmt.Printf("file: %s, name: %s, container: %s\n", expectedMedia.Url, expectedMedia.Name, expectedMedia.Container)

// 	openstackService.MockMethod("UploadFile", func(params ...any) any {
// 		fmt.Println(params)
// 		return expectedMedia.Url
// 	})

// 	mediaModel.MockMethod("Create", func(params ...any) any {
// 		return nil
// 	})

// 	fileContent := "This is a test file"

// 	newMedia, err := mediaService.UploadMedia(
// 		strings.NewReader(fileContent),
// 		expectedMedia.Name,
// 		expectedMedia.Type,
// 		expectedMedia.Size,
// 		expectedMedia.Container,
// 	)
// 	fmt.Printf("file: %s, name: %s, container: %s\n", newMedia.Url, newMedia.Name, newMedia.Container)

// 	if !assert.Nil(t, err, "UploadMedia should not return an error") {
// 		return
// 	}

// 	calledUpload := openstackService.IsMethodCalled("UploadFile")
// 	if !assert.Equal(t, true, calledUpload, "UploadFile method should be called") {
// 		return
// 	}

// 	calledCreate := mediaModel.IsMethodCalled("Create")
// 	if !assert.Equal(t, true, calledCreate, "Create method should be called") {
// 		return
// 	}
// 	// paramsCreate := mediaModel.IsParamsEqual("Create", expectedMedia)
// 	// if !assert.Equal(t, true, paramsCreate, "Create parameter should be the new media") {
// 	// 	return
// 	// }

// 	// testMedia(t, expectedMedia, newMedia)
// }
