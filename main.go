package main

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	
)

type person struct{
	Name string `json:name`
	Age  string `json:age`
	Email string `json:email`
}
func main() {
	r := gin.Default()

	// Define the route for uploading CSV and converting to JSON
	r.POST("/convert", convertHandler)

	r.Run(":8080")
}

func convertHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("csv_file")
	if err != nil {
		c.JSON(400, gin.H{"error": "No CSV file uploaded"})
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	// var jsonData []map[string]string
	persons:=make([]person,0)
	// Read the CSV data line by line
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "Error reading CSV"})
			return
		}

		// Create a map to store the record data
		recordData := make(map[string]string)
        var p person
		// Map the header fields to the record values
		for i := 0; i < len(record); i++ {
			recordData[header[i]] = record[i]
            if header[i]=="header1"{
				p.Name=record[i]
			}else if header[i]=="header2"{
				p.Age=record[i]
			}else{
				p.Email=record[i]
			}
		}
		persons=append(persons, p)

		// Append the record data to the JSON slice
		// jsonData = append(jsonData, recordData)
	}
  
	

	c.JSON(200, persons)
}

var header []string

func init() {
	// Parse the header fields from the first line of the CSV file
	csvHeader := "header1,header2,header3" // Replace with the actual header line of your CSV file
	header = strings.Split(csvHeader, ",")
}
