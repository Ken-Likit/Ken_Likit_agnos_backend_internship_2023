# Ken_Likit_agnos_backend_internship_2023

# Things needed to install.  
1. Have Go programming language installed, code written with 1.20.4
2. Install all files in the github repo into a folder.

# To run on local computer
1. Open folder location in command line.
2. Use the command `go run .` This will start the server on local host port 8081.
3. To input a password for assement use curl command and change init_password field for individual test.
```curl --location 'http://localhost:8081/passwordValidation' \
--header 'Content-Type: application/json' \
--data '{
 "init_password": "sS13ds"
}'
```
4. This will return a json object with the number of changes to make a strong password.

# To run test
1. Open folder location in command line.
2. Use command `go test -v` to run test for strongPassword function. 
