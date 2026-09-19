// Package proxy implements mandatory egress (CONNECT + HTTP L7 + TLS terminate).
// Process logs use github.com/glaciforge/slogx via osg-runtime/logging when wired
// by the sandbox driver; audit lines remain OCSF shorthand on the audit writer.
package proxy
