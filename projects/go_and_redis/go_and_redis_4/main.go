package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisHost =  "192.168.1.14:6379" 
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
	http.HandleFunc("/data", uh.getdata)
	http.HandleFunc("/getdata", uh.getDataByKey)
	http.HandleFunc("/setdata", uh.setdata)

	fmt.Println("server running at 8090")
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

func (uh userHandler) getdata(w http.ResponseWriter, req *http.Request)  {
	var result string
	key := "key_test"
	ctx := context.Background()
	results := uh.client.Get(ctx, key)
	if results.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &response{
			Code: http.StatusInternalServerError,
			Message: results.Err().Error(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
		return
	}

	if results.String() == "" {
		result = "Hello Data Founded"
		time.Sleep(1 * time.Second)
		set := uh.client.Set(ctx,key,result,time.Duration(3) * time.Minute)
		if err := set.Err(); err != nil {
			fmt.Printf("unable to set data. error: %v", err)
			return
		}
	} else {
		result = results.String()
		w.WriteHeader(http.StatusOK)
		resp := &response{
		Code: http.StatusOK,
		Result: result,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
	} 
}

func (uh userHandler) getDataByKey(w http.ResponseWriter, req *http.Request)  {
	var result string
	key := req.URL.Query().Get("key")
	ctx := context.Background()
	results := uh.client.Get(ctx, key)
	if results.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := &response{
			Code: http.StatusInternalServerError,
			Message: results.Err().Error(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
		return
	}

	if results.String() == "" {
		result = "Hello Data Founded"
		time.Sleep(1 * time.Second)
		set := uh.client.Set(ctx,key,result,time.Duration(3) * time.Minute)
		if err := set.Err(); err != nil {
			fmt.Printf("unable to set data. error: %v", err)
			return
		}
	} else {
		result = results.String()
		w.WriteHeader(http.StatusOK)
		resp := &response{
		Code: http.StatusOK,
		Result: result,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
	} 
}

func (uh userHandler) setdata(w http.ResponseWriter, req *http.Request)  {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	ctx := context.Background()
	fmt.Println(key,value)

	if key == "" {
		fmt.Fprintf(w, "key is empty")
        return
	}

	if value == "" {
		fmt.Fprintf(w, "value is empty")
        return
	}

	set, err := uh.client.Set(ctx,key,value,time.Duration(1) * time.Minute).Result()
	fmt.Println(set)
	if err != nil {
		fmt.Printf("unable to set data. error: %v", err)
		return
	} else {
		res := fmt.Sprintf("succesfull set key: %v and value: %v", key, value)
		w.WriteHeader(http.StatusCreated)
		resp := &response{
		Code: http.StatusCreated,
		Result: res,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			return
		}
		w.Write(b)
	}
} 

func hello(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "hello world")
}

type userHandler struct {
	client *redis.Client
}

type response struct {
	Code int `json:"errorr_code"`
	Message string `json:"errorr_message"`
	Result string `json:"result"`
}