/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		db, err := sql.Open("sqlite3", AbsolutePathDB)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query(`
		SELECT history_id, environment_id, entry_id, time, value, hash 
		FROM entry LEFT JOIN environment USING (environment_id)
		`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id, envID int
			var historyID string
			var name string
			var value string
			var hash string
			err = rows.Scan(&historyID, &envID, &id, &name, &value, &hash)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d, %d, %s, %s, %s\n", aurora.Magenta(id), aurora.BrightRed(envID), aurora.BrightRed(hash), aurora.Cyan(name), aurora.Yellow(value))
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
