package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func createFile(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) == false {
		fmt.Printf("Error: file or directory '%s' already exists\n", filename)
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	fmt.Printf("File '%s' created.\n", filename)
}

func createDir(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) == false {
		fmt.Printf("Error: directory or file '%s' already exists\n", directory)
		return
	}

	err := os.MkdirAll(directory, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Directory '%s' created.\n", directory)
}

func copyFile(src string, dst string) error {
	// Open source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create an output file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents of the source file to the output file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Copy the permissions of the source file to the output file
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

func copyDir(src string, dst string) error {
	// Create an output folder if it doesn't exist
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}

	// Get a list of files in the source folder
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each file to the output folder
	for _, file := range files {
		srcPath := filepath.Join(src, file.Name())
		dstPath := filepath.Join(dst, file.Name())

		// Copy file
		if file.IsDir() {
			err = copyDir(srcPath, dstPath)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copyFile(srcPath, dstPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return nil
}

func move(source string, destination string) {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("File or directory not found.")
		return
	}

	if sourceInfo.IsDir() {
		err = os.Rename(source, destination)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Moved directory '%s' to '%s'\n", source, destination)
	} else {
		fileName := filepath.Base(source)
		destFileName := filepath.Join(destination, fileName)

		err = os.Rename(source, destFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Moved file '%s' to '%s'\n", source, destination)
	}
}

func rename(oldName string, newName string) {
	if _, err := os.Stat(newName); os.IsNotExist(err) == false {
		fmt.Printf("Error: '%s' already exists\n", newName)
		return
	}

	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Item '%s' renamed to '%s'.\n", oldName, newName)
}

func delete(item string) {
	itemInfo, err := os.Stat(item)
	if err != nil {
		fmt.Printf("Item '%s' does not exist.\n", item)
		return
	}

	if itemInfo.IsDir() {
		err = os.RemoveAll(item)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Directory '%s' deleted.\n", item)
	} else {
		err = os.Remove(item)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("File '%s' deleted.\n", item)
	}
}

func list() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Current directory:", dir)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Contents:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	list()

	for {
		fmt.Print("\nEnter command (createFile, createDir, copyFile, copyDir, move, rename, delete, list, cd, back, exit): ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		switch command {
		case "createFile":
			fmt.Print("Enter filename: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSuffix(filename, "\n")
			createFile(filename)

		case "createDir":
			fmt.Print("Enter directory name: ")
			directory, _ := reader.ReadString('\n')
			directory = strings.TrimSuffix(directory, "\n")
			createDir(directory)

		case "copyFile":
			fmt.Print("Enter source path: ")
			source, _ := reader.ReadString('\n')
			source = strings.TrimSuffix(source, "\n")
			fmt.Print("Enter destination path: ")
			destination, _ := reader.ReadString('\n')
			destination = strings.TrimSuffix(destination, "\n")

			err := copyFile(source, destination)
			if err != nil {
				fmt.Println(err)
			}

		case "copyDir":
			fmt.Print("Enter source path: ")
			source, _ := reader.ReadString('\n')
			source = strings.TrimSuffix(source, "\n")
			fmt.Print("Enter destination path: ")
			destination, _ := reader.ReadString('\n')
			destination = strings.TrimSuffix(destination, "\n")

			err := copyDir(source, destination)
			if err != nil {
				fmt.Println(err)
			}

		case "move":
			fmt.Print("Enter source path: ")
			source, _ := reader.ReadString('\n')
			source = strings.TrimSuffix(source, "\n")
			fmt.Print("Enter destination path: ")
			destination, _ := reader.ReadString('\n')
			destination = strings.TrimSuffix(destination, "\n")
			move(source, destination)

		case "rename":
			fmt.Print("Enter current item name: ")
			oldName, _ := reader.ReadString('\n')
			oldName = strings.TrimSuffix(oldName, "\n")
			fmt.Print("Enter new item name: ")
			newName, _ := reader.ReadString('\n')
			newName = strings.TrimSuffix(newName, "\n")
			rename(oldName, newName)

		case "delete":
			fmt.Print("Enter item name: ")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSuffix(item, "\n")
			delete(item)

		case "list":
			list()
		case "cd":
			fmt.Print("Enter directory name: ")
			directory, _ := reader.ReadString('\n')
			directory = strings.TrimSuffix(directory, "\n")
			err := os.Chdir(directory)
			if err != nil {
				fmt.Print("Directory not found. Do you want to create it? (y/n): ")
				create, _ := reader.ReadString('\n')
				create = strings.TrimSuffix(create, "\n")
				if create == "y" {
					createDir(directory)
					os.Chdir(directory)
				} else {
					fmt.Println("Navigation cancelled.")
				}
			}
			list()

		case "back":
			os.Chdir("..")
			list()

		case "exit":
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}
