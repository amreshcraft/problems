import java.util.Scanner;

public class Recursion1{

public static void printFirstNNaturalNumber(int n ){
    // negative number
    if (n < 1) {
        System.out.println("Number can't be negative");
        return;
    }
    if(n == 1) {
        System.out.print(n+", ");
        return;
    }
    printFirstNNaturalNumber(n - 1);
    System.out.print(n+", ");

}

public static Scanner scan = new Scanner(System.in) ;
public static void main(String ... args){
int n = scan.nextInt();
printFirstNNaturalNumber(n);
}
}