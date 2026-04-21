package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DegsRed72/gator/internal/config"
	"github.com/DegsRed72/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func main() {
	cfg := config.Read()
	st := state{config: &cfg}
	cmds := commands{list: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	dbURL := st.config.DBUrl
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening psql")
	}
	dbQueries := database.New(db)
	st.db = dbQueries

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments")
	}
	cmdArgs := []string{}
	if len(args) > 2 {
		cmdArgs = append(cmdArgs, args[2])
	}
	cmd := command{name: args[1], args: cmdArgs}
	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
