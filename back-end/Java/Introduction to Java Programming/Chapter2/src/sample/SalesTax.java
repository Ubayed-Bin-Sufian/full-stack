package sample;

import java.util.Scanner;

public class SalesTax {
  public static void main(String[] args) {
    Scanner input = new Scanner(System.in);

    System.out.print("Enter purchase amount: ");
    double purchaseAmount = input.nextDouble();
    
    double tax = purchaseAmount * 0.06;
    System.out.println("Sales tax is " + (int)(tax * 100) / 100.0);  // type casting
    
    // 1. 1300.8
    // 2. 78.04799999999999
    // 3. 7804.799999
    // 4. 7804 (Type casting)
    // 5. 78.04
  }
}