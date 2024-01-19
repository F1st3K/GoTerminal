package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Определяем флаги
	dFlag := flag.String("d", "", "Указание каталога для извлечения файлов")
	oFlag := flag.Bool("o", false, "Перезапись существующих файлов")
	qFlag := flag.Bool("q", false, "Тихий режим без вывода")

	flag.Parse()
	zipFile := flag.Arg(0)

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	// Перебираем файлы в архиве для извлечения
	for _, f := range r.File {
		targetDir := *dFlag
		if targetDir == "" {
			targetDir, _ = os.Getwd()
		}
		if !filepath.IsAbs(targetDir) {
			targetDir, _ = filepath.Abs(targetDir)
		}

		targetPath := filepath.Join(targetDir, f.Name)

		if !*qFlag {
			fmt.Println("unzip:", f.Name)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(targetPath, os.ModePerm)
			continue
		}

		if !*oFlag {
			_, err := os.Stat(targetPath)
			if err == nil {
				var response string
				fmt.Printf("%s is alredy exist. Overwrite? (yes/no/all): ", targetPath)
				_, err := fmt.Scanln(&response)
				if err != nil {
					fmt.Println(err)
					return
				}
				response = strings.ToLower(strings.TrimSpace(response))

				if response != "yes" && response != "y" && response != "all" {
					continue
				} else if response == "all" {
					*oFlag = true
				}
			}
		}

		// Копируем содержимое файла в целевой файл
		fw, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			fmt.Println(err)
			return
		}

		fr, err := f.Open()
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = io.Copy(fw, fr)
		if err != nil {
			fmt.Println(err)
			return
		}

		fw.Close()
		fr.Close()
	}
}
