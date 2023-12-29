#!/bin/bash

# Define the URL to send requests to
URL="http://localhost:8089/test"

# Number of requests to send
NUM_REQUESTS=50

# Array to store responses
responses=()

# Perform 50 HTTP GET requests
for ((i=1; i<=$NUM_REQUESTS; i++)); do
  # Send HTTP GET request using curl and capture response
#   response=$(curl -s "$URL")

  curl -s "$URL" &


  # Store the response in the array
#   responses+=("$response")
done

# echo "Iterating through responses array:"
# for ((i=0; i<${#responses[@]}; i++)); do
#   echo "Response $((i + 1)): ${responses[i]}"
# done