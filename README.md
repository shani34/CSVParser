# CSVParser

Launch Postman and create a new request.

Set the request type to "POST."

Enter the URL for the API endpoint where the conversion is handled. For example, if you are running the server locally on port 8080, the URL would be: http://localhost:8080/convert

In the "Body" tab, select "form-data."

Add a key-value pair in the form-data:

Key: csv_file
Value: Select a CSV file from your local system using the "Choose Files" button.
Click the "Send" button to make the POST request.

Postman will execute the request, and the response will contain the converted JSON data.