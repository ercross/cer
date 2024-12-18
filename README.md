# CER (Currency Exchange Rate)
CER is an implementation of a currency exchange rate API 

# Currency Exchange Rate API

## Overview
The **Currency Exchange Rate API** provides real-time exchange rates for currency pairs by integrating with multiple external exchange rate providers. The API ensures high availability and responsiveness by concurrently fetching exchange rates from two providers and returning the first available result.

---

## Features

1. **Fetch Exchange Rates**
    - Retrieve exchange rates for specified currency pairs (e.g., `USD-EUR`, `NGN-USD`).

2. **Concurrency**
    - Simultaneously query multiple providers to minimize response time.

3. **Error Handling**
    - Gracefully handles provider errors by waiting for the second provider if the first returns an error.

4. **Secure API Key Management**
    - Simulates secure storage and retrieval of API keys for provider authentication.

5. **Mocked External Services**
    - Includes mocked implementations of exchange rate providers for testing and demonstration purposes.

---

## Components
The project is composed of the following:

1. **API Service**
    - Written in Go, this service handles exchange rate requests and interacts with external providers.

2. **Exchange Rate Providers**
    - Implements a `RateProvider` interface to fetch rates from external APIs (mocked in this implementation).

---

## Non-Functional Requirements
This project is designed with the following considerations:

- **Real-Time Processing:**  
  Fetch and return exchange rates on-demand without persistent storage.

- **Concurrency:**  
  Simultaneously query multiple rate providers to ensure fast responses and fault tolerance.

- **Security:**  
  Simulates secure API key management using cloud-based practices (e.g., AWS).

- **Testing:**  
  Includes unit and integration tests for robustness.

---

## Architectural Decisions and Assumptions

1. **Stateless Design:**
    - No database is used; the system fetches and returns exchange rates in real time.

2. **Mocked Providers:**
    - External services are mocked for local development and testing.

3. **Error Handling Strategy:**
    - If the first provider fails, the API waits for the second provider's response.

4. **Technology Choices:**
    - Built with Go for performance and concurrency support.

---

## Getting Started

Follow these steps to set up the project locally:

### Prerequisites
- Go 1.23 or higher installed on your system.

### Setup Instructions

**Clone the Repository:**
   ```bash
   git clone https://github.com/ercross/cer.git
   cd cer
   ```

#### Setup Without Docker

1. **Install Dependencies:**
   `go mod tidy`

2. **Run the API Service:**
   `go run main.go`

3. **Export required environment variables**
   ```
   export AWS_CREDENTIALS="my-secret"
   export API_PORT=15001
   ```

4. **Access the API:**
    - The API will be available at `http://localhost:15001`.
    - Use tools like Postman or cURL to interact with the endpoints.

#### With Docker

The project includes a Makefile to simplify deployment and management using Docker Compose. The following commands are available:

1. **Deploy the Service:**
   `make deplo`
   This command starts the service using the `docker-compose` configuration.

2. **Stop the Service:**
   `make stop`
   Stops the running service.

3. **Restart the Service:**
   `make restart`
   Restarts the service by stopping and then starting it again.

4. **Stop and Clean Up:**
   `make stop_and_cleanup`
   Stops the service and removes all containers, networks, and volumes.


### Example Request

**Endpoint:**
```
GET /exchange-rate?pair=USD-EUR
```

**Response:**
```json
{
  "USD-EUR": 0.92
}
```

---

## Testing

### Run Unit Tests
Execute the following command to run all unit tests:
```bash
go test ./...
```

### Test Cases
- Valid currency pair returns the exchange rate.
- Invalid currency pair returns a `400 Bad Request`.
- Both providers fail to return a response results in a `503 Service Unavailable`.

---

## Limitations

1. **Mocked External Services:**
    - External providers are mocked for demonstration purposes.

2. **Single Environment Support:**
    - Optimized for local development; production-grade features like monitoring and scaling are not included.

---

## Future Enhancements
- Integrate with real external exchange rate APIs.
- Implement caching for frequently requested currency pairs to reduce latency.
- Implement rate limiter and exponential retry on the exchange.RateProviders concrete implementation
