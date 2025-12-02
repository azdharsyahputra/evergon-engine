#!/bin/bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Stopping Nginx..."
"$BASE_DIR/nginx/stop.sh"

echo "Stopping PHP-FPM..."
"$BASE_DIR/php/stop.sh"

echo "✅ Evergon Stopped."
