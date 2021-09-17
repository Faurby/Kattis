import java.util.Scanner;

public class Dicecup{
    public static void main(String[] args) {
    //scanner object
        Scanner sc = new Scanner(System.in);
        int input1 = sc.nextInt();
        int input2 = sc.nextInt();

        if (input1 == input2){
            input1 += 1;
            System.out.println(input1);
        } else if (input1 > input2){
            for (int i = input2; i <= input1; i++){
                input2 += 1;
                System.out.println(input2);
            }
        } else {
            for (int i = input1; i <= input2; i++){
                input1 += 1;
                System.out.println(input1);
            }
        }
        

    }
}