TEXT main.main(SB) /Users/sunbin/study/test_golang/grammar/test_int.go
  test_int.go:5		0x1092f70		65488b0c2530000000		MOVQ GS:0x30, CX			
  test_int.go:5		0x1092f79		488d842478ffffff		LEAQ 0xffffff78(SP), AX			
  test_int.go:5		0x1092f81		483b4110			CMPQ 0x10(CX), AX			
  test_int.go:5		0x1092f85		0f867f030000			JBE 0x109330a				
  test_int.go:5		0x1092f8b		4881ec08010000			SUBQ $0x108, SP				
  test_int.go:5		0x1092f92		4889ac2400010000		MOVQ BP, 0x100(SP)			
  test_int.go:5		0x1092f9a		488dac2400010000		LEAQ 0x100(SP), BP			
  test_int.go:6		0x1092fa2		0f57c0				XORPS X0, X0				
  test_int.go:6		0x1092fa5		0f118424e0000000		MOVUPS X0, 0xe0(SP)			
  test_int.go:6		0x1092fad		0f57c0				XORPS X0, X0				
  test_int.go:6		0x1092fb0		0f118424f0000000		MOVUPS X0, 0xf0(SP)			
  test_int.go:6		0x1092fb8		488d8424e0000000		LEAQ 0xe0(SP), AX			
  test_int.go:6		0x1092fc0		4889442450			MOVQ AX, 0x50(SP)			
  test_int.go:6		0x1092fc5		8400				TESTB AL, 0(AX)				
  test_int.go:6		0x1092fc7		488d0dd2100100			LEAQ type.*+68320(SB), CX		
  test_int.go:6		0x1092fce		48898c24e0000000		MOVQ CX, 0xe0(SP)			
  test_int.go:6		0x1092fd6		488d0d33b60400			LEAQ main.statictmp_0(SB), CX		
  test_int.go:6		0x1092fdd		48898c24e8000000		MOVQ CX, 0xe8(SP)			
  test_int.go:6		0x1092fe5		8400				TESTB AL, 0(AX)				
  test_int.go:6		0x1092fe7		488d05b2080100			LEAQ type.*+66272(SB), AX		
  test_int.go:6		0x1092fee		48898424f0000000		MOVQ AX, 0xf0(SP)			
  test_int.go:6		0x1092ff6		488d0553b20400			LEAQ main.statictmp_1(SB), AX		
  test_int.go:6		0x1092ffd		48898424f8000000		MOVQ AX, 0xf8(SP)			
  test_int.go:6		0x1093005		488b442450			MOVQ 0x50(SP), AX			
  test_int.go:6		0x109300a		8400				TESTB AL, 0(AX)				
  test_int.go:6		0x109300c		eb00				JMP 0x109300e				
  test_int.go:6		0x109300e		48898424b0000000		MOVQ AX, 0xb0(SP)			
  test_int.go:6		0x1093016		48c78424b800000002000000	MOVQ $0x2, 0xb8(SP)			
  test_int.go:6		0x1093022		48c78424c000000002000000	MOVQ $0x2, 0xc0(SP)			
  test_int.go:6		0x109302e		48890424			MOVQ AX, 0(SP)				
  test_int.go:6		0x1093032		48c744240802000000		MOVQ $0x2, 0x8(SP)			
  test_int.go:6		0x109303b		48c744241002000000		MOVQ $0x2, 0x10(SP)			
  test_int.go:6		0x1093044		e84799ffff			CALL fmt.Println(SB)			
  test_int.go:7		0x1093049		0f57c0				XORPS X0, X0				
  test_int.go:7		0x109304c		0f118424e0000000		MOVUPS X0, 0xe0(SP)			
  test_int.go:7		0x1093054		0f57c0				XORPS X0, X0				
  test_int.go:7		0x1093057		0f118424f0000000		MOVUPS X0, 0xf0(SP)			
  test_int.go:7		0x109305f		488d8424e0000000		LEAQ 0xe0(SP), AX			
  test_int.go:7		0x1093067		4889442448			MOVQ AX, 0x48(SP)			
  test_int.go:7		0x109306c		8400				TESTB AL, 0(AX)				
  test_int.go:7		0x109306e		488d0d2b100100			LEAQ type.*+68320(SB), CX		
  test_int.go:7		0x1093075		48898c24e0000000		MOVQ CX, 0xe0(SP)			
  test_int.go:7		0x109307d		488d0d9cb50400			LEAQ main.statictmp_2(SB), CX		
  test_int.go:7		0x1093084		48898c24e8000000		MOVQ CX, 0xe8(SP)			
  test_int.go:7		0x109308c		8400				TESTB AL, 0(AX)				
  test_int.go:7		0x109308e		488d050b080100			LEAQ type.*+66272(SB), AX		
  test_int.go:7		0x1093095		48898424f0000000		MOVQ AX, 0xf0(SP)			
  test_int.go:7		0x109309d		488d05b4b10400			LEAQ main.statictmp_3(SB), AX		
  test_int.go:7		0x10930a4		48898424f8000000		MOVQ AX, 0xf8(SP)			
  test_int.go:7		0x10930ac		488b442448			MOVQ 0x48(SP), AX			
  test_int.go:7		0x10930b1		8400				TESTB AL, 0(AX)				
  test_int.go:7		0x10930b3		eb00				JMP 0x10930b5				
  test_int.go:7		0x10930b5		4889842498000000		MOVQ AX, 0x98(SP)			
  test_int.go:7		0x10930bd		48c78424a000000002000000	MOVQ $0x2, 0xa0(SP)			
  test_int.go:7		0x10930c9		48c78424a800000002000000	MOVQ $0x2, 0xa8(SP)			
  test_int.go:7		0x10930d5		48890424			MOVQ AX, 0(SP)				
  test_int.go:7		0x10930d9		48c744240802000000		MOVQ $0x2, 0x8(SP)			
  test_int.go:7		0x10930e2		48c744241002000000		MOVQ $0x2, 0x10(SP)			
  test_int.go:7		0x10930eb		e8a098ffff			CALL fmt.Println(SB)			
  test_int.go:8		0x10930f0		0f57c0				XORPS X0, X0				
  test_int.go:8		0x10930f3		0f118424e0000000		MOVUPS X0, 0xe0(SP)			
  test_int.go:8		0x10930fb		0f57c0				XORPS X0, X0				
  test_int.go:8		0x10930fe		0f118424f0000000		MOVUPS X0, 0xf0(SP)			
  test_int.go:8		0x1093106		488d8424e0000000		LEAQ 0xe0(SP), AX			
  test_int.go:8		0x109310e		4889442440			MOVQ AX, 0x40(SP)			
  test_int.go:8		0x1093113		8400				TESTB AL, 0(AX)				
  test_int.go:8		0x1093115		488d0d840f0100			LEAQ type.*+68320(SB), CX		
  test_int.go:8		0x109311c		48898c24e0000000		MOVQ CX, 0xe0(SP)			
  test_int.go:8		0x1093124		488d0d05b50400			LEAQ main.statictmp_4(SB), CX		
  test_int.go:8		0x109312b		48898c24e8000000		MOVQ CX, 0xe8(SP)			
  test_int.go:8		0x1093133		8400				TESTB AL, 0(AX)				
  test_int.go:8		0x1093135		488d0564070100			LEAQ type.*+66272(SB), AX		
  test_int.go:8		0x109313c		48898424f0000000		MOVQ AX, 0xf0(SP)			
  test_int.go:8		0x1093144		488d0515b10400			LEAQ main.statictmp_5(SB), AX		
  test_int.go:8		0x109314b		48898424f8000000		MOVQ AX, 0xf8(SP)			
  test_int.go:8		0x1093153		488b442440			MOVQ 0x40(SP), AX			
  test_int.go:8		0x1093158		8400				TESTB AL, 0(AX)				
  test_int.go:8		0x109315a		eb00				JMP 0x109315c				
  test_int.go:8		0x109315c		4889842480000000		MOVQ AX, 0x80(SP)			
  test_int.go:8		0x1093164		48c784248800000002000000	MOVQ $0x2, 0x88(SP)			
  test_int.go:8		0x1093170		48c784249000000002000000	MOVQ $0x2, 0x90(SP)			
  test_int.go:8		0x109317c		48890424			MOVQ AX, 0(SP)				
  test_int.go:8		0x1093180		48c744240802000000		MOVQ $0x2, 0x8(SP)			
  test_int.go:8		0x1093189		48c744241002000000		MOVQ $0x2, 0x10(SP)			
  test_int.go:8		0x1093192		e8f997ffff			CALL fmt.Println(SB)			
  test_int.go:9		0x1093197		0f57c0				XORPS X0, X0				
  test_int.go:9		0x109319a		0f118424e0000000		MOVUPS X0, 0xe0(SP)			
  test_int.go:9		0x10931a2		0f57c0				XORPS X0, X0				
  test_int.go:9		0x10931a5		0f118424f0000000		MOVUPS X0, 0xf0(SP)			
  test_int.go:9		0x10931ad		488d8424e0000000		LEAQ 0xe0(SP), AX			
  test_int.go:9		0x10931b5		4889442438			MOVQ AX, 0x38(SP)			
  test_int.go:9		0x10931ba		8400				TESTB AL, 0(AX)				
  test_int.go:9		0x10931bc		488d0ddd0e0100			LEAQ type.*+68320(SB), CX		
  test_int.go:9		0x10931c3		48898c24e0000000		MOVQ CX, 0xe0(SP)			
  test_int.go:9		0x10931cb		488d0d6eb40400			LEAQ main.statictmp_6(SB), CX		
  test_int.go:9		0x10931d2		48898c24e8000000		MOVQ CX, 0xe8(SP)			
  test_int.go:9		0x10931da		8400				TESTB AL, 0(AX)				
  test_int.go:9		0x10931dc		488d05bd060100			LEAQ type.*+66272(SB), AX		
  test_int.go:9		0x10931e3		48898424f0000000		MOVQ AX, 0xf0(SP)			
  test_int.go:9		0x10931eb		488d0576b00400			LEAQ main.statictmp_7(SB), AX		
  test_int.go:9		0x10931f2		48898424f8000000		MOVQ AX, 0xf8(SP)			
  test_int.go:9		0x10931fa		488b442438			MOVQ 0x38(SP), AX			
  test_int.go:9		0x10931ff		8400				TESTB AL, 0(AX)				
  test_int.go:9		0x1093201		eb00				JMP 0x1093203				
  test_int.go:9		0x1093203		4889442468			MOVQ AX, 0x68(SP)			
  test_int.go:9		0x1093208		48c744247002000000		MOVQ $0x2, 0x70(SP)			
  test_int.go:9		0x1093211		48c744247802000000		MOVQ $0x2, 0x78(SP)			
  test_int.go:9		0x109321a		48890424			MOVQ AX, 0(SP)				
  test_int.go:9		0x109321e		48c744240802000000		MOVQ $0x2, 0x8(SP)			
  test_int.go:9		0x1093227		48c744241002000000		MOVQ $0x2, 0x10(SP)			
  test_int.go:9		0x1093230		e85b97ffff			CALL fmt.Println(SB)			
  test_int.go:11	0x1093235		48c744243000000000		MOVQ $0x0, 0x30(SP)			
  test_int.go:12	0x109323e		48c7042400000000		MOVQ $0x0, 0(SP)			
  test_int.go:12	0x1093246		e81554f7ff			CALL runtime.convT64(SB)		
  test_int.go:12	0x109324b		488b442408			MOVQ 0x8(SP), AX			
  test_int.go:12	0x1093250		4889442460			MOVQ AX, 0x60(SP)			
  test_int.go:12	0x1093255		0f57c0				XORPS X0, X0				
  test_int.go:12	0x1093258		0f118424e0000000		MOVUPS X0, 0xe0(SP)			
  test_int.go:12	0x1093260		0f57c0				XORPS X0, X0				
  test_int.go:12	0x1093263		0f118424f0000000		MOVUPS X0, 0xf0(SP)			
  test_int.go:12	0x109326b		488d8424e0000000		LEAQ 0xe0(SP), AX			
  test_int.go:12	0x1093273		4889442458			MOVQ AX, 0x58(SP)			
  test_int.go:12	0x1093278		8400				TESTB AL, 0(AX)				
  test_int.go:12	0x109327a		488d0d1f0e0100			LEAQ type.*+68320(SB), CX		
  test_int.go:12	0x1093281		48898c24e0000000		MOVQ CX, 0xe0(SP)			
  test_int.go:12	0x1093289		488d0dc0b30400			LEAQ main.statictmp_8(SB), CX		
  test_int.go:12	0x1093290		48898c24e8000000		MOVQ CX, 0xe8(SP)			
  test_int.go:12	0x1093298		8400				TESTB AL, 0(AX)				
  test_int.go:12	0x109329a		488b442460			MOVQ 0x60(SP), AX			
  test_int.go:12	0x109329f		488d0dba060100			LEAQ type.*+66464(SB), CX		
  test_int.go:12	0x10932a6		48898c24f0000000		MOVQ CX, 0xf0(SP)			
  test_int.go:12	0x10932ae		48898424f8000000		MOVQ AX, 0xf8(SP)			
  test_int.go:12	0x10932b6		488b442458			MOVQ 0x58(SP), AX			
  test_int.go:12	0x10932bb		8400				TESTB AL, 0(AX)				
  test_int.go:12	0x10932bd		eb00				JMP 0x10932bf				
  test_int.go:12	0x10932bf		48898424c8000000		MOVQ AX, 0xc8(SP)			
  test_int.go:12	0x10932c7		48c78424d000000002000000	MOVQ $0x2, 0xd0(SP)			
  test_int.go:12	0x10932d3		48c78424d800000002000000	MOVQ $0x2, 0xd8(SP)			
  test_int.go:12	0x10932df		48890424			MOVQ AX, 0(SP)				
  test_int.go:12	0x10932e3		48c744240802000000		MOVQ $0x2, 0x8(SP)			
  test_int.go:12	0x10932ec		48c744241002000000		MOVQ $0x2, 0x10(SP)			
  test_int.go:12	0x10932f5		e89696ffff			CALL fmt.Println(SB)			
  test_int.go:26	0x10932fa		488bac2400010000		MOVQ 0x100(SP), BP			
  test_int.go:26	0x1093302		4881c408010000			ADDQ $0x108, SP				
  test_int.go:26	0x1093309		c3				RET					
  test_int.go:5		0x109330a		e8a1befbff			CALL runtime.morestack_noctxt(SB)	
  test_int.go:5		0x109330f		e95cfcffff			JMP main.main(SB)			
  :-1			0x1093314		cc				INT $0x3				
  :-1			0x1093315		cc				INT $0x3				
  :-1			0x1093316		cc				INT $0x3				
  :-1			0x1093317		cc				INT $0x3				
  :-1			0x1093318		cc				INT $0x3				
  :-1			0x1093319		cc				INT $0x3				
  :-1			0x109331a		cc				INT $0x3				
  :-1			0x109331b		cc				INT $0x3				
  :-1			0x109331c		cc				INT $0x3				
  :-1			0x109331d		cc				INT $0x3				
  :-1			0x109331e		cc				INT $0x3				
  :-1			0x109331f		cc				INT $0x3				
