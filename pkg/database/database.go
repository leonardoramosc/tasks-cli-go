package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Collection struct {
	Name string
	Data []interface{}
}

type emptyData struct{}

type Database struct {
	src string
}

func (d *Database) getSrc() *string {
	if d.src == "" {
		homeDir, e := os.UserHomeDir()

		if e != nil {
			panic(fmt.Sprintf("Error al obtener directorio de usuario: %v", e.Error()))
		}

		d.src = filepath.Join(homeDir, "tasks.json")
	}
	return &d.src
}

func (d *Database) LoadDB(col interface{}) {
	dbDir := *d.getSrc()

	jsonData, e := ioutil.ReadFile(dbDir)

	if e != nil {
		// create DB file
		e := d.Update(col)

		if e != nil {
			fmt.Println("Unable to load DB: ", e)
			return
		}
		return
	}

	e = json.Unmarshal(jsonData, col)

	if e != nil {
		fmt.Println("Error trying to parse JSON db: ", e)
		return
	}

	return
}

func (d *Database) Update(data interface{}) error {
	fmt.Println("Updating database...")
	dir := d.getSrc()
	jsonData, e := json.Marshal(data)

	if e != nil {
		return errors.New(
			fmt.Sprintf("Error al convertir a JSON: %v", e.Error()),
		)
	}

	jsonStr := string(jsonData)

	e = ioutil.WriteFile(*dir, []byte(jsonStr), 0644)
	if e != nil {
		return errors.New(
			fmt.Sprintf("Error al guardar el archivo: %v", e.Error()),
		)
	}

	return nil
}
