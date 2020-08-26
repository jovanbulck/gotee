TEXT ·do_add(SB),$0-24
    VMOVUPS first_arg+0(FP), X1
    VMOVUPS second_arg+8(FP), X2
    VADDPD X1, X2, X3
    VMOVUPS X3, ret+16(FP)
    
    RET

TEXT ·get_fcw(SB),$0-0
    FSTCW ret+0(FP)
    RET

TEXT ·get_mxcsr(SB),$0-0
    STMXCSR ret+0(FP)
    RET
