# Task Management System

## Overview
This project is a To-Do Task Management System designed to facilitate the creation, updating, and tracking of tasks. Each task is defined by a title, description, and status. Tasks that do not reach the "completed" status within one hour of creation are automatically invalidated, with their status updated to "critical".

## API Endpoints
The system provides five RESTful API endpoints to manage tasks:  

### Get All Tasks [GET]
**Description**: Retrieves a list of all tasks stored in the database.  
**API URL**: http://localhost:8080/tasks  
**Returns**: An array of task objects with the following structure:  

```
[
  {
    "Id": number,
    "Title": string,
    "Description": string,
    "Status": string,
    "Created_on": timestamp
  }
]
```

### Get Task [GET]
**Description**: Retrieves a specific task by its unique identifier.  
**API URL**: http://localhost:8080/tasks/{taskId}  
**Parameters**:  
  *taskId*: The unique identifier of the task to be retrieved.  
  *Returns*: A task object with the following structure:  
```
{
  "Id": number,
  "Title": string,
  "Description": string,
  "Status": string,
  "Created_on": timestamp
}
```

### Create Task [POST]
**Description**: Creates a new task in the database.  
**API URL**: http://localhost:8080/tasks  
**Request Body**:  

```
{
  "Title": string,  // Title of the task
  "Description": string,  // Description of the task
  "Status": string  // Status of the task
}
```
    
**Returns**: The created task object with the following structure:  

```
{
  "id": number,  // Unique identifier of the task
  "Title": string,
  "Description": string,
  "Status": string,
  "CreatedOn": timestamp  // Timestamp of task creation
}
```
    


### Update Task [PUT]
**Description**: Updates an existing task in the database.  
**API URL**: http://localhost:8080/tasks/{taskId}  
**Request Body**:  

```
{
  "Id": number,  // Unique identifier of the task
  "Title": string,  // Updated title of the task
  "Description": string,  // Updated description of the task
  "Status": string  // Updated status of the task
}
```
    
**Returns**: The updated task object with the following structure:  

```
{
  "id": number,
  "Title": string,
  "Description": string,
  "Status": string,
  "CreatedOn": timestamp
}
```


### Delete Task [DELETE]
**Description**: Deletes a specific task from the database.  
**API URL**: http://localhost:8080/tasks/{taskId}  
**Parameters**:  
  * *taskId* *: The unique identifier of the task to be deleted.  

## Task Invalidation Mechanism  
A cron job is implemented to automatically invalidate tasks that have not been marked as "completed" within one hour of their creation. Tasks that meet this criterion will have their status updated to "critical."  

### Cron Job Details
**Action**: Updates the status of tasks to "critical" if they are not in the "completed" status and were created more than one hour ago.  
**Frequency**: The cron job runs periodically to check and update task statuses as needed.  
