package tasks

import (
	"os"
	"path/filepath"

	"github.com/ONSdigital/dp-local-data/config"
	"github.com/fatih/color"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"gopkg.in/mgo.v2"
)

func DeleteCollections(cfg *config.Config) error {
	zebRoot := os.Getenv("zebedee_root")
	collectionsDir := filepath.Join(zebRoot, "/zebedee/collections")

	files, err := filepath.Glob(filepath.Join(collectionsDir, "*"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	color.Cyan("[clean] Deleting Zebedee collections: %+v", files)

	for _, f := range files {
		if err := os.RemoveAll(f); err != nil {
			return err
		}
	}
	color.Cyan("[clean] Zebedee collections deleted successfully")
	return nil
}

func DropMongo(cfg *config.Config) error {
	if len(cfg.MongoDBs) == 0 {
		return nil
	}

	sess, err := mgo.Dial(cfg.MongoURL)
	if err != nil {
		return err
	}
	defer sess.Close()

	color.Cyan("[clean] Dropping mongo CMD databases: %+v", cfg.MongoDBs)
	for _, db := range cfg.MongoDBs {
		err := sess.DB(db).DropDatabase()
		if err != nil {
			return err
		}
	}

	color.Cyan("[clean] Mongo CMD databases dropped successfully")
	return nil
}

func DropNeo4j(cfg *config.Config) error {
	color.Cyan("[clean] Dropping neo4j CMD data")
	pool, err := bolt.NewDriverPool(cfg.Neo4jURL, 1)
	if err != nil {
		return err
	}

	conn, err := pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	res, err := conn.ExecNeo("MATCH(n) DETACH DELETE n", nil)
	if err != nil {
		return err
	}

	deletions, _ := res.RowsAffected()
	color.Cyan("[clean] Neo4j nodes deleted: %d", deletions)
	return nil
}
