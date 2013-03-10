/**
 * User: rnentjes
 * Date: 2/26/13
 * Time: 10:19 PM
 */
package bugs

import "sync/atomic"

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

func CreateBugs() *Bugs {
	result := new(Bugs)
	result.Bugs = make(map[uint64]*Bug, 0)

	return result
}

func (bugs *Bugs) SaveBug(bug *Bug) {
	bugs.Bugs[bug.Id()] = bug
}

func (bugs *Bugs) GetBug(id uint64) *Bug {
	return bugs.Bugs[id]
}

var bugs *Bugs

func GetBugs() *Bugs {
	if bugs == nil {
		bugs = CreateBugs()
	}

	return bugs
}
