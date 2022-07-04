#Ch3 - Basic Data Types
* Go’s types fall into four categories: 
1. Basic types: numbers, strings, and booleans;
2. Aggregate types: arrays and struct;
3. Reference types: pointers, slices, maps, functions, and channels, but what they have
  in common is that they all refer to program variables or state in directly, so that the effect of an
  operation applied to one reference is observed by all copies of that reference;
4. Interface types

***

# 3.1 Integers

* Go provides both signed and unsigned integer arithmetic. Four distinct sizes: 8, 16, 32, and 64 bits of signed integers
  represented by the types **int8**, **int16**, **int32**, and **int64**, and corresponding unsigned versions **uint8**, 
  **uint16**, **uint32**, and **uint64**.
* There are also two types called just **int** and **uint** that are the natural or most efficient size 
  integers on a particular platform; int is by far the most widely used numeric type. Both these types 
  have the same size, either 32 or 64 bits, but that varies by compiler, hardware etc.
* The type **rune** is a synonym for int32 and conventionally indicates that a value is a **Unicode**
  code point. The two names may be used interchangeably. 
* Similarly, the type **byte** is a synonym for **uint8**, and emphasizes that the value is a piece of 
  raw data rather than a small numeric quantity.
* Finally, there is an unsigned integer type **uintptr**, used only for low-level programming.
* **int** is not the same type as **int32**, even if the natural size of integers is 32 bits (an explicit 
  conversion is required)