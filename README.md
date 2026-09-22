<h1 align="center">whaleshell-proxy</h1>

<p align="center">
  <strong>Egress proxy for sandboxes</strong><br>
  CONNECT + L7 terminate, secret rewrite, policy.local advisor, middleware pipeline.
</p>
<p align="center">
  <a href="https://github.com/whaleshell/whaleshell-proxy/actions/workflows/ci.yml"><img src="https://github.com/whaleshell/whaleshell-proxy/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/whaleshell/whaleshell-proxy"><img src="https://pkg.go.dev/badge/github.com/whaleshell/whaleshell-proxy.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/whaleshell/whaleshell-proxy"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/whaleshell">whaleshell / whaleshell</a> ecosystem</sub>
</p>

---

## Overview

**whaleshell-proxy** is the mandatory egress sidecar for whaleshell sandboxes. It enforces policy at L4/L7, rewrites credential placeholders, emits OCSF audit lines, and exposes `policy.local` for in-sandbox proposals.

### Key Features

| Category | Capabilities |
|----------|--------------|
| **L4** | CONNECT allowlist from `whaleshell-core` engine |
| **L7** | TLS terminate for REST / GraphQL / MCP rules |
| **Secrets** | Placeholder rewrite + gateway secret refresh |
| **Advisor** | `https://policy.local/v1/{policy,denials,proposals}` |
| **Hot reload** | Watch policy file; Apply under RWMutex |

---

## Installation

```bash
go get github.com/whaleshell/whaleshell-proxy@latest
```

Usually run as the sandbox sidecar (started by `whaleshell-driver`), not as a standalone service.

**Requirements:** Go 1.27+

---

## Quick Start

```bash
# From the CLI (host-side debug proxy):
whaleshell proxy --listen 127.0.0.1:3128 --policy ./policy.yaml
```

---

## Package Structure

| Path | Purpose |
|------|---------|
| `proxy/` | Server, MITM, audit, policy.local |
| `proxy/middleware/` | Request pipeline hooks |


---

## Related

| Resource | Link |
|----------|------|
| Roadmap | [ROADMAP.md](./ROADMAP.md) |
| Organization | [https://github.com/whaleshell](https://github.com/whaleshell) |
| Organization overview | [github.com/whaleshell](https://github.com/whaleshell) |
| pkg.go.dev | [`github.com/whaleshell/whaleshell-proxy`](https://pkg.go.dev/github.com/whaleshell/whaleshell-proxy) |

## License

[MIT](./LICENSE) © whaleshell
