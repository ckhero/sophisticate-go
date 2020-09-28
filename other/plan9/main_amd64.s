TEXT ·main(SB), $16-0
    MOVQ ·helloworld+0(SB), AX
    MOVQ AX, 0(SP)
    MOVQ ·helloworld+8(SB), R8
    MOVQ R8, 8(SP)
    CALL ·print(SB)
    RET
