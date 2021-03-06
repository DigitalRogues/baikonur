// Copyright © 2017 Martin Kim Dung-Pham <kim@elbedev.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	log "github.com/apex/log"
	"github.com/q231950/baikonur/cityparser"
	"github.com/spf13/cobra"
)

var citiesFolderPath string

// citiesCmd represents the cities command
var citiesCmd = &cobra.Command{
	Use:   "cities",
	Short: "Import cities",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Attempting to import cities...")
		if len(citiesFolderPath) == 0 {
			log.Error("`path` to cities csv file is missing. Please provide a path `--path|-p <./path/to/cities.csv>`")
		} else {
			log.WithFields(log.Fields{"path": citiesFolderPath}).Info("Path to cities csv")
			processor := new(cityparser.CityParser)
			reader, error := os.Open(citiesFolderPath)
			defer reader.Close()
			if error != nil {
				log.WithError(error).Error("Failed to open csv.")
				return
			}
			processor.Parse(reader)
		}
	},
}

func init() {
	citiesCmd.Flags().StringVarP(&citiesFolderPath, "path", "p", "", "The path to the cities csv.")

	importCmd.AddCommand(citiesCmd)
}
