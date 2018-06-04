package dao

import (
  "log"

  . "github.com/clcollins/goRestApp/models"

  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type GatosDAO struct {
  Server string
  Database string
}

const COLLECTION = "gatos"

var db *mgo.Database

func (conn *GatosDAO) Connect() {
  session, err := mgo.Dial(conn.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(conn.Database)
}

func (conn *GatosDAO) FindAll() ([]Gato, error) {
  var gatos []Gato
  err := db.C(COLLECTION).Find(bson.M{}).All(&gatos)
  return gatos, err
}

func (conn *GatosDAO) Insert(gato Gato) error {
  err := db.C(COLLECTION).Insert(&gato)
  return err
}

func (conn *GatosDAO) FindById(id string) (Gato, error) {
  var gato Gato
  err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&gato)
  return gato, err
}
