#!/bin/bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 1. Load Library Portable (Biar jalan di distro lain)
export LD_LIBRARY_PATH="$BASE_DIR/libs:$LD_LIBRARY_PATH"

echo "Starting Evergon Nginx..."

# 2. Jalankan Nginx
# -p "$BASE_DIR" : INI KUNCINYA! Memaksa Nginx pake folder ini sbg root, bukan /evergon_dummy
"$BASE_DIR/sbin/nginx" -p "$BASE_DIR"
