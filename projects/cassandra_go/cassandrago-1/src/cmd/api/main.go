package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/google/uuid"
)

func main()  {
	config := read()

	url := os.Getenv("Cassandra_URL")
	if url != "" {
		config.Database.Url = url
	}

	log.Println("Cassadra URL", url)

	Initialize(config)
}

// pkg/api/api.go
func Initialize(config Configuration) {
	// create a gin router with default middleware:
	router := gin.Default()

	
	fmt.Println("%+v\n", config)
	
	// register database, repositories and handler
	session := Connect(config.Database.Url, config.Database.Keyspace)

	repository := NewTodo(session)
	handler := NewTodoHandler(&repository)

	router.GET("/healtz", handler.Healtz)
	router.GET("api/v1/todo/:id", handler.Get)
	router.POST("api/v1/todo", handler.Create)

	// run the server
	router.Run(":" + config.Host.Port + "")
}

// pkg/handler/todohandler.go | package handler
type TodoHandler interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Healtz(*gin.Context)
}

type todoHandler struct {
	repo ITodo
}

func NewTodoHandler(repo *ITodo) TodoHandler {
	return &todoHandler{repo: *repo}
}

func (t *todoHandler) Healtz(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (t *todoHandler) Get(c *gin.Context)  {
	id := c.Param("id")
	todo, err := t.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequest("todo not found"))
	}
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func (t *todoHandler) Create(c *gin.Context)  {
	var todo Todo
	c.BindJSON(&todo)

	// todo := Todo{}

	todo.Id = uuid.New().String()
	// todo.Content = "test content"
	// todo.Title = "test title"
	// use log.info from gommon/log
	log.Println("Data is", todo)
	data, err := t.repo.Save(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, BadRequest("insert operation failed"))
	} else {
		c.JSON(http.StatusCreated, gin.H{"todo": data})
	}
}
// pkg/repository/db/todo.go | package db
// TodoRepository is an interface Save method take model.Todo return *model.Todo, error GetById take id parameter return *model.Todo, error
type ITodo interface {
	Save(todo Todo) (*Todo, error)
	GetById(id string) (*Todo, error)
}

type todo struct {
	session *gocql.Session
}

// NewTodoRepository take session which is use for repository implementation
func NewTodo(session *gocql.Session) ITodo {
	return &todo{session: session}
}
func (t *todo) Save(todo Todo) (*Todo, error){
	var query string = "INSERT INTO todo(id,title,content) VALUES(?,?,?)"
	if err := t.session.Query(query, todo.Id, todo.Title, todo.Content).Exec(); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *todo) GetById(id string) (*Todo, error){
	var todo  Todo
	var query string = "SELECT id,title,content FROM todo where id=?"
	if err := t.session.Query(query, id).Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {
		if err == gocql.ErrNotFound {
			return nil, err
		}
		return nil, err
	}

	return &todo, nil
}
// pkg/clients/cassandra/cassandra_db.go
func Connect(url string, keyspace string) *gocql.Session {
	cluster := gocql.NewCluster(url)
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		// use logrus
		log.Fatal(err)
	}
	return session
}

// pkg/model/todo.go | package model
type Todo struct {
	Id string 
	Title string
	Content string
}
func read() Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return config
}

// utils/errors.go

type RestErr struct {
	Message string
	Status int
	Error string
}

// BadRequest error reurn RessErr with bad_request status and messages
func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad request",
	}
}

// NotFound error reurn RessErr with NotFound status and messages
func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not found",
	}
}

// InternalServerError error reurn RessErr with internal server error status and messages
func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal server error",
	}
}

// utils/config.go
type Configuration struct {
	Database DBSettings
	Host HostSettings
}

type DBSettings struct {
	Url string
	Keyspace string
}

type HostSettings struct {
	Port string
}