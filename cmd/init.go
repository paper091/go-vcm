/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		repoPath := ".gvc"
		dirs := []string{
			filepath.Join(repoPath, "objects"),
			filepath.Join(repoPath, "refs"),
		}

		if _, err := os.Stat(repoPath); err == nil{
			fmt.Println("Repository already initialised.")
			return
		}

		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil{
				fmt.Printf("Failed to create directory %s: %v\n", dir , err)
				return
			}
		}

		// creating head fiel pointing to master
		headPath := filepath.Join(repoPath, "HEAD")
		err := os.WriteFile(headPath, []byte("ref: refs/heads/master\n"), 0644)
		if err != nil{
			fmt.Printf("Failed to crea HEAD fule: %v\n", err)
		}

		fmt.Println("Initialized empty GVC repository in .gvc/")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
