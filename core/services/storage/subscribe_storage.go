package storage

import (
	"os"
	"fmt"

	"github.com/goccy/go-json"
)

type Subscriber struct {
	Id int64		`json:"id"`
	Email string	`json:"email"`
}

type SubscribeStorage interface {
	New(*os.File) *SubscribeStorageService
	Save(Subscriber) error
	FindAll() (*[]Subscriber, error)
	FindByEmail(string) (*Subscriber, error)
}

type SubscribeStorageService struct {
	subsFile *os.File
	subs []Subscriber
}

func New(file *os.File) *SubscribeStorageService {
	return &SubscribeStorageService{
		subsFile: file,
	}
}

func (storageService *SubscribeStorageService) Save(sub Subscriber) error {
	var lastId int64
	
	if err := storageService.readFromFile(); err != nil {
		return err
	}

	if len(storageService.subs) < 1 {
		lastId = 0
	} else {
		lastId = storageService.subs[len(storageService.subs)-1].Id
	}

	lastId++
	sub.Id = lastId
	storageService.subs = append(storageService.subs, sub)

	return storageService.writeToFile()
}

func (storageService *SubscribeStorageService) FindAll() (*[]Subscriber, error) {
	err := storageService.readFromFile()
	
	if err != nil {
		return nil, err
	}

	return &storageService.subs, err
}


func (storageService *SubscribeStorageService) FindByEmail(email string) (*Subscriber, error) {
	err := storageService.readFromFile()
	
	if err != nil {
		return nil, err
	}

	for _, v := range storageService.subs {
		if v.Email == email {
			return &v, nil
		}
	}

	return nil, nil
}

func (storageService *SubscribeStorageService) readFromFile() error {
	byteSubs, err := os.ReadFile(storageService.subsFile.Name())
	
	if err != nil {
		return err
	}
	if len(byteSubs) < 1 {
		return nil
	}

	err = json.Unmarshal(byteSubs, &storageService.subs)

	if err != nil {
		fmt.Println(err)
		return err
	}
	
	return nil
}

func (storageService *SubscribeStorageService) writeToFile() error {
	byteSubs, err := json.Marshal(storageService.subs)

	if err != nil {
		return err
	}

	err = os.WriteFile(storageService.subsFile.Name(), byteSubs, os.FileMode(0775))

	if err != nil {
		return err
	}

	return nil
}