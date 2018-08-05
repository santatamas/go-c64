// Test stack operations
.pc = $4000

		lda #3
		pha
		lda #4
		pha
		pla
		cmp #4
		bne error
		pla
		cmp #3
		bne error
		lda #5
		pha
		lda #6
		pha
		tsx				// stack pointer to x
		cpx #$fd
		bne error
		inx
		txs				// x to stack pointer
		pla
		cmp #5
		bne error
		sec
		lda #$c0		// 11000000
		sta $ee
		lda #$40
		bit $ee			// sets V, N flags (7,6 bit is 1)
		bpl error
		bvc error
		sec
		php				// push status register
		clc
		lda #$20
		sta $ee
		bit $ee			// clear V, N flags (7,6 bit is 0)
		bmi error
		bvs error
		plp
		bcc error
		bpl error
		bvc error

		lda #>passed    // put a return address to stack (high)
		pha
		lda #<passed	// (low)
		sec				// need to substract 1
		sbc #1
		pha
		rts
		
error:
		lda #5
		sta $400
		brk
		
passed:
		ldx #$0
l2:		
		lda message,x
		beq l3
		sec
		sbc #$40   // convert from ascii
		sta $0400,x
		inx
		jmp l2
l3:
		brk
.pc	= $4200
message: .text "PASSED"		 