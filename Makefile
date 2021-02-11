build -ldflags "-s -w":
	cd src ;go build -ldflags "-s -w" -o ../sezarapp ;cd -

clean:
	rm bin/*

compile:
	# 32-Bit
	# FreeBDS
	mkdir -p ../bin
	cd src;GOOS=freebsd GOARCH=386 go build -ldflags "-s -w" -o ../bin/caesar-cipher-freebsd-386 ;cd -
    # MacOS
	cd src;GOOS=darwin GOARCH=386 go build -ldflags "-s -w" -o ../bin/caesar-cipher-darwin-386 ;cd -
    # Linux
	cd src;GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ../bin/caesar-cipher-linux-386 ;cd -
    # Windows
	cd src;GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ../bin/caesar-cipher-windows-386.exe ;cd -
    # 64-Bit
    # FreeBDS
	cd src;GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o ../bin/caesar-cipher-freebsd-amd64 ;cd -
    # MacOS
	cd src;GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ../bin/caesar-cipher-darwin-amd64 ;cd -
    # Linux
	cd src;GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ../bin/caesar-cipher-linux-amd64 ;cd -
    # Winodws
	cd src;GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ../bin/caesar-cipher-windows-amd64.exe ;cd - 
	