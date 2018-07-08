// Fill screen
.pc = $4000

         ldx #$02
         stx $d021    // set background color
		 ldx #0
 
l1:      txa
         and #$01      // even?
		 bne l3       
		 lda #$01      // letter A
         jmp l4
l3:
		 lda #$81      // Inverse A

l4:      sta $0400,x  
         sta $0500,x 
         sta $0600,x 
         sta $06e8,x 
         txa     		 
         sta $d800,x  
         sta $d900,x
         sta $da00,x
         sta $dae8,x
         inx           
         bne l1
l2:
		 inc $D020
		 jmp l2
                               