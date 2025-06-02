# go-alpinejs-mongodb-stack
go alpinejs mongodb stack
### Production Graph
```mermaid

flowchart TD
 subgraph External_Access["External Access"]
        Browser["Browser"]
  end

 subgraph Public_Network["Public Network"]
        NGINX["NGINX Reverse Proxy with Self-signed TLS cert (generated at build time), WAF (Frontend + API /api/)"]
        Grafana["Grafana (port 3000)"]
  end

 subgraph Internal_Backend["Internal Backend"]
        GoAPI["Go API Service"]
        MongoDB["MongoDB Database"]
  end

 subgraph Monitoring_Network["Monitoring Network (Internal Only)"]
        NGINX_Exporter["NGINX Exporter"]
        Prometheus["Prometheus Monitoring"]
        Loki["Loki (logs from NGINX)"]
        Promtail["Promtail (tails logs, sends to Loki)"]
  end

    Browser <-- User Access --> NGINX
    Browser <-- Admin Access --> Grafana
    NGINX -- HTTP (80) / HTTPS (443) --> Browser
    NGINX -- Internal API (8080) only for internal backend network --> GoAPI
    GoAPI <-- DB Connection (27017) --> MongoDB
    NGINX -- Metrics /nginx_status only for monitoring network --> NGINX_Exporter
    GoAPI -- Metrics (9091) /metrics --> Prometheus
    NGINX_Exporter -- "nginx-exporter:9113" --> Prometheus
    Prometheus -- nginx + go monitoring --> Grafana
    NGINX -- ModSecurity logs --> Promtail
    Promtail -- logs --> Loki
    Loki -- logs --> Grafana

    style Browser stroke:#000000
    style Grafana stroke:#FF6D00
    style GoAPI stroke:#2962FF
    style MongoDB stroke:#00C853
    style NGINX_Exporter stroke:#00C853
    style Prometheus stroke:#FFD600
    style Loki stroke:#AA00FF
    style Promtail stroke:#AA00FF
    style External_Access stroke:#D50000
    style Internal_Backend stroke:#2962FF
    style Monitoring_Network stroke:#AA00FF

```
### Development Graph
```mermaid
flowchart TD
  subgraph External_Access["External Access"]
    User["Developer / Browser"]
  end

  subgraph Default_Network["Default Network (Dev Backend + Frontend)"]
    frontend["Frontend (vite + tailwindcss)"]
    api["Go Backend (api)"]
    mongo["MongoDB"]
    grafana["Grafana (port 3000)"]
  end

  subgraph Monitoring_Network["Monitoring Network (Dev Monitoring)"]
    prometheus["Prometheus"]
  end

  %% Connections with ports and labels
  User -- "Access UI (80)" --> frontend
  frontend -- "API Calls (8080)" --> api
  api -- "DB Connection (27017)" --> mongo
  api -- "Metrics Endpoint (9091)" --> prometheus
  prometheus -- "Feeds Data" --> grafana
  User -- "Access Dashboards (3000)" --> grafana

  %% Styling nodes
  style User stroke:#000000,stroke-width:2px,fill:#E3F2FD
  style frontend stroke:#2962FF,stroke-width:2px,fill:#BBDEFB
  style api stroke:#2962FF,stroke-width:2px,fill:#90CAF9
  style mongo stroke:#00C853,stroke-width:2px,fill:#B9F6CA
  style prometheus stroke:#FFD600,stroke-width:2px,fill:#FFF9C4
  style grafana stroke:#FF6D00,stroke-width:2px,fill:#FFE0B2

  %% Styling subgraphs
  style External_Access stroke:#D50000,stroke-width:2px,fill:#FFEBEE
  style Default_Network stroke:#2962FF,stroke-width:3px,fill:#E3F2FD,stroke-dasharray: 5 5
  style Monitoring_Network stroke:#AA00FF,stroke-width:3px,fill:#F3E5F5,stroke-dasharray: 5 5
```

to use grafana with prometheus go to http://localhost:3000/
```bash
http://prometheus:9090
```

start:
```bash
cd frontend/ ; npm install; ../develop.sh;
```
