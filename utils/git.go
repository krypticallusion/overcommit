package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BuildPrefixWithMsg(prefix string, msg string) string {
	if strings.Contains(msg, ":") {
		colonSplit := strings.Split(msg, ":")
		region := strings.Trim(colonSplit[0], " ")
		msg = colonSplit[1]

		prefix += fmt.Sprintf("(%s)", region)
	}

	delimiter := ":"
	return fmt.Sprintf("%s%s %s", prefix, delimiter, msg)
}

func AddToCommitMsg(text string, filename string) error {
	body, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

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