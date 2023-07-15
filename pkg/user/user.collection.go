package user

import (
	"errors"
	"fmt"
	"github.com/leonardoramosc/task-cli/pkg/database"
	"github.com/leonardoramosc/task-cli/pkg/task"
)

type UserSchema struct {
	*User   `json:"user"`
	Tasks []*task.Task `json:"tasks"`
}

func (us *UserSchema) ListTasks() []*task.Task {
	return us.Tasks
}

type userCollection struct {
	Name string        `json:"name"`
	Data []*UserSchema `json:"data"`
	database *database.Database
}

func (uc *userCollection) Init() *userCollection {
	if uc.database == nil {
		uc.database = &database.Database{}
	}
	uc.Name = "users"
	
	return uc
}

func (uc *userCollection) LoadData() {
	uc.database.LoadDB(uc)
}

func (uc *userCollection) Exec() {
	if uc.database == nil {
		panic("Failed to exec query since a database instance was not found")
	}
	uc.database.Update(uc)
}

func (uc *userCollection) CreateUser(u *User) (*UserSchema, error) {
	existingUser, _ := uc.GetByUsername(u.Username)
	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf("El username: %v ya esta tomado\n", u.Username))
	}
	
	us := UserSchema{ User: u, Tasks: []*task.Task{} }

	uc.Data = append(uc.Data, &us)

	return &us, nil
}

func (uc *userCollection) GetByUsername(username string) (*UserSchema, error) {
	for _, us := range uc.Data {
		if us.User.Username == username {
			return us, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("user %v does not exist", username))
}

func (uc *userCollection) AppendTask(username string, task *task.Task) error {
	u, e := uc.GetByUsername(username)
	if e != nil {
		return errors.New(fmt.Sprintf("Cannot add task to user %v because it doesn't exist\n", username))
	}
	u.Tasks = append(u.Tasks, task)
	return nil
}

var (
	instance userCollection
)

func GetUserCollection() *userCollection {
	if instance.Name == "" {
		instance = userCollection{}
		instance.Init().LoadData()
		fmt.Println("USER COLLECTION DATA LOADED")
	}

	return &instance
}
