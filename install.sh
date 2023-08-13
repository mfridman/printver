#!/bin/sh
# Adapted from the Deno installer: Copyright 2019 the Deno authors. All rights reserved. MIT license.
# Ref: https://github.com/denoland/deno_install

set -e

os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

if [ $# -eq 0 ]; then
    download_uri="https://github.com/mfridman/printver/releases/latest/download/printver_${os}_${arch}"
else
    download_uri="https://github.com/mfridman/printver/releases/download/${1}/printver_${os}_${arch}"
fi

printver_install="${PRINTVER_INSTALL:-/usr/local}"
bin_dir="${printver_install}/bin"
exe="${bin_dir}/printver"

if [ ! -d "${bin_dir}" ]; then
    mkdir -p "${bin_dir}"
fi

curl --silent --show-error --location --fail --output "${exe}" "$download_uri"
chmod +x "${exe}"

echo "printver installed successfully to ${exe}"
if command -v printver >/dev/null; then
    echo "Run 'printver'"
fi
