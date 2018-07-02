// subroutine
.pc = $4000 

	ldy #5
l1:
	tya
	jsr mul
	tax
	lda #1
	sta $0400,X
	dey
	bpl l1	
	brk

// multiply A by 40
mul:
	sta $3e
	asl				// *2
	bcs overflow
	asl 			// *2
	bcs overflow
	adc $3e			// +A
	bcs overflow
	asl				// *2
	bcs overflow
	asl				// *2
	bcs overflow
	asl				// *2
	bcs overflow	
	rts
overflow:
	brk
