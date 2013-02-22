/**
 * User: rnentjes
 * Date: 2/22/13
 * Time: 8:46 PM
 */
package main

import (
	"encoding/gob"
	"os"
	"bufio"
)

type Command struct {
 	// boolean remove
	// some id
	// type of data
	// data
}

type PersistentStore struct {
	// map of types of data with therein
	// map of id with data objects
}

func (* PersistentStore) loadData(path string) {
    var command Command
	var (
		file *os.File
		err
	)

	if file, err = os.Open(path); err != nil {
		return nil
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	dec := gob.NewDecoder(&reader);

	dec.Decode(&store);

	return &store
}

func main() {

}
