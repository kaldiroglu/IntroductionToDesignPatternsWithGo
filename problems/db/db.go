// Package db is a faithful port of the Java db example: it demonstrates the
// "telescoping constructor" smell (plus post-construction setters) that the Builder
// pattern is meant to solve. Go has no constructor overloading, so NewDatabaseConnection
// supplies the defaults and callers set exported fields afterwards - exactly as the demo does.
package db

import "fmt"

type DatabaseConnection struct {
	Host              string
	Port              int
	Database          string
	Username          string
	Password          string
	UseSSL            bool
	ConnectionTimeout int
	ReadTimeout       int
	Charset           string
	AutoReconnect     bool
	MaxRetries        int
	ConnectionPool    string
	EnableLogging     bool
	LogLevel          string
	Timezone          string
}

func NewDatabaseConnection() *DatabaseConnection {
	// Set reasonable defaults
	return &DatabaseConnection{
		Port:              3306,
		UseSSL:            true,
		ConnectionTimeout: 30000,
		ReadTimeout:       60000,
		Charset:           "UTF-8",
		AutoReconnect:     true,
		MaxRetries:        3,
		EnableLogging:     false,
		LogLevel:          "INFO",
	}
}

func (c *DatabaseConnection) String() string {
	return fmt.Sprintf(
		"DatabaseConnection[host=%s, port=%d, database=%s, useSSL=%t, "+
			"connectionTimeout=%d, readTimeout=%d, charset=%s, autoReconnect=%t, "+
			"maxRetries=%d, connectionPool=%s, enableLogging=%t, logLevel=%s, timezone=%s]",
		c.Host, c.Port, c.Database, c.UseSSL, c.ConnectionTimeout, c.ReadTimeout,
		c.Charset, c.AutoReconnect, c.MaxRetries, c.ConnectionPool, c.EnableLogging,
		c.LogLevel, c.Timezone)
}

// Run mirrors the Java db/Main: four connection "profiles" built through the default
// constructor + field assignment. The Java Main produced no output; this prints each.
func Run() {
	// Scenario 1: Basic local development connection
	devConn := NewDatabaseConnection()
	devConn.Host = "localhost"
	devConn.Database = "myapp_dev"
	devConn.Username = "dev_user"
	devConn.Password = "dev_pass"

	// Scenario 2: Production connection with custom security and timeouts
	prodConn := NewDatabaseConnection()
	prodConn.Host = "prod-db-cluster.company.com"
	prodConn.Port = 5432 // PostgreSQL
	prodConn.Database = "myapp_production"
	prodConn.Username = "prod_user"
	prodConn.Password = "complex_secure_password"
	prodConn.UseSSL = true
	prodConn.ConnectionTimeout = 10000
	prodConn.ReadTimeout = 120000
	prodConn.ConnectionPool = "HikariCP"
	prodConn.EnableLogging = true

	// Scenario 3: Testing connection with specific charset and no SSL
	testConn := NewDatabaseConnection()
	testConn.Host = "test-server"
	testConn.Database = "test_db"
	testConn.Username = "test_user"
	testConn.Password = "test_pass"
	testConn.UseSSL = false
	testConn.Charset = "UTF-8"
	testConn.AutoReconnect = false

	// Scenario 4: Analytics connection with custom timezone and logging
	analyticsConn := NewDatabaseConnection()
	analyticsConn.Host = "analytics-db"
	analyticsConn.Database = "warehouse"
	analyticsConn.Username = "analytics_user"
	analyticsConn.Password = "analytics_pass"
	analyticsConn.Timezone = "UTC"
	analyticsConn.EnableLogging = true
	analyticsConn.LogLevel = "DEBUG"
	analyticsConn.ReadTimeout = 300000 // 5 minutes for long queries

	fmt.Printf("Dev:       %s\n", devConn)
	fmt.Printf("Prod:      %s\n", prodConn)
	fmt.Printf("Test:      %s\n", testConn)
	fmt.Printf("Analytics: %s\n", analyticsConn)
}

// Why constructors would be problematic here:
//
//  1. COMBINATORIAL EXPLOSION: With 15 properties, we'd need potentially hundreds of
//     constructor overloads to cover meaningful combinations.
//  2. PARAMETER CONFUSION: a constructor with 8+ parameters becomes error-prone -
//     which timeout is which? Easy to mix up port and timeout values.
//  3. MEANINGLESS COMBINATIONS: not every combination makes sense.
//  4. MAINTENANCE NIGHTMARE: adding one new property means many new constructors.
//  5. UNCLEAR INTENT: post-construction initialization makes it clear which properties
//     are being customized for each use case.
//  6. FLEXIBILITY: easy to create different "profiles" without predefined constructors.
