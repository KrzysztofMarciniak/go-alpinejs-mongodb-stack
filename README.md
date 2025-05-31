# go-alpinejs-mongodb-stack
go alpinejs mongodb stack
### Production Graph
```mermaid
flowchart TD
 subgraph External_Access["External Access"]
        Browser["Browser"]
  end
 subgraph Public_Network["Public Network"]
        NGINX["NGINX Reverse Proxy (Frontend + API)"]
        Grafana["Grafana (port 3000)"]
  end
 subgraph Internal_Backend["Internal Backend"]
        GoAPI["Go API Service"]
        MongoDB["MongoDB Database"]
  end
 subgraph Monitoring_Network["Monitoring Network (Internal Only)"]
        NGINX_Exporter["NGINX Exporter"]
        Prometheus["Prometheus Monitoring"]
  end
    Browser <-- User Access --> NGINX
    Browser <-- Admin Access --> Grafana
    NGINX -- HTTP (80) / HTTPS (443) --> Browser
    NGINX -- Internal API (8080) --> GoAPI
    GoAPI -- DB Connection (8080) --> MongoDB
    NGINX -- Metrics (9113) /nginx_status --> NGINX_Exporter
    GoAPI -- Metrics (9091) /metrics --> Prometheus
    NGINX_Exporter --> Prometheus
    Prometheus --> Grafana
    style Browser stroke:#000000
    style NGINX stroke:#00C853
    style Grafana stroke:#FF6D00
    style GoAPI stroke:#2962FF
    style MongoDB stroke:#00C853
    style NGINX_Exporter stroke:#00C853
    style Prometheus stroke:#FFD600
    style External_Access stroke:#D50000
    style Public_Network stroke:#D50000
    style Internal_Backend stroke:#2962FF
    style Monitoring_Network stroke:#AA00FF
```
to use grafana with prometheus go to http://localhost:3000/
```bash
http://prometheus:9090
```

start:
```bash
cd frontend/ ; npm install; ../develop.sh;
```
