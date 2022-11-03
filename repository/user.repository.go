package repository

import (
	"context"
	"elasticsearch/fiber-elasticsearch/entity"
	"elasticsearch/fiber-elasticsearch/interfaces"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

type userRepository struct {
	Es *elasticsearch.Client
}

func NewUserRepository(es *elasticsearch.Client) interfaces.UserRepository {
	return &userRepository{
		Es: es,
	}
}

func (u *userRepository) CreateUser(user *entity.User) (entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := uuid.New().String()
	jsondata, err := json.Marshal(user)
	if err != nil {
		return *user, err
	}

	req := esapi.IndexRequest{
		Index:      "users",
		DocumentID: id,
		Body:       strings.NewReader(string(jsondata)),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, u.Es)
	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()

	return *user, nil
}

func (u *userRepository) FindAllUser() (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := esapi.SearchRequest{
		Index: []string{"users"},
	}

	res, err := req.Do(ctx, u.Es)
	if err != nil {
		log.Panic(err)
	}

	defer res.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Panic(err)
	}

	return result, nil
}

func (u *userRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := esapi.DeleteRequest{
		Index:      "users",
		DocumentID: id,
	}

	res, err := req.Do(ctx, u.Es)
	if err != nil {
		log.Panic(err)
	}

	if res.IsError() {
		return fmt.Errorf("error deleting document")
	}

	defer res.Body.Close()

	return nil
}

func (u *userRepository) FindUserById(id string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := esapi.GetRequest{
		Index:      "users",
		DocumentID: id,
	}

	res, err := req.Do(ctx, u.Es)
	if err != nil {
		log.Panic(err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("error getting document")
	}

	defer res.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Panic(err)
	}

	return result, nil
}

func (u *userRepository) UpdateUser(id string, user *entity.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	channel := make(chan error)
	go func() {
		jsondata, err := json.Marshal(user)
		if err != nil {
			channel <- err
		}

		req := esapi.IndexRequest{
			Index:      "users",
			DocumentID: id,
			Body:       strings.NewReader(string(jsondata)),
			Refresh:    "true",
		}

		res, err := req.Do(ctx, u.Es)
		if err != nil {
			channel <- err
		}

		defer res.Body.Close()

		channel <- nil
	}()

	return <-channel
}

func (u *userRepository) SearchUser() (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// multi match query
	user := `{
		"query": {
			"multi_match": {
				"query": "basket",
				"fields": [
					"hobbies"
				]
			}
		}
	}`

	req := esapi.SearchRequest{
		Index: []string{"users"},
		Body:  strings.NewReader(user),
	}

	fmt.Println(req.Body)

	res, err := req.Do(ctx, u.Es)
	if err != nil {
		log.Panic(err)
	}

	if res.IsError() {
		return nil, fmt.Errorf("error searching document")
	}

	defer res.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Panic(err)
	}

	return result, nil
}
