FROM golang:1.18 AS buildplugin
ENV GO111MODULE on
ENV CGO_ENABLED 0
RUN mkdir /tmp/build
WORKDIR /tmp/build
RUN git clone https://github.com/xinhari/hari
WORKDIR /tmp/build/hari/service/debug/collector
RUN go build .

FROM netdata/netdata:latest
COPY --from=buildplugin /tmp/build/hari/service/debug/collector/collector /usr/libexec/netdata/plugins.d/micro.d.plugin
RUN chown root:netdata /usr/libexec/netdata/plugins.d/micro.d.plugin
RUN chmod 0750 /usr/libexec/netdata/plugins.d/micro.d.plugin
