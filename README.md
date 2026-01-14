# 🌦️ CEP Weather API

API developed in **Go** that receives a **Brazilian ZIP code (CEP)**, identifies the city, and returns the **current temperature** in **Celsius, Fahrenheit, and Kelvin**.

The application is ready for **deployment on Google Cloud Run** and uses **API Key authentication**.

---

## 🚀 Technologies Used

- Go 1.24+
- Google Cloud Run
- Docker (multi-stage build)
- ViaCEP API
- WeatherAPI
- REST / HTTP

---

## 📌 Features

- ZIP code validation (8 digits)
- Location lookup via ViaCEP
- Temperature retrieval via WeatherAPI
- Temperature conversion:

  - Celsius
  - Fahrenheit
  - Kelvin

- API Key authentication
- Standardized HTTP responses
- Ready for local execution or Cloud Run deployment

---

## 🔐 Authentication

This API uses **API Key authentication** via **HTTP header**.

### Required Header

```http
x-api-key: YOUR_API_KEY
```

📌 **Requests without the API key or with an invalid key will be rejected.**

---

## 🌍 Endpoint

### 🔎 Get weather by ZIP code

```http
GET /weather?cep=ZIPCODE
```

### Parameters

| Name | Type   | Description               |
| ---- | ------ | ------------------------- |
| cep  | string | Valid ZIP code (8 digits) |

---

## 📥 Responses

### ✅ Success

**HTTP 200**

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

---

### ❌ Errors

| Code | Message              |
| ---- | -------------------- |
| 401  | API key missing      |
| 403  | Invalid API key      |
| 404  | can not find zipcode |
| 422  | invalid zipcode      |

---

## 🧪 Usage Examples

### 📌 cURL

```bash
curl -X GET "https://YOUR_ENDPOINT/weather?cep=01001000" \
  -H "x-api-key: YOUR_API_KEY"
```

---

### 📌 HTTP File (VS Code / IntelliJ)

File: `weather.http`

```http
### Get weather by ZIP code
GET https://YOUR_ENDPOINT/weather?cep=01001000
x-api-key: YOUR_API_KEY
```

---

## 🐳 Running Locally with Docker

### Build the image

```bash
docker build -t cep-weather .
```

### Run

```bash
docker run -p 8080:8080 \
  -e API_KEY=abc123 \
  -e WEATHER_API_KEY=YOUR_WEATHER_API_KEY \
  cep-weather
```

---

## ☁️ Deploy to Google Cloud Run

### Build & Push the image

```bash
gcloud auth configure-docker us-central1-docker.pkg.dev

docker build -t us-central1-docker.pkg.dev/YOUR_PROJECT/cep-weather/api .
docker push us-central1-docker.pkg.dev/YOUR_PROJECT/cep-weather/api
```

---

### Deploy

```bash
gcloud run deploy cep-weather \
  --image us-central1-docker.pkg.dev/YOUR_PROJECT/cep-weather/api \
  --region us-central1 \
  --platform managed \
  --allow-unauthenticated \
  --set-env-vars API_KEY=abc123,WEATHER_API_KEY=YOUR_KEY
```

---

## 🧪 Automated Tests

The tests cover:

- ZIP code validation
- API Key authentication
- Temperature conversion
- Error scenarios (422, 404, 401, 403)

### Run tests

```bash
go test ./...
```

---

## 📎 Notes

- The service runs on **port 8080**, as required by Cloud Run
- No sensitive keys are committed to the repository
