# shawty
A URL shortener made using Golang, Redis and go-fiber.

# Description

The main goal of shawty is to create a scalable url shortener using redis and golang. The initial relase focuses on creating a simple url shortener with future plans to include user context and saving shortened url's. For this I have planned to use PostgreSQL in the backend.

# Setup Guide

1. **Install Go**: Make sure Go is installed on your system. You can download it from [golang.org](https://golang.org/dl/).

2. **Clone the Repository**:
   ```bash
   git clone git@github.com:2k4sm/shawty.git
   cd shawty
   ```
3. **Install Dependencies**:
   ```bash
   go mod tidy
   ```
4. **env variables.**
      ```json
      DB_ADDR="db:6379"
      DB_PWD=""
      PORT=<port>
      API_QUOTA=<50>
      DOMAIN=<localhost:PORT/hosted domain>
      ```
      - Then Run
        ```bash
        export $(cat .env | xargs)
        ```
5. **To run the application**
   ```bash
   docker compose up --build // podman compose up --build
   ```

# Usage

```json
POST /api/v1 -> creates shortened url.
Payload : {
    "url" : <url-to-shorten>,
    "short": <customshort/ default-random-uuid>,
    "expiry": <duration-for-url-expiry-hours / default-24-hours>
}

Response : {
        "URL":             <original-url>,
		"CustomShort":     <custom-short-url>,
		"Expiry":          <expiry-time>,
		"XRateRemaining":  <rate-remaining>,
		"XRateLimitReset": <limit-reset-time>,
}
```

```json
GET /:url
Response : 301 permanent redirect to :url
```

# V0 Description

The version 0 is a simple url shortener. The main features include creating shortened url's which will be valid for 24 hours. To avoid abuse of the service IP based rate limiting has been used. The limit is set to 50 <can be configured> requests in a time window of 30 mins. The service generates a random uuid for the shortened url is no customShort is specified. and creates a url using it. The server resolves reqests very fast with little to no delay in url resolution.

# V1 Plans

Plans for version 1 are to create a user context. The user context will allow users to signup for the service and all their shortened url's will be associated with this user context. Once the url expires it will have a expired status and the user will be able to regenerate it again / revive it again for 24 hours with a put api call to the server. Specific api endpoints and user model has not been decided yet a schema design will make it more clear as of now this is the plan.

## Thanks for exploring shawty.