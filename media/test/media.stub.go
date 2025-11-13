package media_test

import (
	"golang-api/media"

	"github.com/jaswdr/faker/v2"
)

var fake = faker.New()

func CreateMedia() *media.Media {
	return &media.Media{
		ID:        fake.UUID().V4(),
		Name:      fake.Lorem().Word(),
		Url:       fake.Internet().URL(),
		Type:      fake.Lorem().Word(),
		Size:      fake.Int64Between(100, 10000),
		Container: fake.Lorem().Word(),
		UserID:    fake.UUID().V4(),
	}
}

func CreateManyMedias(n int) []*media.Media {
	medias := make([]*media.Media, n)
	for i := 0; i < n; i++ {
		medias[i] = CreateMedia()
	}
	return medias
}
