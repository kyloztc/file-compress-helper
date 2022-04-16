package pkg

import (
	"file-compression-helper/pkg/helper"
	"testing"
)

func TestDirSize(t *testing.T) {
	helper.DirSizeM("/Users/kylochen/go/src/file-compression-helper")
}