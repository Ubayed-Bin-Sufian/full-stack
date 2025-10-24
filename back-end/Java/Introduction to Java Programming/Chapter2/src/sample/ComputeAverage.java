package sample;

import java.util.Scanner; // Scanner is in the java.util package

public class ComputeAverage {

	public static void main(String[] args) {
		
		// Create a Scanner object
	    Scanner input = new Scanner(System.in);
	    
	    // Prompt the user to enter three numbers
	    System.out.print("Enter first number: ");
	    double number1 = input.nextDouble();
	    
	    System.out.print("Enter second number: ");
	    double number2 = input.nextDouble();
	    
	    System.out.print("Enter third number: ");
	    double number3 = input.nextDouble();

	    // Compute average
	    double average = (number1 + number2 + number3) / 3;

	    // Display result
	    System.out.println("The average of " + number1 + ", " + number2
	      + " & " + number3 + " is " + average);
	  }

}
