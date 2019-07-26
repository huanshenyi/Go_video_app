package defs

//reqeusts
type UserCredential struct {
   Username  string `json:"user_name"`  //タグの変形 "user_name": "xxx"
   Pwd       string `json:"pwd"`             //"pwd": "xxx"
}

// Data model
type VideoInfo struct {
   Id           string
   AuthorId     int
   Name         string
   DisplayCtime string
}

type Comment struct {
   Id string
   VideoId string
   Author string
   Content string
}

type SimpleSession struct {
   Username string //login name
   TTL int64  //賞味期限
}