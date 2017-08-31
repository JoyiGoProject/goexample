package utils

import (
	"bufio"
	"os"
	"strings"
)

func Readfile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}

	defer file.Close()

	var sqls []string
	var tmpline string
	scanner := bufio.NewScanner(file)
	commentBegin := false
	//commentEnd := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "*/") {
			commentBegin = false
			continue
		}
		if len(line) > 0 {
			if strings.HasPrefix(line, "//") {
				continue
			} else if strings.HasPrefix(line, "--") {
				continue
			} else if strings.HasPrefix(line, "/*") {
				if strings.HasSuffix(line, "*/") {
					continue
				}
				commentBegin = true
				continue
			} else {
				if commentBegin {
					continue
				} else {
					// TODO
					if strings.HasSuffix(line, ";") {
						sqls = append(sqls, tmpline+line)
						tmpline = ""
					} else {
						tmpline = tmpline + line
					}
				}
			}
		}
	}
	return sqls
}
