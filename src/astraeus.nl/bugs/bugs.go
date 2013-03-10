/**
 * User: rnentjes
 * Date: 2/26/13
 * Time: 10:19 PM
 */
package bugs

import (
	"sync/atomic"
	"os"
	"encoding/gob"
	"strconv"
)

type Bug struct {
	id 			uint64
	Description string
}

var id uint64 = 0

func createId() uint64 {
	return atomic.AddUint64( &id, 1 )
}

func (b *Bug) Id() uint64 {
	return b.id
}

func CreateBug(description string) *Bug {
	b := new(Bug)
	b.id = createId()
	b.Description = description

	return b
}

type Bugs struct {
	Bugs 	map[uint64]*Bug
}

func createBugs() *Bugs {
	result := new(Bugs)
	result.Bugs = make(map[uint64]*Bug, 0)

	return result
}

func (bugs *Bugs) persistBug(bug *Bug) {
	filename := "data/bugs/bugs/"+strconv.FormatUint(bug.Id(), 10)+".bug"
	file, err := os.Open(filename)
	defer file.Close()

	if err == nil {
		os.Rename(filename, filename+".prev")
	}

	file, err = os.Create(filename)
	defer file.Close()

	gob.NewEncoder(file).Encode(bug)
}

func (bugs *Bugs) loadBug(id uint64) *Bug {
	filename := "data/bugs/bugs/"+strconv.FormatUint(id, 10)+".bug"
	file, err := os.Open(filename)
	defer file.Close()

	if err == nil {
		bug := new(Bug)
		gob.NewDecoder(file).Decode(&bug)

		bugs.Bugs[id] = bug

		return bug
	}

	return nil
}

func (bugs *Bugs) SaveBug(bug *Bug) {
	bugs.persistBug(bug)
	bugs.Bugs[bug.Id()] = bug
}

func (bugs *Bugs) GetBug(id uint64) *Bug {
	bug := bugs.Bugs[id]

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
