package utils

import (
	"fmt"
	"os"
	"strings"
)

func BuildPrefixWithMsg(prefix string, msg string) string {
	if strings.Contains(msg, ":") {
		region := strings.Trim(strings.Split(msg, ":")[0], " ")

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
