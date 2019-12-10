package db

import (
	"gopkg.in/mgo.v2"
	"time"
)

type Session struct {
	session *mgo.Session
	dbname  string
}

func NewSession(addr string, dbname string) (*Session, error) {
	session, err := mgo.Dial(addr)
	if err != nil {
		return nil, err
	}

	session.SetSyncTimeout(10 * time.Second)
	session.SetSocketTimeout(time.Minute)
	session.SetMode(mgo.Monotonic, true)

	clonedSession := session.Clone()

	return &Session{
		session: clonedSession,
		dbname:  dbname,
	}, err
}

func (s *Session) Collection(collection string) *mgo.Collection {
	return s.session.DB(s.dbname).C(collection)
}

func (s *Session) Database(name string) *mgo.Database {
	return s.session.DB(name)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
