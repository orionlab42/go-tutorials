#Ch1 - Binary Numbers

* A binary number is a number expressed in the base-2 numeral system or binary numeral system, a method of mathematical 
  expression which uses only two symbols: typically "0" (zero) and "1" (one).
* The base-2 numeral system is a positional notation with a radix of 2. 
* Each digit is referred to as a **bit**, or **binary digit**.
* Because of its straightforward implementation in digital electronic circuitry using logic gates, the binary system 
  is used by almost all modern computers and computer-based devices.

#1.1 Representation
* Any number can be represented by a sequence of bits (binary digits), which in turn may be represented by any 
  mechanism capable of being in two mutually exclusive states.
* In keeping with customary representation of numerals using Arabic numerals, binary numbers are commonly written 
  using the symbols 0 and 1.
* Common notations: 100101 binary / 100101b / 100101B / bin 100101 / ...

#1.2. Binary counting
* Binary counting follows the exact same procedure, and again the incremental substitution begins with the least 
  significant digit, or bit (the rightmost one, also called the first bit), except that only the two symbols 0 and 1 
  are available. Thus, after a bit reaches 1 in binary, an increment resets it to 0 but also causes an increment of 
  the next bit to the left: 0000, 0001, 0010, 0011, 0100, 0101, 0110, 0111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111 ...
* In the binary system, each bit represents an increasing power of 2, with the rightmost bit representing 2^0, the 
  next representing 2^1, then 2^2, and so on. The value of a binary number is the sum of the powers of 2 represented 
  by each "1" bit.
* Ex.: 
* 1001012 = [ ( 1 ) × 2^5 ] + [ ( 0 ) × 2^4 ] + [ ( 0 ) × 2^3 ] + [ ( 1 ) × 2^2 ] + [ ( 0 ) × 2^1 ] + [ ( 1 ) × 2^0 ]
* 1001012 = [ 1 × 32 ] + [ 0 × 16 ] + [ 0 × 8 ] + [ 1 × 4 ] + [ 0 × 2 ] + [ 1 × 1 ] = 32 + 4 + 1
* 1001012 = 3710

#1.3 Fractions

* Fractions in binary arithmetic terminate only if 2 is the only prime factor in the denominator.
* Ex. As a result, 1/10 does not have a finite binary representation (10 has prime factors 2 and 5). This causes 
  10 × 0.1 not to precisely equal 1 in floating-point arithmetic.
* Ex.1 the binary expression for 1/2 = 0.1 = 1 x 2^-1
* Ex.2 the binary expression for 1/4 = 0.01 = 0 x 2^-1 + 1 x 2^-2

#1.4 Binary arithmetic

* Arithmetic in binary is much like arithmetic in other numeral systems. Addition, subtraction, multiplication, and 
  division can be performed on binary numerals.

#1.4.1 Addition

* The simplest arithmetic operation in binary is addition. Adding two single-digit binary numbers is relatively simple, 
  using a form of carrying:
* 0 + 0 → 0 
* 0 + 1 → 1 
* 1 + 0 → 1 
* 1 + 1 → 0, carry 1 (since 1 + 1 = 2 = 0 + (1 × 2^1))