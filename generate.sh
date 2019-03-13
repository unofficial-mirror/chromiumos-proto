#!/bin/bash -e
#
# Copyright 2019 The Chromium OS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Runs protoc over the protos in this repo to produce generated proto code.

# Version of Protobuf compiler to get from CIPD.
CIPD_PROTOC_VERSION='v3.6.1'

# Move to this script's directory.
cd "$(dirname "$0")"

# Get protobuf compiler from CIPD.
cipd_root=.cipd_bin
cipd ensure \
  -log-level warning \
  -root "${cipd_root}" \
  -ensure-file - \
  <<< "infra/tools/protoc/\${platform} protobuf_version:${CIPD_PROTOC_VERSION}"

protoc="${cipd_root}/protoc"

# Clean up existing generated files.
find go -name '*.pb.go' -exec rm '{}' \;

# Go files need to be processed individually until this is fixed:
# https://github.com/golang/protobuf/issues/39
find src -name '*.proto' -exec \
  "${protoc}" -Isrc --go_out=paths=source_relative:go '{}' \;
