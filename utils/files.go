package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const bytesToKB = 0.001
const bytestoMB = 0.0000001

type FileSize struct {
	Name string
	KB   float64
	MB   float64
}

type DirectorySize struct {
	Files   []FileSize
	Name    string
	TotalKB float64
	TotalMB float64
}

func GetFileSize(file *os.File) int64 {
	info, _ := file.Stat()
	size := info.Size()
	return size
}

func ConvertFileSize(bytes int64) (float64, float64) {
	bytesFloat := float64(bytes)
	KB := bytesFloat * bytesToKB
	MB := bytesFloat * bytestoMB
	return KB, MB
}

func GetPathFromCurrent(path string) string {
	cwd, _ := os.Getwd()
	fromCwd := strings.ReplaceAll(path, cwd, "")
	return fromCwd
}

func GetFilePathsByExt(ext string) []string {
	currentDir, _ := os.Getwd()
	var files []string
	filepath.WalkDir(currentDir, func(path string, info fs.DirEntry, er error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetFileSizeByExt(ext string) map[string]DirectorySize {
	output := make(map[string]DirectorySize)
	var paths []string = GetFilePathsByExt(ext)
	for _, path := range paths {
		dir := filepath.Dir(path)
		fromCwd := GetPathFromCurrent(dir)
		file, _ := os.Open(path)
		bytes := GetFileSize(file)
		KB, MB := ConvertFileSize(bytes)
		filename := GetPathFromCurrent(file.Name())
		fsize := FileSize{filename, KB, MB}
		if dsize, ok := output[fromCwd]; ok {
			dsize.TotalKB += fsize.KB
			dsize.TotalMB += fsize.MB
			dsize.Files = append(dsize.Files, fsize)
			output[fromCwd] = dsize
			output[fromCwd] = dsize
		} else {
			var fslice []FileSize
			fslice = append(fslice, fsize)
			output[fromCwd] = DirectorySize{
				Name:    fromCwd,
				Files:   fslice,
				TotalKB: fsize.KB,
				TotalMB: fsize.MB,
			}
		}
	}
	return output
}
