#!/bin/bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "=== EVERGON FOR LINUX ==="
echo "📍 Detected Path: $BASE_DIR"

# --- CONFIG GENERATOR (SAFE MODE) ---
# Ambil dari Template (.tpl), ganti {{ROOT_PATH}} dengan path asli, simpan ke .conf
# Ini aman dijalankan berkali-kali.
sed "s|{{ROOT_PATH}}|$BASE_DIR/www|g" "$BASE_DIR/nginx/conf/nginx.conf.tpl" > "$BASE_DIR/nginx/conf/nginx.conf"
# ------------------------------------

echo "Starting PHP-FPM..."
"$BASE_DIR/php/start.sh" > /dev/null 2>&1 &
sleep 0.5

echo "Starting Nginx..."
"$BASE_DIR/nginx/start.sh" > /dev/null 2>&1 &
sleep 0.5

echo "------------------------------------------------"
echo "✅ Evergon is RUNNING!"
echo "📂 Document Root: $BASE_DIR/www"
echo "🌐 URL: http://localhost:8080"
echo "------------------------------------------------"