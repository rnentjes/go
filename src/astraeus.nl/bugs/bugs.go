/**
 * User: rnentjes
 * Date: 2/26/13
 * Time: 10:19 PM
 */
package bugs

import (
	"sync/atomic"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Persistent interface {
	Id() uint64
	PersistentTypeName() (string)
}

type Bug struct {
	id 			uint64
	Title		string
	Description string
}

var id uint64 = 0

func createId() uint64 {
	return atomic.AddUint64( &id, 1 )
}

func (b *Bug) Id() uint64 {
	return b.id
}

func CreateBug(title string, description string) *Bug {
	b := new(Bug)
	b.id = createId()
	b.Title = title
	b.Description = description

	return b
}

type Bugs struct {
	Bugs 	map[uint64]*Bug
}

func createBugs() *Bugs {
	result := new(Bugs)
	result.Bugs = make(map[uint64]*Bug, 0)

	result.loadBugs()

	return result
}
func (bugs *Bugs) loadBugs() {
	var maxid uint64 = 0

	dirname := "data/bugs/bugs/"
	entries, err := ioutil.ReadDir(dirname)

	if err != nil {
		panic("bugs directory not found")
	}

	for i := range entries {
		if !entries[i].IsDir() {
			name := entries[i].Name()
			if strings.HasSuffix(name, ".bug") {
				id, _ := strconv.ParseUint(name[0:len(name)-4], 10, 64)
				bugs.loadBug(id)

				if id > maxid {
					maxid = id
				}
			}
		}
	}

	atomic.AddUint64( &id, maxid - id )
}

func (bugs *Bugs) saveBug(bug *Bug) {
	filename := "data/bugs/bugs/"+strconv.FormatUint(bug.Id(), 10)+".bug"
	file, err := os.Open(filename)

	if err == nil {
		defer file.Close()
		os.Rename(filename, filename+".prev")
	}

	file, err = os.Create(filename)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	json.NewEncoder(file).Encode(bug)
}

func (bugs *Bugs) loadBug(id uint64) *Bug {
	filename := "data/bugs/bugs/"+strconv.FormatUint(id, 10)+".bug"
	file, err := os.Open(filename)
	defer file.Close()

	if err == nil {
		var bug Bug
		json.NewDecoder(file).Decode(&bug)
		bug.id = id

		bugs.Bugs[id] = &bug

		return &bug
	}

	return nil
}

func (bugs *Bugs) SaveBug(bug *Bug) {
	bugs.saveBug(bug)
	bugs.Bugs[bug.Id()] = bug
}

func (bugs *Bugs) DeleteBug(bug *Bug) {
	filename := "data/bugs/bugs/"+strconv.FormatUint(bug.Id(), 10)+".bug"
	err := os.Remove(filename)

	if err != nil {
		panic("Unable to remove buf file: "+strconv.FormatUint(bug.Id(), 10))
	}

	delete(bugs.Bugs, bug.Id())
}

func (bugs *Bugs) GetBug(id uint64) *Bug {
	var bug *Bug

	bug = bugs.Bugs[id]

	if bug == nil {
		bug = bugs.loadBug(id)
	}

	return bug
}

var bugs *Bugs

func GetBugs() *Bugs {
	if bugs == nil {
		bugs = createBugs()
	}

	return bugs
}
