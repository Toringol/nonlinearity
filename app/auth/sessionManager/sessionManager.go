package sessionManager

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Toringol/nonlinearity/tools"
	"github.com/spf13/viper"

	"github.com/Toringol/nonlinearity/app/model"
	"github.com/gomodule/redigo/redis"
)

// SessionManager - manager of sessions using redis
type SessionManager struct {
	redisConn redis.Conn
}

// NewSessionManager - create new session manager
func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

// Create - create new session in redisDB
func (sm *SessionManager) Create(in *model.Session) (*model.SessionID, error) {
	id := model.SessionID{
		ID: tools.RandStringBytesMaskImprSrcSB(viper.GetInt("sessionKeylen")),
	}

	dataSerialized, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	mkey := "sessions:" + id.ID

	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 86400))
	if err != nil {
		return nil, err
	}

	if result != "OK" {
		return nil, fmt.Errorf("result not OK")
	}

	return &id, nil
}

// Check - checking session in redisDB
func (sm *SessionManager) Check(in *model.SessionID) (*model.Session, error) {
	mkey := "sessions:" + in.ID

	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		log.Println("cant get data:", err)
		return nil, err
	}

	sess := &model.Session{}

	err = json.Unmarshal(data, sess)
	if err != nil {
		log.Println("cant unpack session data:", err)
		return nil, err
	}

	return sess, nil
}

// Delete - delete record in redisDB
func (sm *SessionManager) Delete(in *model.SessionID) error {
	mkey := "sessions:" + in.ID
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		log.Println("redis error:", err)
	}
	return nil
}
