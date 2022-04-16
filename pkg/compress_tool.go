package pkg

import (
	"errors"
	"file-compression-helper/pkg/helper"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func SplitFolder(path string, limit float64, outDir string, subFolderName string) error {
	path, outDir, subFolderName, err := pathFix(path, outDir, subFolderName)
	if err != nil {
		fmt.Printf("pre check error|%v\n", err)
		return err
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("read dir error: %v|dir path: %v\n", err, path)
		return err
	}
	fmt.Printf("read folder: %v|file/dir num: %v\n", path, len(files))
	totalSize := float64(0)
	fileList := make([]*os.File, 0)
	zipSuffix := 0
	for _, file := range files {
		filePath := fmt.Sprintf("%s%s", path, file.Name())
		fs, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("open file error: %v|file path: %v\n", err, filePath)
			return err
		}
		fileList = append(fileList, fs)
		fileSize := helper.Byte2M(file.Size())
		if file.IsDir() {
			fileSize, _ = helper.DirSizeM(filePath)
		}
		totalSize += fileSize
		fmt.Printf("add file: %v|size: %vM\n", filePath, fileSize)
		if totalSize <= limit {
			continue
		}
		err = compressFile(fileList, subFolderName, zipSuffix, outDir, totalSize)
		if err != nil {
			return err
		}
		zipSuffix += 1
		fileList = make([]*os.File, 0)
		totalSize = 0
	}
	err = compressFile(fileList, subFolderName, zipSuffix, outDir, totalSize)
	if err != nil {
		return err
	}
	return nil
}

func compressFile(fileList []*os.File, subFolderName string, zipSuffix int, outDir string, totalSize float64) error {
	outputFile := fmt.Sprintf("%v%v-%d.zip", outDir, subFolderName, zipSuffix)
	fmt.Printf("compress file: %v|files size: %v\n", outputFile, totalSize)
	err := helper.Compress(fileList, outputFile)
	if err != nil {
		fmt.Printf("compress file err: %v\n", err)
		return err
	}
	return nil
}

func pathFix(path string, outDir string, subFolderName string) (string, string, string, error) {
	if path == "" || outDir == "" {
		return "", "", "", errors.New("path or out dir empty")
	}
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	if !strings.HasSuffix(outDir, "/") {
		outDir = outDir + "/"
	}
	_, err := os.Stat(outDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(outDir, os.ModePerm)
		if err != nil {
			return "", "", "", err
		}
	}
	if subFolderName == "" {
		rootDir := strings.Split(path, "/")
		subFolderName = fmt.Sprintf("%s-%d", rootDir[len(rootDir) - 2], time.Now().Unix())
	}
	return path, outDir, subFolderName, nil
}


