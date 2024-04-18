package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	address() string
	isAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleError((err))
	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type loadBalancer struct {
	port       string
	roundRobin int32
	servers    []Server
}

func newLoadBalancer(port string, servers []Server) *loadBalancer {
	return &loadBalancer{
		port:       port,
		roundRobin: 0,
		servers:    servers,
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error occured %v/n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) address() string {
	return s.addr
}
func (s *simpleServer) isAlive() bool {
	return true
}
func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)

}

func (lb *loadBalancer) getNextAvailalbaleServer() Server {
	server := lb.servers[lb.roundRobin%int32(len(lb.servers))]
	if !server.isAlive(){
		lb.roundRobin++
		server = lb.servers[lb.roundRobin%int32(len(lb.servers))]
	}
	lb.roundRobin++
	return server
}
func (lb *loadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailalbaleServer()
	fmt.Printf("target server is %s\n",targetServer.address())
	targetServer.Serve(w, r)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.google.com"),
		newSimpleServer("https://www.cuet.ac.bd"),
	}
	lb := newLoadBalancer("8000", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)
	http.ListenAndServe(":"+lb.port, nil)
}
