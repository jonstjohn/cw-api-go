# cw-api-go
ClimbingWeather.com API service written in Go

## Deploy to Google Cloud Run

```
gcloud run deploy cw-api-go --region us-central1 --allow-unauthenticated --source=.
```

## Database URLs

gcloud sql instances describe cw-pg-dev | grep connectionName
Pg: user=cwapp password=XXXXXXX database=cw host=/cloudsql/api-project-736062072361:us-central1:cw-pg-dev
CRDB: postgresql://cwapp:XXXXX@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/cw?sslmode=verify-full&sslrootcert=/certs/CW_CRDB_CRT&options=--cluster%3Dcw-test-4902

Squamish
INSERT INTO area (name, latitude, longitude) values ('Squamish', 49.7016, -123.1558);
