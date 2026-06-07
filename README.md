# LensLift 

I built this because I needed it.

Been going to the gym consistently lately and realized I was missing something — an app that actually handles everything in one place. Tracking what I eat, logging my workouts, monitoring my weight progress. Most apps out there either do one thing well or try to do everything and end up being bloated and annoying to use.

So I built LensLift. A gym companion app that covers the full picture: nutrition tracking powered by AI, workout logging, and body progress monitoring. The hero feature? You just take a photo of your food and the AI figures out the calories and macros for you. No more manually searching through databases for "nasi goreng ayam" and guessing the portion size.

---

## What it does

**AI Food Analyzer** — Take a photo of your meal, get instant calorie and macro breakdown. Powered by Claude Vision API on the backend, so your API key never touches the client.

**Nutrition Tracker** — Daily food log with calorie and protein progress tracking against your personal target. Targets are auto-calculated based on your body stats and goal (cut, bulk, or maintain) using the Harris-Benedict TDEE formula.

**Workout Templates** — Build your own training programs (Push Day, Pull Day, whatever works for you) with exercises, sets, and reps. Use them as blueprints for your actual sessions.

**Workout Sessions** — Start a session from a template, log your actual sets and weights in real time, and get a summary when you're done — total volume, duration, everything.

**Gym Schedule** — Assign your templates to days of the week. Monday is Push Day, Wednesday is Pull Day, you get the idea.

**Body Weight Log** — Track your weight daily and see your progress over time.

---

## Tech Stack

**Backend**
- Go + Gin (REST API)
- PostgreSQL (database)
- GORM (ORM)
- JWT (authentication)
- Claude Vision API (AI food analysis)

**Mobile**
- iOS (Swift + SwiftUI) — in progress

**Architecture**
The backend follows a layered architecture: Handler → Service → Repository → Database. JWT auth middleware protects all user endpoints. The Claude API is called server-side only — the API key is never exposed to the client.

---

## API Endpoints

```
Auth
POST   /auth/register
POST   /auth/login

Profile
GET    /api/profile
PUT    /api/profile

Workout Templates
POST   /api/workout-templates
GET    /api/workout-templates
GET    /api/workout-templates/:id
PUT    /api/workout-templates/:id
DELETE /api/workout-templates/:id

Gym Schedule
POST   /api/schedules
GET    /api/schedules
DELETE /api/schedules/:day

Workout Sessions
POST   /api/sessions/start
POST   /api/sessions/:id/log
POST   /api/sessions/:id/end
GET    /api/sessions
GET    /api/sessions/:id

Food Log
POST   /api/food-logs
POST   /api/food-logs/analyze
GET    /api/food-logs
DELETE /api/food-logs/:id

Body Weight
POST   /api/body-weights
GET    /api/body-weights
GET    /api/body-weights/latest
DELETE /api/body-weights/:id
```

---

## Running locally

**Prerequisites:** Go 1.22+, PostgreSQL

```bash
git clone https://github.com/nanda/lenslift-backend.git
cd lenslift-backend

cp .env.example .env
# fill in your DB credentials and API keys

go mod tidy
go run cmd/main.go
```

The server will start on port 8080. Database tables are created automatically on first run via GORM AutoMigrate.

---

## Environment Variables

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=lenslift_db
JWT_SECRET=your_jwt_secret
ANTHROPIC_API_KEY=your_claude_api_key
APP_PORT=8080
APP_ENV=development
```

---

Built by Nanda — Informatics Engineering student at Telkom University Bandung, currently learning iOS development and Go backend while building things I actually want to use.
