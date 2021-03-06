package users

import (
	"github.com/nightlegend/apigateway/core/module"
	"github.com/nightlegend/apigateway/core/utils"
	"github.com/nightlegend/apigateway/core/utils/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

/*
* Register
* * param [UserInfo]
* * return bool type.
 */
func Register(userInfo module.UserInfo) bool {
	session := db.Connectmon()
	defer session.Close()
	c := session.DB("test").C("userInfo")
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	err := c.Insert(userInfo)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

/*
* Login
* * param [userName, password]
* * return bool type.
 */
func Login(userName string, password string) int {
	/*
	* Get db connection
	 */
	session := db.Connectmon()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("userInfo")

	var userInfo module.UserInfo
	err := c.Find(bson.M{"username": userName}).One(&userInfo)
	if err != nil {
		log.Println(err)
		return 204
	}
	if password == utils.DeCryptedStr([]byte(userInfo.PASSWORD)) {
		return 200
	} else {
		return 201
	}

	return 205
}
