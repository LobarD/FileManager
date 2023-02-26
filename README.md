# FileManager

This is a Go program that provides a basic command-line interface for file and directory operations such as creating files and directories, copying files and directories, moving files and directories, renaming files and directories, and deleting files and directories.

The createFile function takes a filename as input and creates a new file with that name if it does not already exist. If the file or directory already exists, it prints an error message.

The copyDir function takes the name of the directory to be copied with its contents and the new name of the directory, this function creates a folder with the specified name and copies all the files and folders inside the source directory

The copyFile function copies a file with its contents to a new file with a new name.

The copyDir function takes two directory paths as input, the source directory and the destination directory, and recursively copies all files and directories in the source directory to the destination directory.

The move function takes two file or directory paths as input, the source and the destination, and moves the source file or directory to the destination directory. If the destination directory does not exist, it is created.

The rename function takes two file or directory names as input, the old name and the new name, and renames the file or directory with the old name to the new name.

The delete function takes a file or directory path as input and deletes the file or directory if it exists.

The list function lists the files and directories in the current directory.

The main function provides a command-line interface to the user to execute the functions listed above.