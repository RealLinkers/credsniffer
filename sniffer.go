package main

import (
    "fmt"
    "log"
	"flag"
	"path/filepath"
	"bufio"
	"os"
	"strings"
	"strconv"
	"runtime"
)

type Color string

type file struct {
	fileName string
	filesData   []filedata
}

type filedata struct {
	line     int
	text  string
}

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

func scanFile(path string, phrases []string) (file, error) {
	result := file{}
	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	line := 1
	filedatalist := []filedata{}

	for scanner.Scan() {
		for _, phrase := range phrases {
			if strings.Contains(scanner.Text(), phrase) {
				value := filedata{line, scanner.Text()}
				filedatalist = append(filedatalist, value)
			}
			line++
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	result = file{path, filedatalist}
	return result, nil
}

func main() {
	phrasePtr := flag.String("pattern", "password,Password,AsSecureString", "Specify patterns to scan separated by comma")
	dirPtr := flag.String("dir", ".", "Specify a directory to scan Example: \\\\localhost\\c$\\Users\\Public")
	extPtr := flag.String("ext", "ps1", "Specify an extension to scan Example: ps1,txt")
	flag.Parse()
	
	extensions := strings.Split(*extPtr, ",")
	phrases := strings.Split(*phrasePtr, ",")
    err := filepath.Walk(*dirPtr, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		fileExtension := filepath.Ext(path)
		extCheck := false
		for _, ext := range extensions {
			if fileExtension == "." + ext {
				extCheck = true
			}
		}

		if !extCheck {
			return nil
		}

		result, err := scanFile(path, phrases)
		if err != nil {
			return err
		}

		for _, val := range result.filesData {
			left := Green + result.fileName + ":" +  strconv.Itoa(val.line) + Reset
			right := Red + val.text + Reset
			final := left + "  " + right
			fmt.Println(final)
			fmt.Println("---------------------------")
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}
}

func init() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}
}
