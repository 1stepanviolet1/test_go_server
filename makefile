PROJECT_NAME = server
BIN_DIR = ./bin/
MAIN_SRC_FILE = server.go

COMP = go


all:
	@ ${COMP} run ${MAIN_SRC_FILE}


build:
	@ ${COMP} build -o ${BIN_DIR}${PROJECT_NAME}.exe ${MAIN_SRC_FILE}
