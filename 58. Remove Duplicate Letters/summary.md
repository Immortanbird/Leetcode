We want bigger letters appear after smaller letters.

Based on greedy principle,
- we take in whatever letter that is not in the final result but is bigger than the last letter taken.
- If the current letter is smaller than the previous taken one, gradually throw away previous letters if they appear later in the pattern.

Considering the previous theory, we need to
- use a stack
- record the last occurrence of each letter in the given pattern
- keep track of the singularity of each letter in the final result