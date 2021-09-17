import java.util.Scanner;

public class Quadrant{
    public static void main(String[] args) {
    //scanner object
    Scanner sc = new Scanner(System.in);

    int x = sc.nextInt();
    int y = sc.nextInt();
    sc.close();

    int output;

    if(x>0){
        if(y>0){
            output = 1;
        } 
        else {
            output = 4;
        }
    
    }
    else {
        if(y>0){
            output = 2;
        }
        else {
            output = 3;
        }
    }             
    System.out.println(output);

    }
}