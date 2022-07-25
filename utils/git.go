package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extractRegionAndMsg(str string) (string, string) {
	if strings.Contains(str, ":") {
		colonSplit := strings.Split(str, ":")
		region := strings.TrimSpace(colonSplit[0])
		msg := strings.TrimSpace(colonSplit[1])

		return region, msg
	}

	return "", str
}

func BuildPrefixWithMsg(prefix string, msg string) string {
	region, msg := extractRegionAndMsg(msg)

	if region != "" {
		prefix += fmt.Sprintf("(%s)", region)
	}

	delimiter := ":"
	return fmt.Sprintf("%s%s %s", prefix, delimiter, msg)
}

func AddToCommitMsg(text string, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	firstLiner, err := reader.ReadSlice('\n')

	// seek to next line
	if _, err := file.Seek(int64(len(firstLiner)), 0); err != nil {
		return err
	}

	body := make([]byte, 0)

	if _, err := file.Read(body); err != nil {
		return err
	}

	// Truncate the file
	_ = file.Truncate(0)
	_, _ = file.Seek(0, 0)

	if _, err = file.WriteString(text); err != nil {
		return err
	}

	if _, err = file.Write(body); err != nil {
		return err
	}

	if err = file.Sync(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	return nil
}

func GetCommitMsgFromFile(fileName string) (string, error) {
	// the first line of the file should be the commit msg

	file, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return line, nil

}
