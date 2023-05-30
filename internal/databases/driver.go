package driver

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	// "github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	redisClientOnce sync.Once
	redisClient     RedisClient
)

func init() {
	if err := godotenv.Load("./config/.env"); err != nil {
		panic(err)
	}
}

// Config return database with listed configuration below
func Config() *sql.DB {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	// migration
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./internal/databases/schemas/",
		"postgres",
		driver,
	)
	if err != nil {
		fmt.Println(err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("error applying migrations: %v\n", err)
	}

	return db
}

func Redis() RedisClient {
	redisClientOnce.Do(func() {
		hostValue := strings.Split(os.Getenv("REDIS_HOST"), ",")
		var hosts []string
		for _, v := range hostValue {
			v = strings.TrimSpace(v)
			if v != "" {
				hosts = append(hosts, v)
			}
		}
		useSentinel := false
		if len(hosts) > 1 {
			useSentinel = true
		}

		timeout, _ := strconv.Atoi(os.Getenv("REDIS_POOL_TIMEOUT"))
		port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
		database, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
		pool_size, _ := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
		min_idle_conn, _ := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE_CONN"))

		redisConn := NewRedisConnection(
			useSentinel,
			strings.TrimSpace(os.Getenv("REDIS_MASTER")),
			hosts,
			port,
			os.Getenv("REDIS_PASSWORD"),
			database,
			pool_size,
			time.Duration(timeout)*time.Second,
			min_idle_conn,
		)

		redisClient = NewRedis(redisConn)
	})

	return redisClient
}
