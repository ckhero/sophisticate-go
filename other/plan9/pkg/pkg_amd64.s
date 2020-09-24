#include "textflag.h"

GLOBL ·Id(SB), $8

DATA ·Id(SB)/1,$0xe8
DATA ·Id+1(SB)/1,$0x03
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

GLOBL ·NameData(SB),NOPTR,$8
DATA  ·NameData(SB)/8, $"gopher"

GLOBL ·Name(SB),$24
DATA  ·Name(SB)/8,$·NameData(SB)
DATA  ·Name+8(SB)/8,$6
DATA ·Name+16(SB)/8,$"gopher"


