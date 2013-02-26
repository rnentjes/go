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
	Bugs 	[]*Bug
}

func CreateBugs() *Bugs {
	result := new(Bugs)
	result.Bugs = make([]*Bug, 0)

	return result
}

func (bugs *Bugs) AddBug(bug *Bug) {
	bugs.Bugs = append(bugs.Bugs, bug)
}
