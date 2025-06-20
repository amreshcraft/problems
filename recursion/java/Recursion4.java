import java.util.Scanner;

public class Recursion4 {

    public static void printFirstNEvenNaturalNumberReverse(int n) {
        // negative number
        if (n <= 1) {

            return;
        }

        if (n == 2) {
            System.out.print(n + ", ");
            return;
        }
        // make even
        if ((n & 1) >= 1) {
            n = n - 1;
        }

        System.out.print(n + ", ");
        printFirstNEvenNaturalNumberReverse(n - 2);

    }

    public static Scanner scan = new Scanner(System.in);

    public static void main(String... args) {
        int n = scan.nextInt();
        printFirstNEvenNaturalNumberReverse(n);
    }
}