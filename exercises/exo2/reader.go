package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"

func(reader spaceEraser) Read(buffer []byte) (int,error){
	n, err := reader.r.Read(buffer)
	nbCaract :=0
	for i := 0; i < n ; i++{
		if buffer[i] != 32 {
			buffer[nbCaract] = buffer[i] //le buffer, ne va pas changer de taille mais les éléments dans le buffer vont changer de place.
			nbCaract++
		}
	}
	return nbCaract, err
}