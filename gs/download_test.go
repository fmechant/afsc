package gs

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/afs/asset"
	"github.com/viant/afs/url"
	"io/ioutil"
	"testing"
)

func TestStorager_Download(t *testing.T) {
	jwtConfig, err := NewTestJwtConfig()
	if err != nil {
		t.Skip(err)
		return
	}
	ctx := context.Background()
	var useCases = []struct {
		description string
		URL         string
		assets      []*asset.Resource
	}{
		{
			description: "single asset download",
			URL:         fmt.Sprintf("gs://%v/", TestBucket),
			assets: []*asset.Resource{
				asset.NewFile("download/asset1.txt", []byte("test is test 1 "), 0655),
			},
		},
		{
			description: "multi asset download",
			URL:         fmt.Sprintf("gs://%v/", TestBucket),
			assets: []*asset.Resource{
				asset.NewFile("download/folder1/asset1.txt", []byte("test is test 2"), 0655),
				asset.NewFile("download/folder1/asset2.txt", []byte("test is test 3"), 0655),
			},
		},
	}
	mgr := New(jwtConfig)
	defer mgr.Delete(ctx, fmt.Sprintf("gs://%v/", TestBucket))
	for _, useCase := range useCases {

		err = asset.Create(mgr, useCase.URL, useCase.assets)
		assert.Nil(t, err, useCase.description)

		for _, asset := range useCase.assets {
			reader, err := mgr.DownloadWithURL(ctx, url.Join(useCase.URL, asset.Name))
			if !assert.Nil(t, err, useCase.description) {
				continue
			}
			data, err := ioutil.ReadAll(reader)
			assert.EqualValues(t, asset.Data, data, useCase.description+" "+asset.Name)

		}

	}

}
