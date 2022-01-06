FROM --platform=$BUILDPLATFORM "traefik:2.5.6"

ARG PLUGIN_MODULE="github.com/fdevillard/plugin"
RUN mkdir -p /plugins-local/src/${PLUGIN_MODULE}

COPY . /plugins-local/src/${PLUGIN_MODULE}
