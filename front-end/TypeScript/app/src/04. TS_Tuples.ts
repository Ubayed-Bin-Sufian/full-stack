// Typed Arrays
// define our tuple
let ourTuple: [number, boolean, string];

// initialize correctly
ourTuple = [5, false, 'Coding God was here'];

// initialize incorrectly throws an error
// ourTuple = [false, 'Coding God was mistaken', 5];
// Type 'boolean' is not assignable to type 'number'.
// Type 'string' is not assignable to type 'boolean'.
// Type 'number' is not assignable to type 'string'.

// We have no type safety in our tuple for indexes 3+
ourTuple.push('Something new and wrong');

console.log(ourTuple);


// Readonly Tuple
// instead use our readonly tuple
const ourReadonlyTuple: readonly [number, boolean, string] = [5, true, 'The Real Coding God'];
// Property 'push' does not exist on type 'readonly [number, boolean, string]'.

// throws error as it is readonly.
// ourReadonlyTuple.push('Coding God took a day off');
// Property 'push' does not exist on type 'readonly [number, boolean, string]'.

// Named Tuples
const graph: [x: number, y: number] = [55.2, 41.3];