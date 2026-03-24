# syntax=docker/dockerfile:1

# ── Stage 1: build secmux ────────────────────────────────────────────────────
FROM golang:1.26-alpine AS builder

ARG VERSION=dev
ARG COMMIT=none

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
    -o /secmux ./cmd/secmux

# ── Stage 2: runtime image ───────────────────────────────────────────────────
FROM python:3.12-slim

# Pin scanner versions — bump these to upgrade bundled tools
ARG GITLEAKS_VERSION=8.21.2
ARG TRUFFLEHOG_VERSION=3.88.1

# Build-time arch vars (set automatically by `docker buildx`)
ARG TARGETARCH
ARG TARGETOS=linux

RUN apt-get update && apt-get install -y --no-install-recommends \
        git \
        curl \
        ca-certificates \
        make \
    && rm -rf /var/lib/apt/lists/*

# ── gitleaks ──────────────────────────────────────────────────────────────────
# gitleaks uses "x64" for amd64, "arm64" for arm64
RUN set -eux; \
    case "${TARGETARCH}" in \
        amd64) GL_ARCH="x64" ;; \
        arm64) GL_ARCH="arm64" ;; \
        *) echo "unsupported arch: ${TARGETARCH}" && exit 1 ;; \
    esac; \
    curl -sSfL \
        "https://github.com/gitleaks/gitleaks/releases/download/v${GITLEAKS_VERSION}/gitleaks_${GITLEAKS_VERSION}_${TARGETOS}_${GL_ARCH}.tar.gz" \
        | tar -xz -C /usr/local/bin gitleaks; \
    gitleaks version

# ── trufflehog ───────────────────────────────────────────────────────────────
RUN set -eux; \
    curl -sSfL \
        "https://github.com/trufflesecurity/trufflehog/releases/download/v${TRUFFLEHOG_VERSION}/trufflehog_${TRUFFLEHOG_VERSION}_${TARGETOS}_${TARGETARCH}.tar.gz" \
        | tar -xz -C /usr/local/bin trufflehog; \
    trufflehog --version

# ── Python-based scanners ────────────────────────────────────────────────────
# Pin detect-secrets to 1.4.x: 1.5.0 actively verifies secrets via external APIs,
# which filters out all fake/test secrets. 1.4.0 reports all detected candidates.
RUN pip install --no-cache-dir "detect-secrets==1.4.0" ggshield

# ── git-secrets ───────────────────────────────────────────────────────────────
RUN git clone --depth=1 https://github.com/awslabs/git-secrets /tmp/git-secrets \
    && make -C /tmp/git-secrets install \
    && rm -rf /tmp/git-secrets \
    && git secrets --register-aws --global

# ── secmux ────────────────────────────────────────────────────────────────────
COPY --from=builder /secmux /usr/local/bin/secmux

# Drop to non-root
RUN useradd -r -u 1000 -g nogroup -d /scan -s /sbin/nologin secmux \
    && mkdir /scan && chown secmux /scan
USER secmux

WORKDIR /scan

ENTRYPOINT ["secmux"]
CMD ["--help"]
