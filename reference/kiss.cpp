// Keep it Simple Stupid, 64-bit MWC version (2011 version)
// Postname : RNGs with periods exceeding 10^(40million).
// https://www.thecodingforums.com/threads/rngs-with-periods-exceeding-10-40million.742134/

#include <stdio.h>
static unsigned long long Q[2097152], carry = 0;

unsigned long long B64MWC(void)
{
    unsigned long long t, x;
    static int j = 2097151;
    j = (j + 1) & 2097151;
    x = Q[j];
    t = (x << 28) + carry;
    carry = (x >> 36) - (t < x);
    return (Q[j] = t - x);
}

#define CNG (cng = 6906969069LL * cng + 13579)
#define XS (xs ^= (xs << 13), xs ^= (xs >> 17), xs ^= (xs << 43))
#define KISS (B64MWC() + CNG + XS)

int main(void)
{
    unsigned long long i, x, cng = 123456789987654321LL,
                             xs = 362436069362436069LL;
    /* First seed Q[] with CNG+XS: */
    for (i = 0; i < 2097152; i++)
        Q[i] = CNG + XS;
    /* Then generate 10^9 B64MWC()s */
    for (i = 0; i < 1000000000; i++)
        x = B64MWC();
    printf("Does x=13596816608992115578 ?\n x=%llu\n", x);
    /* followed by 10^9 KISSes: */
    for (i = 0; i < 1000000000; i++)
        x = KISS;
    printf("Does x=5033346742750153761 ?\n x=%llu\n", x);
    return 0;
}