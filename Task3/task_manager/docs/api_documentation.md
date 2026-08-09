# Task Management API

A simple **Task Management REST API** built with **Go** and the **Gin Web Framework**.

## Features

- Create a new task
- Get all tasks
- Get a task by ID
- Update a task
- Delete a task

---

## Technologies

- **Go**
- **Gin Web Framework**
- **MongoDB**
- **MongoDB Go Driver**
- **Postman**

The API uses **MongoDB for persistent data storage**, replacing the previous in-memory slice-based storage.

---

## Project Structure

```text
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
├── go.mod
└── go.sum
```

### Folder Responsibilities

#### `main.go`

The entry point of the application. It initializes the service, controller, router, and starts the HTTP server.

#### `models/`

Contains the data structures used by the application.

#### `data/`

Contains the task service and the business logic for interacting with MongoDB. It implements the CRUD operations using the MongoDB Go Driver.

#### `controllers/`

Handles HTTP requests and responses and communicates with the task service.

#### `router/`

Defines the API routes and connects them to the appropriate controller methods.

#### `docs/`

Contains additional API documentation.

---

## Task Model

Each task contains the following fields:

| Field         | Type     | Description                           |
| ------------- | -------- | ------------------------------------- |
| `id`          | ObjectID | Unique MongoDB identifier of the task |
| `title`       | string   | Title of the task                     |
| `description` | string   | Description of the task               |
| `dueDate`     | time     | Deadline for the task                 |
| `status`      | string   | Current status of the task            |

MongoDB automatically generates the `id` (`_id`) when a new task is created.

Example:

```json
{
  "id": "6a78e0aace3ad04c3ba7aa66",
  "title": "Learn Gin/go",
  "description": "Study Gin framework",
  "dueDate": "2026-08-10T00:00:00Z",
  "status": "pending"
}
```

When creating a task, the `id` does not need to be included in the request body because MongoDB generates it automatically.

---

## API Endpoints

Base URL:

```text
http://localhost:8080
```

## 1. Get All Tasks

```http
GET /tasks
```

Returns all tasks currently stored in memory.

### Example Request

```text
GET http://localhost:8080/tasks
```

### Success Response

**200 OK**

```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "description": "Learn Gin framework",
    "dueDate": "2026-08-15T00:00:00Z",
    "status": "pending"
  }
]
```

---

## 2. Get Task by ID

```http
GET /tasks/:id
```

Returns the details of a specific task.

### Example Request

```text
GET http://localhost:8080/tasks/1
```

### Success Response

**200 OK**

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Learn Gin framework",
  "dueDate": "2026-08-15T00:00:00Z",
  "status": "pending"
}
```

### Invalid ID

**400 Bad Request**

```json
{
  "error": "invalid task ID"
}
```

### Task Not Found

**404 Not Found**

```json
{
  "error": "task with this ID does not exist"
}
```

---

## 3. Create a Task

```http
POST /tasks
```

Creates a new task.

### Example Request

```text
POST http://localhost:8080/tasks
```

### Request Body

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Learn Gin framework",
  "dueDate": "2026-08-15T00:00:00Z",
  "status": "pending"
}
```

### Success Response

**201 Created**

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Learn Gin framework",
  "dueDate": "2026-08-15T00:00:00Z",
  "status": "pending"
}
```

### Invalid Request Body

**400 Bad Request**

```json
{
  "error": "invalid request body"
}
```

---

## 4. Update a Task

```http
PUT /tasks/:id
```

Updates an existing task.

### Example Request

```text
PUT http://localhost:8080/tasks/1
```

### Request Body

```json
{
  "id": 1,
  "title": "Learn Gin",
  "description": "Build a Task Management API",
  "dueDate": "2026-08-20T00:00:00Z",
  "status": "completed"
}
```

### Success Response

**200 OK**

```json
{
  "id": 1,
  "title": "Learn Gin",
  "description": "Build a Task Management API",
  "dueDate": "2026-08-20T00:00:00Z",
  "status": "completed"
}
```

### Invalid Task ID

**400 Bad Request**

```json
{
  "error": "invalid task ID"
}
```

### Task Not Found

**404 Not Found**

```json
{
  "error": "task with this ID does not exist"
}
```

---

## 5. Delete a Task

```http
DELETE /tasks/:id
```

Deletes a task from the in-memory store.

### Example Request

```text
DELETE http://localhost:8080/tasks/1
```

### Success Response

**204 No Content**

No response body is returned.

### Invalid Task ID

**400 Bad Request**

```json
{
  "error": "invalid task ID"
}
```

### Task Not Found

**404 Not Found**

```json
{
  "error": "task with this ID does not exist"
}
```

---

# HTTP Status Codes

The API uses the following HTTP status codes:

| Status            | Meaning                        |
| ----------------- | ------------------------------ |
| `200 OK`          | Request completed successfully |
| `201 Created`     | A new task was created         |
| `204 No Content`  | Task was successfully deleted  |
| `400 Bad Request` | Invalid request or task ID     |
| `404 Not Found`   | Requested task does not exist  |

---

# Architecture

The app follows a simple layered structure:

```text
Client
   │
   ▼
Router
   │
   ▼
Controller
   │
   ▼
Task Service
   │
   ▼
In-Memory Slice
```

### Router

Receives HTTP requests and directs them to the appropriate controller.

### Controller

Handles HTTP-specific concerns such as:

- Request parsing
- JSON binding
- HTTP status codes
- HTTP responses

### Service

Contains the task management logic:

- Add task
- Get tasks
- Get task by ID
- Update task
- Delete task

### Model

Defines the `Task` structure used throughout the application.

---
