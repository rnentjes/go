/**
 * User: rnentjes
 * Date: 2/22/13
 * Time: 8:46 PM
 */
package persistence

import (
	"encoding/gob"
	"os"
	"bufio"
	"fmt"
	"reflect"
	"sync/atomic"
)

type Command struct {
	remove 		bool
	data   		Persistent
}

type Persistent interface {
	Type() string
	Id() uint64
}

var id uint64 = 0

func NextId() uint64 {
	return atomic.AddUint64( &id, 1 )
}

type PersistentParent struct {
	id uint64
}

func (p *PersistentParent) Type() string {
	return reflect.TypeOf(p).String()
}

func (p *PersistentParent) Id() uint64 {
	if (p.id == 0) {
		p.id = NextId()
	}

	return p.id
}

type PersistentStore struct {
	// map of types of data with therein
	// map of id with data objects
	Data 	map[string]map[uint64]Persistent
	Writer 	gob.Decoder
}

func Open(path string) *PersistentStore {
	return new(PersistentStore)

}

func (p *PersistentStore) LoadData(path string) {
	var (
		file *os.File
		err error
		command Command
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	dec := gob.NewDecoder(reader)

	if err := dec.Decode(&command); err != nil {
		p.execute(&command)
	}
}

func (p *PersistentStore) SaveCommand(path string, command Command) {
	var (
		file *os.File
		err error
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	dec := gob.NewDecoder(reader)

	if err := dec.Decode(&command); err != nil {
		p.execute(&command)
	}
}

func(p *PersistentStore) Save(pers Persistent) {
	if (p.Data == nil) {
		fmt.Println("p.Data == nil")
		p.Data = make(map[string]map[uint64]Persistent)
	}

	if (p.Data[pers.Type()] == nil) {
		p.Data[pers.Type()] = make(map[uint64]Persistent)
	}

	p.Data[pers.Type()][pers.Id()] = pers
}

func(p *PersistentStore) Find(dataType string, id uint64) Persistent {
	var result Persistent

	if (p.Data != nil) {
		if (p.Data[dataType] != nil) {
			result = p.Data[dataType][id]
		}
	}

	return result
}

func(p *PersistentStore) FindMap(dataType string) map[uint64]Persistent {
	if (p.Data != nil) {
		return p.Data[dataType]
	}

	return nil
}

func (p *PersistentStore) saveAndExecute(command *Command) {
	// decode to outputstream
	// if successful, execute command
	p.execute(command)
}

func (* PersistentStore) execute(command *Command) {
}
