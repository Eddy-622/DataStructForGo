package test

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	// 访问百度
	request, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	client := http.Client{}
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	//for {
	//	line, _, err := reader.ReadLine()
	//	if err != nil && err == io.EOF {
	//		break
	//	}
	//	t.Log(string(line))
	//}

	s := make([]byte, 1024)
	file, _ := os.Create("test.txt")
	defer file.Close()
	for {
		n, err := reader.Read(s)
		if err != nil && err == io.EOF {
			break
		}
		io.WriteString(file, string(s[:n]))
	}
}
