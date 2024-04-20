<h1><strong>📎 Load Balancing with Proxy Server in Go</strong></h1>


This project showcases the implementation of a simple load balancing mechanism using a proxy server written in Go (Golang). Load balancing is a critical aspect of modern web architecture, facilitating the efficient distribution of incoming network traffic across multiple servers. This distribution is vital for ensuring high availability and reliability of web applications.

---

## Features

- **Proxy Server:** Implements a lightweight proxy server in Go to distribute incoming requests to multiple backend servers.
- **Load Balancing Algorithm:** Utilizes the Round Robin algorithm to evenly distribute requests among the available backend servers.
- **Dynamic Configuration:** The proxy server allows for dynamic addition or removal of backend servers without requiring a restart.
- **Health Checks:** Includes a basic health check mechanism to determine the availability of backend servers. If a server becomes unavailable, the load balancer skips it and forwards the request to the next available server.

## Server Structure

The `Simple Server` structure represents a simple backend server. It contains the following properties:

- **Address**: The address of the backend server.
- **Proxy**: A reverse proxy for the backend server.

```go
type SimpleServer struct {
    Address string               // Address of the backend server
    Proxy   *httputil.ReverseProxy // Reverse proxy for the backend server
}
```



To create a simple server structure, you can define a `newSimpleServer` function. This function takes the address of the backend server as input and returns a pointer to a `simpleServer` structure.

## Load Balancer Structure

The `Load Balancer` structure represents the core of the load balancing functionality in this project. It consists of the following properties:
```go
type LoadBalancer struct {
    Port       string     // Represents the port number on which the load balancer listens for incoming requests. This property is essential for configuring the network interface of the load balancer.
    RoundRobin int32      // An integer counter used for round-robin load balancing. It keeps track of the next available backend server to distribute incoming requests evenly.
    Servers    []Server   // This property is a slice of type Server, holding references to the backend servers registered with the load balancer. Each server in this slice is responsible for handling a portion of the incoming traffic.
}
```

## Server Interface

The `Server` interface defines the contract that all backend servers must adhere to in order to be compatible with the load balancer. It includes the following methods:

```go
type Server interface {
    address() string  // Returns the address of the backend server.
    isAlive() bool    // Checks if the backend server is alive and available.
    Serve(http.ResponseWriter, *http.Request) // Serves HTTP requests by proxying them to the backend server.
}
```

### Load Balancer Methods

- **`GetNextAvailableServer() Server`**: Retrieves the next available backend server based on the round-robin algorithm.
  - Returns:
    - The next available backend server.
  
- **`ServeProxy(w http.ResponseWriter, r *http.Request)`**: Serves HTTP requests by proxying them to the next available backend server, ensuring load balancing across all registered servers.

These methods provide the functionality to initialize and run the simple server and load balancer, as well as to retrieve the next available server for load balancing purposes and to serve HTTP requests by proxying them to the backend servers.



