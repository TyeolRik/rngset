      DOUBLE PRECISION FUNCTION ACORNJ(XDUMMY)
C
C          The additive congruential random number generator—A special case of a multiple recursive generator
C          Roy S.Wikramaratna 
C          https://doi.org/10.1016/j.cam.2007.05.018
C          Page. 383
C          
C          
C          Fortran implementation of ACORN random number generator
C          of order less than or equal to 120 (higher orders can be
C          obtained by increasing the parameter value MAXORD) and
C          modulus less than or equal to 2^60.
C
C          After appropriate initialization of the common block /IACO2/
C          each call to ACORNJ generates a single variate drawn from
C          a uniform distribution over the unit interval.
C
      IMPLICIT DOUBLE PRECISION (A-H,O-Z)
      PARAMETER (MAXORD=120,MAXOP1=MAXORD+1)
      COMMON /IACO2/ KORDEJ,MAXJNT,IXV1(MAXOP1),IXV2(MAXOP1)
      DO 7 I=1,KORDEJ
        IXV1(I+1)=(IXV1(I+1)+IXV1(I))
        IXV2(I+1)=(IXV2(I+1)+IXV2(I))
        IF (IXV2(I+1).GE.MAXJNT) THEN
          IXV2(I+1)=IXV2(I+1)-MAXJNT
          IXV1(I+1)=IXV1(I+1)+1
        ENDIF
      IF (IXV1(I+1).GE.MAXJNT) IXV1(I+1)=IXV1(I+1)-MAXJNT
    7 CONTINUE
      ACORNJ=(DBLE(IXV1(KORDEJ+1)) 
     1          +DBLE(IXV2(KORDEJ+1))/MAXJNT)/MAXJNT
      RETURN
      END