import java.util.Scanner;

public class Zamka{
    public static void main(String[] args) {
    //scanner object
    Scanner sc = new Scanner(System.in);

    int n = sc.nextInt();
    int m = sc.nextInt();
    int x = sc.nextInt();
    
    int a = n;
    int b = m;
    int dSum = 0;
    int dSumN = 0;

    sc.close();

    while(dSum != x){
        for(int i = n; i<m; i++){
            while(a>0){
                dSum += a%10;
                a /= 10;
            } a = n + i; 
        }
        
    }

    while(dSumN != x){
        for(int j = n; j<m; j++){
            while(b>0){
                dSumN += b%10;
                b /= 10;
            } b = m + j;
        }
    }
    System.out.println(dSum);
    System.out.println(dSumN);
    }
}