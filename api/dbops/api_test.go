package dbops

import "testing"

var tempvid string

func clearTables()  {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
     t.Run("Add",testAddUser)
     t.Run("Get",testGetUser)
     t.Run("Del",testDeleteUser)
     t.Run("Reget",testRegutUser)
}

func testAddUser(t *testing.T)  {
    err := AddUserCredential("avenssi", "123")
    if err != nil {
    	t.Errorf("Error of AddUser: %v",err)
	}
}

func testGetUser(t *testing.T)  {
   pwd, err := GetUserCredential("avenssi")
   if pwd != "123" || err !=nil{
	   t.Errorf("Error of GetUser")
   }

}

func testDeleteUser(t *testing.T) {
    err := DeleteUser("avenssi", "123")
    if err != nil{
    	t.Errorf("Error of DeleteUser: %v",err)
	}
}

func testRegutUser(t *testing.T)  {
     pwd ,err := GetUserCredential("avenssi")
     if err != nil{
     	t.Errorf("error of RegetUser: %v", err)
	 }
     if pwd != ""{
     	t.Error("Deleting user test failed")
	 }
}

func TestVideoWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo",testAddVideoInfo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DelVideo",testDeleteVideo)
	t.Run("RegetVideo",testRegeVideoInfo)
}

func testAddVideoInfo(t *testing.T)  {
	vi, err := AddNewVideo(1,"my-video")
	if err != nil{
		t.Errorf("Error of AddVideoInfo %v",err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T)  {
	_, err := GetVideoInfo(tempvid)
	if err != nil{
		t.Errorf("Error of DelereVideoInfo %v",err)
	}
}

func testDeleteVideo(t *testing.T)  {
	err := DeleteVideoInfo(tempvid)
	if err != nil{
		t.Errorf("Error of DeleteVideoInfo %v",err)
	}
}
func testRegeVideoInfo(t *testing.T)  {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegeVideoInfo %v",err)
	}
}