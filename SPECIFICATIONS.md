# API Specifications
Todo api specifications

> Base URL : http://localhost:5000

> Auth : - (No Auth)

## Table of contents
- [Endpoints](#endpoints)
    - todos
      - [[GET] /api/todos](#get-apitodos)
      - [[POST] /api/todos](#post-apitodos)
      - [[PUT] /api/todos](#put-apitodostodoid)
      - [[GET] /api/todos/:todoId](#get-apitodostodoid)
      - [[DELETE] /api/todos/:todoId](#delete-apitodostodoid)
- [Error Response](#error-response)

## Endpoints

### [GET] /api/todos
#### request body:
`no request body`

#### response status:
`200 OK`

#### response body:
```js
{
    message: string,
    result: [
        {
            id: string,
            title: string,
            isDone: boolean
        },
        ...
    ]  
}
```

### [POST] /api/todos
#### request body:
```js
{
    title: string,
    isDone: boolean
}
```

#### response status:
`201 CREATED`

#### response body:
```js
{
    message: string,
    result: {
        id: string,
        title: string,
        isDone: boolean
    }
}
```

### [PUT] /api/todos/:todoId
#### request params:
- `todoId`

#### request body:
```js
{
    title: string,
    isDone: boolean
}
```

#### response status:
`200 OK`

#### response body:
```js
{
    message: string,
    result: {
        id: string,
        title: string,
        isDone: boolean
    }    
}
```

### [GET] /api/todos/:todoId
#### request params:
- `todoId`

#### request body:
`no request body`

#### response status:
`200 OK`

#### response body:
```js
{
    message: string,
    result: {
        id: string,
        title: string,
        isDone: boolean
    }    
}
```

### [DELETE] /api/todos/:todoId
#### request params:
- `todoId`

#### request body:
`no request body`

#### response status:
`200 OK`

#### response body:
```js
{
    message: string,
    result: null 
}
```

## Error response
### response status:
`any error status code depending on the error that occurred`

### response body:
```js
{
    message: string, // error message
    result: null 
}
```


