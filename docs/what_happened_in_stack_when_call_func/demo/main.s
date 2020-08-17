# command-line-arguments
"".sum STEXT nosplit size=19 args=0x18 locals=0x0
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	TEXT	"".sum(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	PCDATA	$0, $-2
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	PCDATA	$1, $-2
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	PCDATA	$0, $0
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	PCDATA	$1, $0
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	MOVQ	"".b+16(SP), AX
	0x0005 00005 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	MOVQ	"".a+8(SP), CX
	0x000a 00010 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	ADDQ	CX, AX
	0x000d 00013 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	MOVQ	AX, "".~r2+24(SP)
	0x0012 00018 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:4)	RET
	0x0000 48 8b 44 24 10 48 8b 4c 24 08 48 01 c8 48 89 44  H.D$.H.L$.H..H.D
	0x0010 24 18 c3                                         $..
"".main STEXT nosplit size=1 args=0x0 locals=0x0
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	PCDATA	$0, $-2
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	PCDATA	$1, $-2
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:7)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:9)	PCDATA	$0, $-1
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:9)	PCDATA	$1, $-1
	0x0000 00000 (/Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go:9)	RET
	0x0000 c3                                               .
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info."".sum$abstract SDWARFINFO dupok size=29
	0x0000 04 6d 61 69 6e 2e 73 75 6d 00 01 01 11 61 00 00  .main.sum....a..
	0x0010 00 00 00 00 11 62 00 00 00 00 00 00 00           .....b.......
	rel 16+4 t=29 go.info.int+0
	rel 24+4 t=29 go.info.int+0
go.loc."".sum SDWARFLOC size=71
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 02 00 91 08 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00                             .......
	rel 0+8 t=53 "".sum+0
	rel 8+8 t=53 "".sum+19
	rel 35+8 t=53 "".sum+0
	rel 43+8 t=53 "".sum+19
go.info."".sum SDWARFINFO size=54
	0x0000 05 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 01 9c 13 00 00 00 00 00 00 00 00  ................
	0x0020 13 00 00 00 00 00 00 00 00 0f 7e 72 32 00 01 03  ..........~r2...
	0x0030 00 00 00 00 00 00                                ......
	rel 0+0 t=24 type.int+0
	rel 1+4 t=29 go.info."".sum$abstract+0
	rel 5+8 t=1 "".sum+0
	rel 13+8 t=1 "".sum+19
	rel 24+4 t=29 go.info."".sum$abstract+12
	rel 28+4 t=29 go.loc."".sum+0
	rel 33+4 t=29 go.info."".sum$abstract+20
	rel 37+4 t=29 go.loc."".sum+35
	rel 48+4 t=29 go.info.int+0
go.range."".sum SDWARFRANGE size=0
go.debuglines."".sum SDWARFMISC size=11
	0x0000 04 02 12 06 41 04 01 03 7d 06 01                 ....A...}..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=35
	0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 00                                         ...
	rel 11+8 t=1 "".main+0
	rel 19+8 t=1 "".main+1
	rel 29+4 t=30 gofile../Users/zy/GOPATH/src/github.com/xpzouying/golang-notes/docs/what_happened_in_stack_when_call_func/demo/main.go+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=10
	0x0000 04 02 03 03 14 04 01 03 78 01                    ........x.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
