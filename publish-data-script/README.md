# Publish Data Script

This folder contains a Go program to publish random vehicle location data to an MQTT broker for testing and simulation purposes.

## Features

- Publishes random vehicle location data (vehicle_id, latitude, longitude, timestamp)
- Configurable via environment variables
- Uses MQTT protocol

## Prerequisites

- Go 1.18 or newer
- Access to an MQTT broker (local or remote)
- Set up the required environment variables

## Setup

1. **Install dependencies**

   ```powershell
   go mod tidy
   ```

2. **Configure environment variables**
   Create an `.env` file (rename .env.example to .env or .env.production for container setup), then set the following variables:

   - `MQTT_PROTOCOL` (e.g. `mqtt`)
   - `MQTT_BROKER_URL` (e.g. `localhost`)
   - `MQTT_BROKER_PORT` (e.g. `1883`)
   - `MQTT_BROKER_USERNAME` (if needed)
   - `MQTT_BROKER_PASSWORD` (if needed)
   - `MQTT_CLIENT_ID` (any unique string)

   **Example `.env.example`:**

   ```env
   MQTT_PROTOCOL=mqtt
   MQTT_BROKER_URL=localhost
   MQTT_BROKER_PORT=1883
   MQTT_BROKER_USERNAME=guest
   MQTT_BROKER_PASSWORD=guest
   MQTT_CLIENT_ID=publish-script
   ```

## How to Run

1. **Build the program**

   ```powershell
   go build -o publish-data-script.exe

   or

   go build
   ```

2. **Run the program**

   ```powershell
   .\publish-data-script.exe
   ```

   Or, run directly with Go:

   ```powershell
   go run main.go
   ```

## Customization

- You can adjust the data generation logic in `app/common/util/generate_random_vehicle_location_data.go`.
- Update the MQTT topic or payload format as needed in `app/common/constant/mqtt.go`
