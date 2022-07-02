Ch2 - Arrays
-----------------------

# 1.1 Basics
* Arrays are basic data structures present in almost every programming language,
we use them to store a list of items (list of strings, integers, objects), these items
gets stored SEQUENTIALLY in the memory, because of this looking up items in an array
is fast, the runtime complexity of the LOOKUP of an item is O( 1 ) 
* Limitations: arrays are static (once created we need to specify the size, and it 
cannot be changed)
* If we need to create a bigger array and copy all the items: INSERT the complexity 
will be O( n ), DELETE complexity is O( n ) in worst case scenario when we are deleting 
the first item, because we will have to shift all of them one space to fill the first 
memory address
* 