// Copyright © 2018 Alex Kolbasov
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gometastore_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/jahnestacado/gometastore"
)

func ExampleOpen() {
	client, err := gometastore.Open("localhost", 9083, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.GetAllDatabases(context.Background()))
}

func ExampleMetastoreClient_GetAllDatabases() {
	client, err := gometastore.Open("localhost", 9083, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println(client.GetAllDatabases(context.Background()))
}

func TestOpenBadHost(t *testing.T) {
	t.Log("connecting to fake host")
	client, err := gometastore.Open("foobar", 1, nil)
	if err == nil {
		t.Error("connection to bad host succeeded")
	}
	if client != nil {
		t.Error("connecting to bad host returned valid client")
	}
}

func getClient(t *testing.T) (*gometastore.MetastoreClient, error) {
	host := os.Getenv("HMS_SERVER")
	port := os.Getenv("HMS_PORT")
	if port == "" {
		port = "9083"
	}
	if host == "" {
		host = "localhost"
	}
	portVal, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		t.Error("invalid port", portVal, err)
		return nil, err

	}
	t.Log("connecting to", host)
	client, err := gometastore.Open(host, int(portVal), nil)
	if err != nil {
		t.Error("failed connection to", host, err)
		return nil, err
	}
	return client, nil
}

func TestGetDatabases(t *testing.T) {
	client, err := getClient(t)
	if err != nil {
		return
	}
	defer client.Close()
	databases, err := client.GetAllDatabases(context.Background())
	if err != nil {
		t.Error("failed to get databases", err)
		return
	}
	if len(databases) == 0 {
		t.Error("no databases available")
		return
	}
}

func TestMetastoreClient_CreateDatabase(t *testing.T) {
	dbName := os.Getenv("HMS_TEST_DATABASE")
	owner := os.Getenv("HADOOP_USER_NAME")
	if dbName == "" {
		dbName = "hms_test_database"
	}
	t.Log("Testing creating database", dbName, "and owner", owner)
	client, err := getClient(t)
	if err != nil {
		return
	}
	defer client.Close()
	description := "test database"
	err = client.CreateDatabase(context.Background(), &gometastore.Database{Name: dbName, Description: description, Owner: owner})
	if err != nil {
		t.Error("failed to create database:", err)
		return
	}
	db, err := client.GetDatabase(context.Background(), dbName)
	if err != nil {
		t.Error("failed to get database:", err)
		return
	}
	if db.Name != dbName {
		t.Errorf("dbname %s is not equal %s", db.Name, dbName)
	}
	if description != db.Description {
		t.Errorf("description %s is not equal %s", db.Description, description)
	}
	if owner != db.Owner {
		t.Errorf("owner %s is not equal %s", db.Owner, owner)
	}
	err = client.DropDatabase(context.Background(), dbName, true, false)
	if err != nil {
		t.Error("failed to drop database", err)
		return
	}
}
