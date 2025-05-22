# Sistem Manajemen Armada - Project Overview

This repository contains a Fleet Management System with two main components:

- **backend/**: The main RESTful API service for managing and monitoring vehicle locations.
- **publish-data-script/**: A script for simulating and publishing random vehicle location data to the system via MQTT.

## Backend

- Provides RESTful APIs for vehicle location data (create, retrieve, history, latest location)
- Integrates with MQTT and RabbitMQ for real-time and event-driven messaging
- Connects to a PostgreSQL database
- See [backend/README.md](backend/README.md) for setup, environment variables, and API documentation

## Publish Data Script

- Standalone Go script to generate and publish random vehicle location data
- Useful for testing and simulating incoming data
- See [publish-data-script/README.md](publish-data-script/README.md) for usage instructions

## API Documentation

- Full API documentation and example requests are available in the public Postman collection:
  [https://documenter.getpostman.com/view/26314293/2sB2qZG3di](https://documenter.getpostman.com/view/26314293/2sB2qZG3di)

## Getting Started

1. Clone the repository
2. Follow the setup instructions in each component's README
3. Start the backend services
4. (Optional) Run the publish-data-script to simulate data sending data via MQTT