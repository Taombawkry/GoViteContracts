#!/bin/bash

# Check for any process running on port 8080 and kill it
echo "Checking for processes running on port 8080..."
PID=$(lsof -t -i:8080)

if [ -n "$PID" ]; then
  echo "Killing process $PID running on port 8080..."
  kill -9 $PID
else
  echo "No process running on port 8080."
fi

# Change to the server directory
cd /server || { echo "Failed to change directory to /server"; exit 1; }

# Build the Go application
echo "Building the Go application..."
go build -o go-webserver .

if [ $? -eq 0 ]; then
  echo "Build succeeded."

  # Run the web server in the background
  echo "Starting the web server..."
  #nohup ./go-webserver -env production >& output.log &
  nohup ./go-webserver -env production > output.log 2>&1 &

  echo "Web server started and running in the background."
else
  echo "Build failed. Exiting."
  exit 1
fi
