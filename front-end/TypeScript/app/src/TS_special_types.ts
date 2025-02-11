// Type: any
// let u = true;
// u = "false"; // Error: Type 'string' is not assignable to type 'boolean'.

// u.runANonExistentMethod(); // Error: Property 'runANonExistentMethod' does not exist on type 'boolean'.

// console.log(Math.round(u)); // Error: Argument of type 'boolean' is not assignable to parameter of type 'number'.


// let v: any = true;
// v = "string"; // no error as it can be "any" type

// v.runANonExistentMethod();
// no type error in the editor, but will still throw an error if commented in

// console.log(Math.round(v)); // no error as it can be "any" type


// Type: unknown
// let w: unknown = 1; 
// w = "string"; // no error

// w = { 
//   runANonExistentMethod: () => {
//     console.log("I think therefore I am");
//   } 
// } as { runANonExistentMethod: () => void }

// // How can we avoid the error for the code commented out below when we don't know the type? 
// // w.runANonExistentMethod(); // Error: Object is of type 'unknown'. 

// if(typeof w === 'object' && w !== null) {
//   (w as { runANonExistentMethod: Function }).runANonExistentMethod(); 
// } 
// // Although we have to cast multiple times we can do a check in the if to secure our type and have a safer casting


// Type: never
// let x: never = true; // Error: Type 'boolean' is not assignable to type 'never'.


// Type: undefined and null
let y: undefined = undefined;
console.log(typeof y);

let z: null = null;
console.log(typeof z);