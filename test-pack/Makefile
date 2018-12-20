SUPPORTED_TEST_PACKS=port

run-all-tests:
	for testpack in $(SUPPORTED_TEST_PACKS) ; do \
		make -C $$testpack run-tests; \
	done
