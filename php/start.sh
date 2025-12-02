#!/bin/bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# LOAD PORTABLE LIBRARIES
# Ini mantra supaya Linux baca library dari folder 'libs' kita dulu
export LD_LIBRARY_PATH="$BASE_DIR/libs:$LD_LIBRARY_PATH"

echo "Starting Evergon PHP (Portable Mode)..."

# JALANKAN PHP-FPM
# Menggunakan config relative terhadap folder ini
"$BASE_DIR/sbin/php-fpm" \
    -p "$BASE_DIR" \
    -y "$BASE_DIR/etc/php-fpm.conf" \
    -c "$BASE_DIR/etc/php.ini" \
    --nodaemonize
