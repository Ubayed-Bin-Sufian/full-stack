let firstName: string = "Dylan"; // type string; Explicit type
console.log(typeof firstName);

let secondName = "Dylan"; // Implicit - TypeScript will "guess" the type, based on the assigned value:
console.log(typeof secondName);

// let thirdName: string = "Dylan"; // type string
// thirdName = 33; // attempts to re-assign the value to a different type
// console.log(thirdName);

// let fourthName = "Dylan"; // inferred to type string
// fourthName = 33; // attempts to re-assign the value to a different type
// console.log(fourthName);

// Implicit any as JSON.parse doesn't know what type of data it returns so it can be "any" thing... 
const json = JSON.parse("55");

// Most expect json to be an object, but it can be a string or a number like this example
console.log(typeof json);

// const json1 = JSON.parse("hello"); // SyntaxError: Unexpected token h
// console.log(typeof json1)

const json2 = JSON.parse("\"hello\""); // Outputs: "hello" (string)
console.log(typeof json2)