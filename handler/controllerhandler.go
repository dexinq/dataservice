package handler

import (
	"context"

	proto1 "github.com/dexinq/utils/proto/controller"
	proto2 "github.com/dexinq/utils/proto/global"
	"github.com/micro/go-micro/util/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ControllerHandler struct {
	database   string
	collection string
}

func NewControllerHandler() *ControllerHandler {
	return &ControllerHandler{}
}

func (ch *ControllerHandler) GetController(ctx context.Context, req *proto2.Empty, rsp *proto1.ControllerList) error {

	session, err := getSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	db := session.DB(ch.database)
	collection := db.C(ch.collection)
	res := make([]*proto1.ControllerObj, 0)
	err = collection.Find(nil).All(&res)

	rsp.Co = res

	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil

}

func (ch *ControllerHandler) AddController(ctx context.Context, req *proto1.ControllerObj, rsp *proto2.Response) error {
	session, err := getSession()
	defer session.Close()

	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	db := session.DB(ch.database)
	collection := db.C(ch.collection)

	err = collection.Insert(req)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}

	return nil
}

func (ch *ControllerHandler) UpdateController(ctx context.Context, req *proto1.ControllerObj, rsp *proto2.Response) error {
	session, err := getSession()
	defer session.Close()

	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	db := session.DB(ch.database)
	collection := db.C(ch.collection)

	query := bson.M{"name": req.Name}
	_, err = collection.Upsert(query, req)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func getSession() (*mgo.Session, error) {
	url := "get from config"
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return session, err
}
