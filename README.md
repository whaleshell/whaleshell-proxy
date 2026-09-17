<h1 align="center">osg-proxy</h1>

<p align="center">
  <strong>Egress proxy for sandboxes</strong><br>
  CONNECT + L7 terminate, secret rewrite, policy.local advisor, middleware pipeline.
</p>
<p align="center">
  <a href="https://github.com/zorneth/osg-proxy/actions/workflows/ci.yml"><img src="https://github.com/zorneth/osg-proxy/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://pkg.go.dev/github.com/zorneth/osg-proxy"><img src="https://pkg.go.dev/badge/github.com/zorneth/osg-proxy.svg" alt="Go Reference"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="License"></a>
  <a href="https://github.com/zorneth/osg-proxy"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/zorneth">zorneth / osg</a> ecosystem</sub>
</p>

---

## Overview

**osg-proxy** is the mandatory egress sidecar for osg sandboxes. It enforces policy at L4/L7, rewrites credential placeholders, emits OCSF audit lines, and exposes `policy.local` for in-sandbox proposals.

### Key Features

| Category | Capabilities |
|----------|--------------|
| **L4** | CONNECT allowlist from `osg-core` engine |
| **L7** | TLS terminate for REST / GraphQL / MCP rules |
| **Secrets** | Placeholder rewrite + gateway secret refresh |
| **Advisor** | `https://policy.local/v1/{policy,denials,proposals}` |
| **Hot reload** | Watch policy file; Apply under RWMutex |

---

## Installation

```bash
go get github.com/zorneth/osg-proxy@latest
```

Usually run as the sandbox sidecar (started by `osg-driver`), not as a standalone service.

**Requirements:** Go 1.27+

---

## Quick Start

```bash
# From the CLI (host-side debug proxy):
osg proxy --listen 127.0.0.1:3128 --policy ./policy.yaml
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
| Organization | [https://github.com/zorneth](https://github.com/zorneth) |
| Organization overview | [github.com/zorneth](https://github.com/zorneth) |
| pkg.go.dev | [`github.com/zorneth/osg-proxy`](https://pkg.go.dev/github.com/zorneth/osg-proxy) |

## License

[MIT](./LICENSE) © zorneth
