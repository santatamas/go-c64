// Test addressing modes
.pc = $4000

		lda #0			// immidiate
		cmp #0
		bne error
		
		lda #1			// zero page
		sta $30
		lda $30
		cmp #1
		bne error
		
		lda #2			// absolute
		sta $5000
		lda $5000
		cmp #2
		bne error
		
		lda #3			// absolute indexed
		ldx #4
		sta $5000,x
		cmp $5004
		bne error
		
		lda #$12		// indexed indirect
		sta $42
		lda #$50
		sta $43
		ldx #3
		lda #4
		sta ($3f,x)
		cmp $5012
		bne error
		
		lda #$FF		//  indirect indexed
		sta $42
		lda #$50
		sta $43
		ldy #3
		lda #5
		sta ($42),y
		cmp $5102
		bne error
		
		lda #6			// zero page indexed
		ldx #4
		sta $50,x
		lda #6
		cmp $54
		bne error
		
		lda #<passed	// low byte of address passed
		sta $5000
		lda #>passed	// high byte of address passed
		sta $5001
		jmp ($5000)     // indirect
		brk
error:
		clc
		adc #$30
		sta $0400
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