#!/usr/bin/env bash

install_buf() {
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Substitute BIN for your bin directory.
        # Substitute VERSION for the current released version.
        BIN="/usr/local/bin" && \
        VERSION="1.49.0" && \
        sudo curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
        -o "${BIN}/buf" && \
        sudo chmod +x "${BIN}/buf"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew tap bufbuild/buf && \
        brew install buf
    else
        echo "Unsupported OS: $OSTYPE"
        exit 1
    fi
}

install_npm_dependencies() {
    npm install -g @redocly/cli@1.9.1
}

install_go_dependencies() {
    echo "Installing buf..."
    install_buf

    echo "Installing go dependencies..."
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.14
    go install go.uber.org/mock/mockgen@v0.3.0
    go install  -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
}

check_version() {
    echo "Checking version..."
    go version
    oapi-codegen --version
    mockgen --version
    migrate --version
    buf --version
    protoc --version
}

install_npm_dependencies
install_go_dependencies
check_version