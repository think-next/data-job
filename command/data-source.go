package command

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const DefaultSize = 1000

type DataSource struct {
	Size int
}

func (source *DataSource) SetSize(size int) {
	source.Size = size
}

func (source *DataSource) GetSize() int {
	if source.Size == 0 {
		source.Size = DefaultSize
	}
	return source.Size
}

func (source *DataSource) MakeDataBuffer() chan string {
	return make(chan string, source.Size)
}

func (source *DataSource) GetFromFileByLine(file string, dataBuffer chan string) error {
	sourceFile, err := os.Open(file)
	if err != nil {
		fmt.Println("source file error ", err)
		return err
	}
	defer sourceFile.Close()

	r := bufio.NewReader(sourceFile)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read line error ", err)
			continue
		}

		line = strings.TrimSpace(line)
		dataBuffer <- line
	}

	close(dataBuffer)
	return nil
}
