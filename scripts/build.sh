#!/usr/bin/env bash
set -euo pipefail

# ----- Config -----
VERSION="${VERSION:-v0.1.0}"
OUT_DIR="${OUT_DIR:-dist}"

# ----- Metadata -----
COMMIT="$(git rev-parse --short HEAD)"
BUILD_DATE="$(date -u +"%Y-%m-%dT%H:%M:%SZ")"

GOOS="$(go env GOOS)"
GOARCH="$(go env GOARCH)"

BIN_NAME="leakhound-${GOOS}-${GOARCH}"
OUT_PATH="${OUT_DIR}/${BIN_NAME}"

mkdir -p "${OUT_DIR}"

echo "Building ${OUT_PATH}"
echo "  Version   : ${VERSION}"
echo "  Commit    : ${COMMIT}"
echo "  Build date: ${BUILD_DATE}"
echo "  Platform  : ${GOOS}-${GOARCH}"

go build -o "${OUT_PATH}" -ldflags "\
-X main.Version=${VERSION} \
-X main.Commit=${COMMIT} \
-X main.BuildDate=${BUILD_DATE}" .

echo "Done."
echo "Verifying:"
"${OUT_PATH}" -v
