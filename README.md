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


