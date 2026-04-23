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

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		log.Fatal("addfeed expects exactly two arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	dbUser, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return errors.New("User not found")
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    dbUser.ID,
	})
	fmt.Printf("%v", feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	return nil
}
