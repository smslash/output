package logic

import (
	"fmt"
	"os"
)

func Check(s []string) {
	fileName := s[1][9:]
	stringValue := s[2]
	bannerName := s[3]

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("./banners/" + bannerName + ".txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileHash := MD5(string(data))
	if !CheckHash(fileHash, bannerName) {
		fmt.Println("Error: hash code of", bannerName+".txt", "has been changed")
		return
	}

	result := Process(stringValue, data)

	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
