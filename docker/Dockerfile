# The octobus binary is not compiled here. Compile it for the target platform
# first (scripts/build-octobus.sh, GOOS=linux GOARCH=<arch>) and put the result
# in dist/ in the build context — the image only packages it. Keeping the
# toolchain out of the image lets CI build linux/arm64 natively instead of
# running the whole toolchain under QEMU.
FROM ubuntu:26.04

ENV DEBIAN_FRONTEND=noninteractive \
	OCTOBUS_ADDR=0.0.0.0:9000 \
	OCTOBUS_DATA_DIR=/var/lib/octobus

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
		ca-certificates \
		curl \
		git \
		nodejs \
		npm \
		protobuf-compiler \
		libprotobuf-dev \
	&& rm -rf /var/lib/apt/lists/* \
	&& groupadd --system octobus \
	&& useradd --system --gid octobus --home-dir /var/lib/octobus --create-home --shell /usr/sbin/nologin octobus

ARG TARGETARCH
COPY dist/octobus-linux-${TARGETARCH} /usr/local/bin/octobus

USER octobus
WORKDIR /var/lib/octobus
VOLUME ["/var/lib/octobus"]
EXPOSE 9000

ENTRYPOINT ["octobus"]
CMD ["serve"]
