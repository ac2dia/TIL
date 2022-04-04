# OpenAPI Guide

## Basic Structure

```yaml
openapi: 3.0.0
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
servers:
  - url: http://api.example.com/v1
    description: Optional server description, e.g. Main (production) server
  - url: http://staging-api.example.com
    description: Optional server description, e.g. Internal staging server for testing
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML.
      responses:
        '200': # status code
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string
```

- OpenAPI 형식은 YAML 또는 JSON 으로 정의할 수 있습니다.

### Metadata

```yaml
openapi: 3.0.0
```

- OpenAPI 버전으로 SemVer 표기법을 따릅니다.
- 현재 사용하는 버전은 3.0.0 ~ 3.0.3 으로 기능상으로는 동일합니다.

```yaml
info:
  title: Sample API
  description: Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.
  version: 0.1.9
```

- info 영역의 경우 API 정보를 가지고 있습니다.
  - title = 정의하는 API 이름
  - description = 정의하는 API 에 대한 정보
  - version = 정의하는 API 의 버전

### Servers

```yaml
servers:
  - url: http://api.example.com/v1
    description: Optional server description, e.g. Main (production) server
  - url: http://staging-api.example.com
    description: Optional server description, e.g. Internal staging server for testing
```

- servers 영역의 경우 API 서버 및 base URL 정보를 가지고 있습니다.
  - product, staging 과 같은 여러 서버 정보들을 작성할 수 있습니다.

### Paths

```yaml
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in CommonMark or HTML
      responses:
        '200':
          description: A JSON array of user names
          content:
            application/json:
              schema:
                type: array
                items:
                  type: string
```

- paths 영역의 경우 API 의 개별 endpoint 와 해당 endpoint 가 지원하는 HTTP Method 를 정의합니다.

### Parameters

```yaml
paths:
  /users/{userId}:
    get:
      summary: Returns a user by ID.
      parameters:
        - name: userId
          in: path
          required: true
          description: Parameter description in CommonMark or HTML.
          schema:
            type: integer
            format: int64
            minimum: 1
      responses:
        '200':
          description: OK
```

- parameters 의 경우 URL path (/users/{userId}), query string (/users?role=admin), headers (x-CustomHeader: Value), cookies (Cookie: debug=0) 과 같은 형태로 사용됩니다.

### Request Body

```yaml
paths:
  /users:
    post:
      summary: Creates a user.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                username:
                  type: string
      responses:
        '201':
          description: Created
```

- requestBody 키워드의 경우 body content 및 media type 을 정의하는데 사용될 수 있습니다.

### Responses

```yaml
paths:
  /users/{userId}:
    get:
      summary: Returns a user by ID.
      parameters:
        - name: userId
          in: path
          required: true
          description: The ID of the user to return.
          schema:
            type: integer
            format: int64
            minimum: 1
      responses:
        '200':
          description: A user object.
          content:
            application/json:
              schema:
                type: object
                properties:
                  id:
                    type: integer
                    format: int64
                    example: 4
                  name:
                    type: string
                    example: Jessica Smith
        '400':
          description: The specified user ID is invalid (not a number).
        '404':
          description: A user with the specified ID was not found.
        default:
          description: Unexpected error
```

- status code 로 상태 값을 표기하고, schema 를 사용하여 response body 를 표기합니다.
- 내부에 직접 정의하거나 $ref 키워드를 통해 schema 를 참조합니다.

### Input and Output Models

```yaml
{
  "id": 4,
  "name": "Arthur Dent"
}

components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: integer
          example: 4
        name:
          type: string
          example: Arthur Dent
      # Both properties are required
      required:
        - id
        - name
```

```yaml
paths:
  /users/{userId}:
    get:
      summary: Returns a user by ID.
      parameters:
        - in: path
          name: userId
          required: true
          schema:
            type: integer
            format: int64
            minimum: 1
      responses:
        '200':
          description: OK
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User' # <-------
  /users:
    post:
      summary: Creates a new user.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User' # <-------
      responses:
        '201':
          description: Created
```

- components/schemas 영역의 경우 API 에서 사용되는 데이터 구조에 대해 정의합니다.
- schema 가 필요한 경우 $ref 키워드를 통해 parameters, request bodies, response bodies 가 참조될 수 있습니다.

### Authentication

```yaml
components:
  securitySchemes:
    BasicAuth:
      type: http
      scheme: basic
security:
  - BasicAuth: []
```

- securitySchemes, security 키워드의 경우 API 에서 사용되는 인증 방법에 대해 정의합니다.
- 현재 지원하는 인증 방법의 경우 아래와 같습니다.
  - HTTP authentication: Basic, Bearer, and so on.
  - API key as a header or query parameter or in cookies
  - OAuth 2
  - OpenID Connect Discovery

# Reference

[1] OpenAPI Guide, https://swagger.io/docs/specification/about
