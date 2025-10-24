package sample;

import java.util.Scanner; // Scanner is in the java.util package

public class ComputeBMI {

	public static void main(String[] args) {
		
		// Create a Scanner object
	    Scanner input = new Scanner(System.in);
	    
	    // Prompt the user to enter the weight in pounds
	    System.out.print("Enter weight in pounds: ");
	    double weight = input.nextDouble();
	    
	    // Prompt the user to enter the height in inches
	    System.out.print("Enter height in inches: ");
	    double height = input.nextDouble();
	    
	    // Calculate BMI in pounds and inches
	    double bmi = weight * 0.45359237 / (height * 0.0254 * height * 0.0254);
	    
	    // Display the result
	    System.out.println("BMI is " + bmi);
	}

}
