#!/bin/bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
export LD_LIBRARY_PATH="$BASE_DIR/libs:$LD_LIBRARY_PATH"

echo "Stopping Evergon Nginx..."
# Perintah -s stop menyuruh Nginx berhenti baik-baik
"$BASE_DIR/sbin/nginx" -p "$BASE_DIR" -s stop
