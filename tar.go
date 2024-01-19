package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Объявление флагов для команды tar
	create := flag.Bool("c", false, "Создать новый архив")
	extract := flag.Bool("x", false, "Извлечь файлы из архива")
	archiveName := flag.String("f", "", "Имя файла архива")

	flag.Parse()

	// Создание нового архива
	if *create && *archiveName != "" {
		files := flag.Args()
		err := createArchive(files, *archiveName)
		if err != nil {
			fmt.Printf("tar: %s: %v\n", *archiveName, err)
		} else {
			fmt.Println("tar: arhive created:", *archiveName)
		}
	}

	// Извлечение файлов из архива
	if *extract && *archiveName != "" {
		err := extractArchive(*archiveName)
		if err != nil {
			fmt.Printf("tar: %s: %v\n", *archiveName, err)
		} else {
			fmt.Println("tar: arhive is dearhivate:", *archiveName)
		}
	}
}

// Создание нового архива
func createArchive(files []string, archiveName string) error {
	archiveFile, err := os.Create(archiveName)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	tw := tar.NewWriter(archiveFile)
	defer tw.Close()

	for _, file := range files {
		err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			header.Name = path

			if err := tw.WriteHeader(header); err != nil {
				return err
			}

			if !info.Mode().IsDir() {
				fileToArchive, err := os.Open(path)
				if err != nil {
					return err
				}
				defer fileToArchive.Close()

				if _, err := io.Copy(tw, fileToArchive); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Извлечение файлов из архива
func extractArchive(archiveName string) error {
	archiveFile, err := os.Open(archiveName)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	tr := tar.NewReader(archiveFile)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("tar: dearhivate... %s\n", header.Name)

		targetFilePath := filepath.Join(".", header.Name)
		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(targetFilePath, 0755); err != nil {
				return err
			}
			continue
		}

		// Создание файла и копирование данных
		fileToExtract, err := os.Create(targetFilePath)
		if err != nil {
			return err
		}
		defer fileToExtract.Close()

		if _, err := io.Copy(fileToExtract, tr); err != nil {
			return err
		}
	}
	return nil
}
