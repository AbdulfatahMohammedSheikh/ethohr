go test  ./test/checkhealth_test/checkhealth_test.go | sed ''/ok/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | sed ''/error/s//$(printf "\033[31merror\033[0m")/''
go test ./test/backend_test/*/*** | sed ''/ok/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | sed ''/error/s//$(printf "\033[31merror\033[0m")/''
