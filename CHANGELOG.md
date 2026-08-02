# Changelog

All notable changes to the **LibrePod** fork of `frp-operator` are documented here.

This project is a fork of [`zufardhiyaulhaq/frp-operator`](https://github.com/zufardhiyaulhaq/frp-operator), adopted by LibrePod after the upstream project became unmaintained and stopped accepting pull requests. Versioning continues from upstream's `v0.7.0`.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.8.0] - 2026-08-02

First LibrePod release. Adds HTTPS subdomain support and rebrands the chart and documentation to the fork.

### Added
- `subdomain` field on HTTPS `Upstream` (mirrors the existing HTTP field), exposed as `<subdomain>.<subdomainHost>` when the FRP server has `subdomainHost` configured.
- `TestTemplateHTTPSUpstreamWithSubdomain` template test.
- `CHANGELOG.md` (this file).

### Changed
- The Helm chart now deploys `ghcr.io/librepod/frp-operator:v0.8.0` (was the inherited upstream image `ghcr.io/zufardhiyaulhaq/frp-operator:v0.7.0`, which lacked the subdomain feature).
- `customDomains` is no longer required on HTTPS `Upstream` — a subdomain-only HTTPS proxy is now valid.
- Helm chart `version` `1.5.0` → `1.6.0`; operator `appVersion` `0.7.0` → `0.8.0`.
- Rebranded chart and docs to the LibrePod fork: `Chart.yaml` `home`/`maintainers`, README badges and links, `helm repo add` URL, `CONTRIBUTING.md`, and the ansible clone URL now point at `librepod/frp-operator`; added an "About this fork" notice crediting the original author.

### Notes
- Versioning is a **minor** bump rather than major: the change is additive and the CRDs are still `v1alpha1`. A `1.0.0` / chart `2.0.0` release is reserved for the eventual `v1alpha1` → `v1` CRD graduation.
- The `frp.zufardhiyaulhaq.com` API group, the Go module path, and the leader-election lease ID are intentionally unchanged to preserve compatibility with existing installs.

## Pre-fork (upstream)

Changes prior to the fork are recorded in the upstream repository, up to tag [`v0.7.0`](https://github.com/zufardhiyaulhaq/frp-operator/releases/tag/v0.7.0).

[Unreleased]: https://github.com/librepod/frp-operator/compare/v0.8.0...HEAD
[0.8.0]: https://github.com/librepod/frp-operator/releases/tag/v0.8.0
