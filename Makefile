SOURCE=$(shell find . -name '*.go')
RPI_ADDR=pi@192.168.1.17

dev:
	go build ./cmd/volctl-web/ && ./volctl-web

clean:
	rm -f volctl-web
	rm -f volctl-web-arm

build: volctl-web-arm

volctl-web-arm: $(SOURCE)
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./volctl-web-arm ./cmd/volctl-web

setup-service:
	scp deploy/volctl.service '$(RPI_ADDR):volctl.service'
	ssh $(RPI_ADDR) sudo mv volctl.service /lib/systemd/system/volctl.service
	ssh $(RPI_ADDR) sudo systemctl enable volctl

service-status:
	ssh $(RPI_ADDR) sudo systemctl status volctl

deploy-service: volctl-web-arm
	scp volctl-web-arm '$(RPI_ADDR):volctl-web'

restart-service:  deploy-service
	ssh $(RPI_ADDR) sudo systemctl stop volctl
	ssh $(RPI_ADDR) sudo mv volctl-web /usr/local/bin/volctl-web
	ssh $(RPI_ADDR) sudo systemctl start volctl
	ssh $(RPI_ADDR) sudo systemctl status volctl
