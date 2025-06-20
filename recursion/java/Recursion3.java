import java.util.Scanner;

public class Recursion3 {

    public static void printFirstNEvenNaturalNumber(int n) {
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

        printFirstNEvenNaturalNumber(n - 2);
        System.out.print(n + ", ");

    }

    public static Scanner scan = new Scanner(System.in);

    public static void main(String... args) {
        int n = scan.nextInt();
        printFirstNEvenNaturalNumber(n);
    }
}