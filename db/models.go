package main

import (
	"time"
	_ "database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)


// http://jinzhu.me/gorm/models.html#model-definition

type ModelID struct {
  ID        uint `gorm:"primary_key"`
}

type ModelTime struct {
  // UserID	uint
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}

type User struct {
	ModelID

	ModelTime
	Job				[]Job
	Tool				[]Tool
	UserSetting		[]UserSetting
}

type Job struct {
	ModelID
	Name        string  	`gorm:"not null"` // Default size for string is 255, reset it with this tag
	Command			string		`gorm:"unique_index:idx_command_arg;not null"`
	Arguments		string		`gorm:"unique_index:idx_command_arg"`
	ModelTime
	Schedule		[]Schedule
}

// Schedule The scheduling and history of all jobs ran.
type Schedule struct {
	ModelID
	JobID				uint			`gorm:"index"`
	StartTime		time.Time
	EndTime			time.Time
	Status			string		`gorm:"index"`
	PID					uint
	Completed		bool
	ReturnCode  int
	Output			string
	ModelTime
}

type Tool struct {
	ModelID
	Name				string
	Position		uint
	ModelTime
	ToolSetting 		[]ToolSetting
	Dropdown 		[]Dropdown
}

type UserSetting struct {
	ModelID
	UserID			uint		`gorm:"index"`
	Key					string		`gorm:"unique_index:idx_settings_key_value"`
	Value				string		`gorm:"unique_index:idx_settings_key_value"`
	ModelTime
}

type ToolSetting struct {
	ModelID
	ToolID			uint		`gorm:"index"`
	Key					string		`gorm:"unique_index:idx_settings_key_value"`
	Value				string		`gorm:"unique_index:idx_settings_key_value"`
	ModelTime
}

type Dropdown struct {
	ModelID
	ToolID			uint			`gorm:"index"`
	Name				string		`gorm:"index"`
	Position		uint			`gorm:"index"`
	Value				string
	ModelTime
}


func main() {
  db, _ := gorm.Open("sqlite3", "gorm.db")
  defer db.Close()

	db.AutoMigrate(&Job{}, &Schedule{})

	job := Job{Name: "Test1"}
	db.Create(&job)

	jobs := []Job{}
	db.First(&job)
	db.Find(&jobs)
	fmt.Println(&job)
	for _, j := range jobs {
		fmt.Println(j)
	}
}