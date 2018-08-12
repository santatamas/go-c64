// Write alphabet
.pc = $4000 

	lda #$01   // letter A
	ldx #0
l1:
	sta $0400,X
	inx
	tay
	iny
	tya
	cpx #26  // Z
	bne l1
	brk
