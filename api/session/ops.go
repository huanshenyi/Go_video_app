package session

import (
	"Go_video_app/api/dbops"
	"Go_video_app/api/defs"
	"Go_video_app/api/utils"
	"sync"
	"time"
)


var sessionMap *sync.Map //sync.Map 要研究

func init()  {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deletExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DEleteSession(sid)
}

func LoadSessionsFromDB()  {
    r,err := dbops.RetrieveAllSessions()
    if err != nil{
		return
	}
    r.Range(func(k,v interface{}) bool{
    	ss := v.(*defs.SimpleSession)
    	sessionMap.Store(k,ss)
    	return true
	})
}

func GenerateNewSessionId(un string) string {
    id, _ := utils.NewUUID()
    ct := nowInMilli()
    ttl := ct + 30*60*1000//Severside session valid time:30 min
    ss := &defs.SimpleSession{Username:un,TTL:ttl}
    sessionMap.Store(id,ss)
    dbops.InserSession(id, ttl,un)
    return id
}

// ExpiredしてなければUserNameを返す,Expiredしていればstring=nil,bool=False
func IsSessionExpired(sid string) (string, bool) {
   ss,ok := sessionMap.Load(sid)
   if ok {
	   ct := nowInMilli()
	   if ss.(*defs.SimpleSession).TTL < ct {
	   	   //delete expired session
		   deletExpiredSession(sid)
	   	   return "",true
	   }
	   return ss.(*defs.SimpleSession).Username,false
   }
   return  "",true
}

