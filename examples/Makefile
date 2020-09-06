integration:
	# test with testman
	testman test -run ^TestStable ./...
	@# FIXME: test unstable tests
	@# FIXME: test broken tests

	# test with default tools
	go install moul.io/retry
	go test -run ^TestStable -count=20 ./... >/dev/null                                                     # should always work
	retry -m=5  --interval=0 -- "(go test -run ^TestBroken -count=1 ./... >/dev/null)" && exit 1 || exit 0  # should always fail
	retry -m=50 --interval=0 -- "(go test -run ^TestUnstable -count 1 ./... >/dev/null)"                    # should work at least 1/50
	go mod tidy

	@echo "SUCCESS."
