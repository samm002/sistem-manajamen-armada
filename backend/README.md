# Backend API - Sistem Manajemen Armada

This folder contains the backend service for the Fleet Management System. It provides RESTful APIs for managing vehicle locations, including creation, update, retrieval, and history queries.

## Features

- CRUD operations for vehicle locations, including:
  - Query vehicle location history by timestamp range
  - Query latest vehicle location
  - (Exclude Update & Delete, more focus on monitoring data and handle event)
- MQTT and RabbitMQ integration for real-time and event-driven messaging

## Prerequisites

- Go 1.18 or newer
- PostgreSQL database
- MQTT broker and RabbitMQ server for messaging features
- Environment variables configured
- Setup Atlas Migration for GORM, reference : <https://atlasgo.io/guides/orms/gorm/getting-started>

## Setup

1. **Install dependencies**

   ```powershell
   go mod tidy
   ```

2. **Configure environment variables**
   Create a `.env` file in the root directory (rename .env.example to .env or .env.production for container setup), then set the following variables:

   - `PORT` (e.g. `3000`)
   - `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSL_MODE`, `DB_TIME_ZONE`
   - `MQTT_PROTOCOL`, `MQTT_BROKER_URL`, `MQTT_BROKER_PORT`, `MQTT_BROKER_USERNAME`, `MQTT_BROKER_PASSWORD`, `MQTT_CLIENT_ID`
   - `RABBITMQ_URL`

   Example `.env`:

   ```env
   PORT=3000
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=admin
   DB_PASSWORD=yourpassword
   DB_NAME=sistem_manajemen_armada
   DB_SSL_MODE=disable
   DB_TIME_ZONE="Asia/Shanghai"
   MQTT_PROTOCOL=mqtt
   MQTT_BROKER_URL=localhost
   MQTT_BROKER_PORT=1883
   MQTT_BROKER_USERNAME=guest
   MQTT_BROKER_PASSWORD=guest
   MQTT_CLIENT_ID=backend
   RABBITMQ_URL=amqp://guest:guest@localhost:5672/
   ```

3. **Run database migrations**
   (If using Atlas or another migration tool, see the `database/migration` folder for instructions.)

   ```powershell
   # If Atlas Migration configured already, you can simply run
   atlas migrate apply --env gorm

   # Or just run the migration with other migration tools or directly via PSQL or PGADMIN query tool
   ```

## How to Run

1. **Build the backend**

   ```powershell
   go build -o sistem-manajemen-armada.exe
   ```

2. **Run the backend**

   ```powershell
   .\sistem-manajemen-armada.exe
   ```

   Or, run directly with Go:

   ```powershell
   go run main.go
   ```

## API Endpoints

See the full API documentation and example requests in the public Postman collection:

**Postman Collection:**
<https://documenter.getpostman.com/view/26314293/2sB2qZG3di>

### Main Endpoints

- `POST   /vehicle-locations` : Create a new vehicle location
- `GET    /vehicle-locations` : Get all vehicle locations (ordered by vehicle_id asc, timestamp desc)
- `GET    /vehicle-locations/:vehicleId/history?start=...&end=...` : Get location history for a vehicle in a timestamp range
- `GET    /vehicle-locations/:vehicleId/location` : Get the latest location for a vehicle
- `PUT    /vehicle-locations/:vehicleId` : Update a vehicle location (Not Used)
- `DELETE /vehicle-locations/:vehicleId` : Delete a vehicle location (Not Used)
