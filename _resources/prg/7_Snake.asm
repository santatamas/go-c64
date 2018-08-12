
.pc = $4000

// pseudo random colors
	ldx #0
	lda #$53
l1:
    sta $d800,x  
    sta $d900,x
    sta $da00,x
    sta $dae8,x
	ror
	ror
	adc #$a6
	ror
	eor #$c2
	ror	
	inx
	bne l1
	
// set address of snake positions
	ldx #20
	ldy #10				// snake position (x,y)
	stx $3000			// x coordinates 3000-3003
	sty $3010			// y coordinates 3010-3013
	stx $3001
	sty $3011
	stx $3002
	sty $3012
	stx $3003
	sty $3013

	
//
	lda #0				// 0 = Right Up, 1 = Left Up, 2 = Right Down, 3= Left Down
	sta $30
	
mainloop:

	lda #$20
	ldx #3				// space at the tail
	jsr putchar

	ldx $3000			// save head
	stx $31
	ldy $3010
	sty $32
	
	ldx #2
l6:						// move snake coordinates
	lda $3000,x
	sta $3001,x
	lda $3010,x
	sta $3011,x
	dex
	bpl l6
		
	jsr calculate_next_position	
	ldx $31				// set head to next position
	stx $3000
	ldy $32
	sty $3010
	
	lda #1
	ldx #0
	jsr putchar
	
	lda #0
wait:
	iny
	nop
	nop
	nop
	nop
	bne wait
	inx
	bne wait
	
	jmp mainloop
	
// $30 direction
// $31 x
// $32 y
calculate_next_position:
	ldx $31
	ldy $32
y:
	lda $30
	and #2				// down?
	bne up
down:
	iny
	cpy #24				// bottom of the screen?
	bmi x
	lda $30
	eor #2				// turn up
	sta $30
	jmp x
up:
	dey
	cpy #0
	bne x
	lda $30
	eor #2
	sta $30
x:
	lda $30
	and #1				// left?
	bne right
left:
	dex
	cpx #0
	bne exit
	lda $30
	eor #1				// turn right
	sta $30
	jmp exit
right:
	inx
	cpx #39
	bmi exit
	lda $30
	eor #1
	sta $30

exit:
	stx $31
	sty $32
	rts
	
	

// x - position index
// a - character
putchar:
	sta $40
	lda #$00
	sta $41				// 41, 42, screen topleft ($0400)
	lda #04
	sta $42
	lda $3010,x
	tay					// y
addy:
	beq addx			// add y*40 to 41,42
	lda #40
	clc
	adc $41
	sta $41
	bcc l4
	inc $42
l4:	dey
	jmp addy
	
addx:
	lda $3000,x
	clc
	adc $41
	sta $41
	bcc l5
	inc $42

l5:
	lda $40
	ldy #00
	sta ($41),y
	rts
