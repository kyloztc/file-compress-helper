package cmd

import (
	"file-compression-helper/pkg"
	"github.com/spf13/cobra"
)

var path string
var limit float64
var outDir string
var subFolder string

var splitCmd = &cobra.Command{
	Use: "split",
	Run: func(cmd *cobra.Command, args []string) {
		_ = pkg.SplitFolder(path, limit, outDir, subFolder)
	},
}

func init() {
	splitCmd.Flags().StringVarP(&path, "path", "p", "", "folder path")
	splitCmd.Flags().Float64VarP(&limit, "limit", "l", 1024, "sub zip size limit")
	splitCmd.Flags().StringVarP(&outDir, "out", "o", "", "output path")
	splitCmd.Flags().StringVarP(&subFolder, "suffix", "s", "", "sub zup suffix name")
	rootCmd.AddCommand(splitCmd)
}
