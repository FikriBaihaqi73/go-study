# Error Handling Examples

File ini menjelaskan berbagai contoh error handling yang telah ditambahkan ke aplikasi Go Study.

## 📋 Daftar Error Codes yang Tersedia

### 1. **400 Bad Request** 
Error ini terjadi ketika client mengirim data yang tidak valid.

**Contoh di User Handler:**
- Format JSON tidak valid
- Email tidak valid (tidak ada @ atau .)
- Nama kosong atau kurang dari 3 karakter
- ID bukan angka atau <= 0

**Test Endpoint:**
```bash
# Test error 400
curl http://localhost:8080/examples/400?invalid=invalid

# Test success
curl http://localhost:8080/examples/400

# Test di User Handler
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "AB", "email": "invalid-email"}'
```

**Response Error:**
```json
{
  "error": "Bad Request",
  "message": "Name must be at least 3 characters long",
  "code": 400
}
```

---

### 2. **401 Unauthorized**
Error ini terjadi ketika authentication diperlukan tetapi tidak disediakan atau token invalid.

**Test Endpoint:**
```bash
# Tanpa header Authorization
curl http://localhost:8080/examples/401

# Dengan token invalid
curl -H "Authorization: Bearer invalid-token" \
  http://localhost:8080/examples/401

# Dengan token valid
curl -H "Authorization: Bearer valid-token" \
  http://localhost:8080/examples/401
```

**Response Error:**
```json
{
  "error": "Unauthorized",
  "message": "Unauthorized: Missing authentication token",
  "code": 401
}
```

---

### 3. **403 Forbidden**
Error ini terjadi ketika user sudah terautentikasi tetapi tidak memiliki permission untuk mengakses resource.

**Test Endpoint:**
```bash
# User tanpa role admin
curl http://localhost:8080/examples/403?role=user

# User dengan role admin
curl http://localhost:8080/examples/403?role=admin
```

**Response Error:**
```json
{
  "error": "Forbidden",
  "message": "Forbidden: You don't have permission to access this resource. Admin role required.",
  "code": 403
}
```

---

### 4. **404 Not Found**
Error ini terjadi ketika resource yang diminta tidak ditemukan.

**Contoh di User Handler:**
- User dengan ID tertentu tidak ada

**Test Endpoint:**
```bash
# Resource tidak ditemukan
curl http://localhost:8080/examples/404?id=999

# Resource ditemukan
curl http://localhost:8080/examples/404?id=123

# Test di User Handler
curl http://localhost:8080/users/9999
```

**Response Error:**
```json
{
  "error": "Not Found",
  "message": "User not found",
  "code": 404
}
```

---

### 5. **409 Conflict**
Error ini terjadi ketika terdapat konflik dengan state saat ini (misalnya duplicate data).

**Contoh di User Handler:**
- Email sudah terdaftar

**Test Endpoint:**
```bash
# Email sudah ada
curl http://localhost:8080/examples/409?email=john@example.com

# Email available
curl http://localhost:8080/examples/409?email=new@example.com

# Test di User Handler
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "existing@email.com"}'
```

**Response Error:**
```json
{
  "error": "Conflict",
  "message": "Email already exists",
  "code": 409
}
```

---

### 6. **422 Unprocessable Entity**
Error ini terjadi ketika server memahami request tetapi tidak bisa memprosesnya karena business logic validation.

**Test Endpoint:**
```bash
# Age kurang dari 18
curl http://localhost:8080/examples/422?age=15

# Age valid
curl http://localhost:8080/examples/422?age=25
```

**Response Error:**
```json
{
  "error": "Unprocessable Entity",
  "message": "Unprocessable Entity: Age must be 18 or older",
  "code": 422
}
```

---

### 7. **429 Too Many Requests**
Error ini terjadi ketika rate limit terlampaui.

**Test Endpoint:**
```bash
# Melebihi rate limit
curl http://localhost:8080/examples/429?requests=15

# Dalam batas normal
curl http://localhost:8080/examples/429?requests=5
```

**Response Error:**
```json
{
  "error": "Too Many Requests",
  "message": "Too Many Requests: Rate limit exceeded. Please try again later.",
  "code": 429
}
```

---

### 8. **500 Internal Server Error**
Error ini terjadi ketika ada error yang tidak terduga di server.

**Contoh di User Handler:**
- Error saat encoding JSON response

**Test Endpoint:**
```bash
# Trigger error 500
curl http://localhost:8080/examples/500?trigger=error

# Normal response
curl http://localhost:8080/examples/500
```

**Response Error:**
```json
{
  "error": "Internal Server Error",
  "message": "Internal Server Error: An unexpected error occurred. Please try again later.",
  "code": 500
}
```

---

### 9. **503 Service Unavailable**
Error ini terjadi ketika service sedang maintenance atau temporarily unavailable.

**Test Endpoint:**
```bash
# Service maintenance
curl http://localhost:8080/examples/503?maintenance=true

# Service available
curl http://localhost:8080/examples/503
```

**Response Error:**
```json
{
  "error": "Service Unavailable",
  "message": "Service Unavailable: System is under maintenance. Please try again later.",
  "code": 503
}
```

---

## 🔍 Melihat Daftar Semua Error Examples

```bash
curl http://localhost:8080/examples
```

Response:
```json
{
  "available_examples": [
    {
      "code": "400",
      "name": "Bad Request",
      "endpoint": "/examples/400?invalid=invalid",
      "description": "Client sent invalid data"
    },
    ...
  ],
  "message": "Test these endpoints to see different error responses"
}
```

---

## 📝 Struktur Error Response

Semua error response mengikuti struktur yang konsisten:

```go
type ErrorResponse struct {
    Error   string `json:"error"`    // HTTP Status Text
    Message string `json:"message"`  // Detail error message
    Code    int    `json:"code"`     // HTTP Status Code
}
```

---

## 🚀 Cara Testing

1. **Jalankan Server:**
   ```bash
   go run main.go
   ```

2. **Test dengan cURL atau Postman:**
   - Gunakan endpoint error examples untuk melihat berbagai contoh error
   - Test CRUD operations di `/users` endpoint dengan data invalid

3. **Lihat Response:**
   - Semua error response akan dalam format JSON
   - HTTP Status Code akan sesuai dengan jenis error

---

## 💡 Best Practices

1. **Selalu gunakan status code yang tepat** untuk setiap jenis error
2. **Berikan pesan error yang jelas dan informatif**
3. **Gunakan struktur error response yang konsisten**
4. **Validasi input data** sebelum processing
5. **Handle error dengan graceful** - jangan biarkan aplikasi crash

---

## 📚 Referensi HTTP Status Codes

- **2xx Success**: Request berhasil
- **4xx Client Error**: Error dari sisi client
- **5xx Server Error**: Error dari sisi server

Dokumentasi lengkap: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
