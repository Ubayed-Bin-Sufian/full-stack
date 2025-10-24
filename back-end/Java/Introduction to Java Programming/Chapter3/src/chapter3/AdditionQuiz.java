package chapter3;

import java.util.Scanner;

public class AdditionQuiz {

	public static void main(String[] args) {
		
		int number1 = (int)(System.currentTimeMillis() % 10);  // 10%10=0; 11%10=1; 12%10=2; .... 19%10=9;  
	    int number2 = (int)(System.currentTimeMillis() /6 % 10);

	    // Create a Scanner
	    Scanner input = new Scanner(System.in);

	    System.out.print(
	      "What is " + number1 + " + " + number2 + "? ");

	    int answer = input.nextInt();

	    System.out.println(
	      number1 + " + " + number2 + " = " + answer + " is " +
	      (number1 + number2 == answer));  
	  }

}
