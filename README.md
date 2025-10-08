# Stocky Golang Assignment

This repository contains an example Golang implementation for the **Stocky assignment**.

---

## Project Overview
Stocky is a hypothetical platform where users earn shares of Indian stocks (e.g., Reliance, TCS, Infosys) as rewards for onboarding, referrals, or trading milestones. The system tracks rewards, maintains a ledger for stock purchases, and fetches stock prices to compute the INR value of user holdings.

---

## Requirements
- Go 1.20+
- PostgreSQL database named `assignment`
- Optional: Postman for testing APIs

---

## Setup Instructions

1. **Clone the repository:**
```bash
git clone <your-repo-url>
cd <repo-folder>
```

2. **Create the PostgreSQL database:**
```bash
psql -U postgres
CREATE DATABASE assignment;
```

3. **Set environment variables:**
Create a `.env` file (or export variables) with:
```env
DATABASE_URL=postgres://username:password@localhost:5432/assignment?sslmode=disable
PORT=8080
```

4. **Run the application:**
```bash
go run main.go
```
Server should start on `http://localhost:8080`.

5. **Run migrations:**
Migrations are run automatically on server startup.

---

## API Endpoints

### 1. `POST /reward`
Record a stock reward for a user.

**Request:**
```json
{
  "eventId": "unique-event-uuid",
  "userId": 12345,
  "symbol": "RELIANCE",
  "quantity": "1.000000",
  "unitPriceInr": "2500.00",
  "feesInr": "12.50",
  "timestamp": "2025-10-08T12:00:00Z"
}
```

**Response:**
```json
{ "status": "ok", "rewardId": 123 }
```
If duplicate:
```json
{"status": "duplicate", "eventId": "..."}
```

### 2. `GET /today-stocks/{userId}`
Returns all reward events for today for the given user.

### 3. `GET /historical-inr/{userId}`
Returns INR value of user’s stock rewards for all past days (up to yesterday).

### 4. `GET /stats/{userId}`
Returns:
- Total shares rewarded today grouped by stock symbol.
- Current INR value of the user’s portfolio.
- Portfolio breakdown by symbol.

### 5. `GET /portfolio/{userId}`
Returns holdings per stock symbol with current INR value.

---

## Database Schema
Key tables:
- `users`: stores user info.
- `reward_events`: immutable record of stock rewards.
- `ledger_entries`: double-entry ledger for assets, cash, and fees.
- `holdings`: user holdings per stock.
- `prices_hourly`: hourly stock prices.
- `prices_daily`: daily closing prices.

---

## Edge Cases & Notes
- **Duplicate rewards / replay attacks:** Prevented via `event_id` unique constraint.
- **Stock splits/mergers/delisting:** In production, handle via a `corporate_actions` table.
- **Rounding errors:** Quantities stored as `NUMERIC(18,6)`, INR as `NUMERIC(18,4)`.
- **Price API downtime:** Use last known price and mark data as stale.
- **Adjustments/refunds:** Use negative reward events or adjustments table with compensating ledger entries.

---

## Scaling & Best Practices
- Shard `reward_events` by user or time for high-volume writes.
- Use a dedicated service for price updates, publish via Kafka.
- Cache read-heavy endpoints with Redis.
- Keep `ledger_entries` and `reward_events` immutable for audit.

---

## Postman Collection
Attach the included `postman_collection.json` file in the repository root for testing all endpoints.

---

## Collaborator
Add `ashutosh@021.trade` as a collaborator in your private GitHub repository.

---

## License
For internal academic purposes. Do not distribute publicly.
