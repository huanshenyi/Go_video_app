package defs

//reqeusts
type UserCredential struct {
   Username string `json:"user_name"`  //タグの変形 "user_name": "xxx"
   Pwd string `json:"pwd"`             //"pwd": "xxx"
}
