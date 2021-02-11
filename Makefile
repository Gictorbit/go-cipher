build:
	cd src ;go build -o ../sezarapp ;cd -

clean:
	rm bin/*

compile:
	# 32-Bit
	# FreeBDS
	mkdir -p ../bin
	cd src;GOOS=freebsd GOARCH=386 go build -o ../bin/main-freebsd-386 ;cd -
    # MacOS
	cd src;GOOS=darwin GOARCH=386 go build -o ../bin/main-darwin-386 ;cd -
    # Linux
	cd src;GOOS=linux GOARCH=386 go build -o ../bin/main-linux-386 ;cd -
    # Windows
	cd src;GOOS=windows GOARCH=386 go build -o ../bin/main-windows-386.exe ;cd -
    # 64-Bit
    # FreeBDS
	cd src;GOOS=freebsd GOARCH=amd64 go build -o ../bin/main-freebsd-amd64 ;cd -
    # MacOS
	cd src;GOOS=darwin GOARCH=amd64 go build -o ../bin/main-darwin-amd64 ;cd -
    # Linux
	cd src;GOOS=linux GOARCH=amd64 go build -o ../bin/main-linux-amd64 ;cd -
    # Winodws
	cd src;GOOS=windows GOARCH=amd64 go build -o ../bin/main-windows-amd64.exe ;cd - 
	