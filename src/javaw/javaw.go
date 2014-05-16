package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "<configureFile> <encryptFile>")
		os.Exit(1)
	}
	confPath := os.Args[1]
	orgPath := os.Args[2]
	javaPath := orgPath + ".java"
	txtPath := orgPath + ".txt"

	var confItems map[string]string
	confItems = make(map[string]string)

	confFile, err := os.Open(confPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	confReader := bufio.NewReader(confFile)
	for {
		if line, err := confReader.ReadString('\n'); err != nil {
			break
		} else {
			items := strings.Split(line, "=")
			confItems[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
		}
	}
	confFile.Close()

	isSkip := false
	ind := strings.LastIndex(orgPath, ".")
	if ind != -1 {
		suffix := orgPath[ind+1:]
		if strings.Index(confItems["SkpiSuffix"], suffix) != -1 {
			isSkip = true
		}
	}
	if isSkip {
		fmt.Println("Skip", orgPath)
		os.Exit(1)
	}

	if err := os.Rename(orgPath, javaPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	infile, err := os.Open(javaPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outfile, err := os.OpenFile(txtPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err := io.Copy(outfile, infile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	outfile.Close()
	infile.Close()

	if err := os.Remove(javaPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	renameCmd := exec.Command(confItems["RenameCmd"], txtPath, orgPath)
	if err := renameCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

