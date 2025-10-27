package chapter4;

public class StringDemo {

	public static void main(String[] args) {
		
		String s = "     Welcome to Java! ";
		String s1 = "Tahsin is dedicated";
		
		// Display the length & the character at index 20 in the string
		System.out.println(s.length());
		System.out.println(s.charAt(20));
		
		// Concatenate two strings
		System.out.println(s + s1);
		// System.out.println(s.concat(s1));
		
		// Remove leading and trailing spaces
		System.out.println(s.trim());
		System.out.println(s);
		
		
		String t = "Hello";
		String t1 = "Hello";
		String t2 = "hello";
		
		// Compare strings for equality (case sensitive)
		if (t.equals(t1))
			System.out.println("They are the same but case sensitive");
		
		// Compare strings for equality (case insensitive)
		if (t.equalsIgnoreCase(t2))
			System.out.println("They are the same but case insensitive");
		
		// Compare strings lexicographically
		if (t2.compareTo(t) < 0)
			System.out.println(t + "\n" + t2);
		else
			System.out.println(t2 + "\n" + t);
		
	}

}
