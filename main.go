package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

var rjs *rejson.Handler

type Data struct {
	Key   string                 `json:"key"`
	Value map[string]interface{} `json:"value"`
}

var certFile string
var keyFile string

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})
	/* client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}) */
	err := waitForRedis(client, 5*time.Second, 5)
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		os.Exit(1)
	}

	rjs = rejson.NewReJSONHandler()
	rjs.SetGoRedisClient(client)

	router := gin.Default()

	router.POST("/api/data", setData)
	router.GET("/api/data/:key", getData)
	router.GET("/:tenant/:key", getRedis)
	router.POST("/:tenant", putRedis)
	certFile = "cert.pem"
	keyFile = "key.pem"
	//router.RunTLS(":443", "/app/cert.pem", "/app/key.pem")
	router.RunTLS(":443", certFile, keyFile)

	for {
		fmt.Println("App is running...")
		time.Sleep(5 * time.Second)
	}
}

func waitForRedis(client *redis.Client, timeout time.Duration, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		err := client.Ping().Err()
		if err == nil {
			return nil
		}
		time.Sleep(timeout)
	}
	return fmt.Errorf("failed to connect to Redis after %d retries", maxRetries)
}
func setData(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	key := data["key"].(string)
	value := data["value"].(string)

	// Convert the value into a JSON string
	jsonValue := fmt.Sprintf(`{"value": "%s"}`, value)

	// Store the JSON string in Redis
	rjs.JSONSet(key, ".", jsonValue)

	c.JSON(200, gin.H{"status": "success"})
}

func getData(c *gin.Context) {
	key := c.Param("key")

	data, err := rjs.JSONGet(key, ".")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(data.(string)), &result)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to unmarshal data"})
		return
	}

	c.JSON(200, result)
}
func getRedis(c *gin.Context) {
	tenant := c.Param("tenant")
	key := c.Param("key")

	// Construct the Redis key using the tenant ID and key
	redisKey := tenant + ":" + key

	// Retrieve the JSON value from Redis as bytes
	jsonBytes, err := rjs.JSONGet(redisKey, ".")
	if err != nil {
		c.JSON(404, gin.H{"error": "Key not found"})
		return
	}

	// Perform a type assertion to convert the bytes to []byte
	bytes, ok := jsonBytes.([]byte)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid data type"})
		return
	}

	// Unmarshal the JSON bytes into a map[string]interface{}
	var jsonData map[string]interface{}
	err = json.Unmarshal(bytes, &jsonData)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error decoding JSON"})
		return
	}

	// Return the JSON data as a response
	c.JSON(200, jsonData)
}
func putRedis(c *gin.Context) {
	tenant := c.Param("tenant")
	var data Data

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Construct the Redis key using the tenant ID and key
	redisKey := tenant + ":" + data.Key
	_, err := rjs.JSONSet(redisKey, ".", data.Value)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error saving JSON data to Redis"})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
