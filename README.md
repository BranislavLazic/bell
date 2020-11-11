# Bell

A programming language based on Lisp

## Basic operators

`(operator expression1 expression2 expressionn)`

### Operators

Example of area calculation

```
(let area [length width] (* length width))
(writeln (area 3 4))
```

Reverse word:

```
(let reverse [word]
    (if (> (size word) 0)
        (+ (reverse (tail word)) (head word))))

(writeln (reverse "hello"))
```

Fold left implementation:

```
(let foldL [init lst func]
    (if (not= nil lst)
        (func (head lst) (foldL init (tail lst) func))
        init))

(writeln "Sum:" (foldL 0 (list 1 2 3 4 5) (let _ [a b] (+ a b))))

(writeln "Product:" (foldL 1 (list 1 2 3 4 5) (let _ [a b] (* a b))))
```

Map implementation with invocation which takes a function
that multiplies list members with 2:

```
(let map [lst func]
    (if (not= nil lst)
    (list ((func (head lst)) (map (tail lst) func)))))

(let lst (list 1 2 3 5 6))

(writeln (map lst (let _ [x] (* x 2))))
```

#### Arithmetic operators

| Operator | Description                                                   | Example                      |
| :------: | ------------------------------------------------------------- | ---------------------------- |
|   `+`    | Adds expressions                                              | (+ 1 2 3) will evaluate to 6 |
|   `-`    | Subtracts expressions                                         | (- 3 1) will evaluate to 2   |
|   `*`    | Multiplies expressions                                        | (\* 3 4) will evaluate to 12 |
|   `/`    | Divides expressions                                           | (/ 6 2) will evaluate to 3   |
|   `%`    | Modulo division (produces a remainder of an integer division) | (% 6 2) will evaluate to 0   |
|   `^`    | Value raised to the power of the second argument              | (^ 2 3) will evaluate to 8   |

#### Relational operators

| Operator | Description                                                                             | Example                            |
| :------: | --------------------------------------------------------------------------------------- | ---------------------------------- |
|   `=`    | Checks equality between expressions (all values have to be equal to evaluate to true)   | (= 2 2 3) will evaluate to false   |
|  `not=`  | Checks the difference between expressions (one value has to differ to evaluate to true) | (not= 2 2 3) will evaluate to true |
|   `<`    | Checks whether the left value is less than the right value                              | (< 2 3) will evaluate to true      |
|   `<=`   | Checks whether the left value is less than or equal to the right value                  | (<= 3 3) will evaluate to true     |
|   `>`    | Checks whether the left value is greater than the right value                           | (> 2 3) will evaluate to false     |
|   `>=`   | Checks whether the left value is greater than or equal to the right value               | (>= 3 3) will evaluate to true     |

#### Logical operators

| Operator | Description                                                                                      | Example                                      |
| :------: | ------------------------------------------------------------------------------------------------ | -------------------------------------------- |
|  `and`   | Logical "and" operator                                                                           | (and true true false) will evaluate to false |
|   `or`   | Logical "or" operator                                                                            | (or true false false) will evaluate to true  |
|  `not`   | Logical "or" operator. It can have only a single expression. Otherwise, it will return an error. | (not true) will evaluate to false            |

#### Types and type rules

Bell supports following types:

- 64-bit signed integers

- booleans - `true` or `false`

- nil - a null value

- strings

- lists - a sequence which can contain all previous values

Arithmetic operations can only accept numbers. Meaning, following expression:
`(+ 3 true)` will give an error `Operation (+ 3 true) cannot be performed for types: INTEGER and BOOLEAN`.

On the other hand, logical operations can only accept booleans. Expression: `(and true 1)` will give an error
`Operation (and true 1) cannot be performed for types: BOOLEAN and INTEGER`

Relational operations produce boolean values and can be mixed with logical operations.
Expression `(and true (> 4 5))` will evaluate to `false`.

#### Value assignment

`(let x 3)`

`let` is a keyword

`x` is a variable identifier

`3` is an assigned value

More complex expressions can be also assigned

`(let x (* 3 6))`

#### Conditional expressions

If expression contains three parts.

- Condition - any expression which evaluates to BOOLEAN type.
  Expressions which evaluate to any other type will result in an error.
- Then expression - when the condition is true (4 in the following example)
- Optional, else expression - when the condition is false (3 in the following example)

`(if (> 4 3) 4 3)`

#### Function assignment

Similarly to value assignment, `let` keyword is being used for a function assignment.
Function contains a name, list of parameters (it can be empty) and a body.

`(let double [number] (* number 2))`

To call the function:

`(double 6)` which will produce `12`.

Higher order functions are also supported. Which means that we can pass a function
as an argument:

`(let map [arg func] (func arg))`

Call the function with a function as the argument

`(map 10 (let _ [a] (* a 2)))`

Gives: `20`.

##### Builtin functions

`head` - returns a first element in a list or a string.

Expression `(head (list 1 2 3))` will produce `1`.

Expression `(head "hello")` will give `h`.

`tail` - returns all elements in a list or a string except of the first element.

Expression `(tail (list 1 2 3))` will produce `2 3`.

Expression `(tail "hello")` will produce `ello`.

`size` - returns a size of a list or a length of a string.

Expression `(size (list 1 2 3 4))` will produce `4`.

Expression `(size "hello")` will produce `5`.

#### Open files

Use `open` to import variables and functions from another file relative to bell executable file.

```
(open "stdlib/collection")
```
