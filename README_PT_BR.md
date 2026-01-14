# 🌦️ CEP Weather API

API desenvolvida em **Go** que recebe um **CEP brasileiro**, identifica a cidade e retorna a **temperatura atual** em **Celsius, Fahrenheit e Kelvin**.

A aplicação está preparada para **deploy no Google Cloud Run** e utiliza autenticação via **API Key**.

---

## 🚀 Tecnologias Utilizadas

- Go 1.24+
- Google Cloud Run
- Docker (multi-stage build)
- ViaCEP API
- WeatherAPI
- REST / HTTP

---

## 📌 Funcionalidades

- Validação de CEP (8 dígitos)
- Busca de localização via ViaCEP
- Consulta de temperatura via WeatherAPI
- Conversão de temperatura:

  - Celsius
  - Fahrenheit
  - Kelvin

- Autenticação via API Key
- Respostas HTTP padronizadas
- Pronto para execução local ou Cloud Run

---

## 🔐 Autenticação

Esta API utiliza **API Key** enviada via **header HTTP**.

### Header obrigatório

```http
x-api-key: SUA_API_KEY
```

📌 **Requisições sem a chave ou com chave inválida serão rejeitadas.**

---

## 🌍 Endpoint

### 🔎 Obter clima por CEP

```http
GET /weather?cep=CEP
```

### Parâmetros

| Nome | Tipo   | Descrição                |
| ---- | ------ | ------------------------ |
| cep  | string | CEP válido com 8 dígitos |

---

## 📥 Respostas

### ✅ Sucesso

**HTTP 200**

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

---

### ❌ Erros

| Código | Mensagem             |
| ------ | -------------------- |
| 401    | API key missing      |
| 403    | Invalid API key      |
| 404    | can not find zipcode |
| 422    | invalid zipcode      |

---

## 🧪 Exemplos de Uso

### 📌 cURL

```bash
curl -X GET "https://SEU_ENDPOINT/weather?cep=01001000" \
  -H "x-api-key: SUA_API_KEY"
```

---

### 📌 Arquivo HTTP (VS Code / IntelliJ)

Arquivo: `weather.http`

```http
### Get weather by CEP
GET https://SEU_ENDPOINT/weather?cep=01001000
x-api-key: SUA_API_KEY
```

---

## 🐳 Executando Localmente com Docker

### Build da imagem

```bash
docker build -t cep-weather .
```

### Executar

```bash
docker run -p 8080:8080 \
  -e API_KEY=abc123 \
  -e WEATHER_API_KEY=SUA_CHAVE_WEATHER_API \
  cep-weather
```

---

## ☁️ Deploy no Google Cloud Run

### Build e Push da imagem

```bash
gcloud auth configure-docker us-central1-docker.pkg.dev

docker build -t us-central1-docker.pkg.dev/SEU_PROJECT/cep-weather/api .
docker push us-central1-docker.pkg.dev/SEU_PROJECT/cep-weather/api
```

---

### Deploy

```bash
gcloud run deploy cep-weather \
  --image us-central1-docker.pkg.dev/SEU_PROJECT/cep-weather/api \
  --region us-central1 \
  --platform managed \
  --allow-unauthenticated \
  --set-env-vars API_KEY=abc123,WEATHER_API_KEY=SUA_CHAVE
```

---

## 🧪 Testes Automatizados

Os testes cobrem:

- Validação de CEP
- Autenticação via API Key
- Conversão de temperatura
- Cenários de erro (422, 404, 401, 403)

### Executar testes

```bash
go test ./...
```

## 📎 Observações

- O serviço roda na porta **8080**, conforme exigido pelo Cloud Run
- Nenhuma chave sensível é versionada no repositório

---

## 👨‍💻 Autor

**Ramires Marques**
Backend Engineer | Go | .NET | Cloud | Distributed Systems
