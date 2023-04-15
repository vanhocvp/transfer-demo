package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/dambaquyen96/smartivr-backend-go/pkg/setting"
	"github.com/gabriel-vasile/mimetype"
	"github.com/go-audio/wav"

	"github.com/twinj/uuid"
)

// SaveReceivedFile ...
func SaveReceivedFile(file *multipart.FileHeader, desPath string) error {
	inStream, err := file.Open()
	if err != nil {
		return err
	}
	folderPath := strings.LastIndex(desPath, "/")
	log.Printf("[info] SaveReceivedFile | folder: %s", desPath[:folderPath])
	err = os.MkdirAll(desPath[:folderPath], os.ModePerm)
	if err != nil {
		return err
	}
	outStream, err := os.Create(desPath)
	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	for {
		n, err := inStream.Read(buf)
		if err != nil && err != io.EOF {
			break
		}
		if n == 0 {
			break
		}
		if _, err := outStream.Write(buf[:n]); err != nil {
			return err
		}
	}
	defer inStream.Close()
	defer outStream.Close()
	return nil
}

// GenReceivedFilePath ...
func GenReceivedFilePath(file *multipart.FileHeader, subFolder string) (string, error) {
	fileNameFrags := strings.Split(file.Filename, ".")
	fileExt := fileNameFrags[len(fileNameFrags)-1]
	// Check userID folder
	userFolder := fmt.Sprintf("%s/%s", setting.FileManagerSetting.AudioReceivedRoot, subFolder)
	err := os.Mkdir(userFolder, 0777)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	newFilePath := fmt.Sprintf(
		"%s/%s.%s", userFolder, uuid.NewV4().String(), fileExt)
	return newFilePath, nil
}

// GenReceivedFilePath ...
func GenBlacklistFilePath(file *multipart.FileHeader, subFolder string) (string, error) {
	fileNameFrags := strings.Split(file.Filename, ".")
	fileExt := fileNameFrags[len(fileNameFrags)-1]
	// Check userID folder
	userFolder := fmt.Sprintf("%s/%s", setting.FileManagerSetting.BlacklistFileRoot, subFolder)
	err := os.Mkdir(userFolder, 0777)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	newFilePath := fmt.Sprintf(
		"%s/%s.%s", userFolder, uuid.NewV4().String(), fileExt)
	return newFilePath, nil
}

// GenConvertedFilePath ...
func GenConvertedFilePath(file *multipart.FileHeader, subFolder string) (string, error) {
	// Check userID folder
	userFolder := fmt.Sprintf("%s/%s", setting.FileManagerSetting.AudioConvertedRoot, subFolder)
	err := os.Mkdir(userFolder, 0777)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	newFilePath := fmt.Sprintf(
		"%s/%s.%s", userFolder, uuid.NewV4().String(), "wav")
	return newFilePath, nil
}

// GenPlanDetailFilePath ...
func GenPlanDetailFilePath(file *multipart.FileHeader, subFolder string) (string, error) {
	userFolder := fmt.Sprintf("%s/%s", setting.FileManagerSetting.PlanDetailRoot, subFolder)
	err := os.Mkdir(userFolder, 0777)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	fileExtension, err := GetFileNameExtension(file.Filename)
	if err != nil {
		return "", err
	}
	newFilePath := fmt.Sprintf(
		"%s/%s.%s", userFolder, uuid.NewV4().String(), fileExtension)
	return newFilePath, nil
}

// DeleteFile ...
func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

// GetAudioDuration ...
func GetAudioDuration(filePath string) (float64, error) {
	inStream, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	duration, err := wav.NewDecoder(inStream).Duration()
	if err != nil {
		return 0, err
	}

	defer inStream.Close()
	return duration.Seconds(), nil
}

// GetFileNameWithoutExtension ...
func GetFileNameWithoutExtension(fileName string) string {
	fileNameFrags := strings.Split(fileName, ".")
	return strings.Join(fileNameFrags[:len(fileNameFrags)-1], ".")
}

// GetFileNameExtension ...
func GetFileNameExtension(fileName string) (string, error) {
	fileNameFrags := strings.Split(fileName, ".")
	if len(fileNameFrags) < 2 {
		return "", fmt.Errorf("file has no extension")
	}

	return fileNameFrags[len(fileNameFrags)-1], nil
}

func ValidateFileName(fileNameRaw string) bool {
	fileNameSplit := strings.Split(fileNameRaw, ".")
	fileType := fileNameSplit[len(fileNameSplit)-1]
	if fileType == "xlsx" || fileType == "xls" {
		return true
	}
	return false
}
func ValidateMagicHeaderBytesFile(fileNamePathSave string) bool {
	buf, _ := ioutil.ReadFile(fileNamePathSave)
	mtype := mimetype.Detect(buf)
	mimeType := mtype.String()
	extension := mtype.Extension()
	if (extension == ".xlsx" || extension == ".xls") && (mimeType == "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {
		return true
	}
	return false
}
