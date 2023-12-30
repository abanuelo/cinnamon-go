#!/bin/bash

# Define the URL to send requests to
URL="http://localhost:8089/world"

# Number of requests to send
NUM_REQUESTS=600

# Array to store responses
responses=()

# Perform 50 HTTP GET requests in background
for ((i=1; i<=$NUM_REQUESTS; i++)); do
  curl -s "$URL" &
done