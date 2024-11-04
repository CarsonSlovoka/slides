#!/bin/bash
set -e

HOST="${HOST:-}"
PORT="${PORT:-9000}"
MD_DIR="${MD_DIR:-/md}"

# -e FS_DIRS="pages1,pages2"
FS_ARGS=()
if [ -n "$FS_DIRS" ]; then
    IFS=',' read -ra DIRS <<< "$FS_DIRS"
    for dir in "${DIRS[@]}"; do
        FS_ARGS+=("-fs" "$dir")
    done
# else
#    FS_ARGS+=("-fs" "pages")
fi

exec slides \
    -host "$HOST" \
    -port "$PORT" \
    -md "$MD_DIR" \
    "${FS_ARGS[@]}"
