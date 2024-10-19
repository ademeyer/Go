package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(FileName string) (io.Reader, error) {
	rdr, err := os.Open(FileName)
	//rdr, err := os.OpenFile(FileName, os.O_RDONLY, 0777)
	defer rdr.Close()
	if err != nil {
		return nil, nil
	}
	return rdr, nil
}

/*
	var device map[string]interface{} //inv{}
	if err := json.NewDecoder(strings.NewReader(jsonString)).Decode(&device); err != nil {
		fmt.Println("Err: ", err)
		return
	}

	fmt.Printf("device: %+v\n", device)
*/

func main() {
	data, err := ReadFile("./resp.txt")
	if err != nil {
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(data).Decode(&result); err != nil {
		fmt.Println("Err: ", err)
		return
	}

	inventory, ok := result["inventory"].(map[string]interface{})
	if !ok {
		fmt.Println("Err: ", ok)
		return
	}

	serial_nums, ok := inventory["serial_nums"].(map[string]interface{})
	if !ok {
		fmt.Println("Err: ", ok)
		return
	}

	for key, value := range serial_nums {
		encharge_capacity, ok := value.(map[string]interface{})["encharge_capacity"]
		if !ok {
			fmt.Printf("Key: %s does not have a cap field\n", key)
			continue
		}
		u, ok := encharge_capacity.(float64)
		fmt.Printf("Key: %s, cap: %v, type: %T\n", key, u, uint32(u))
	}
}
