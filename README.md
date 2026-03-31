# Laravel + Golang Observability Service

Projet personnel pour construire un système d’observabilité moderne sur Kubernetes, avec:
- **Laravel** pour le dashboard & API frontend.
- **Golang** pour le backend observability (metrics, logs, alerting).
- **Kubernetes** pour le déploiement (Prometheus + Grafana + Loki + Jaeger).

## 📦 Stack

- Laravel (PHP 8.x)
- Golang (Go 1.20+)
- Kubernetes (Prometheus, Grafana, Loki, Jaeger)

## �� Comment déployer

1. Construire l’image Laravel (ex: `Dockerfile` dans laravel/).
2. Construire l’image Go (ex: `Dockerfile` dans golang/).
3. Déployer les manifests dans k8s/:

```bash
kubectl apply -f k8s/
```

## 🧩 Répertoire

- `laravel/`     → code Laravel + Dockerfile
- `golang/`      → service observability Go
- `k8s/`         → deployments, services, service monitors
