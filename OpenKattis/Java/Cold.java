import java.util.Scanner;

public class Cold{
    public static void main(String[] args) {
    //scanner object

    int sum = 0;

        Scanner sc = new Scanner(System.in);
        int input1 = sc.nextInt();
        for (int i = 0; i < input1; i++){
            int tal = sc.nextInt();
            if (tal < 0){
                sum += 1;
            }
        }
            System.out.println(sum);
    }
}