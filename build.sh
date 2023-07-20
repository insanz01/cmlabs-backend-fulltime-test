#!/bin/bash

# Jalankan perintah go test
echo "Running go test..."
go test ./test

# Periksa apakah perintah go test selesai dengan sukses (exit code 0)
if [ $? -eq 0 ]; then
  # Jika go test berhasil, jalankan perintah go build
  echo "Running go build..."
  go build -o ./dist/app
  echo "Build successful! Executable binary created as app"
else
  # Jika go test gagal, tampilkan pesan kesalahan dan hentikan build
  echo "Unit tests failed! Build aborted."
fi
