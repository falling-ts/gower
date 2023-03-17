package console

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func init() {
	cliApp.Commands = append(cliApp.Commands, &cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "初始化应用内容",
		UsageText: `
key 初始化秘钥
`,
		Action: func(c *cli.Context) error {
			argsNum := c.Args().Len()
			for i := 0; i < argsNum; i++ {
				return execInit(c.Args().Get(i))
			}
			return nil
		},
	})
}

func execInit(arg string) error {
	switch arg {
	case "key":
		key, err := generateSecretKey(16)
		if err != nil {
			return err
		}
		err = updateEnvVariable(".env", "APP_KEY", key)
		if err != nil {
			return err
		}

		fmt.Println("秘钥生成成功.")
	}

	return nil
}

func generateSecretKey(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

func updateEnvVariable(filePath, key, newValue string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = fmt.Sprintf("%s=%s", key, newValue)
		}
		lines = append(lines, line)
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(outputFile *os.File) {
		_ = outputFile.Close()
	}(outputFile)

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
