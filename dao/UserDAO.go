package dao

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"login_service/config"
	"login_service/model"
)

type UserDAO interface {
	FindUser(username string) (*model.UserModel, error)
	AddNewUser(user *model.UserModel) error
	SaveUser(user *model.UserModel) error
	UserNotFoundError(err error) bool
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//       use Mongo as DAO
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//UserDAOMongo structure
type UserDAOMongo struct {
	userDBName string
	db         *mongo.Database
}

var instance_mongo *UserDAOMongo

//GetUserDAOMongoinstance method for accessing singleton UserDAOMongo instance
func GetUserDAOMongoinstance() *UserDAOMongo {

	if instance_mongo == nil {
		instance_mongo = &UserDAOMongo{
			userDBName: "users",
		}
	}
	return instance_mongo
}

//FindUser finds a user with given username in the Database, will return error if user not found
func (uDao *UserDAOMongo) FindUser(username string) (*model.UserModel, error) {
	collection := uDao.db.Collection(uDao.userDBName)
	var user model.UserModel
	err := collection.FindOne(context.Background(), bson.D{{"username", username}}).Decode(&user)
	return &user, err
}

//AddNewUser adds a new user to the Database
func (uDao *UserDAOMongo) AddNewUser(user *model.UserModel) error {
	collection := uDao.db.Collection(uDao.userDBName)

	_, err := collection.InsertOne(context.Background(), user)
	return err
}

//SaveUser saves a user to the Databs
func (uDao *UserDAOMongo) SaveUser(user *model.UserModel) error {
	collection := uDao.db.Collection(uDao.userDBName)
	update := bson.M{
		"$set": user,
	}
	doc := collection.FindOneAndUpdate(context.Background(), bson.D{{"username", user.Username}}, update, nil)
	return doc.Err()
}

//UserNotFoundError convenience method for checking if given error is no user found user.
func (uDao *UserDAOMongo) UserNotFoundError(err error) bool {
	return err.Error() == "mongo: no documents in result"
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//       use MYSQL as DAO
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//UserDAOMysql structure
type UserDAOMysql struct {
	userDBName string
}

var instance_mysql *UserDAOMysql

//GetUserDAOMysqlinstance method for accessing singleton UserDAOMysql instance
func GetUserDAOMysqlInstance() *UserDAOMysql {

	if instance_mysql == nil {
		instance_mysql = &UserDAOMysql{
			userDBName: "users",
		}
	}
	return instance_mysql
}

//FindUser finds a user with given username in the Database, will return error if user not found
func (uDao *UserDAOMysql) FindUser(username string) (*model.UserModel, error) {

	var user model.UserModel

	return &user, nil
}

//AddNewUser adds a new user to the Database
func (uDao *UserDAOMysql) AddNewUser(user *model.UserModel) error {

	return nil
}

//SaveUser saves a user to the Databs
func (uDao *UserDAOMysql) SaveUser(user *model.UserModel) error {

	return nil
}

//UserNotFoundError convenience method for checking if given error is no user found user.
func (uDao *UserDAOMysql) UserNotFoundError(err error) bool {
	return err.Error() == "mysql: no documents in result"
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//  init
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//Called only once during package inititalization
func init() {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetInstance().DBServerURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	UserDAOMongo := GetUserDAOMongoinstance()
	UserDAOMongo.db = client.Database("LoginService")

}
