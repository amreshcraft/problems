import java.util.Scanner;

public class Recursion2{

public static void printFirstNNaturalNumberReverse(int n ){
    // negative number
    if (n < 1) {
        System.out.println("Number can't be negative");
        return;
    }
    if(n == 1) {
        System.out.print(n+", ");
        return;
    }
    System.out.print(n+", ");
    printFirstNNaturalNumberReverse(n - 1);

}

public static Scanner scan = new Scanner(System.in) ;
public static void main(String ... args){
int n = scan.nextInt();
printFirstNNaturalNumberReverse(n);
}
}