package database

type Database struct {
}

func New(host, user, password, database string) (db *Database, err error) {
	return db, err
}

func (db *Database) Close() {

}
