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
	"os"
	"strings"

	"github.com/svanellewee/history-manager/environment"

	"github.com/spf13/cobra"
)

func insertEnv(entryID int64) string {
	s := os.Environ()
	//sort.Strings(s)
	stringData := make([]string, 0, 100)
	for _, i := range s {
		parts := strings.Split(i, "=")
		stringData = append(stringData, fmt.Sprintf("(%d, '%s', '%s')", entryID, parts[0], parts[1]))
	}
	return fmt.Sprintf("INSERT INTO environment(entry_id, key, value) VALUES %s;", strings.Join(stringData, ","))
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", AbsolutePathDB)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		environmentID, err := environment.Insert(db, environment.NewEnvArray())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("....", environmentID)
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("INSERT INTO entry(environment_id, history_id, time, value) VALUES (?, ?, DATETIME(), ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		//i := 0
		values := strings.Trim(args[0], " ")
		parts := strings.Split(values, " ")
		theRest := strings.Trim(strings.Join(parts[1:], " "), " ")
		res, err := stmt.Exec(environmentID, parts[0], theRest)
		if err != nil {
			log.Fatal(err)
		}
		_ = res
		// lastID, err := res.LastInsertId()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		tx.Commit()

		// envUpdate := insertEnv(lastID)
		// //fmt.Println(envUpdate)
		// _, err = db.Exec(envUpdate)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
