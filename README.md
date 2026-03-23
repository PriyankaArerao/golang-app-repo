# Coach Booking System

Backend service where:

* Coaches set availability
* Users view 30-min slots
* Users book appointments
* No double booking allowed

---

## Tech Stack

* Go (Echo framework)
* PostgreSQL

---

### 1. Set Availability

**POST** `/coaches/availability`

```json
{
  "coach_id": 1,
  "day": "Tuesday",
  "start_time": "09:00",
  "end_time": "14:00"
}
```

---

### 2. Get Available Slots

**GET** `/users/slots?coach_id=1&date=2025-10-28`

**Response**

```json
[
  "2025-10-28T09:00:00Z",
  "2025-10-28T09:30:00Z"
]
```

---

### 3. Book Slot

**POST** `/users/bookings`

```json
{
  "user_id": 101,
  "coach_id": 1,
  "datetime": "2025-10-28T09:30:00Z"
}
```

---

### 4. Get User Bookings

**GET** `/users/bookings?user_id=101`

---

## Setup

```bash
git clone https://github.com/PriyankaArerao/golang-app-repo.git
go mod tidy
go run main.go
```
## NOTE:
Tables are automatically created by GORM. No manual database setup is required.

## Swagger Docs Instructions
* Automatic Docs – echo-swagger generates documentation automatically from endpoint comments.
* Comment Format – Use the declarative comment format: https://github.com/swaggo/swag#declarative-comments-format
* View UI – While the application is running locally, open:
http://localhost:8080/swagger/index.html to see the interactive Swagger UI.