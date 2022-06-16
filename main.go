package main

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-common/currency"
	"github.com/WolffunGame/theta-shared-common/eventbus"
	"github.com/WolffunGame/theta-shared-common/thetalog"
	"github.com/ansel1/merry/v2"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"time"
)

func main() {
	var stats currency.Stats
	var client currency.Client
	client.Init(GetSession(), &stats, &currency.Settings{
		DeleteHistory: false,
		Check:         false,
	})

	t := currency.NewTransfer(currency.System, "sotuanhoang", 100000000, 13)

	err := client.MakeTransfer(&t)

	fmt.Println(merry.Details(err))
}

// S Singleton session
var session gocqlx.Session

func init() {
	//var cluster = gocql.NewCluster("scylladb-poc-client.infrastructure.svc.cluster.local")
	cluster := gocql.NewCluster("10.8.2.5", "10.8.8.5", "10.8.1.37")
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("GCE_ASIA_SOUTHEAST_1")
	cluster.Timeout = 5 * time.Second
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 5}
	cluster.Consistency = gocql.Quorum
	cluster.Timeout, _ = time.ParseDuration("30s")
	cluster.Keyspace = "thetancurrency"

	sess, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		thetalog.Err(err).Msg("Cannot connect to Scylla clusters")
	} else {
		err = sess.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS thetan WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 3}`)

		err := currency.BootstrapDatabase(sess.Session)

		fmt.Println(err)

		setSession(sess)
	}

	//var query = session.Query("SELECT * FROM system.clients", nil)

	//defer session.Close()
}

func GetSession() gocqlx.Session {
	return session
}

func setSession(s gocqlx.Session) {
	session = s
	fmt.Println("Scylla session has been set")
	//Publish 1 cái event thể hiện scylla đã init xong để những service nào depend on scylla có thể init (create keyspace, index, vv))
	eventbus.S.Publish("scylla", true)
}
