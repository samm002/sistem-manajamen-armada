# Sistem Manajemen Armada - Project Overview

This repository contains a Fleet Management System with two main components:

- **backend/**: The main RESTful API service for managing and monitoring vehicle locations.
- **publish-data-script/**: A script (service) written in go for simulating and publishing random vehicle location data to the system via MQTT.

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

## Running with Docker Compose

To run the entire stack (backend, database, MQTT, RabbitMQ, and publisher) using Docker Compose:

```powershell
# Build and start all services
docker-compose up -d

# Access Root Main API Endpoint (Local Example)
http://localhost:3000 # or other port specified based on env
```

### Subscribed MQTT Topics

The system subscribes to the following 3 MQTT topics (one for each vehicle):

- `/fleet/vehicle/AB1234CDE/location`
- `/fleet/vehicle/L5432AB/location`
- `/fleet/vehicle/H7890SM/location`

You can view or modify these topics in `common/constant/mqtt.go`.

### Sending Payload to a Subscribed Topic

To send a payload to one of the subscribed topics, use the following JSON format:

```json
{
  "vehicle_id": "L1234AB", // optional, leave blank for randomized (example: L1234AB)
  "latitude": -6.2012,
  "longitude": 106.8168,
  "timestamp": 174798309 // optional, default is current timestamp
}
```

- You can use any MQTT client (e.g., MQTT Explorer, mosquitto_pub) to publish to one of the topics above.
- If `vehicle_id` is omitted or blank, the backend will generate a random vehicle ID.
- If `timestamp` is omitted, the backend will use the current timestamp.

Example using `mosquitto_pub`:

```powershell (default without .env.production example change)
mosquitto_pub -h localhost -p 1884 -u default -P default -t "/fleet/vehicle/L1234AB/location" -m '{"latitude":-6.201200,"longitude":106.816800}'
```

This will send a message to the topic `/fleet/vehicle/L1234AB/location`.

### Subscribing to Topics with mosquitto_sub

To monitor messages on a topic from your local machine, use:

```powershell (default without .env.production example change)
mosquitto_sub -h localhost -p 1883 -u default -P default -t "/fleet/vehicle/L1234AB/location" -v
```

- Replace the topic as needed to monitor other vehicles.
- Use `-t "#"` to subscribe to all topics.

- This will start all required services as defined in `docker-compose.yaml`.
- The backend will be available on the port specified in your `.env.production` (default: 3000).
- The publish-data-script will automatically send simulated data to the MQTT broker.

To stop all services:

```powershell
docker-compose down
```
