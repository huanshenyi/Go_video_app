# APi設計

### user

|id|説明|URL:|Method|StatusCode|
|---|---|---|---|---|
|1|新規ユーザー|/user|POST|201,400,500|
|2|ユーザーログイン|/user/:username|POST|200,400,500|
|3|ユーザーのデーターを取得|/user/:username|GET|200,400,401,403,500|
|4|ユーザー削除|/user/:username|DELETE|204,400,401,403,500|

### video

|id|説明|URL:|Method|StatusCode|
|---|---|---|---|---|
|1|List_all_videos|/user/:username/videos|GET|200,400,500|
|2|Get_one_video|/user/:username/videos/:vid-id|GET|200,400,500|
|3|Delete_one_video|/user/:username/videos/:vid-id|DELETE|204,400,401,403,500|

### comment
|id|説明 |URL|Method|StatusCode|
|---|---|---|---|---|
|1|Show_comments|/videos/:vid-id/comments|GET|200,400,500|
|2|Show_comments|/videos/:vid-id/comments|POST|201,400,500|
|3|Delete_a_comment|/videos/:vid-id/comment/:comment-id|DELETE|204,400,401,403,500|


### リクエストデーターフロー

```cassandraql

handler->validation{1.request,2.user}->business logic->reponse.

```
 - data model
 - error handling
 
 ### DB設計
  
 #### users
 
 |カラム名|タイプ|||
 |---|---|---|---|
 |id|UNSIGNED INT|PRIMARY KEY|AUTO_INCREMENT|
 |login_name|VARCHAR(64)|UNIQUE KEY|
 |pwd|TEXT|||
 
 #### video_info

  |カラム名|タイプ|||
  |---|---|---|---|
  |id|VARCHAR(64)|PRIMARY KEY|NOT NULL|
  |author_id|UNSIGNED INT||
  |name|TEXT|||
  |display_ctime|TEXT|
  |create_time|DATETIME|
  

 - idをvarcharにするのはオーバー問題を防ぐ為に
 - author_idは外部キー、でもテーブル関連は敢えてcodeで実現
 
 #### comment
 
|カラム名|タイプ|||
  |---|---|---|---|
  |id|VARCHAR(64)|PRIMARY KEY|NOT NULL|
  |video_id|VARCHAR(64)||
  |author_id|UNSIGNED INT|||
  |content|TEXT|
  |time|DATETIME|
  
#### sessions

|カラム名|タイプ|||||
|---|---|---|---|---|---|
|session_id|TINYTEXT|PRIMARY KEY|NOT NULL|
|TTL |TINYTEXT|
|login_name|VARCHAR(64)|  

- TTL セッション切れる時間

