#!/bin/bash

# Define the base URL
BASE_URL="http://localhost:3000"

# Test the user registration
echo "Testing user registration..."
curl -X post -H "Content-Type: application/json" -d '{"username":"JohnDoe","orgName":"Org1"}' $BASE_URL/users
echo -e "\n"

# Sleep for 2 seconds
sleep 2

# Test adding a BPHR record
echo "Testing adding a BPHR record..."
curl -X post -H "Content-Type: application/json" -d '{"bphrId":"BPHR001","systolic":120,"diastolic":80,"pulseRate":72,"username":"JohnDoe"}' $BASE_URL/addRecord
echo -e "\n"

# Sleep for 2 seconds
sleep 2

# Test getting a BPHR record by ID
echo "Testing getting a BPHR record by ID..."
curl -X get $BASE_URL/getRecord/BPHR001
echo -e "\n"

# Sleep for 2 seconds
sleep 2

# Test updating a user's data
echo "Testing updating a user's data..."
curl -X put -H "Content-Type: application/json" -d '{"username":"JohnDoe","newUsername":"JohnDoeUpdated","orgName":"Org1"}' $BASE_URL/users
echo -e "\n"

# Sleep for 2 seconds
sleep 2

# Test getting all records for a user
echo "Testing getting all records for a user..."
curl -X get $BASE_URL/getAllRecords/JohnDoeUpdated
echo -e "\n"

# Add any other tests for functionalities you've added

# End of script
