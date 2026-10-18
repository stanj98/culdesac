package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

const Version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)
	opts := Options{}
	if options != nil {
		opts = *options
	}
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err != nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)

}

func stat(path string) (fi os.FileInfo, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte("\n"))
	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}
}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	emps := []User{
		{"John", "23", "709823451", "Seaworld", Address{"bangalore", "karnata", "india", "876293"}},
		{"Melanie", "32", "709823451", "Google", Address{"chennai", "tamil. adu", "india", "876293"}},
		{"Tim", "30", "109823451", "MS", Address{"wayanad", "kerala", "india", "876293"}},
		{"Paul", "19", "709823341", "Apple", Address{"pune", "maharashtra", "india", "876293"}},
		{"Sam", "23", "709093451", "Remot-teas", Address{"sik", "skikkim", "india", "876293"}},
		{"Kam", "43", "293843451", "Lopeer", Address{"bangalore", "karnata", "india", "876293"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	//fetch all users
	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println((allusers))

	//delete a user from DB
	// if err := db.Delete("user", "john"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// //delete all users from DB
	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
