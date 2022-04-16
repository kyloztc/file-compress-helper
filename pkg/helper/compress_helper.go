package helper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//getFileSize get file size by path(M)
func DirSizeM(path string) (float64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//fmt.Printf("path: %v|size: %v\n", info.Name(), info.Size())
			size += info.Size()
		}
		return err
	})
	return Byte2M(size), err
}

func Byte2M(size int64) float64 {
	return float64(size) / 1024 /1024
}