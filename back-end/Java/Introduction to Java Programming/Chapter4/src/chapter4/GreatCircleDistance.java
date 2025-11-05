package chapter4;

import java.util.Scanner;

public class GreatCircleDistance {

	public static void main(String[] args) {
		
		Scanner input = new Scanner(System.in);
	    
	    // Prompt the user to enter point 1
	    System.out.print("Enter point 1 in (lat and long) in degrees: ");
	    double x1 = input.nextDouble();
	    double y1 = input.nextDouble();
	    
	    // Prompt the user to enter point 2
	    System.out.print("Enter point 2 in (lat and long) in degrees: ");
	    double x2 = input.nextDouble();
	    double y2 = input.nextDouble();
	    
	    double d = 6371.01 * Math.acos(Math.sin(Math.toRadians(x1)) * Math.sin(Math.toRadians(x2)) 
	    		+ Math.cos(Math.toRadians(x1)) * Math.cos(Math.toRadians(x2)) *
	    		Math.cos(Math.toRadians(y1 - y2)));
	    
	    System.out.print("The distance between the two points is " + d + " km");
	}

}
