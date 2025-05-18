package cmd

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use: "add [file]",
	Short: "Add a file to the staging area",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		filePath := args[0]

		// readign the file contents
		content, err:= os.ReadFile(filePath)
		if err != nil{
			fmt.Printf("failed to read file: %v\n", err)
			return
		}

		// computing hash
		hash := sha1.Sum(content)
		hashStr := fmt.Sprintf("%x", hash)


		// saving to that dir
		// .gvc/objects/<hash>
		objectPath := filepath.Join(".gvc", "objects", hashStr)
		err = os.WriteFile(objectPath, content, 0644)
		if err != nil{
			fmt.Printf("failed to write object: %v\n", err)
			return
		}

		fmt.Printf("Added %s as object %s\n", filePath, hashStr)
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}