TMPDIR ?= /tmp/testman-integration

integration:
	# test with testman
	@mkdir -p $(TMPDIR)
	testman test -timeout=10s -run ^TestStable ./...  # should always work
	testman test -timeout=60s -run ^TestUnstable -retry=50 ./...  # should work at least 1/50
	(testman test -timeout=10s -continue-on-error -retry=10 ./...; echo "EXIT CODE: $$?") | tee $(TMPDIR)/got.continue-on-error.log	  # should fail, but also run all tests
	diff ./testdata/expected.continue-on-error.log $(TMPDIR)/got.continue-on-error.log

	# test with default tools
	go install moul.io/retry
	go test -run ^TestStable -count=20 ./... >/dev/null                                                     # should always work
	retry -m=5  --interval=0 -- "(go test -run ^TestBroken -count=1 ./... >/dev/null)" && exit 1 || exit 0  # should always fail
	retry -m=50 --interval=0 -- "(go test -run ^TestUnstable -count 1 ./... >/dev/null)"                    # should work at least 1/50
	go mod tidy

	@echo "SUCCESS."
	@rm -rf $(TMPDIR)
