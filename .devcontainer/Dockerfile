# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.191.1/containers/debian/.devcontainer/base.Dockerfile

# [Choice] Debian version: bullseye, buster, stretch
ARG VARIANT="bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# ** [Optional] Uncomment this section to install additional packages. **
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get install -y golang \
    && apt-get install -y libnspr4 libnss3 libexpat1 libfontconfig1 libuuid1 \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# you set GOPATH. install go package 
ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

# これが無いとvscodeから書き込めない。公式のDockerfileもこれなので大丈夫なはず。
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


# [Optional] Uncomment the next line to use go get to install anything else you need
# install vscode inteligent
RUN go get -v golang.org/x/tools/gopls github.com/ramya-rao-a/go-outline

# chromeで使うamd64bitに依存するパッケージをインストールするためにamd64bitを有効化
RUN dpkg --add-architecture amd64 \
    && dpkg --print-foreign-architectures \
    && apt-get update

# RUN apt install -y libappindicator1 fonts-liberation libasound2 libnspr4 libnss3 libxss1 lsb-release xdg-utils
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && apt install -y ./google-chrome-stable_current_amd64.deb \
    && rm ./google-chrome-stable_current_amd64.deb

# 日本語のフォントを入れるために必要。
RUN apt install -y task-japanese \
    && locale-gen ja_JP.UTF-8 \
    && localedef -f UTF-8 -i ja_JP ja_JP

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
