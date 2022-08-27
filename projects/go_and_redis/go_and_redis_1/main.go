package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

const (
	redisHost =  "192.168.1.7:6379" 
)

	
func main()  {
	
	opt := &redis.Options{Addr: redisHost}
	client := redis.NewClient(opt)

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to connect with redis instance at %s - %v", redisHost,err)
	}

	uh := userHandler{client: client}

	http.HandleFunc("/ping", uh.ping)
	http.HandleFunc("/foo", uh.foo)
	http.HandleFunc("/", hello)


	log.Fatal("server running at 8090", http.ListenAndServe(":8090", nil))
}

func (uh userHandler) ping(w http.ResponseWriter, req *http.Request)  {
	pong, _ := uh.client.Ping(context.Background()).Result()
	fmt.Fprintf(w, "result: %s", pong)
}

func (uh userHandler) foo(w http.ResponseWriter, req *http.Request)  {
	foo := uh.client.Get(context.Background(), "foo")
	if err := foo.Err(); err != nil {
		str := fmt.Sprintf("unable get data. error : %v ", err)	
		http.Error(w, str,http.StatusInternalServerError)
		// fmt.Fprintf(w, "unable get data. error : %v ", err)	
		return
	}

	res, err := foo.Result()
	if err != nil {
		fmt.Fprintf(w, "unable get data. error : %v ", err)	
		return
	}

	fmt.Fprintf(w, "result: %s", res)
}

func hello(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello world")
}

type userHandler struct {
	client *redis.Client
}