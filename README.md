# Flight Seat Map Response (Bookcabin Backend Take Home Hiring Test)

---

## 📂 Project Overview

This project is a **backend + frontend coding test** for the Bookcabin OTA booking system.

**Tech Stack:**

- **Backend:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL
- **Frontend:** React (Vite)
- **ORM/SQL:** Custom repositories (no GORM)

---

## 🚀 How to Run Backend

### 1️⃣ Install Dependencies

```
go mod tidy
```

### 2️⃣ Configure Environment Variables

Create a .env file in the project root:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=alif123
DB_NAME=bookcabin_flight_seat_app_db
```

### 3️⃣ Run Migrations

Using golang-migrate:

```
migrate -path db/migrations -database "postgres://postgres:alif123@localhost:5432/bookcabin_flight_seat_app_db?sslmode=disable" up
```

### 4️⃣ Start Backend Server

```
make run
```

The backend will start at http://localhost:8080.

#### 🧩 API Endpoint

```
GET /seat-map/seats-itinerary-part/:id
```

Example:

```
GET /seat-map/seats-itinerary-part/1
```

Response:

Returns JSON containing the full seat map structure.

## 🖥️ How to Run Frontend

### 1️⃣ Install Dependencies

```
make web-install
```

### 2️⃣ Start Development Server

```
make web
```

#### 🎨 Frontend Features

Seat grid with dynamic seat status (available/unavailable)

Tooltip displaying seat details (price, code, refund indicator)

Passenger details panel

Segment details panel (origin, destination, departure/arrival, terminals)

Responsive 3-column layout

#### 💾 How to Get Database

Provided the database in `db/seeds` folder

#### 🛠️ Notes

The backend uses raw SQL queries for better control.

The frontend is developed with Vite for fast builds.

Seat map data can be extended with more fields as needed.
