For later:
https://en.wikipedia.org/wiki/Signed_number_representations
https://en.wikipedia.org/wiki/Bitwise_operation
https://en.wikipedia.org/wiki/Two%27s_complement


#Ch1 - Binary Numbers and Bitwise operation

## Binary Numbers

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
* This is similar to what happens in decimal when certain single-digit numbers are added together; if the result 
  equals or exceeds the value of the radix (10), the digit to the left is incremented with 1.

#1.4.1.1 Long carry method /  Brookhouse Method of Binary Addition

* This method is generally useful in any binary addition in which one of the numbers contains a long "string" of ones. 
  When given a "string" of digits composed entirely of **n** ones (where n is any integer length), adding 1 will result in 
  the number 1 followed by a string of n zeros. (Ex. 1111 + 1 = 10000)

#1.4.2 Subtraction

* 0 − 0 → 0
* 0 − 1 → 1, borrow 1
* 1 − 0 → 1
* 1 − 1 → 0
* Subtracting a "1" digit from a "0" digit produces the digit "1", while 1 will have to be subtracted from the next 
  column. This is known as borrowing. When the result of a subtraction is less than 0, the procedure is to "borrow" 
  the deficit divided by the radix from the left, subtracting it from the next positional value.
* Subtracting a positive number is equivalent to adding a negative number of equal absolute value. Computers use 
  signed number representations to handle negative numbers—most commonly the two's complement notation. 

#1.4.3 Multiplication

* Same like in decimal: two numbers A and B can be multiplied by partial products: for each digit in B, the product of 
  that digit in A is calculated and written on a new line, shifted leftward so that its rightmost digit lines up with 
  the digit in B that was used. The sum of all these partial products gives the final result.
* In binary there are only two possible outcomes of each partial multiplication: if the digit in B is 0, the partial 
  product is also 0 else the partial product is equal to A.
* Ex. 101 x 100 = 100 + 0 + 10000 = 10100

#1.4.4 Division

* Long division in binary is again similar to its decimal counterpart. Starting from the left we check how many times is 
  the divisor in the dividend first few digits(the number of digits depends on the divisor size), so on...

#1.5 Bitwise operations

* Though not directly related to the numerical interpretation of binary symbols, sequences of bits may be manipulated 
  using Boolean logical operators. When a string of binary symbols is manipulated in this way, it is called a 
  **bitwise operation**; the logical operators **AND**, **OR**, and **XOR** may be performed on corresponding bits 
  in two binary numerals provided as input. The logical **NOT** operation may be performed on individual bits in a 
  single binary numeral provided as input. Sometimes, such operations may be used as arithmetic short-cuts, and may 
  have other computational benefits as well.

#1.6 Conversion to and from other numeral systems
#1.6.2 Decimal to Binary

* To convert from a base-10 integer to its base-2 (binary) equivalent, the number is divided by two. The remainder is 
  the least-significant bit. The quotient is again divided by two; its remainder becomes the next least significant bit. 
  This process repeats until a quotient of one is reached.
* Ex. 2x0 R1 - 2x1 R0 - 2x2 R1 - 2x5 R1 - 2x11 R0 - 2x22 R0 - 2x44 R1 - 2x89 R0 -  2x178 R1 - 357 => 101100101

#1.6.2 Binary to Decimal

* Conversion from base-2 to base-10 simply inverts the preceding algorithm. The bits of the binary number are used one 
  by one, starting with the most significant (leftmost) bit. Beginning with the value 0, the prior value is doubled, 
  and the next bit is then added to produce the next value.
* Ex. 10010101101 = ((((((((((0x2 + 1)x2 + 0)x2 + 0.)x2 + 1.)x2 + 0.)x2 + 1.)x2 + 0.)x2 + 1.)x2 + 1.)x2 + 0.)x2 + 1 
* 1 => 2 => 4 => 9 => 18 => 37 => 74 => 149 => 299 => 598 => 1197

#1.6.3 Hexadecimal

* Binary may be converted to and from hexadecimal more easily. This is because the radix of the hexadecimal system (16) 
  is a power of the radix of the binary system (2). More specifically, 16 = 24, so it takes four digits of binary to 
  represent one digit of hexadecimal. Ex. 1hex = 0001 b, 3 hex = 0011 b, 7 hex = 111 b, A hex = 10 dec = 1010 b, E hex = 14 dec = 1110 b
* To convert a hexadecimal number into its binary equivalent, simply substitute the corresponding binary digits:
  Ex. 3A hex = 58 dec = 0011 10102 b, E7 hex = 231 dec = 1110 01112 b
* To convert a binary number into its hexadecimal equivalent, divide it into groups of four bits. If the number of bits 
  isn't a multiple of four, simply insert extra 0 bits at the left (called padding). For example: 
  1010010 b = 0101 0010 grouped with padding = 52 hex;  11011101 b = 1101 1101 grouped = DD hex
* Binary is also easily converted to the octal numeral system, since octal uses a radix of 8, which is a power of two
  (namely, 2^3, so it takes exactly three binary digits to represent an octal digit).

## Bitwise operation

* In computer programming, a bitwise operation operates on a bit string, a bit array or a binary numeral (considered as a bit string) 
  at the level of its individual bits. It is a fast and simple action, basic to the higher-level arithmetic operations 
  and directly supported by the processor. Most bitwise operations are presented as two-operand instructions where the 
  result replaces one of the input operands.
* On simple low-cost processors, typically, bitwise operations are substantially faster than division, several times 
  faster than multiplication, and sometimes significantly faster than addition. While modern processors usually perform 
  addition and multiplication just as fast as bitwise operations due to their longer instruction pipelines and other 
  architectural design choices, bitwise operations do commonly use less power because of the reduced use of resources.

#1.7 Bitwise operators

* In the explanations below, any indication of a bit's position is counted from the right (least significant) side LSB, 
  advancing left. For example, the binary value 0001 (decimal 1) has zeroes at every position but the first (i.e., the rightmost) one.
* **NOT** - bitwise NOT, or **bitwise complement**, is a unary operation that performs logical negation on each bit, 
  forming the ones' complement of the given binary value. Bits that are 0 become 1, and those that are 1 become 0. 
  For example: NOT 0111 b (decimal 7) = 1000 b (decimal 8)

* **AND** - bitwise AND is a binary operation that takes two equal-length binary representations and performs the 
  logical AND operation on each pair of the corresponding bits. Thus, if both bits in the compared position are 1, the 
  bit in the resulting binary representation is 1 (1 × 1 = 1); otherwise, the result is 0 (1 × 0 = 0 and 0 × 0 = 0). 
  For example: 0101 b (decimal 5) AND 0011 b (decimal 3) = 0001 b (decimal 1)


