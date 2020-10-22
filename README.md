# Bell

A programming language based on Lisp

## Basic operators

`(operator expression1 expression2 expressionn)`

### Operators

Example of area calculation

```
(let length 5)
(let width 4)

(let area (* length width))
```

#### Arithmetic operators

|Operator|Description|Example|
|:---:|---|---|
|`+`|Adds expressions|(+ 1 2 3) will evaluate to 6|
|`-`|Subtracts expressions|(- 3 1) will evaluate to 2|
|`*`|Multiplies expressions|(* 3 4) will evaluate to 12|
|`/`|Divides expressions|(/ 6 2) will evaluate to 3|
|`%`|Modulo division (produces a remainder of an integer division)|(% 6 2) will evaluate to 0|

#### Relational operators
|Operator|Description|Example|
|:---:|---|---|
|`=`|Checks equality between expressions (all values have to be equal to evaluate to true)|(= 2 2 3) will evaluate to false|
|`not=`|Checks the difference between expressions (one value has to differ to evaluate to true)|(not= 2 2 3) will evaluate to true|
|`<`|Checks whether the left value is less than the right value|(< 2 3) will evaluate to true|
|`<=`|Checks whether the left value is less than or equal to the right value|(<= 3 3) will evaluate to true|
|`>`|Checks whether the left value is greater than the right value|(> 2 3) will evaluate to false|
|`>=`|Checks whether the left value is greater than or equal to the right value|(>= 3 3) will evaluate to true|

#### Logical operators
|Operator|Description|Example|
|:---:|---|---|
|`and`|Logical "and" operator|(and true true false) will evaluate to false|
|`or`|Logical "or" operator|(or true false false) will evaluate to true|
|`not`|Logical "or" operator. It can have only a single expression. Otherwise, it will return an error.|(not true) will evaluate to false|

#### Types and type rules

Bell supports following types:

- 64-bit signed integers

- booleans - `true` or `false`

Arithmetic operations can only accept numbers. Meaning, following expression:
`(+ 3 true)` will give an error `Operation (+ 3 true) cannot be performed for types: INTEGER and BOOLEAN`.

On the other hand, logical operations can only accept booleans. Expression: `(and true 1)` will give an error
`Operation (and true 1) cannot be performed for types: BOOLEAN and INTEGER`

Relational operations produce boolean values and can be mixed with logical operations.
Expression `(and true (> 4 5))` will evaluate to `false`.

#### Assignment

`(let x 3)`

`let` is a keyword

`x` is a variable identifier

`3` is an assigned value

More complex expressions can be also assigned

`(let x (* 3 6))`