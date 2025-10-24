package chapter3;

import java.util.Scanner; 

public class InsideCircle {

	public static void main(String[] args) {
		
		// Create a Scanner
	    Scanner input = new Scanner(System.in);

		// Prompt the user to enter a point with two coordinates
		System.out.print("Enter a point with two coordinates: ");
		
		double x = input.nextDouble();
		double y = input.nextDouble();
		
		double distance = Math.pow(x*x + y*y, 0.5);
		// distance = Math.sqrt(x*x + y*y);
		
		if (distance <= 10)  // assume radius is 10 units
			System.out.println("Point is inside the circle");
		else
			System.out.println("Point is outside the circle");
	}

}
