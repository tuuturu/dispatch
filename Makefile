BUILD_DIR=./build
INSTALL_DIR=/usr/bin
ASSETS_DIR=./assets

${BUID_DIR}/dispatcher:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/dispatcher

build: ${BUID_DIR}/dispatcher

install: build
	mkdir -p ${INSTALL_DIR}
	cp ${BUILD_DIR}/dispatcher ${INSTALL_DIR}
	cp ${ASSETS_DIR}/dispatcher.service /etc/systemd/system
	systemctl daemon-reload

uninstall:
	rm ${INSTALL_DIR}/dispatcher
	rm /etc/systemd/system/dispatcher.service

clean:
	rm -r ${BUILD_DIR}
