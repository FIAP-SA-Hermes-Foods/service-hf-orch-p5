#!/bin/bash

echo -ne "\n\n[38;2;150;214;255m      ============================= TEST COVERAGE =============================[0m\n\n"

go test -v -count=1 -cover -failfast -coverprofile cover.out ./
go tool cover -html=cover.out -o coverage.html

echo -ne "\n\n[38;2;150;214;255m      ============================= COVERAGE DONE =============================[0m\n\n"
