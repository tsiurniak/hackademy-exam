
# Hachademy exam project

To start you need:

    1. Change directory to ./backend
    2. Run "go run ./..." to start server
    3. Then change directory to ./frontend
    4. Run "npm install"
    5. Then run "npm run dev"

Enjoy🤝


## API Reference

#### Sign up

```http
  POST /user/signup
```

Creates user in-memory

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. Your email|
| `password` | `string` | **Required**. Your password|

Returns "registered" if your data is valid

#### Sign in

```http
  POST /user/signin
```

Returns JWT access token

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email` | `string` | **Required**. Your email|
| `password` | `string` | **Required**. Your password|

Returns your JWT Auth token if your data is valid


#### Create TODO list

```http
  POST /todo/lists
```

Creates a new TODO list

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | **Required**. List name|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Creates a new TODO list and return its body if your data is valid

#### Update TODO list

```http
  PUT /todo/lists/{list_id}
```

Update an available TODO list

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` | **Required**. List id|
| `name` | `string` |New list name|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Update an available TODO list and return its body if your data is valid

#### Delete TODO list

```http
  DELETE /todo/lists/{list_id}
```

Delete an available TODO list

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` | **Required**. List id|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Delete an available TODO list

#### Get TODO lists

```http
  GET /todo/lists
```

Returns all the lists for the given user

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Returns all the lists bodies for the given user

#### Create task

```http
  POST /todo/lists/{list_id}/tasks
```

Creates a task in the specified list.

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` | **Required**. List id|
| `task_name` | `string` | **Required**. Task name|
| `description` | `string` | **Required**. Task description|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Creates a new task and return its body if your data is valid

#### Create task

```http
  PUT /todo/lists/{list_id}/tasks/{task_id}
```

Updates a task in the specified list.

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` |**Required**. List id|
| `task_id` | `int` |**Required**. Task id|
| `task_name` | `string` |Task name|
| `description` | `string` |Task description|
| `status` | `string` |Task status|
| `update_name` | `string` | **Required** if you want to change task name|
| `update_description` | `string` | **Required** if you want to change task description|
| `update_status` | `string` | **Required** if you want to change task status|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Update an available task and return its body if your data is valid

#### Delete task

```http
  DELETE /todo/lists/{list_id}/tasks/{task_id}
```

Deletes a task in the specified list.

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` | **Required**. List id|
| `task_id` | `int` | **Required**. Task id|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Delete an available task

#### Get tasks

```http
  GET /todo/lists/{list_id}/tasks/
```

Returns all the tasks for the given user in the specified list.

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `list_id` | `int` | **Required**. List id|

| Header | Value     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `"Authotization"` | `"Bearer {yourToken}"` |**Required**. Your JWT token|

Returns all the tasks bodies for the given user in the specified list.