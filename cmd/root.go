/*
Copyright © 2022 nitin mewar <nitinmewar28@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"github.com/common-nighthawk/go-figure"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert file types inside CLI",
	Long:  `convert can (As name suggest) convert one file type to another. It has support for many file extension that you can use inside your cli. no need to go to any webpage or anything.`,

	Run: func(cmd *cobra.Command, args []string) {
		figure.NewFigure("Convert", "isometric1", true).Print()
		err := godotenv.Load(".env")
		pwd, _ := os.Getwd()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
		config.Default.Secret = os.Getenv("CONVERTAPI_SECRET")

		arg1, arg2, arg3 := args[0], args[1], args[2]
		fmt.Println("converting " + arg1 + " to " + arg2 + " inside " + pwd)

		convertedFile := convertapi.ConvDef(arg1, arg2,
			param.NewPath("file", arg3, nil))

		if files, errs := convertedFile.ToPath(pwd); errs == nil {
			fmt.Println("converted file saved to: ", files[0].Name())
		} else {
			fmt.Println(errs)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
