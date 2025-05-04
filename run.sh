#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Starting Tumdum Backend Setup...${NC}"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed. Please install Go first.${NC}"
    exit 1
fi

# Check if PostgreSQL is installed
if ! command -v psql &> /dev/null; then
    echo -e "${RED}Error: PostgreSQL is not installed. Please install PostgreSQL first.${NC}"
    exit 1
fi

# Check if .env file exists
if [ ! -f .env ]; then
    echo -e "${YELLOW}Creating .env file from .env.example...${NC}"
    cp .env.example .env
    echo -e "${GREEN}Please update the .env file with your configuration values.${NC}"
fi

# Check if config.yaml exists
if [ ! -f config/config.yaml ]; then
    echo -e "${YELLOW}Creating config.yaml from config.yaml.example...${NC}"
    cp config/config.yaml.example config/config.yaml
    echo -e "${GREEN}Please update the config.yaml file with your configuration values.${NC}"
fi

# Install dependencies
echo -e "${YELLOW}Installing dependencies...${NC}"
go mod download

# Run database migrations
echo -e "${YELLOW}Running database migrations...${NC}"
psql -U postgres -f database/sql/schema.sql

# Build the project
echo -e "${YELLOW}Building the project...${NC}"
go build -o tumdum-backend

# Run the application
echo -e "${GREEN}Starting the application...${NC}"
./tumdum-backend 