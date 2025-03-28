// Try playing around with modifying properties and adding ones to see what happens
// const car: { type: string, model: string, year: number } = {
//     type: "Toyota",
//     model: "Corolla",
//     year: 2009
//   };


// Type inference
// const car = {
//     type: "Toyota",
//   };
//   car.type = "Ford"; // no error
  // car.type = 2; // Error: Type 'number' is not as


// Example without an optional property
// const car: { type: string, mileage: number } = { // Error: Property 'mileage' is missing in type '{ type: string; }' but required in type '{ type: string; mileage: number; }'.
//   type: "Toyota",
// };
// car.mileage = 2000;


// Example with an optional property
// no error on optional property, remove it and see what happens
// const car: { type: string, mileage?: number } = {
//   type: "Toyota"
// };
// car.mileage = 2000;

// console.log(car);


// Index signatures
const nameAgeMap: { [index: string]: number } = {};
nameAgeMap.Jack = 25; // no error
// nameAgeMap.Mark = "Fifty"; // Error: Type 'string' is not assignable to type 'number'.
console.log(nameAgeMap);