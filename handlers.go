package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/DegsRed72/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects only one argument")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatal("User not in db")
	}
	s.config.SetUser(cmd.args[0])
	fmt.Println("Username has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("register expects only one argument")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		log.Fatal("Name already in db")
	}
	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return errors.New("error creating user")
	}
	s.config.SetUser(dbUser.Name)
	fmt.Println("user has been created")
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New(err.Error())
	}
	for _, user := range dbUsers {
		if user.Name == s.config.CurrentUserName {
			fmt.Printf("%s (current)", user.Name)
		} else {
			fmt.Println(user.Name)
		}

	}

	return nil
}
