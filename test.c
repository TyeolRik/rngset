#include <stdio.h>

static unsigned long long x;

#define CNG (cng = 6906969069LL * cng + 13579)
#define XS (xs ^= (xs << 13), xs ^= (xs >> 17), xs ^= (xs << 43))

#define TEST (a = (1), a = (2), a = (3))

int main() {
    x = (1ULL << 38) + 10;
    printf("%lld \n", x);
    printf("%d \n", (1 < 3));

    unsigned long long cng = 0;
    printf("%llu \n", CNG);
    printf("%llu \n", cng);

    unsigned long long xs = 5;
    int a = 5;
    printf("%llu \n", XS);
    printf("%d\n", TEST);

    unsigned int k = (1 << 31);
    printf("k = %u\n", k);

    unsigned long long m = (1 << 38);
    printf("m = %llu \n", m);
}