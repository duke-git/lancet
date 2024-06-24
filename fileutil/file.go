// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package fileutil implements some basic functions for file operations
package fileutil

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/duke-git/lancet/v2/validator"
)

// FileReader is a reader supporting offset seeking and reading one
// line at a time, this is especially useful for large files
type FileReader struct {
	*bufio.Reader
	file   *os.File
	offset int64
}

// NewFileReader creates the FileReader struct for reading
func NewFileReader(path string) (*FileReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FileReader{
		file:   f,
		Reader: bufio.NewReader(f),
		offset: 0,
	}, nil
}

// ReadLine reads and returns one line at a time excluding the trailing '\r' and '\n'
func (f *FileReader) ReadLine() (string, error) {
	data, err := f.Reader.ReadBytes('\n')
	f.offset += int64(len(data))
	if err == nil || err == io.EOF {
		for len(data) > 0 && (data[len(data)-1] == '\r' || data[len(data)-1] == '\n') {
			data = data[:len(data)-1]
		}
		return string(data), err
	}
	return "", err
}

// Offset returns the current offset of the file
func (f *FileReader) Offset() int64 {
	return f.offset
}

// SeekOffset sets the current offset of the reading
func (f *FileReader) SeekOffset(offset int64) error {
	_, err := f.file.Seek(offset, 0)
	if err != nil {
		return err
	}
	f.Reader = bufio.NewReader(f.file)
	f.offset = offset
	return nil
}

// Close takes care of the opened file
func (f *FileReader) Close() error {
	return f.file.Close()
}

// IsExist checks if a file or directory exists.
// Play: https://go.dev/play/p/nKKXt8ZQbmh
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateFile create a file in path.
// Play: https://go.dev/play/p/lDt8PEsTNKI
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// CreateDir create directory in absolute path. param `absPath` like /a/, /a/b/.
// Play: https://go.dev/play/p/qUuCe1OGQnM
func CreateDir(absPath string) error {
	// return os.MkdirAll(path.Dir(absPath), os.ModePerm)
	return os.MkdirAll(absPath, os.ModePerm)
}

// CopyDir copy	src directory to dst directory, it will copy all files and directories recursively.
// the access permission will be the same as the source directory.
// if dstPath exists, it will return an error.
// Play: https://go.dev/play/p/YAyFTA_UuPb
func CopyDir(srcPath string, dstPath string) error {
	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return fmt.Errorf("failed to get source directory info: %w", err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source path is not a directory: %s", srcPath)
	}

	err = os.MkdirAll(dstPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	entries, err := os.ReadDir(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	for _, entry := range entries {
		srcDir := filepath.Join(srcPath, entry.Name())
		dstDir := filepath.Join(dstPath, entry.Name())

		if entry.IsDir() {
			err := CopyDir(srcDir, dstDir)
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(srcDir, dstDir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// IsDir checks if the path is directory or not.
// Play: https://go.dev/play/p/WkVwEKqtOWk
func IsDir(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}

// RemoveFile remove the path file.
// Play: https://go.dev/play/p/P2y0XW8a1SH
func RemoveFile(path string) error {
	return os.Remove(path)
}

// CopyFile copy src file to dest file.
// Play: https://go.dev/play/p/Jg9AMJMLrJi
func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	distFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer distFile.Close()

	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		_, err = distFile.Write(tmp[:n])
		if err != nil {
			return err
		}
	}
}

// ClearFile write empty string to path file.
// Play: https://go.dev/play/p/NRZ0ZT-G94H
func ClearFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

// ReadFileToString return string of file content.
// Play: https://go.dev/play/p/cmfwp_5SQTp
func ReadFileToString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileByLine read file line by line.
// Play: https://go.dev/play/p/svJP_7ZrBrD
func ReadFileByLine(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	result := make([]string, 0)
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := string(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		result = append(result, l)
	}

	return result, nil
}

// ListFileNames return all file names in the path.
// Play: https://go.dev/play/p/Tjd7Y07rejl
func ListFileNames(path string) ([]string, error) {
	if !IsExist(path) {
		return []string{}, nil
	}

	fs, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	result := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			result = append(result, fs[i].Name())
		}
	}

	return result, nil
}

// IsZipFile checks if file is zip or not.
// Play: https://go.dev/play/p/9M0g2j_uF_e
func IsZipFile(filepath string) bool {
	f, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// Zip create zip file, fpath could be a single file or a directory.
// Play: https://go.dev/play/p/j-3sWBp8ik_P
func Zip(path string, destPath string) error {
	if IsDir(path) {
		return zipFolder(path, destPath)
	}

	return zipFile(path, destPath)
}

func zipFile(filePath string, destPath string) error {
	zipFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return addFileToArchive1(filePath, archive)
}

func zipFolder(folderPath string, destPath string) error {
	outFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	err = addFileToArchive2(w, folderPath, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func addFileToArchive1(fpath string, archive *zip.Writer) error {
	err := filepath.Walk(fpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(fpath)+"/")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(writer, file); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func addFileToArchive2(w *zip.Writer, basePath, baseInZip string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := os.ReadFile(basePath + file.Name())
			if err != nil {
				return err
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			addFileToArchive2(w, newBase, baseInZip+file.Name()+"/")
		}
	}

	return nil
}

// UnZip unzip the file and save it to destPath.
// Play: https://go.dev/play/p/g0w34kS7B8m
func UnZip(zipFile string, destPath string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		// issue#62: fix ZipSlip bug
		path, err := safeFilepathJoin(destPath, f.Name)
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ZipAppendEntry append a single file or directory by fpath to an existing zip file.
// Play: https://go.dev/play/p/cxvaT8TRNQp
func ZipAppendEntry(fpath string, destPath string) error {
	tempFile, err := os.CreateTemp("", "temp.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	zipReader, err := zip.OpenReader(destPath)
	if err != nil {
		return err
	}

	archive := zip.NewWriter(tempFile)

	for _, zipItem := range zipReader.File {
		zipItemReader, err := zipItem.Open()
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(zipItem.FileInfo())
		if err != nil {
			return err
		}
		header.Name = zipItem.Name
		targetItem, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(targetItem, zipItemReader)
		if err != nil {
			return err
		}
	}

	err = addFileToArchive1(fpath, archive)

	if err != nil {
		return err
	}

	err = zipReader.Close()
	if err != nil {
		return err
	}
	err = archive.Close()
	if err != nil {
		return err
	}
	err = tempFile.Close()
	if err != nil {
		return err
	}

	return CopyFile(tempFile.Name(), destPath)
}

func safeFilepathJoin(path1, path2 string) (string, error) {
	relPath, err := filepath.Rel(".", path2)
	if err != nil || strings.HasPrefix(relPath, "..") {
		return "", fmt.Errorf("(zipslip) filepath is unsafe %q: %v", path2, err)
	}
	if path1 == "" {
		path1 = "."
	}
	return filepath.Join(path1, filepath.Join("/", relPath)), nil
}

// IsLink checks if a file is symbol link or not.
// Play: https://go.dev/play/p/TL-b-Kzvf44
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// FileMode return file's mode and permission.
// Play: https://go.dev/play/p/2l2hI42fA3p
func FileMode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// MiMeType return file mime type
// param `file` should be string(file path) or *os.File.
// Play: https://go.dev/play/p/bd5sevSUZNu
func MiMeType(file any) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}
	return mediatype
}

// CurrentPath return current absolute path.
// Play: https://go.dev/play/p/s74a9iBGcSw
func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = filepath.Dir(filename)
	}

	return absPath
}

// FileSize returns file size in bytes.
// Play: https://go.dev/play/p/H9Z05uD-Jjc
func FileSize(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// DirSize walks the folder recursively and returns folder size in bytes.
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.WalkDir(path, func(_ string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			size += info.Size()
		}
		return err
	})
	return size, err
}

// MTime returns file modified time.
// Play: https://go.dev/play/p/s_Tl7lZoAaY
func MTime(filepath string) (int64, error) {
	f, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// Sha returns file sha value, param `shaType` should be 1, 256 or 512.
// Play: https://go.dev/play/p/VfEEcO2MJYf
func Sha(filepath string, shaType ...int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha1.New()
	if len(shaType) > 0 {
		if shaType[0] == 1 {
			h = sha1.New()
		} else if shaType[0] == 256 {
			h = sha256.New()
		} else if shaType[0] == 512 {
			h = sha512.New()
		} else {
			return "", errors.New("param `shaType` should be 1, 256 or 512")
		}
	}

	_, err = io.Copy(h, file)

	if err != nil {
		return "", err
	}

	sha := fmt.Sprintf("%x", h.Sum(nil))

	return sha, nil

}

// ReadCsvFile read file content into slice.
// Play: https://go.dev/play/p/OExTkhGEd3_u
func ReadCsvFile(filepath string, delimiter ...rune) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	if len(delimiter) > 0 {
		reader.Comma = delimiter[0]
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// WriteCsvFile write content to target csv file.
// append: append to existing csv file
// delimiter: specifies csv delimiter
// Play: https://go.dev/play/p/dAXm58Q5U1o
func WriteCsvFile(filepath string, records [][]string, append bool, delimiter ...rune) error {
	flag := os.O_RDWR | os.O_CREATE

	if append {
		flag = flag | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	writer := csv.NewWriter(f)
	// 设置默认分隔符为逗号，除非另外指定
	if len(delimiter) > 0 {
		writer.Comma = delimiter[0]
	} else {
		writer.Comma = ','
	}

	// 遍历所有记录并处理包含分隔符或双引号的单元格
	for i := range records {
		for j := range records[i] {
			records[i][j] = escapeCSVField(records[i][j], writer.Comma)
		}
	}

	return writer.WriteAll(records)
}

// WriteStringToFile write string to target file.
// Play: https://go.dev/play/p/GhLS6d8lH_g
func WriteStringToFile(filepath string, content string, append bool) error {
	var flag int
	if append {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	f, err := os.OpenFile(filepath, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

// WriteBytesToFile write bytes to target file.
// Play: https://go.dev/play/p/s7QlDxMj3P8
func WriteBytesToFile(filepath string, content []byte) error {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)
	return err
}

// ReadFile get file reader by a url or a local file
// Play: https://go.dev/play/p/uNep3Tr8fqF
func ReadFile(path string) (reader io.ReadCloser, closeFn func(), err error) {
	switch {
	case validator.IsUrl(path):
		resp, err := http.Get(path)
		if err != nil {
			return nil, func() {}, err
		}
		return resp.Body, func() { resp.Body.Close() }, nil
	case IsExist(path):
		reader, err := os.Open(path)
		if err != nil {
			return nil, func() {}, err
		}
		return reader, func() { reader.Close() }, nil
	default:
		return nil, func() {}, errors.New("unknown file type")
	}
}

// escapeCSVField 处理单元格内容，如果包含分隔符，则用双引号包裹
func escapeCSVField(field string, delimiter rune) string {
	// 替换所有的双引号为两个双引号
	escapedField := strings.ReplaceAll(field, "\"", "\"\"")

	// 如果字段包含分隔符、双引号或换行符，用双引号包裹整个字段
	if strings.ContainsAny(escapedField, string(delimiter)+"\"\n") {
		escapedField = fmt.Sprintf("\"%s\"", escapedField)
	}

	return escapedField
}

// WriteMapsToCsv write slice of map to csv file.
// Play: https://go.dev/play/p/umAIomZFV1c
// filepath: Path to the CSV file.
// records: Slice of maps to be written. the value of map should be basic type.
// the maps will be sorted by key in alphabeta order, then be written into csv file.
// appendToExistingFile: If true, data will be appended to the file if it exists.
// delimiter: Delimiter to use in the CSV file.
// headers: order of the csv column headers, needs to be consistent with the key of the map.
func WriteMapsToCsv(filepath string, records []map[string]any, appendToExistingFile bool, delimiter rune,
	headers ...[]string) error {
	for _, record := range records {
		for _, value := range record {
			if !isCsvSupportedType(value) {
				return errors.New("unsupported value type detected; only basic types are supported: \nbool, rune, string, int, int64, float32, float64, uint, byte, complex128, complex64, uintptr")
			}
		}
	}

	var columnHeaders []string
	if len(headers) > 0 {
		columnHeaders = headers[0]
	} else {
		for key := range records[0] {
			columnHeaders = append(columnHeaders, key)
		}
		// sort keys in alphabeta order
		sort.Strings(columnHeaders)
	}

	var datasToWrite [][]string
	if !appendToExistingFile {
		datasToWrite = append(datasToWrite, columnHeaders)
	}

	for _, record := range records {
		var row []string
		for _, h := range columnHeaders {
			row = append(row, fmt.Sprintf("%v", record[h]))
		}
		datasToWrite = append(datasToWrite, row)
	}

	return WriteCsvFile(filepath, datasToWrite, appendToExistingFile, delimiter)
}

// check if the value of map which to be written into csv is basic type.
func isCsvSupportedType(v interface{}) bool {
	switch v.(type) {
	case bool, rune, string, int, int64, float32, float64, uint, byte, complex128, complex64, uintptr:
		return true
	default:
		return false
	}
}

// ChunkRead reads a block from the file at the specified offset and returns all lines within the block
// Play: https://go.dev/play/p/r0hPmKWhsgf
func ChunkRead(file *os.File, offset int64, size int, bufPool *sync.Pool) ([]string, error) {
	buf := bufPool.Get().([]byte)[:size] // 从Pool获取缓冲区并调整大小
	n, err := file.ReadAt(buf, offset)   // 从指定偏移读取数据到缓冲区
	if err != nil && err != io.EOF {
		return nil, err
	}
	buf = buf[:n] // 调整切片以匹配实际读取的字节数

	var lines []string
	var lineStart int
	for i, b := range buf {
		if b == '\n' {
			line := string(buf[lineStart:i]) // 不包括换行符
			lines = append(lines, line)
			lineStart = i + 1 // 设置下一行的开始
		}
	}

	if lineStart < len(buf) { // 处理块末尾的行
		line := string(buf[lineStart:])
		lines = append(lines, line)
	}
	bufPool.Put(buf) // 读取完成后，将缓冲区放回Pool
	return lines, nil
}

// ParallelChunkRead reads the file in parallel and send each chunk of lines to the specified channel.
// filePath 文件路径
// chunkSizeMB 分块的大小（单位MB，设置为0时使用默认100MB）,设置过大反而不利，视情调整
// maxGoroutine 并发读取分块的数量，设置为0时使用CPU核心数
// linesCh用于接收返回结果的通道。
// Play: https://go.dev/play/p/teMXnCsdSEw
func ParallelChunkRead(filePath string, linesCh chan<- []string, chunkSizeMB, maxGoroutine int) error {
	if chunkSizeMB == 0 {
		chunkSizeMB = 100
	}
	chunkSize := chunkSizeMB * 1024 * 1024
	// 内存复用
	bufPool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, chunkSize)
		},
	}

	if maxGoroutine == 0 {
		maxGoroutine = runtime.NumCPU() // 设置为0时使用CPU核心数
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	chunkOffsetCh := make(chan int64, maxGoroutine)

	// 分配工作
	go func() {
		for i := int64(0); i < info.Size(); i += int64(chunkSize) {
			chunkOffsetCh <- i
		}
		close(chunkOffsetCh)
	}()

	// 启动工作协程
	for i := 0; i < maxGoroutine; i++ {
		wg.Add(1)
		go func() {
			for chunkOffset := range chunkOffsetCh {
				chunk, err := ChunkRead(f, chunkOffset, chunkSize, &bufPool)
				if err == nil {
					linesCh <- chunk
				}

			}
			wg.Done()
		}()
	}

	// 等待所有解析完成后关闭行通道
	wg.Wait()
	close(linesCh)

	return nil
}
