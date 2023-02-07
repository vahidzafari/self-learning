package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// In this subsection, you learn how to read from the /dev/random system device. The
// purpose of the /dev/random system device is to generate random data, which you
// might use for testing your programs or, in this case, as the seed for a random number
// generator.

// There are two representations named little endian and big endian that have to do
// with the byte order in the internal representation. In our case, we are using little
// endian. The endian-ness has to do with the way different computing systems order
// multiple bytes of information.

// In a big endian representation, bytes are read from left to right, while little endian
// reads bytes from right to left. For the 0x01234567 value, which requires 4 bytes for
// storing, the big endian representation is 01 | 23 | 45 | 67 whereas the little endian
// representation is 67 | 45 | 23 | 01.

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var seed int64
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
