import java.util.Scanner;

public class Taisformula{
    public static void main(String[] args) {
    //scanner object
    Scanner sc = new Scanner(System.in);

    int x1;
    int x2;
    double a;
    double b;
    double sum = 0.0;

    int periode = sc.nextInt();
    
    x1 = sc.nextInt();
        a = sc.nextDouble();
    
    for(int i = 0; i < periode-1; i++){
        

        x2 = sc.nextInt();
        b = sc.nextDouble();
        
        sum += ((a+b)/2)*(x2-x1);

        x1 = x2;
        a = b;
    }
    System.out.println(sum/1000);
    }
}