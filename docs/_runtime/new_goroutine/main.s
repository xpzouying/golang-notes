# command-line-arguments
"".sum STEXT nosplit size=25 args=0x18 locals=0x0
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	TEXT	"".sum(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	PCDATA	$0, $-2
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	PCDATA	$1, $-2
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	PCDATA	$0, $0
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	PCDATA	$1, $0
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:3)	MOVQ	$0, "".~r2+24(SP)
	0x0009 00009 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:4)	MOVQ	"".a+8(SP), AX
	0x000e 00014 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:4)	ADDQ	"".b+16(SP), AX
	0x0013 00019 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:4)	MOVQ	AX, "".~r2+24(SP)
	0x0018 00024 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:4)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 8b 44 24 08 48 03  H.D$.....H.D$.H.
	0x0010 44 24 10 48 89 44 24 18 c3                       D$.H.D$..
"".main STEXT size=107 args=0x0 locals=0x40
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	TEXT	"".main(SB), ABIInternal, $64-0
	0x0000 00000 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	MOVQ	(TLS), CX
	0x0009 00009 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	CMPQ	SP, 16(CX)
	0x000d 00013 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$0, $-2
	0x000d 00013 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	JLS	100
	0x000f 00015 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$0, $-1
	0x000f 00015 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	SUBQ	$64, SP
	0x0013 00019 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	MOVQ	BP, 56(SP)
	0x0018 00024 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	LEAQ	56(SP), BP
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$0, $-2
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$1, $-2
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:8)	PCDATA	$0, $0
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:8)	PCDATA	$1, $0
	0x001d 00029 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:8)	MOVQ	$1, "".v1+48(SP)
	0x0026 00038 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:9)	MOVQ	$2, "".v2+40(SP)
	0x002f 00047 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	MOVL	$24, (SP)
	0x0036 00054 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	PCDATA	$0, $1
	0x0036 00054 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	LEAQ	"".sum·f(SB), AX
	0x003d 00061 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	PCDATA	$0, $0
	0x003d 00061 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	MOVQ	AX, 8(SP)
	0x0042 00066 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	MOVQ	"".v1+48(SP), AX
	0x0047 00071 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	MOVQ	AX, 16(SP)
	0x004c 00076 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	MOVQ	$2, 24(SP)
	0x0055 00085 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:11)	CALL	runtime.newproc(SB)
	0x005a 00090 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:12)	MOVQ	56(SP), BP
	0x005f 00095 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:12)	ADDQ	$64, SP
	0x0063 00099 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:12)	RET
	0x0064 00100 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:12)	NOP
	0x0064 00100 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$1, $-1
	0x0064 00100 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$0, $-2
	0x0064 00100 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x0069 00105 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	PCDATA	$0, $-1
	0x0069 00105 (/Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go:7)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 55 48  dH..%....H;a.vUH
	0x0010 83 ec 40 48 89 6c 24 38 48 8d 6c 24 38 48 c7 44  ..@H.l$8H.l$8H.D
	0x0020 24 30 01 00 00 00 48 c7 44 24 28 02 00 00 00 c7  $0....H.D$(.....
	0x0030 04 24 18 00 00 00 48 8d 05 00 00 00 00 48 89 44  .$....H......H.D
	0x0040 24 08 48 8b 44 24 30 48 89 44 24 10 48 c7 44 24  $.H.D$0H.D$.H.D$
	0x0050 18 02 00 00 00 e8 00 00 00 00 48 8b 6c 24 38 48  ..........H.l$8H
	0x0060 83 c4 40 c3 e8 00 00 00 00 eb 95                 ..@........
	rel 5+4 t=17 TLS+0
	rel 57+4 t=16 "".sum·f+0
	rel 86+4 t=8 runtime.newproc+0
	rel 101+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFINFO dupok size=0
	0x0000 2d 4e 20 2d 6c                                   -N -l
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".sum SDWARFLOC size=0
go.info."".sum SDWARFINFO size=71
	0x0000 03 6d 61 69 6e 2e 73 75 6d 00 00 00 00 00 00 00  .main.sum.......
	0x0010 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
	0x0020 01 0f 61 00 00 03 00 00 00 00 01 9c 0f 62 00 00  ..a..........b..
	0x0030 03 00 00 00 00 02 91 08 0f 7e 72 32 00 01 03 00  .........~r2....
	0x0040 00 00 00 02 91 10 00                             .......
	rel 0+0 t=24 type.int+0
	rel 10+8 t=1 "".sum+0
	rel 18+8 t=1 "".sum+25
	rel 28+4 t=30 gofile../Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go+0
	rel 38+4 t=29 go.info.int+0
	rel 49+4 t=29 go.info.int+0
	rel 63+4 t=29 go.info.int+0
go.range."".sum SDWARFRANGE size=0
go.debuglines."".sum SDWARFMISC size=12
	0x0000 04 02 11 6a 06 41 04 01 03 7d 06 01              ...j.A...}..
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=59
	0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 0a 76 32 00 09 00 00 00 00 02 91 60 0a 76  ...v2........`.v
	0x0030 31 00 08 00 00 00 00 02 91 68 00                 1........h.
	rel 0+0 t=24 type.int+0
	rel 11+8 t=1 "".main+0
	rel 19+8 t=1 "".main+107
	rel 29+4 t=30 gofile../Users/zouying/GOPATH/src/github.com/xpzouying/golang-notes/docs/new_goroutine/main.go+0
	rel 39+4 t=29 go.info.int+0
	rel 51+4 t=29 go.info.int+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=23
	0x0000 04 02 03 01 14 0a a5 9c 6a 6b 06 55 06 08 88 03  ........jk.U....
	0x0010 7f 6f 04 01 03 7a 01                             .o...z.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
"".sum·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".sum+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
