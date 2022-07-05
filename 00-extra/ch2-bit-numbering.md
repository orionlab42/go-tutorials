#Ch2 - Bit Numbering

* In computing, bit numbering is the convention used to identify the bit positions in a binary number.

#2.1 Bit significance and indexing

* In computing, the **least significant bit** (LSB) is the bit position in a binary integer representing 
  the **binary 1s place** of the integer. The LSB is sometimes referred to as the **low-order bit** or **right-most bit**, 
  due to the convention in positional notation of writing less significant digits further to the right.
* Similarly, the **most significant bit** (MSB) represents the highest-order place of the binary integer. The MSB is 
  referred to as the **high-order bit** or **left-most bit**. 
* In both cases, the LSB and MSB correlate directly to the least significant digit and most significant digit of 
  a decimal integer.
* Bit indexing correlates to the positional notation of the value in base 2. For this reason, bit index is not affected 
  by how the value is stored on the device, such as the value's byte order. Rather, it is a property of the numeric 
  value in binary itself. This is often utilized in programming via bit shifting: A value of 1 << n corresponds to the 
  nth bit of a binary integer (with a value of 2^n).

#2.2 Unsigned integer example

* This table illustrates an example of decimal value of 149 and the location of LSB. In this particular example, 
* the position of unit value (decimal 1 or 0) is located in bit position 0 (n = 0). MSB stands for most significant 
* bit, while LSB stands for least significant bit.
* Ex. For the decimal number 149: 1 0 0 1 0 1 0 1 b; 
  Bit weight for given bit position n (2^n): 2^7 2^6 2^5 2^4 2^3 2^2 2^1 2^0; where the bit in position 7 is MSB and 
  in position 0 is LSB

#2.3 LSB 0 bit numbering

* When the bit numbering starts at zero for the least significant bit (LSB) the numbering scheme is called LSB 0.
  This bit numbering method has the advantage that for any unsigned number the value of the number can be calculated 
  by using exponentiation with the bit number and a base of 2.

#2.4 MSB 0 bit numbering

* When the bit numbering starts at zero for the most significant bit (MSB) the numbering scheme is called MSB 0.
  