package media_manager

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"service/pkg/errors"
	"strings"
)

type mediaManager struct {
	address string
}

type MediaManager interface {
	// Joins current path of the MediaManager with passed path and returns the result
	JoinPath(path string) string
	// Joins two passed path into one path
	JoinPaths(path1, path2 string) string
	// Goes back one directory
	GoBack() MediaManager
	// Goes into another folder
	//
	// # Note: You can create them if they don't exist with passing true for createPath variable
	GoTo(folderName string, createPath ...bool) (MediaManager, bool)
	// Creates a folder, if exist, ignore
	CreateFolder(folderName ...string) MediaManager
	// Removes a folder, if not found, ignore
	RemoveFolder(folderName ...string) MediaManager
	// Creates a file, if exists, ignore
	CreateFile(file multipart.File, fileName string) MediaManager
	// Removes a file, if not found, ignore
	RemoveFile(fileName string) MediaManager
	// Removes existing file(if exist) and replaces the new one
	OverwriteFile(file multipart.File, fileName string) MediaManager
	// Returns true if passed address(file or folder) exists
	Exists(name string) bool
	// Returns true if passed filename with different extension exists
	FindFileWithJustName(name string) (string, bool)
	// Return current directory
	GetAddress() string
	// Returns directory from the media root to the end file targes
	GetHostAddress(fileName string) string
	// Searchs for a file in the directory and returns it if found
	GetFile(fileName string) (file *os.File)
}

func getPwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("couldn't get address: %s", err.Error()))
	}
	for parent := pwd; true; parent = filepath.Dir(parent) {
		if _, err := os.Stat(filepath.Join(parent, "go.mod")); err == nil {
			pwd = parent
			break
		}
	}
	return pwd
}

func (m *mediaManager) Exists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *mediaManager) FindFileWithJustName(name string) (string, bool) {
	files, err := ioutil.ReadDir(m.address)
	if err != nil {
		return "", false
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		basename := filepath.Base(fileName)
		extension := filepath.Ext(fileName)
		if basename[0:len(basename)-len(extension)] == name {
			return fileName, true
		}
	}

	return "", false
}

func (m *mediaManager) TrimEnd(name string) string {
	output, _ := strings.CutSuffix(name, "/")
	return strings.TrimSpace(output)
}

func (m *mediaManager) TrimStart(name string) string {
	output, _ := strings.CutPrefix(name, "/")
	output, _ = strings.CutPrefix(output, "./")
	return strings.TrimSpace(output)
}

func (m *mediaManager) JoinPaths(path1, path2 string) string {
	return m.TrimEnd(path1) + "/" + m.TrimStart(path2)
}

func (m *mediaManager) JoinPath(path string) string {
	return m.address + "/" + m.TrimStart(path)
}

func (m *mediaManager) AddressIsValid() {
	if !m.Exists(m.address) {
		panic(fmt.Sprintf("media_manager: %s does not exist", m.address))
	}
}

func (m *mediaManager) GoBack() MediaManager {
	splitPlace := -1
	for i := len(m.address) - 1; i > -1; i-- {
		if m.address[i] == '/' {
			splitPlace = i
			break
		}
	}
	if splitPlace > 0 {
		return NewMediaManager(m.address[:splitPlace+1])
	}
	return m
}

func (m *mediaManager) GoTo(folderName string, createPath ...bool) (MediaManager, bool) {
	if m.Exists(m.JoinPath(folderName)) || (!m.Exists(m.JoinPath(folderName)) && len(createPath) > 0 && createPath[0]) {
		return NewMediaManager(m.JoinPath(folderName), createPath...), true
	}
	return nil, false
}

func (m *mediaManager) CreateFolder(folderName ...string) MediaManager {
	name := ""
	if len(folderName) > 0 {
		name = folderName[0]
	}

	path := m.address
	if name != "" {
		path = m.JoinPath(name)
	}

	if !m.Exists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("media_manager: error on creating %s directory: %s", path, err.Error()))
		}
	}
	return m
}

func (m *mediaManager) RemoveFolder(folderName ...string) MediaManager {
	name := ""
	if len(folderName) > 0 {
		name = folderName[0]
	}

	path := m.address
	if name != "" {
		path = m.JoinPath(name)
	}

	if m.Exists(path) {
		err := os.RemoveAll(path)
		if err != nil {
			panic(fmt.Sprintf("media_manager: error on removing %s directory: %s", path, err.Error()))
		}
	}
	return m
}

func (m *mediaManager) CreateFile(file multipart.File, fileName string) MediaManager {
	address := m.JoinPath(fileName)
	if !m.Exists(address) {
		outFile, err := os.OpenFile(address, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(fmt.Sprintf("media_manager: error on trying to record %s file, err: %s", fileName, err.Error()))
		}
		defer outFile.Close()

		if _, err = io.Copy(outFile, file); err != nil {
			panic(fmt.Sprintf("media_manager: error on trying to copy content of %s file into %s, err: %s", fileName, address, err.Error()))
		}
	}
	return m
}

func (m *mediaManager) RemoveFile(fileName string) MediaManager {
	address := m.JoinPath(fileName)
	if m.Exists(address) {
		err := os.Remove(address)
		if err != nil {
			panic(fmt.Sprintf("media_manager: error on trying to remove %s file, err: %s", fileName, err.Error()))
		}
	}
	return m
}

func (m *mediaManager) OverwriteFile(file multipart.File, fileName string) MediaManager {
	m.RemoveFile(fileName)
	m.CreateFile(file, fileName)
	return m
}

func (m *mediaManager) GetAddress() string {
	return m.address
}

func (m *mediaManager) GetHostAddress(filename string) string {
	return m.TrimEnd(strings.Replace(m.address, m.GetAddress(), "", 1)) + "/" + filename
}

func (m *mediaManager) GetFile(fileName string) (file *os.File) {
	file, err := os.Open(m.JoinPath(fileName))
	if err != nil {
		if os.IsNotExist(err) {
			panic(errors.New(errors.InvalidStatus, "FileDoesNotExist", fmt.Sprintf("file %s not found", fileName)))
		}

		panic(err.Error())
	}

	return file
}

func NewMediaManager(address string, createPath ...bool) MediaManager {
	create := false
	if len(createPath) > 0 {
		create = createPath[0]
	}
	media := &mediaManager{}

	if len(address) > 0 && address[0] == '/' {
		media.address = media.TrimEnd(address)
	} else if len(address) > 1 && (address[0] != '/' || address[0:2] == "./") {
		media.address = media.JoinPaths(getPwd(), address)
	} else {
		panic(fmt.Sprintf("couldn't connect %s with %s", getPwd(), address))
	}

	media.address = media.TrimEnd(media.address)

	if create {
		media.CreateFolder()
	} else {
		media.AddressIsValid()
	}
	return media
}
