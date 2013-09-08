package main

import (
    "flag"
    "github.com/skydb/sky.go"
    "github.com/skydb/skygen/core"
    "github.com/skydb/skygen/parser"
    "math/rand"
    "log"
    "os"
)

var host string
var port uint
var tableName string
var iterations int
var overwrite bool
var verbose bool

func init() {
    flag.StringVar(&host, "host", "localhost", "")
    flag.StringVar(&host, "h", "localhost", "")
    flag.UintVar(&port, "port", 8585, "")
    flag.UintVar(&port, "p", 8585, "")
    flag.StringVar(&tableName, "table", "", "")
    flag.StringVar(&tableName, "t", "", "")
    flag.BoolVar(&overwrite, "overwrite", false, "overwrite existing table")
    flag.IntVar(&iterations, "i", 1, "the number of iterations")
}

func main() {
    flag.Parse()

    // Load script and setup client.
    script := load()
    _, table := setup()

    // Sync schema if provided.
    schema := script.Schema()
    if schema != nil {
        if err := schema.Sync(table); err != nil {
            log.Fatalf("Schema error: %s\n", err.Error())
        }
    }

    // Generate an object for each iteration of the script.
    for i := 0; i < iterations; i++ {
        objectId := genid()
        if err := script.Generate(table, objectId); err != nil {
            log.Fatalf("Generation error: %s\n", err.Error())
        }
    }
}

// Reads a script from a file argument.
func load() *core.Script {
    if flag.NArg() == 0 {
        log.Fatalln("Script filename required")
    }
    file, err := os.Open(flag.Arg(0))
    if err != nil {
        log.Fatalf("Script load error: %s", err.Error())
    }
    script, err := parser.New().Parse(file)
    if err != nil {
        log.Fatalf("Script parse error: %s", err.Error())
    }
    return script
}

// Create a Sky client and grab a reference to the table.
func setup() (*sky.Client, *sky.Table) {
    client := sky.NewClient(host)
    client.Port = port
    if !client.Ping() {
        log.Fatal("Server is not running: %s:%d", host, port)
    }

    // Retrieve the table.
    table, _ := client.GetTable(tableName)
    if table != nil && overwrite {
        if err := client.DeleteTable(table); err != nil {
            log.Fatalf("Unable to overwrite existing table '%v': %v", tableName, err)
        }
        table = nil
    }
    if table == nil {
        table = sky.NewTable(tableName, nil)
        if err := client.CreateTable(table); err != nil {
            log.Fatalf("Unable to create table '%v': %v", tableName, err)
        }
    }

    return client, table
}

// Generates an 8 character random object identifier string.
func genid() string {
    length := 8
    b := make([]byte, length)
    for i := 0; i < length; i++ {
        b[i] = byte(rand.Intn(122-97) + 97)
    }
    return string(b)
}
