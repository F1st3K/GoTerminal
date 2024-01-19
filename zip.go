package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Определяем флаги
	rFlag := flag.Bool("r", false, "Рекурсивное добавление файлов и каталогов")
	nineFlag := flag.Bool("9", false, "Использовать наивысший уровень сжатия")
	oFlag := flag.String("o", "output.zip", "Имя архива")

	flag.Parse()
	files := flag.Args()

	// Создаем новый архив
	newZipFile, err := os.Create(*oFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Добавляем файлы или директории в архив
	for _, file := range files {
		basePath := ""
		if *rFlag {
			fi, err := os.Stat(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			if fi.IsDir() {
				basePath = file
			}
		}

		err = addFilesToZip(file, basePath, zipWriter, *nineFlag)
		if err != nil {
			fmt.Printf("zip: %s: %v\n", file, err)
			return
		}
	}

	fmt.Println("zip: archive created:", *oFlag)
}

// Функция для рекурсивного добавления файлов
func addFilesToZip(file string, base string, zipWriter *zip.Writer, isNineZip bool) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	if info.IsDir() {
		files, err := os.ReadDir(file)
		if err != nil {
			return err
		}
		for _, f := range files {
			err = addFilesToZip(filepath.Join(file, f.Name()), base, zipWriter, isNineZip)
			if err != nil {
				return err
			}
		}
	} else {
		fileToZip, err := os.Open(file)
		if err != nil {
			return err
		}
		defer fileToZip.Close()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		if isNineZip {
			header.Method = zip.Deflate
		}

		if base != "" {
			rel, err := filepath.Rel(base, file)
			if err != nil {
				return err
			}
			header.Name = filepath.ToSlash(rel)
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			return err
		}
	}
	return nil
}
