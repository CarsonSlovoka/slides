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

# slides -md=docs -fs pages -fs static -fs tmpl
docker run -p 8080:8080 -e PORT=8080 -e MD_DIR="docs" -e FS_DIRS="pages,static,tmpl" -v ./docs:/docs -v ./static:/static -v ./tmpl/:/tmpl --name SlidesCmd slides-cmd:v0.2.0



exec slides \
    -host "$HOST" \
    -port "$PORT" \
    -md "$MD_DIR" \
    "${FS_ARGS[@]}"
