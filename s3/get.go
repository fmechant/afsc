package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/viant/afs/file"
	"github.com/viant/afs/option"
	"github.com/viant/afs/storage"
	"os"
	"path"
	"strings"
	"time"
)

//Get returns an object for supplied location
func (s *storager) get(ctx context.Context, location string, options []storage.Option) (os.FileInfo, error) {
	location = strings.Trim(location, "/")
	_, name := path.Split(location)

	object, err := s.GetObject(&s3.GetObjectInput{
		Bucket: &s.bucket,
		Key:    &location,
	})
	if err != nil {
		return nil, err
	}
	hasObject := object != nil && (object.ContentLength != nil || object.LastModified != nil)
	if !hasObject {
		return nil, fmt.Errorf(noSuchKeyMessage + " " + location)
	}
	contentLength := int64(0)
	modified := time.Now()
	if object.LastModified != nil {
		modified = *object.LastModified
	}
	if object.ContentLength != nil {
		contentLength = *object.ContentLength
	}
	if object.Body != nil {
		_ = object.Body.Close()
	}
	return file.NewInfo(name, contentLength, file.DefaultFileOsMode, modified, false, object), nil
}

//Get returns an object for supplied location
func (s *storager) Get(ctx context.Context, location string, options ...storage.Option) (os.FileInfo, error) {
	location = strings.Trim(location, "/")
	info, err := s.get(ctx, location, options)
	if err == nil {
		return info, err
	}
	if isNotFound(err) {
		objectKind := &option.ObjectKind{}
		if _, ok := option.Assign(options, objectKind); ok && objectKind.File {
			return nil, err
		}
	}
	options = append(options, option.NewPage(0, 1))
	objects, err := s.List(ctx, location, options...)
	if err != nil {
		return nil, err
	}
	if len(objects) == 0 {
		return nil, errors.Errorf("%v %v", location, doesNotExistsMessage)
	}
	return objects[0], err
}
