import java.util.Scanner;

public class StarPattern1{
public static void star1(int n ){
	for (int i = 0; i < n ; i++ ) {
		for (int j = 0; j <= i ; j++ ) {
			
			System.out.print("*");
		}
		System.out.println();
		
	}
}

public static void star2(int n ){
	for (int i = 0; i < n ; i++ ) {
		for (int j = 0; j <= n-i-1 ; j++ ) {
			
			System.out.print("*");
		}
		System.out.println();
		
	}
}

public static Scanner sc = new Scanner(System.in);
    public static void main(String[] args) {
        	int t = sc.nextInt();
	sc.nextLine();
	while(t-- > 0){
		int n = sc.nextInt();
		 StarPattern1.star2(n);
	}
    }
}