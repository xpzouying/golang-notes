# command-line-arguments
"".main STEXT size=116 args=0x0 locals=0x30
	0x0000 00000 (.../make_chan.go:3)	TEXT	"".main(SB), ABIInternal, $48-0
	0x0000 00000 (.../make_chan.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (.../make_chan.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (.../make_chan.go:3)	PCDATA	$0, $-2
	0x000d 00013 (.../make_chan.go:3)	JLS	109
	0x000f 00015 (.../make_chan.go:3)	PCDATA	$0, $-1
	0x000f 00015 (.../make_chan.go:3)	SUBQ	$48, SP
	0x0013 00019 (.../make_chan.go:3)	MOVQ	BP, 40(SP)
	0x0018 00024 (.../make_chan.go:3)	LEAQ	40(SP), BP
	0x001d 00029 (.../make_chan.go:3)	PCDATA	$0, $-2
	0x001d 00029 (.../make_chan.go:3)	PCDATA	$1, $-2
	0x001d 00029 (.../make_chan.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (.../make_chan.go:3)	FUNCDATA	$1, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
	0x001d 00029 (.../make_chan.go:3)	FUNCDATA	$2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (.../make_chan.go:4)	PCDATA	$0, $1
	0x001d 00029 (.../make_chan.go:4)	PCDATA	$1, $0
	0x001d 00029 (.../make_chan.go:4)	LEAQ	type.chan int(SB), AX
	0x0024 00036 (.../make_chan.go:4)	PCDATA	$0, $0
	0x0024 00036 (.../make_chan.go:4)	MOVQ	AX, (SP)
	0x0028 00040 (.../make_chan.go:4)	MOVQ	$10, 8(SP)
	0x0031 00049 (.../make_chan.go:4)	CALL	runtime.makechan(SB)
	0x0036 00054 (.../make_chan.go:4)	PCDATA	$0, $1
	0x0036 00054 (.../make_chan.go:4)	MOVQ	16(SP), AX
	0x003b 00059 (.../make_chan.go:4)	PCDATA	$0, $0
	0x003b 00059 (.../make_chan.go:4)	MOVQ	AX, "".c1+32(SP)
	0x0040 00064 (.../make_chan.go:5)	PCDATA	$0, $1
	0x0040 00064 (.../make_chan.go:5)	LEAQ	type.chan int(SB), AX
	0x0047 00071 (.../make_chan.go:5)	PCDATA	$0, $0
	0x0047 00071 (.../make_chan.go:5)	MOVQ	AX, (SP)
	0x004b 00075 (.../make_chan.go:5)	MOVQ	$0, 8(SP)
	0x0054 00084 (.../make_chan.go:5)	CALL	runtime.makechan(SB)
	0x0059 00089 (.../make_chan.go:5)	PCDATA	$0, $1
	0x0059 00089 (.../make_chan.go:5)	MOVQ	16(SP), AX
	0x005e 00094 (.../make_chan.go:5)	PCDATA	$0, $0
	0x005e 00094 (.../make_chan.go:5)	MOVQ	AX, "".c2+24(SP)
	0x0063 00099 (.../make_chan.go:8)	MOVQ	40(SP), BP
	0x0068 00104 (.../make_chan.go:8)	ADDQ	$48, SP
	0x006c 00108 (.../make_chan.go:8)	RET
	0x006d 00109 (.../make_chan.go:8)	NOP
	0x006d 00109 (.../make_chan.go:3)	PCDATA	$1, $-1
	0x006d 00109 (.../make_chan.go:3)	PCDATA	$0, $-2
	0x006d 00109 (.../make_chan.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0072 00114 (.../make_chan.go:3)	PCDATA	$0, $-1
	0x0072 00114 (.../make_chan.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 5e 48  eH..%....H;a.v^H
	0x0010 83 ec 30 48 89 6c 24 28 48 8d 6c 24 28 48 8d 05  ..0H.l$(H.l$(H..
	0x0020 00 00 00 00 48 89 04 24 48 c7 44 24 08 0a 00 00  ....H..$H.D$....
	0x0030 00 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24 20  ......H.D$.H.D$ 
	0x0040 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24 08  H......H..$H.D$.
	0x0050 00 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48 89  .........H.D$.H.
	0x0060 44 24 18 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00  D$.H.l$(H..0....
	0x0070 00 00 eb 8c                                      ....
	rel 5+4 t=17 TLS+0
	rel 32+4 t=16 type.chan int+0
	rel 50+4 t=8 runtime.makechan+0
	rel 67+4 t=16 type.chan int+0
	rel 85+4 t=8 runtime.makechan+0
	rel 110+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.producer.main SDWARFINFO dupok size=0
	0x0000 2d 4e 20 2d 6c                                   -N -l
go.cuinfo.packagename.main SDWARFINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=59
	0x0000 03 6d 61 69 6e 2e 6d 61 69 6e 00 00 00 00 00 00  .main.main......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00  ................
	0x0020 00 01 0a 63 32 00 05 00 00 00 00 02 91 60 0a 63  ...c2........`.c
	0x0030 31 00 04 00 00 00 00 02 91 68 00                 1........h.
	rel 0+0 t=24 type.chan int+0
	rel 11+8 t=1 "".main+0
	rel 19+8 t=1 "".main+116
	rel 29+4 t=30 gofile...../make_chan.go+0
	rel 39+4 t=29 go.info.chan int+0
	rel 51+4 t=29 go.info.chan int+0
go.range."".main SDWARFRANGE size=0
go.debuglines."".main SDWARFMISC size=27
	0x0000 04 02 11 0a a5 9c 06 55 06 08 38 06 55 06 91 06  .......U..8.U...
	0x0010 41 06 76 03 7f 6f 04 01 03 7e 01                 A.v..o...~.
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*chan int- SRODATA dupok size=12
	0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ed 7b ed 3b 08 08 08 36 00 00 00 00 00 00 00 00  .{.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 48+8 t=1 type.chan int+0
type.chan int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 91 55 cb 71 0a 08 08 32 00 00 00 00 00 00 00 00  .U.q...2........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan int-+0
	rel 44+4 t=6 type.*chan int+0
	rel 48+8 t=1 type.int+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·f207267fbf96a0178e8758c6e3e0ce28 SRODATA dupok size=9
	0x0000 01 00 00 00 02 00 00 00 00                       .........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
