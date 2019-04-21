package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func initOut(logDir, commandName string) (stdout, stderr io.Writer, err error) {
	if logDir == "" {
		stdout = os.Stdout
		stderr = os.Stderr
	} else {
		ts := time.Now().Unix()

		stdoutFileName := fmt.Sprintf("%s-%v-stdout.log", getFileNameWithoutExt(commandName), ts)
		stdoutFile, err := os.Create(filepath.Join(logDir, stdoutFileName))
		if err != nil {
			return nil, nil, err
		}
		stdout = io.MultiWriter(os.Stdout, stdoutFile)

		stderrFileName := fmt.Sprintf("%s-%v-stderr.log", getFileNameWithoutExt(commandName), ts)
		stderrFile, err := os.Create(filepath.Join(logDir, stderrFileName))
		if err != nil {
			return nil, nil, err
		}
		stderr = io.MultiWriter(os.Stderr, stderrFile)
	}
	return
}
