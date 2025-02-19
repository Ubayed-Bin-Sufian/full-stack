// Try playing around with modifying properties and adding ones to see what happens
// const car: { type: string, model: string, year: number } = {
//     type: "Toyota",
//     model: "Corolla",
//     year: 2009
//   };


// Type inference
const car = {
    type: "Toyota",
  };
  car.type = "Ford"; // no error
  // car.type = 2; // Error: Type 'number' is not as

  console.log(car);