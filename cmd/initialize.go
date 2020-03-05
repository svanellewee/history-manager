package cmd

import (
	_ "github.com/mattn/go-sqlite3"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initializeCmd)
}

var initializeCmd = &cobra.Command{
	Use: "initialize",
	Short: "Initializing the database	",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		// //fmt.Println("Hiya")

		// db, err := sql.Open("sqlite3", AbsolutePathDB)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// defer db.Close()

		// sqlStmt := `
		// CREATE TABLE IF NOT EXISTS entry (
		// 	entry_id INTEGER,
		// 	time TIMESTAMP,
		// 	VALUE VARCHAR
		// );
		// `
		// _, err = db.Exec(sqlStmt)
		// if err != nil {
		// 	log.Printf("%q: %s\n", err, sqlStmt)
		// 	return
		// }
	},
}
