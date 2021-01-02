# How to run this project?
1. Clone the repo and `cd api-based-resiliency/pay-bill-resilience-service`
2. `go mod download`
3. `docker-cmopose up`
4. `127.0.0.1:8000/api/update-biller-health` with POST req body ```
{
    "biller_1": "UP",
    "biller_2": "DOWN",
    "biller_3": "UP"
}```
5. `127.0.0.1:8000/api/get-biller-status?code=biller_2`, GET req

# How to read the project
1. When you are hitting the POST request the matrix will be stored in the Redis with a TTL. It's described in the file, `pay-bill-resilience-service`.
2. When you are making a GET request to get the status of a specific code, it will collect the matrix from the Redis and send it back to you. It's described in the file, `biller-status.go`
3. If no entry is found in Redis, or if there is any error a default value `UP` will be sent.

# System Design
<img width="1258" alt="Screenshot 2021-01-02 at 10 38 31 PM" src="https://user-images.githubusercontent.com/19304394/103461808-488cb900-4d4b-11eb-9f25-15609e13217c.png">

# Description
If you are in `Layer 0` and wants to know the health status of `Layer 2` and if there is no circuit breaker in `Layer 1`
then this approach is helpful.

# Dependency
1. gin-gonic: Used in Transport layer
2. redi-go: redis client in go