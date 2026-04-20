#!/usr/bin/env bash

set -euo pipefail

STORAGE_DIR="/opt/render/project/.render"
INSTALL_DIR="${STORAGE_DIR}/chrome"
DOWNLOAD_DIR="${STORAGE_DIR}/downloads"
DEB_PATH="${DOWNLOAD_DIR}/google-chrome-stable_current_amd64.deb"
CHROME_BIN="${INSTALL_DIR}/opt/google/chrome/google-chrome"

mkdir -p "${INSTALL_DIR}" "${DOWNLOAD_DIR}"

if [ -x "${CHROME_BIN}" ]; then
  echo "Chrome already installed at ${CHROME_BIN}"
  exit 0
fi

echo "Downloading Google Chrome for Render..."
curl -fsSL "https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb" -o "${DEB_PATH}"

echo "Extracting Chrome..."
dpkg-deb -x "${DEB_PATH}" "${INSTALL_DIR}"

if [ ! -x "${CHROME_BIN}" ]; then
  echo "Chrome install failed: binary not found at ${CHROME_BIN}" >&2
  exit 1
fi

echo "Chrome installed successfully at ${CHROME_BIN}"
