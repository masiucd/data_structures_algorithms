package problem_slving;

public class McdStrings {

  public static void main(String[] args) {
    System.out.println(reverseString("Hello"));
  }

  private static String reverseString(String s) {
    String reversed = "";
    for (char ch : s.toCharArray()) {
      reversed = Character.toString(ch) + reversed;
    }
    return reversed;
  }
}
