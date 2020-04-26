/**
Partitioned Byte Array

This package wrapper around a byte array.
It allow you to manage nibbles and stuff like that.
*/

package pbytearray

import "github.com/golang-collections/go-datastructures/bitarray"

type PByteArray struct {
	bytes []byte
	pSize uint8
}

func (a PByteArray) validate() {
	if e.pSize > 8 {
		panic("PByteArray should never recieve partition size larger than 8")
	}
}

func (a PByteArray) write(pos uint, val int) {
	
}
