PROGRAMS = nc

nc: main.go
	go build -o $@ -ldflags "-w -s" $<

clean:
	rm -f $(PROGRAMS)
