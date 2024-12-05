## Curl

## 1. Create a Movie (POST)
```bash
curl -X POST http://localhost:8080/v1/movies \
-H "Content-Type: application/json" \
-d '{
    "title": "An Amazing Movie",
    "genre": "Drama",
    "description": "An incredible story of perseverance.",
    "release_date": "2024-12-05T00:00:00Z"
}'
```

### 2. Get a Movie by ID (GET)


```bash
curl -X GET http://localhost:8080/v1/movies/1 \
-H "Content-Type: application/json"
```

### 3. Update a Movie (PATCH)

```bash
curl -X PATCH http://localhost:8080/v1/movies/1 \
-H "Content-Type: application/json" \
-d '{
    "description": "An updated description for the movie."
}'
```

### 4. Delete Movie (DELETE)

```bash
curl -X DELETE http://localhost:8080/v1/movies/1 \
-H "Content-Type: application/json"
```

### 5. Undelete a Movie (POST)

```bash
curl -X POST http://localhost:8080/v1/movies:undelete \
-H "Content-Type: application/json" \
-d '{
    "id": 1
}'
```