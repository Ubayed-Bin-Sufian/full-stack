package sample;

public class ShowCurrentTime {
	
	  public static void main(String[] args) {
		
	    // Obtain the total milliseconds since midnight, Jan 1, 1970 (Unix Epoch)
	    long totalMilliseconds = System.currentTimeMillis();
	    System.out.println("Total Milliseconds till now: " + totalMilliseconds);

	    // Obtain the total seconds since midnight, Jan 1, 1970
	    long totalSeconds = totalMilliseconds / 1000;
	    System.out.println("Total seconds till now: " + totalSeconds);

	    // Compute the current second in the minute in the hour
	    long currentSecond = totalSeconds % 60;
	    System.out.println("Total current seconds: " + currentSecond);

	    // Obtain the total minutes
	    long totalMinutes = totalSeconds / 60;

	    // Compute the current minute in the hour
	    long currentMinute = totalMinutes % 60;

	    // Obtain the total hours
	    long totalHours = totalMinutes / 60;

	    // Compute the current hour
	    long currentHour = totalHours % 24;

	    // Display results
	    System.out.println("Current time in Bangladesh is " + currentHour + ":"
	      + currentMinute + ":" + currentSecond + " GMT");
	  }
	  
}
