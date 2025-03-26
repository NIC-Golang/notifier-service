# Notifier Service

Notifier Service is a microservice responsible for sending notifications to users. It handles various types of notifications such as email, SMS, and push notifications. The service is built with Go and integrates with external services for message delivery, providing an efficient and scalable solution for notifying users in real-time.

## Main features

- Email notifications: Send welcome emails, order confirmations, password reset links, and other types of email notifications.
- SMS notifications: Send SMS alerts for critical actions like order status updates, security alerts, etc.
- Push notifications: Notify users via push messages on mobile devices or browsers.
- Event-driven architecture: Listens to events from other microservices (e.g., user registration, order creation) to trigger notifications.
- Scalable and extensible: Easily extendable to support additional notification channels or services.
- Integration with external providers: Uses external services (e.g., SendGrid for emails, Twilio for SMS) for message delivery.
- Real-time notifications: Ensures timely delivery of notifications based on events and user preferences.

## Requirements:

- Golang 1.23+
- Docker & Docker Compose
- RabbitMQ (preferably)

## Setup Instructions  

Clone the repository and set up the environment:  

```sh
git clone https://github.com/NIC-Golang/notifier-service.git
cd notifier-service
cp .env.example .env
```
## Available Makefile Commands:
```sh
make help       # Show available commands  
make install    # Install dependencies (go mod tidy)  
make run        # Run the server  
make stop       # Stop the server  
make restart    # Restart the server  
make compile    # Compile the application  
make clean      # Clean the build cache  
make test       # Run the test container (docker compose -f docker-compose.test.yml up -d)  
make build      # Build the docker container  
make up         # Run the docker container  
make down       # Stop the docker container
```

## Running the service:
Make sure Docker Engine is running, then execute:
```
docker compose build
docker compose up
```
To stop the service:
```
docker compose down
```
If you are using an older Docker version, use docker-compose instead of docker compose.