#!/bin/bash
echo "Generate Swagger"
swag init
echo "Build Golang Folder"
go build -o ./bin/startup
echo "Done..."