package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
)

// TODO:
//  add option to use custom directory as base for shell instead of .
//    add option to make output directory if it doesn't exist
//  add option to follow symlinks?
//  add option to specify data file
//  add ability to clone from git and template from that

// TODO: change to something more "clever" like "out" (eg., stamp out someDirectory)
var (
	genCommand = &cobra.Command{
		Use: "gen [path to template]",
		Short: "Generate shell from template",
		Args: cobra.ExactArgs(1),
		RunE: gen,
	}
	outputDirectory string
)

func init() {
	genCommand.Flags().StringVarP(&outputDirectory, "output", "o", ".",
		"Path to output directory for generated shell")
}

func gen(_ *cobra.Command, args []string) error {
	pathToTemplate := args[0]
	fileInfo, err := os.Stat(pathToTemplate)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("'%v': no such file or directory", pathToTemplate)
		}
		return err
	}

	if ! fileInfo.Mode().IsRegular() && ! fileInfo.Mode().IsDir() {
		return fmt.Errorf("template path '%v' is not a directory or regular file", pathToTemplate)
	}

	// TODO: verify pathToTemplate is not cwd

	return template(pathToTemplate, outputDirectory)

	/*
	if directory:
	  - template name
	  - mkdir (if not already exist)
	  - add new entry to relative path stack for generated files
	  - loop over each entry in directory calling function
	if reg file:
	  - template name
	  - template file contents
	  - write to new location
    else
	  - emit warning
	  - continue processing
	 */

	//return filepath.WalkDir(pathToTemplate, callback)
}

func template(inputFile string, outputPath string) error {
	fileInfo, err := os.Lstat(inputFile)
	if err != nil {
		return err
	}

	fmt.Printf("inputFile: %v, outputPath: %v -- ", inputFile, outputPath)

	if fileInfo.Mode().IsDir() {
		fmt.Printf("Directory: %v\n", fileInfo.Name())
		dirEntries, err := os.ReadDir(inputFile)
		if err != nil {
			return err
		}

		// TODO: template name

		newOutputPath := filepath.Join(outputPath, fileInfo.Name())

		err = os.Mkdir(newOutputPath, fileInfo.Mode())
		if err != nil {
			return err
		}

		for _, entry := range dirEntries {
			newInputPath := filepath.Join(inputFile, entry.Name())
			err = template(newInputPath, newOutputPath)
			if err != nil {
				return err
			}
		}
	} else if fileInfo.Mode().IsRegular() {
		fmt.Printf("File: %v\n", fileInfo.Name())
	} else {
		// TODO: look into supporting symlinks
		_, _ = fmt.Fprintf(os.Stderr, "Skipping %v as it is not a regular file or directory.", fileInfo.Name())
	}

	return nil
}

func callback(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return fmt.Errorf("fatal error walking template: %v", err)
	}

	abs, _ := filepath.Abs(d.Name())
	base := filepath.Base(abs)

	// Looking at .idea - .idea - /home/albersm/Projects/go/stamp/.idea - .idea

	fmt.Printf("Looking at %v - %v - %v - %v\n", d.Name(), path, abs, base)
	return nil
}
