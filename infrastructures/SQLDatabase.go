package infrastructures

import (
	"database/sql"
	"sync"
	"fmt"
	"time"

	config "github.com/spf13/viper"
)

type SQLConnection struct{}

type alias struct {
	DataSource string
	DB         *sql.DB
	Active     bool
}

type dbCache struct {
	mux   sync.RWMutex
	cache map[string]*alias
}

var (
	databaseCache = &dbCache{cache: make(map[string]*alias)}
)

func (sql *SQLConnection) InitDB(aliasName string) string {
	host := config.GetString(aliasName + ".host")
	port := config.GetString(aliasName + ".port")
	user := config.GetString(aliasName + ".user")
	password := config.GetString(aliasName + ".password")
	dbName := config.GetString(aliasName + ".name")
	charset := config.GetString(aliasName + ".charset")

	address := fmt.Sprintf("tcp(%s:%s)", host, port)
	dataSource := fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=true&loc=Local", user, password, address, dbName, charset)

	return dataSource
}

func (sql *SQLConnection) RegisterDataBase(aliasName, dataSource string) (err error) {
	err = addAliasWithDB(aliasName, dataSource)
	if err != nil {
		err = fmt.Errorf("%v", err)
	}

	return err
}

func (sql *SQLConnection) GetDB(aliasName string) (*sql.DB, error) {
	var (
		err  error
		al   *alias
		isOk bool
	)

	al, isOk = databaseCache.get(aliasName)

	if false == isOk {
		// New register database //
		sql.InitDB(aliasName)
		db, err := sql.GetDB(aliasName)
		return db, err
	}

	if !al.Active { // Re-open connection
		db, err := openConnection(al.DataSource)
		if err != nil {
			newDatasource := sql.InitDB(aliasName)
			db, err := openConnection(newDatasource)
			if err != nil {
				return db, fmt.Errorf(fmt.Sprintf("Failed to open database, err : %v\n", err))
			}
		}
		al.DB = db
		al.Active = true // set connection active status = true

	} else { // Re-check connection status with ping function
		err := al.DB.Ping()
		if err != nil {
			db, err := openConnection(al.DataSource)
			if err != nil {
				return db, fmt.Errorf(fmt.Sprintf("Failed to open database, err : %v\n", err))
			}
		}
	}

	return al.DB, err

}

func openConnection(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSource)
	db.SetConnMaxLifetime(config.GetDuration("database.max_life_time") * time.Second)
	db.SetMaxIdleConns(config.GetInt("database.max_idle_connection"))
	db.SetMaxOpenConns(config.GetInt("database.max_open_connection"))
	return db, err
}

func addAliasWithDB(aliasName, dataSource string) error {
	al := new(alias)
	al.DataSource = dataSource
	al.Active = false

	// Add connection into database cache
	if !databaseCache.add(aliasName, al) {
		return fmt.Errorf("DataBase alias name `%s` already registered, cannot reuse", aliasName)
	}

	return nil
}

func (ac *dbCache) add(name string, al *alias) (added bool) {
	ac.mux.Lock()
	defer ac.mux.Unlock()
	if _, ok := ac.cache[name]; !ok {
		ac.cache[name] = al
		added = true
	}
	return
}

func (ac *dbCache) get(name string) (al *alias, ok bool) {
	ac.mux.RLock()
	defer ac.mux.RUnlock()
	al, ok = ac.cache[name]
	return
}
