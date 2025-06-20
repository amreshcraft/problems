import java.util.Scanner;

public class Recursion5 {

    public static void printFirstNOddNaturalNumber(int n) {
        // negative number
        if (n < 1) {
            return;
        }

        if (n == 1) {
            System.out.print(n + ", ");
            return;
        }
        // make Odd
        if (n % 2 == 0) {
            n = n - 1;
        }

        printFirstNOddNaturalNumber(n - 2);
        System.out.print(n + ", ");

    }

    public static Scanner scan = new Scanner(System.in);

    public static void main(String... args) {
        int n = scan.nextInt();
        printFirstNOddNaturalNumber(n);
    }
}