package environment

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"

	"log"
	"os"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Environment interface {
	GetHash() string
	GetLookup() *map[string]string
	String() string
}

type EnvArray struct {
	variables            []string
	stringRepresentation string
	hash                 string
	lookup               *map[string]string
}

func NewEnvArray(input ...string) *EnvArray {
	lookup := make(map[string]string)
	var variables []string
	if len(input) == 0 {
		variables = os.Environ()
	} else {
		copy(variables, input)
	}
	sort.Strings(variables)
	return &EnvArray{
		hash:                 "",
		variables:            variables,
		stringRepresentation: strings.Join(variables, "\n"),
		lookup:               &lookup,
	}
}

func (env *EnvArray) String() string {
	return env.stringRepresentation
}

func (env *EnvArray) GetLookup() *map[string]string {
	populate(env)
	return env.lookup
}

func (env *EnvArray) GetHash() string {
	populate(env)
	return env.hash
}

func populate(env *EnvArray) {
	if len(*env.lookup) != 0 && env.hash != "" && env.stringRepresentation != "" {
		return
	}
	results := env.variables
	hash := sha256.New()
	for _, value := range results {
		_, err := hash.Write([]byte(value))
		if err != nil {
			log.Fatal("should not happen", err)
		}
		parts := strings.Split(value, "=")
		(*env.lookup)[parts[0]] = parts[1]
	}
	(*env).hash = hex.EncodeToString(hash.Sum(nil))
}

func InitEnv(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS environment (
		   environment_id INTEGER PRIMARY KEY AUTOINCREMENT,
		   hash VARCHAR UNIQUE,
		   environment VARCHAR
	);
	`)
	if err != nil {
		log.Fatal("Also sux", err)
	}
}

func Insert(db *sql.DB, env Environment) (int64, error) {
	tx, err := db.Begin()
	defer tx.Commit()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO environment(hash, environment) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec(env.GetHash(), env.String())
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if lastID == 0 {
		rows, err := tx.Query("SELECT environment_id FROM environment WHERE hash = ?", env.GetHash())
		if err != nil {
			return 0, err
		}
		for rows.Next() {
			err = rows.Scan(&lastID)
			if err != nil {
				return 0, err
			}
		}
	}
	return lastID, nil
}

// datastore key values. Same as sqlite? use bbolt?
// func main() {
// 	fmt.Println("Hiya")
// 	db, err := sql.Open("sqlite3", "./db")
// 	if err != nil {
// 		log.Fatal("Sux", err)
// 	}
// 	defer db.Close()

// 	initDB(db)
// 	//	envHash()
// 	s := NewEnvArray()
// 	fmt.Println("...", s.GetHash())
// 	for k, i := range *s.GetLookup() {
// 		fmt.Println(k, "...", i)
// 	}

// 	fmt.Printf("...%s \n", s)
// 	insertValue(db, s)
// 	//fmt.Println("....", insertEnv(123))
// 	// h := sha256.New()
// 	// if _, err := io.Copy(h, f); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Printf("%x", h.Sum(nil))
// }
