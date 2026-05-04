# alloy-config-generator

`alloygen` is a local Go CLI that generates starter Grafana Alloy configuration for Kubernetes observability pipelines.

It helps DevOps, SRE, Kubernetes, Cloud MSP, and platform engineers create readable Alloy configs when migrating from Prometheus Agent, Grafana Agent, Promtail, or OpenTelemetry Collector.

This is a deterministic local generator. It is not a SaaS product and does not call Kubernetes, Grafana, Loki, Mimir, Tempo, cloud provider, database, authentication, or AI APIs.

## Build

```sh
go build -o alloygen .
```

Run from source:

```sh
go run . version
```

## Install from GitHub Releases

Download a prebuilt binary from the GitHub Releases page.

Linux/macOS:

```sh
tar -xzf <archive>.tar.gz
chmod +x alloygen
./alloygen version
```

Windows:

Download the Windows archive, extract it, and run:

```powershell
alloygen.exe version
```

## Commands

```sh
alloygen init --output ./alloy-demo

alloygen generate --logs loki --metrics mimir --traces tempo --output ./generated

alloygen generate --logs loki --metrics none --traces none --output ./logs-only

alloygen generate --logs none --metrics mimir --traces none --remote-write-url http://mimir-nginx.monitoring.svc:80/api/v1/push --output ./metrics-only

alloygen generate --profile production --namespace monitoring --cluster-name prod-kr --output ./prod-alloy
```

## Generated Output

With `--format all`:

```text
generated/
├── README.md
├── config/
│   └── config.alloy
├── helm/
│   └── values.yaml
└── examples/
    ├── install.sh
    └── uninstall.sh
```

With `--format config`, only `README.md` and `config/config.alloy` are generated.

With `--format helm`, only `README.md` and `helm/values.yaml` are generated.

## Pipeline Options

- Logs: `none`, `loki`
- Metrics: `none`, `mimir`, `prometheus`
- Traces: `none`, `tempo`, `otlp`

At least one pipeline must be enabled.

## Profiles

- `dev`: simple readable config, lower Helm resource requests, useful for labs, test clusters, and PoCs.
- `production`: production-oriented starter configuration with review comments for endpoints, RBAC, service accounts, and resource sizing. It is not fully production-certified.

## Modes

- `kubernetes`: generates Kubernetes pod discovery, relabeling, log forwarding, metrics scraping, and OTLP trace receiver starters.
- `standalone`: generates local starter patterns for running Alloy outside the cluster.

## Default Endpoints

- Loki: `http://loki-gateway.monitoring.svc:80/loki/api/v1/push`
- Remote write: `http://mimir-nginx.monitoring.svc:80/api/v1/push`
- Tempo: `tempo-distributor.monitoring.svc:4317`
- OTLP: `otel-collector.monitoring.svc:4317`

## Helm Install Example

```sh
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update
kubectl create namespace monitoring
helm upgrade --install alloy grafana/alloy -n monitoring -f helm/values.yaml
```

## Limitations

- Does not deploy Alloy directly.
- Does not install Loki, Mimir, Tempo, Prometheus, or OpenTelemetry Collector.
- Does not validate config against a running Alloy instance.
- Does not generate credentials, tokens, secrets, or kubeconfig.
- Does not integrate with Grafana Cloud, Vault, cloud providers, or Kubernetes APIs.

Use Kubernetes Secrets, an external secret operator, or Helm secret management when credentials are required.

## Roadmap

- More template variants for common observability-stack-generator workflows.
- Optional GitHub Action integration.
- Optional config validation against local Alloy.
- Optional Grafana Cloud, Vault, and Helm chart version helpers.
