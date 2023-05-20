package utils

import (
  "os"
  "log"
  "encoding/json"
)

func StoreMapAsJSONFile(response interface{}, id string) (error) { 
  jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
		return err
	}

	file, err := os.Create(id + ".json")
	if err != nil {
		log.Fatal("Error creating JSON file:", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file:", err)
		return err
	}
  return nil
}
