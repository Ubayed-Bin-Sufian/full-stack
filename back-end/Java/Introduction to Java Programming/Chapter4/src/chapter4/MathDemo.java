package chapter4;

public class MathDemo {
	
	public static void main (String[] args) {
		
		// Trigonometric methods
		System.out.println(Math.cos(Math.PI / 60));
		System.out.println(Math.toRadians(60));
		System.out.println(Math.toDegrees(0.52));
		
		// Exponential methods
		System.out.println(Math.exp(1));
		System.out.println(Math.log(Math.E));
		System.out.println(Math.pow(2,3));
		System.out.println(Math.sqrt(25));

		// Rounding methods
		System.out.println(Math.ceil(78.1));  // round up
		System.out.println(Math.floor(96.2));  // round down
		// round to nearest integer, for 0.5 even -> round down; odd -> round up
		System.out.println(Math.rint(56.4));    
		System.out.println(Math.sqrt(25));
		
		// Max, min, absolute
		System.out.println(Math.max(2, 10));
		System.out.println(Math.min(-6, 78.2));
		System.out.println(Math.abs(-9.65));  // absolute number
		
		// random method
		System.out.println(Math.random()); // 0.0 - 1.0
		System.out.println((int) (1 + Math.random() * 10)); // 1 - 10
		
	}

}
