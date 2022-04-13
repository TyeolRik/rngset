# rngset
This repository is the set of all kinds of Random Generators.

## Implemented

1. Go default PRNG (math/rand)
2. Go default CSPRNG (crypto/rand)
3. [MiddleSquare](https://en.wikipedia.org/wiki/Middle-square_method)
4. [ZX81](https://en.wikipedia.org/wiki/ZX81)
5. [minstd_rand](http://www.cplusplus.com/reference/random/minstd_rand/) from C++
6. [Wichmann-Hill](https://en.wikipedia.org/wiki/Wichmann%E2%80%93Hill)
7. [Rule30](https://en.wikipedia.org/wiki/Rule_30)
8. [Inversive congruential generator](https://en.wikipedia.org/wiki/Inversive_congruential_generator)
9. [ACORN](https://en.wikipedia.org/wiki/ACORN_(PRNG))
10. [Add with Carry](https://www.kurims.kyoto-u.ac.jp/~kyodo/kokyuroku/contents/pdf/0850-07.pdf)
11. [Subtract with Borrow](https://www.kurims.kyoto-u.ac.jp/~kyodo/kokyuroku/contents/pdf/0850-07.pdf)
12. [KISS](https://www.thecodingforums.com/threads/rngs-with-periods-exceeding-10-40million.742134/) 2011 version, 64bit
13. [Mersenne Twister](https://en.wikipedia.org/wiki/Mersenne_Twister) (mt19937)
14. [Well equidistributed long-period linear](https://en.wikipedia.org/wiki/Well_equidistributed_long-period_linear) (512, 1024, 19937)

|       TEST NAME |       SEED |     PASSED |       WEAK |     FAILED |
|-----------------|------------|------------|------------|------------|
|     TogetheRand |      99999 |        113 |          1 |          0 |
|            zx81 |      99999 |          1 |          0 |        113 |
|             awc |      99999 |         15 |          6 |         93 |
|           awc_c |      99999 |         17 |          2 |         95 |
|            swb1 |      99999 |         15 |          6 |         93 |
|            swb2 |      99999 |         17 |          3 |         94 |
|          c_rand |      12345 |         28 |          0 |         86 |
| cpp_minstd_rand |      12345 |         26 |          2 |         86 |
|     cpp_mt19937 |      12345 |        112 |          2 |          0 |
|   go_cryptorand |      99999 |        112 |          2 |          0 |
|     go_mathrand |      99999 |        111 |          3 |          0 |
| kiss_2011_64bit |      12345 |        113 |          1 |          0 |
|  mt19937_64bits |      12345 |        110 |          4 |          0 |
|     WELL512a    |      12345 |        110 |          4 |          0 |
|     WELL1024a   |      12345 |        111 |          3 |          0 |
|    WELL19937a   |      12345 |        108 |          6 |          0 |
|    WELL19937c   |      12345 |        114 |          0 |          0 |
|    Keccak256    |      12345 |        111 |          3 |          0 |

All PRNGs are tested by [Dieharder tests](https://webhome.phy.duke.edu/~rgb/General/dieharder.php) (Actually there is ubuntu package in [here](https://linux.die.net/man/1/dieharder))

## How did I test?

```
$ go run . 12345 | pv | dieharder -a -g 200 >> ./test_results.txt
```

[pv](http://manpages.ubuntu.com/manpages/focal/man1/pv.1.html) shows  the  progress  of  data  through  a pipeline by giving information such as time elapsed, percentage completed (with progress bar), current throughput rate, total data transferred, and ETA.

So, using pipeline, generate all pseudo-random bytes from code, send to pv, and finally send to dieharder for test. Of course, you can just execute (executable compiled) binary file instead of ```go run . 12345```. It's just example.

```-a``` means "All tests". Below is list of tests which are in dieharder tests.

```
$ dieharder -a -l
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
Installed dieharder tests:
 Test Number                         Test Name                Test Reliability
===============================================================================
  -d 0                            Diehard Birthdays Test              Good
  -d 1                               Diehard OPERM5 Test              Good
  -d 2                    Diehard 32x32 Binary Rank Test              Good
  -d 3                      Diehard 6x8 Binary Rank Test              Good
  -d 4                            Diehard Bitstream Test              Good
  -d 5                                      Diehard OPSO           Suspect
  -d 6                                 Diehard OQSO Test           Suspect
  -d 7                                  Diehard DNA Test           Suspect
  -d 8                Diehard Count the 1s (stream) Test              Good
  -d 9                  Diehard Count the 1s Test (byte)              Good
  -d 10                         Diehard Parking Lot Test              Good
  -d 11         Diehard Minimum Distance (2d Circle) Test             Good
  -d 12         Diehard 3d Sphere (Minimum Distance) Test             Good
  -d 13                             Diehard Squeeze Test              Good
  -d 14                                Diehard Sums Test        Do Not Use
  -d 15                                Diehard Runs Test              Good
  -d 16                               Diehard Craps Test              Good
  -d 17                     Marsaglia and Tsang GCD Test              Good
  -d 100                                STS Monobit Test              Good
  -d 101                                   STS Runs Test              Good
  -d 102                   STS Serial Test (Generalized)              Good
  -d 200                       RGB Bit Distribution Test              Good
  -d 201           RGB Generalized Minimum Distance Test              Good
  -d 202                           RGB Permutations Test              Good
  -d 203                             RGB Lagged Sum Test              Good
  -d 204                RGB Kolmogorov-Smirnov Test Test              Good
  -d 205                               Byte Distribution              Good
  -d 206                                         DAB DCT              Good
  -d 207                              DAB Fill Tree Test              Good
  -d 208                            DAB Fill Tree 2 Test              Good
  -d 209                              DAB Monobit 2 Test              Good
```

```-g``` means "Generate number". You can see that there is ```200```, meaning "From binaries using pipeline."

```
$ dieharder -g -l
#=============================================================================#
#            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
#=============================================================================#
#    Id Test Name           | Id Test Name           | Id Test Name           #
#=============================================================================#
|   000 borosh13            |001 cmrg                |002 coveyou             |
|   003 fishman18           |004 fishman20           |005 fishman2x           |
|   006 gfsr4               |007 knuthran            |008 knuthran2           |
|   009 knuthran2002        |010 lecuyer21           |011 minstd              |
|   012 mrg                 |013 mt19937             |014 mt19937_1999        |
|   015 mt19937_1998        |016 r250                |017 ran0                |
|   018 ran1                |019 ran2                |020 ran3                |
|   021 rand                |022 rand48              |023 random128-bsd       |
|   024 random128-glibc2    |025 random128-libc5     |026 random256-bsd       |
|   027 random256-glibc2    |028 random256-libc5     |029 random32-bsd        |
|   030 random32-glibc2     |031 random32-libc5      |032 random64-bsd        |
|   033 random64-glibc2     |034 random64-libc5      |035 random8-bsd         |
|   036 random8-glibc2      |037 random8-libc5       |038 random-bsd          |
|   039 random-glibc2       |040 random-libc5        |041 randu               |
|   042 ranf                |043 ranlux              |044 ranlux389           |
|   045 ranlxd1             |046 ranlxd2             |047 ranlxs0             |
|   048 ranlxs1             |049 ranlxs2             |050 ranmar              |
|   051 slatec              |052 taus                |053 taus2               |
|   054 taus113             |055 transputer          |056 tt800               |
|   057 uni                 |058 uni32               |059 vax                 |
|   060 waterman14          |061 zuf                 |                        |
#=============================================================================#
|   200 stdin_input_raw     |201 file_input_raw      |202 file_input          |
|   203 ca                  |204 uvag                |205 AES_OFB             |
|   206 Threefish_OFB       |207 XOR (supergenerator)|208 kiss                |
|   209 superkiss           |                        |                        |
#=============================================================================#
|   400 R_wichmann_hill     |401 R_marsaglia_multic. |402 R_super_duper       |
|   403 R_mersenne_twister  |404 R_knuth_taocp       |405 R_knuth_taocp2      |
#=============================================================================#
|   500 /dev/random         |501 /dev/urandom        |                        |
#=============================================================================#
#=============================================================================#
```
