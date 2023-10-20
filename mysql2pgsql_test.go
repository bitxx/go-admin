package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	bytes, err := os.ReadFile("./app.sql")
	require.Nil(t, err)
	content := string(bytes)
	content = strings.Replace(content, "`", "\"", -1)
	content = strings.Replace(content, "SET NAMES utf8mb4;", "", -1)
	content = strings.Replace(content, "SET FOREIGN_KEY_CHECKS = 0;", "", -1)
	content = strings.Replace(content, "ENGINE=InnoDB", "", -1)
	content = strings.Replace(content, " DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin", "", -1)
	content = strings.Replace(content, " CHARACTER SET utf8mb4 COLLATE utf8mb4_bin", "", -1)
	content = strings.Replace(content, "datetime", "timestamp", -1)
	re := regexp.MustCompile("\\/\\*[\\s\\S]*?\\*\\/")
	content = re.ReplaceAllString(content, "")
	_ = os.WriteFile("app_pgsql.sql", []byte(content), 0666)
}
