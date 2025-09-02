package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"mian/models"
	"os"
)

type Queue struct {
	items    []models.User
	filename string
}

func NewQueue(filename string) *Queue {
	queue := &Queue{items: make([]models.User, 0), filename: filename}
	err := queue.setup()
	if err != nil {
		fmt.Print("Error creating queue")
	}
	return queue
}

func (queue *Queue) setup() error {
	file, err := os.Open(queue.filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(queue.filename)
			if err != nil {
				return fmt.Errorf("error creating file: %s", err)
			}
			emptyData := []models.User{}
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(emptyData); err != nil {
				return fmt.Errorf("error initializing file: %v", err)
			}
			queue.items = emptyData
			return nil
		}
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %v", err)
	}

	if info.Size() == 0 {
		queue.items = []models.User{}
		return nil
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&queue.items); err != nil {
		return fmt.Errorf("error decoding file: %v", err)
	}

	return nil
}

func (queue *Queue) save() error {

	file, err := os.Create(queue.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(queue.items)
}

func (queue *Queue) Enqueue(user models.User) error {
	if cheakID(queue, user.ID) == true {
		return errors.New("user already exists")
	}
	queue.items = append(queue.items, user)
	return queue.save()
}

func (queue *Queue) Dequeue() (*models.User, error) {
	if len(queue.items) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}
	user := queue.items[0]
	queue.items = queue.items[1:]
	return &user, nil
}

func (queue *Queue) SearchID(id int) (*models.User, error) {
	for _, user := range queue.items {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf(`user not found`)
}

func (queue *Queue) SearchEmail(email string) (*models.User, error) {
	for _, user := range queue.items {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf(`user not found`)
}

func (queue *Queue) DeleteID(id int) error {
	for i, user := range queue.items {
		if user.ID == id {
			queue.items = append(queue.items[:i], queue.items[i+1:]...)
			return queue.save()
		}
	}
	return fmt.Errorf(`user not found`)
}

func (queue *Queue) DeleteEmail(email string) error {
	for i, user := range queue.items {
		if user.Email == email {
			queue.items = append(queue.items[:i], queue.items[i+1:]...)
			return queue.save()
		}
	}
	return fmt.Errorf(`user not found`)
}

func (queue *Queue) Update(id int, user models.User) error {
	for i, j := range queue.items {
		if j.ID == id {
			queue.items[i] = user
			return queue.save()
		}
	}
	return errors.New("user not found")
}

func (queue *Queue) All() []models.User {
	return queue.items
}

func (queue *Queue) PrintAll() {
	fmt.Print("-----------------List users------------------\n")
	for _, user := range queue.items {
		fmt.Printf("User ID: %d\n", user.ID)
		fmt.Printf("User Name: %s\n", user.Name)
		fmt.Printf("User Email: %s\n", user.Email)
		fmt.Printf("User Position: %s\n", user.Position)
		fmt.Printf("_____________________________________________\n")
	}
}

func (queue *Queue) Len() int {
	return len(queue.items)
}

func cheakID(queue *Queue, id int) bool {
	for _, user := range queue.items {
		if id == user.ID {
			return true
		}
	}
	return false
}
