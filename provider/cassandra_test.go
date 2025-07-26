// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copyright 2024 FeatureForm Inc.
//

package provider

import (
	"testing"

	pc "github.com/featureform/provider/provider_config"
	pt "github.com/featureform/provider/provider_type"
)

func TestOnlineStoreCassandra(t *testing.T) {
	//if testing.Short() {
	//	t.Skip("skipping integration tests")
	//}
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	t.Logf("could not open .env file... Checking environment: %s", err)
	//}
	//cassandraUsername, ok := os.LookupEnv("CASSANDRA_USER")
	//if !ok {
	//	t.Fatalf("missing CASSANDRA_USER variable")
	//}
	//cassandraPassword, ok := os.LookupEnv("CASSANDRA_PASSWORD")
	//if !ok {
	//	t.Fatalf("missing CASSANDRA_PASSWORD variable")
	//}
	cassandraAddr := "localhost:9042"
	cassandraConfig := &pc.CassandraConfig{
		//Keyspace:    "f",
		Addr:        cassandraAddr,
		Username:    "",
		Consistency: "ONE",
		Password:    "",
		Replication: 3,
	}
	//featureform__fb69cfe644eb42339867c3e50a015852__2dad735de78e444ba8bc930cd502df7b
	//my_very_long_table_name_exceeding_48_characters
	//CREATE TABLE featureform__fb69cfe644eb42339867c3e50a015852__2dad735de78e444ba8bc930cd502df7b (
	//	id UUID PRIMARY KEY,
	//	name TEXT
	//);
	store, err := GetOnlineStore(pt.CassandraOnline, cassandraConfig.Serialized())
	if err != nil {
		t.Fatalf("could not initialize store: %s\n", err)
	}

	test := OnlineStoreTest{
		t:     t,
		store: store,
	}
	test.Run()
}
