.DEFAULT_GOAL := hello

DIRECTORY?=
FILE?=

# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error Undefined $1$(if $2, ($2))))

hello:
	@echo "Hello, this is the makefile for justforfun repository, please check the existing katas on Makefile"
# new example: make new DIRECTORY=directory_name FILE=file_name
new:
	$(call check_defined, DIRECTORY, h)
	$(call check_defined, FILE, h)
	@echo "creating new kata directory: ${DIRECTORY}, title: '\# ${DIRECTORY} katas' and file: ${FILE}.go"
	mkdir ${DIRECTORY} && \
	cd ${DIRECTORY} && \
	echo "package ${DIRECTORY}" > ${FILE}.go && \
	echo "package ${DIRECTORY}_test" > ${FILE}_test.go && \
	echo "# ${DIRECTORY} katas" > README.md && \
	cd ..
	