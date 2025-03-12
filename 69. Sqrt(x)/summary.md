- Overflow

    Checking condition using `mid * mid < x` may result in overflow. Use division instead.

- Zero

    Rule out situations where the divisor may be zero.

- Boundary

    If the loop condition is `l <= r`, return r instead of l, since we need the square root rounded down to the nearest integer.
