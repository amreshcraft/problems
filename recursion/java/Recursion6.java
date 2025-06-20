import java.util.Scanner;

public class Recursion6 {

    public static void printFirstNOddNaturalNumberReverse(int n) {
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

        System.out.print(n + ", ");
        printFirstNOddNaturalNumberReverse(n - 2);

    }

    public static Scanner scan = new Scanner(System.in);

    public static void main(String... args) {
        int n = scan.nextInt();
        printFirstNOddNaturalNumberReverse(n);
    }
}